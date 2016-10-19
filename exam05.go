package main
import "fmt"
func main(){
  var m,s,tmp int64
  fmt.Scanf("%d %d",&s,&m)
    for m != 0{
      tmp = s % m
      s = m
      m = tmp
    }
  fmt.Println(m)
}
