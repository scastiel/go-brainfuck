package memory_test

import (
	"sync"
	"testing"

	"github.com/scastiel/go-brainfuck/lib/memory"
)

func TestMemoryAccessInParallel(t *testing.T) {
	const n = 10000
	memory := memory.NewMemory()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			memory.Inc(0)
			memory.Dec(0)
			wg.Done()
		}()
	}

	wg.Wait()
}
