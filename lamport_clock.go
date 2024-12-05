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

import "sync/atomic"

type LamportClock struct {
	// Logical time
	time uint64
}

func NewLamportClock() *LamportClock {
	return &LamportClock{}
}

// NewLamportClockFrom create a new LamportClock from the given time
func NewLamportClockFrom(t uint64) *LamportClock {
	return &LamportClock{
		time: atomic.LoadUint64(&t),
	}
}

// Increase increase the logical time by 1
func (c *LamportClock) Increase() {
	atomic.AddUint64(&c.time, 1)
}

// Update update the logical time to the maximum of the current time and the other time
func (c *LamportClock) Update(other uint64) {
	currentTime := atomic.LoadUint64(&c.time)
	if other > currentTime {
		atomic.CompareAndSwapUint64(&c.time, currentTime, other)
	}
}

// Time return the current logical time
func (c *LamportClock) Time() uint64 {
	return atomic.LoadUint64(&c.time)
}
