package mapper

import "res-collector/internal/model"

type Management struct {
	StatsId string `db:"stats_id"`
	Year    int    `db:"year"`
	Month   int    `db:"month"`
	Day     int    `db:"day"`
	Hour    int    `db:"hour"`
	Minute  int    `db:"minute"`
	Second  int    `db:"second"`
}

func NewManagement() *Management {
	return &Management{}
}

func (m *Management) From(mng *model.Management) *Management {
	year, month, day, hour, minute, second := mng.Timestamp.Parse()
	m.StatsId = mng.StatsId.String()
	m.Year = year
	m.Month = month
	m.Day = day
	m.Hour = hour
	m.Minute = minute
	m.Second = second
	return m
}
