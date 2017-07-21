package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				stopValue := c.Int("stop")

				if stopValue < 0 {
					fmt.Println("specify a value greater than 0 to count up to")
				}

				for i := 0; i <= stopValue; i++ {
					fmt.Println(i)
				}

				return nil
			},
		},

		{
			Name:      "down",
			ShortName: "d",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {

				start := c.Int("start")

				if start <= 0 {
					fmt.Println("start value can not be less than 0")
				}

				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
