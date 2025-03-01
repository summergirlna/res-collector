package model

import (
	"fmt"
	"time"
)

type Cpu struct {
	Timestamp time.Time
	User      float64
	System    float64
	IOWait    float64
}

func (c Cpu) String() string {
	return fmt.Sprintf("GotAt: %s, User: %f, System: %f, IOWait: %f",
		c.formatTimestamp(), c.User, c.System, c.IOWait)
}

func (c Cpu) formatTimestamp() string {
	return c.Timestamp.Format("2006-01-02 15:04:05")
}

func (c Cpu) ParseTimestamp() (int, int, int, int, int, int) {
	t := c.Timestamp
	return t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()
}
