package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/forget-the-bright/j/internal/pkg/collector"
	"github.com/k0kubun/go-ansi"
	"github.com/urfave/cli/v2"
)

func remoteVersionLength(version string) string {
	yu := 8 - len(version)
	for i := 0; i < yu; i++ {
		version += " "
	}
	return version
}

func listRemote(*cli.Context) (err error) {
	use_version := inuse(goroot)
	out := ansi.NewAnsiStdout()
	rs := collector.ConvertCollectorToUrlItem(collector.GetOpenJDKArchiveReleasesInfo(), false)
	color.New(color.FgGreen).Fprintf(out, " %s\n", " version                    info")
	for _, v := range rs {
		if v.SimpleName == use_version { //strings.Contains(v.SimpleName, version)
			color.New(color.FgGreen).Fprintf(out, "*  %s\n", remoteVersionLength(v.SimpleName)+"      "+v.Expected)
		} else {
			fmt.Fprintf(out, "   %s\n", remoteVersionLength(v.SimpleName)+"      "+v.Expected)
		}
	}
	return nil
}
