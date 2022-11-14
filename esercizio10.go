package main

import (
  "ftm";
  "math"
)
const EPSIOLON = 1.0e-9

func XMaggioreDiY(x, y float64) bool{
  return (x - y) > EPSIOLON
}

func XUgualeDiY (x,y float64) bool {
  return (x - y) <= EPSIOLON
}

func XMinoreDiY (x,y float64) bool {
  return (x - y) < -EPSIOLON
}

func main() {
  var s int
  var m,q float32

  fmt.Scan(&s, &m, &q)


}
