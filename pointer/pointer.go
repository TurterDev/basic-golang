// pointer คือการชี้ไปที่ adress ว่าข้อมูลนั้นอยู่ที่ไหน
//  คือ * บอกว่าเป็น pointer
package main

import "fmt"

func zerovalue(ivalue int){
	ivalue = 0
}

func zeropointer(ipointer *int){
	*ipointer = 0
}

func main() {

	i := 1
	fmt.Println("i =", i)

	zerovalue(i)
	fmt.Println("i form function zerovalue =", i)

	zeropointer(&i)
	fmt.Println("i value form function zerovalue =", i)
	fmt.Println("i address form function zerovalue =", &i)
}