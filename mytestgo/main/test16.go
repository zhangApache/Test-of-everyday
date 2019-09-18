package main

import "log"

func main() {
	var m map[string]int
	m["one"] = 1 //error
	log.Fatal(m)
}
