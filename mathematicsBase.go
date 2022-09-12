package main

import (
  "fmt"
  "math"
)

func main(){
  var(
    a int8 = 3
    b int8 = 5
  )

  sum := a + b
  difference := a - b
  multiplication := a * b
  division := a / b

  fmt.Print(sum,"\n",
              difference ,"\n",
              multiplication, "\n",
              division, "\n")

  var (
    c float64 = 32
    d float64 = 5
  )

  divisionV2 := c / d
  ost := int32(c) % int32(d)

  fmt.Print(divisionV2,"\n",
           ost,"\n")

  var number float64 = 144
  result := math.Sqrt(number)
  fmt.Println(result)

  var number1 float64 = 13.132141241
  result1 := math.Ceil(number1)
  fmt.Println(result1)

  var number2 float64 = 13.932141241
  result2 := math.Floor(number2)
  fmt.Println(result2)

  var number3 float64 = 13.59
  result3 := math.Round(number3)
  fmt.Println(result3)

}