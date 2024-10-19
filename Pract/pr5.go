package main

import (
	"fmt"
	"math"
	"strconv"
)

var a = [...]float64{-3.0, -2.0, -1.0}

func main()  {
	binary("1100", 2, 10)
	sq(2.0, 2.0, 2.0)
	sort(a)
}

func binary(a string, b int, c int)  {
	
	x, err := strconv.ParseInt(a, b, 64)
	fmt.Println(strconv.FormatInt(x, c))
	fmt.Println(err)
}

func sq(a float64, b float64, c float64)  {
	D := (b * b) - 4*a*c
	x1 := (-b + math.Sqrt(D)) / 2
	x2 := (-b - math.Sqrt(D)) / 2
	if (D < 0) {
		x1 = (-b) / 2
		i1 := math.Sqrt(math.Abs(D)) / 2
		x2 = (-b) / 2
		i2 := math.Sqrt(math.Abs(D)) / 2
		fmt.Printf("%f + %fi, %f - %fi", x1, i1, x2, i2)
	} else {
		fmt.Println(x1, x2)
	}
}

func sort(a [])  {
	for i := 0; i < len(a); i++ {
		if math.Abs(a[i]) < math.Abs(a[i+1]) {
			b := a[i]
			a[i] = a[i+1]
			a[i+1] = b
		}
	}
	fmt.Println(a)
}