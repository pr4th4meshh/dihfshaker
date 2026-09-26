package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: diffcheck <file1> <file2>")
		os.Exit(1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	content1, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("error reading", file1, ":", err)
		os.Exit(1)
	}

	content2, err := os.ReadFile(file2)
	if err != nil {
		fmt.Println("error reading", file2, ":", err)
		os.Exit(1)
	}

	lines1 := strings.Split(string(content1), "\n")
	lines2 := strings.Split(string(content2), "\n")

	compare(lines1, lines2)
}

func compare(a, b []string) {
	max := len(a)
	if len(b) > max {
		max = len(b)
	}

	for i := 0; i < max; i++ {
		var oldLine, newLine string

		if i < len(a) {
			oldLine = a[i]
		}

		if i < len(b) {
			newLine = b[i]
		}

		if oldLine == newLine {
			fmt.Println(" ", oldLine)
		} else {
			if oldLine != "" {
				fmt.Println("-", oldLine)
			}

			if newLine != "" {
				fmt.Println("+", newLine)
			}
		}
	}
}
