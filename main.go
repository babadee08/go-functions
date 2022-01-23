package main

import "fmt"

func main() {
	fmt.Printf("%f\n", add(2, 6))
	fmt.Printf("%f\n", subtract(2, 6))
	fmt.Printf("%f\n", multiply(2, 6))
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func multiply(p1, p2 float64) float64 {
	return p1 * p2
}
