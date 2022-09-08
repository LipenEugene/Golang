package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64
		b float64
		c float64
	)

	fmt.Println("Реши квадратное уравнение")
	fmt.Print("Введи а: ")
	fmt.Scan(&a)
	fmt.Print("Введи b: ")
	fmt.Scan(&b)
	fmt.Print("Введи c: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {

		var (
			x1 float64
			x2 float64
		)

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет 2 корня\nD" + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Printf("Ваше уравнение имеет 1 корень\n D" + fmt.Sprint(D))
		fmt.Println("X: " + fmt.Sprint(x))

	} else if D < 0 {
		fmt.Println("Ваше уравнение не имеет корней \nD: " + fmt.Sprint(D))
	}
}
