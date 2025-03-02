package model

import "github.com/google/uuid"

type StatsId struct {
	Value uuid.UUID
}

func NewStatsId() *StatsId {
	return &StatsId{}
}

func (i *StatsId) Generate() *StatsId {
	i.Value = uuid.New()
	return i
}

func (i *StatsId) String() string {
	return i.Value.String()
}
