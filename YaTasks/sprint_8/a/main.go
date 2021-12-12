package main

import (
	"bufio"
	"os"
	"strings"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func getData() []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		return strings.Split(scanner.Text(), " ")
	}

	return nil
}


func main() {

	data := getData()

	sb := strings.Builder{}
	for i := len(data) - 1; i >= 0; i-- {
		sb.WriteString(data[i])
		sb.WriteString(" ")
	}

	writeData(sb.String())

}
