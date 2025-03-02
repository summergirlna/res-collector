package model

import "time"

type Timestamp struct {
	Value time.Time
}

func NewTimestamp() *Timestamp {
	return &Timestamp{}
}

func (t Timestamp) String() string {
	return t.Value.Format("2006-01-02 15:04:05")
}

func (t Timestamp) Parse() (int, int, int, int, int, int) {
	ts := t.Value
	return ts.Year(), int(ts.Month()), ts.Day(), ts.Hour(), ts.Minute(), ts.Second()
}

func (t Timestamp) Now() *Timestamp {
	ts := NewTimestamp()
	ts.Value = time.Now()
	return ts
}
