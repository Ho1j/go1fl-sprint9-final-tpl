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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		// Добавляем единицу, чтобы все числа были строго положительными.
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

// maximum returns the largest element, or zero for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks searches eight chunks concurrently and returns their maximum.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// Последняя часть включает остаток при делении длины на CHUNKS.
			end = len(data)
		}
		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			// Для пустой части используем элемент исходного слайса:
			// ноль мог бы исказить результат для отрицательных чисел.
			maxima[index] = data[0]
			if len(chunk) > 0 {
				maxima[index] = maximum(chunk)
			}
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
