package main

import (
	"bufio"
	"fmt"
	"os"
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

func getData(arrLen int) []int {
	banknotes := make([]int, 0, arrLen)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		splittedString := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(splittedString); i++ {
			val, _ := strconv.Atoi(splittedString[i])
			banknotes = append(banknotes, val)
		}
	}

	return banknotes
}

func main() {
	summ, bankCount := 0, 0
	fmt.Scanf("%d", &summ)
	fmt.Scanf("%d", &bankCount)

	banknotes := getData(bankCount)
	dp := make([]int,summ+1)
	dp[0] = 1
	for i:=0;i<bankCount;i++ {
		start := banknotes[i]
		for j:=start;j<=summ;j++ {
			dp[j] += dp[j-start]
		}
	}

	writeData(strconv.Itoa(int(dp[summ])))
}
