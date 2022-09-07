package main

import "fmt"

func main() {
	var age int8 = 23
	var number float64 = 254321.321453
	fmt.Println(age)
	fmt.Println(number)

	lastName := "Lipen"
	var firstName string = "Eugene"
	fmt.Println(firstName, lastName)
	fmt.Println(len(lastName))

	var name string

	fmt.Println("What is you name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!!!")

	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years!")
}
