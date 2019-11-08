package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var i, j int
var ans int
var pre float64
var prei, prej int
var state float64

func main() {

	var N = bufio.NewScanner(os.Stdin)
	N.Scan()

	buf, err := strconv.Atoi(N.Text())
	if err != nil {
		fmt.Println(err)
	}
	prei = 1
	prej = buf / 1
	pre = math.Abs(float64(prei - prej))
	//3

	for i = 2; i < 1000000000000; i++ {
		if buf%i == 0 {
			j = (buf / i)
			if pre == math.Abs(float64(j-i)) || pre <= math.Abs(float64(j-i)) {
				ans = (prej - 1) + (prei - 1)
				fmt.Println(ans)
				break
			}
			state = float64(j - i)
			pre = math.Abs(state)
			prei = i
			prej = j
		}
	}

}
