package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count := make(map[string]int)
	var files []string

	for _, arg := range os.Args[1:] {
		files = append(files, arg)
	}

	if len(files) == 0 {
		CountLines(os.Stdin, count)
		PrintMap(count)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				log.Fatal(err)
			}

			CountLines(f, count)
			fmt.Println("---")
			PrintMap(count)
			fmt.Println("---")
		}
	}
}

func CountLines(file *os.File, count map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		count[input.Text()]++
	}
}

func PrintMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("[%d] %s\n", value, key)
	}
}
