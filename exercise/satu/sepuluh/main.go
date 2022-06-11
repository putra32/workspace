package main

import "fmt"

func main() {
	count := 5
	
	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	name := "Jhon"

	name += " Smith"
	fmt.Println("Hello, ", name)
}
