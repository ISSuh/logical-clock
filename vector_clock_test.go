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

func Test_VectorClock_NewVectorClock(t *testing.T) {
	vc := NewVectorClock()
	if vc == nil {
		t.Errorf("Expected new VectorClock to be created")
	}
	if len(vc.Times()) != 0 {
		t.Errorf("Expected new VectorClock to have no times")
	}
}

func Test_VectorClock_NewVectorClockFrom(t *testing.T) {
	initialClock := map[string]uint64{"node1": 1, "node2": 2}
	vc := NewVectorClockFrom(initialClock)
	if vc == nil {
		t.Errorf("Expected new VectorClock to be created")
	}
	if len(vc.Times()) != 2 {
		t.Errorf("Expected VectorClock to have 2 times")
	}
	if vc.Time("node1") != 1 || vc.Time("node2") != 2 {
		t.Errorf("Expected VectorClock times to match initial values")
	}
}

func Test_VectorClock_Increase(t *testing.T) {
	vc := NewVectorClock()
	vc.Increase("node1")
	if vc.Time("node1") != 1 {
		t.Errorf("Expected node1 time to be 1")
	}
	vc.Increase("node1")
	if vc.Time("node1") != 2 {
		t.Errorf("Expected node1 time to be 2")
	}
}

func Test_VectorClock_ConcurrentIncrease(t *testing.T) {
	vc := NewVectorClock()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vc.Increase("node1")
		}()
	}
	wg.Wait()
	if vc.Time("node1") != 100 {
		t.Errorf("Expected node1 time to be 100")
	}
}

func Test_VectorClock_Update(t *testing.T) {
	vc1 := NewVectorClockFrom(map[string]uint64{"node1": 1, "node2": 2})
	vc2 := NewVectorClockFrom(map[string]uint64{"node1": 3, "node2": 1, "node3": 4})

	vc1.Update(vc2)

	if vc1.Time("node1") != 3 {
		t.Errorf("Expected node1 time to be 3")
	}
	if vc1.Time("node2") != 2 {
		t.Errorf("Expected node2 time to be 2")
	}
	if vc1.Time("node3") != 4 {
		t.Errorf("Expected node3 time to be 4")
	}
}

func Test_VectorClock_Time(t *testing.T) {
	vc := NewVectorClockFrom(map[string]uint64{"node1": 1})
	if vc.Time("node1") != 1 {
		t.Errorf("Expected node1 time to be 1")
	}
	if vc.Time("node2") != 0 {
		t.Errorf("Expected node2 time to be 0")
	}
}

func Test_VectorClock_Times(t *testing.T) {
	initialClock := map[string]uint64{"node1": 1, "node2": 2}
	vc := NewVectorClockFrom(initialClock)
	times := vc.Times()
	if len(times) != 2 {
		t.Errorf("Expected 2 times")
	}
	if times["node1"] != 1 || times["node2"] != 2 {
		t.Errorf("Expected times to match initial values")
	}
}
