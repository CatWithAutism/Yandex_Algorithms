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

Это не решение, а кусок кода.
Надеюсь никто никогда это использовать не будет.
Если ты читаешь это, то не стоит даже время тратить.
У меня не было времени это в порядок приводить и код обвешал фичами с прошлых заданий.

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

func createGraph(data []int, vertexCount, edgeCount int) [][]Vertex {
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

		vertexes[to] = append(vertexes[to], Vertex{
			nodeNumber: from,
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

func BFS(colorArray, previousArray, distanceArray []int, vertexArray [][]Vertex, currentWay int) []int {
	colorArray[currentWay] = gray
	distanceArray[currentWay] = 0
	queue := []int{currentWay}
	path := make([]int, 0)

	for len(queue) > 0 {
		currentWay = queue[0]
		queue = queue[1:]
		ways := vertexArray[currentWay]
		path = append(path, currentWay)
		for i := range ways {
			if colorArray[ways[i].nodeNumber] == white {
				distanceArray[ways[i].nodeNumber] = distanceArray[currentWay] + 1
				previousArray[ways[i].nodeNumber] = currentWay
				colorArray[ways[i].nodeNumber] = gray
				queue = append(queue, ways[i].nodeNumber)
			}
		}

		colorArray[currentWay] = black
	}

	return path
}

func readEdges(count int) []int {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	data := make([]int, count)
	i := 0

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
	countOfVertex, countOfEdges := 0, 0
	fmt.Scanf("%d %d", &countOfVertex, &countOfEdges)

	data := readEdges(countOfEdges*2 + 1)

	startVertex := data[len(data)-1]
	vertexArray := createGraph(data[:len(data)-1], countOfVertex, countOfEdges)

	colorArray := make([]int, len(vertexArray))
	previousArray := make([]int, len(vertexArray))
	distanceArray := make([]int, len(vertexArray))

	result := BFS(colorArray, previousArray, distanceArray, vertexArray, startVertex)
	writeData(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))

}
