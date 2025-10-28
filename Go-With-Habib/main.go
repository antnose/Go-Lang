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
// package main

// import "fmt"

// func main() {
// 	age := 15
// 	if age > 18 {
// 		fmt.Println("You are eligible for married ")
// 	} else if age < 18 {
// 		fmt.Println("You are not eligible to be married, but you can love someone")
// 	} else {
// 		fmt.Println("You are just a teenager, not eligible for married")
// 	}
// }

// and => &&
// or => ||
// not => !

// package main

// import "fmt"

// func main() {
// 	age := 10
// 	gender := "female"

// 	if age >= 18 && gender == "male" {
// 		fmt.Println("You are responsible for everything!")
// 	} else {
// 		fmt.Println("Nothing!..")
// 	}

// 	if age >= 18 || gender == "female" {
// 		fmt.Println("You are good to gooooo!")
// 	} else {
// 		fmt.Println("We are so sorry! We can't do anything!")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	a := 30

// 	switch a {
// 	case 1:
// 		fmt.Println("a is 1")
// 	case 2:
// 		fmt.Println("a is 2")
// 	case 3:
// 		fmt.Println("A is 3")
// 	default:
// 		fmt.Println("It's out of range!")
// 	}
// }

// -------------> Function <-------------//

// package main

// import "fmt"

// func main() {
// 	a := 10
// 	b := 20

// 	sum := a + b
// 	fmt.Println(sum)
// }

package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func main() {
	add(12, 44)

	add(5, 8)
}
