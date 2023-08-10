//การใช้งาน function มี 4 รูปแบบ

//1. การใช้งานแบบไม่รับค่า และ return ค่า
//2. การใช้การแบบมีการรับค่า
//3. การใช้งานที่มีการ return ค่า
//4. การใช้งานแบบที่มีการรับค่าและ return ค่า

package main

import "fmt"

func hello(){
	fmt.Println("Hello Turter")
}

func plus(value1 int, value2 int) int{
	// result := value1 + value2
	// fmt.Println("result =",result)
	return value1 + value2
}

func plus3value(value1,value2,value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	// plus(1,2)
	result := plus(5,5)

	fmt.Println("result =", result)

	result2 := plus3value(5,5,10)

	fmt.Println("result =", result2)
}