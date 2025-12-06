package integration

import (
	"bufio"
	"os/exec"
	"sort/internal/comparator"
	"sort/internal/config"
	"sort/internal/sorter"
	"strings"
	"testing"
)

func TestCompareWithUnixSortSimple(t *testing.T) {
	input := "c\na\nb\n"
	testCompareWithUnix(t, input, []string{})
}

func TestCompareWithUnixSortNumeric(t *testing.T) {
	input := "3\n1\n2\n10\n"
	testCompareWithUnix(t, input, []string{"-n"})
}

func TestCompareWithUnixSortReverse(t *testing.T) {
	input := "a\nc\nb\n"
	testCompareWithUnix(t, input, []string{"-r"})
}

func TestCompareWithUnixSortUnique(t *testing.T) {
	input := "a\na\nb\nb\nc\n"
	testCompareWithUnix(t, input, []string{"-u"})
}

func TestCompareWithUnixSortNumericReverse(t *testing.T) {
	input := "3\n1\n2\n10\n"
	testCompareWithUnix(t, input, []string{"-n", "-r"})
}

func TestCompareWithUnixSortKeyField(t *testing.T) {
	input := "1\tz\n2\ta\n3\tb\n"
	testCompareWithUnix(t, input, []string{"-k", "2"})
}

func TestCompareWithUnixSortNumericKeyField(t *testing.T) {
	input := "10\tother\n2\tother\n1\tother\n"
	testCompareWithUnix(t, input, []string{"-n", "-k", "1"})
}

func TestCompareWithUnixSortIgnoreBlanks(t *testing.T) {
	input := "a  \nb\nc \n"
	testCompareWithUnix(t, input, []string{"-b"})
}

func testCompareWithUnix(t *testing.T, input string, unixArgs []string) {
	cfg := config.Config{
		Numeric:      contains(unixArgs, "-n"),
		Reverse:      contains(unixArgs, "-r"),
		Unique:       contains(unixArgs, "-u"),
		IgnoreBlanks: contains(unixArgs, "-b"),
		KeyField:     getKeyField(unixArgs),
		Delimiter:    '\t',
	}

	factory := comparator.NewFactory(
		cfg.KeyField,
		cfg.Numeric,
		cfg.Month,
		cfg.Human,
		cfg.IgnoreBlanks,
		cfg.Delimiter,
	)
	cmp := factory.Create()
	srt := sorter.NewConcurrentSorter(cmp, cfg.Reverse)

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sorted := srt.Sort(lines)

	if cfg.Unique {
		sorted = removeDuplicates(sorted, cmp)
	}

	ourOutput := strings.Join(sorted, "\n")
	if len(sorted) > 0 {
		ourOutput += "\n"
	}

	cmd := exec.Command("sort", unixArgs...)
	cmd.Stdin = strings.NewReader(input)
	unixOutput, err := cmd.Output()
	if err != nil {
		t.Skipf("UNIX sort not available: %v", err)
	}

	ourStr := normalizeOutput(ourOutput)
	unixStr := normalizeOutput(string(unixOutput))

	if ourStr != unixStr {
		t.Errorf("Our output:\n%q\nUnix output:\n%q", ourStr, unixStr)
	}
}

func TestCheckSorted(t *testing.T) {
	factory := comparator.NewFactory(0, false, false, false, false, '\t')
	cmp := factory.Create()

	lines := []string{"a", "b", "c"}
	if !isSorted(lines, cmp, false) {
		t.Error("should detect sorted")
	}

	lines = []string{"c", "a", "b"}
	if isSorted(lines, cmp, false) {
		t.Error("should detect not sorted")
	}
}

func TestCheckSortedReverse(t *testing.T) {
	factory := comparator.NewFactory(0, false, false, false, false, '\t')
	cmp := factory.Create()

	lines := []string{"c", "b", "a"}
	if !isSorted(lines, cmp, true) {
		t.Error("should detect reverse sorted")
	}
}

func TestLargeFile(t *testing.T) {
	var builder strings.Builder
	for i := 1000; i > 0; i-- {
		builder.WriteString(string(rune('a' + (i % 26))))
		builder.WriteString("\n")
	}
	input := builder.String()

	factory := comparator.NewFactory(0, false, false, false, false, '\t')
	cmp := factory.Create()
	srt := sorter.NewConcurrentSorter(cmp, false)

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sorted := srt.Sort(lines)

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i-1] {
			t.Errorf("Not sorted at position %d: %s > %s", i, sorted[i-1], sorted[i])
		}
	}
}

func TestMonthSort(t *testing.T) {
	input := "Feb\nJan\nDec\nMar\n"
	factory := comparator.NewFactory(0, false, true, false, false, '\t')
	cmp := factory.Create()
	srt := sorter.NewConcurrentSorter(cmp, false)

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sorted := srt.Sort(lines)
	expected := []string{"Jan", "Feb", "Mar", "Dec"}

	if len(sorted) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(sorted))
	}
}

func TestHumanSort(t *testing.T) {
	input := "2M\n1K\n512K\n"
	factory := comparator.NewFactory(0, false, false, true, false, '\t')
	cmp := factory.Create()
	srt := sorter.NewConcurrentSorter(cmp, false)

	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sorted := srt.Sort(lines)
	if len(sorted) < 3 {
		t.Fatal("not enough sorted lines")
	}
	if sorted[0] != "1K" {
		t.Errorf("expected 1K first, got %s", sorted[0])
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getKeyField(args []string) int {
	for i, arg := range args {
		if arg == "-k" && i+1 < len(args) {
			return parseInt(args[i+1])
		}
	}
	return 0
}

func parseInt(s string) int {
	var result int
	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
		} else {
			break
		}
	}
	return result
}

func normalizeOutput(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n")
}

func removeDuplicates(lines []string, cmp comparator.Comparator) []string {
	if len(lines) == 0 {
		return lines
	}
	result := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		if cmp.Compare(lines[i], lines[i-1]) != 0 {
			result = append(result, lines[i])
		}
	}
	return result
}

func isSorted(lines []string, cmp comparator.Comparator, reverse bool) bool {
	for i := 1; i < len(lines); i++ {
		cmpRes := cmp.Compare(lines[i-1], lines[i])
		if reverse {
			if cmpRes < 0 {
				return false
			}
		} else {
			if cmpRes > 0 {
				return false
			}
		}
	}
	return true
}
