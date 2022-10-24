package main
import (
    "fmt"
)
func main() {
    var a int
    var b int

    fmt.Println("inserisci un numero")
    fmt.Scan(&a)
    fmt.Println("inserisci un numero")
    fmt.Scan(&b)

    var i int
    var ris int

    if a > b{
      for i = b; i < a; i++{
        if i % 2 == 0{
          ris += i
        }
      }
    } else {
      for i = a; i < b; i++{
        if i % 2 == 0{
          ris += i
        }
      }
    }

    fmt.Println(ris)
}
