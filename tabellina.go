package main

import "fmt"

func tabelllina(numero int) {

  if 1 <= numero && numero <= 9 {
    for i := 0; i < 11; i++{
      fmt.Println(numero * i)

    }
  }
}


func main() {

  var numero int

  fmt.Println("Inserisci un numero")
  fmt.Scan(&numero)

  tabelllina(numero)
}
