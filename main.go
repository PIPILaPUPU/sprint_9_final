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

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// done
	if size <= 0 {
		return []int{}
	}

	rand.New(rand.NewSource(time.Now().Unix()))

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = int(rand.Int63())
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// done
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// done
	if len(data) == 0 {
		return 0
	}

	cutSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * cutSize
		end := start + cutSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(index, start, end int) {
			defer wg.Done()

			maxValues[index] = maximum(data[start:end])
		}(i, start, end)
	}

	wg.Wait()
	finalMax := maxValues[0]
	for i := 0; i < len(maxValues); i++ {
		if maxValues[i] > finalMax {
			finalMax = maxValues[i]
		}
	}

	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
