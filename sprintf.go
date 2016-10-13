package main
import "fmt"

func main(){
  var aa int = 10
  var bb int = 3
  var cc float64 = float64(aa)/float64(bb)
  s:= fmt.Sprintf("%.5f",cc)
  fmt.Println(s)
}
