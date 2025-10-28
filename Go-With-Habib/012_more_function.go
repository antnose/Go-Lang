package main

import "fmt"

func printSomething() {
	fmt.Println(("Education must be free"))
}

func sayHello(name string) {
	fmt.Println("Welcome to the go lang course", name)
}

func main() {
	printSomething()
	sayHello("Habib vai")
}
