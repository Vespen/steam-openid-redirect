package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// Error handler structure.
type errorHandler struct {
	wrapped http.Handler
}

// http.Handler.ServeHTTP
func (h *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)

			status := http.StatusInternalServerError

			http.Error(w, http.StatusText(status), status)
		}
	}()

	h.wrapped.ServeHTTP(w, r)
}

// Creates error handler with the provided parameters.
func NewErrorHandler(wrapped http.Handler) http.Handler {
	if wrapped == nil {
		panic("wrapped cannot be nil")
	}

	return &errorHandler{
		wrapped: wrapped,
	}
}
