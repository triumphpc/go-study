package main

func main() {
	m := make(map[float64]int, 8)
	k := 0.0
	m[k] = 0
	k /= k // k is NaN now
	m[k] = 1
	m[k] = 2
	for k := range m {
		delete(m, k)
	}
	print(len(m)) // 2
	//clear(m)
	println(len(m))
}
