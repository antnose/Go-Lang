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

// package main

// import "fmt"

// func add(num1 int, num2 int) {
// 	sum := num1 + num2
// 	fmt.Println(sum)
// }

// func main() {
// 	add(12, 44)

// 	add(5, 8)
// }

// package main

// import "fmt"

// func printSomething() {
// 	fmt.Println(("Education must be free"))
// }

// func sayHello(name string) {
// 	fmt.Println("Welcome to the go lang course", name)
// }

// func main() {
// 	printSomething()
// 	sayHello("Habib vai")
// }

// -----------------> [Golang] 013 - Why Functions Are Needed <----------------- //

// package main

// import "fmt"

// func printWelcomeMessage() {
// 	fmt.Println("Welcome to the application")
// }

// func getUserName() string {
// 	var name string
// 	fmt.Print("Enter your name : ")
// 	fmt.Scan(&name)
// 	return name
// }

// func getTwoNumber() int {
// 	var num1 int
// 	var num2 int
// 	fmt.Print("Enter your first number : ")
// 	fmt.Scan(&num1)

// 	fmt.Print("Enter your second number : ")
// 	fmt.Scan(&num2)

// 	return (num1 + num2)
// }

// func showUserMessage(name string, sum int) {
// 	fmt.Printf("Welcome %s your sum is %d", name, sum)
// }

// func main() {
// 	printWelcomeMessage()
// 	name := getUserName()
// 	sum := getTwoNumber()

// 	showUserMessage(name, sum)
// }

// -----------------> [Golang] 014 - What Is Scope <----------------- //
// package main

// import "fmt"

// var (
// 	a = 20
// 	b = 30
// )

// func add(x int, y int) {
// 	fmt.Println(x + y)
// }

// func main() {
// 	var p = 30
// 	var q = 40

// 	add(p, q)
// 	add(a, b)
// 	add(q, a)
// }

// Primarily scopes are 3 types:
// 1. Block Scope / Local Scope
// 2. Package Scope / File Scope
// 3. Global Scope

// -----------------> [Golang] 016 - Package Scope <----------------- //
// package main

// import (
// 	"example.com/mathlib"
// )

// func main() {
// 	mathlib.Add(6, 4)
// }

// -----------------> [Golang] 017 - Scope With Another Boring Example <----------------- //
// package main

// import "fmt"

// var (
// 	a = 10
// 	b = 20
// )

// func printNum(num int) {
// 	fmt.Println(num)

// }

// func add(a int, b int) {
// 	res := a + b
// 	printNum(res)
// }

// func main() {
// 	add(a, b)
// }

// ----------------->[Golang] 018 - Variable Shadowing <----------------- //
package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age >= 18 {
		a := 47
		fmt.Println(a)
	}

	fmt.Println(a)
}
