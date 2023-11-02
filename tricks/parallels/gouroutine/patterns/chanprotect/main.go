package main

import "fmt"

// Если мы захотим сделать так, чтобы в канала handle никто больше не могу писать после его инициализации, кроме нашей горутины?

func main() {
	chanOwner := func() <-chan int {
		res := make(chan int, 5)

		go func() {
			defer close(res)
			for i := 0; i <= 5; i++ {
				res <- i
			}

		}()
		return res
	}

	consumer := func(res <-chan int) {
		for r := range res {
			fmt.Println(r)
		}

		fmt.Println("Done")
	}

	res := chanOwner()
	//res <- 5 // не скомпилируется
	consumer(res)

}
