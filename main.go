package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(Generator(n))
	fmt.Println(Solution(n))

}

func Generator(n int) (arr []int) {
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return
}

func Solution(n int) (l int) {
	var arr = strings.Split(strconv.FormatInt(int64(n), 2), "1")
	for _, s := range arr {
		var ls = len(s)
		if ls > l {
			l = ls
		}
	}
	return
}
