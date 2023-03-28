package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/forget-the-bright/j/internal/pkg/config"
	"github.com/k0kubun/go-ansi"
	"github.com/urfave/cli/v2"
)

func reverseArray(arr []*config.UrlItem) []*config.UrlItem {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
func mathVersionLength(version string) string {
	if len(version) <= 1 {
		return version + " "
	}
	return version
}
func listAll(*cli.Context) (err error) {
	use_version := inuse(goroot)
	out := ansi.NewAnsiStdout()
	color.New(color.FgGreen).Fprintf(out, " %s\n", "version      info")
	for _, v := range reverseArray(config.Url_Items) {
		if v.SimpleName == use_version { //strings.Contains(v.SimpleName, version)
			color.New(color.FgGreen).Fprintf(out, "*  %s\n", mathVersionLength(v.SimpleName)+"      "+v.Expected)
		} else {
			fmt.Fprintf(out, "   %s\n", mathVersionLength(v.SimpleName)+"      "+v.Expected)
		}
	}
	return nil
}
