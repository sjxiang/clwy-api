package handlers

import (
	"encoding/json"
	"net/http"
)


func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "OK",
		"env":    "dev",
	}

	h.logger.Infow("Health check passed")	
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
