package build

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var errInvalidJobNumber = errors.New("Error: missing or invalid job number.")

var buildStopCmd = cli.Command{
	Name:      "stop",
	Usage:     "stop a build",
	ArgsUsage: "<repo/name> [build] [job]",
	Action:    buildStop,
}

func buildStop(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}

	number, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return errInvalidBuildNumber
	}

	var job int
	jobArg := c.Args().Get(2)
	if jobArg == "" {
		job = 1
	} else if job, err = strconv.Atoi(jobArg); err != nil {
		return errInvalidJobNumber
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	err = client.BuildStop(owner, name, number, job)
	if err != nil {
		return err
	}

	fmt.Printf("Stopping build %s/%s#%d.%d\n", owner, name, number, job)
	return nil
}
