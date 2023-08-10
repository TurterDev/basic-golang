// channel เป็นตัวแปรที่ใช้ในการสื่อสารระหว่าง go rutine กับ go rutine
// go rutine คือ Process หรือการทำงานย่อยๆ โดยประมวลผลไปพร้อมๆกัน

package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := 0; i < 100; i++ {
		fmt.Println(from ,":", i)
		
	}
}

func main() {
	// go คือ gorutine
	go f("Hello Turter")
	go f("message2")
	time.Sleep(5*time.Second)
}