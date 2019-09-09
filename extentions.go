package extentions

import (
	"encoding/json"
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// RespondwithJSON help to write response to json format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetQueryParam Help to get query param by key
func GetQueryParam(r *http.Request, key string) string {
	query := r.URL.Query().Get(key)
	return query
}

// IsEmpty validate value is empty
func IsEmpty(str string) bool {
	if len(str) <= 0 || str == "" {
		return true
	}
	return false
}

// IsEqual compare two string
func IsEqual(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

// ValidateStruct help to validate struct
func ValidateStruct(model interface{}) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(model); err != nil {
		return false, err
	}
	return true, nil
}
