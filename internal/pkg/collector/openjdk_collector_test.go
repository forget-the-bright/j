package collector

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_getOpenJdkVersion(t *testing.T) {
	resp, _ := http.Get(Collector_Archive_Url)
	colls := make([]*Collector, 0)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	t.Run("", func(t *testing.T) {
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
					colls = append(colls, test_getVersionByUrl(version, version_url))
					//fmt.Println(version)
					//fmt.Println(version_url)
				})
			}
			fmt.Printf("link_about: %v\n", link_about)
		})
		temp := colls
		fmt.Printf("len(temp): %v\n", len(temp))
		//fmt.Printf("divs.Length(): %v\n", divs.Length())
	})
}

func test_getVersionByUrl(version, url string) *Collector {
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

func Test_getArchiveVersion(t *testing.T) {
	Collectors = make([]*Collector, 0)
	resp, _ := http.Get(Collector_Archive_Url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	c, _ := goquery.NewDocumentFromReader(resp.Body)
	t.Run("", func(t *testing.T) {
		s := c.Find(".builds").Find("tbody").Children()
		/* 		s.Each(func(j int, tr *goquery.Selection) {
			th := tr.Find("th")
			if th.Length() == 1 {
				val := th.Text()
				if val != "Source" {
					s2 := strings.Split(val, "(build ")
					fmt.Println(strings.Trim(strings.Trim(s2[0], " "), " GA"))
				}
			}
		}) */
		var collector_Item *Collector
		//flag := false
		s.Each(func(j int, tr *goquery.Selection) {
			th := tr.Find("th")
			if th.Length() == 1 {
				val := th.Text()
				if val != "Source" {
					s2 := strings.Split(val, "(build ")
					version_str := strings.Trim(strings.Trim(s2[0], " "), " GA")
					fmt.Println("\n" + version_str)
					collector_Item = &Collector{
						Version: version_str,
					}
					Collectors = append(Collectors, collector_Item)
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
					collector_Item.Mac_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
				case "Mac":
					if len(releases) == 1 || releases[1] == "x64" {
						collector_Item.Mac_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					} else {
						collector_Item.Mac_AArch64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					}
				case "Linux":
					if len(releases) == 1 || releases[1] == "x64" {
						collector_Item.Linux_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					} else {
						collector_Item.Linux_X64 = build_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					}
				}
				//fmt.Printf("%v %v\n", strings.Trim(a_item.First().Text(), "\n"), a_item.Last().Text())
				//fmt.Printf("%v %v\n", th.First().Text(), th.Last().Text())
				//fmt.Printf("%v\n%v\n", a_item.First().AttrOr("href", ""), a_item.Last().AttrOr("href", ""))
				//fmt.Printf("th.Text(): %v th.len:%v\n", th.Text(), thLen)
			}
		})
	})
	c2 := Collectors
	fmt.Printf("len(c2): %v\n", len(c2))
	fmt.Println(c2)
}
