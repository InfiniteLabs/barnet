package stylist

import (
	"time"
)

type Shift struct {
	startTime time.Time
	duration  time.Duration
}
