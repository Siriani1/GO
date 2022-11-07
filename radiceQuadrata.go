package main

import (
  "fmt";
  "math"
)

func RadiceQuadrata(numero float64) (float64, bool){
  var x bool
  var radiceQuadrata float64
  if numero > 0 {
    radiceQuadrata = math.Sqrt(numero)
    x = true
  }else{
    radiceQuadrata = 0
  }
  return radiceQuadrata, x
}

func main() {
  var numero,radice float64
  var x bool

  fmt.Println("Inserisci un numero")
  fmt.Scan(&numero)

  radice,x = RadiceQuadrata(numero)

  if x == true{
    fmt.Println("Radice quadrata: ", radice)
  }else{
    fmt.Println("Il valore inserito non appartiene al dominio della funzione")
  }

}
