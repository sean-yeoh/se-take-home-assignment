package order

import (
	"fmt"
	"sync"
	"time"
)

type Controller struct {
	mu             sync.Mutex
	orderIDCounter int
	botIDCounter   int
	pendingVIP     []*Order
	pendingNormal  []*Order
	processing     []*Order
	completed      []*Order
	bots           []Bot
	events         []string
}

func NewController() *Controller {
	controller := &Controller{orderIDCounter: 1000}
	controller.logEvent("System initialized with 0 bots")
	return controller
}

func (c *Controller) logEvent(format string, args ...any) {
	line := fmt.Sprintf("[%s] %s",
		time.Now().Format("15:04:05"),
		fmt.Sprintf(format, args...),
	)

	c.events = append(c.events, line)
	fmt.Println(line)
}

func (c *Controller) nextBotID() int {
	c.botIDCounter++
	return c.botIDCounter
}

func (c *Controller) nextOrderID() int {
	c.orderIDCounter++
	return c.orderIDCounter
}
