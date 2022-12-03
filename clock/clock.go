package clock

import (
	"time"
)

// Define the Clock type here.

type Clock time.Time

func New(h, m int) Clock {
	return Clock(time.Date(2000, 01, 01, h, m, 00, 0, time.Local))
}

func (c Clock) Add(m int) Clock {
	return Clock(time.Time(c).Add(time.Duration(m) * time.Minute))
}

func (c Clock) Subtract(m int) Clock {
	return Clock(time.Time(c).Add(time.Duration(-m) * time.Minute))
}

func (c Clock) String() string {
	return time.Time(c).Format("15:04")
}
