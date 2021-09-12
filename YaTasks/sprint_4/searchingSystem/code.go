/*
	Объясню сразу пока не забыл как это работает.
	https://contest.yandex.ru/contest/24414/run-report/52851884/
	Для голанга 1.16 лимит времени 4 секунды, когда 1.14 - 6 секунд
	Последний тест еле пролезает за 5 секунд.

	<!-- Принцип работы -->
	Он получает строки(документы) на вход.
	Создает общую мапу слов, которая индексирует слова в строках по принципу map[word][]{id, count}
	Где word - слово, id - номер документа во вхождении и count количество слов для текущего документа
	В таком случае мы можем получить количество слов в нужном нам документе дернув его по ID.
	Само собой классно, что мы не храним все эти строки целиком, а только уникальные слова.

	Далее берем поисковый запрос и разбиваем его на уникальные слова(писал бы это я на шарпе, сделал бы через HashSet чес слово) через мапу.
	Заводим большой словарик map[id]count где id - номер документа, а count количество релевантных слов в нем.
	Начинаем перечислять уникальные слова для поисковой фразы.
	Вспоминаем про наш царский словарик, который мы создали из всех доступных слов и дергаем из него значения.
	Дергаем значения по принципу howmanywordscanyousave[word] и приплюсовываем значения count для значения map[howmanywordscanyousave[word].id] += count

	Иными словами мы пробегаемся по всем уникальным словам поискового запроса и достаем для них значения из словаря,
	где мы храним информацию о том, где эти слова содержаться. Набиваем исходный словарик и отправляем его обратно.
	Мапим его в массив(т.к. голанг не позволяет использовать сортированные словари) и сортируем по нескольким условиям

	???
	А все, тут конец, история оборвалась на самом интересном.

	<!-- Сложность -->
	Тут на самом деле очень сложно оценить сложность работы, ибо, чтобы перебрать все документы это минимум O(N) без создания самого словаря
	А поисковая система работает, наверное, за O(1) * wordCount + сколько-то там, что мы тратим на перебор мест, где встречается значение.
	Там наверное надо оценивать худший и лучший, но я как-то здраво оценить это не могу.
	Вроде как не квадрат и то радует.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IndexingElement struct {
	id    int
	count int
}

type RelevantElement struct {
	id    int
	count int
}

func main() {
	docs, phrases := getOrthodoxData(readStrings())
	hashMap := createHashMap(docs)
	for _, phrase := range phrases {
		relevant := getRelevant(hashMap, phrase)

		relevantArr := make([]RelevantElement, len(relevant))
		counter := 0
		for id, count := range relevant {
			relevantArr[counter] = RelevantElement{
				id:    id,
				count: count,
			}

			counter++
		}

		sort.Slice(relevantArr, func(i, j int) bool {
			if relevantArr[i].count == relevantArr[j].count {
				return relevantArr[i].id < relevantArr[j].id
			}
			return relevantArr[i].count > relevantArr[j].count
		})

		sizeOf := len(relevant)
		if sizeOf > 5 {
			sizeOf = 5
		}

		for i := 0; i < sizeOf; i++ {
			fmt.Printf("%d ", relevantArr[i].id)
		}
		fmt.Print("\n")
	}
}

func getRelevant(hashMap map[string][]IndexingElement, phrase string) map[int]int {

	pWords := strings.Fields(phrase)
	phraseWords := make(map[string]bool)
	for _, word := range pWords {
		phraseWords[word] = true
	}

	relElements := make(map[int]int)
	for phraseWord, _ := range phraseWords {
		if indexedElements, existing := hashMap[phraseWord]; existing {
			for _, element := range indexedElements {
				if count, existing := relElements[element.id]; existing {
					relElements[element.id] = count + element.count
				} else {
					relElements[element.id] = element.count
				}
			}
		}
	}

	return relElements
}

func createHashMap(docs []string) map[string][]IndexingElement /* у меня щас глаза вытекут */ {
	wordsMap := make(map[string][]IndexingElement)
	for i, doc := range docs {
		for word, count := range wordCount(doc) {
			newElement := IndexingElement{
				id:    i + 1,
				count: count,
			}

			if iElement, existing := wordsMap[word]; existing {
				iElement = append(iElement, newElement)
				wordsMap[word] = iElement
			} else {
				wordsMap[word] = []IndexingElement{newElement}
			}
		}
	}

	return wordsMap
}

func wordCount(str string) map[string]int {
	wordArr := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordArr {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

//region Basic

func getOrthodoxData(data []string) ([]string, []string) {
	docsCount, _ := strconv.Atoi(data[0])
	docs := data[1 : docsCount+1]
	phrases := data[docsCount+2:]

	return docs, phrases
}

func readStrings() []string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 512 * 15625
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)
	var data []string

	for scanner.Scan() {
		rawString := scanner.Text()
		data = append(data, rawString)
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}

//endregion
