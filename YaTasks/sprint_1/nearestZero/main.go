package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//мне голанг не родной, так что если о каком-то функционале не знаю, то можно бить, чтобы узнал :D
// https://contest.yandex.ru/contest/22450/run-report/52315993/
func main(){
	initReader()
	rawArray := readInts()
	numbers := rawArray[1:len(rawArray)] //на не очень нужно первое число
	outputArray := make([]int, 0)
	counter := 0
	foundFirstZero := false
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0{
			outputArray = append(outputArray, numbers[i])
			innerCounter := 1
			//слайсы для слабых :P
			if foundFirstZero {
				for j := i - 1; j >= len(outputArray) - 1 - counter / 2; j-- {
					outputArray[j] = innerCounter
					innerCounter++
				}
			} else {
				foundFirstZero = true
				for j := i - 1; j >= 0; j-- {
					outputArray[j] = innerCounter
					innerCounter++
				}
			}
			counter = 1
			continue
		}

		outputArray = append(outputArray, counter)
		counter++
	}

	writeData(strings.Trim(fmt.Sprintf("%v", outputArray), "[]"))
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func initReader(){
	scanner.Split(bufio.ScanWords)
}

func readInts() []int{
	var data []int

	for scanner.Scan(){
		strNumber := scanner.Text()
		number, err := strconv.Atoi(strNumber)
		if err == nil{
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
