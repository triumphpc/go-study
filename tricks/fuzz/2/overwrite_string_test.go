package overwrite_string

import "testing"

func TestOverwriteString(t *testing.T) {
	//OverwriteString("Hello, World!", 'A', 5)
	OverwriteString("0000000̠̠", -33, 11)
}

func FuzzBasicOverwriteString(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string, value rune, n int) {
		OverwriteString(str, value, n)
	})
}
