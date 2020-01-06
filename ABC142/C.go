package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var N int = nextInt()
	fmt.Println(N)
	Ai := map[int]int{}
	var n int = 0
	for count := 1; count <= N; count++ {
		n = nextInt()
		if n == 0 {
			break
		}
		Ai[count] = n
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if Ai[j] == i {
				fmt.Print(j)
				if i == N {
					fmt.Print("\n")
				}
				break
			}
		}
	}
}

/*
今回の課題は貪欲法に基づく問題だった
コツとしては、そもそも１が数列に存在しないと始まらないというところに着目できるかどうかである。
*/
