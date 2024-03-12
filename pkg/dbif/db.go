package dbif

import (
	"fmt"

	"github.com/justseemore/ip_loc/pkg/cdn"
	"github.com/justseemore/ip_loc/pkg/geoip"
	"github.com/justseemore/ip_loc/pkg/ip2location"
	"github.com/justseemore/ip_loc/pkg/ip2region"
	"github.com/justseemore/ip_loc/pkg/ipip"
	"github.com/justseemore/ip_loc/pkg/qqwry"
	"github.com/justseemore/ip_loc/pkg/zxipv6wry"
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
