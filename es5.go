package main
import (
    "fmt"
)
func main() {

  var n int
  fmt.Scan(&n)
  var i int

  for i = 1; i <= n; i++ {
    if i % 3 == 0 && i % 5 == 0{
      fmt.Print("Fizz Buzz", " ")
    } else if i % 5 == 0 {
      fmt.Print("Buzz", " ")
    } else if i % 3 == 0 {
      fmt.Print("Fizz", " ")
    }else{
      fmt.Print(i, " ")
    }
  }

}
