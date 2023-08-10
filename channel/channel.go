// chan คือบอกว่าตัวแปรนั้นใช้ channel
//channel เป็นช่องที่ใข้ในการสือสารกันระหว่าง retine1 และ rutine2

package main

import (
	"fmt"
	"time"
)

//c, data คือตัวแปร
func process1(c chan string,data string){
	c <- data
}

func main() {
	ch := make(chan string, 1)
	go process1(ch, "Data1")
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
}