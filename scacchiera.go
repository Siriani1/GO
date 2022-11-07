package main

import "fmt"

func StampaRigainizioAsterisco(lunghezza int) {
  var scacchiera string
  for i := 0; i < lunghezza; i++{
    if i % 2 == 0{
      scacchiera += "*"
    }else{
      scacchiera += "+"
    }
  }
  fmt.Println(scacchiera)
}

func StampaRigaInizioPiu(lunghezza int) {
  var scacchiera string
  for i := 0; i < lunghezza; i++{
    if i % 2 == 0{
      scacchiera += "+"
    }else{
      scacchiera += "*"
    }
  }
  fmt.Println(scacchiera)
}

func StampaScacchiera(dimensione int){
  for i:= 0; i < dimensione; i++{
    if i % 2 == 0{
      StampaRigainizioAsterisco(dimensione)
    }else{
      StampaRigaInizioPiu(dimensione)
    }
  }
}


func main() {
  var dimensione int

  fmt.Println("inserisci la diensione")
  fmt.Scan(&dimensione)
  if dimensione >= 0{
    StampaScacchiera(dimensione)
  }

}
