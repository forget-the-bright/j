package config

type JavaFileItem struct {
	FileName string
	URL      string
}
type UrlItem struct {
	In         *JavaFileItem
	Expected   string
	SimpleName string
}

var Url_Items = []*UrlItem{
	{
		In: &JavaFileItem{
			FileName: "openjdk-19+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk19/ri/openjdk-19+36_windows-x64_bin.zip",
		},
		Expected:   "jdk-19",
		SimpleName: "19",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-18+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk18/ri/openjdk-18+36_windows-x64_bin.zip",
		},
		Expected:   "jdk-18",
		SimpleName: "18",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-17+35_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk17/ri/openjdk-17+35_windows-x64_bin.zip",
		},
		Expected:   "jdk-17",
		SimpleName: "17",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-16+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk16/ri/openjdk-16+36_windows-x64_bin.zip",
		},
		Expected:   "jdk-16",
		SimpleName: "16",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-15+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk15/ri/openjdk-15+36_windows-x64_bin.zip",
		},
		Expected:   "jdk-15",
		SimpleName: "15",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-14+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk14/ri/openjdk-14+36_windows-x64_bin.zip",
		},
		Expected:   "jdk-14",
		SimpleName: "14",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-13+33_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk13/ri/openjdk-13+33_windows-x64_bin.zip",
		},
		Expected:   "jdk-13",
		SimpleName: "13",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-12+32_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk12/ri/openjdk-12+32_windows-x64_bin.zip",
		},
		Expected:   "jdk-12",
		SimpleName: "12",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-11+28_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk11/ri/openjdk-11+28_windows-x64_bin.zip",
		},
		Expected:   "jdk-11",
		SimpleName: "11",
	},
	{
		In: &JavaFileItem{
			FileName: "jdk-9+181_windows-x64_ri.zip",
			URL:      "https://download.java.net/openjdk/jdk9/ri/jdk-9+181_windows-x64_ri.zip",
		},
		Expected:   "jdk-9",
		SimpleName: "9",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
			URL:      "https://download.java.net/openjdk/jdk8u42/ri/openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
		},
		Expected:   "jdk-8",
		SimpleName: "8",
	},
}
