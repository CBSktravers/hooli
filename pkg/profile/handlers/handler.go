package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/CBSktravers/hooli/pkg/profile"
	"github.com/CBSktravers/hooli/pkg/profile/models"
)

// ProfileResource manages endpoints for profile
type Handlers struct {
	logger  *log.Logger
	service profile.Service
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/create", h.Create)
}

func NewHandlers(logger *log.Logger, service profile.Service) *Handlers {
	return &Handlers{
		logger:  logger,
		service: service,
	}
}

// Create Endpoint that creates a profile
func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	// check user credientaion were passed
	keys, ok := r.URL.Query()["UserName"]
	if !ok || len(keys[0]) < 1 {
		// Failed to get user
		err := errors.New("Forbidden: Invlaid credientails")
		DisplayAppError(w, err, "Url Param 'UserName' is missing or account not authorized", 500)
		log.Println("Unathorized account called")
		return
	}
	user := keys[0]
	log.Println("Create profile called by user:", user)
	// autherization of user request has permission

	decoder := json.NewDecoder(r.Body)
	var profile models.Profile
	err := decoder.Decode(&profile)
	log.Printf("%s wants to create new Profile %s", user, profile)

	//handle error better
	if err != nil {
		// return and log response
		DisplayAppError(w, err, "Failed to decode query", 500)
		log.Println("Failed to decode query", err)
		return

	}
	// call create to add profile to database
	err = h.service.Create(profile)
	if err != nil {
		// return and log response
		DisplayAppError(w, err, "Failed add profile into database", 500)
		log.Println("Failed add profile into database", err)
		return

	}
	// Send response back
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile Created"))
}

type appError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HttpStatus int    `json:"status"`
}
type errorResource struct {
	Data appError `json:"data"`
}

// Function return error message
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}
