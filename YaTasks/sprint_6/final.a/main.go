/*
Посылка: https://contest.yandex.ru/contest/25070/run-report/60044758/

O(∣E∣⋅log∣V∣), где |E| — количество рёбер в графе, а |V| — количество вершин.

Очень долго мучался с этой задачей.

Читает ребра, вершины хранит в хештаблице.
Если вершина пройдена, то он удаляет запись из нее.
Так сделано чтобы в цикле он не во всей хештаблице искал записи с значением false(может можно лучше).
Для каждой вершины выгружаем ребра в приоритетную очередь, которая сортируется по весу этих ребер.
Таким образом мы перебираем только максимальные значения.

 */

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	from   int
	to     int
	weight int
}

//region Basic

func getEdges(countOfEdges int) []Edge {
	edges := make([]Edge, 0, countOfEdges)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), " ")
		if len(strNumbers) > 0 {
			from, _ := strconv.Atoi(strNumbers[0])
			to, _ := strconv.Atoi(strNumbers[1])
			if from == to {
				continue
			}
			weight, _ := strconv.Atoi(strNumbers[2])

			edges = append(edges, Edge{
				from:   from,
				to:     to,
				weight: weight,
			})
		}
	}

	return edges
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

func remove(s []Edge, i int) []Edge {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func getMax(spanningEdges *PriorityQueue) *Edge {
	return heap.Pop(spanningEdges).(*Edge)
}

func containsNotAdded(added map[int]bool) bool {
	return len(added) > 0
}

func getFreeRoutes(added map[int]bool, edges []Edge, spanningEdges *PriorityQueue, vertex int) []Edge {
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		if edge.from == vertex || edge.to == vertex {
			valTo, okTo := added[edge.to]
			valFrom, okFrom := added[edge.from]

			if okTo && !valTo || okFrom && !valFrom {
				heap.Push(spanningEdges, &Edge{
					from:   edges[i].from,
					to:     edges[i].to,
					weight: edges[i].weight,
				})
				edges = remove(edges, i)
				i--
			}
		}
	}

	return edges
}

func addVertex(added map[int]bool, edges []Edge, spanningEdges *PriorityQueue, vertex int) []Edge {
	added[vertex] = true
	return getFreeRoutes(added, edges, spanningEdges, vertex)
}

func findMST(edges []Edge, countOfVertex int) string {
	var maxSpanningTree int
	added := make(map[int]bool, countOfVertex)
	for i := 1; i < countOfVertex+1; i++ {
		added[i] = false
	}

	spanningEdges := make(PriorityQueue, 0)
	heap.Init(&spanningEdges)

	edges = addVertex(added, edges, &spanningEdges, 1)

	delete(added, 1)
	for containsNotAdded(added) && len(spanningEdges) > 0 {
		edge := getMax(&spanningEdges)

		valTo, okTo := added[edge.to]
		valFrom, okFrom := added[edge.from]

		if okTo && !valTo {
			maxSpanningTree += edge.weight
			edges = addVertex(added, edges, &spanningEdges, edge.to)
			delete(added, edge.to)
		} else if okFrom && !valFrom {
			maxSpanningTree += edge.weight
			edges = addVertex(added, edges, &spanningEdges, edge.from)
			delete(added, edge.from)
		}
	}

	if len(added) > 0 {
		return "Oops! I did it again"
	}

	return strconv.Itoa(maxSpanningTree)
}

func main() {
	countOfVertex, countOfEdges := 0, 0
	fmt.Scanf("%d %d", &countOfVertex, &countOfEdges)

	vertexArray := getEdges(countOfEdges)

	writeData(findMST(vertexArray, countOfVertex))
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight > pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}