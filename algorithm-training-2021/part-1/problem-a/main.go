package main

import (
	"fmt"
)

func main() {
	var s, j string

	_, _ = fmt.Scanf("%s", &s)
	_, _ = fmt.Scanf("%s", &j)

	seen := map[rune]struct{}{}
	for _, letter := range s {
		seen[letter] = struct{}{}
	}

	result := 0
	for _, stone := range j {
		if _, ok := seen[stone]; ok {
			result++
		}
	}

	fmt.Println(result)
}
