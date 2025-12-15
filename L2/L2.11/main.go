package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// алгоритм нахождения анаграмм
// каждое слово приводим к нижнему регистру и сортируем символы в слове(по возрастанию)
// перебираем все слова, для каждого ключевого слова создаем множество анаграмм
// если слово уже есть в мапе, добавляем его в множество
// если слова нет в мапе, создаем новую запись в структуре
// возвращаем мапу(ключ - отсортированное слово, значение - множество слов)
func searchForAnagrams(words []string) map[string][]string {
	type dataSet struct {
		originalOrder []string
		set           map[string]struct{}
	}

	sets := make(map[string]*dataSet)

	for _, w := range words {
		if w != "" {
			lower := strings.ToLower(w)
			reassembled := sortedRunes(lower)

			s, ok := sets[reassembled]
			if !ok {
				s = &dataSet{
					originalOrder: make([]string, 0),
					set:           make(map[string]struct{}),
				}
				sets[reassembled] = s
			}

			if _, exists := s.set[lower]; !exists {
				s.set[lower] = struct{}{}
				s.originalOrder = append(s.originalOrder, lower)
			}
		}
	}

	result := make(map[string][]string)

	for _, st := range sets {
		if len(st.originalOrder) < 2 {
			continue
		}

		k := st.originalOrder[0]

		values := make([]string, len(st.originalOrder))
		copy(values, st.originalOrder)
		sort.Strings(values)

		result[k] = values
	}

	return result
}

func sortedRunes(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func main() {
	wrds := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	angs := searchForAnagrams(wrds)

	for k, sets := range angs {
		fmt.Printf("%s: %v\n", k, sets)
	}
}
