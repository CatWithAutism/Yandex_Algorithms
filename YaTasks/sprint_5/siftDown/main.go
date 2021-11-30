package main

func siftDown(heap []int, idx int) int {
	leftIndex, rightIndex, heapSize := 2*idx, 2*idx+1, len(heap)-1

	if heapSize < leftIndex {
		return idx
	}

	largest := 0
	if rightIndex <= heapSize && heap[leftIndex] < heap[rightIndex] {
		largest = rightIndex
	} else {
		largest = leftIndex
	}

	if heap[idx] < heap[largest] {
		heap[idx], heap[largest] = heap[largest], heap[idx]

		return siftDown(heap, largest)
	}

	return idx
}

func test() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		panic("WA")
	}
}

func main() {
	test()
}
