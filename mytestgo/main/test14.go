package main

import (
	"net/rpc"
	"log"
)

const HelloServiceName  = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main()  {
	client , err := rpc.Dial("tcp","localhost:1122")
	if err != nil{
		log.Fatal("dialing:", err)
	}
	log.Fatal(client)

}
