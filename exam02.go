package main
import ("fmt";"math")
func main(){
  var half,face,all float64
  var pi float64 = 3.14159
  fmt.Scanf("%f",&half)
  face = 4*pi*(math.Pow(half,2))
  all  = 4*pi*math.Pow(half,3)/3
  fmt.Printf("%.2f\n%.2f\n",all,face)
}
