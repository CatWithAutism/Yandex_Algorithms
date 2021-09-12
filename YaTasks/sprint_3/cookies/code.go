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
	data := readInts()
	babies := data[1 : data[0]+1]
	cookies := data[data[0]+2:]

	sort.Ints(babies)
	sort.Ints(cookies)
	count := 0
	for i := 0; i < len(babies); i++ {
		for j := count; j < len(cookies); j++ {
			if babies[i] <= cookies[j] {
				cookies[j] = -1001
				count++
				break
			}
		}
	}

	writeData(strconv.Itoa(count))
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
