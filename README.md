# logical-clock

Implements simple logical clocks, including Lamport Clock, Vector Clock, and Hybrid Logical Clock (HLC).

## Clocks Implemented

- **Lamport Clock**: A logical clock that uses a single counter to order events in a distributed system.
- **Vector Clock**: A logical clock that uses a vector of counters to order events in a distributed system.
- **Hybrid Logical Clock (HLC)**: A logical clock that combines physical and logical clocks to order events in a distributed system.

## Usage

```bash
go get github.com/ISSuh/logical-clock
```

#### Lamport Clock

```go
import (
    logicalclock "github.com/ISSuh/logical-clock"
)

func func1() {
    clock := logicalclock.NewLamportClock()
    clock.Increase()

    currentTime := clock.Time()
    clock.Update(5)
}

func func2(received uint64) {
    clock := logicalclock.NewLamportClockFrom(received)
    clock.Increase()
    currentTime := clock.Time()
}
```

#### Vector Clock

```go
import (
    logicalclock "github.com/ISSuh/logical-clock"
)

func func1() {
    clock := logicalclock.NewVectorClock()
    clock.Increase("node1")
    currentTime := clock.Time("node1")

    otherClock := logicalclock.NewVectorClockFrom(map[string]uint64{"node1": 2})
    clock.Update(otherClock)
}
```

#### Hybrid Logical Clock (HLC)

```go
import (
    logicalclock "github.com/ISSuh/logical-clock"
)

func func1() {
    clock := logicalclock.NewHLC()
    clock.Increase()
    physicalTime, logicalTime := clock.Time()
    clock.Update(time.Now().UnixNano(), 10)
}

func func2(receivedPT, receivedLT uint64) {
    clock := logicalclock.NewHLCFrom(receivedPT, receivedLT)
    clock.Increase()
    physicalTime, logicalTime := clock.Time()
}
```