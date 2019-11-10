package main

import (
	"fmt"
	"math"
)

func pow(x, n, flim float64) float64 {
	if v:= math(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
