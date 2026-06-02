package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "order-controller",
		Usage: "McDonald's order controller CLI",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "demo",
				Usage: "run the deterministic demo scenario",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			mode := "interactive"
			if cmd.Bool("demo") {
				mode = "demo"
			}
			fmt.Printf("mode: %s\n", mode)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
