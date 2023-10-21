package entity

type Transaction struct {
	ID               string
	DeviceID         string
	SignatureCounter int
	Data             []byte
	LastSignatureID  []byte
}

type Device struct {
	ID               string
	Label            string
	Algorithm        string
	PrivateKey       []byte
	SignatureCounter int
}
