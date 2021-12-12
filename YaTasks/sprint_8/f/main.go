package main

import (
	"bufio"
	"os"
	"sort"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

type Helper struct {
	key   string
	value int
}

func solve(strs ...string) string {
	strs = strs[1:]

	strMap := make(map[string]int, 0)
	for i := range strs {
		if val, ok := strMap[strs[i]]; ok {
			val++
			strMap[strs[i]] = val
		} else {
			strMap[strs[i]] = 1
		}
	}

	keys := make([]Helper, 0)
	for key, val := range strMap {
		keys = append(keys, Helper{key: key, value: val})
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].value > keys[j].value
	})

	resultArr := make([]string, 0)
	element := keys[0].value
	for i := 0; i < len(keys); i++ {
		if keys[i].value != element {
			break
		}

		resultArr = append(resultArr, keys[i].key)
	}

	sort.Strings(resultArr)

	return resultArr[0]
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
