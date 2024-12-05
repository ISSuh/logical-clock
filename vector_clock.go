/*
MIT License

Copyright (c) 2024 ISSuh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package logicalclock

import "sync"

type VectorClock struct {
	m sync.RWMutex

	// Logical time
	clock map[string]uint64
}

func NewVectorClock() *VectorClock {
	return &VectorClock{
		clock: make(map[string]uint64),
	}
}

// NewVectorClockFrom create a new VectorClock from the given clock
func NewVectorClockFrom(clock map[string]uint64) *VectorClock {
	return &VectorClock{
		clock: clock,
	}
}

// Increase increase the logical time by 1 on the given node
func (c *VectorClock) Increase(node string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.clock[node]++
}

// Update update the logical time to the maximum of the current time and the other time
func (c *VectorClock) Update(other *VectorClock) {
	c.m.Lock()
	defer c.m.Unlock()

	times := other.Times()
	for node, time := range times {
		if time > c.clock[node] {
			c.clock[node] = time
		}
	}
}

// Time return the current logical time on the given node
func (c *VectorClock) Time(node string) uint64 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.clock[node]
}

// Times return the current logical time
func (c *VectorClock) Times() map[string]uint64 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.clock
}
