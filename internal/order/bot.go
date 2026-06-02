package order

import "time"

type BotState string

const (
	Idle BotState = "IDLE"
	Busy BotState = "PROCESSING"
)

type Bot struct {
	ID    int
	State BotState
	Timer *time.Timer
}
