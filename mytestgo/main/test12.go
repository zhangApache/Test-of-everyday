package main

import (
	"net/rpc"
	"net"
	"fmt"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error{
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil{
		fmt.Printf("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil{
		fmt.Printf("Accept error:", err)
	}
	rpc.ServeConn(conn)
}

