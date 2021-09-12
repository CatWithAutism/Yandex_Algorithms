//D. Хаотичность погоды
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readIntegers(" ")
	weatherData := data[1:]
	counter := 0
	for i := 0; i < len(weatherData); i++ {
		if isChaoticDay(weatherData, i) {
			counter++
		}
	}

	writeData(strconv.Itoa(counter))
}

func isChaoticDay(data []int, n int) bool {
	if len(data) == 1 {
		return true
	} else if n == 0 && data[n] > data[n+1] {
		return true
	} else if n == len(data)-1 && data[n] > data[n-1] {
		return true
	} else if n > 0 && n < len(data)-1 && data[n] > data[n-1] && data[n] > data[n+1] {
		return true
	}

	return false
}

func readIntegers(separator string) []int {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)

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
