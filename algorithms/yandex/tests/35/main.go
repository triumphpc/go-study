package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Печеньки
// К Васе в гости пришли одноклассники. Его мама решила угостить ребят печеньем.
//Но не всё так просто. Печенья могут быть разного размера. А у каждого ребёнка есть фактор жадности —– минимальный размер печенья, которое он возьмёт.
//Нужно выяснить, сколько ребят останутся довольными в лучшем случае, когда они действуют оптимально.
//Каждый ребёнок может взять не больше одного печенья.
// В первой строке записано n —– количество детей.
//Во второй —– n чисел, разделённых пробелом, каждое из которых –— фактор жадности ребёнка. Это натуральные числа, не превосходящие 1000.
//В следующей строке записано число m –— количество печенек.
//Далее —– m натуральных чисел, разделённых пробелом —– размеры печенек. Размеры печенек не превосходят 1000.
//Оба числа n и m не превосходят 10000.

func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	peoples := make([]int, 0, size)

	for idx := 0; idx < size; idx++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		peoples = append(peoples, num)
	}

	scanner.Scan()
	sizeTwo, _ := strconv.Atoi(scanner.Text())
	cookies := make([]int, 0, sizeTwo)

	for idx := 0; idx < sizeTwo; idx++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		cookies = append(cookies, num)
	}

	quickSort(peoples, 0, size-1)    // Используем быструю сортировку для сортировки жадности
	quickSort(cookies, 0, sizeTwo-1) // Используем быструю сортировку для сортировки размера печенек

	result := quickSearch(peoples, cookies)

	fmt.Fprintf(writer, "%s ", fmt.Sprint(result))
}

func quickSearch(peoples, cookies []int) (result int) {
	for idx := range peoples {
		i := binarySearch(cookies, 0, len(cookies)-1, peoples[idx])
		//fmt.Println(peoples[idx], i, cookies, len(cookies)-1)
		if i >= 0 {
			result++

			cookies = append(cookies[:i], cookies[i+1:]...) // Обрезаем набор печенек
			if len(cookies) == 0 {
				break
			}
		}
	}

	return result
}

// binarySearch - определяет позицию оптимальную
func binarySearch(arr []int, left, right, need int) int {
	// Частный случай
	if right-left <= 1 {
		if arr[left] >= need {
			return left
		}

		if arr[right] >= need {
			return right
		}

		return -1
	}

	mid := (left + right) / 2

	if need <= arr[mid] { // Проверяем в какой части искомая печенька
		return binarySearch(arr, left, mid, need)
	} else {
		return binarySearch(arr, mid, right, need)
	}
}

func quickSort(arr []int, left, right int) {
	if right-left <= 1 {
		if arr[right] < arr[left] {
			arr[left], arr[right] = arr[right], arr[left]
		}

		return
	}

	mid := left + (right-left)/2
	leftPt := left
	rightPt := right

	for leftPt < rightPt {
		if arr[leftPt] < arr[mid] {
			leftPt++
			continue
		}

		if arr[mid] < arr[rightPt] {
			rightPt--
			continue
		}

		// swap
		arr[leftPt], arr[rightPt] = arr[rightPt], arr[leftPt]
		leftPt++
		rightPt--
	}

	quickSort(arr, left, mid)
	quickSort(arr, mid, right)
}
