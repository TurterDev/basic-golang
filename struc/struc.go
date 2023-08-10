package main

import "fmt"

type employee struct {
	employeeID string
	employeeName string
	phone string
}

func main() {
	// employee1 := employee{
	// 	employeeID: "101",
	// 	employeeName: "Turter",
	// 	phone: "1234567890",
	// }
	// fmt.Println("employee1 = ", employee1)

	// add to array
	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID: "101",
		employeeName: "Turter",
		phone: "1234567890",
	}
	employeeList[1] = employee{
		employeeID: "102",
		employeeName: "Victor",
		phone: "1234567890",
	}
	employeeList[2] = employee{
		employeeID: "103",
		employeeName: "9james",
		phone: "1234567890",
	}
	// add to slice
	employeeList2 := []employee{}
	employee2 := employee{
		employeeID: "101",
		employeeName: "Turter",
		phone: "1234567890",
	}
	employee3 := employee{
		employeeID: "102",
		employeeName: "Victor",
		phone: "1234567890",
	}
	employeeList2 = append(employeeList2, employee2)
	employeeList2 = append(employeeList2, employee3)
	fmt.Println("employee = ", employeeList2)
}