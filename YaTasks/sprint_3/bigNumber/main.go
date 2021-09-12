package main

import (
	"bufio"
	"os"
	"sort"
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
	initReader()
	data := readInts()[1:]
	sort.Strings(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	for i, j := 0, 1; i < len(data) - 1; i, j = i+1, j+1 {
		if data[i][0] == data[j][0] && len(data[i]) != len(data[j]){
			data[j], data[i] = data[i], data[j]
		}
	}
	writeData(strings.Join(data, ""))
}

func maxInt(first int, second int) int{
	if first > second{
		return first
	}

	return second
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader(){
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readInts() []string{
	var data []string

	for scanner.Scan(){
		strNumbers := strings.Split(scanner.Text(), " ")
		for _, strNumber := range strNumbers {
			data = append(data, strNumber)
		}
	}

	return data
}


func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + " ")
	writer.Flush()
}

//endregion