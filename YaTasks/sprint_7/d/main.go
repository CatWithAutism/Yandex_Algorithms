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

func fibonacci(n int64) int64 {
	a, b, r := int64(1), int64(1), int64(0)
	for i := int64(2); i <= n; i++ {
		r = (a + b) % 1000000007
		if i+1 <= n {
			a, b = b, r
		}
	}

	return r
}

func main() {
	fib := int64(0)
	fmt.Scanf("%d", &fib)

	writeData(strconv.Itoa(int(fibonacci(fib))))
}
