package main

import (
  "fmt"
  "math"
)

func main(){
  marks := [5] float64 {10,7,8,5,9}
  var sum float64 = 0

  for i := 0 ; i < len(marks) ; i++ {
    sum += marks[i]
  }
  var result float64 = sum / float64(len(marks))
  fmt.Println(result)
  fmt.Println(math.Round(result))
}