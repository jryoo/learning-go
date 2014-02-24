package main

import (
    "fmt"
    "math"
)

func Sqrt(x, delta float64) float64 {
    z := 1.0
    for math.Abs(math.Sqrt(x) - z) > delta {
        z = z - ((math.Pow(z,2) - x)/(2*z))
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2, 0.000000000000000000001))
}