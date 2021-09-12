package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	initReader()
	vars := readInts()[1:]
	sort.Slice(vars, func(i, j int) bool {
		if vars[i] < vars[j] {
			return false
		}

		return true
	})
	for i := 0; i < len(vars); i++ {
		for j := i + 1; j < len(vars); j++ {
			for k := j + 1; k < len(vars); k++ {
				if vars[i] < vars[j]+vars[k] {
					writeData(strconv.Itoa(vars[i] + vars[j] + vars[k]))
					os.Exit(0)
				}
			}
		}
	}
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func initReader() {
	scanner.Split(bufio.ScanLines)
}

func readInts() []int {
	var data []int

	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), " ")
		for _, strNumber := range strNumbers {
			number, err := strconv.Atoi(strNumber)
			if err == nil {
				data = append(data, number)
			}
		}
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion
