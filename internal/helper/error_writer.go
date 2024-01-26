package helper

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	responseData := ResponseData{}

	if errors.Is(ErrNotFound, err) {
		responseData.StatusCode = http.StatusNotFound
		responseData.Message = "data not found"
	} else if errors.Is(ErrLoginFailed, err) || err.Error() == "token has invalid claims: token is expired" || errors.Is(ErrBearerTokenMissing, err) {
		responseData.StatusCode = http.StatusUnauthorized
		responseData.Message = "unauthorized"
	} else if _, ok := err.(validator.ValidationErrors); ok {
		responseData.StatusCode = http.StatusBadRequest
		responseData.Message = "validation error"
	} else {
		responseData.StatusCode = http.StatusInternalServerError
		responseData.Message = "internal server error"
	}

	responseData.Err = err

	WriteResponse(w, responseData)
}
