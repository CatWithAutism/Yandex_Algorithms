/*
Этот ужас я сначала сделал на слайсах.
Но из-за того что мне голанг не родной и я не особо знаю как фурычат слайсы под капотом, он работал десять тысяч и один год.
Пытался перекидывать опорный элемент в конец и от него уже сортировать
https://contest.yandex.ru/contest/23815/run-report/52628958/

Задание переделал т.к. он на слайсах работал в 10? раз хуже
Итоговый вариант:
https://contest.yandex.ru/contest/23815/run-report/52658711/

<!-- Принцип работы -->
Берет опорный элемент с конца и рекурсивно пробегается по массиву совершая перестановки
ориентируясь на опорный элемент и функцию сравнения поданную на вход.

Функция сравнения проверяет больше ли первый элемент второго отталкиваясь от условий задачи.

<!-- Сложность -->
Описанным выше способом он проходит весь массив
Итого сложность: O(N*log(N))
O(n) на хранение, где N количество записей

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	members := readMembers()
	start := time.Now()

	inPlaceQuickSort(members, 0, len(members) - 1, greaterMember)

	sprintf := fmt.Sprintf("%s", time.Since(start))

	var sb strings.Builder
	for _, member := range members {
		sb.WriteString(member.Name)
		sb.WriteString("\r\n")
	}

	writeData(sb.String())
	println(sprintf)
}

type Member struct {
	Name          string
	ResolvedTasks int
	Penalty       int
}

func inPlaceQuickSort(arr []Member, start, end int, greater func(first, second *Member) bool)[]Member{
	if start > end{
		return arr
	}

	pivot := arr[end]
	left := start
	for i := start; i <= end; i++ {
		if greater(&arr[i], &pivot) {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[end] = arr[end], arr[left]

	inPlaceQuickSort(arr, start, left - 1, greater)
	inPlaceQuickSort(arr, left + 1, end, greater)

	return arr
}

func greaterMember(first, second *Member) bool{
	if first.ResolvedTasks == second.ResolvedTasks {
		if first.Penalty == second.Penalty {
			return first.Name < second.Name
		}
		return first.Penalty < second.Penalty
	}

	return first.ResolvedTasks > second.ResolvedTasks
}

//region Basic
func readMembers() []Member{
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var data []Member
	counter := 0

	for scanner.Scan(){
		values := strings.Split(scanner.Text(), " ")
		if len(values) == 1{
			size, fError := strconv.Atoi(values[0])
			if fError != nil{
				panic(fError)
			}
			data = make([]Member, size)
		} else if len(values) == 3{
			resolvedTasks, fError := strconv.Atoi(values[1])
			penalty, sError := strconv.Atoi(values[2])
			if sError != nil || fError != nil{
				continue
			}
			data[counter] = Member{
				Name:          values[0],
				ResolvedTasks: resolvedTasks,
				Penalty:       penalty,
			}

			counter++
		}
	}

	return data
}


func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(data)
	_ = writer.Flush()
}

//endregion