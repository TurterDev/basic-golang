// slice เป็นตัวแปรที่ใช้งานเหมือนอาเรย์ แต่มีความยึดหยุ่นกว่า โดยไม่จำเป็นต้องกำหนดขนาดของข้อมูลก่อน

package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "C", "C#", "HTML", "CSS")

	fmt.Println(courseName)

	courseWeb := courseName[4:6]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}