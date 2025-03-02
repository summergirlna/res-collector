package mapper

import (
	"github.com/google/uuid"
	"res-collector/internal/model"
)

type Memory struct {
	SerialNumber string  `db:"serial_number"`
	Year         int     `db:"year"`
	Month        int     `db:"month"`
	Day          int     `db:"day"`
	Hour         int     `db:"hour"`
	Minute       int     `db:"minute"`
	Second       int     `db:"second"`
	Used         float64 `db:"application_used"`
}

func NewMemory() *Memory {
	return &Memory{}
}

func (mm *Memory) From(m *model.Memory) *Memory {
	year, month, day, hour, minute, second := m.Timestamp.Parse()
	mm.SerialNumber = uuid.NewString()
	mm.Year = year
	mm.Month = month
	mm.Day = day
	mm.Hour = hour
	mm.Minute = minute
	mm.Second = second
	mm.Used = m.Used
	return mm
}
