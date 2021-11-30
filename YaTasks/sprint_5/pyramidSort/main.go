/*
	10.11.2021 1:46, скажи, я успею за две недели, да? :)

	Горит дедлайн, я стараюсь следить за кодом, но чето не очень.
	Там доступен Шарп стал и я бы очень хотел писать на нем, ибо родной, но поздняк метаться.
	Не стал сильно задумываться над самой задачей и собрал Франкенштейна.

	Тут я намешал код из 4 уроков.
	1. Одна из финальных задач квик сорта, которая парсит данные
	2. SiftUp - https://contest.yandex.ru/contest/24809/run-report/56018210/
	3. SiftDown - https://contest.yandex.ru/contest/24809/run-report/55885690/
	4. И пирамидальную сортировку

	Сложность: O(n log n)
	Описано в теории и не вижу смысла сюда это копировать :)

	Посылка - https://contest.yandex.ru/contest/24810/run-report/56563716/
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	members := readMembers()
	copyMembers := members
	for i := 1; i < len(members); i++ {
		var member Member
		member, copyMembers = heapPopMax(copyMembers)
		printData(member.Name + "\r\n")
	}
}

type Member struct {
	Name          string
	ResolvedTasks int
	Penalty       int
}

//region push/pop

func heapPush(heap []Member, val Member) []Member {
	currentLength := len(heap)

	if currentLength == 0 {
		heap = make([]Member, 1)
		currentLength = 1
	}

	heap = append(heap, val)
	return siftUp(heap, currentLength)
}

func heapPopMax(heap []Member) (Member, []Member) {
	result := heap[1]
	heap[1] = heap[len(heap)-1]
	return result, siftDown(heap[:len(heap)-1], 1)
}

//endregion

//region sift

func siftUp(heap []Member, idx int) []Member {
	if idx == 1 {
		return heap
	}

	parentIndex := idx / 2
	if lessMember(heap[parentIndex], heap[idx]) {
		heap[parentIndex], heap[idx] = heap[idx], heap[parentIndex]
		return siftUp(heap, parentIndex)
	}

	return heap
}

func siftDown(heap []Member, idx int) []Member {
	leftIndex, rightIndex, heapSize := 2*idx, 2*idx+1, len(heap)-1

	if heapSize < leftIndex {
		return heap
	}

	largest := 0
	if rightIndex <= heapSize && lessMember(heap[leftIndex], heap[rightIndex]) {
		largest = rightIndex
	} else {
		largest = leftIndex
	}

	if lessMember(heap[idx], heap[largest]) {
		heap[idx], heap[largest] = heap[largest], heap[idx]

		return siftDown(heap, largest)
	}

	return heap
}

//endregion

func lessMember(first, second Member) bool {
	if first.ResolvedTasks == second.ResolvedTasks {
		if first.Penalty == second.Penalty {
			return first.Name > second.Name
		}
		return first.Penalty > second.Penalty
	}

	return first.ResolvedTasks < second.ResolvedTasks
}

//region Basic
func readMembers() []Member {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	data := make([]Member, 0)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		currentLength := len(values)
		if currentLength == 1 {
			//не интересно сколько там он нам даст
			continue
		} else if currentLength == 3 {
			resolvedTasks, fError := strconv.Atoi(values[1])
			penalty, sError := strconv.Atoi(values[2])
			if sError != nil || fError != nil {
				continue
			}

			data = heapPush(data, Member{
				Name:          values[0],
				ResolvedTasks: resolvedTasks,
				Penalty:       penalty,
			})
		}
	}

	return data
}

func printData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	_, error := writer.WriteString(data)
	if error != nil {
		panic(error)
	}

	error = writer.Flush()
	if error != nil {
		panic(error)
	}
}

//endregion
