package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/forget-the-bright/j/internal/pkg/archiver"
	"github.com/forget-the-bright/j/internal/pkg/config"
	"github.com/forget-the-bright/j/internal/pkg/download"

	//"github.com/mholt/archiver/v3"
	"github.com/urfave/cli/v2"
)

func fundVersion(version string) *config.UrlItem {
	for _, v := range config.Url_Items {
		if v.SimpleName == version { //strings.Contains(v.SimpleName, version)
			return v
		}
	}
	return nil
}

func downloadAndInstall(version string) (err error) {
	ui := fundVersion(version)
	if ui == nil {
		return cli.Exit(errors.New(version+" version is not supported"), 1)
	}

	filename := filepath.Join(downloadsDir, ui.In.FileName)
	//判断本地有没有安装包 没有就进入下载
	if _, err := os.Stat(filename); err != nil {
		DownloadWithProgress(ui.In.URL, filename)
	}

	//获取解压目标目录
	targetV := filepath.Join(versionsDir, ui.SimpleName)

	//判断判断解压目录是否存在，不存在就解压
	if _, err := os.Stat(targetV); err != nil {
		// 解压安装包
		if err = archiver.Unarchive(filename, targetV, true); err != nil {
			return cli.Exit(errstring(err), 1)
		}
	}
	/* 	// 解压安装包
	   	if err = archiver.Unarchive(filename, versionsDir); err != nil {
	   		fmt.Println(err.Error())
	   		return cli.Exit(errstring(err), 1)
	   	}
	   	// 目录重命名
	   	if err = os.Rename(filepath.Join(versionsDir, ui.Expected), targetV); err != nil {
	   		fmt.Println(err.Error())
	   		return cli.Exit(errstring(err), 1)
	   	} */

	// 重新建立软链接
	_ = os.Remove(goroot)
	if err = mkSymlink(targetV, goroot); err != nil {
		return cli.Exit(errstring(err), 1)
	}
	fmt.Printf("Now using %s\n", ui.Expected)
	return nil
}

func install(ctx *cli.Context) (err error) {
	version := ctx.Args().First()
	if version == "" {
		return cli.ShowSubcommandHelp(ctx)
	}
	return downloadAndInstall(version)

}

func Install(version string) (err error) {
	return downloadAndInstall(version)
}

// DownloadWithProgress 下载版本另存为指定文件且显示下载进度
func DownloadWithProgress(url, dst string) (size int64, err error) {
	return download.Download(url, dst, os.O_CREATE|os.O_WRONLY, 0644, true)
}
