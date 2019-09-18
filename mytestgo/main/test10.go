package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {
}

func (iPhone *ApplePhone) call() {
	fmt.Println("I am Apple Phone, I can call you!")
}

type Test struct {
}

func (test *Test) testcall() {
	fmt.Println("I am test")
}

func main() {
	var phone Phone
	//phone = new(NokiaPhone)
	phone = &NokiaPhone{}
	phone.call()

	phone = new(ApplePhone)
	phone.call()


	/*var testphone TestPhone
	testphone = &Test{}
	testphone.testcall()*/
}