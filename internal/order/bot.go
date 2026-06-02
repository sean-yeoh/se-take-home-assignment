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

func (c *Controller) RemoveBot() {
	if len(c.bots) == 0 {
		c.logEvent("No bots available to remove")
		return
	}

	index := len(c.bots) - 1
	bot := c.bots[index]
	c.bots = c.bots[:index]

	if bot.Timer != nil {
		bot.Timer.Stop()
	}

	c.logEvent("Bot #%d destroyed while %s", bot.ID, bot.State)
}
