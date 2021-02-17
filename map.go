package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
Needles and pins
Needles and pins
Sew me a sail
To catch me the wind
`
	fieldsFunc := strings.Fields(text)
	count := map[string]int{}
	for _, word := range fieldsFunc {
		k := count[word]
		k++
		count[word] = k
	}
	for word, count := range count {
		fmt.Printf("%s: %d,\n ", word, count)
	}
}
