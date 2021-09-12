//B. Чётные и нечётные числа
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readIntegers(" ")
	divider := 2
	remains := (data[0]%divider + divider) % divider
	for i := 1; i < len(data); i++ {
		if (data[i]%divider+divider)%divider != remains {
			writeData("FAIL")
			os.Exit(0)
		}
	}

	writeData("WIN")
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
