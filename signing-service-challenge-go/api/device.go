package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/service"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/validation"
	"github.com/gorilla/mux"
)

func (s *Server) handleCreateSignatureDevice(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &validation.CreateSignatureDeviceInput{}
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			WriteErrorResponse(w, http.StatusInternalServerError, nil)
			return
		}
		if err := json.Unmarshal(body, input); err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		output, err := service.CreateSignatureDevice(input)

		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		WriteAPIResponse(w, http.StatusCreated, output)
	}
}

// handleListSignatureDevices handles the listing of signature devices.
func (s *Server) handleListSignatureDevices(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &validation.ListSignatureDeviceInput{
			ID:        r.URL.Query().Get("id"),
			Label:     r.URL.Query().Get("label"),
			Algorithm: r.URL.Query().Get("algorithm"),
		}
		// Implement the logic to list signature devices
		devices, err := service.ListSignatureDevice(input)
		if err != nil {
			// Handle errors and return a response
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// Write a success response
		WriteAPIResponse(w, http.StatusOK, devices)
	}
}

// handleGetSignatureDevices handles the retrieval of a specific signature device.
func (s *Server) handleGetSignatureDevices(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the device ID from the URL or request parameters
		vars := mux.Vars(r)

		input := &validation.GetSignatureDeviceInput{ID: vars["id"]}

		// Implement the logic to get a specific signature device
		device, err := service.GetSignatureDevice(input)
		if err != nil {
			// Handle errors and return a response
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// Write a success response
		WriteAPIResponse(w, http.StatusOK, device)
	}
}

func (s *Server) handleSignTransaction(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &validation.SignTransactionInput{}
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			WriteErrorResponse(w, http.StatusInternalServerError, nil)
			return
		}
		if err := json.Unmarshal(body, input); err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		output, err := service.SignTransaction(input)

		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		WriteAPIResponse(w, http.StatusOK, output)
	}
}

// handleListTransactions handles the listing of transactions.
func (s *Server) handleListTransactions(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the device ID from the URL or request parameters
		deviceID := r.URL.Query().Get("device_id")

		input := &validation.ListTransactionInput{DeviceID: deviceID}

		// Implement the logic to list transactions for a specific device
		transactions, err := service.ListTransaction(input)
		if err != nil {
			// Handle errors and return a response
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// Write a success response
		WriteAPIResponse(w, http.StatusOK, transactions)
	}
}

// handleGetTransaction handles the retrieval of a specific transaction.
func (s *Server) handleGetTransaction(service service.DeviceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the transaction ID from the URL or request parameters
		vars := mux.Vars(r)

		input := &validation.GetTransactionInput{ID: vars["id"]}

		// Implement the logic to get a specific transaction
		transaction, err := service.GetTransaction(input)
		if err != nil {
			// Handle errors and return a response
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// Write a success response
		WriteAPIResponse(w, http.StatusOK, transaction)
	}
}
