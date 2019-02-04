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
		los := NewLOS(pair.maxNumber)
		assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
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
}

func TestLOS_Sort1(t *testing.T) {
	for _, pair := range testsSort {
		los := NewLOS(pair.maxNumber)
		los.Sort1()
		assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
	}
}

func BenchmarkLOS_Sort1(b *testing.B) {
	b.ReportAllocs()
	los := NewLOS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort1()
	}
}

func TestLOS_Sort2(t *testing.T) {
	for _, pair := range testsSort {
		los := NewLOS(pair.maxNumber)
		los.Sort2()
		assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
	}
}

func BenchmarkLOS_Sort2(b *testing.B) {
	b.ReportAllocs()
	los := NewLOS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort2()
	}
}

func TestLOS_Sort3(t *testing.T) {
	for _, pair := range testsSort {
		los := NewLOS(pair.maxNumber)
		los.Sort3()
		assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
	}
}

func BenchmarkLOS_Sort3(b *testing.B) {
	b.ReportAllocs()
	los := NewLOS(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		los.Sort3()
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
		los := NewLOS(pair.maxNumber)
		assert.Equal(t, pair.stringLOS, (*los).String())
	}
}
