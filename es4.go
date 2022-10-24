package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {

  rand.Seed(int64(time.Now().Nanosecond()))
  var numeroGenerato int = rand.Intn(101)
  var n int
  var i int = 1


  for {
    fmt.Print("Tentativo n°", i, ": ")
    fmt.Scan(&n)

    if n == numeroGenerato {
      fmt.Println("Hai indovinato in", i, "tentativi")
      break
    }else if n > numeroGenerato {
      fmt.Println("Troppo alto! Riprova!")
    }else {
      fmt.Println("Troppo basso! Riprova!")
    }
    i++
  }

}
