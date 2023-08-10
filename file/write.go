package main

import (
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("hello\n turter")
	err := os.WriteFile("file/data.txt", data1, 0644)
	check(err)
	// if err != nil {
	// 	panic(err)
	// }
	//สร้างไฟล์และเขียนไฟล์
	f, err := os.Create("employeerName")
	check(err)

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile("file/employee.txt", data2, 0644)
}