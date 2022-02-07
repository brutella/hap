package hkdf

import (
	"crypto/sha512"
	"golang.org/x/crypto/hkdf"
	"io"
)

// Sha512 returns a 256-bit hash from a key, salt and info.
func Sha512(key, salt, info []byte) ([32]byte, error) {
	hash := sha512.New
	hkdf := hkdf.New(hash, key, salt, info)

	var buf [32]byte
	_, err := io.ReadFull(hkdf, buf[:])

	return buf, err
}
