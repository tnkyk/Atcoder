package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getScanner(fp *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 500001), 500000)
	return scanner
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

func getNextInt64(scanner *bufio.Scanner) int64 {
	i, _ := strconv.ParseInt(getNextString(scanner), 10, 64)
	return i
}

func getNextUint64(scanner *bufio.Scanner) uint64 {
	i, _ := strconv.ParseUint(getNextString(scanner), 10, 64)
	return i
}

func getNextFloat64(scanner *bufio.Scanner) float64 {
	i, _ := strconv.ParseFloat(getNextString(scanner), 64)
	return i
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout

	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	n := getNextInt(scanner)
	aa := make([]AA, n)
	for i := 0; i < n; i++ {
		a := getNextInt(scanner)
		aa[i].a = a
		//初期化
		aa[i].x = make([]int, a)
		aa[i].y = make([]int, a)
		//データの格納
		for j := 0; j < a; j++ {
			aa[i].x[j] = getNextInt(scanner) - 1
			aa[i].y[j] = getNextInt(scanner)
		}
	}

	ans := 0
	for i := 0; i < 1<<uint(n); i++ {
		sum := 0
		valid := true
		for j := 0; j < n && valid; j++ {
			for k := 0; k < aa[j].a && valid; k++ {
				if i>>uint(j)&1 == 1 {
					if i>>uint(aa[j].x[k])&1 != aa[j].y[k] {
						valid = false
					}
					continue
				}

			}
			sum += i >> uint(j) & 1
		}
		if valid && ans < sum {
			ans = sum
		}
	}

	fmt.Fprintln(writer, ans)
	writer.Flush()
}

type AA struct {
	a    int
	x, y []int
}
