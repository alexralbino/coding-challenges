package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/repository"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain/service"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/gorilla/mux"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors string `json:"errors"`
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		// TODO: add services / further dependencies here ...
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	mux := mux.NewRouter()

	log := log.New(os.Stdout, "[SIGNING CHALLENGE] ", log.LstdFlags)
	db := persistence.NewDatabase()
	repo := repository.NewRepository(db)
	deviceSvc := service.NewDeviceService(log, repo)
	mux.Handle("/api/v0/health", http.HandlerFunc(s.Health)).Methods(http.MethodGet)

	mux.Handle("/api/v0/signature-device", http.HandlerFunc(s.handleCreateSignatureDevice(deviceSvc))).Methods(http.MethodPost)
	mux.Handle("/api/v0/signature-device/list", http.HandlerFunc(s.handleListSignatureDevices(deviceSvc))).Methods(http.MethodGet)
	mux.Handle("/api/v0/signature-device/{id}", http.HandlerFunc(s.handleGetSignatureDevices(deviceSvc))).Methods(http.MethodGet)
	mux.Handle("/api/v0/sign-transaction", http.HandlerFunc(s.handleSignTransaction(deviceSvc))).Methods(http.MethodPost)
	mux.Handle("/api/v0/sign-transaction/list", http.HandlerFunc(s.handleListTransactions(deviceSvc))).Methods(http.MethodGet)
	mux.Handle("/api/v0/sign-transaction/{id}", http.HandlerFunc(s.handleGetTransaction(deviceSvc))).Methods(http.MethodGet)

	return http.ListenAndServe(s.listenAddress, mux)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: err.Error(),
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(errorResponse)
}

// WriteAPIResponse writes an HTTP response with the provided status code and data.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Use the json.NewEncoder to write the data and handle encoding errors
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
