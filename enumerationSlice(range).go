package main

import (
  "fmt"
)

func main(){
  slice := [] int {2,3,4,1,6,5}
  fmt.Println("id","number")
  for i, element := range slice{
    fmt.Println(i,"   ",element)
  }

  fmt.Println("____________")
  for _, element := range slice{
    fmt.Println(element)
  }
}