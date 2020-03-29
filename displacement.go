package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(ao, vo, so float64)  func (float64) float64 {
	fn := func (t float64) float64 {
		return .5 * ao * math.Pow(t, 2) + vo * t + so
	}
	return fn
}

func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
}