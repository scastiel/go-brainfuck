package memory

import "sync"

type Memory struct {
	mu sync.Mutex
	v  map[int]byte
}

func (m *Memory) Read(position int) byte {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.v[position]
}

func (m *Memory) Write(position int, value byte) {
	m.mu.Lock()
	m.v[position] = value
	m.mu.Unlock()
}

func (m *Memory) Inc(position int) {
	m.mu.Lock()
	m.v[position]++
	m.mu.Unlock()
}

func (m *Memory) Dec(position int) {
	m.mu.Lock()
	m.v[position]--
	m.mu.Unlock()
}

func NewMemory() Memory {
	return Memory{v: make(map[int]byte)}
}
