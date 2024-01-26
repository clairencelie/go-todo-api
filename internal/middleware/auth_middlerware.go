package middleware

import (
	"errors"
	"go_todo_api/internal/helper"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authorizationHeader := r.Header.Get("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			helper.WriteErrorResponse(w, errors.New("bearer token is missing"))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		_, err := helper.ValidateJWT(tokenString)

		if err != nil {
			helper.WriteErrorResponse(w, err)
			return
		}

		next(w, r, params)
	}
}
