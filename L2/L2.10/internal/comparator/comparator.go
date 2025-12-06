package comparator

import (
	"strings"
	"unicode"
)

type Comparator interface {
	Compare(a, b string) int
}

type Factory struct {
	keyField     int
	numeric      bool
	month        bool
	human        bool
	ignoreBlanks bool
	delimiter    rune
}

func NewFactory(keyField int, numeric, month, human, ignoreBlanks bool, delimiter rune) *Factory {
	return &Factory{
		keyField:     keyField,
		numeric:      numeric,
		month:        month,
		human:        human,
		ignoreBlanks: ignoreBlanks,
		delimiter:    delimiter,
	}
}

func (f *Factory) Create() Comparator {
	if f.month {
		return &MonthComparator{
			keyField:     f.keyField,
			ignoreBlanks: f.ignoreBlanks,
			delimiter:    f.delimiter,
		}
	}
	if f.human {
		return &HumanComparator{
			keyField:     f.keyField,
			ignoreBlanks: f.ignoreBlanks,
			delimiter:    f.delimiter,
		}
	}
	if f.numeric {
		return &NumericComparator{
			keyField:     f.keyField,
			ignoreBlanks: f.ignoreBlanks,
			delimiter:    f.delimiter,
		}
	}
	return &LexicoComparator{
		keyField:     f.keyField,
		ignoreBlanks: f.ignoreBlanks,
		delimiter:    f.delimiter,
	}
}

func extractKey(line string, field int, delim rune) string {
	if field == 0 {
		return line
	}
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == delim
	})
	if field-1 < len(fields) {
		return fields[field-1]
	}
	return ""
}

func cleanKey(key string, ignoreBlanks bool) string {
	if ignoreBlanks {
		return strings.TrimRightFunc(key, unicode.IsSpace)
	}
	return key
}

type LexicoComparator struct {
	keyField     int
	ignoreBlanks bool
	delimiter    rune
}

func (c *LexicoComparator) Compare(a, b string) int {
	keyA := extractKey(a, c.keyField, c.delimiter)
	keyB := extractKey(b, c.keyField, c.delimiter)
	keyA = cleanKey(keyA, c.ignoreBlanks)
	keyB = cleanKey(keyB, c.ignoreBlanks)
	return strings.Compare(keyA, keyB)
}
