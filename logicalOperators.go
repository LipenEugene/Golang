package main

import (
  "fmt"
)

func main(){
  var(
    a int32 = 1
    b int32 = 10
  )
  if a != 2 && b > 5 || a > 6 {
    fmt.Println("True!")
  }
}