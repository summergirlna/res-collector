package model

import (
	"fmt"
)

type Cpu struct {
	User   float64
	System float64
	IOWait float64
}

func (c Cpu) String() string {
	return fmt.Sprintf("User: %f, System: %f, IOWait: %f", c.User, c.System, c.IOWait)
}
