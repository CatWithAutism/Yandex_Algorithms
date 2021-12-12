/*

https://contest.yandex.ru/contest/25070/run-report/60104785/
O(V+E)

Развернул синие ребра в обратную сторону чтобы найти цикл(и тут нам цвет ребер становится уже не важен).
Получается что красные ребра идут от меньших к большим, а синие от больших к меньшим.
Поиск цикла реализован DFS из теории + одно условие внутри.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	White = iota
	Gray  = iota
	Black = iota
	Red   = 82
	Blue  = 66
)

type Edge struct {
	to    int16
	color int16
}

//region Basic

func getEdge(currentVertex int, text string) map[int][]Edge {
	roads := make(map[int][]Edge, 0)

	for index, val := range text {
		to := currentVertex + index + 1
		if val == Red {
			roads[currentVertex] = append(roads[currentVertex], Edge{
				to:    int16(to),
				color: Red,
			})
		} else {
			roads[to] = append(roads[to], Edge{
				to:    int16(currentVertex),
				color: Blue,
			})
		}
	}

	return roads
}

func getVertexes() [][]Edge {
	vertexCount := 0
	fmt.Scanf("%d", &vertexCount)
	vertexCount++

	edges := make([][]Edge, vertexCount)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for i := 1; scanner.Scan(); i++ {
		roads := getEdge(i, scanner.Text())
		for key, val := range roads {
			edges[key] = append(edges[key], val...)
		}
	}

	return edges
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func DFS(colorArray []int8, edgeArray [][]Edge, currentWay int16) bool {
	colorArray[currentWay] = Gray
	ways := edgeArray[currentWay]
	gotResult := false

	for _, way := range ways {
		if colorArray[way.to] == White {
			if DFS(colorArray, edgeArray, way.to) {
				return true
			}
		} else if colorArray[way.to] == Gray {
			return true
		}
	}

	colorArray[currentWay] = Black

	return gotResult
}

func main() {
	edges := getVertexes()
	lenEdges := len(edges)

	colorArray := make([]int8, lenEdges)
	for i := 1; i < lenEdges; i++ {
		if colorArray[i] == White {
			if DFS(colorArray, edges, int16(i)) {
				writeData("NO")
				return
			}
		}
	}

	writeData("YES")
}
