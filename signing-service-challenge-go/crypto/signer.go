package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct {
	Device *entity.Device
}

func (r *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	rsaMarshaler := NewRSAMarshaler()
	keyPair, err := rsaMarshaler.Unmarshal(r.Device.PrivateKey)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPSS(
		rand.Reader,
		keyPair.Private.(*rsa.PrivateKey),
		crypto.SHA256,
		dataToBeSigned,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

type ECCSigner struct {
	Device *entity.Device
}

func (e *ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	eccMarshaler := NewECCMarshaler()
	keyPair, err := eccMarshaler.Unmarshal(e.Device.PrivateKey)
	if err != nil {
		return nil, err
	}
	signature, err := ecdsa.SignASN1(
		rand.Reader,
		keyPair.Private.(*ecdsa.PrivateKey),
		dataToBeSigned,
	)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
