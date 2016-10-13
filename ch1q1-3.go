package main
import "fmt"
func main(){
i:=-1
first:
i++
  if i<6{
    fmt.Println(i)
    goto first
  }
}
