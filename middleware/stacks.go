package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)


type Request struct {
	ID string
	Date string
}

func ContextMidleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		ctx := context.WithValue(r.Context(),"request",&Request{ID: uuid.New().String(),Date: time.Now().String()})
		next.ServeHTTP(w,r.WithContext(ctx))
	}
}

type warappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *warappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}


func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()

		wrapped := &warappedWriter{
			statusCode: http.StatusOK,
			ResponseWriter: w,
		}
		next.ServeHTTP(wrapped,r)

		log.Println(wrapped.statusCode,r.Context().Value("request"),r.Method,r.URL.Path,time.Since(start))
	}
}