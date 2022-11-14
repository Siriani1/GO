package main

import (
  "fmt";
  "math"
)

func Troncamento(a float64, n int){
  x := a * math.Pow(10, float64(n))
  b := math.Trunc(x)
  e := b / math.Pow(10, float64(n))

  fmt.Println("Il valore troncato è:", e)
}

func Arrotondamento(a float64, n int){
  x := a * math.Pow(10, float64(n))
  b := math.Round(x)
  e := b / math.Pow(10, float64(n))

  fmt.Println("Il valore arrotondato è:", e)
}

func main() {
  var a float64
  var n int

  fmt.Println("inserisci una numero")
  fmt.Scan(&a)
  fmt.Println("inserici il numero di cifre dopo la virgola")
  fmt.Scan(&n)

  Troncamento(a,n)
  Arrotondamento(a,n)
}
