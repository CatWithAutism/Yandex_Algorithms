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
	maxLadder, maxJump := int64(0), int64(0)
	fmt.Scanf("%d %d", &maxLadder, &maxJump)

	ladder := make([]int64, maxLadder+1)
	ladder[0] = 0
	ladder[1] = 1

	for i := int64(2); i < int64(len(ladder)); i++ {
		if maxJump >= i {
			for j := int64(1); j < i; j++ {
				ladder[i] += ladder[j]
			}
		} else {
			for j := i - maxJump; j < i; j++ {
				ladder[i] += ladder[j]
			}
		}

		ladder[i] = ladder[i] % 1000000007
	}

	writeData(strconv.Itoa(int(ladder[maxLadder])))
}
