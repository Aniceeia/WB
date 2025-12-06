package comparator

import (
	"strconv"
	"strings"
	"unicode"
)

type HumanComparator struct {
	keyField     int
	ignoreBlanks bool
	delimiter    rune
}

var multipliers = map[byte]float64{
	'k': 1024,
	'K': 1024,
	'm': 1024 * 1024,
	'M': 1024 * 1024,
	'g': 1024 * 1024 * 1024,
	'G': 1024 * 1024 * 1024,
	't': 1024 * 1024 * 1024 * 1024,
	'T': 1024 * 1024 * 1024 * 1024,
}

func (c *HumanComparator) Compare(a, b string) int {
	keyA := extractKey(a, c.keyField, c.delimiter)
	keyB := extractKey(b, c.keyField, c.delimiter)
	keyA = cleanKey(keyA, c.ignoreBlanks)
	keyB = cleanKey(keyB, c.ignoreBlanks)

	valA, errA := parseHumanSize(keyA)
	valB, errB := parseHumanSize(keyB)

	if errA != nil && errB != nil {
		return strings.Compare(keyA, keyB)
	}
	if errA != nil {
		return -1
	}
	if errB != nil {
		return 1
	}

	if valA < valB {
		return -1
	}
	if valA > valB {
		return 1
	}
	return 0
}

func parseHumanSize(s string) (float64, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, nil
	}

	lastChar := s[len(s)-1]
	multiplier, hasMultiplier := multipliers[lastChar]

	if hasMultiplier {
		numStr := s[:len(s)-1]
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, err
		}
		return num * multiplier, nil
	}

	var numStr strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || r == '.' || r == '-' || r == '+' {
			numStr.WriteRune(r)
		} else {
			break
		}
	}

	if numStr.Len() == 0 {
		return 0, nil
	}

	return strconv.ParseFloat(numStr.String(), 64)
}
