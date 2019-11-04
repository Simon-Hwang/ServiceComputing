package service

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler http.Handler, content string) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		handler.ServeHTTP(w,r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			content,
			time.Since(start),
		)//output connect info 
	})
}