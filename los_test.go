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

func TestLOS_Sort(t *testing.T) {
	var tests = []testpairs{
		{1, []int{1}, "only one element: no sort"},
		{2, []int{1, 2}, "two elements: no sort"},
		{3, []int{1, 3, 2}, "three elements: one swap"},
		{4, []int{1, 4, 2, 3}, "four elements: one swap"},
		{5, []int{1, 5, 2, 4, 3}, "usually sort"},
		{8, []int{1, 8, 2, 7, 3, 6, 4, 5}, "from readme"},
	}
	for _, pair := range tests {
		los := NewLOS(pair.maxNumber)
		los.Sort()
		assert.Equal(t, newTestLOS(pair.elements), los, pair.name)
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