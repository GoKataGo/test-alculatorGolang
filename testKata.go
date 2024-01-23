package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// римские числа
var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
}

func isRomanNumeral(str string) bool {
	for _, char := range str {
		if _, ok := romanNumerals[char]; !ok {
			return false
		}
	}
	return true
}

// в арабские числа
func toArabicNumeral(roman string) int {
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result
}

// из арабских цифр в римские
func arabicToRoman(num int) string {
	if num <= 0 || num > 3999 {
		return "Недопустимое число"
	}

	var result strings.Builder

	for num > 0 {
		switch {
		case num >= 1000:
			result.WriteString("M")
			num -= 1000
		case num >= 900:
			result.WriteString("CM")
			num -= 900
		case num >= 500:
			result.WriteString("D")
			num -= 500
		case num >= 400:
			result.WriteString("CD")
			num -= 400
		case num >= 100:
			result.WriteString("C")
			num -= 100
		case num >= 90:
			result.WriteString("XC")
			num -= 90
		case num >= 50:
			result.WriteString("L")
			num -= 50
		case num >= 40:
			result.WriteString("XL")
			num -= 40
		case num >= 10:
			result.WriteString("X")
			num -= 10
		case num >= 9:
			result.WriteString("IX")
			num -= 9
		case num >= 5:
			result.WriteString("V")
			num -= 5
		case num >= 4:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num--
		}
	}
	return result.String()
}

func calculateExpressionArab(expression string) (int, error) {

	tokens := strings.Fields(expression)
	/*
		if len(tokens) != 3 {
			return 0, fmt.Errorf("Неверный формат выражения")
		}
	*/

	num1Str := tokens[0]
	operator := tokens[1]
	num2Str := tokens[2]

	var num1, num2 int
	var err error

	num1, err = strconv.Atoi(num1Str)
	if err != nil {
		return 0, fmt.Errorf("Ошибка парсинга первого числа")
	}

	num2, err = strconv.Atoi(num2Str)
	if err != nil {
		return 0, fmt.Errorf("Ошибка парсинга второго числа")
	}

	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		return 0, fmt.Errorf("Числа должны быть от 1 до 10")
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
		return 0, fmt.Errorf("Неверный арифметический оператор")
	}

	return result, nil
}

func calculateExpressionRome(expression string) (string, error) {

	tokens := strings.Fields(expression)

	num1Str := tokens[0]
	operator := tokens[1]
	num2Str := tokens[2]

	var num1, num2 int

	if isRomanNumeral(num1Str) {
		num1 = toArabicNumeral(num1Str)
	}

	if isRomanNumeral(num2Str) {
		num2 = toArabicNumeral(num2Str)
	}

	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		return "", fmt.Errorf("Числа должны быть от 1 до 10")
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
		return "", fmt.Errorf("Неверный арифметический оператор")
	}

	if result < 0 {
		panic("В римской системе нет отрицательных чисел.")
	}
	resultStr := arabicToRoman(result)
	return resultStr, nil
}

func schitatStr(text string) {

	expression3 := text

	tokens := strings.Fields(expression3)

	if len(tokens) < 3 {
		panic("Строка не является математической операцией.")
	} else if len(tokens) > 4 {
		fmt.Println(len(text))
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else {
		num1Str := tokens[0]
		num2Str := tokens[2]

		if isRomanNumeral(num1Str) && isRomanNumeral(num2Str) && isRomanNumeral(num1Str) == isRomanNumeral(num2Str) {
			result3, err3 := calculateExpressionRome(expression3)
			if err3 != nil {
				panic("Используются одновременно разные системы счисления.")
			} else {
				fmt.Printf("Результат: %s\n", result3)
			}
		} else {
			result2, err2 := calculateExpressionArab(expression3)
			if err2 != nil {
				panic("Используются одновременно разные системы счисления.")
			} else {
				fmt.Printf("Результат: %v\n", result2)
			}
		}
	}
}

func main() {

	fmt.Println("Введите выражение ввиде 1 + 1 или I + I")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	schitatStr(text)
}
