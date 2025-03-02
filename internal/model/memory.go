package model

import "fmt"

type Memory struct {
	Used float64
}

func (m Memory) String() string {
	return fmt.Sprintf("Used: %f", m.Used)
}
