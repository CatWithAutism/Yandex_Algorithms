package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//region Basic

func readVertexes() [][]int {
	var vertexCount, edgeCount int
	fmt.Scanf("%d %d", &vertexCount, &edgeCount)
	edgeCount++
	vertexCount++

	vertexArray := make([][]int, vertexCount)
	for i := range vertexArray {
		vertexArray[i] = make([]int, vertexCount)
	}

	var from, to int
	for i := 1; i < edgeCount; i++ {
		fmt.Scanf("%d %d", &from, &to)
		vertexArray[from][to] = 1
	}

	return vertexArray
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

func getVertexInfo(vertexArray []int) (int, []int) {
	countOfWays := 0
	vertexWithWays := make([]int, 0)
	for i := range vertexArray {
		if vertexArray[i] == 1 {
			countOfWays++
			vertexWithWays = append(vertexWithWays, i)
		}
	}

	return countOfWays, vertexWithWays
}

//endregion

func main() {
	vertexArray := readVertexes()
	sb := strings.Builder{}
	for i := 1; i < len(vertexArray); i++ {
		sb.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vertexArray[i][1:])), " "), "[]") + "\n")
	}

	writeData(sb.String())
}
