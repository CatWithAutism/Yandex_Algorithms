package main

import (
	"bufio"
	"os"
	"strconv"
)

func main(){
	vals := readStrings()
	qInt64, _ := strconv.ParseUint(vals[0], 10, 64)
	modInt64, _ := strconv.ParseUint(vals[1], 10, 64)

	str := vals[2]
	strLen := len(str)

	if strLen == 0 {
		writeData("0")
		os.Exit(0)
	}

	hash := uint64(str[0])
	for i := 1; i < strLen; i++ {
		hash = hash % modInt64
		hash = hash * qInt64 + uint64(str[i])
	}
	hash = hash % modInt64
	writeData(strconv.FormatUint(hash, 10))
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