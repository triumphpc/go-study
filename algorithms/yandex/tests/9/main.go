package test

import (
	"fmt"
	"math"
)

// Вычисления значения функции
// Вася делает тест по математике: вычисляет значение функций в различных точках.
// Стоит отличная погода, и друзья зовут Васю гулять. Но мальчик решил сначала закончить тест и только после этого идти к друзьям.
//К сожалению, Вася пока не умеет программировать. Зато вы умеете. Помогите Васе написать код функции, вычисляющей y = ax2 + bx + c.
//Напишите программу, которая будет по коэффициентам a, b, c и числу x выводить значение функции в точке x.

type TestRun struct {
}

func (*TestRun) Name() string {
	return "Вычисления значения функции"
}

func (*TestRun) Run(k int, in ...float32) (result []float32, err error) {

	run(float64(in[0]), float64(in[1]), float64(in[2]), float64(in[3]))

	return result, nil
}

// Реализация алгоритма
func run(a, x, b, c float64) {

	result := a*math.Pow(x, 2) + b*x + c

	fmt.Println(result)
}