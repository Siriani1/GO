package main
import (
    "fmt"
)
func main() {

  var n,i int

  fmt.Scan(&n)

  fmt.Print("i divisori di ", n, ": ")

  for i = 1; i < n; i++{
    if n % i == 0{
      fmt.Print(i, " ")
    }
  }
}
