package middleware

import (
	"bytes"
	"wallet/internal/core/helper"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type customWriter struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (writer *customWriter) WriteHeader(status int) {
	writer.status = status
	// writer.ResponseWriter.WriteHeader(status)
}

func (writer *customWriter) Write(b []byte) (int, error) {
	writer.body.Write(b)
	return writer.ResponseWriter.Write(b)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logWriter := &customWriter{
			ResponseWriter: w,
			body:           bytes.NewBufferString(""),
			status:         http.StatusOK,
		}
		statusCode := logWriter.status
		response := helper.NO_ERRORS_FOUND
		level := "INFO"
		if statusCode >= http.StatusBadRequest {
			response = logWriter.body.String()
			level = "ERROR"
		}

		data, err := json.Marshal(&helper.LogStruct{
			Method:        r.Method,
			Level:         level,
			StatusCode:    strconv.Itoa(statusCode),
			Path:          r.URL.String(),
			UserAgent:     r.UserAgent(),
			RemoteIP:      r.RemoteAddr,
			ResponseTime:  time.Since(time.Now()).String(),
			Message:       http.StatusText(statusCode) + ": " + response,
			Version:       "1",
			CorrelationId: uuid.New().String(),
			AppName:       helper.Config.AppName,
			// AppName:         os.Getenv("service_name"),
			ApplicationHost: r.Host,
			LoggerName:      "",
			TimeStamp:       time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("%s\n", data)
		next.ServeHTTP(logWriter, r)

	})
}
