package comparator

import (
	"strconv"
	"strings"
)

type NumericComparator struct {
	keyField     int
	ignoreBlanks bool
	delimiter    rune
}

func (c *NumericComparator) Compare(a, b string) int {
	keyA := extractKey(a, c.keyField, c.delimiter)
	keyB := extractKey(b, c.keyField, c.delimiter)
	keyA = cleanKey(keyA, c.ignoreBlanks)
	keyB = cleanKey(keyB, c.ignoreBlanks)

	numA, errA := parseNumber(keyA)
	numB, errB := parseNumber(keyB)

	if errA != nil && errB != nil {
		return strings.Compare(keyA, keyB)
	}
	if errA != nil {
		return -1
	}
	if errB != nil {
		return 1
	}

	if numA < numB {
		return -1
	}
	if numA > numB {
		return 1
	}
	return 0
}

func parseNumber(s string) (float64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseFloat(s, 64)
}
