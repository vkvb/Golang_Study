package main

import (

    "fmt"

)

func main() {
		var num int
		fmt.Scan(&num)
		if num <= 20 && num >= 0 {
			fmt.Print("您獲得",15000+(num*380),"元~~")
		}else if num > 20 && num <= 40 {
			fmt.Print("您獲得",15000+(num*420),"元~~")
		}else{
			fmt.Print("您獲得",15000+(num*420)+((((num-40)/10))*1500),"元~~")
		}
}
