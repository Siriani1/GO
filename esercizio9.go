package main

import (
  "fmt";
  "math"
)

func main(){
  var x float64

  fmt.Scan(&x)

  radice := math.Sqrt(x)

  switch{
  default:
    fmt.Println(x, "uguale a", radice, "*", radice, "\n")
  case x != radice * radice:
    fmt.Println(x, "diverso da", radice, "*", radice, "\n")
  }

}
