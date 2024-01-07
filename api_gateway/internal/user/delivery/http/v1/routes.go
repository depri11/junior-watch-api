package v1

import (
	"net/http"
)

func (h *userHandlers) Routes() {
	h.group.HandleFunc("/create", h.CreateUser).Methods("POST")
	h.group.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	}).Methods("GET")
}
