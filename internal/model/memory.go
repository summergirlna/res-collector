package model

import "fmt"

type Memory struct {
	Timestamp *Timestamp
	Used      float64
}

func (m Memory) String() string {
	return fmt.Sprintf("Timestamp: %s, Used: %f", m.Timestamp.String(), m.Used)
}
