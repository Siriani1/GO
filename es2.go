package main
import (
    "fmt"
)
func main() {

  var n int
  fmt.Scan(&n)

  var a int
  var somma, minimo, massimo, Mzero, minzero, nullo int

  var i int
  for i = 0; i < n; i++ {
    fmt.Scan(&a)
    somma += a
    if i == 0{
      minimo = a
      massimo = a
    }
    if a < minimo {
      minimo = a
    } else if a > massimo {
      massimo = a
    }
    if a > 0 {
      Mzero ++
    } else if a < 0 {
      minzero ++
    } else {
      nullo ++
    }

  }
  fmt.Println(somma, minimo, massimo, Mzero, minzero, nullo)

}
