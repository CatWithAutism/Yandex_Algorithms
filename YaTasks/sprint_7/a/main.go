package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//region Basic

func getData() []int {
	countOfDays := 0
	fmt.Scanf("%d", &countOfDays)

	costs := make([]int, 0, countOfDays)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), " ")
		for i := range splittedString {
			val, _ := strconv.Atoi(splittedString[i])
			costs = append(costs, val)
		}
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

	totalProfit := 0
	purchasePrice := -1

	today := data[0]
	data = data[1:]

	for len(data) > 0 {
		tomorrow := data[0]
		data = data[1:]

		if today < tomorrow && purchasePrice == -1 {
			purchasePrice = today
		} else if purchasePrice != -1 && purchasePrice < today {
			totalProfit += today - purchasePrice
			purchasePrice = -1
			if today < tomorrow && purchasePrice == -1 {
				purchasePrice = today
			}
		}

		today = tomorrow
	}

	if purchasePrice != -1 && purchasePrice < today {
		totalProfit += today - purchasePrice
		purchasePrice = -1
	}

	writeData(strconv.Itoa(totalProfit))
}
