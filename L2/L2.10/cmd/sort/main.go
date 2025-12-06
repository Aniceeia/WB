package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"sort/internal/comparator"
	"sort/internal/config"
	"sort/internal/sorter"
)

func main() {
	cfg := config.ParseFlags()

	var input io.Reader
	if cfg.InputFile != "" {
		file, err := os.Open(cfg.InputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	err := sort(input, os.Stdout, cfg)
	if err != nil {
		os.Exit(1)
	}
}

// sort выполняет основную логику сортировки.
// Алгоритм:
// 1. Читаем все строки из входного потока
// 2. Создаем компаратор на основе флагов конфигурации
// 3. Если флаг -c установлен, проверяем отсортированность и выходим
// 4. Сортируем строки используя concurrent merge sort
// 5. Если флаг -u установлен, удаляем дубликаты
// 6. Выводим результат в выходной поток
func sort(input io.Reader, output io.Writer, cfg config.Config) error {
	lines, err := readLines(input)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
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

	if cfg.Check {
		if isSorted(lines, cmp, cfg.Reverse) {
			return nil
		}
		fmt.Fprintln(output, "disorder")
		return fmt.Errorf("not sorted")
	}

	srt := sorter.NewConcurrentSorter(cmp, cfg.Reverse)
	sorted := srt.Sort(lines)

	if cfg.Unique {
		sorted = removeDuplicates(sorted, cmp)
	}

	for _, line := range sorted {
		fmt.Fprintln(output, line)
	}

	return nil
}

func readLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
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
