package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	for j := 0; j < len(s); j++ {
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				swap(s, i-1, i)
			}
		}
	}
	return s
}
func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
