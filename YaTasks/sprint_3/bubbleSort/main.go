package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var keys = map[int]string{
	2:"abc",
	3:"def",
	4:"ghi",
	5:"jkl",
	6:"mno",
	7:"pqrs",
	8:"tuv",
	9:"wxyz",
}

func main(){
	initalizeReader()
	data := readIntegers()[1:]
	notReady := true
	firstShown := false
	for notReady{
		notReady = false
		for i := 0; i < len(data) - 1; i++ {
			next := i + 1
			if data[i] > data[next]{
				data[i], data[next] = data[next], data[i]
				notReady = true
			}
		}

		if notReady || !firstShown{
			writeData(strings.Trim(fmt.Sprintf("%v", data), "[]"))
			firstShown = true
		}
	}
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initalizeReader(){
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readIntegers() []int{
	var data []int

	for scanner.Scan(){
		strNumbers := strings.Split(scanner.Text(), " ")
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

//endregion