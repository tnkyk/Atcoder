package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Score struct {
	score int
}

func main() {
	var s string
	var Ai int

	if sc.Scan() {
		s = sc.Text()
	}
	ss := strings.Fields(s)
	N, _ := strconv.Atoi(ss[0])
	K, _ := strconv.Atoi(ss[1])
	Q, _ := strconv.Atoi(ss[2])
	Scores := make([]Score, N)
	for i := 1; i <= Q; i++ {
		Ai = Scantext(sc)
		Scores[Ai-1].score++
	}

	for i := 0; i < N; i++ {
		if (Q - Scores[i].score) < K {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func Scantext(sc *bufio.Scanner) (i int) {
	sc.Scan()
	var s string
	s = sc.Text()
	i, _ = strconv.Atoi(s)
	return i
}
