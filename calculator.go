package main

import "fmt"

func main(){
  fmt.Println("Calculator")
  fmt.Println("Select an action +,-,* or /")
  var action string
  fmt.Scan(&action)

  var(
    a float64
    b float64
  )

  fmt.Println("enter the first number")
  fmt.Scan(&a)

  fmt.Println("enter the second number")
  fmt.Scan(&b)

  switch{
    case action == "+":
      fmt.Println(a+b)
    case action == "-":
      fmt.Println(a-b)
    case action == "*":
      fmt.Println(a*b)
    case action == "/":
      fmt.Println(a/b)

  }
}