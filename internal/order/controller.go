package order

import (
	"fmt"
	"strings"
	"time"
)

type Controller struct {
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

func (c *Controller) StatusTable() string {
	pending := orderLabels(c.pendingOrders(), false)
	completed := orderLabels(c.completed, false)
	processing := orderLabels(c.processing, true)
	bots := botLabels(c.bots)

	rowCount := max(len(pending), len(completed), len(processing), len(bots))
	if rowCount == 0 {
		rowCount = 1
	}

	rows := [][]string{
		{"Pending", "Completed", "Processing", "Bots"},
		{"-------", "---------", "----------", "----"},
	}

	for i := 0; i < rowCount; i++ {
		rows = append(rows, []string{
			valueAt(pending, i),
			valueAt(completed, i),
			valueAt(processing, i),
			valueAt(bots, i),
		})
	}

	widths := columnWidths(rows)
	lines := make([]string, 0, len(rows))
	for _, row := range rows {
		lines = append(lines, fmt.Sprintf(
			"%-*s | %-*s | %-*s | %-*s",
			widths[0], row[0],
			widths[1], row[1],
			widths[2], row[2],
			widths[3], row[3],
		))
	}

	lines = append(lines, "")
	lines = append(lines, fmt.Sprintf(
		"Summary: %d pending, %d completed, %d processing, %d bots",
		len(c.pendingVIP)+len(c.pendingNormal),
		len(c.completed),
		len(c.processing),
		len(c.bots),
	))

	return strings.Join(lines, "\n")
}

func (c *Controller) pendingOrders() []*Order {
	pending := make([]*Order, 0, len(c.pendingVIP)+len(c.pendingNormal))
	pending = append(pending, c.pendingVIP...)
	pending = append(pending, c.pendingNormal...)
	return pending
}

func orderLabels(orders []*Order, includeBot bool) []string {
	labels := make([]string, 0, len(orders))
	for _, order := range orders {
		label := fmt.Sprintf("Order #%d %s", order.ID, order.Kind)
		if includeBot && order.BotID != 0 {
			label = fmt.Sprintf("%s - Bot #%d", label, order.BotID)
		}
		labels = append(labels, label)
	}
	return labels
}

func botLabels(bots []Bot) []string {
	labels := make([]string, 0, len(bots))
	for _, bot := range bots {
		labels = append(labels, fmt.Sprintf("Bot #%d - %s", bot.ID, bot.State))
	}
	return labels
}

func valueAt(values []string, index int) string {
	if index >= len(values) {
		return "-"
	}
	return values[index]
}

func columnWidths(rows [][]string) []int {
	widths := []int{0, 0, 0, 0}
	for _, row := range rows {
		for i, value := range row {
			if len(value) > widths[i] {
				widths[i] = len(value)
			}
		}
	}
	return widths
}

func max(values ...int) int {
	result := 0
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result
}
