package utils

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)
var Validator *validator.Validate

func init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
}
func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}


func WriteJsonResponse(w http.ResponseWriter , status int , data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.WriteHeader(status)


	return json.NewEncoder(w).Encode(data)

}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	}
	return WriteJsonResponse(w, status, response)
}
func WriteJsonErrorResponse(w http.ResponseWriter, status int ,errorMessage string,err error ) error {
	response := map[string]string{
		"status": "error",
		"message": errorMessage,
		"error": err.Error(),
	}
	return WriteJsonResponse(w, status, response)
}

func ReadJsonBody(r *http.Request, result any) error {

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields() //Prevent unknown fields from being sent in the request
	return decoder.Decode(result)


}