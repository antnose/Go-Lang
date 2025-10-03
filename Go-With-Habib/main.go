// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// 	a := 10
// 	a = 45
// 	fmt.Println(a)
// 	const b = "Rohim"
// 	// b = "Karim"
// 	fmt.Println(b)
// }

// If-else and switch
package main

import "fmt"

func main() {
	age := 15
	if age > 18 {
		fmt.Println("You are eligible for married ")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married, but you can love someone")
	} else {
		fmt.Println("You are just a teenager, not eligible for married")
	}
}
