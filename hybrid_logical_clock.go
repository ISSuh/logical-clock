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

import (
	"sync"
	"time"
)

// hybrid logical clock
type HLC struct {
	m sync.RWMutex

	// Physical time. Unix nano
	physicalTime int64

	// Logical time
	logicalTime uint64
}

func (c *HLC) NewHLC() *HLC {
	return &HLC{
		physicalTime: time.Now().UnixNano(),
		logicalTime:  0,
	}
}

// NewHLCFrom create a new HLC from the given time
func (c *HLC) NewHLCFrom(physicalTime int64, logicalTime uint64) *HLC {
	return &HLC{
		physicalTime: physicalTime,
		logicalTime:  logicalTime,
	}
}

// Increase increase the physical time and logical time by 1
func (c *HLC) Increase() {
	c.m.Lock()
	defer c.m.Unlock()

	c.physicalTime = time.Now().UnixNano()
	c.logicalTime++
}

// Update update the physical time and logical time to the maximum of the current time and the other time
func (c *HLC) Update(physicalTime int64, logicalTime uint64) {
	c.m.Lock()
	defer c.m.Unlock()

	currentPhysicalTime := time.Now().UnixNano()
	if physicalTime < currentPhysicalTime {
		c.physicalTime = currentPhysicalTime
		c.logicalTime = 0
	} else {
		c.physicalTime = physicalTime

		if logicalTime >= c.logicalTime {
			c.logicalTime = logicalTime + 1
		} else {
			c.logicalTime++
		}
	}
}

// Time return the physical time and logical time
func (c *HLC) Time() (int64, uint64) {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.physicalTime, c.logicalTime
}
