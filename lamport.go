package main

import "sync"

type Lamport struct {
	T  int32
	Mu sync.Mutex
}

func (l *Lamport) Increment() {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	l.T++
}

func (l *Lamport) MaxAndIncrement(other int32) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	if l.T < other {
		l.T = other
	}

	l.T++
}
