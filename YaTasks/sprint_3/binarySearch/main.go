package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	initReader()
	data := readInts()

	cash := data[1:len(data) - 1]
	bicycleCost := data[len(data) - 1]

	index := binarySearch(cash, bicycleCost, 0, len(cash))
	if index == -1{
		writeData(strconv.Itoa(index))
	} else {
		writeData(strconv.Itoa(index + 1))
	}

	index = binarySearch(cash, bicycleCost * 2, 0, len(cash))
	if index == -1{
		writeData(strconv.Itoa(index))
	} else {
		writeData(strconv.Itoa(index + 1))
	}
}

func binarySearch(collection []int, element int, left int, right int) int{
	if right <= left{
		return -1
	}

	pivot := (left + right) / 2
	if collection[pivot] >= element{
		for i := left; i < right; i++ {
			if collection[i] >= element{
				return i
			}
		}
	} else if element < collection[pivot]{
		return binarySearch(collection, element, left, pivot)
	} else {
		return binarySearch(collection, element, pivot + 1, right)
	}

	return -1
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader(){
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readInts() []int{
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
	writer.WriteString(data + " ")
	writer.Flush()
}

//endregion