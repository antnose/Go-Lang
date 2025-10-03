package main

import "fmt"

// const age = 45

func main() {
	// const name = "golang"
	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}
