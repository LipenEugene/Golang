package main

import (
  "fmt"
)

func main(){
  var names [3] string

  names[0] = "Jon"
  names[1] = "Dan"
  names[2] = "Kate"
  fmt.Println(names)

  names1 := [3] string {"Jon", "Dan", "Kate"}
  fmt.Println(names1,len(names))

  for i := 0; i<len(names);i++{
    fmt.Println(names[i])
  }
}