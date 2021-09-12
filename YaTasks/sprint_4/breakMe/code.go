package main

import (
	"bufio"
	"os"
)

func main(){
	var qInt64 uint64 = 1000
	var modInt64 uint64 = 123987123

	str := "sdfgkjseroigudfkjgld"
	strLen := len(str)
	var hash uint64 = 0
	hash = uint64(str[0])
	for i := 1; i < strLen; i++ {
		hash = hash % modInt64
		hash = hash * qInt64 + uint64(str[i])
	}
	hash = hash % modInt64


	for str := range GenerateCombinations("abcdefghijklmnopqrstuvwxyz", 20) {
		strLen := len(str)

		var hash uint64 = 0
		hash = uint64(str[0])
		for i := 1; i < strLen; i++ {
			hash = hash % modInt64
			hash = hash * qInt64 + uint64(str[i])
		}
		hash = hash % modInt64

		if hash == 48649258{
			writeData(str)
			os.Exit(0)
		}
	}

}

//region Basic

func readStrings() []string{
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 512 * 15625
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)
	var data []string

	for scanner.Scan(){
		rawString := scanner.Text()
		data = append(data, rawString)
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}

//endregion

func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)

		AddLetter(c, "", alphabet, length)
	}(c)

	return c
}

func AddLetter(c chan string, combo string, alphabet string, length int) {
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}
}