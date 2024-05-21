package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currValue, ok := romanMap[roman[i]]
		if ok {
			if currValue >= prevValue {
				arabic += currValue
			} else {
				arabic -= currValue
			}
			prevValue = currValue
		} else {
			panic("Invalid Roman number!")
		}
	}

	return arabic
}

func arabicToRoman(arabic int) string {
	if arabic > 100 || arabic < 0 {
		panic("Arabic number out of range!")
	}
	romanValues := [...]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicValues := [...]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := ""
	i := 0

	for arabic != 0 {
		if arabic >= arabicValues[i] {
			arabic -= arabicValues[i]
			roman += romanValues[i]
		} else {
			i++
		}
	}
	return roman
}

func main() {
	fmt.Print("Enter an expression: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		panic("Invalid expression length!")
	}
	isRoman := false
	num1, err1 := strconv.Atoi(parts[0])
	operator := parts[1]
	num2, err2 := strconv.Atoi(parts[2])

	if err1 != nil {
		if err2 != nil {
			num1 = romanToArabic(parts[0])
			num2 = romanToArabic(parts[2])
			isRoman = true
		} else {
			panic("Invalid formatting!")
		}
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Invalid numbers!")
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic("Invalid operator!")
	}

	if isRoman {
		fmt.Print(arabicToRoman(result))
	} else {
		fmt.Print(result)
	}
}
