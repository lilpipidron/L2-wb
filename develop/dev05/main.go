package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func match(line, pattern string, fixed, ignoreCase bool) bool {
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if fixed {
		return line == pattern
	}
	return strings.Contains(line, pattern)
}

func main() {
	after := flag.Int("A", 0, "print +N lines after match")
	before := flag.Int("B", 0, "print +N lines before match")
	context := flag.Int("C", 0, "print +-N lines around match")
	count := flag.Bool("c", false, "print count of lines with match")
	ignoreCase := flag.Bool("i", false, "ignore register")
	inverted := flag.Bool("v", false, "exclude instead of matching")
	fixed := flag.Bool("F", false, "exact match to string, not a pattern")
	lineNum := flag.Bool("n", false, "print line number")

	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Only 2 argument is allowed: match string and filename is required")
	}

	pattern := args[0]
	filename := args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matches []int
	var lines []string
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		str := scanner.Text()
		if match(str, pattern, *fixed, *ignoreCase) != *inverted {
			matches = append(matches, i)
		}
		lines = append(lines, str)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *count {
		fmt.Println(len(matches))
		return
	}

	if *context > 0 {
		*before = *context
		*after = *context
	}

	printed := make(map[int]struct{})
	for _, num := range matches {
		start := num - *before
		if start < 0 {
			start = 0
		}
		end := num + *after
		if end >= len(lines) {
			end = len(lines) - 1
		}

		for i := start; i <= end; i++ {
			if _, ok := printed[i]; !ok {
				if *lineNum {
					fmt.Print(i+1, " ")
				}
				fmt.Println(lines[i])
				printed[i] = struct{}{}
			}
		}
	}
}
