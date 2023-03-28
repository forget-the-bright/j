package collector

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type Collector struct {
	version       string
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

var Collectors []*Collector

func Test_getGoVersion(t *testing.T) {
	Collectors = make([]*Collector, 0)
	resp, _ := http.Get("https://jdk.java.net/archive/")
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
						version: version_str,
					}
					Collectors = append(Collectors, collector_Item)
				}
			}
			if th.Length() == 2 {
				a_item := tr.Find("td").Find("a")
				file_type := strings.Trim(a_item.First().Text(), "\n")
				download_url := a_item.First().AttrOr("href", "")
				downloads := strings.Split(download_url, "/")
				file_name := downloads[len(downloads)-1]
				sha256_url := a_item.Last().AttrOr("href", "")
				releases := strings.Split(th.First().Text(), "/")

				switch releases[0] {
				case "Windows":
					collector_Item.Mac_X64 = Make_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
				case "Mac":
					if len(releases) == 1 || releases[1] == "x64" {
						collector_Item.Mac_X64 = Make_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					} else {
						collector_Item.Mac_AArch64 = Make_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					}
				case "Linux":
					if len(releases) == 1 || releases[1] == "x64" {
						collector_Item.Linux_X64 = Make_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
					} else {
						collector_Item.Linux_X64 = Make_Op_Item(file_type, "x64", download_url, sha256_url, file_name)
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

func Make_Op_Item(file_type, arch, download_url, sha256_url, file_name string) *Op_Item {
	return &Op_Item{
		FileType:  file_type,
		Arch:      "x64",
		Url:       download_url,
		Sha256Url: sha256_url,
		FileName:  file_name,
	}
}
