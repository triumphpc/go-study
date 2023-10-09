package main

import (
	"io"
	"strings"
	"testing"
)

func BenchmarkBasic(b *testing.B) {
	in := strings.NewReader("8\nвышивание крестиком\nрисование мелками на парте\nнастольный керлинг\nнастольный керлинг\nкухня африканского племени ужасмай\nтяжелая атлетика\nтаракановедение\nтаракановедение\n")
	for i := 0; i < b.N; i++ {
		task(in, io.Discard)
	}
}

func BenchmarkBasic2(b *testing.B) {
	in := strings.NewReader("8\nвышивание крестиком\nрисование мелками на парте\nнастольный керлинг\nнастольный керлинг\nкухня африканского племени ужасмай\nтяжелая атлетика\nтаракановедение\nтаракановедение\n")
	for i := 0; i < b.N; i++ {
		task2(in, io.Discard)
	}
}
