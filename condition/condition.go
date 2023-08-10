package main

import "fmt"

var score int

func main() {

	// myMoney := 100
	// if myMoney > 100 {
	// 	fmt.Println("ขออภัยคุณสามารถซื้อกระเป๋าใบนี้ได้")
	// }else {
	// 	fmt.Println("ขออภัยเงินคุณไม่เพียงพอ")
	// }
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("score =", score)
	if score >= 80 {
		fmt.Println("grade A")
	} else if score >= 70 {
		fmt.Println("grade B")
	} else if score >= 60 {
		fmt.Println("grade C")
	} else if score >= 50 {
		fmt.Println("grade D")
	} else {
		fmt.Println("grade F")
	}

}
