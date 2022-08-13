package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 31

	if input >= 1 && input <= 100 {
		for i:= 1; i <= input; i++ {
			if i % 3 == 0 && i % 5 == 0 {
				fmt.Println("FizzBuzz")
			} else if i % 5 == 0 {
				fmt.Println("Buzz")
			} else if i % 3 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println(strconv.Itoa(i))
			}
		}
	}else {
		fmt.Println("Must great than 0 and last than 100")
}
