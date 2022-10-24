package main
import (
    "fmt"
)
func main() {

  var n float32
  var i bool = false
  var x float32
  var somma float32
  var media float32

  for i == false {
    fmt.Scan(&n)

    if n <= 0{
      i = true
    } else {
      somma += n
      x++

    }


  }

  media = somma / x

  fmt.Println(media)

}
