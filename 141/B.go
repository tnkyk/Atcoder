package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}

	ans := "Yes"
	for k, v := range s {
		if (k+1)%2 == 0 {

			if string(v) == "R" {
				ans = "No"
				break
			}
		} else {

			if string(v) == "L" {
				ans = "No"
				break
			}
		}
	}
	fmt.Println(ans)
}
