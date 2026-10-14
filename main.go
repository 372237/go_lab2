package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс случайных положительных чисел заданного размера.
// Если size <= 0, возвращает nil.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum возвращает максимальный элемент слайса.
// Если слайс пустой, возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks находит максимум в слайсе, разбивая его на CHUNKS частей
// и обрабатывая каждую часть в отдельной горутине.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Если элементов меньше, чем частей, обрабатываем в один поток.
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// Последний срез забирает все оставшиеся элементы.
			end = len(data)
		}

		go func(idx int, chunk []int) {
			defer wg.Done()
			maxValues[idx] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n", max, elapsed)
}
