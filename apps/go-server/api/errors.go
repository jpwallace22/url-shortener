package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type JSONError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func sendJSONError(w http.ResponseWriter, message interface{}, code int) {
	w.WriteHeader(code)
	jsonErr := JSONError{Code: code, Message: message}
	err := json.NewEncoder(w).Encode(jsonErr)
	if err != nil {
		log.Printf("Could not encode error message: %v\n", err)
	}
}

func TranslateErrors(err error) []string {
	formatToSnake := func(f string) string {
		snake := regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(f, "${1}_${2}")
		return strings.ToLower(snake)
	}

	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		var errorMessage string
		field := formatToSnake(err.Field())
		switch err.Tag() {
		case "required":
			errorMessage = fmt.Sprintf("%s is required", field)
		case "min":
			errorMessage = fmt.Sprintf("%s must be at least %s", field, err.Param())
		case "max":
			errorMessage = fmt.Sprintf("%s cannot be more than %s", field, err.Param())
		case "email":
			errorMessage = fmt.Sprintf("%s must be a valid email address", field)
		case "url":
			errorMessage = fmt.Sprintf("%s must be a valid url", field)
		default:
			errorMessage = fmt.Sprintf("%s is not valid", field)
		}
		errors = append(errors, errorMessage)
	}
	return errors
}
