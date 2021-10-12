package main

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := int64(ReadIntegers()[0])
	big2NFactorial := big.NewInt(n).MulRange(1, n*2)
	bigN1Factorial := big.NewInt(n).MulRange(1, n+1)
	bigN := big.NewInt(n).MulRange(1, n)
	result := big2NFactorial.Div(big2NFactorial, bigN1Factorial.Mul(bigN1Factorial, bigN))
	writeData(result.String())
}

func factorial(n uint64) uint64 {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func ReadIntegers() []int {
	var data []int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		for _, value := range values {
			number, _ := strconv.Atoi(value)
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
