package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	initReader()
	data := readInts()[0]
	printBrackets(data, "", 0, 0)
}

func printBrackets(pairs int, prefix string, opened int, closed int) {
	if pairs == opened && closed == pairs {
		writeData(prefix)
		return
	}

	if opened < pairs {
		printBrackets(pairs, prefix+"(", opened+1, closed)
	}

	if closed < opened {
		printBrackets(pairs, prefix+")", opened, closed+1)
	}
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader() {
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readInts() []int {
	var data []int

	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), " ")
		for _, strNumber := range strNumbers {
			number, err := strconv.Atoi(strNumber)
			if err == nil {
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
