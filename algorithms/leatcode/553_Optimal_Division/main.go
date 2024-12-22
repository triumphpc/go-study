package main

// Задача: 553. Optimal Division
//
//Дано целочисленный массив nums. Соседние целые числа в nums будут выполнять деление с плавающей запятой.
//
//Например, для nums = [2,3,4] мы будем вычислять выражение "2/3/4".
//Однако, вы можете добавить любое количество скобок в любое место, чтобы изменить приоритет операций. Вы хотите добавить эти скобки так, чтобы значение выражения после вычисления было максимальным.
//
//Верните соответствующее выражение, которое имеет максимальное значение в строковом формате.
// Алгоритм:
//
//1⃣Разверните оба числа. Инициализируйте массив ans с (N+M) нулями. Для каждой цифры во втором числе: держите переменную переноса, изначально равную 0. Инициализируйте массив (currentResult), начинающийся с некоторых нулей в зависимости от места цифры во втором числе.
//
//2⃣Для каждой цифры первого числа: умножьте цифру второго числа на цифру первого числа и добавьте предыдущий перенос к результату умножения. Возьмите остаток от деления на 10, чтобы получить последнюю цифру. Добавьте последнюю цифру к массиву currentResult. Разделите результат умножения на 10, чтобы получить новое значение переноса.
//
//3⃣После итерации по каждой цифре в первом числе, если перенос не равен нулю, добавьте перенос к массиву currentResult. Добавьте currentResult к массиву ans. Если последняя цифра в ans равна нулю, перед разворотом ans удалите этот ноль, чтобы избежать ведущего нуля в окончательном ответе. Разверните ans и верните его.
import (
	"fmt"
	"strconv"
)

func addStrings(num1 []int, num2 []int) []int {
	var ans []int
	carry := 0
	n1, n2 := len(num1), len(num2)

	for i := 0; i < max(n1, n2) || carry != 0; i++ {
		digit1, digit2 := 0, 0
		if i < n1 {
			digit1 = num1[i]
		}
		if i < n2 {
			digit2 = num2[i]
		}
		sum := digit1 + digit2 + carry
		carry = sum / 10
		ans = append(ans, sum%10)
	}
	return ans
}

func multiplyOneDigit(firstNumber string, secondNumberDigit rune, numZeros int) []int {
	var currentResult []int
	for i := 0; i < numZeros; i++ {
		currentResult = append(currentResult, 0)
	}
	carry := 0

	for _, digit := range firstNumber {
		multiplication := (int(secondNumberDigit-'0') * int(digit-'0')) + carry
		carry = multiplication / 10
		currentResult = append(currentResult, multiplication%10)
	}
	if carry != 0 {
		currentResult = append(currentResult, carry)
	}
	return currentResult
}

func multiply(firstNumber string, secondNumber string) string {
	if firstNumber == "0" || secondNumber == "0" {
		return "0"
	}

	firstNumberRev := reverseString(firstNumber)
	secondNumberRev := reverseString(secondNumber)

	ans := make([]int, len(firstNumber)+len(secondNumber))

	for i, digit := range secondNumberRev {
		ans = addStrings(multiplyOneDigit(firstNumberRev, digit, i), ans)
	}

	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}

	result := ""
	for i := len(ans) - 1; i >= 0; i-- {
		result += strconv.Itoa(ans[i])
	}

	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(multiply("123", "456"))
}
