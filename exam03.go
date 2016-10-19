package main
import ("fmt";"math")
func main(){
  var x1,y1,x2,y2 float64
  fmt.Scanf("%f,%f",&x1,&y1)
  fmt.Scanf("%f,%f",&x2,&y2)
  far := math.Sqrt((math.Pow(x2-x1,2)+(math.Pow(y2-y1,2))))
  line := (y2-y1)/(x2-x1)
  fmt.Printf("%.2f\n%.2f\n",far,line)
}
