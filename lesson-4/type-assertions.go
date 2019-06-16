package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	s, ok = i.(float64)
	fmt.Println(s, ok)

	s = i.(float64) // Panic
	fmt.Println(s, ok)

}
