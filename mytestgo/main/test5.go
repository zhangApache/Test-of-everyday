package main

import "fmt"
func main()  {
	a := 100
	b := 200
	b, a = a, b
	fmt.Println(a, b)
}
