package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

type StrIndex struct {
	str   string
	index int
}

func convertToStrIndex(strs ...string) []StrIndex {
	result := make([]StrIndex, 0, len(strs))

	for i := range strs {
		split := strings.Split(strs[i], " ")
		val, _ := strconv.Atoi(split[1])
		result = append(result, StrIndex{
			str:   split[0],
			index: val,
		})
	}

	return result
}

func solve(strs ...string) string {
	strIndexes := convertToStrIndex(strs[2:]...)
	sort.Slice(strIndexes, func(i, j int) bool {
		return strIndexes[i].index < strIndexes[j].index
	})

	sb := strings.Builder{}
	prev := 0
	for i := range strIndexes {
		sb.WriteString(strs[0][prev:strIndexes[i].index])
		prev = strIndexes[i].index
		sb.WriteString(strIndexes[i].str)

		if i == len(strIndexes) - 1 && strIndexes[i].index != len(strs[0]){
			sb.WriteString(strs[0][prev:])
			prev = strIndexes[i].index
		}
	}

	return sb.String()
}

func getMin(strs ...string) int {
	min := len(strs[0])
	for i := range strs {
		current := len(strs[i])
		if min > current {
			min = current
		}
	}

	return min
}

func getData() []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1024*1024*512), 1024*1024*512)
	data := make([]string, 0, 2)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func main() {
	data := getData()
	writeData(solve(data...))
}





