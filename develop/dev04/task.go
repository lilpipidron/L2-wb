package main

import (
	"fmt"
	"sort"
	"strings"
)

// Алгоритм просто из слов получает строки, где все буквы отсортированы и закидывает само слово в мапу
// Далее идет по мапе, и все слайсы с длинной более 1 закидывает в ответ

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	seenWords := make(map[string]bool)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortString(lowerWord)

		if !seenWords[lowerWord] {
			anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
			seenWords[lowerWord] = true
		}
	}

	result := make(map[string][]string)
	for _, anagramGroup := range anagramMap {
		if len(anagramGroup) > 1 {
			sort.Strings(anagramGroup)
			result[anagramGroup[0]] = anagramGroup
		}
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слово"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}
