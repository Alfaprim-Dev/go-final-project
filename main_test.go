package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"zero", 0, 0},
		{"negative", -5, 0},
		{"one", 1, 1},
		{"typical", 1000, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("len = %d, want %d", len(got), tt.want)
			}
			for i, v := range got {
				if v <= 0 {
					t.Errorf("elem[%d] = %d, want positive", i, v)
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
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{42}, 42},
		{"max at start", []int{100, 1, 2}, 100},
		{"max at end", []int{1, 2, 100}, 100},
		{"max in middle", []int{1, 100, 2}, 100},
		{"all equal", []int{7, 7, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum = %d, want %d", got, tt.want)
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
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{42}, 42},
		{"less than chunks", []int{3, 1, 4, 1, 5}, 5},
		{"exactly chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"more than chunks", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(10_000)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Errorf("maxChunks = %d, maximum = %d", got, want)
	}
}
