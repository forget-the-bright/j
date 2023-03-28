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
			Name:      "ls-all",
			Usage:     "List All versions",
			UsageText: "j ls-all",
			Action:    listAll,
		},
		{
			Name:      "install",
			Usage:     "install versions",
			UsageText: "j install <version>",
			Action:    install,
		},
		{
			Name:      "use",
			Usage:     "Switch to specified version",
			UsageText: "j use <version>",
			Action:    use,
		},
		{
			Name:      "uninstall",
			Usage:     "Uninstall a version",
			UsageText: "j uninstall <version>",
			Action:    uninstall,
		},
		{
			Name:      "clean",
			Usage:     "Remove files from the package download directory",
			UsageText: "j clean",
			Action:    clean,
		},
	}
)
