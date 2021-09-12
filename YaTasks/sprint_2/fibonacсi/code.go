package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	data := readInts()
	writeData(strconv.Itoa(fibonacci(1, 1, data[0])))
}

func fibonacci(a, b, n int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	}

	return fibonacci(b, a+b, n-1)
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func initReader() {
	scanner.Split(bufio.ScanLines)
}

func readInts() []int {
	var data []int

	for scanner.Scan() {
		strNumber := scanner.Text()
		number, err := strconv.Atoi(strNumber)
		if err == nil {
			data = append(data, number)
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
