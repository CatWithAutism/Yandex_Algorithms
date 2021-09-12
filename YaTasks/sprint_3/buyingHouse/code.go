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
	budget := data[1]
	buildingCosts := data[2:]
	sort.Ints(buildingCosts)
	count := 0
	for _, buildingCost := range buildingCosts {
		if buildingCost > budget {
			break
		}

		budget -= buildingCost
		count++
	}
	writeData(strconv.Itoa(count))
	os.Exit(0)
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
