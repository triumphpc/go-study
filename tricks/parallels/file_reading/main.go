package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/disintegration/imaging"
)

// Напишем функцию, которая для из каталога pics получит список файлов и
// создаст плавью для них в этом же каталоге.
// 1. Реализовать параллельное чтение

func main() {
	files := []string{
		"files/1.jpg",
		"files/2.jpg",
		"files/3.jpg",
	}
	//makeThumbnails(files)

	run(files)

}

func makeThumbnails(filenames []string) { // Наивное решение
	for _, f := range filenames {
		src, err := imaging.Open(f)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		// Resize srcImage to size = 128x128px using the Lanczos filter.
		dst := imaging.Resize(src, 128, 128, imaging.Lanczos)
		err = imaging.Save(dst, f+"_small.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}

	}
}

func run(fileNames []string) {
	filenames := make(chan string, len(fileNames))
	var wg sync.WaitGroup

	// 1. Создаем читателей. Читают данные и пишут в канал
	for _, f := range fileNames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			filenames <- f
		}(f)
	}

	go func() {
		wg.Wait()
		close(filenames)
	}()

	// 2. Создатель/писатель
	size := makeThumbnails6(filenames)
	fmt.Println(size)

}

func makeThumbnails6(filenames <-chan string) int64 { // Решение с паралельной обраоткой
	sizes := make(chan int64) // Создаем канал
	var wg sync.WaitGroup     // Count of goroutines
	errCh := make(chan error)

	for f := range filenames {
		wg.Add(1)

		go func(f string) { // Выполняет паралельно обработку
			defer wg.Done()
			thumb, err := makeThumb3(f)
			if err != nil {
				errCh <- err

				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
		close(errCh)
	}()

	go func() {
		for err := range errCh {
			log.Println(err)
		}
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func makeThumb3(f string) (thumbName string, err error) {
	src, err := imaging.Open(f)
	if err != nil {
		return "", err
	}
	// Resize srcImage to size = 128x128px using the Lanczos filter.
	dst := imaging.Resize(src, 128, 128, imaging.Lanczos)
	name := f + "_small.jpg"
	err = imaging.Save(dst, name)
	if err != nil {
		return "", err
	}
	return name, nil
}
