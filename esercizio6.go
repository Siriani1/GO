package main

import "fmt"

func main(){
  var a,b float32

  fmt.Scan(&a, &b)

  switch {
  default:
    fmt.Println(-b/a)
  case a == 0 && b != 0:
    fmt.Println("Impossibile")
  case a == 0 && b==0:
    fmt.Println("Impossibile")
  }



}
