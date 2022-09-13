package main

import (
  "fmt"
  "sort"
)

func main(){
  slice := [] int{3,4,7,12,32,5,231,1}
  slice = append(slice,0)
  fmt.Println(slice)
  sort.Ints(slice)
  fmt.Println(slice)

  slice1 := [] string{"h","f","a","b"}
  fmt.Println(slice1)
  sort.Strings(slice1)
  fmt.Println(slice1)
}