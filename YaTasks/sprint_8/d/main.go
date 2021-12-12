package main

import (
	"bufio"
	"os"
	"strconv"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func solve(strs ...string) int {
	min := getMin(strs...)
	maxLen := 0
	found := false
	for i := 0; i < min; i++ {
		if found {
			break
		}

		current := strs[0][i]
		for j := range strs {
			if current != strs[j][i] {
				maxLen = i
				found = true
				break
			}
		}
	}

	if !found {
		return min
	}

	return maxLen
}

func getMin(strs ...string) int {
	min := len(strs[0])
	for i := range strs {
		current := len(strs[i])
		if min > current {
			min = current
		}
	}

	return min
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
	writeData(strconv.Itoa(solve(data[1:]...)))
}
