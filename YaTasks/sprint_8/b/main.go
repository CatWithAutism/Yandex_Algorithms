package main

import (
	"bufio"
	"math"
	"os"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func solve(first, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	minLen := getMin(len(first), len(second))
	for i := 0; i < minLen; i++ {
		if first[i] != second[i] {
			if len(first) > len(second) {
				return first[i+1:] == second[i:]
			} else if len(first) < len(second) {
				return first[i:] == second[i+1:]
			} else {
				return first[i+1:] == second[i+1:]
			}

		}
	}

	return true
}

func getMin(first, second int) int {
	if first < second {
		return first
	}

	return second
}

func getData() []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1024*1024*512), 1024*1024*512)
	data := make([]string, 0, 2)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func main() {
	data := getData()
	if solve(data[0], data[1]) {
		writeData("OK")
		os.Exit(0)
	}
	writeData("FAIL")
}
