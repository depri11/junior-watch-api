package v1

import (
	"net/http"
)

func (h *userHandlers) Routes() {
	h.group.Path("/create").Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.CreateUser(w, r)
	})
	h.group.Path("/health").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
}
