package main

import "fmt"

func primo(n int) bool{
  var x bool
  var div, count int = 2,0
  for div <= n/2 && count < 1{
    if n % div == 0{
      count += 1
    }
    div += 1
  }
  if count==0 {
    x = true
  }else{
    x = false
  }

  return x
}

func NumeriPrimi(limite int) {
  var i int
  fmt.Println("Numeri primi inferiori a", limite)
  for i = 2; i < limite; i++{
    if primo(i) == true{
      fmt.Print(i, " ")
    }
  }
}



func main() {
  var n int

  fmt.Println("Inserisci un numero")
  fmt.Scan(&n)

  if n > 0 {
    NumeriPrimi(n)
  }else{
    fmt.Println("Il numero non è positivo")
  }


}
