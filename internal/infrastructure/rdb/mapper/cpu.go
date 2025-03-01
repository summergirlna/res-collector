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

func FromCpu(m *model.Cpu) *Cpu {
	year, month, day, hour, minute, second := m.ParseTimestamp()
	return &Cpu{
		SerialNumber: uuid.NewString(),
		Year:         year,
		Month:        month,
		Day:          day,
		Hour:         hour,
		Minute:       minute,
		Second:       second,
		User:         m.User,
		System:       m.System,
		IOWait:       m.IOWait,
	}
}
