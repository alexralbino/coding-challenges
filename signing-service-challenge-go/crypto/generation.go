package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
)

// RSAGenerator generates a RSA key pair.
type RSAGenerator struct{}

// Generate generates a new KeyPair.
func (g *RSAGenerator) Generate() (*KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

// ECCGenerator generates an ECC key pair.
type ECCGenerator struct{}

// Generate generates a new KeyPair.
func (g *ECCGenerator) Generate() (*KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}
