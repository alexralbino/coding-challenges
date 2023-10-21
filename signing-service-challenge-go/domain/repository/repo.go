package repository

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

type repository struct {
	repo *persistence.Database
}

func NewRepository(db *persistence.Database) *repository {
	return &repository{repo: db}
}







type Repository interface {
	CreateSignatureDevice(device *entity.Device) (*entity.Device, error)
	GetSignatureDevice(id string) (*entity.Device, error)
	GetTransaction(id string) (*entity.Transaction, error)
	ListSignatureDevices(id, label, algorithm string) ([]*entity.Device, error)
	ListTransactions(deviceID string) ([]*entity.Transaction, error)
	SignTransaction(signature *entity.Transaction) (*entity.Transaction, error)
}
