/*
https://contest.yandex.ru/contest/26133/run-report/61150608/

Рекурсивно распаковываем используем стринг билдеры чтобы избежать лишних выделений памяти.
Пробегаемся в поисках нашего постфикса i, j где i это буква, а j это слово

По сложности:
Распаковку не имею понятия как оценить, ибо все зависит от строки, которая пришла к нам на вход.
Может быть и строка без запаковки, а может быть строка запакованная int32.Max раз

Поиск префикса k * min(n), где k количество строк, а min(n) это минимальное вхождение строки из всех.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func unpackString(val string, count int) string {
	sb := strings.Builder{}
	lastIndex := 0
	opening, closing := uint8('['), uint8(']')
	opened, closed := 0, 0
	for i := 0; i < len(val); i++ {
		if unicode.IsDigit(rune(val[i])) {
			sb.WriteString(val[lastIndex:i])
			for j := i; j < len(val); j++ {
				if val[j] == opening {
					opened++
				} else if val[j] == closing {
					closed++
				}

				if opened != 0 && closed != 0 && opened == closed {
					sb.WriteString(unpackString(val[i+2:j], int(val[i]-48)))
					lastIndex, i = j+1, j
					opened, closed = 0, 0
					break
				}
			}
		}
	}

	sb.WriteString(val[lastIndex:])

	old := sb.String()
	for i := 0; i < count-1; i++ {
		sb.WriteString(old)
	}

	return sb.String()
}

func getData(length int) []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1024*1024*512), 1024*1024*512)
	data := make([]string, 0, length)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func getMin(data []string) int {
	min := len(data[0])
	for i := 1; i < len(data); i++ {
		if len(data) < min {
			min = len(data)
		}
	}

	return min
}

func main() {
	length := 0
	fmt.Scanf("%d", &length)
	data := getData(length)

	for i := 0; i < len(data); i++ {
		data[i] = unpackString(data[i], 1)
	}

	sort.Slice(data, func(i, j int) bool {
		return len(data[i]) < len(data[j])
	})

	min := len(data[len(data)-1]) - 1

	till := 0
	for i := 0; i < min; i++ {
		r := data[0][i]
		for j := 1; j < len(data); j++ {
			if r != data[j][i] {
				till = i
				break
			}
		}

		if till != 0 {
			break
		}
	}

	if till == 0 {
		writeData(data[0])
	} else {
		writeData(data[0][:till])
	}

}
