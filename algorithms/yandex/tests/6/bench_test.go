package test

import (
	"testing"
	"time"

	y "github.com/triumphpc/go-study/algorithms/yandex"
	"golang.org/x/exp/slices"
)

// go test -run none -bench . -benchmem -benchtime 3s
// BenchmarkBasic-16        	 4080832	      1675 ns/op	       8 B/op	       1 allocs/op
// BenchmarkWithPoint-16    	  920319	      4137 ns/op	       8 B/op	       1 allocs/op
// BenchmarkWithMap-16      	 1313520	      2422 ns/op	   16416 B/op	       3 allocs/op

import (
	"math/rand"
)

var N []float32
var K = 3
var runner = y.Runner{}

func setup() {
	rand.Seed(time.Now().Unix())

	N = make([]float32, 0, 1000)
	for i := 0; i < 1000; i++ {
		n := rand.Intn(9-1) + 1
		N = append(N, float32(n))
	}
}

func al1(k int, in ...float32) (result []float32, err error) {

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if (in[i] + in[j]) == float32(k) {
				result = append(result, in[i], in[j])

				return result, nil
			}

		}
	}

	return result, nil
}

func al2(k int, in ...float32) (result []float32, err error) {
	// 1. Сортируем список
	slices.Sort(in)

	// 2. Установка и обход по указателям
	left := 0
	right := len(in) - 1

	for left < right {
		currentSum := in[left] + in[right]
		if in[left]+in[right] == float32(k) {
			result = append(result, in[left], in[right])

			return result, nil
		}

		if currentSum > float32(k) {
			right--
		} else {
			left++
		}
	}

	return result, nil
}

func al3(k int, in ...float32) (result []float32, err error) {
	set := make(map[float32]bool, len(in))

	for idx := range in {
		y := float32(k) - in[idx]

		// Если уже есть в карте, то возвращаем
		if set[y] {
			result = append(result, in[idx], y)

			return
		}

		set[in[idx]] = true
	}

	return result, nil
}

func BenchmarkBasic(b *testing.B) {
	setup()

	for i := 0; i < b.N; i++ {

		al1(K, N...)
	}
}

func BenchmarkWithPoint(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		al2(K, N...)
	}
}

func BenchmarkWithMap(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		al3(K, N...)
	}

}
