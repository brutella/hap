package hap

// Pairing is the pairing of a controller with the server.
type Pairing struct {
	Name       string
	PublicKey  []byte
	Permission byte
}
