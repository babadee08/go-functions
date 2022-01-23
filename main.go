package main

import (
	"fmt"
	"github.com/babadee08/go-functions/simplemath"
)

func main() {
	fmt.Printf("%f\n", simplemath.Add(2, 6))
	fmt.Printf("%f\n", simplemath.Subtract(2, 6))
	fmt.Printf("%f\n", simplemath.Multiply(2, 6))

	answer, err := simplemath.Divide(6, 3)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
	// ignoring the error
	answer, _ = simplemath.Divide(6, 0)
	fmt.Printf("%f\n", answer)

	//total := sum(12.2, 14, 16, 22.4)
	numbers := []float64{12.2, 14, 16, 22.4}
	total := simplemath.Sum(numbers...)
	fmt.Printf("total of sum: %f\n", total)

}
