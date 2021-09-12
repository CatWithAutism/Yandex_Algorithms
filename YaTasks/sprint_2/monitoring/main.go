package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	initReader()
	rawArray := readInts()

	height := rawArray[0]
	width := rawArray[1]
	numbers := rawArray[2:]
	matrix := make([][]int, width)

	for i := 0; i < width; i++ {
		matrix[i] = make([]int, height)
		for j := 0; j < height; j++ {
			matrix[i][j] = numbers[j * width + i]
		}
		writeData(strings.Trim(fmt.Sprintf("%v", matrix[i]), "[]"))
	}
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
