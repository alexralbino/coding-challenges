package validation

import (
	"errors"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
)

// CreateSignatureDeviceInput is the body expected from the CreateSignatureDevice request
type CreateSignatureDeviceInput struct {
	ID        string `json:"id"`
	Algorithm string `json:"algorithm"`
	Label     string `json:"label,omitempty"`
}

// Validate if CreateSignatureDeviceInput is correct
func (c *CreateSignatureDeviceInput) IsValid() error {
	if c.ID == "" || c.Algorithm == "" {
		return errors.New("id and algorithm are required fields")
	}
	if c.Algorithm != "ECC" && c.Algorithm != "RSA" {
		return errors.New("unsupported algorithm")
	}
	return nil
}

// SignTransactionInput is the body expected from the SignTransaction request
type SignTransactionInput struct {
	DeviceID string `json:"device_id"`
	Data     []byte `json:"data"`
}

// Validate if SignTransactionInput is correct
func (s *SignTransactionInput) IsValid() error {
	if s.DeviceID == "" || s.Data == nil {
		return errors.New("id and algorithm are required fields")
	}
	return nil
}

type ListSignatureDeviceInput struct {
	ID        string `json:"id,omitempty"`
	Label     string `json:"label,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
}

type GetSignatureDeviceInput struct {
	ID string
}

type ListTransactionInput struct {
	DeviceID string `json:"device_id,omitempty"`
}

type GetTransactionInput struct {
	ID string
}

// CreateSignatureDeviceOutput handles which data is returned by the API
type CreateSignatureDeviceOutput struct {
	Status string `json:"status"`
}

type ListSignatureDeviceOutput struct {
	Device []*entity.Device `json:"devices"`
}

type GetSignatureDeviceOutput struct {
	Device *entity.Device `json:"device"`
}

// SignTransactionOutput handles which data is returned by the API
type SignTransactionOutput struct {
	Transaction string `json:"signature"`
	SignedData  string `json:"signed_data"`
}

type ListTransactionOutput struct {
	Transaction []*entity.Transaction `json:"transactions"`
}

type GetTransactionOutput struct {
	Transaction *entity.Transaction `json:"transaction"`
}
