package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/forget-the-bright/j/internal/pkg/config"
	"github.com/forget-the-bright/j/internal/pkg/download"
	"github.com/mholt/archiver/v3"
	"github.com/urfave/cli/v2"
)

func fundVersion(version string) *config.UrlItem {
	for _, v := range config.Url_Items {
		if strings.Contains(v.Expected, version) {
			return v
		}
	}
	return nil
}

func Install(version string) (err error) {
	ui := fundVersion(version)
	if ui == nil {
		return nil
	}

	filename := filepath.Join(downloadsDir, ui.In.FileName)
	DownloadWithProgress(ui.In.URL, filename)
	targetV := filepath.Join(versionsDir, ui.SimpleName)

	// 解压安装包
	if err = archiver.Unarchive(filename, versionsDir); err != nil {
		fmt.Println(err.Error())
		return cli.Exit(errstring(err), 1)
	}
	// 目录重命名
	if err = os.Rename(filepath.Join(versionsDir, ui.Expected), targetV); err != nil {
		fmt.Println(err.Error())
		return cli.Exit(errstring(err), 1)
	}
	// 重新建立软链接
	_ = os.Remove(goroot)

	if err = mkSymlink(targetV, goroot); err != nil {
		return cli.Exit(errstring(err), 1)
	}
	fmt.Printf("Now using %s\n", ui.Expected)
	return nil
}

// DownloadWithProgress 下载版本另存为指定文件且显示下载进度
func DownloadWithProgress(url, dst string) (size int64, err error) {
	return download.Download(url, dst, os.O_CREATE|os.O_WRONLY, 0644, true)
}
