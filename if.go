package main

import "fmt"

func main() {
  var num int = 100

  if num <= 101 || num==100 {
    fmt.Println("num")
  } else if num < 100  {
    fmt.Println("太大了")
  } else{
    fmt.Println("error")
  }
}
