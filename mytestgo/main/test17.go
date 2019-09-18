package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func echo(wr http.ResponseWriter, r *http.Request)  {
	msg, err := ioutil.ReadAll(r.Body)
	log.Println("test17....")
	if err != nil{
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg){
		log.Println(err, "writen len:", writeLen)
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8001", nil)
	log.Println("test17______")
	if err != nil {
		log.Fatal(err)
	}
}