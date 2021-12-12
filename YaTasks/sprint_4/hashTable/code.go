/*
	https://contest.yandex.ru/contest/24414/run-report/52992059/
	У 19 теста есть хвост, где записаны лишние данные

	<!-- Принцип работы -->
	Обычная хеш таблица.
	Реализация такая же как и в теории.
	Коллизии решаю цепочками.
	Номер корзины определяют умножением и взял за a обратное золотому сечению.

	Для корзин сделал ссылки на структуры чтобы он не сжирал память просто потому-что

	<!-- Сложность -->
	O(1) лучший
	O(N) худший

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type KeyValPair struct {
	key   int
	value int
	next  *KeyValPair
}

type HashTable struct {
	table []*KeyValPair
	size  int
}

func NewHashTable(tableSize int) HashTable {
	table := make([]*KeyValPair, tableSize)
	return HashTable{
		table: table,
		size:  tableSize,
	}
}

const alpha = 0.6180339887

func (hashTable *HashTable) computeHash(key int) int {
	hA := float64(key) * alpha
	return int(float64(hashTable.size) * (hA - float64(int(hA))))
}

func (hashTable *HashTable) Set(key int, value int) {
	index := hashTable.computeHash(key)
	hashNode := hashTable.table[index]
	node := KeyValPair{key, value, nil}
	if hashNode == nil {
		hashTable.table[index] = &node
		return
	}

	previousPair := &KeyValPair{}
	for hashNode != nil {
		if hashNode.key == key {
			hashNode.value = value
			return
		}
		previousPair = hashNode
		hashNode = hashNode.next
	}
	previousPair.next = &node
}

func (hashTable *HashTable) Get(key int) (int, bool) {
	index := hashTable.computeHash(key)
	currentPair := hashTable.table[index]
	for currentPair != nil {
		if currentPair.key == key {
			return currentPair.value, true
		}
		currentPair = currentPair.next
	}
	return 0, false
}

func (hashTable *HashTable) Delete(key int) (int, bool) {
	index := hashTable.computeHash(key)
	currentPair := hashTable.table[index]
	if currentPair == nil {
		return 0, false
	} else if currentPair.key == key {
		hashTable.table[index] = currentPair.next
		return currentPair.value, true
	}

	previousPair := currentPair
	currentPair = currentPair.next
	for currentPair != nil {
		if currentPair.key == key {
			previousPair.next = currentPair.next
			return currentPair.value, true
		}
		previousPair = currentPair
		currentPair = currentPair.next
	}
	return 0, false
}

func main() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	builder := strings.Builder{}
	hashTable := NewHashTable(100003)

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())

		key, _ := strconv.Atoi(fields[1])
		if fields[0] == "get" {
			if val, existing := hashTable.Get(key); existing {
				builder.WriteString(strconv.Itoa(val) + "\n")
				continue
			}
		} else if fields[0] == "put" {
			val, _ := strconv.Atoi(fields[2])
			hashTable.Set(key, val)
			continue
		} else {
			if deletedVal, existing := hashTable.Delete(key); existing {
				builder.WriteString(strconv.Itoa(deletedVal) + "\n")
				continue
			}
		}

		builder.WriteString("None" + "\n")
	}

	writeData(builder.String())
}

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion
