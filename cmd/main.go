package main

import (
	"fmt"
	"os"
	"strings"

	"diffcheck/internal/diff"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: diffcheck <file1> <file2>")
		os.Exit(2)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	content1, err := os.ReadFile(file1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading", file1, ":", err)
		os.Exit(2)
	}

	content2, err := os.ReadFile(file2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading", file2, ":", err)
		os.Exit(2)
	}

	lines1 := strings.Split(string(content1), "\n")
	lines2 := strings.Split(string(content2), "\n")

	printEdits(diff.Diff(lines1, lines2))
}

func printEdits(edits []diff.Edit) {
	for _, e := range edits {
		switch e.Op {
		case diff.Equal:
			fmt.Println(" ", e.Text)
		case diff.Delete:
			fmt.Println("-", e.Text)
		case diff.Insert:
			fmt.Println("+", e.Text)
		}
	}
}
