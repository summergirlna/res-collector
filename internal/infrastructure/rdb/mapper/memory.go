package mapper

import (
	"res-collector/internal/model"
)

type Memory struct {
	StatsId string  `db:"stats_id"`
	Used    float64 `db:"application_used"`
}

func NewMemory() *Memory {
	return &Memory{}
}

func (mm *Memory) From(m *model.Memory, mng *model.Management) *Memory {
	mm.StatsId = mng.StatsId.String()
	mm.Used = m.Used
	return mm
}
