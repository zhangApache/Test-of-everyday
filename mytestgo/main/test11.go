package main

import "fmt"

type TestPhone interface {
	Phone
	testcall()
}

type Test111 struct {
}

func (test111 *Test111) call() {
	fmt.Println("I am a test!!!")
}

func main() {
	//var phone1  Phone
	var phone TestPhone
	//phone1 = &Test111{}
	phone.call()
}
