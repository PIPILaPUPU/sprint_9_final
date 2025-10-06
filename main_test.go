package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
// Тесты для generateRandomElements
func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	assert.Equal(t, size, len(result), "the slice length must match the requested size")

	for _, num := range result {
		assert.GreaterOrEqual(t, num, 0, "number must be positive")
		assert.Less(t, num, 100_000_000, "numbers must be lower than 100_000_000")
	}
}

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	result := generateRandomElements(0)
	assert.Empty(t, result, "for size 0, an empty slice should be returned.")
}

func TestGenerateRandomElements_NegativeSize(t *testing.T) {
	result := generateRandomElements(-5)
	assert.Empty(t, result, "for a negative size, an empty slice should be returned.")
}

func TestGenerateRandomElements_SingleElement(t *testing.T) {
	result := generateRandomElements(1)
	assert.Len(t, result, 1, "for size 1, a slice with one element should be returned.")
}

func TestGenerateRandomElements_Content(t *testing.T) {
	data := generateRandomElements(100)

	assert.NotEmpty(t, data, "the generated data must not be empty.")
	assert.Len(t, data, 100, "100 elements should be generated")

	for i, num := range data {
		assert.GreaterOrEqual(t, num, 0, "number %d by index %d must be positive", num, i)
		assert.Less(t, num, 100_000_000, "number %d пby index %d must be lower than 100_000_000", num, i)
	}
}

// Тесты для maximum (однопоточный поиск)
func TestMaximum_NormalCase(t *testing.T) {
	data := []int{1, 5, 3, 9, 2, 8, 4, 6, 7}
	result := maximum(data)
	assert.Equal(t, 9, result, "the maximum must be found correctly.")
}

func TestMaximum_EmptySlice(t *testing.T) {
	result := maximum([]int{})
	assert.Equal(t, 0, result, "for an empty slice, 0 should be returned.")
}

func TestMaximum_SingleElement(t *testing.T) {
	result := maximum([]int{42})
	assert.Equal(t, 42, result, "for one element, this element must be returned")
}

func TestMaximum_AllSameElements(t *testing.T) {
	data := []int{7, 7, 7, 7}
	result := maximum(data)
	assert.Equal(t, 7, result, "for identical elements, this element should be returned")
}

func TestMaximum_LargeNumbers(t *testing.T) {
	data := []int{1000, 500, 1500, 2000}
	result := maximum(data)
	assert.Equal(t, 2000, result, "should handle large numbers correctly")
}

// Тесты для maxChunks (многопоточный поиск)
func TestMaxChunks_NormalCase(t *testing.T) {
	data := []int{1, 5, 3, 9, 2, 8, 4, 6, 7}
	result := maxChunks(data)
	assert.Equal(t, 9, result, "multithreaded search must find the correct maximum")
}

func TestMaxChunks_EmptySlice(t *testing.T) {
	result := maxChunks([]int{})
	assert.Equal(t, 0, result, "for an empty slice, 0 should be returned.")
}

func TestMaxChunks_SingleElement(t *testing.T) {
	result := maxChunks([]int{42})
	assert.Equal(t, 42, result, "for one element, this element must be returned")
}

func TestMaxChunks_LessThanChunks(t *testing.T) {
	data := []int{1, 5, 3} // меньше чем CHUNKS (8)
	result := maxChunks(data)
	assert.Equal(t, 5, result, "should handle slices smaller than CHUNKS correctly")
}

func TestMaxChunks_ExactlyChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8} // ровно CHUNKS элементов
	result := maxChunks(data)
	assert.Equal(t, 8, result, "should handle slices of size CHUNKS correctly")
}

func TestMaxChunks_MoreThanChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	result := maxChunks(data)
	assert.Equal(t, 16, result, "should correctly handle slices larger than CHUNKS")
}

func TestMaxChunks_AllSameElements(t *testing.T) {
	data := []int{7, 7, 7, 7, 7, 7, 7, 7, 7}
	result := maxChunks(data)
	assert.Equal(t, 7, result, "Should handle identical elements correctly")
}
