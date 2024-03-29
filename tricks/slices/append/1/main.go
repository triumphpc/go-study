// Изменения функции append в Go 1.18

package main

import "fmt"

// Как это было:
//1. Если требуемая емкость cap больше двойного размера старой емкости old.cap,
// то требуемая емкость cap будет использована в качестве новой newcap .
//
//2. В противном случае, если старая емкость old.cap меньше 1024 (Байт = Кб).
//Конечной емкостью newcap будет увеличение в 2 раза старой емкости (old.cap), то есть newcap = doublecap
//
// 3. Если оба предыдущих условия не выполнены, а длина старого среза больше или равна 1024,
// окончательная емкость newcap начинается со старой емкости old.cap и циклически увеличивается
//на 1/4 от исходной, где newcap = old.cap, для {newcap + = newcap / 4} до тех пор,
// пока конечной емкостью newcap не станет емкость большая требуемой емкости cap, то есть newcap >= cap

func main() {
	// < 1.8
	for i := 0; i < 2000; i += 100 {
		fmt.Println(i, cap(append(make([]bool, i), true)))
		// 0 8
		//100 208
		//200 416
		//300 640
		//400 896
		//500 1024
		//600 1280
		//700 1408
		//800 1792
		//900 2048
		//1000 2048
		//1100 1408 <- нарушение монотонности (предыдущее значение больше)
		//1200 1536
	}
}
