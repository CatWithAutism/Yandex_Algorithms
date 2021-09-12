//C. Соседи
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := readIntegers(" ")

	posX := data[len(data)-2]
	posY := data[len(data)-1]

	flatMatrix := data[2 : cap(data)-2]
	matrix := make([][]int, data[0])
	for i := 0; i < data[0]; i++ {
		firstIndex := i * data[1]
		secondIndex := firstIndex + data[1]
		matrix[i] = flatMatrix[firstIndex:secondIndex]
	}

	var numbers []int
	for i := 0; i < data[0]; i++ {
		for j := 0; j < data[1]; j++ {
			if ((posY-1 == j || posY+1 == j) && posX == i) || ((posX-1 == i || posX+1 == i) && posY == j) {
				numbers = append(numbers, matrix[i][j])
			}
		}
	}

	sort.Ints(numbers)
	writeData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "[]"))
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
