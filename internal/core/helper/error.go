package helper

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	VALIDATION_ERROR    = "the request has validation errors"
	REQUEST_NOT_FOUND   = "the requested resource was NOT found"
	AUTHORIZATION_ERROR = "you do NOT have adequate permission to access this resource"
	NO_PRINCIPAL        = "principal identifier NOT provided"
	MONGO_DB_ERROR      = "mongodb error"
	NO_RESOURCE_FOUND   = "this resource does not exist"
	NO_RECORD_FOUND     = "sorry. no record found"
	NO_ERRORS_FOUND     = "no errors at the moment"
)

func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)
}

func ErrorArrayToError(errorBody []validator.FieldError) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()

	for _, value := range errorBody {
		body := ErrorBody{Code: VALIDATION_ERROR, Source: Config.AppName, Message: value.Error()}
		// body := ErrorBody{Code: VALIDATION_ERROR, Source: os.Getenv("service_name"), Message: value.Error()}

		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}

func ErrorMessage(code string, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{Code: code, Source: "wallet", Message: message})

	return errorResponse
}
