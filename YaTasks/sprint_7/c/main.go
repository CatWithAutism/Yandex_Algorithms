package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EasyMoney struct {
	coeff  float32
	weight int32
	cost   int32
}

//region Basic

func getData() (int32, []EasyMoney) {
	maxWeight, nop := int32(0), int32(0)
	fmt.Scanf("%d", &maxWeight)

	coeffs := make([]EasyMoney, 0, maxWeight)

	//just skip a number
	fmt.Scanf("%d", &nop)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), " ")
		val, _ := strconv.ParseInt(splittedString[0], 10, 32)
		val1, _ := strconv.ParseInt(splittedString[1], 10, 32)
		coeffs = append(coeffs, EasyMoney{
			weight: int32(val1),
			cost:   int32(val),
		})
	}

	return maxWeight, coeffs
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func main() {
	availableWeight, data := getData()
	sort.Slice(data, func(i, j int) bool {
		if data[i].cost > data[j].cost {
			return true
		}

		if data[i].cost < data[j].cost {
			return false
		}

		return data[i].weight > data[j].weight
	})

	totalCost := int64(0)
	for i := 0; i < len(data); i++ {
		if availableWeight >= data[i].weight {
			totalCost += int64(data[i].cost) * int64(data[i].weight)
			availableWeight -= data[i].weight
		} else {
			totalCost += int64(data[i].cost) * int64(availableWeight)
			break
		}
	}

	writeData(strconv.Itoa(int(totalCost)))
}
