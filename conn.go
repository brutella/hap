package hap

import (
	"sync"

	"github.com/brutella/hap/log"

	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
)

type conn struct {
	net.Conn

	// s and ss are used to encrypt data. s is used to temporarily store the session.
	// After the next read, ss becomes s and the session is encrypted from then on.
	// ------------------------------------------------------------------------------------
	// 2022-02-17 (mah) This workaround is needed because switching to encryption is done
	// after sending a response. But Write() on http.ResponseWriter is not immediate.
	// So therefore we wait until the next read.
	s   *session
	smu sync.Mutex
	ss  *session

	readBuf io.Reader
}

func newConn(c net.Conn) *conn {
	return &conn{
		Conn: c,
		smu:  sync.Mutex{},
	}
}

func (c *conn) Upgrade(s *session) {
	c.smu.Lock()
	c.s = s
	c.smu.Unlock()
}

// Write writes bytes to the connection.
// The written bytes are encrypted when possible.
func (c *conn) Write(b []byte) (int, error) {
	if c.ss == nil {
		return c.Conn.Write(b)
	}

	var buf bytes.Buffer
	buf.Write(b)
	enc, err := c.ss.Encrypt(&buf)

	if err != nil {
		log.Debug.Println("encryption failed:", err)
		err = c.Conn.Close()
		return 0, err
	}

	encB, err := ioutil.ReadAll(enc)
	if err != nil {
		return 0, err
	}
	_, err = c.Conn.Write(encB)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}

const (
	packetSize = 0x400
)

// Read reads bytes from the connection.
// The read bytes are decrypted when possible.
func (c *conn) Read(b []byte) (int, error) {
	c.smu.Lock()
	if c.s != nil {
		c.ss = c.s
		c.s = nil
	}
	c.smu.Unlock()

	if c.ss == nil {
		return c.Conn.Read(b)
	}

	if c.readBuf == nil {
		r := bufio.NewReader(c.Conn)
		buf, err := c.ss.Decrypt(r)
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				// Ignore timeout error #77
			} else if errors.Is(err, net.ErrClosed) {
				// Ignore close errors
			} else {
				log.Debug.Println("decryption failed:", err)
				c.Conn.Close()
			}
			return 0, err
		}

		c.readBuf = buf
	}

	n, err := c.readBuf.Read(b)

	if n < len(b) || err == io.EOF {
		c.readBuf = nil
	}

	return n, err
}
