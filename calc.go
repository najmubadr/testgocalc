package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	numbers        = "0123456789"
	operators      = "+-*/"
	romanSymbols   = "IVXLC"
	allowedSymbols = operators + numbers + romanSymbols
	romanMode      = false
	hasError       = false
	action         string
)

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func toNumber(n string) string {
	out := 0
	ln := len(n)
	for i := 0; i < ln; i++ {
		c := string(n[i])
		vc := num[c]
		if i < ln-1 {
			cnext := string(n[i+1])
			vcnext := num[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return strconv.Itoa(out)
}

func toRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, string(10))
	input = strings.ReplaceAll(input, " ", "")
	return input
}

func resetState() {
	romanMode = false
	hasError = false
	action = ""
}
func isValid(input string, sep string) bool {

	ops := strings.Split(input, sep)

	if isRoman(ops[0]) && isRoman(ops[1]) {
		romanMode = true
		return true
	}
	if isNum(ops[0]) && isNum(ops[1]) {
		romanMode = false
		return true
	}
	return false
}

func isRoman(s string) bool {
	r := strings.Index(romanSymbols, string(s[0]))

	if r > -1 {
		return true
	} else {
		return false
	}
}

func isNum(s string) bool {
	r := strings.Index(numbers, string(s[0]))
	if r > -1 {
		return true
	} else {
		return false
	}
}

func calc(operation string, a, b int) int {
	switch operation {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mul(a, b)
	default:
		return div(a, b)
	}
}

func stringToNumber(s string) int {
	str, _ := strconv.Atoi(s)
	return str
}

func throwError(msg string) {
	err := errors.New(msg)
	fmt.Println(err)
}

func main() {

	for true {
		fmt.Print("\nВведите данные: ")

		input := readInput()

		for _, s := range []rune(input) {

			if ok := strings.Index(allowedSymbols, string(s)); ok < 0 {
				hasError = true
				throwError("Ошибка! Введены недопустимые данные.")
				break
			}
		}

		if hasError == false {

			for _, operator := range operators {
				s := strings.Index(input, string(operator))
				if s > -1 {
					action = string(input[s])
					break
				}
			}

			if isValid(input, action) {
				var operands [2]int
				ops := strings.Split(input, action)

				if romanMode == true {
					for i, operand := range ops {
						ops[i] = toNumber(operand)
					}
				}

				for i, operand := range ops {
					operands[i] = stringToNumber(operand)
				}

				result := calc(action, operands[0], operands[1])

				if romanMode == true {
					if result > 0 {
						fmt.Println(toRoman(result))
					} else {
						throwError("Вне диапазона римского счета")
					}
				} else {
					fmt.Println(result)
				}

			} else {
				throwError("Ошибка! Оба операнда должны быть в одной и той же системе.")
			}
		}

		resetState()

	}

}
