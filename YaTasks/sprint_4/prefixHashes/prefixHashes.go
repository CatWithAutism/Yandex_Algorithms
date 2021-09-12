package main

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main(){
	vals := readStrings()
	qInt64, _ := strconv.ParseInt(vals[0], 10, 64)
	modInt64, _ := strconv.ParseInt(vals[1], 10, 64)
	buffer := make([]int64, 0)

	str := vals[2]
	strLen := len(str)

	if strLen == 0 {
		writeData("0")
		os.Exit(0)
	}

	hash := int64(str[0])
	hash = hash % modInt64
	buffer = append(buffer, hash % modInt64)

	for i := 1; i < strLen; i++ {
		hash = hash * qInt64 + int64(str[i])
		hash = hash % modInt64
		buffer = append(buffer, hash % modInt64)
	}

	substringPointers := vals[4:]
	for _, rawVal := range substringPointers {
		rawVals := strings.Split(rawVal," ")
		pointer1, _ := strconv.ParseInt(rawVals[0], 10, 64)
		pointer2, _ := strconv.ParseInt(rawVals[1], 10, 64)

		var pointer1Hash int64
		if int(pointer1) - 2 < 0 {
			pointer1Hash = 0
		} else {
			pointer1Hash = buffer[pointer1 - 2]
		}

		powOf := powInt(qInt64, (pointer2 - 1) - (pointer1 - 1) + 1, modInt64)
		mulOf := (pointer1Hash * powOf) % modInt64
		subHash := ((buffer[pointer2 - 1] - mulOf) % modInt64 + modInt64) % modInt64
		writeData(strconv.FormatInt(subHash, 10))
	}
}

func powInt(base, index, modulus int64) int64 {
	bigBase := big.NewInt(base)
	bigIndex := big.NewInt(index)
	bigModulus := big.NewInt(modulus)

	return bigBase.Exp(bigBase, bigIndex, bigModulus).Int64()
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