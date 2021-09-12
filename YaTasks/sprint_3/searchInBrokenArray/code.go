/*
https://contest.yandex.ru/contest/23815/run-report/52620407/
Алгоритм в этой задачи как я в пятницу после бара,
раскачивается на ветру влево вправо и пытается понять где он

<!-- Принцип работы -->
В основе само собой лежит бинарный поиск, т.е. рекурсивный спуск вниз в отсортированной последовательности данных.
Задача усложнилась тем, что массив как бы отсортирован, но как бы нет.
Сначала надо найти кусок, в котором будем искать, опираясь на элементы слева и справа.
Когда нашли, то обрадоваться и продолжить бинарный поиск
Т.е. спускаться используя принцип больше меньше от опорного элемента

<!-- Сложность -->
Итого сложность: log(N)
O(n) на хранение, где N количество записей

*/

package main

func binarySearch(collection []int, element int, left int, right int) int{
	pivot := (left + right) / 2
	midVal := collection[pivot]
	leftVal := collection[left]
	rightVal := collection[right]

	if right - left < 2{
		if leftVal == element{
			return left
		} else if rightVal == element{
			return right
		}
		return -1
	}

	/*
		Этот кусок кода определяет где он находится
		отталкиваясь от крайних элементов.

		Радует, что оно хотя бы работает(но это не точно).
	 */
	if midVal > leftVal && midVal < rightVal{
		if element <= midVal{
			return binarySearch(collection, element, left, pivot)
		}

		return binarySearch(collection, element, pivot, right)
	} else if leftVal > midVal && midVal < rightVal{
 		if element <= midVal{
			return binarySearch(collection, element, left, pivot)
		}

		return binarySearch(collection, element, pivot, right)
	} else {
		if element >= midVal{
			return binarySearch(collection, element, pivot, right)
		} else if rightVal >= element{
			return binarySearch(collection, element, pivot, right)
		}

		return binarySearch(collection, element, left, pivot)
	}
}

func brokenSearch(arr []int, k int) int {
	return binarySearch(arr, k, 0, len(arr) - 1)
}

func test() {
	//arr := []int{12, 41, 122, 411, 412, 1222, 3000, 12222, 122222}
	arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	if brokenSearch(arr, 5) != 6 {
		panic("WA")
	} else {
		println("Gotcha")
	}
}
