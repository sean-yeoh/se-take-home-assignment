package order

import (
	"time"
)

const processingDuration = 10 * time.Second

type OrderKind string

const (
	Normal OrderKind = "Normal"
	VIP    OrderKind = "VIP"
)

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Processing OrderStatus = "PROCESSING"
	Complete   OrderStatus = "COMPLETE"
)

type Order struct {
	ID          int
	Kind        OrderKind
	Status      OrderStatus
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	BotID       int
}

func (c *Controller) AddNormalOrder() {
	c.addOrder(Normal)
}

func (c *Controller) AddVIPOrder() {
	c.addOrder(VIP)
}

func (c *Controller) addOrder(kind OrderKind) {
	order := &Order{
		ID:        c.nextOrderID(),
		Kind:      kind,
		Status:    Pending,
		CreatedAt: time.Now(),
	}

	switch kind {
	case VIP:
		c.pendingVIP = append(c.pendingVIP, order)
	default:
		c.pendingNormal = append(c.pendingNormal, order)
	}

	c.logEvent("Created %s Order #%d - Status: %s", order.Kind, order.ID, order.Status)
}
