package main

import "fmt"

func main() {

	counter := 0
	for i := 1000; i <= 9999; i++ {
		for k:= i; k <= 9999; k ++ {
			result := i * k
			s := fmt.Sprintf("%d", result)
			if s[0] == s[len(s) - 1] {
				counter++
			}
		}
	}
	fmt.Printf("Result %d", counter)

}
