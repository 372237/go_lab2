package main

import (
	"reflect"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int // ожидаемая длина слайса
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"размер 1", 1, 1},
		{"размер 100", 100, 100},
		{"размер 10000", 10000, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d) длина = %d, ожидалось %d",
					tt.size, len(got), tt.want)
			}
			// Проверяем, что все элементы положительные.
			for i, v := range got {
				if v < 0 {
					t.Errorf("generateRandomElements(%d) содержит отрицательный элемент по индексу %d: %d",
						tt.size, i, v)
					break
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"несколько элементов, максимум в конце", []int{1, 2, 3, 4, 5}, 5},
		{"несколько элементов, максимум в начале", []int{9, 1, 2, 3}, 9},
		{"несколько элементов, максимум в середине", []int{1, 2, 100, 3, 4}, 100},
		{"все элементы одинаковые", []int{7, 7, 7, 7}, 7},
		{"два элемента", []int{10, 3}, 10},
		{"большой слайс", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"меньше CHUNKS элементов", []int{1, 2, 3}, 3},
		{"ровно CHUNKS элементов", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"больше CHUNKS элементов", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"максимум в первом чанке", []int{100, 1, 2, 3, 4, 5, 6, 7, 8}, 100},
		{"максимум в последнем чанке", []int{1, 2, 3, 4, 5, 6, 7, 100}, 100},
		{"все одинаковые", []int{5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// Проверяем, что maxChunks и maximum дают одинаковый результат
// на одних и тех же данных.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(10000)
	want := maximum(data)
	got := maxChunks(data)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("maxChunks = %d, maximum = %d — результаты должны совпадать", got, want)
	}
}
