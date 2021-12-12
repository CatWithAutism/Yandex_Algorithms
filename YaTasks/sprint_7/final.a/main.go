/*

https://contest.yandex.ru/contest/25597/run-report/60468778/

Я такое использовал уже, правда дамерау-левенштейна. Мне нужно было вычислять разницу между названиями треков в ВК и Спотифая :)
Не совсем тривиальная задача, когда в одном у тебя может быть Louna - Полюса, а в другом L0una BassBoostedSmesharikiEdition - п0люса
Ну, это все шутки ^^^ :)

Сложность: как гласит википедия  (log(n)) ^ O(1/ϵ) ϵ > 0
Память: в случае если мы храним матрицу то O(n*m)
Как я понял хранить всю матрицу не совсем обязательно, мы можем обойтись только предпоследней вычисляемой строкой этой матрицы.
Тут можно глянуть более наглядно - https://www.youtube.com/watch?v=MiqoA-yF-0M где-то с 10-й минуты

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion

func solve(first, second string) int {
	//откидываем базовые случаи
	if len(first) == 0 {
		return len(first)
	} else if len(second) == 0 {
		return len(second)
	} else if first == second {
		return 0
	}

	runeArray1 := []rune(first)
	runeArray2 := []rune(second)

	runeLen1 := len(runeArray1)
	runeLen2 := len(runeArray2)

	ops := make([]uint16, runeLen1+1)

	for i := 1; i < len(ops); i++ {
		ops[i] = uint16(i)
	}

	for i := 1; i <= runeLen2; i++ {
		prev := uint16(i)
		for j := 1; j <= runeLen1; j++ {
			if runeArray2[i-1] != runeArray1[j-1] {
				ops[j-1] = min(ops[j]+1, min(ops[j-1]+1, prev+1))
			}
			ops[j-1], prev = prev, ops[j-1]
		}
		ops[runeLen1] = prev
	}
	return int(ops[runeLen1])
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var str1 string
	fmt.Scanf("%s", &str1)

	var str2 string
	fmt.Scanf("%s", &str2)

	writeData(strconv.Itoa(solve(str1, str2)))
}
