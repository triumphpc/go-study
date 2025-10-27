package _8_find_index

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sadfd"

	k := strStr(haystack, needle)
	fmt.Println(k)

}

func strStr(haystack string, needle string) int {
	for readIdx := 0; readIdx < len(haystack); readIdx++ {
		if len(needle)+readIdx > len(haystack) {
			return -1
		}

		if haystack[readIdx:readIdx+len(needle)] == needle {
			return readIdx
		}
	}

	return -1
}
