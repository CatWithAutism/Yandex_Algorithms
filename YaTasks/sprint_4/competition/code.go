package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	raw := ReadIntegers()[1:]
	rawLength := len(raw)

	numbers := make([]int, rawLength + 1)
	numbersLength := rawLength + 1

	numbers[0] = 0
	copy(numbers[1:], raw)

	positionsOfNumbers := make(map[int][]int)
	positionsOfNumbers[0] = []int{ 0 }
	currentVal := 0

	for i := 1; i < numbersLength; i++ {
		if numbers[i] == 0 {
			currentVal--
		} else {
			currentVal++
		}

		if values, existing := positionsOfNumbers[currentVal]; existing{
			if len(values) > 1{
				values[1] = i
			} else {
				positionsOfNumbers[currentVal] = append(values, i)
			}
			continue
		}

		positionsOfNumbers[currentVal] = []int{ i }
	}

	if currentVal == 0{
		writeData(strconv.Itoa(rawLength))
		os.Exit(0)
	}

	maxStreak := 0
	for _, positionsOfNumber := range positionsOfNumbers {
		if len(positionsOfNumber) > 1{
			streak := numbers[positionsOfNumber[0] : positionsOfNumber[len(positionsOfNumber) - 1]]

			if len(streak) > maxStreak{
				maxStreak = len(streak)
			}
		}
	}

	writeData(strconv.Itoa(maxStreak))
}

//region Basic

func ReadIntegers() []int{
	var data []int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 512*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		values := strings.Split(scanner.Text(), " ")
		for _, value := range values {
			number, _ := strconv.Atoi(value)
			data = append(data, number)
		}
	}
	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}

//endregion