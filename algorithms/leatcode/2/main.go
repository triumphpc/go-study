// Конвертер римских цифр
// https://leetcode.com/problems/integer-to-roman/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(3))

}

func intToRoman(num int) string {
	type Pair struct {
		Arab  int
		Roman string
	}

	var pairs = [13]Pair{
		0:  {Arab: 1000, Roman: "M"},
		1:  {Arab: 900, Roman: "CM"},
		2:  {Arab: 500, Roman: "D"},
		3:  {Arab: 400, Roman: "CD"},
		4:  {Arab: 100, Roman: "C"},
		5:  {Arab: 90, Roman: "XC"},
		6:  {Arab: 50, Roman: "L"},
		7:  {Arab: 40, Roman: "XL"},
		8:  {Arab: 10, Roman: "X"},
		9:  {Arab: 9, Roman: "IX"},
		10: {Arab: 5, Roman: "V"},
		11: {Arab: 4, Roman: "IV"},
		12: {Arab: 1, Roman: "I"},
	}
	var res string
	var pairIdx int

	for {
		curRes := num - pairs[pairIdx].Arab
		if curRes < 0 {
			pairIdx++

			if pairIdx == len(pairs) {
				break
			}
			continue
		}

		num = curRes
		res += pairs[pairIdx].Roman
	}

	return res
}

func intToRoman2(num int) string {
	return getRoman(num, "")
}

func getRoman(num int, roman string) string {
	if num == 0 {
		return roman
	}

	if num-1000 >= 0 {
		return getRoman(num-1000, roman+"M")
	}

	if num-900 >= 0 {
		return getRoman(num-900, roman+"CM")
	}

	if num-500 >= 0 {
		return getRoman(num-500, roman+"D")
	}

	if num-400 >= 0 {
		return getRoman(num-400, roman+"CD")
	}

	if num-100 >= 0 {
		return getRoman(num-100, roman+"C")
	}

	if num-90 >= 0 {
		return getRoman(num-90, roman+"XC")
	}

	if num-50 >= 0 {
		return getRoman(num-50, roman+"L")
	}

	if num-40 >= 0 {
		return getRoman(num-40, roman+"XL")
	}

	if num-10 >= 0 {
		return getRoman(num-10, roman+"X")
	}

	if num-9 >= 0 {
		return getRoman(num-9, roman+"IX")
	}

	if num-5 >= 0 {
		return getRoman(num-5, roman+"V")
	}

	if num-4 >= 0 {
		return getRoman(num-4, roman+"IV")
	}

	if num-1 >= 0 {
		return getRoman(num-1, roman+"I")
	}

	return roman
}
