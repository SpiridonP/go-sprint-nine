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

	mainSlice := make([]int, size)

	for i := 0; i < size; i++ {
		mainSlice[i] = rand.Intn(1000)
	}

	return mainSlice

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return  0
	}

	maxNumber := data[0]

    for _, n := range data {
        if n > maxNumber {
            maxNumber = n
        }
    }
	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Вычисляем размер одной части.
	chunkSize := len(data) / CHUNKS

	// Слайс для хранения максимумов по частям.
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		// Вычисляем границы текущего среза.
		start := i * chunkSize
		end := start + chunkSize

		// Для последней части захватываем всё, что осталось (на случай, если len(data) не делится на CHUNKS).
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Запускаем горутину для поиска максимума в этом срезе.
		go func(s, e int, idx int) {
			defer wg.Done()

			// Находим максимум в срезе data[s:e].
			localMax := data[s]
			for j := s + 1; j < e; j++ {
				if data[j] > localMax {
					localMax = data[j]
				}
			}

			maxValues[idx] = localMax
		}(start, end, i)
	}

	wg.Wait()

	// Теперь находим максимум среди найденных максимумов частей.
	overallMax := maxValues[0]
	for i := 1; i < len(maxValues); i++ {
		if maxValues[i] > overallMax {
			overallMax = maxValues[i]
		}
	}

	return overallMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	chunkedMax := maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", chunkedMax, elapsed)
}
