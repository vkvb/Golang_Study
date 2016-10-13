package main
import "fmt"
func main(){
  var sum uint64 = 2
  first:
  fmt.Println(sum)
  sum = sum * 2
  goto first
}
