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

func getData(n, m int) [][]int16 {
	flowers := make([][]int16, n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		splittedString := strings.Split(scanner.Text(), "")
		flowers[i] = make([]int16, 0, m)
		for j := 0; j < len(splittedString); j++ {
			val, _ := strconv.Atoi(splittedString[j])
			flowers[i] = append(flowers[i], int16(val))
		}
	}

	return flowers
}

func getArrayVal(points [][]int16, i, j int) int16 {
	if (i < 0 || i > len(points)-1) || (j < 0 || j > len(points[i])-1) {
		return 0
	}

	return points[i][j]
}

func solve(points [][]int16, n, m int) int16 {
	dp := make([][]int16, len(points))
	for i := range points {
		dp[i] = make([]int16, len(points[i]))
	}

	dp[n-1][0] = getArrayVal(points, n-1, 0)

	for i := len(points) - 1; i >= 0; i-- {
		for j := 0; j < len(points[i]); j++ {
			dp[i][j] = max(getArrayVal(dp, i+1, j), getArrayVal(dp, i, j-1)) + points[i][j]
		}
	}

	return dp[0][len(points[0])-1]
}

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d", &n, &m)
	if n == 1 && m == 1{
		writeData("1")
		os.Exit(0)
	}

	data := getData(n, m)
	writeData(strconv.Itoa(int(solve(data, n, m))))
}

func max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
