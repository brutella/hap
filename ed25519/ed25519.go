package ed25519

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
)

// ValidateSignature return true when the  signature
// is a valid signature of data based on the provided
// key. Otherwise false is returned.
func ValidateSignature(key, data, signature []byte) bool {
	if len(key) != ed25519.PublicKeySize || len(signature) != ed25519.SignatureSize {
		return false
	}

	return ed25519.Verify(ed25519.PublicKey(key), data, signature)
}

// Signature returns the signature of data using the provided key.
func Signature(key, data []byte) ([]byte, error) {
	if len(key) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("Invalid size of key (%v)", len(key))
	}

	signature := ed25519.Sign(ed25519.PrivateKey(key), data)

	return signature[:], nil
}

// GenerateKey return a public and private key pair from the provided str.
func GenerateKey(str string) (public [32]byte, private [64]byte, err error) {
	b := bytes.NewBuffer([]byte(str))
	if len(str) < 32 {
		zeros := make([]byte, 32-len(str))
		b.Write(zeros)
	}

	pub, priv, err := ed25519.GenerateKey(bytes.NewReader(b.Bytes()))

	copy(public[:], pub)
	copy(private[:], priv)

	return public, private, err
}
