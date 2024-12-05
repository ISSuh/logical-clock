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
	"testing"
)

func Test_LamportClock_NewLamportClock(t *testing.T) {
	clock := NewLamportClock()
	if clock.Time() != 0 {
		t.Errorf("Expected initial time to be 0, got %d", clock.Time())
	}
}

func Test_LamportClock_NewLamportClockFrom(t *testing.T) {
	initialTime := uint64(10)
	clock := NewLamportClockFrom(initialTime)
	if clock.Time() != initialTime {
		t.Errorf("Expected initial time to be %d, got %d", initialTime, clock.Time())
	}
}

func Test_LamportClock_Increase(t *testing.T) {
	clock := NewLamportClock()
	clock.Increase()
	if clock.Time() != 1 {
		t.Errorf("Expected time to be 1 after increase, got %d", clock.Time())
	}
}

func Test_LamportClock_ConcurrentIncrease(t *testing.T) {
	clock := NewLamportClock()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clock.Increase()
		}()
	}
	wg.Wait()
	if clock.Time() != 100 {
		t.Errorf("Expected time to be 100 after 100 concurrent increases, got %d", clock.Time())
	}
}

func Test_LamportClock_Update(t *testing.T) {
	clock := NewLamportClock()
	clock.Update(5)
	if clock.Time() != 5 {
		t.Errorf("Expected time to be 5 after update, got %d", clock.Time())
	}
	clock.Update(3)
	if clock.Time() != 5 {
		t.Errorf("Expected time to remain 5 after update with smaller value, got %d", clock.Time())
	}
}
