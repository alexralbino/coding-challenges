package service

import (
	"fmt"
	"log"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/repository"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/validation"
	"github.com/google/uuid"
)

type DeviceService interface {
	CreateSignatureDevice(input *validation.CreateSignatureDeviceInput) (*validation.CreateSignatureDeviceOutput, error)
	ListSignatureDevice(input *validation.ListSignatureDeviceInput) (*validation.ListSignatureDeviceOutput, error)
	GetSignatureDevice(input *validation.GetSignatureDeviceInput) (*validation.GetSignatureDeviceOutput, error)
	SignTransaction(input *validation.SignTransactionInput) (*validation.SignTransactionOutput, error)
	ListTransaction(input *validation.ListTransactionInput) (*validation.ListTransactionOutput, error)
	GetTransaction(input *validation.GetTransactionInput) (*validation.GetTransactionOutput, error)
}

type deviceService struct {
	repo   repository.Repository
	logger *log.Logger
}

func NewDeviceService(logger *log.Logger, repo repository.Repository) DeviceService {
	return &deviceService{
		logger: logger,
		repo:   repo,
	}
}

func (d *deviceService) CreateSignatureDevice(input *validation.CreateSignatureDeviceInput) (*validation.CreateSignatureDeviceOutput, error) {
	if err := input.IsValid(); err != nil {
		return nil, err
	}

	device := &entity.Device{
		ID:        input.ID,
		Label:     input.Label,
		Algorithm: input.Algorithm,
	}

	var generator crypto.KeyPairGenerator
	var marshaler crypto.KeyPairMarshaler

	switch input.Algorithm {
	case "ECC":
		generator = &crypto.ECCGenerator{}
		marshaler = crypto.NewECCMarshaler()
	case "RSA":
		generator = &crypto.RSAGenerator{}
		marshaler = crypto.NewRSAMarshaler()
	}

	keys, err := generator.Generate()
	if err != nil {
		return nil, err
	}
	_, device.PrivateKey, err = marshaler.Marshal(*keys)
	if err != nil {
		return nil, err
	}

	_, err = d.repo.CreateSignatureDevice(device)
	if err != nil {
		return nil, err
	}

	return &validation.CreateSignatureDeviceOutput{Status: "Device Created"}, nil
}

func (d *deviceService) ListSignatureDevice(input *validation.ListSignatureDeviceInput) (*validation.ListSignatureDeviceOutput, error) {
	devices, err := d.repo.ListSignatureDevices(input.ID, input.Label, input.Algorithm)
	if err != nil {
		return nil, err
	}
	return &validation.ListSignatureDeviceOutput{Device: devices}, nil
}

func (d *deviceService) GetSignatureDevice(input *validation.GetSignatureDeviceInput) (*validation.GetSignatureDeviceOutput, error) {
	device, err := d.repo.GetSignatureDevice(input.ID)
	if err != nil {
		return nil, err
	}
	return &validation.GetSignatureDeviceOutput{Device: device}, nil
}

func (d *deviceService) SignTransaction(input *validation.SignTransactionInput) (*validation.SignTransactionOutput, error) {
	if err := input.IsValid(); err != nil {
		return nil, err
	}

	device, err := d.repo.GetSignatureDevice(input.DeviceID)
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(device)
	if err != nil {
		return nil, err
	}

	signature, err := signer.Sign(input.Data)
	if err != nil {
		return nil, err
	}

	lastSignature, err := signer.Sign([]byte(input.DeviceID))
	if err != nil {
		return nil, err
	}

	if device.SignatureCounter > 0 {
		var max int
		transactions, err := d.repo.ListTransactions(input.DeviceID)
		if err != nil {
			return nil, err
		}
		for _, transaction := range transactions {
			if transaction.SignatureCounter > max {
				lastSignature = transaction.LastSignatureID
			}
		}
	}

	signatureEntity := &entity.Transaction{
		ID:               uuid.New().String(),
		DeviceID:         input.DeviceID,
		SignatureCounter: device.SignatureCounter,
		Data:             input.Data,
		LastSignatureID:  lastSignature,
	}

	_, err = d.repo.SignTransaction(signatureEntity)
	if err != nil {
		return nil, err
	}

	return &validation.SignTransactionOutput{
		Transaction: string(signature),
		SignedData:  fmt.Sprintf("%d_%s_%s", device.SignatureCounter, input.Data, lastSignature),
	}, nil
}

func (d *deviceService) ListTransaction(input *validation.ListTransactionInput) (*validation.ListTransactionOutput, error) {
	transactions, err := d.repo.ListTransactions(input.DeviceID)
	if err != nil {
		return nil, err
	}
	return &validation.ListTransactionOutput{Transaction: transactions}, nil
}

func (d *deviceService) GetTransaction(input *validation.GetTransactionInput) (*validation.GetTransactionOutput, error) {
	transaction, err := d.repo.GetTransaction(input.ID)
	if err != nil {
		return nil, err
	}
	return &validation.GetTransactionOutput{Transaction: transaction}, nil
}

func getSigner(device *entity.Device) (crypto.Signer, error) {
	switch device.Algorithm {
	case "ECC":
		return &crypto.ECCSigner{Device: device}, nil
	case "RSA":
		return &crypto.RSASigner{Device: device}, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", device.Algorithm)
	}
}
