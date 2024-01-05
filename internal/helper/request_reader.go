package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}
