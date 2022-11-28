package main

import (
  "fmt"
  "bufio"
  "os"
)


func LeggiTesto() string {
  var testo string
  scanner := bufio.NewScanner(os.Stdin)
  for  scanner.Scan(){
      testo = testo + scanner.Text() + "\n"
  }
  return testo
}

func main(){

  s := LeggiTesto()
  fmt.Println(s)


}
