package main
import "fmt"

func main() {
  var (
    name string = "Jon"
  )
  switch name {
    case "Jonatan":
      fmt.Println("Hello Jonatan")
    case "Kate":
      fmt.Println("Hello Kate")
    case "Jon":
      fmt.Println("Hello Jon")
    default:
      fmt.Println("No name")
  }

}