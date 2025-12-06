package sorter

import (
	"sort/internal/comparator"
	"testing"
)

func TestConcurrentSorterSmall(t *testing.T) {
	cmp := comparator.NewFactory(0, false, false, false, false, '\t').Create()
	srt := NewConcurrentSorter(cmp, false)
	input := []string{"c", "a", "b"}
	result := srt.Sort(input)
	expected := []string{"a", "b", "c"}
	if !equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConcurrentSorterReverse(t *testing.T) {
	cmp := comparator.NewFactory(0, false, false, false, false, '\t').Create()
	srt := NewConcurrentSorter(cmp, true)
	input := []string{"a", "b", "c"}
	result := srt.Sort(input)
	expected := []string{"c", "b", "a"}
	if !equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConcurrentSorterNumeric(t *testing.T) {
	cmp := comparator.NewFactory(0, true, false, false, false, '\t').Create()
	srt := NewConcurrentSorter(cmp, false)
	input := []string{"10", "2", "1"}
	result := srt.Sort(input)
	expected := []string{"1", "2", "10"}
	if !equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConcurrentSorterLarge(t *testing.T) {
	cmp := comparator.NewFactory(0, false, false, false, false, '\t').Create()
	srt := NewConcurrentSorter(cmp, false)
	input := make([]string, 50000)
	for i := 0; i < 50000; i++ {
		input[i] = string(rune('a' + (i % 26)))
	}
	result := srt.Sort(input)
	if !isSorted(result, cmp) {
		t.Error("result should be sorted")
	}
}

func TestStringPool(t *testing.T) {
	pool := NewStringPool()
	builder1 := pool.Get()
	builder1.WriteString("test")
	pool.Put(builder1)
	builder2 := pool.Get()
	if builder2.Len() != 0 {
		t.Error("pool should reset builder")
	}
	pool.Put(builder2)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSorted(lines []string, cmp comparator.Comparator) bool {
	for i := 1; i < len(lines); i++ {
		if cmp.Compare(lines[i-1], lines[i]) > 0 {
			return false
		}
	}
	return true
}
