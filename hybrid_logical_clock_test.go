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
	"testing"
	"time"
)

func Test_HLC_NewHLC(t *testing.T) {
	hlc := (&HLC{}).NewHLC()
	if hlc.physicalTime == 0 {
		t.Errorf("Expected non-zero physical time, got %d", hlc.physicalTime)
	}
	if hlc.logicalTime != 0 {
		t.Errorf("Expected logical time to be 0, got %d", hlc.logicalTime)
	}
}

func Test_HLC_NewHLCFrom(t *testing.T) {
	physicalTime := time.Now().UnixNano()
	logicalTime := uint64(10)
	hlc := (&HLC{}).NewHLCFrom(physicalTime, logicalTime)
	if hlc.physicalTime != physicalTime {
		t.Errorf("Expected physical time %d, got %d", physicalTime, hlc.physicalTime)
	}
	if hlc.logicalTime != logicalTime {
		t.Errorf("Expected logical time %d, got %d", logicalTime, hlc.logicalTime)
	}
}

func Test_HLC_Increase(t *testing.T) {
	hlc := (&HLC{}).NewHLC()
	initialPhysicalTime := hlc.physicalTime
	initialLogicalTime := hlc.logicalTime

	hlc.Increase()
	if hlc.physicalTime != initialPhysicalTime+1 {
		t.Errorf("Expected physical time %d, got %d", initialPhysicalTime+1, hlc.physicalTime)
	}
	if hlc.logicalTime != initialLogicalTime+1 {
		t.Errorf("Expected logical time %d, got %d", initialLogicalTime+1, hlc.logicalTime)
	}
}

func Test_HLC_Update(t *testing.T) {
	hlc := (&HLC{}).NewHLC()
	physicalTime := time.Now().UnixNano() + 1000
	logicalTime := uint64(10)

	hlc.Update(physicalTime, logicalTime)
	if hlc.physicalTime != physicalTime {
		t.Errorf("Expected physical time %d, got %d", physicalTime, hlc.physicalTime)
	}
	if hlc.logicalTime != logicalTime+1 {
		t.Errorf("Expected logical time %d, got %d", logicalTime+1, hlc.logicalTime)
	}
}

func Test_HLC_Time(t *testing.T) {
	hlc := (&HLC{}).NewHLC()
	physicalTime, logicalTime := hlc.Time()
	if physicalTime != hlc.physicalTime {
		t.Errorf("Expected physical time %d, got %d", hlc.physicalTime, physicalTime)
	}
	if logicalTime != hlc.logicalTime {
		t.Errorf("Expected logical time %d, got %d", hlc.logicalTime, logicalTime)
	}
}
