// Есть робот, который может ходить только вправо и вверх. Есть поле n = 5, m = 5. Нужно найти количество возможных от A путей до точки B
// https://www.youtube.com/watch?v=GhiRlhPlJ9Q
//
// Нужно понимать, что в точку Б робот может войти только с точек n-1 и m-1, для каждой такой точки мы можем вывести формулу:
// количество путей paths(n, m) = paths(n-1, m) + paths(n, m -1)
// Нужно выделить бизу для функции. Мы понимаем, что количество добраться до клеток, находящихся слева от робота и снизу робота = 0 ( он не может ходить в эти направлекния)
// paths(x, 0) = paths(0, x) = 0
// Если выход будет находиться в стартовой клетке с роботом - то будет только один маршрут до выхода: paths(1,1) = 1
//

package main

import "fmt"

func main() {
	fmt.Println(paths(4, 5))
	fmt.Println(paths2(4, 5))

}

// Возвращает количество путей
func paths(n, m int) int {
	if n < 1 || m < 1 { // если выход находится слева или снизу от робота
		return 0
	}

	if n == 1 && m == 1 { // Если в точке старта робота
		return 1
	}

	// В противном случае количество путей то левой точки выхода + до нижней от выхода
	// И рекурсивно пробегаем по всем путям
	return paths(n-1, m) + paths(n, m-1)
}

// Реализация динамическое программирование
func paths2(n, m int) int {
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
	}

	return helper(n, m, arr)
}

// Релизация с динамическим программированием с кешированием подсчета динамически
func helper(n, m int, arr [][]int) int {
	if n < 1 || m < 1 { // если выход находится слева или снизу от робота
		return 0
	}

	if n == 1 && m == 1 { // Если в точке старта робота
		return 1
	}

	if arr[n-1][m-1] != 0 {
		return arr[n-1][m-1]
	}

	// В противном случае количество путей то левой точки выхода + до нижней от выхода
	// И рекурсивно пробегаем по всем путям
	arr[n-1][m-1] = helper(n-1, m, arr) + helper(n, m-1, arr)

	return arr[n-1][m-1]
}
