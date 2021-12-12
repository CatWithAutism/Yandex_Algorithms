/*
https://contest.yandex.ru/contest/25597/run-report/60624045/

Неплохое объяснение нашел тут https://www.youtube.com/watch?v=n3LS2p7xoE8, но в принципе оно решается почти также как и предыдущее.
Откидываем базовое значение, когда сумма не делится на два без остатка и заполняем матрицу от половины суммы.

По памяти O(summ(N)/2*N + N)
Т.е. сама матрица + выделенная память под сам массив
Сложность O(summ(n))

Если честно, я бы еще прочитал раза два теорию, но время поджимает :(

*/

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

func summOf(arr []int) int {
	summ := 0
	for i := range arr {
		summ += arr[i]
	}

	return summ
}

func getArrayVal(points [][]bool, i, j int) *bool {
	if (i < 0 || i > len(points)-1) || (j < 0 || j > len(points[i])-1) {
		return nil
	}

	return &points[i][j]
}

func solve(arr []int, n int) bool {
	summ := summOf(arr)
	if summ%2 != 0 {
		return false
	}

	parts := make([][]bool, summ/2+1)
	for i := 0; i < len(parts); i++ {
		parts[i] = make([]bool, n+1)
		if i == 0 {
			for j := range parts[i] {
				parts[i][j] = true
			}
		}
	}

	for i := 1; i <= summ/2; i++ {
		for j := 1; j <= n; j++ {
			parts[i][j] = parts[i][j-1]
			if i >= arr[j-1] {
				parts[i][j] = parts[i][j-1] || parts[i-arr[j-1]][j-1]
			}
		}
	}

	return parts[summ/2][n]
}

func main() {
	var count int
	fmt.Scanf("%d", &count)

	data := getData(count)

	if solve(data, count) {
		writeData("True")
	} else {
		writeData("False")
	}

}
