package main
import "fmt"
func main(){
  var coin, coin10, coin5, coin1 int

  fmt.Scanf("%d",&coin)

  coin10 = coin/10
  coin5  = (coin%10)/5
  coin1  = coin%5

  fmt.Println(coin10,"\n",coin5,"\n",coin1)
}
