package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// -----
	// UTF-8
	// -----
	fmt.Printf("\n=> UTF-8\n")

	// Everything in Go is based on UTF-8 character sets.
	// If we use different encoding scheme, we might have a problem.

	// Declare a string with both Chinese and English characters.
	// For each Chinese character, we need 3 byte for each one.
	// The UTF-8 is built on 3 layers: bytes, code point and character. From Go perspective, string
	// are just bytes. That is what we are storing.
	// In our example, the first 3 bytes represents a single code point that represents that single
	// character. We can have anywhere from 1 to 4 bytes representing a code point (a code point is
	// a 32 bit value) and anywhere from 1 to multiple code points can actually represent a
	// character. To keep it simple, we only have 3 bytes representing 1 code point representing 1
	// character. So we can read s as 3 bytes, 3 bytes, 1 byte, 1 byte,... (since there are only 2
	// Chinese characters in the first place, the rest are English)
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune -> maximum number of bytes we need to
	// represent any code point is 4.
	// Rune is its own type. It is an alias for int32 type. Similar to type byte we are using, it
	// is just an alias for uint8.
	var buf [utf8.UTFMax]byte

	// When we are ranging over a string, are we doing it byte by byte or code point by code point or
	// character by character?
	// The answer is code point by code point.
	// On the first iteration, i is 0. On the next one, i is 3 because we are moving to the next
	// code point. Then i is 6.
	for i, r := range s {
		// Capture the number of bytes for this rune/code point.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		// Copy rune from the string to our buffer.
		// We want to go through every code point and copy them into our array buf, and display
		// them on the screen.
		// "Every array is just a slice waiting to happen." - Go saying
		// We are using the slicing syntax, creating our slice header where buf becomes the backing
		// array. All of them are on the stack. There is no allocation here.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
