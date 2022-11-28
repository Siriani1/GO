package main

import "fmt"


func main() {

  var s string

  fmt.Scan(&s)

  fmt.Printf("Output: %c", s[0])
  for i := 1; i < len(s); i++{
    if s[i] != s[i-1]{
      fmt.Printf("%c", s[i])
    }

  }
  fmt.Println()
}
