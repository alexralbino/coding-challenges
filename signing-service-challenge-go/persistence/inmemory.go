package persistence

import (
	"sync"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
)

type Database struct {
	Device          map[string]*entity.Device
	DeviceRWLock    sync.RWMutex
	Transaction     map[string]*entity.Transaction
	SignatureRWLock sync.RWMutex
}

func NewDatabase() *Database {
	deviceMap := make(map[string]*entity.Device, 0)
	signatureMap := make(map[string]*entity.Transaction, 0)
	return &Database{Device: deviceMap, Transaction: signatureMap}
}
