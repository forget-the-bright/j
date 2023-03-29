package collector

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/forget-the-bright/j/internal/pkg/config"
)

func ConvertCollectorToUrlItem(colls []*Collector, isGetSha256 bool) []*config.UrlItem {
	var rs = make([]*config.UrlItem, 0)
	for _, coll := range colls {
		var item *Op_Item
		switch runtime.GOOS {
		case "linux":
			if runtime.GOARCH == "aarch64" {
				item = coll.Linux_AArch64
			} else {
				item = coll.Linux_X64
			}
		case "windows":
			item = coll.Windows_X64
		case "darwin":
			if runtime.GOARCH == "aarch64" {
				item = coll.Mac_AArch64
			} else {
				item = coll.Mac_X64
			}
		default:
			item = nil
		}
		if item != nil {
			rs = append(rs, &config.UrlItem{
				In: &config.JavaFileItem{
					FileName: item.FileName,
					URL:      item.Url,
					Sha256:   GetSha256ByUrl(item.Sha256Url, isGetSha256),
				},
				SimpleName: coll.Version,
				Expected:   getFileNameNoSuffix(item.FileName),
			})
		}
	}
	switch runtime.GOOS {
	case "windows":
		rs = append(rs, &config.UrlItem{
			In: &config.JavaFileItem{
				FileName: "openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
				URL:      "https://download.java.net/openjdk/jdk8u42/ri/openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
				Sha256:   "0314134bd981db63c7ca68d262ef896383b5694307a14bac81af88b5ad926279",
			},
			Expected:   "openjdk-8u42-b03-windows-i586-14_jul_2022",
			SimpleName: "8",
		})
	case "linux":
		rs = append(rs, &config.UrlItem{
			In: &config.JavaFileItem{
				FileName: "openjdk-8u42-b03-linux-x64-14_jul_2022.tar.gz",
				URL:      "https://download.java.net/openjdk/jdk8u42/ri/openjdk-8u42-b03-linux-x64-14_jul_2022.tar.gz",
				Sha256:   "dd5fc6ef5ebffb88cd66af5258226c31f6c719fdcd855d95464fdb2cab051baa",
			},
			Expected:   "openjdk-8u42-b03-linux-x64-14_jul_2022",
			SimpleName: "8",
		})

	}
	return config.ReverseArray(rs)
}

func GetOpenJDKVesionUrlInfo() []*Collector {
	resp, _ := http.Get(Collector_Archive_Url)
	Collectors = make([]*Collector, 0)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	divs := doc_selector.Find("#sidebar").Find(".links")
	divs.Each(func(j int, div *goquery.Selection) {
		link_about := div.Find(".about").Text()
		if link_about == "Reference Implementations" {
			a_docs := div.Find("a")
			a_docs.Each(func(j int, a_doc *goquery.Selection) {
				info_url := a_doc.AttrOr("href", "")
				versions := strings.Split(a_doc.Text(), " ")
				version := versions[len(versions)-1]
				version_url := Collector_Url + strings.ReplaceAll(info_url, ".", "")
				Collectors = append(Collectors, getVersionByUrl(version, version_url))
				//fmt.Println(version)
				//fmt.Println(version_url)
			})
		}
		//fmt.Printf("link_about: %v\n", link_about)
	})
	return Collectors
}

func GetOpenJDKArchiveReleasesInfo() []*Collector {
	Archive_Releases_Collectors = make([]*Collector, 0)
	resp, _ := http.Get(Collector_Archive_Url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	docs, _ := goquery.NewDocumentFromReader(resp.Body)
	tbody_docs := docs.Find(".builds").Find("tbody").Children()
	var collector_Item *Collector
	tbody_docs.Each(func(j int, tr *goquery.Selection) {
		th := tr.Find("th")
		if th.Length() == 1 {
			val := th.Text()
			if val != "Source" {
				s2 := strings.Split(val, "(build ")
				version_str := strings.Trim(strings.Trim(s2[0], " "), " GA")
				//fmt.Println("\n" + version_str)
				collector_Item = &Collector{
					Version: version_str,
				}
				Archive_Releases_Collectors = append(Archive_Releases_Collectors, collector_Item)
			}
		}
		if th.Length() == 2 {
			a_item := tr.Find("td").Find("a")
			file_type := strings.Trim(a_item.First().Text(), "\n")
			download_url := a_item.First().AttrOr("href", "")
			file_name := getFileNameByDownLoadUrl(download_url)
			sha256_url := a_item.Last().AttrOr("href", "")
			releases := strings.Split(th.First().Text(), "/")

			switch releases[0] {
			case "Windows":
				collector_Item.Windows_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
			case "Mac":
				if len(releases) == 1 || releases[1] == "x64" {
					collector_Item.Mac_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
				} else {
					collector_Item.Mac_AArch64 = build_Op_Item(file_type, releases[1], download_url, sha256_url, file_name)
				}
			case "Linux":
				if len(releases) == 1 || releases[1] == "x64" {
					collector_Item.Linux_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
				} else {
					collector_Item.Linux_X64 = build_Op_Item(file_type, releases[1], download_url, sha256_url, file_name)
				}
			}
		}
	})
	return Archive_Releases_Collectors
}

func getVersionByUrl(version, url string) *Collector {
	var coll = Collector{
		Version: version,
	}
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)

	divs := doc_selector.Find("#main").Find("ul")

	if divs.Length() <= 1 || version == "8" {
		li_docs := divs.Find("li")
		linux_x64_a := li_docs.Eq(0).Find("a")
		windows_x64_a := li_docs.Eq(1).Find("a")

		linux_url := linux_x64_a.Eq(0).AttrOr("href", "")
		linux_file_name := getFileNameByDownLoadUrl(linux_url)
		linux_file_type := getFileTypeByFileName(linux_file_name)
		linux_sha256_url := linux_x64_a.Eq(1).AttrOr("href", "")

		windows_url := windows_x64_a.Eq(0).AttrOr("href", "")
		windows_file_name := getFileNameByDownLoadUrl(windows_url)
		windows_file_type := getFileTypeByFileName(windows_file_name)
		windows_sha256_url := windows_x64_a.Eq(1).AttrOr("href", "")

		coll.Linux_X64 = &Op_Item{
			Arch:      "x64",
			Url:       linux_url,
			Sha256Url: linux_sha256_url,
			FileName:  linux_file_name,
			FileType:  linux_file_type,
		}
		coll.Windows_X64 = &Op_Item{
			Arch:      "x64",
			Url:       windows_url,
			Sha256Url: windows_sha256_url,
			FileName:  windows_file_name,
			FileType:  windows_file_type,
		}
	} else {
		divs_eq0 := divs.Eq(0)
		li_docs := divs_eq0.Find("li")
		linux_x64_a := li_docs.Eq(0).Find("a")

		linux_url := linux_x64_a.Eq(0).AttrOr("href", "")
		linux_file_name := getFileNameByDownLoadUrl(linux_url)
		linux_file_type := getFileTypeByFileName(linux_file_name)
		linux_sha256_url := linux_x64_a.Eq(1).AttrOr("href", "")
		coll.Linux_X64 = &Op_Item{
			Arch:      "x64",
			Url:       linux_url,
			Sha256Url: linux_sha256_url,
			FileName:  linux_file_name,
			FileType:  linux_file_type,
		}

	}
	return &coll
}
