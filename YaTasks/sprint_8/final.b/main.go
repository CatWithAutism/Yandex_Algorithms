/*
https://contest.yandex.ru/contest/26133/run-report/61312625/

Самая сложная задача из всех что были.
Для начала взял trie в его классическом понимании и модифицировал для того, чтобы он мне возвращал статус.
Статус - это текущий ответ, который дал trie по подстроке. Он может быть
NO_WORDS - нету даже такого префикса
HAS_PREFIX - когда есть префикс, но это не полное слово
IS_WORD - когда имеем полное слово
Это сделано для оптимизации запросов к нему. Чтобы мы не перебирали подстроки дальше, когда префиксов таких даже нету.
Далее используем DP из прошлых лекций.
W - количество слов
L - их длинна
N - длинна делимой строки

Построение Trie O(W*L)
В худшем случае имеем (L*N)*N

*/
package main

import (
	"bufio"
	"os"
)

//region Basic

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data)
	writer.Flush()
}

//endregion
const (
	NO_WORDS      = iota
	HAS_PREFIX    = iota
	IS_WORD       = iota
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) int {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return NO_WORDS
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return IS_WORD
	}
	return HAS_PREFIX
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

func solve(str string, tr *trie) bool {
	if len(str) == 0 {
		return true
	}

	n := len(str) + 1
	dp := make([]bool, n)
	dp[0] = true

	for i := 1; i < n; i++ {
		if !dp[i] && tr.find(str[0:i]) == IS_WORD {
			dp[i] = true
		}

		if dp[i] {
			if i == n-1 {
				return true
			}

			for j := i + 1; j < n; j++ {
				status := tr.find(str[i:j])
				if status == NO_WORDS {
					break
				}

				if !dp[j] && status == IS_WORD {
					dp[j] = true
				}

				if j == n-1 && dp[j] {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	data := getData(1)
	mainStr, words := data[0], data[2:]
	trieCollection := initTrie()
	for i := range words {
		trieCollection.insert(words[i])
	}
	if solve(mainStr, trieCollection) {
		writeData("YES")
	} else {
		writeData("NO")
	}
}
