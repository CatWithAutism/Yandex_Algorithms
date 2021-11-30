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
	nextVertex *Vertex
	nodeNumber int
	color      int
}

//region Basic

func readVertexes(data []int, vertexCount, edgeCount int) []*Vertex {
	vertexCount++

	vertexes := make([]*Vertex, vertexCount)
	for i := range vertexes {
		vertexes[i] = &Vertex{
			nodeNumber: i,
		}
	}

	for i := 0; i < edgeCount*2; i += 2 {
		from, to := data[i], data[i+1]

		currentVertex := vertexes[from]

		for currentVertex.nextVertex != nil {
			currentVertex = currentVertex.nextVertex
		}

		currentVertex.nextVertex = &Vertex{
			nextVertex: nil,
			nodeNumber: to,
		}

		currentVertex = vertexes[to]

		for currentVertex.nextVertex != nil {
			currentVertex = currentVertex.nextVertex
		}

		currentVertex.nextVertex = &Vertex{
			nextVertex: nil,
			nodeNumber: from,
		}

	}

	return vertexes
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

func getVertexInfo(vertex *Vertex) (int, []int) {
	if vertex == nil {
		return 0, nil
	}

	countOfWays := 0
	vertexWithWays := make([]int, 0)

	for vertex.nextVertex != nil {
		vertex = vertex.nextVertex

		countOfWays++
		vertexWithWays = append(vertexWithWays, vertex.nodeNumber)
	}

	return countOfWays, vertexWithWays
}

func getWays(vertexArray []*Vertex, way int) []int {
	currentVertex := vertexArray[way]
	if currentVertex == nil {
		return nil
	}

	result := make([]int, 0)
	for currentVertex.nextVertex != nil {
		currentVertex = currentVertex.nextVertex
		result = append(result, currentVertex.nodeNumber)
	}

	sort.Ints(result)
	return result
}

func DFS(colorArray, prevResult []int, vertexArray []*Vertex, currentWay int) []int {
	colorArray[currentWay] = gray
	ways := getWays(vertexArray, currentWay)
	prevResult = append(prevResult, currentWay)

	for _, way := range ways {
		if colorArray[way] == white {
			prevResult = DFS(colorArray, prevResult, vertexArray, way)
		}
	}

	colorArray[currentWay] = black

	return prevResult
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
	startVertex := data[len(data)-1]
	data = data[2 : len(data)-1]

	vertexArray := readVertexes(data, countOfVertex, countOfEdges)

	colorArray := make([]int, len(vertexArray))

	result := DFS(colorArray, make([]int, 0), vertexArray, startVertex)
	writeData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
}
