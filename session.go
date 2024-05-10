package hap

import (
	"time"

	"github.com/brutella/hap/chacha20poly1305"
	"github.com/brutella/hap/hkdf"

	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var mux = &sync.Mutex{}
var cons = make(map[string]*conn)

func setConn(addr string, conn *conn) {
	mux.Lock()
	defer mux.Unlock()
	cons[addr] = conn
}

func getConn(req *http.Request) *conn {
	mux.Lock()
	defer mux.Unlock()

	if con, ok := cons[req.RemoteAddr]; !ok {
		return nil
	} else {
		return con
	}
}

func conns() map[string]*conn {
	copy := map[string]*conn{}
	mux.Lock()
	for k, v := range cons {
		copy[k] = v
	}
	mux.Unlock()

	return copy
}

type session struct {
	Pairing Pairing

	encryptKey   [32]byte
	decryptKey   [32]byte
	encryptCount uint64
	decryptCount uint64
	mu           sync.Mutex

	twr *TimedWrite
}

type TimedWrite struct {
	deadline time.Time
	pid      uint64
}

func newSession(shared [32]byte, p Pairing) (*session, error) {
	salt := []byte("Control-Salt")
	out := []byte("Control-Read-Encryption-Key")
	in := []byte("Control-Write-Encryption-Key")

	s := &session{
		Pairing: p,
	}
	var err error
	s.encryptKey, err = hkdf.Sha512(shared[:], salt, out)
	s.encryptCount = 0
	if err != nil {
		return nil, err
	}

	s.decryptKey, err = hkdf.Sha512(shared[:], salt, in)
	s.decryptCount = 0

	return s, err
}

// Encrypt return the encrypted data by splitting it into packets
// [ length (2 bytes)] [ data ] [ auth (16 bytes)]
func (s *session) Encrypt(r io.Reader) (io.Reader, error) {
	packets := packetsFromBytes(r)
	var buf bytes.Buffer
	for _, p := range packets {
		var nonce [8]byte
		s.mu.Lock()
		binary.LittleEndian.PutUint64(nonce[:], s.encryptCount)
		s.encryptCount++
		s.mu.Unlock()

		bLength := make([]byte, 2)
		binary.LittleEndian.PutUint16(bLength, uint16(p.length))

		encrypted, mac, err := chacha20poly1305.EncryptAndSeal(s.encryptKey[:], nonce[:], p.value, bLength[:])
		if err != nil {
			return nil, err
		}

		buf.Write(bLength[:])
		buf.Write(encrypted)
		buf.Write(mac[:])
	}

	return &buf, nil
}

// Decrypt returns the decrypted data
func (s *session) Decrypt(r io.Reader) (io.Reader, error) {
	var buf bytes.Buffer
	for {
		var length uint16
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var b = make([]byte, length)
		if err := binary.Read(r, binary.LittleEndian, &b); err != nil {
			return nil, err
		}

		var mac [16]byte
		if err := binary.Read(r, binary.LittleEndian, &mac); err != nil {
			return nil, err
		}

		var nonce [8]byte
		s.mu.Lock()
		binary.LittleEndian.PutUint64(nonce[:], s.decryptCount)
		s.decryptCount++
		s.mu.Unlock()

		lengthBytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(lengthBytes, uint16(length))

		decrypted, err := chacha20poly1305.DecryptAndVerify(s.decryptKey[:], nonce[:], b, mac, lengthBytes)

		if err != nil {
			return nil, fmt.Errorf("Data encryption failed %s", err)
		}

		buf.Write(decrypted)

		// Finish when all bytes fit in b
		if length < packetLengthMax {
			break
		}
	}

	return &buf, nil
}

const (
	// packetLengthMax is the max length of encrypted packets
	packetLengthMax = 0x400
)

type packet struct {
	length int
	value  []byte
}

// packetsWithSizeFromBytes returns lv (tlv without t(ype)) packets
func packetsWithSizeFromBytes(length int, r io.Reader) []packet {
	var packets []packet
	for {
		var value = make([]byte, length)
		n, err := r.Read(value)
		if n == 0 {
			break
		}

		if n > length {
			panic("Invalid length")
		}

		p := packet{length: n, value: value[:n]}
		packets = append(packets, p)

		if n < length || err == io.EOF {
			break
		}
	}

	return packets
}

// packetsFromBytes returns packets with length packetLengthMax
func packetsFromBytes(r io.Reader) []packet {
	return packetsWithSizeFromBytes(packetLengthMax, r)
}
