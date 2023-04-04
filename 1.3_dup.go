package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		printRepeats(counts)
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}

		countLines(f, counts)
		f.Close()
	}

	printRepeats(counts)

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "eof" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printRepeats(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
