package main

import "fmt"

func CreateSet(words []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, w := range words {
		if _, ok := set[w]; !ok {
			set[w] = struct{}{}
		}
	}
	return set
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(CreateSet(words))
}
