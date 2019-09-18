package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main() {
	clint, err := rpc.Dial("tcp","localhost:1234")
	if err != nil{
		log.Fatal("dialing:", err)
	}

	var reply string
	for true  {
		err = clint.Call("HelloService.Hello", "test111", &reply)
		if err != nil{
			log.Fatal("call:", err)
		}

		fmt.Println("data: ",reply)
	}
}