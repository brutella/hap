package curve25519

import (
	"crypto/rand"
	"golang.org/x/crypto/curve25519"
)

func GenerateKeyPair() (public, private [32]byte) {
	rand.Read(private[:])
	curve25519.ScalarBaseMult(&public, &private)
	return
}

// SharedSecret returns a Curve25519 shared secret derived from privateKey and otherPublicKey.
func SharedSecret(privateKey, otherPublicKey [32]byte) [32]byte {
	var k [32]byte
	curve25519.ScalarMult(&k, &privateKey, &otherPublicKey)

	return k
}
