package main
import "fmt"
func main(){
  var numA,numB float64
  fmt.Scanf("%f",&numA)
  fmt.Scanf("%f",&numB)
  tmp:= fmt.Sprintf("%.4f",(numA*100)/numB)
  fmt.Println(tmp)
}
