package client

import (
	"errors"
	"sync"
	"time"
)

// ErrNoNodeAvailable is returned if all the nodes are blocked
var ErrNoNodeAvailable = errors.New("no ecs nodes available")

type node struct {
	host         string
	blockedUntil time.Time
}

// NewEcs ...
func NewEcs(nodes []string) *Ecs {
	ecs := &Ecs{nodes: make([]node, len(nodes))}

	for i, n := range nodes {
		ecs.nodes[i] = node{n, time.Time{}}
	}
	return ecs
}

// Ecs ...
type Ecs struct {
	sync.Mutex
	nodes   []node
	current int
}

// NextAvailableNode ...
func (ns *Ecs) NextAvailableNode() (string, int, error) {
	ns.Lock()
	defer ns.Unlock()
	now := time.Now()
	for i := 0; i < len(ns.nodes); i++ {
		ns.current = (ns.current + 1) % len(ns.nodes)
		if now.After(ns.nodes[ns.current].blockedUntil) {
			return ns.nodes[ns.current].host, ns.current, nil
		}
	}
	return "", 0, ErrNoNodeAvailable
}

// BlockNode is to block node in prefined duration
func (ns *Ecs) BlockNode(i int, dur time.Duration) {
	if dur > 0 && i >= 0 && i < len(ns.nodes) {
		ns.Lock()
		defer ns.Unlock()
		ns.nodes[i].blockedUntil = time.Now().Add(dur)
	}
}
