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
	Error   error       `json:"-"`
}

func (j *JSONError) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(j.Code)
	if j.Message != nil {
		err := json.NewEncoder(w).Encode(j)
		if err != nil {
			log.Printf("Could not encode error message: %s\n", err)
		}
		return
	}

	if j.Error != nil {
		j.Code = 500
		j.Message = getStatusCodeDescription(j.Code)
		err := json.NewEncoder(w).Encode(j)
		if err != nil {
			log.Printf("Could not encode error message: %s\n", err)
		}
		return
	}

	j.Message = getStatusCodeDescription(j.Code)
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		log.Printf("Could not encode error message: %s\n", err)
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

func getStatusCodeDescription(statusCode int) string {
	switch statusCode {
	case 100:
		return "Continue"
	case 101:
		return "Switching Protocols"
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 202:
		return "Accepted"
	case 203:
		return "Non-Authoritative Information"
	case 204:
		return "No Content"
	case 205:
		return "Reset Content"
	case 206:
		return "Partial Content"
	case 300:
		return "Multiple Choices"
	case 301:
		return "Moved Permanently"
	case 302:
		return "Found"
	case 303:
		return "See Other"
	case 304:
		return "Not Modified"
	case 305:
		return "Use Proxy"
	case 307:
		return "Temporary Redirect"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 402:
		return "Payment Required"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 406:
		return "Not Acceptable"
	case 407:
		return "Proxy Authentication Required"
	case 408:
		return "Request Timeout"
	case 409:
		return "Conflict"
	case 410:
		return "Gone"
	case 411:
		return "Length Required"
	case 412:
		return "Precondition Failed"
	case 413:
		return "Payload Too Large"
	case 414:
		return "URI Too Long"
	case 415:
		return "Unsupported Media Type"
	case 416:
		return "Range Not Satisfiable"
	case 417:
		return "Expectation Failed"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Not Implemented"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"
	case 505:
		return "HTTP Version Not Supported"
	default:
		return "Unknown status code"
	}
}
