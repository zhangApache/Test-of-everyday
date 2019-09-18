package main

import "fmt"

func main()  {
	str1 := "acfg"
	str2 :="dsxas"
	for _, v1 := range str1 {
		for _, v2 := range str2{
			if v1 == v2 {
				fmt.Printf(string(v1))
			}
		}
	}
}
