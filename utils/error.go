package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
)

func RespondInternalServerError(err error, msg string, w http.ResponseWriter, r *http.Request) {
	if err == nil {
		return
	}

	log := logger.Get(r.Context())
	log.WithError(err).Error(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if err = json.NewEncoder(w).Encode(map[string]string{"message": msg}); err != nil {
		log.WithError(err).Error("Failed to encode JSON")
	}
}
