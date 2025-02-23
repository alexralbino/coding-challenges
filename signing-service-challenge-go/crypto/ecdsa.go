package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// ECCKeyPair is a DTO that holds ECC private and public keys.
type ECCKeyPair struct {
	Public  *ecdsa.PublicKey
	Private *ecdsa.PrivateKey
}

// ECCMarshaler can encode and decode an ECC key pair.
type ECCMarshaler struct{}

// NewECCMarshaler creates a new ECCMarshaler.
func NewECCMarshaler() ECCMarshaler {
	return ECCMarshaler{}
}

// Marshal takes an ECCKeyPair and encodes it to be written on disk.
// It returns the public and the private key as a byte slice.
func (m ECCMarshaler) Marshal(keyPair KeyPair) ([]byte, []byte, error) {
	privateKeyBytes, err := x509.MarshalECPrivateKey(keyPair.Private.(*ecdsa.PrivateKey))
	if err != nil {
		return nil, nil, err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(keyPair.Public.(*ecdsa.PublicKey))
	if err != nil {
		return nil, nil, err
	}

	encodedPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE_KEY",
		Bytes: privateKeyBytes,
	})

	encodedPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC_KEY",
		Bytes: publicKeyBytes,
	})

	return encodedPublic, encodedPrivate, nil
}

// Unmarshal assembles an ECCKeyPair from an encoded private key.
func (m ECCMarshaler) Unmarshal(privateKeyBytes []byte) (*KeyPair, error) {
	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Private: privateKey,
		Public:  &privateKey.PublicKey,
	}, nil
}
