package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
  fmt.Println(os.Args)
  var m int = 1
  for _,a := range os.Args[1:]{
    fmt.Printf("%T, %V\n", a, a)
    i, err := strconv.Atoi(a)
    if err == nil {
      m *= i
    }
  }
  fmt.Println("Il prodotto è:", m)



}
