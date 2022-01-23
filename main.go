package main

import (
	"fmt"
	"github.com/babadee08/go-functions/semantics"
	"github.com/babadee08/go-functions/simplemath"
	"math"
	"net/http"
	"strings"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	// uncomment to for all the code for the Second module
	// moduleTwo()

	// uncomment to for all the code for the Third module
	// moduleThree()

	// moduleFourA()
	addExpr := mathExpression(AddExpr)
	println(addExpr(2, 3))

	fmt.Printf("%f\n", double(3, 2, mathExpression(AddExpr)))
	fmt.Printf("%f\n", double(3, 2, mathExpression(MultiplyExpr)))
	fmt.Printf("%f\n", double(3, 2, mathExpression(SubtractExpr)))

	p2 := powerOfTwo()
	value := p2()
	println(value)
	value = p2()
	println(value)
	value = p2()
	println(value)

	var funcs []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	for _, f := range funcs {
		println(f())
	}

}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	/*return func(f float64, f2 float64) float64 {
		return f + f2
	}*/
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
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
