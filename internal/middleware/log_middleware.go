package middleware

import (
	"go_todo_api/internal/helper"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type LogMiddlewareHandler http.Handler

type LogMiddleware struct {
	handler http.Handler
}

func NewLogMiddleware(handler http.Handler) LogMiddlewareHandler {
	return &LogMiddleware{
		handler: handler,
	}
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.URL.Path
	requestMethod := r.Method

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		helper.WriteErrorResponse(w, err)
	}

	logger.SetOutput(file)

	entry := logger.WithFields(logrus.Fields{
		"request_url":    requestUrl,
		"request_method": requestMethod,
	})

	entry.Info("API Request Occured")

	middleware.handler.ServeHTTP(w, r)
}
