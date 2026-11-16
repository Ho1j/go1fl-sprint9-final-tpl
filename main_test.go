package main

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 7, 8, 9, 10_000} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if size < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("len = %d, want %d", len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"single negative", []int{-42}, -42},
		{"zeros", []int{0, 0, 0}, 0},
		{"first maximum", []int{9, 1, 2}, 9},
		{"last maximum", []int{1, 2, 9}, 9},
		{"middle maximum", []int{1, 9, 2}, 9},
		{"duplicates", []int{5, 5, 5, 5}, 5},
		{"eight elements", []int{1, 2, 3, 4, 8, 6, 7, 5}, 8},
		{"remainder maximum", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"negative remainder", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"integer limits", []int{minInt, 0, maxInt}, maxInt},
		{"minimum integer", []int{minInt}, minInt},
	}
	functions := []struct {
		name string
		find func([]int) int
	}{
		{"maximum", maximum},
		{"maxChunks", maxChunks},
	}
	for _, function := range functions {
		for _, tt := range tests {
			t.Run(function.name+"/"+tt.name, func(t *testing.T) {
				before := slices.Clone(tt.data)
				if got := function.find(tt.data); got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
				if !slices.Equal(tt.data, before) {
					t.Error("input slice was modified")
				}
			})
		}
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	random := rand.New(rand.NewSource(1))
	for size := 0; size <= 1024; size++ {
		data := make([]int, size)
		for i := range data {
			data[i] = random.Intn(2001) - 1000
		}
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Fatalf("size %d: got %d, want %d", size, got, want)
		}
	}
}
