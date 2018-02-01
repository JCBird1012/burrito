package main

import (
	"os"

	// Other burrito libraries
	"burrito/utils/account"

	// Command line creation helper tool
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "burrito"
	app.Usage = "Chipotle Online Ordering from the comfort of your command line."

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Begin a new order.",
			Action: func(c *cli.Context) error {
				order()
				return nil
			},
		},
		{
			Name:  "logout",
			Usage: "Logout out of Chipotle Online Ordering",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "login",
			Usage: "Login to Chipotle Online Ordering",
			Action: func(c *cli.Context) error {
				account.Login()
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func order() {

}
