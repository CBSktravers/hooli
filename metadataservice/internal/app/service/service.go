package service

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) GetRow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is the row"))
}

func (h *Handlers) AddRow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Row added to table"))
}

func (h *Handlers) UpdateRow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Row updated"))
}

func (h *Handlers) DeleteRow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("row deleted"))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/GetRow", h.Logger(h.GetRow))
	mux.HandleFunc("/AddRow", h.Logger(h.AddRow))
	mux.HandleFunc("/UpdateRow", h.Logger(h.UpdateRow))
	mux.HandleFunc("/DeleteRow", h.Logger(h.DeleteRow))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
