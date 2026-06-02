package cli

import (
	"fmt"
	"io"
	"order-controller/internal/order"
	"time"
)

func RunDemo(output io.Writer) error {
	if _, err := fmt.Fprintln(output, "McDonald's Order Management System - Simulation Results"); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(output); err != nil {
		return err
	}

	controller := order.NewController()
	controller.AddNormalOrder()
	controller.AddVIPOrder()
	controller.AddNormalOrder()
	controller.AddBot()
	controller.AddBot()
	controller.AddVIPOrder()
	time.Sleep(30 * time.Second)
	controller.RemoveBot()

	if _, err := fmt.Fprintf(output, "\n%s\n", controller.FinalStatus()); err != nil {
		return err
	}

	return nil
}
