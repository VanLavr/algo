package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	m := make(map[string]int)
	for input.Scan() {
		m[input.Text()]++
	}

	for line, n := range m {
		fmt.Println(n, line)
	}
}
