package cli

import "github.com/urfave/cli/v2"

var (
	commands = []*cli.Command{
		{
			Name:      "ls",
			Usage:     "List installed versions",
			UsageText: "j ls",
			Action:    list,
		},
		{
			Name:      "use",
			Usage:     "Switch to specified version",
			UsageText: "j use <version>",
			Action:    use,
		},
	}
)
