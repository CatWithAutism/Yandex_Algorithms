//a. Значения функции
package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readIntegers(" ")
	square := int(math.Pow(float64(data[1]), 2))
	writeData(strconv.Itoa((data[0] * square) + (data[2] * data[1]) + data[3]))
}

func readIntegers(separator string) []int {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var data []int

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), separator)
		for _, value := range values {
			number, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			data = append(data, number)
		}
	}
	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(data + "\n")
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
