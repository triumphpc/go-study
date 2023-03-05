package metrics

import (
	"testing"
	"time"
)

const maxSize = 10000000

func TestDuration(t *testing.T) {
	defer Duration(Track(t.Name()))

	go foo()

	time.Sleep(10 * time.Second)
}

func foo() {
	for {
		s := make([]int, 0, maxSize)
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}

}
