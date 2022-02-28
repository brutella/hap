package hap

import (
	"github.com/brutella/hap/hkdf"
	"github.com/tadglines/go-pkgs/crypto/srp"

	"crypto/sha512"
	"errors"
)

type pairSetupSession struct {
	Identifier    []byte
	Salt          []byte   // s
	PublicKey     []byte   // A
	PrivateKey    []byte   // S
	EncryptionKey [32]byte // K

	session *srp.ServerSession
}

// newPairSetupSession return a new setup server session.
func newPairSetupSession(id, pin string) (*pairSetupSession, error) {
	var err error
	pairName := []byte("Pair-Setup")
	srp, err := srp.NewSRP(srpGroup, sha512.New, keyDerivativeFuncRFC2945(sha512.New, []byte(pairName)))

	if err == nil {
		srp.SaltLength = 16
		salt, v, err := srp.ComputeVerifier([]byte(pin))
		if err == nil {
			session := srp.NewServerSession([]byte(pairName), salt, v)
			pairing := pairSetupSession{
				session:    session,
				Salt:       salt,
				PublicKey:  session.GetB(),
				Identifier: []byte(id),
			}
			return &pairing, nil
		}
	}

	return nil, err
}

// ProofFromClientProof validates client proof (`M1`) and returns authenticator or error if proof is not valid.
func (p *pairSetupSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("client proof is invalid")
	}

	return p.session.ComputeAuthenticator(clientProof), nil
}

// SetupPrivateKeyFromClientPublicKey calculates and internally sets secret key `S` based on client public key `A`
func (p *pairSetupSession) SetupPrivateKeyFromClientPublicKey(key []byte) error {
	key, err := p.session.ComputeKey(key) // S
	if err == nil {
		p.PrivateKey = key
	}

	return err
}

// SetupEncryptionKey calculates and internally sets encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *pairSetupSession) SetupEncryptionKey(salt []byte, info []byte) error {
	hash, err := hkdf.Sha512(p.PrivateKey, salt, info)
	if err == nil {
		p.EncryptionKey = hash
	}

	return err
}

// Main SRP algorithm is described in http://srp.stanford.edu/design.html
// The HAP uses the SRP-6a Stanford implementation with the following characteristics
//      x = H(s | H(I | ":" | P)) -> called the key derivative function
//      M1 = H(H(N) xor H(g), H(I), s, A, B, K)
const (
	srpGroup = "rfc5054.3072" // N (modulo) => 384 byte
)

// keyDerivativeFuncRFC2945 returns the SRP-6a key derivative function which does
//      x = H(s | H(I | ":" | P))
func keyDerivativeFuncRFC2945(h srp.HashFunc, id []byte) srp.KeyDerivationFunc {
	return func(salt, pin []byte) []byte {
		h := h()
		h.Write(id)
		h.Write([]byte(":"))
		h.Write(pin)
		t2 := h.Sum(nil)
		h.Reset()
		h.Write(salt)
		h.Write(t2)
		return h.Sum(nil)
	}
}
