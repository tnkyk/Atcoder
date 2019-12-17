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

	if s == "Sunny" {
		fmt.Println("Cloudy")
	} else if s == "Cloudy" {
		fmt.Println("Rainy")
	} else if s == "Rainy" {
		fmt.Println("Sunny")
	}

}
