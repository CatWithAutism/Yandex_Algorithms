package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const separator = " "

func main(){
	data := ReadIntegers()
	var simpleDimple []int
	diviner := 2
	for data > 1{
		reminder := data % diviner

		if reminder == 0 {
			simpleDimple = append(simpleDimple, diviner)
			data = data / diviner
			diviner = 2
			continue
		}

		quotient := data / diviner
		quotient++

		if quotient < diviner || quotient == diviner {
			simpleDimple = append(simpleDimple, data)
			break
		}

		if diviner == 2{
			diviner++
		} else {
			diviner +=2
		}
	}

	WriteData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(simpleDimple)), " "), "[]"))
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