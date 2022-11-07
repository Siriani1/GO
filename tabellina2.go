package main

import "fmt"

func tabelllina(numero int) bool{
  var x bool
  if 1 <= numero && numero <= 9 {
    for i := 0; i < 11; i++{
      fmt.Println(numero * i)
      x = true
    }
  }
  return x
}


func main() {

  var numero int

  fmt.Println("Inserisci un numero")
  fmt.Scan(&numero)

  if tabelllina(numero) == false {
    fmt.Println("Numero non valido")
  }
}
