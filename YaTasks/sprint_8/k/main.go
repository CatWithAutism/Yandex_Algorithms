package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func solve(first, second string) int {
	data := normalizeStrings(first, second)
	return strings.Compare(data[0], data[1])
}

func normalizeStrings(strs ...string) []string {
	result := make([]string, 0, len(strs))
	for i := range strs {
		sb := strings.Builder{}
		for j := range strs[i] {
			if strs[i][j]%2 == 0 {
				sb.WriteString(string(strs[i][j]))
			}
		}

		result = append(result, sb.String())
	}

	return result
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
	writeData(strconv.Itoa(solve(data[0], data[1])))
}
