package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"os"
	"strconv"
)

func NumIsValid(num string) (int, string) {
	arabic := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	numeral := ""
	for i := 0; i < len(arabic); i++ {
		if num == arabic[i] {
			numeral = "arabic"
			return ParserInt(num), numeral
		} else if num == roman[i] {
			numeral = "roman"
			return RomanToInt(num), numeral
		}
	}
	return 0, ""
}

func Calculator(x, y int, oper string) (int, error) {
	switch oper {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "/":
		return x / y, nil
	case "*":
		return x * y, nil
	default:
		return 0, errors.New("Неверный оператор")
	}
}

func ParserInt(num string) int {
	var result int
	result, _ = strconv.Atoi(num)
	return result
}

func RomanToInt(num string) int {
	romanToArabic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}

	result := 0

	for i := len(num) - 1; i >= 0; i-- {
		current := romanToArabic[string(num[i])]
		if i == len(num)-1 || current >= romanToArabic[string(num[i+1])] {
			result += current
		} else {
			result -= current
		}
	}
	return result
}

func IntToRoman(number int) string {
	conv := []struct {
		arabic int
		roman  string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	result := ""
	for _, conv := range conv {
		for number >= conv.arabic {
			result += conv.roman
			number -= conv.arabic
		}
	}
	return result
}

func OutputResult(x, oper, y string) {
	num1, numerical1 := NumIsValid(x)
	num2, numerical2 := NumIsValid(y)

	if num1 == 0 || num2 == 0 {
		fmt.Println("Недопустимые значения. Используйте арабские или римские цифры от 1 до 10")
		return
	}
	if (numerical1 != numerical2) && (numerical1 != "") && (numerical2 != "") {
		fmt.Println("Используются одновременно разные системы счисления")
		return
	}

	result, err := Calculator(num1, num2, oper)
	if err != nil {
		fmt.Println(err)
		return
	}
	if numerical1 == "roman" {
		if result < 1 {
			fmt.Println("В римской системе нет отрицательных чисел или нуля")
			return
		}
		fmt.Println(IntToRoman(result))
		return
	}
	fmt.Println(result)
}

func main() {
	var x, oper, y string
	var inputtext string
	fmt.Println("Введите данные:")
	inputex := bufio.NewScanner(os.Stdin)
	inputex.Scan()
	inputtext = inputex.Text()
	arr := strings.Split(inputtext, " ")
	if len(arr) != 3 {
		fmt.Println("Математическая операция должна состоять из двух элементов")
		return
	}
	x, oper, y = arr[0], arr[1], arr[2]
	OutputResult(x, oper, y)
}
