package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestLOS(elements []int) *LOS {
	if len(elements) == 0 {
		return nil
	}
	los := newElement(elements[0])
	ptr := los
	for _, element := range elements[1:] {
		ptr.next = newElement(element)
		ptr = ptr.next
	}
	return los
}

type testpairs struct {
	maxNumber int
	elements  []int
	name      string
}

func TestNewLOS(t *testing.T) {
	var tests = []testpairs{
		{0, []int{}, "negative: no elements"},
		{1, []int{1}, "only one element"},
		{2, []int{1, 2}, "two elements"},
		{5, []int{1, 2, 3, 4, 5}, "usually list"},
	}
	for _, pair := range tests {
		t.Run(pair.name, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
		})
	}
}

var testsSort = []testpairs{
	{1, []int{1}, "only one element: no sort"},
	{2, []int{1, 2}, "two elements: no sort"},
	{3, []int{1, 3, 2}, "three elements: one swap"},
	{4, []int{1, 4, 2, 3}, "four elements: one swap"},
	{5, []int{1, 5, 2, 4, 3}, "usually sort"},
	{8, []int{1, 8, 2, 7, 3, 6, 4, 5}, "from readme"},
	{9, []int{1, 9, 2, 8, 3, 7, 4, 6, 5}, "long list"},
	{102, []int{1, 102, 2, 101, 3, 100, 4, 99, 5, 98, 6, 97, 7, 96,
		8, 95, 9, 94, 10, 93, 11, 92, 12, 91, 13, 90, 14, 89, 15, 88, 16, 87, 17, 86,
		18, 85, 19, 84, 20, 83, 21, 82, 22, 81, 23, 80, 24, 79, 25, 78, 26, 77, 27, 76,
		28, 75, 29, 74, 30, 73, 31, 72, 32, 71, 33, 70, 34, 69, 35, 68, 36, 67, 37, 66,
		38, 65, 39, 64, 40, 63, 41, 62, 42, 61, 43, 60, 44, 59, 45, 58, 46, 57, 47, 56,
		48, 55, 49, 54, 50, 53, 51, 52}, "very long list"},
}

func TestLOS_Sort1(t *testing.T) {
	for _, pair := range testsSort {
		t.Run(pair.name, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			los.Sort1()
			assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
		})
	}
}

func TestLOS_Sort2(t *testing.T) {
	for _, pair := range testsSort {
		t.Run(pair.name, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			los.Sort2()
			assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
		})
	}
}

func TestLOS_Sort3(t *testing.T) {
	for _, pair := range testsSort {
		t.Run(pair.name, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			los.Sort3()
			assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
		})
	}
}

func TestLOS_Sort4(t *testing.T) {
	for _, pair := range testsSort {
		t.Run(pair.name, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			los.Sort4()
			assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
		})
	}
}

type testpairs2 struct {
	maxNumber int
	stringLOS string
}

func TestLOS_String(t *testing.T) {
	var tests = []testpairs2{
		{1, "1"},
		{2, "1, 2"},
		{5, "1, 2, 3, 4, 5"},
	}
	for _, pair := range tests {
		t.Run("print: " + pair.stringLOS, func(t *testing.T) {
			los := NewLOS(pair.maxNumber)
			assert.Equal(t, pair.stringLOS, (*los).String())
		})
	}
}

func benchmarkLOS_Sort1_N(b *testing.B, n int) {
	b.ReportAllocs()
	los := NewLOS(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort1()
	}
}

func benchmarkLOS_Sort2_N(b *testing.B, n int) {
	b.ReportAllocs()
	los := NewLOS(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort2()
	}
}

func benchmarkLOS_Sort3_N(b *testing.B, n int) {
	b.ReportAllocs()
	los := NewLOS(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort3()
	}
}

func benchmarkLOS_Sort4_N(b *testing.B, n int) {
	b.ReportAllocs()
	los := NewLOS(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort4()
	}
}

func BenchmarkLOS_Sort1_10(b *testing.B) {
	benchmarkLOS_Sort1_N(b, 10)
}

func BenchmarkLOS_Sort2_10(b *testing.B) {
	benchmarkLOS_Sort2_N(b, 10)
}

func BenchmarkLOS_Sort3_10(b *testing.B) {
	benchmarkLOS_Sort3_N(b, 10)
}

func BenchmarkLOS_Sort4_10(b *testing.B) {
	benchmarkLOS_Sort4_N(b, 10)
}

func BenchmarkLOS_Sort1_100(b *testing.B) {
	benchmarkLOS_Sort1_N(b, 100)
}

func BenchmarkLOS_Sort2_100(b *testing.B) {
	benchmarkLOS_Sort2_N(b, 100)
}

func BenchmarkLOS_Sort3_100(b *testing.B) {
	benchmarkLOS_Sort3_N(b, 100)
}

func BenchmarkLOS_Sort4_100(b *testing.B) {
	benchmarkLOS_Sort4_N(b, 100)
}

func BenchmarkLOS_Sort1_1000(b *testing.B) {
	benchmarkLOS_Sort1_N(b, 1000)
}

func BenchmarkLOS_Sort2_1000(b *testing.B) {
	benchmarkLOS_Sort2_N(b, 1000)
}

func BenchmarkLOS_Sort3_1000(b *testing.B) {
	benchmarkLOS_Sort3_N(b, 1000)
}

func BenchmarkLOS_Sort4_1000(b *testing.B) {
	benchmarkLOS_Sort4_N(b, 1000)
}

func BenchmarkLOS_Sort1_10000(b *testing.B) {
	benchmarkLOS_Sort1_N(b, 10000)
}

func BenchmarkLOS_Sort2_10000(b *testing.B) {
	benchmarkLOS_Sort2_N(b, 10000)
}

func BenchmarkLOS_Sort3_10000(b *testing.B) {
	benchmarkLOS_Sort3_N(b, 10000)
}

func BenchmarkLOS_Sort4_10000(b *testing.B) {
	benchmarkLOS_Sort4_N(b, 10000)
}