package main

import (
	"time"
)

type Ping struct {
	ID         uint64
	SourceIP   string
	DeviceName string
	ReceivedAt time.Time `gorm:"index"`
}
