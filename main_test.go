package main

import (
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{name: "negative size", size: -1},
		{name: "zero size", size: 0},
		{name: "single element", size: 1, wantSize: 1},
		{name: "many elements", size: 1000, wantSize: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.wantSize {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", tt.size, len(got), tt.wantSize)
			}

			// Проверяем диапазон значений: в твоей функции rand.Intn(1000) даёт [0, 999]
			for i, value := range got {
				if value < 0 || value >= 1000 {
					t.Errorf("element %d = %d, want in range [0, 1000)", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name      string
		data      []int
		want      int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "all same values", data: []int{5, 5, 5}, want: 5},
		{name: "max at start", data: []int{10, 3, 8, 2}, want: 10},
		{name: "max in middle", data: []int{3, 7, 9, 5}, want: 9},
		{name: "max at end", data: []int{2, 1, 4, 9}, want: 9},
		{name: "negative numbers", data: []int{-5, -2, -10, -1}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

