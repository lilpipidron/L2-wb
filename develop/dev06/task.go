package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Используем "flag" для получения всех флагов, и "os" для работы с файлом
func selectSeparatedStrings(separated bool, str []string, delimiter string) []string {
	if !separated {
		return str
	}
	var result []string
	for _, s := range str {
		if strings.Contains(s, delimiter) {
			result = append(result, s)
		}
	}

	return result
}

func selectFields(words [][]string, fields int) [][]string {
	if fields < 1 {
		return words
	}
	result := make([][]string, len(words))
	for i, str := range words {
		for position, word := range str {
			if position == fields-1 {
				result[i] = append(result[i], word)
			}
		}
	}

	return result
}

func splitLinesByDelimiter(delimiter string, str []string) [][]string {
	var result [][]string
	for _, line := range str {
		result = append(result, strings.Split(line, delimiter))
	}
	return result
}

func main() {
	fields := flag.Int("fields", -1, "Number of fields")
	delimiter := flag.String("delimiter", " ", "Field delimiter (optional)")
	separated := flag.Bool("separated", false, "Separated fields")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Only 1 argument is allowed: filename is required")
	}

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	var str []string
	for reader.Scan() {
		str = append(str, reader.Text())
	}

	if err = reader.Err(); err != nil {
		log.Fatal(err)
	}

	str = selectSeparatedStrings(*separated, str, *delimiter)
	words := splitLinesByDelimiter(*delimiter, str)

	for _, s := range selectFields(words, *fields) {
		if len(s) == 0 {
			continue
		}
		for _, word := range s {
			fmt.Print(word, " ")
		}
		fmt.Println()
	}
}
