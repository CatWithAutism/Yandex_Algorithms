package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//region Basic

type StartEnd struct {
	start  float64
	end    float64
	length float64
}

func getData() []StartEnd {
	countOfDays := 0
	fmt.Scanf("%d", &countOfDays)

	costs := make([]StartEnd, 0, countOfDays)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), " ")
		val, _ := strconv.ParseFloat(splittedString[0], 64)
		val1, _ := strconv.ParseFloat(splittedString[1], 64)
		costs = append(costs, StartEnd{
			start:  val,
			end:    val1,
			length: val1 - val,
		})
	}

	return costs
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func main() {
	data := getData()
	sort.Slice(data, func(i, j int) bool {
		if data[i].end < data[j].end {
			return true
		}

		if data[i].end > data[j].end {
			return false
		}

		return data[i].start < data[j].start
	})

	counter := 1
	buffer := make([]StartEnd, 0)
	currentEvent := data[0]
	buffer = append(buffer, currentEvent)
	for i := 1; i < len(data); i++ {
		if data[i].start >= currentEvent.end {
			currentEvent = data[i]
			buffer = append(buffer, currentEvent)
			counter++
		}
	}

	writeData(strconv.Itoa(counter) + "\n")
	for i := range buffer {
		writeData(fmt.Sprintf("%s %s\n", strconv.FormatFloat(buffer[i].start, 'f', -1, 64), strconv.FormatFloat(buffer[i].end, 'f', -1, 64)))
	}
}
