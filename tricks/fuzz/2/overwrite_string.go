package overwrite_string

import (
	"strings"
	"unicode/utf8"
)

//go test -fuzz FuzzBasicOverwriteString
//   Failing input written to testdata/fuzz/FuzzBasicOverwriteString/25fd1c157a004ec82704ba6d1b43ca3a09f7d31f760bc5c2e702a0714b017620
//    To re-run:
//    go test -run=FuzzBasicOverwriteString/25fd1c157a004ec82704ba6d1b43ca3a09f7d31f760bc5c2e702a0714b017620
//  cat testdata/fuzz/FuzzBasicOverwriteString/25fd1c157a004ec82704ba6d1b43ca3a09f7d31f760bc5c2e702a0714b017620                 1 ↵ triumphpc@MacBook-Pro-triumphpc
// go test fuzz v1
// string("")
// int32(-19)
// int(0)
//
// https://blog.fuzzbuzz.io/go-fuzzing-basics/

// OverwriteString overwrites the first 'n' characters in a string with
// the rune 'value'
func OverwriteString(str string, value rune, n int) string {
	// If asked to overwrite more than the entire string then no need to loop,
	// just return string length * the rune
	//if n > len(str) {
	//	return strings.Repeat(string(value), len(str))
	//}

	if n >= utf8.RuneCountInString(str) { // это корректно
		return strings.Repeat(string(value), len(str))
	}

	result := []rune(str)

	for i := 0; i <= n; i++ {
		result[i] = value
	}
	return string(result)
}
