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

type data struct {
	a int
	b int
}

func main() {
	var N int
	Data := []data{}
	Count := map[int]int{}
	var Max int = 0
	var Color int = 1

	N = nextInt()
	for n := 1; n <= N; n = nextInt() {
		Data[n].a = n
		Count[Data[n].a]++
		Data[n].b = nextInt()
		Count[Data[n].b]++
	}
	PreDataA := Data[1].a
	PreDataB := Data[1].b
	for i := 1; i <= N; i++ {
		if Count[Max] < Count[i] {
			Max = i
		}
	}
	for i := 0; i <= N; i++ {
		if Data[i].a == Max || Data[i].b == Max {
			fmt.Println(Color)
			PreDataA = Data[i].a
			PreDataB = Data[i].b
			Color++
		} else if Data[i].a != PreDataA && Data[i].b != PreDataB {
			Color = 1
			fmt.Println(Color)
			Color++
		} else if PreDataA == Data[i].a || PreDataB == Data[i].b {
			fmt.Println(Color)
			Color++
		}
	}

}
