package mapper

import (
	"res-collector/internal/model"
)

type Cpu struct {
	StatsId string  `db:"stats_id"`
	User    float64 `db:"user_time"`
	System  float64 `db:"system_time"`
	IOWait  float64 `db:"iowait_time"`
}

func NewCpu() *Cpu {
	return &Cpu{}
}

func (c *Cpu) From(m *model.Cpu, mng *model.Management) *Cpu {
	c.StatsId = mng.StatsId.String()
	c.User = m.User
	c.System = m.System
	c.IOWait = m.IOWait
	return c

}
