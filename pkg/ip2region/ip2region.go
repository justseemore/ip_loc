package ip2region

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/justseemore/ip_loc/pkg/wry"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var DownloadUrls = []string{
	"https://cdn.jsdelivr.net/gh/lionsoul2014/ip2region/data/ip2region.xdb",
	"https://raw.githubusercontent.com/lionsoul2014/ip2region/master/data/ip2region.xdb",
}

type Ip2Region struct {
	seacher *xdb.Searcher
}

func NewIp2Region(filePath string) (*Ip2Region, error) {
	_, err := os.Stat(filePath)

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	searcher, err := xdb.NewWithBuffer(data)
	if err != nil {
		fmt.Printf("无法解析 ip2region xdb 数据库: %s\n", err)
		return nil, err
	}
	return &Ip2Region{
		seacher: searcher,
	}, nil
}

func (db Ip2Region) Find(query string, params ...string) (result fmt.Stringer, err error) {
	if db.seacher != nil {
		res, err := db.seacher.SearchByStr(query)
		if err != nil {
			return nil, err
		} else {
			return wry.Result{
				Country: strings.ReplaceAll(res, "|0", ""),
			}, nil
		}
	}

	return nil, errors.New("ip2region 未初始化")
}

func (db Ip2Region) Name() string {
	return "ip2region"
}
