package main

import (
	"bufio"
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
	initReader()
	data := readInts()[0]
	comboPhone(strconv.Itoa(data), "");
}

func getDigit(data string) (int, string){
	if len(data) > 0{
		atoi, err := strconv.Atoi(string(data[0]))
		if err != nil {
			return -1, ""
		}
		return atoi, data[1:]
	}

	return -1, ""
}

func comboPhone(numberStr string, current string){
	if number, numberStr := getDigit(numberStr); number == -1 {
		return
	} else {
		for i := 0; i < len(keys[number]); i++ {
			if numberStr == ""{
				writeData(current + string(keys[number][i]))
			}

			comboPhone(numberStr, current + string(keys[number][i]))
		}
	}
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