package model

import "fmt"

type Management struct {
	StatsId   *StatsId
	Timestamp *Timestamp
}

func NewManagement() *Management {
	return &Management{}
}

func (m *Management) Generate() *Management {
	m.StatsId = NewStatsId().Generate()
	m.Timestamp = NewTimestamp().Now()
	return m
}

func (m *Management) String() string {
	return fmt.Sprintf("StatsId: %s, Timestamp: %s", m.StatsId.String(), m.Timestamp.String())
}
