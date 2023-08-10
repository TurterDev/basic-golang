// maps เป็นตัวแปรชนิดหนึ่งที่เก็บข้อมูลเป็นชุดๆ  เก็บข้อมูลในลักษณะ Key และ Value

package main

import "fmt"

//สร้างตัวแปรสำหรับ add ข้อมูล
var product = make(map[string]float64)

func main() {
	fmt.Println("Product =", product)

	//add data
	product["Macbook"] = 40000
	product["IPhone"] = 30000
	product["IPad"] = 25000
	fmt.Println("Product =", product)

	//delete data
	delete(product, "IPad")
	fmt.Println("Product =", product)

	//update data
	product["IPhone"] = 20000
	fmt.Println("Product =", product)

	//เข้าถึงข้อมูล
	value1 := product["IPhone"]
	fmt.Println("Value1 =", value1)

	courseName := map[string]string{"101":"Java","102":"Python"}

	fmt.Println("courseName =", courseName)

}