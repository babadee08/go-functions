package main

import (
	"fmt"
	"github.com/babadee08/go-functions/semantics"
	"github.com/babadee08/go-functions/simplemath"
	"net/http"
	"strings"
)

func main() {
	// uncomment to for all the code for the Second module
	// moduleTwo()

	// uncomment to for all the code for the Third module
	// moduleThree()

	// moduleFourA()

}

func moduleFourA() {
	func() {
		println("my first anonymous function")
	}()

	a := func() {
		println("my first anonymous function")
	}
	a()

	b := func(name string) {
		fmt.Printf("my first %s function\n", name)
	}
	b("sample crap")

	c := func(name string) string {
		return fmt.Sprintf("my first %s function", name)
	}
	value := c("sample crap")
	println(value)
}

func moduleThree() {
	sv := semantics.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	println(sv.String())

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "https://plutalsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(request *http.Request) (*http.Response, error) {
	rt.count++
	return nil, nil
}

func moduleTwo() {
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
