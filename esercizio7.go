package main

import (
  "fmt";
  "math"
)

func main() {
  var a,b,c float64

  fmt.Scan(&a, &b, &c)

  switch{
  default:
    fmt.Println( (-b + math.Sqrt((b*b)-4*a*c))/(2*a) )
  case (b*b)-4*a*c < 0:
    fmt.Println("Impossibile")
  }
}
