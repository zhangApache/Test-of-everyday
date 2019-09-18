package main

import (
	"net/http"
	"log"
	"time"
)

func hello(wr http.ResponseWriter, rs *http.Request)  {
	wr.Write([]byte("hello"))
}
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, rs *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(wr, rs)
		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}
func main() {
	//http.HandleFunc("/", hello)
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	servcer := http.Server{
		Addr: "0.0.0.0:8004",
		Handler: nil,
	}
	err := servcer.ListenAndServe()
	if err != nil{
		log.Println("Error")
	}
}
