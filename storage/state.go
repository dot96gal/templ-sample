package storage

import "sync"

type State struct {
	count int
	mu    sync.Mutex
}

func NewState() *State {
	return &State{
		count: 1,
	}
}

func (s *State) Count() int {
	return s.count
}

func (s *State) CountUp() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count++
}

func (s *State) CountDown() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count--
}
