package main

import "fmt"

// Разные деревья.
// Ребятам стало интересно, сколько может быть различных деревьев поиска, содержащих в своих узлах все уникальные числа
// от 1 до n. Помогите им найти ответ на этот вопрос.
//	1
//		2
//			3
//
// 	1
//		3
//	2
//
//	2
//		3
//	1
//
//  	3
//	2
//1
//
//		3
//	1
//		2

func main() {
	fmt.Println(task(4)) // 14

}

func task(n int) (cnt int) {
	for i := 1; i <= n; i++ {

		for j := i + 1; j <= n; j++ {
			cnt++

			if i > 1 {
				for k := 1; k < j; k++ {
					cnt++
				}
			}
		}
	}

	return cnt
}
