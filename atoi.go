package main

import (
	"fmt"
	"strconv"
)

func main() {
	s2 := "1234"

	num1, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	} else {
		fmt.Println("Converted integer:", num1)
	}

	sl := "-1234"
	num2, err := strconv.Atoi(sl)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Println("Converted integer:", num2)

	s3 := "-567"
	num3, err := strconv.Atoi(s3)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	} else {
		fmt.Println("Converted integer:", num3)
	}
}