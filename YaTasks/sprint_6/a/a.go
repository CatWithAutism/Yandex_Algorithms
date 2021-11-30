package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	nextVertex *Vertex
	nodeNumber int
}

//region Basic

func readVertexes() []*Vertex {
	var vertexCount, edgeCount int
	fmt.Scanf("%d %d", &vertexCount, &edgeCount)
	vertexCount++
	edgeCount++

	vertexes := make([]*Vertex, vertexCount)
	for i := 1; i < edgeCount; i++ {
		var from, to int
		fmt.Scanf("%d %d", &from, &to)

		currentVertex := vertexes[from]
		if currentVertex == nil {
			currentVertex = &Vertex{
				nextVertex: &Vertex{
					nodeNumber: to,
				},
				nodeNumber: from,
			}
			vertexes[from] = currentVertex
			continue
		}

		for currentVertex.nextVertex != nil {
			currentVertex = currentVertex.nextVertex
		}

		currentVertex.nextVertex = &Vertex{
			nextVertex: nil,
			nodeNumber: to,
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

//endregion

func main() {
	vertexArray := readVertexes()
	sb := strings.Builder{}
	for i := 1; i < len(vertexArray); i++ {
		countOfWays, vertexWithWays := getVertexInfo(vertexArray[i])
		if countOfWays > 0 {
			sb.WriteString(fmt.Sprintf("%d ", countOfWays) + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vertexWithWays)), " "), "[]") + "\n")
		} else {
			sb.WriteString("0\n")
		}
	}

	writeData(sb.String())
}
