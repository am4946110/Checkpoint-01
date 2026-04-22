package main

import (
	"fmt"
	"strconv"
)

func safeAtoi(s string) int{
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0

	}
	return i
}

func basicatoi2(){
	fmt.Println(safeAtoi("1234"))
	fmt.Println(safeAtoi("1234abc"))
	fmt.Println(safeAtoi("abc1234"))
	fmt.Println(safeAtoi("abc"))
	fmt.Println(safeAtoi("100"))
	fmt.Println(safeAtoi(""))
}