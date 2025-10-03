package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// ---> Classic For Loop
	// for d := 0; d < 3; d++ {
	// 	fmt.Println(d)
	// }

	for i := range 3 {
		fmt.Println(i)
	}
}
