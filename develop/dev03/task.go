package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// В решении используются стандартные пакеты, есть вспомогательная структура
// для хранения всех опций сортировки, реализованы функции, чтобы можно было
// использовать паке sort на нашей структуре

type customSort struct {
	Lines             []string
	Column            int
	Numeric           bool
	Reverse           bool
	Unique            bool
	Month             bool
	Ignore            bool
	Sorted            bool
	NumericWithSuffix bool
}

func (cs *customSort) Len() int {
	return len(cs.Lines)
}

func (cs *customSort) Swap(i, j int) {
	cs.Lines[i], cs.Lines[j] = cs.Lines[j], cs.Lines[i]
}

func (cs *customSort) Less(i, j int) bool {
	fieldI := strings.Fields(cs.Lines[i])[cs.Column-1]
	fieldJ := strings.Fields(cs.Lines[j])[cs.Column-1]

	if cs.Numeric {
		numI, errI := strconv.ParseFloat(fieldI, 64)
		numJ, errJ := strconv.ParseFloat(fieldJ, 64)
		if errI != nil {
			log.Fatal(errI)
		}
		if errJ != nil {
			log.Fatal(errJ)
		}

		return numI < numJ != cs.Reverse
	}

	if cs.Month {
		timeI, errI := time.Parse("Jan", fieldI)
		timeJ, errJ := time.Parse("Jan", fieldJ)
		if errI != nil {
			log.Fatal(errI)
		}

		if errJ != nil {
			log.Fatal(errJ)
		}

		return timeI.Before(timeJ) != cs.Reverse
	}

	if cs.NumericWithSuffix {
		numI := parseWithSuffix(fieldI)
		numJ := parseWithSuffix(fieldJ)

		return numI < numJ != cs.Reverse
	}

	return fieldI < fieldJ != cs.Reverse
}

func parseWithSuffix(s string) float64 {
	s = strings.TrimSpace(s)
	var i int
	for i = 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) && s[i] != '.' {
			break
		}
	}
	numPart := s[:i]
	suffixPart := s[i:]
	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToUpper(suffixPart) {
	case "K":
		num *= 1e3
	case "M":
		num *= 1e6
	case "G":
		num *= 1e9
	case "T":
		num *= 1e12
	case "P":
		num *= 1e15
	case "E":
		num *= 1e18
	}

	return num
}

func removeDuplicates(data *customSort) {
	lines := make(map[string]struct{})
	var uniqueLines []string
	for _, line := range data.Lines {
		if _, ok := lines[line]; !ok {
			lines[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	data.Lines = uniqueLines
}

func main() {
	data := customSort{}
	data.Column = *flag.Int("k", 1, "coulmn for sort")
	data.Numeric = *flag.Bool("n", false, "sort by numeric value")
	data.Reverse = *flag.Bool("r", false, "sort in reverse order")
	data.Unique = *flag.Bool("u", false, "output only unique lines")
	data.Month = *flag.Bool("M", false, "sort by month name")
	data.Ignore = *flag.Bool("b", false, "ignore trailing spaces")
	data.Sorted = *flag.Bool("c", false, "check if sorted")
	data.NumericWithSuffix = *flag.Bool("h", false, "sort by numeric value with suffix")

	flag.Parse()

	/*args := flag.Args()
	if len(args) != 1 {
		log.Fatal("only the file name is allowed")
	}

	filename := args[0]
	*/
	filename := "test"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if data.Ignore {
			line = strings.TrimRight(line, " ")
		}
		data.Lines = append(data.Lines, line)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if data.Sorted {
		sorted := sort.IsSorted(&data)
		if sorted {
			fmt.Println("Sorted")
		} else {
			fmt.Println("Not sorted")
		}
		return
	}

	sort.Sort(&data)

	if data.Unique {
		removeDuplicates(&data)
	}

	for _, line := range data.Lines {
		fmt.Println(line)
	}
}
