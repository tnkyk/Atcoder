package main

import "fmt"

var count int = 0
var N int

func main() {

	var BufA string
	var BufB string
	var A []string
	var B []string
	var ans string = "Yes"

	fmt.Scan(&N)
	fmt.Scan(&BufA)
	fmt.Scan(&BufB)

	for _, v := range BufA {
		A = append(A, string(v))
	}
	for _, v := range BufB {
		B = append(B, string(v))
	}
	
	for i := 0; i < len(A); i++ {
		if A[i] <= B[i] {
			continue
		} else {
			for j:=0;j<i;i++{
				if A[j] > B[i] {
					if A[i] >= A[j] 
					A[i] = A[j]

				}
			}
			ans = "No"
			break
		}
	}
	fmt.Println(A, B)
	fmt.Println(ans)
}
