package main

import (
	"net/http"
	"log"
)

func hello1(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", hello1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("start server fail", err)
	}
}