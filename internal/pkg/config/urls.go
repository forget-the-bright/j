package config

import (
	"errors"
	"runtime"

	cli "github.com/urfave/cli/v2"
)

type JavaFileItem struct {
	FileName string
	URL      string
	Sha256   string
}
type UrlItem struct {
	In         *JavaFileItem
	Expected   string
	SimpleName string
}

var Url_Items []*UrlItem

func ReverseArray(arr []*UrlItem) []*UrlItem {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func init() {
	switch runtime.GOOS {
	case "linux":
		Url_Items = linux_Url_Items
	case "windows":
		Url_Items = windows_Url_Items
	default:
		cli.Exit(errors.New(runtime.GOOS+" OS is not supported"), 1)
	}
}
