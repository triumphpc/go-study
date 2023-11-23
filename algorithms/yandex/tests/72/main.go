// Конвертер римских цифр
// https://leetcode.com/problems/integer-to-roman/

package main

import "fmt"

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
