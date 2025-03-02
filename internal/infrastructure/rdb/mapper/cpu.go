package mapper

import (
	"github.com/google/uuid"
	"res-collector/internal/model"
)

type Cpu struct {
	SerialNumber string  `db:"serial_number"`
	Year         int     `db:"year"`
	Month        int     `db:"month"`
	Day          int     `db:"day"`
	Hour         int     `db:"hour"`
	Minute       int     `db:"minute"`
	Second       int     `db:"second"`
	User         float64 `db:"user_time"`
	System       float64 `db:"system_time"`
	IOWait       float64 `db:"iowait_time"`
}

func NewCpu() *Cpu {
	return &Cpu{}
}

func (c *Cpu) From(m *model.Cpu) *Cpu {
	year, month, day, hour, minute, second := m.Timestamp.Parse()
	c.SerialNumber = uuid.NewString()
	c.Year = year
	c.Month = month
	c.Day = day
	c.Hour = hour
	c.Minute = minute
	c.Second = second
	c.User = m.User
	c.System = m.System
	c.IOWait = m.IOWait
	return c

}
