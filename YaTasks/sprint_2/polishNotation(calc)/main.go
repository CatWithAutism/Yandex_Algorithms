package main

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	initReader()
	values := readStrings()
	stack := NewStack()

	for i := 0; i < len(values); i++ {
		if val, err := strconv.Atoi(values[i]); err == nil {
			stack.Push(val)
		} else {
			second, err := stack.Pop()
			if err != nil {
				return
			}

			first, err := stack.Pop()
			if err != nil {
				return
			}
			stack.Push(calculate(first.(int), second.(int), values[i]))
		}
	}

	pop, err := stack.Pop()
	if err != nil {
		return
	}
	writeData(strconv.Itoa(pop.(int)))
}

func calculate(firstNumber int, secondNumber int, operand string) int {
	switch strings.TrimSpace(operand) {
	case "*":
		return firstNumber * secondNumber
	case "/":
		return int(math.Floor(float64(firstNumber) / float64(secondNumber)))
	case "+":
		return firstNumber + secondNumber
	case "-":
		return firstNumber - secondNumber
	}

	return -1
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader() {
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readStrings() []string {
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

//region Stack

type Stack struct {
	buffer  []interface{}
	pointer int
}

func NewStack() *Stack {

	return &Stack{
		buffer:  make([]interface{}, 0),
		pointer: 0,
	}
}

func (stack *Stack) Push(val interface{}) {
	//Это на случай если pop размер номинально, но не фактически
	if len(stack.buffer) == stack.pointer {
		stack.buffer = append(stack.buffer, val)
	} else {
		stack.buffer[stack.pointer] = val
	}

	stack.pointer++
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.pointer == 0 {
		return "", errors.New("empty")
	}

	stack.pointer--
	return stack.buffer[stack.pointer], nil
}

//endregion
