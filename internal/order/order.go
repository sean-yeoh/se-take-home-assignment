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
