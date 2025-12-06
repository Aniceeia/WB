package comparator

import (
	"strings"
)

type MonthComparator struct {
	keyField     int
	ignoreBlanks bool
	delimiter    rune
}

var monthNames = []string{
	"jan", "feb", "mar", "apr", "may", "jun",
	"jul", "aug", "sep", "oct", "nov", "dec",
}

func (c *MonthComparator) Compare(a, b string) int {
	keyA := extractKey(a, c.keyField, c.delimiter)
	keyB := extractKey(b, c.keyField, c.delimiter)
	keyA = cleanKey(keyA, c.ignoreBlanks)
	keyB = cleanKey(keyB, c.ignoreBlanks)

	monthA := parseMonth(keyA)
	monthB := parseMonth(keyB)

	if monthA == -1 && monthB == -1 {
		return strings.Compare(keyA, keyB)
	}
	if monthA == -1 {
		return -1
	}
	if monthB == -1 {
		return 1
	}

	if monthA < monthB {
		return -1
	}
	if monthA > monthB {
		return 1
	}
	return 0
}

func parseMonth(s string) int {
	s = strings.ToLower(strings.TrimSpace(s))
	for i, month := range monthNames {
		if strings.HasPrefix(s, month) {
			return i
		}
	}
	return -1
}
