package main
import (
    "fmt"
)
func main() {

  var x,y int
  fmt.Scan(&x)
  fmt.Scan(&y)

  var max,min,somma,diff,prodotto,rapporto,media int

  if x > y {
    max = x
    min = y
    somma = x + y
    diff = x - y
    prodotto = x * y
    rapporto = x / y
    media = (x + y) / 2
  } else {
    max = y
    min = x
    somma = x + y
    diff = y - x
    prodotto = x * y
    rapporto = y / x
    media = (x + y) / 2
  }


  fmt.Println("il maggiore è : ", max)
  fmt.Println("Il minore è :", min)
  fmt.Println("La somma è: ", somma)
  fmt.Println("La differenza è: ", diff)
  fmt.Println("il prodotto è: ", prodotto)
  fmt.Println("Il rapporto è: ", rapporto)
  fmt.Println("La media è: ", media)



}
