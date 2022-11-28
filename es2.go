package main

import (
  "fmt"
  "bufio"
  "os"
  "unicode"
)

func LeggiTesto() string {
  var testo string
  scanner := bufio.NewScanner(os.Stdin)
  for  scanner.Scan(){
      testo = testo + scanner.Text() + "\n"
  }
  return testo
}

func StatisticheParole(s string) (int, int){
  var contatore int
  var contLettera int
  var isSep bool = true

  for _, r := range s {
    if unicode.IsLetter(r) && isSep == true{
      contatore ++
    }
    if unicode.IsLetter(r){
      contLettera ++
      isSep = false
    }else{
      isSep = true
    }
  }

  return contatore, contLettera
}

func main(){


  s := LeggiTesto()
  contatore, contLettera := StatisticheParole(s)

  fmt.Println(contatore, contLettera)

}
