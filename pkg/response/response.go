package response

import "net/http"

// APIResponse is the standard response envelope for all API responses.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Status codes for reference:
var _ = http.StatusOK // 200
// http.StatusCreated      // 201
// http.StatusBadRequest   // 400
// http.StatusUnauthorized // 401
// http.StatusForbidden    // 403
// http.StatusNotFound     // 404
// http.StatusConflict     // 409
// http.StatusInternalServerError // 500
