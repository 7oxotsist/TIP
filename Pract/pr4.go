package main

import "fmt"

func main() {
	//fmt.Print(sumOfNumbers(1234))
	//fmt.Println(toC(25))
	//fmt.Println(toF(77))
	arr := [4]int32{1, 2, 3, 4}
	fmt.Print(double(arr))
}

// Сумма цифр
func sumOfNumbers(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + sumOfNumbers(num/10)
}

// Температура из градусов Цельсия в Фаренгейты и обратно
func toC(num float64) float64 {
	return (num*2 - (num * 2 * 0.1)) + 32
}

func toF(num float64) float64 {
	return ((num - 32) / 1.8)
}

// Удвоение массива
func double(a [4]int32) [4]int32 {
	for i := range a {
		a[i] = a[i] * 2
	}
	return a
}
