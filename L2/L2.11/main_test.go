package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagramsBasic(t *testing.T) {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	got := searchForAnagrams(input)

	if len(got) != 2 {
		t.Fatalf("получили %d: %#v", len(got), got)
	}

	var groups [][]string
	for _, words := range got {
		sorted := append([]string(nil), words...)
		sort.Strings(sorted)
		groups = append(groups, sorted)
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})

	want := [][]string{
		{"листок", "слиток", "столик"},
		{"пятак", "пятка", "тяпка"},
	}

	if !reflect.DeepEqual(groups, want) {
		t.Fatalf("\nожидали: %#v\nполучили: %#v", want, groups)
	}
}

func TestFindAnagramsSingleWordsIgnored(t *testing.T) {
	input := []string{"стол", "окно", "дверь"}

	got := searchForAnagrams(input)

	if len(got) != 0 {
		t.Fatalf(" получили %#v", got)
	}
}

func TestFindAnagramsCaseInsensitive(t *testing.T) {
	input := []string{"Пятак", "пЯткА", "ТЯПКА"}

	got := searchForAnagrams(input)

	if len(got) != 1 {
		t.Fatalf("получили %d: %#v", len(got), got)
	}

	var group []string
	for _, words := range got {
		group = words
		break
	}

	sort.Strings(group)
	want := []string{"пятак", "пятка", "тяпка"}

	if !reflect.DeepEqual(group, want) {
		t.Fatalf("ожидали: %#v\nполучили: %#v", want, group)
	}
}

func TestFindAnagramsWithDuplicates(t *testing.T) {
	input := []string{"пятак", "пятак", "тяпка", "тяпка", "пятка"}

	got := searchForAnagrams(input)

	if len(got) != 1 {
		t.Fatalf("получили %d: %#v", len(got), got)
	}

	var group []string
	for _, words := range got {
		group = words
		break
	}

	sort.Strings(group)
	want := []string{"пятак", "пятка", "тяпка"}

	if !reflect.DeepEqual(group, want) {
		t.Fatalf("ожидали: %#v\nполучили: %#v", want, group)
	}
}
