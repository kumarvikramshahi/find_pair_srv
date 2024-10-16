package domain

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var Validator *validator.Validate

func ValidatorSingletonClient() {
	Validator = validator.New()
}

func SendHttpResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
	response interface{},
) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(response)
}

func SendHttpError(
	w http.ResponseWriter,
	errEntity *Entity,
) {
	resp := map[string]interface{}{
		"message": errEntity.Message,
		"code":    errEntity.Code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errEntity.HttpStatusCode)
	json.NewEncoder(w).Encode(resp)
}
