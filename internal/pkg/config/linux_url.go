package config

var linux_Url_Items = []*UrlItem{
	{
		In: &JavaFileItem{
			FileName: "openjdk-19+36_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk19/ri/openjdk-19+36_linux-x64_bin.tar.gz",
			Sha256:   "f47aba585cfc9ecff1ed8e023524e8309f4315ed8b80100b40c7dcc232c12f96",
		},
		Expected:   "jdk-19",
		SimpleName: "19",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-18+36_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk18/ri/openjdk-18+36_linux-x64_bin.tar.gz",
			Sha256:   "0f60aef7b8504983d6e374fe94d09a7bedcf05ec559e812d801a33bd4ebd23d0",
		},
		Expected:   "jdk-18",
		SimpleName: "18",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-17+35_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk17/ri/openjdk-17+35_linux-x64_bin.tar.gz",
			Sha256:   "aef49cc7aa606de2044302e757fa94c8e144818e93487081c4fd319ca858134b",
		},
		Expected:   "jdk-17",
		SimpleName: "17",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-16+36_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk16/ri/openjdk-16+36_linux-x64_bin.tar.gz",
			Sha256:   "e952958f16797ad7dc7cd8b724edd69ec7e0e0434537d80d6b5165193e33b931",
		},
		Expected:   "jdk-16",
		SimpleName: "16",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-15+36_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk15/ri/openjdk-15+36_linux-x64_bin.tar.gz",
			Sha256:   "bb67cadee687d7b486583d03c9850342afea4593be4f436044d785fba9508fb7",
		},
		Expected:   "jdk-15",
		SimpleName: "15",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-14+36_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk14/ri/openjdk-14+36_linux-x64_bin.tar.gz",
			Sha256:   "c7006154dfb8b66328c6475447a396feb0042608ee07a96956547f574a911c09",
		},
		Expected:   "jdk-14",
		SimpleName: "14",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-13+33_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk13/ri/openjdk-13+33_linux-x64_bin.tar.gz",
			Sha256:   "5f547b8f0ffa7da517223f6f929a5055d749776b1878ccedbd6cc1334f4d6f4d",
		},
		Expected:   "jdk-13",
		SimpleName: "13",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-12+32_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk12/ri/openjdk-12+32_linux-x64_bin.tar.gz",
			Sha256:   "4bf3d9f4bbbb8cb9a0d96ceade42df8b2ca85f7853fbcd08274df2b7d2cef074",
		},
		Expected:   "jdk-12",
		SimpleName: "12",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-11+28_linux-x64_bin.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk11/ri/openjdk-11+28_linux-x64_bin.tar.gz",
			Sha256:   "3784cfc4670f0d4c5482604c7c513beb1a92b005f569df9bf100e8bef6610f2e",
		},
		Expected:   "jdk-11",
		SimpleName: "11",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-10+44_linux-x64_bin_ri.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk10/ri/openjdk-10+44_linux-x64_bin_ri.tar.gz",
			Sha256:   "0b0e778f2c935dae32c71dd78f3ad921e0972ffc536e64703f0e9855e570abe2",
		},
		Expected:   "jdk-10",
		SimpleName: "10",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-9+181_linux-x64_ri.zip",
			URL:      "https://download.java.net/openjdk/jdk9/ri/openjdk-9+181_linux-x64_ri.zip",
			Sha256:   "f4f9c75f9f1f2d1888b9be013c96210e017fd22547366c8e0929085dfe8e38aa",
		},
		Expected:   "jdk-9",
		SimpleName: "9",
	},
	{
		In: &JavaFileItem{
			FileName: "openjdk-8u42-b03-linux-x64-14_jul_2022.tar.gz",
			URL:      "https://download.java.net/openjdk/jdk8u42/ri/openjdk-8u42-b03-linux-x64-14_jul_2022.tar.gz",
			Sha256:   "dd5fc6ef5ebffb88cd66af5258226c31f6c719fdcd855d95464fdb2cab051baa",
		},
		Expected:   "jdk-8",
		SimpleName: "8",
	},
}
