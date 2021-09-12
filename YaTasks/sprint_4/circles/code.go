package main

import (
	"bufio"
	"os"
)

func main(){
	strings := readStrings()[1:]
	unique := make(map[string]int, 0)
	for _, str := range strings {
		if _, ok := unique[str]; ok {
			continue
		}

		unique[str] = 0
		writeData(str)
	}
}

//region Basic
func readStrings() []string{
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
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