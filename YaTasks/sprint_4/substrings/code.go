package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	data := readStrings()
	longest(data[0])
}

func longest(str string) {
	maxLen := 0
	for i, _ := range str {
		subStr := str[i:]
		chars := make(map[int32]bool)

		var j int
		for j = 0; j < len(subStr); j++ {
			if _, existing := chars[int32(subStr[j])]; existing {
				if j > maxLen {
					maxLen = j
				}
				break
			}

			chars[int32(subStr[j])] = true
		}

		if j > maxLen {
			maxLen = j
		}
	}

	writeData(strconv.Itoa(maxLen))
}

//region Basic
func readStrings() []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)
	var data []string

	for scanner.Scan() {
		rawString := scanner.Text()
		data = append(data, rawString)
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion
