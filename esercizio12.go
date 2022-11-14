package main

import (
  "fmt";
  "math"
)

func primo(a float64){
  x := int(a*100 + 0.5)
  c := float32(x)
  d := c / 100.0

  fmt.Println(d)

}

func secondo(a float64){
  x := a * 100
  b := math.Round(x)
  e := b / 100

  fmt.Println(e)
}

func main() {
  var a float64

  fmt.Scan(&a)

  primo(a)
  secondo(a)
}
