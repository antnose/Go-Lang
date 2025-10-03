package main

import "fmt"

func main() {
	// simple switch
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")

	// case 3:
	// 	fmt.Println("Three")

	// case 4:
	// 	fmt.Println("Out of range!")

	// default:
	// 	fmt.Println("Others")
	// }

	// multiple conditional switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Off day!")

	// default:
	// 	fmt.Println("Work day!")
	// }

	// type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer")

		case string:
			fmt.Println("It's an string!")

		case bool:
			fmt.Println("It's a boolean")

		default:
			fmt.Println("Out of range!")
		}
	}

	whoAmI(5)
}
