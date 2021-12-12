package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func DFS(colorArray, entryArray, leaveArray, prevResult []int, vertexArray [][]Vertex, currentWay, time int) ([]int, int) {
	colorArray[currentWay] = gray

	entryArray[currentWay] = time
	time++

	ways := vertexArray[currentWay]
	prevResult = append(prevResult, currentWay)

	for _, way := range ways {
		if colorArray[way.nodeNumber] == white {
			prevResult, time = DFS(colorArray, entryArray, leaveArray, prevResult, vertexArray, way.nodeNumber, time)
		}
	}

	colorArray[currentWay] = black

	leaveArray[currentWay] = time
	time++

	return prevResult, time
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
	entryArray := make([]int, len(vertexArray))
	leaveArray := make([]int, len(vertexArray))

	DFS(colorArray, entryArray, leaveArray, make([]int, 0), vertexArray, 1, 0)
	for i := 1; i < len(entryArray); i++ {
		writeData(strconv.Itoa(entryArray[i]) + " " + strconv.Itoa(leaveArray[i]) + "\n")
	}
	//writeData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
}
