// https://leetcode.com/problems/hamming-distance/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingDistance(4, 14))

}

func hammingDistance(x int, y int) int {
	xb := strconv.FormatInt(int64(x), 2)
	yb := strconv.FormatInt(int64(y), 2)

	lenXb := len(xb)
	lenYb := len(yb)

	if lenXb > lenYb {
		for i := 0; i < lenXb-lenYb; i++ {
			yb = "0" + yb
		}
	} else if lenXb < lenYb {
		for i := 0; i < lenYb-lenXb; i++ {
			xb = "0" + xb
		}
	}

	dst := 0
	for k, v := range xb {
		if uint8(v) != yb[k] {
			dst++
		}
	}

	return dst
}
