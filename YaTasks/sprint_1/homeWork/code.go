package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const separator = " "

func main() {
	data := ReadIntegers()
	toBinary := data[0]
	var answer string
	for toBinary != 0 {
		answer += strconv.Itoa(toBinary % 2)
		toBinary /= 2
	}

	WriteData(Reverse(answer))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReadIntegers() []int {
	var data []int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), separator)
		for _, value := range values {
			number, _ := strconv.Atoi(value)
			data = append(data, number)
		}
	}
	return data
}

func WriteData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}
