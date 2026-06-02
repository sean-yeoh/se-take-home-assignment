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

func (c *Controller) AddBot() {
	bot := Bot{
		ID:    c.nextBotID(),
		State: Idle,
	}

	c.bots = append(c.bots, bot)
	c.logEvent("Bot #%d created - Status: %s", bot.ID, "ACTIVE")
}
