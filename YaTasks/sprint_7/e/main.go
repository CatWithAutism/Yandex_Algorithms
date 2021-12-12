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
	sort.Ints(banknotes)
	dp := make([]int, summ+1)
	dp[0] = 1

	writeData(strconv.Itoa(solve(summ, banknotes)))
}

func solve(n int, coins []int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	max := 9999999999
	for i := 1; i <= n; i++ {
		dp[i] = max
	}

	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
