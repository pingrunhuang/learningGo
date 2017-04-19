package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		temp := z
		z = z - (math.Pow(z, 2)-x)/(2*z)
		fmt.Println(z)
		if math.Abs(z-temp) < 0.0001 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
