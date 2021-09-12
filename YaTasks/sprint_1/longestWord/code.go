//E. Самое длинное слово
package main

import (
	"bufio"
	"os"
	"strconv"
)

const whiteSpace = 32

func main() {
	data := readStrings()
	maxLength, indexOfMax, currentLength, indexOfCurrent := -1, -1, -1, -1
	arrayLength := len(data)
	for i := 0; i < arrayLength; i++ {
		if data[i] == whiteSpace {
			if indexOfCurrent != -1 && currentLength != -1 && maxLength < currentLength {
				maxLength = currentLength
				indexOfMax = indexOfCurrent
			}
			indexOfCurrent = -1
			currentLength = -1
		} else if currentLength != -1 && indexOfCurrent != -1 {
			currentLength++
		} else {
			indexOfCurrent = i
			currentLength = 1
		}

		if i == arrayLength-1 && indexOfCurrent != -1 && currentLength != -1 && maxLength < currentLength {
			maxLength = currentLength
			indexOfMax = indexOfCurrent
		}
	}

	writeData(data[indexOfMax : indexOfMax+maxLength])
	writeData(strconv.Itoa(maxLength))
}

func readStrings() string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)

	for scanner.Scan() {
		scanner.Text()
		return scanner.Text()
	}

	return ""
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(data + "\n")
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
