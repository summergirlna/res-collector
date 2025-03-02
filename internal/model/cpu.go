package model

import (
	"fmt"
)

type Cpu struct {
	Timestamp *Timestamp
	User      float64
	System    float64
	IOWait    float64
}

func (c Cpu) String() string {
	return fmt.Sprintf("Timestamp: %s, User: %f, System: %f, IOWait: %f",
		c.Timestamp.String(), c.User, c.System, c.IOWait)
}
