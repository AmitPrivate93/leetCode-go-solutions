package main

import (
	"fmt"
	"math"
)

func myAtoi1(str string) int {
	s := []rune(str)
	i, result, sign, l := 0, 0, 1, len(s)
	MAX, MIN := (1<<31)-1, 1<<31
	// skip spaces
	for i < l && s[i] == ' ' {
		i++
	}
	// check if first non-space value is a sign
	if i < l && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	// iterate over digits
	for i < l && isNum(s[i]) {
		d := int(s[i] - '0')
		if sign > 0 && result > MAX/10 || sign < 0 && result > MIN/10 {
			if sign == 1 {
				return MAX
			} else {
				return -MIN
			}
		}
		result = result*10 + d
		i++
	}
	return result * sign
}

func myAtoi2(s string) int {
	started := false
	negative := false

	result := 0
exit:
	for _, char := range s {
		switch {
		case !started && char == ' ':
			continue

		case !started && (char == '+' || char == '-'):
			negative = char == '-'
			started = true

		case !isNum(char):
			break exit

		default:
			started = true
			digit := int(char - '0')
			if negative && result > ((-1*math.MinInt32)-digit)/10 {
				return math.MinInt32
			} else if !negative && result > ((math.MaxInt32-digit)/10) {
				return math.MaxInt32
			}
			result = result*10 + digit
		}
	}

	if negative {
		return -1 * result
	}

	return result
}

func isNum(c rune) bool { return c >= '0' && c <= '9' }

func main() {
	fmt.Println(myAtoi2("-90"))
	fmt.Println(myAtoi1("-9009 99"))
}
