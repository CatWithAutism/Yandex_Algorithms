package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


// https://contest.yandex.ru/contest/22450/run-report/52316076/
func main(){
	initReader()
	rawArray := readInts()
	numbers := rawArray[1:len(rawArray)]
	dictionary := make(map[int]int)
	for _, number := range numbers {
		if val, exists := dictionary[number]; exists {
			val++
			dictionary[number] = val
			continue
		}

		dictionary[number] = 1
	}

	fingers := rawArray[0] * 2
	maxScores := 0
	for _, val := range dictionary {
		if val <= fingers{
			maxScores++
		}
	}


	writeData(strconv.Itoa(maxScores))
}



var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func initReader(){
	scanner.Split(bufio.ScanLines)
}

func readInts() []int{
	var data []int

	for scanner.Scan(){
		strNumbers := strings.Split(scanner.Text(), "")
		for _, strNumber := range strNumbers {
			number, err := strconv.Atoi(strNumber)
			if err == nil{
				data = append(data, number)
			}
		}
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}
