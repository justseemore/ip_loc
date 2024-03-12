package db

import (
	"github.com/justseemore/ip_loc/pkg/qqwry"
	"os"
)

func getCwd() string {
	dir, _ := os.Getwd()
	return dir
}

func GetDefaultDBList() List {
	os.Getwd()
	return List{
		&DB{
			Name: "qqwry",
			NameAlias: []string{
				"chunzhen",
			},
			Format:       FormatQQWry,
			File:         getCwd() + "/data/qqwry.dat",
			Languages:    LanguagesZH,
			Types:        TypesIPv4,
			DownloadUrls: qqwry.DownloadUrls,
		},
		&DB{
			Name: "zxipv6wry",
			NameAlias: []string{
				"zxipv6",
				"zx",
			},
			Format:    FormatZXIPv6Wry,
			File:      getCwd() + "/data/zxipv6wry.db",
			Languages: LanguagesZH,
			Types:     TypesIPv6,
		},
		&DB{
			Name:      "cdn",
			Format:    FormatCDNYml,
			File:      getCwd() + "/data/cdn.yml",
			Languages: LanguagesZH,
			Types:     TypesCDN,
			//DownloadUrls: cdn.DownloadUrls,
		},
	}
}
