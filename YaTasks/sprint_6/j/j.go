package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

Сколько извращений лишь бы реализацию стека сюда не добавлять. Ужс

*/

const (
	white = iota
	gray  = iota
	black = iota
)

type Vertex struct {
	nodeNumber int
	color      int
}

//region Basic

func readVertexes(data []int, vertexCount, edgeCount int) [][]Vertex {
	vertexCount++

	vertexes := make([][]Vertex, vertexCount)
	for i := range vertexes {
		vertexes[i] = make([]Vertex, 0)
	}

	for i := 0; i < edgeCount*2; i += 2 {
		from, to := data[i], data[i+1]
		vertexes[from] = append(vertexes[from], Vertex{
			nodeNumber: to,
		})
	}

	for i := range vertexes {
		sort.Slice(vertexes[i], func(first, second int) bool {
			return vertexes[i][first].nodeNumber < vertexes[i][second].nodeNumber
		})
	}

	return vertexes
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

func TopSort(colorArray, prevResult []int, vertexArray [][]Vertex, currentWay, startFrom int) ([]int, int) {
	colorArray[currentWay] = gray

	ways := vertexArray[currentWay]

	for _, way := range ways {
		if colorArray[way.nodeNumber] == white {
			prevResult, startFrom = TopSort(colorArray, prevResult, vertexArray, way.nodeNumber, startFrom)
		}
	}

	colorArray[currentWay] = black

	prevResult[startFrom] = currentWay
	startFrom--
	return prevResult, startFrom
}

func readInts() []int {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	vertex, edges, i := 0, 0, 0
	fmt.Scanf("%d %d", &vertex, &edges)
	data := make([]int, edges*2+3)

	data[i] = vertex
	i++

	data[i] = edges
	i++

	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), " ")
		for _, strNumber := range strNumbers {
			number, err := strconv.Atoi(strNumber)
			if err == nil {
				data[i] = number
				i++
			}
		}
	}

	return data
}

func main() {
	data := readInts()
	countOfVertex := data[0]
	countOfEdges := data[1]
	data = data[2:len(data)]

	vertexArray := readVertexes(data, countOfVertex, countOfEdges)
	colorArray := make([]int, len(vertexArray))
	//topSorted := make([]int, 0)
	startFrom := len(vertexArray) - 1
	resultArray := make([]int, len(vertexArray))

	for i := len(vertexArray) - 1; i > 0; i-- {
		if colorArray[i] == white {
			_, startFrom = TopSort(colorArray, resultArray, vertexArray, i, startFrom)
		}
	}

	writeData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(resultArray[1:])), " "), "[]"))
}
