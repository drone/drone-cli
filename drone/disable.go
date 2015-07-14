package main

import (
	"github.com/codegangsta/cli"
	"github.com/drone/drone-go/drone"
)

// NewDisableCommand returns the CLI command for "disable".
func NewDisableCommand() cli.Command {
	return cli.Command{
		Name:  "disable",
		Usage: "disable a repository",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) {
			handle(c, disableCommandFunc)
		},
	}
}

// disableCommandFunc executes the "disable" command.
func disableCommandFunc(c *cli.Context, client *drone.Client) error {
	host, owner, name := parseRepo(c.Args())
	return client.Repos.Disable(host, owner, name)
}
