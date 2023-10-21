package crypto

// KeyPairGenerator defines a contract for generating key pairs.
type KeyPairGenerator interface {
	Generate() (*KeyPair, error)
}

// KeyPairGenerator defines a contract for generating key pairs.
type KeyPairMarshaler interface {
	Marshal(keyPair KeyPair) ([]byte, []byte, error)
	Unmarshal(privateKeyBytes []byte) (*KeyPair, error)
}

type KeyPair struct {
	Public  any
	Private any
}
