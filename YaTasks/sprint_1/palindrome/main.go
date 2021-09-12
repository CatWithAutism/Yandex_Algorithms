package main

import (
	"bufio"
	"os"
	"unicode"
)

func main(){
	data := ReadText()
	arrLength := len(data)
	startIndex := 0
	endIndex := arrLength - 1
	for endIndex >= startIndex{
		if (unicode.IsLetter(rune(data[startIndex])) && unicode.IsLetter(rune(data[endIndex]))) ||
			(unicode.IsDigit(rune(data[startIndex])) && unicode.IsDigit(rune(data[endIndex]))){
			firstByte := data[startIndex]
			secondByte := data[endIndex]
			if firstByte > 91{
				firstByte -= 32
			}

			if secondByte > 91{
				secondByte -= 32
			}

			if firstByte != secondByte{
				WriteData("False")
				os.Exit(0)
			}

			startIndex++
			endIndex--
		} else {
			if !unicode.IsLetter(rune(data[startIndex])) || !unicode.IsDigit(rune(data[startIndex])){
				startIndex++
			} else if !unicode.IsLetter(rune(data[startIndex])) || !unicode.IsDigit(rune(data[startIndex])){
				endIndex++
			}
		}
	}

	WriteData("True")
}

func ReadText() string{
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)
	scanner.Scan()
	return scanner.Text()
}


func WriteData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}