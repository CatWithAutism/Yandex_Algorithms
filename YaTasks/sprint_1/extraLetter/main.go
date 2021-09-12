package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const separator = " "

func main(){
	Init()
	str1 := ReadLine()
	str2 := ReadLine()

	var slice1 RuneSlice
	slice1 = []rune(str1)
	sort.Sort(slice1)

	var slice2 RuneSlice
	slice2 = []rune(str2)
	sort.Sort(slice2)

	length1 := len(slice1)
	length2 := len(slice2)
	maxLength := Max(length1, length2)
	for i := 0; i < maxLength; i++ {
		if i == length1{
			WriteData(string(slice2[length2 - 1]))
			break
		} else if slice1[i] != slice2[i]{
			WriteData(string(slice2[i]))
			break
		}
	}

	//WriteData(strings.Trim(strings.Join(strings.Split(strconv.Itoa(xNumber + kNumber), ""), separator), "[]"))
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(reader)

func Init(){
	scanner.Split(bufio.ScanLines)
	buffer := make([]byte, 600000)
	scanner.Buffer(buffer, 600000)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func ReadIntegers() []int{
	var data []int

	for scanner.Scan(){
		strNumbers := strings.Split(scanner.Text(), separator)
		for _, strNumber := range strNumbers {
			number, _ := strconv.Atoi(strNumber)
			data = append(data, number)
		}
	}

	return data
}

func ReadLine() string{
	scanner.Scan()
	return scanner.Text()
}


func WriteData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func unique(intSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}