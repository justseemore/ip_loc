package dbif

import (
	"fmt"

	"ip_loc/pkg/cdn"
	"ip_loc/pkg/geoip"
	"ip_loc/pkg/ip2location"
	"ip_loc/pkg/ip2region"
	"ip_loc/pkg/ipip"
	"ip_loc/pkg/qqwry"
	"ip_loc/pkg/zxipv6wry"
)

type QueryType uint

const (
	TypeIPv4 = iota
	TypeIPv6
	TypeDomain
)

type DB interface {
	Find(query string, params ...string) (result fmt.Stringer, err error)
	Name() string
}

var (
	_ DB = &qqwry.QQwry{}
	_ DB = &zxipv6wry.ZXwry{}
	_ DB = &ipip.IPIPFree{}
	_ DB = &geoip.GeoIP{}
	_ DB = &ip2region.Ip2Region{}
	_ DB = &ip2location.IP2Location{}
	_ DB = &cdn.CDN{}
)
