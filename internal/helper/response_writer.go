package helper

import (
	"encoding/json"
	"go_todo_api/internal/model/response"
	"net/http"
)

type ResponseData struct {
	StatusCode int
	Message    string
	Data       any
	Err        error
}

func WriteResponse(w http.ResponseWriter, responseData ResponseData) error {
	w.WriteHeader(responseData.StatusCode)

	if responseData.StatusCode == 304 || responseData.StatusCode == 204 {
		return nil
	}

	standardResponse := response.StandardResponse{
		Message: responseData.Message,
		Data:    responseData.Data,
	}

	if responseData.Err != nil {
		standardResponse.Data = responseData.Err.Error()
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(standardResponse); err != nil {
		return err
	}

	return nil
}
