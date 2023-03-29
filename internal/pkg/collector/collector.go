package collector

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Collector struct {
	Version       string
	Windows_X64   *Op_Item
	Linux_X64     *Op_Item
	Linux_AArch64 *Op_Item
	Mac_X64       *Op_Item
	Mac_AArch64   *Op_Item
}

type Op_Item struct {
	FileType  string
	Arch      string
	Url       string
	Sha256Url string
	FileName  string
}

var Archive_Releases_Collectors []*Collector
var Collectors []*Collector

var Collector_Archive_Url string = Collector_Url + "/archive/"
var Collector_Url string = "https://jdk.java.net"

func build_Op_Item(file_type, arch, download_url, sha256_url, file_name string) *Op_Item {
	return &Op_Item{
		FileType:  file_type,
		Arch:      arch,
		Url:       download_url,
		Sha256Url: sha256_url,
		FileName:  file_name,
	}
}

func getFileNameByDownLoadUrl(url string) string {
	downloads := strings.Split(url, "/")
	file_name := downloads[len(downloads)-1]
	return file_name
}
func getFileNameNoSuffix(file_name string) string {
	return strings.ReplaceAll(file_name, "."+getFileTypeByFileName(file_name), "")
}

func GetSha256ByUrl(url string, isGetSha256 bool) string {
	if isGetSha256 {
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		return string(bytes)
	} else {
		return url
	}
}

func getFileTypeByFileName(filename string) string {
	filenames := strings.Split(filename, ".")
	switch filenames[len(filenames)-1] {
	case "zip":
		return "zip"
	case "gz":
		return "tar.gz"
	default:
		return filenames[len(filenames)-1]
	}
}
