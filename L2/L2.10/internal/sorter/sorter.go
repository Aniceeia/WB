package sorter

import (
	"container/heap"
	"sort"
	"sync"

	"sort/internal/comparator"
)

type Sorter interface {
	Sort(lines []string) []string
}

type ConcurrentSorter struct {
	comparator comparator.Comparator
	reverse    bool
	chunkSize  int
	pool       *StringPool
}

func NewConcurrentSorter(cmp comparator.Comparator, reverse bool) *ConcurrentSorter {
	return &ConcurrentSorter{
		comparator: cmp,
		reverse:    reverse,
		chunkSize:  10000,
		pool:       NewStringPool(),
	}
}

func (s *ConcurrentSorter) Sort(lines []string) []string {
	if len(lines) <= s.chunkSize {
		return s.sortChunk(lines)
	}

	chunks := s.splitChunks(lines)
	sortedChunks := s.sortChunksConcurrently(chunks)
	return s.mergeChunks(sortedChunks)
}

func (s *ConcurrentSorter) sortChunk(lines []string) []string {
	result := make([]string, len(lines))
	copy(result, lines)
	sort.Slice(result, func(i, j int) bool {
		return s.comparator.Compare(result[i], result[j]) < 0
	})
	if s.reverse {
		s.reverseSlice(result)
	}
	return result
}

func (s *ConcurrentSorter) splitChunks(lines []string) [][]string {
	var chunks [][]string
	for i := 0; i < len(lines); i += s.chunkSize {
		end := i + s.chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		chunks = append(chunks, lines[i:end])
	}
	return chunks
}

func (s *ConcurrentSorter) sortChunksConcurrently(chunks [][]string) [][]string {
	var wg sync.WaitGroup
	chunkChan := make(chan int, len(chunks))
	workerCount := 4

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range chunkChan {
				chunks[idx] = s.sortChunk(chunks[idx])
			}
		}()
	}

	for i := range chunks {
		chunkChan <- i
	}
	close(chunkChan)
	wg.Wait()

	return chunks
}

type mergeItem struct {
	line  string
	chunk int
	pos   int
}

type mergeHeap struct {
	items      []mergeItem
	comparator comparator.Comparator
}

func (h mergeHeap) Len() int { return len(h.items) }
func (h mergeHeap) Less(i, j int) bool {
	return h.comparator.Compare(h.items[i].line, h.items[j].line) < 0
}
func (h mergeHeap) Swap(i, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *mergeHeap) Push(x interface{}) {
	h.items = append(h.items, x.(mergeItem))
}

func (h *mergeHeap) Pop() interface{} {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[:n-1]
	return item
}

func (s *ConcurrentSorter) mergeChunks(chunks [][]string) []string {
	if len(chunks) == 0 {
		return []string{}
	}
	if len(chunks) == 1 {
		return chunks[0]
	}

	h := &mergeHeap{
		items:      make([]mergeItem, 0),
		comparator: s.comparator,
	}
	heap.Init(h)

	for i, ch := range chunks {
		if len(ch) > 0 {
			heap.Push(h, mergeItem{
				line:  ch[0],
				chunk: i,
				pos:   0,
			})
		}
	}

	var result []string
	builder := s.pool.Get()
	defer s.pool.Put(builder)

	for h.Len() > 0 {
		item := heap.Pop(h).(mergeItem)
		result = append(result, item.line)

		ch := chunks[item.chunk]
		if item.pos+1 < len(ch) {
			heap.Push(h, mergeItem{
				line:  ch[item.pos+1],
				chunk: item.chunk,
				pos:   item.pos + 1,
			})
		}
	}

	if s.reverse {
		s.reverseSlice(result)
	}

	return result
}

func (s *ConcurrentSorter) reverseSlice(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}
