package main

import (
	"bufio"
	"os"
	"strconv"
)

const separator = " "

func main(){
	data := ReadIntegers()
	if data == 1{
		WriteData("True")
		os.Exit(0)
	}
	current := 4
	for current <= data {
		if current == data{
			WriteData("True")
			os.Exit(0)
		}
		current *= 4
	}

	WriteData("False")
}

func ReadIntegers() int{
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)
	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	return number
}


func WriteData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}