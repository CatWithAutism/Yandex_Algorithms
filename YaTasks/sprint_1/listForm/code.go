package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const separator = " "

func main() {
	Init()
	digits := ReadIntegers()
	kNumber := digits[len(digits)-1]
	xNumber := Join(digits[1 : len(digits)-1])

	WriteData(strings.Trim(strings.Join(strings.Split(strconv.Itoa(xNumber+kNumber), ""), separator), "[]"))
}

func Join(digits []int) int {
	total := 0
	order := 1
	for i := len(digits) - 1; i >= 0; i-- {
		total += digits[i] * order
		order *= 10
	}

	return total
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func Init() {
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)
}

func ReadIntegers() []int {
	var data []int

	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), separator)
		for _, strNumber := range strNumbers {
			number, _ := strconv.Atoi(strNumber)
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
