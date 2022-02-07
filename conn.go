package hap

import (
	"github.com/brutella/hap/log"

	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net"
)

type Conn struct {
	net.Conn

	// s and ss are used to encrypt data. s is used to temporarily store the session.
	// After the next read, ss becomes s and the session is encrypted from then on.
	// ------------------------------------------------------------------------------------
	// 2022-02-17 (mah) This workaround is needed because switching to encryption is done
	// after sending a response. But Write() on http.ResponseWriter is not immediate.
	// So therefore we wait until the next read.
	s  *Session
	ss *Session

	readBuf io.Reader
}

func (conn *Conn) Upgrade(s *Session) {
	conn.s = s
}

// Write writes bytes to the connection.
// The written bytes are encrypted when possible.
func (conn *Conn) Write(b []byte) (int, error) {
	if conn.ss == nil {
		return conn.Conn.Write(b)
	}

	var buf bytes.Buffer
	buf.Write(b)
	enc, err := conn.ss.Encrypt(&buf)

	if err != nil {
		log.Debug.Println("encryption failed:", err)
		err = conn.Conn.Close()
		return 0, err
	}

	encB, err := ioutil.ReadAll(enc)
	n, err := conn.Conn.Write(encB)

	return n, err
}

const (
	packetSize = 0x400
)

// Read reads bytes from the connection.
// The read bytes are decrypted when possible.
func (conn *Conn) Read(b []byte) (int, error) {
	if conn.s != nil {
		conn.ss = conn.s
		conn.s = nil
	}

	if conn.ss == nil {
		return conn.Conn.Read(b)
	}

	if conn.readBuf == nil {
		r := bufio.NewReader(conn.Conn)
		buf, err := conn.ss.Decrypt(r)
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				// Ignore timeout error #77
			} else if err == net.ErrClosed {
				// Ignore close errors
			} else {
				log.Debug.Println("decryption failed:", err)
				conn.Conn.Close()
			}
			return 0, err
		}

		conn.readBuf = buf
	}

	n, err := conn.readBuf.Read(b)

	if n < len(b) || err == io.EOF {
		conn.readBuf = nil
	}

	return n, err
}
