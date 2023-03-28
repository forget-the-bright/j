package config

var windows_Url_Items = []*UrlItem{
	{
		In: &JavaFileItem{
			FileName: "openjdk-19+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk19/ri/openjdk-19+36_windows-x64_bin.zip",
			Sha256:   "8fabcee7c4e8d3b53486777ecd27bb906d67d7c1efd1bf22a8290cf659afa487",
		},
		Expected:   "jdk-19",
		SimpleName: "19",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-18+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk18/ri/openjdk-18+36_windows-x64_bin.zip",
			Sha256:   "a5b91d4c12752d44aa75df70ae3e2311287b3e60c288b07dade106376c688277",
		},
		Expected:   "jdk-18",
		SimpleName: "18",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-17+35_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk17/ri/openjdk-17+35_windows-x64_bin.zip",
			Sha256:   "e88b0df00021c9d266bb435c9a95fdc67d1948cce4518daf85c234907bd393c5",
		},
		Expected:   "jdk-17",
		SimpleName: "17",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-16+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk16/ri/openjdk-16+36_windows-x64_bin.zip",
			Sha256:   "a78bdeaad186297601edac6772d931224d7af6f682a43372e693c37020bd37d6",
		},
		Expected:   "jdk-16",
		SimpleName: "16",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-15+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk15/ri/openjdk-15+36_windows-x64_bin.zip",
			Sha256:   "764e39a71252a9791118a31ae56a4247c049463bda5eb72497122ec50b1d07f8",
		},
		Expected:   "jdk-15",
		SimpleName: "15",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-14+36_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk14/ri/openjdk-14+36_windows-x64_bin.zip",
			Sha256:   "6b56c65c2ebb89eb361f47370359f88c4b87234dc073988a2c33e7d75c01e488",
		},
		Expected:   "jdk-14",
		SimpleName: "14",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-13+33_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk13/ri/openjdk-13+33_windows-x64_bin.zip",
			Sha256:   "053d8c87bb34347478512911a6218a389720bffcde4e496be5a54d51ad7c9c2f",
		},
		Expected:   "jdk-13",
		SimpleName: "13",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-12+32_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk12/ri/openjdk-12+32_windows-x64_bin.zip",
			Sha256:   "d6a550477754289e5bc0a635974b40bf5bc0515db441381414303ae954d8d6b8",
		},
		Expected:   "jdk-12",
		SimpleName: "12",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-11+28_windows-x64_bin.zip",
			URL:      "https://download.java.net/openjdk/jdk11/ri/openjdk-11+28_windows-x64_bin.zip",
			Sha256:   "fde3b28ca31b86a889c37528f17411cd0b9651beb6fa76cac89a223417910f4b",
		},
		Expected:   "jdk-11",
		SimpleName: "11",
	},
	{
		In: &JavaFileItem{
			FileName: "jdk-9+181_windows-x64_ri.zip",
			URL:      "https://download.java.net/openjdk/jdk9/ri/jdk-9+181_windows-x64_ri.zip",
			Sha256:   "51948d69c7b770b376162ec5b88f6ec8a266bd3c9e6da21c4e834b6d0d661897",
		},
		Expected:   "jdk-9",
		SimpleName: "9",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
			URL:      "https://download.java.net/openjdk/jdk8u42/ri/openjdk-8u42-b03-windows-i586-14_jul_2022.zip",
			Sha256:   "0314134bd981db63c7ca68d262ef896383b5694307a14bac81af88b5ad926279",
		},
		Expected:   "java-se-8u42-ri",
		SimpleName: "8",
	},
}
