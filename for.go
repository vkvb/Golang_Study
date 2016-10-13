package main
import "fmt"
func main(){
  var base2 int64 = 2
  for i := 0; i< 10; i++{
    fmt.Println(base2)
    base2  = base2 * 2
  }
}
