package main
import ("fmt";"math")
func main(){
  var aa float64 = 64
  a:=math.Sqrt(aa)
  b:=fmt.Sprintf("%.5f",a)
  fmt.Println(b)
}
