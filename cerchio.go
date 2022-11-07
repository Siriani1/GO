package main

import "fmt"

func Calcoloarea(raggio float32) float32 {
  var area float32
  area = 3.14 * (raggio*raggio)
  return area
}

func main(){
  var raggio,area float32

  fmt.Scan(&raggio)

  area = Calcoloarea(raggio)

  fmt.Println(area)
}
