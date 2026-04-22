package main

import (
	"fmt"
	"strconv"
)

func basicatoi() {
	
	testCases := []string{"1234", "-5678", "0", "-1"}
	
	for _, textNumber := range testCases {
		number, err := strconv.Atoi(textNumber)
		
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue
		}
		
		fmt.Println("Converted integer:", number)
	}
}