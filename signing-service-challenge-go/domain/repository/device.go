package repository

import (
	"errors"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/entity"
)

func (r *repository) CreateSignatureDevice(device *entity.Device) (*entity.Device, error) {

	r.repo.DeviceRWLock.Lock()
	defer r.repo.DeviceRWLock.Unlock()

	// Check if the device with the same ID already exists
	if _, exists := r.repo.Device[device.ID]; exists {
		return nil, errors.New("Device with the same ID already exists")
	}

	r.repo.Device[device.ID] = device
	return device, nil
}

func (r *repository) ListSignatureDevices(id, label, algorithm string) ([]*entity.Device, error) {
	r.repo.DeviceRWLock.RLock()
	defer r.repo.DeviceRWLock.RUnlock()

	var devices []*entity.Device

	for _, device := range r.repo.Device {
		// Filter devices based on the provided criteria
		if (id == "" || device.ID == id) &&
			(label == "" || device.Label == label) &&
			(algorithm == "" || device.Algorithm == algorithm) {
			devices = append(devices, device)
		}
	}

	return devices, nil
}

func (r *repository) GetSignatureDevice(id string) (*entity.Device, error) {
	r.repo.DeviceRWLock.RLock()
	defer r.repo.DeviceRWLock.RUnlock()

	// Retrieve the device by ID
	device, exists := r.repo.Device[id]
	if !exists {
		return nil, errors.New("Device not found")
	}

	return device, nil
}

func (r *repository) SignTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	r.repo.SignatureRWLock.Lock()
	defer r.repo.SignatureRWLock.Unlock()

	// Check if the signature with the same ID already exists
	if _, exists := r.repo.Transaction[transaction.ID]; exists {
		return nil, errors.New("Transaction with the same ID already exists")
	}

	r.repo.Transaction[transaction.ID] = transaction
	r.incrementCounter(transaction.DeviceID)
	return transaction, nil
}

func (r *repository) ListTransactions(deviceID string) ([]*entity.Transaction, error) {
	r.repo.SignatureRWLock.RLock()
	defer r.repo.SignatureRWLock.RUnlock()

	var signatures []*entity.Transaction

	for _, transaction := range r.repo.Transaction {
		// Filter signatures based on the provided criteria
		if deviceID == "" || transaction.DeviceID == deviceID {
			signatures = append(signatures, transaction)
		}
	}

	return signatures, nil
}

func (r *repository) GetTransaction(id string) (*entity.Transaction, error) {
	r.repo.SignatureRWLock.RLock()
	defer r.repo.SignatureRWLock.RUnlock()

	// Retrieve the signature by ID
	signature, exists := r.repo.Transaction[id]
	if !exists {
		return nil, errors.New("Transaction not found")
	}

	return signature, nil
}

func (r *repository) incrementCounter(id string) error {
	// Retrieve the device by ID
	_, exists := r.repo.Device[id]
	if !exists {
		return errors.New("Device not found")
	}
	r.repo.Device[id].SignatureCounter += 1
	return nil
}
