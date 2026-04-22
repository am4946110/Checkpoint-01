package main

import (
	"fmt"
	"strconv"
)

func main() {
	textNumber := "1234"
	number, err := strconv.Atoi(textNumber)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	fmt.Println("Converted integer:", number)
}