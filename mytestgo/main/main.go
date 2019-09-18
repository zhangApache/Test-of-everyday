package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	p1 := increment(&count)
	println("p1:\tValue Of[", p1, "]")

	p2 := increment(&count)
	println("p2:\tValue Of[", p2, "]")

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	println(p1 == p2)
}

//go:noinline
func increment(inc *int) *int{
	v := 3
	// Increment the "value of" inc.
	*inc++
	println("inc:\tValue Of[", *inc, "]\tAddr Of[", &inc, "]")
	return &v
}
