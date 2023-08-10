package main
import "fmt"

var input string
var colorCode string

func main() {
	// input := 2
	// switch input {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3: 
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Invalid value")
	// }
	fmt.Scan(&input)
	switch input {
	case "blue":
		fmt.Println("#000FF")
	case "pink":
		fmt.Println("#FFC0CB")
	case "green":
		fmt.Println("#008000")
	case "yellow":
		fmt.Println("#FFFF00")

	default:
		fmt.Println("No color")
		
	}
}