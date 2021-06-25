package repo

import (
	"fmt"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var repoMoveCmd = cli.Command{
	Name:      "mv",
	Usage:     "move the repository to a new location",
	ArgsUsage: "<repo/name> <newRepo/newName>",
	Action:    repoMove,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "unsafe",
			Usage: "validate updating the full repo name is unsafe - build status in repos will reflect old name",
		},
	},
}

func repoMove(c *cli.Context) error {
	repo := c.Args().Get(0)
	newRepo := c.Args().Get(1)
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}
	_, _, newErr := internal.ParseRepo(newRepo)
	if newErr != nil {
		return newErr
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	unsafe := c.Bool("unsafe");

	if !unsafe {
		fmt.Printf("Moving the repo is an unsafe operation that could put your repository in an inconsistent state. You must also manually sync the repos before running this operation. Please use --unsafe to proceed\n")
		return nil
	}

	if err := client.RepoMove(owner, name, newRepo); err != nil {
		return err
	}

	fmt.Printf("Successfully moved repository from %s/%s to %s\n", owner, name, newRepo)
	return nil
}
