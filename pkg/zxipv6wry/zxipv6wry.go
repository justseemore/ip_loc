package zxipv6wry

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/justseemore/ip_loc/pkg/wry"
)

type ZXwry struct {
	wry.IPDB[uint64]
}

func NewZXwry(filePath string) (*ZXwry, error) {
	var fileData []byte

	fileBase, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}
	defer fileBase.Close()

	fileData, err = io.ReadAll(fileBase)
	if err != nil {
		return nil, err
	}
	header := fileData[:24]
	offLen := header[6]
	ipLen := header[7]

	start := binary.LittleEndian.Uint64(header[16:24])
	counts := binary.LittleEndian.Uint64(header[8:16])
	end := start + counts*11

	return &ZXwry{
		IPDB: wry.IPDB[uint64]{
			Data: fileData,

			OffLen:   offLen,
			IPLen:    ipLen,
			IPCnt:    counts,
			IdxStart: start,
			IdxEnd:   end,
		},
	}, nil
}

func (db *ZXwry) Find(query string, _ ...string) (result fmt.Stringer, err error) {
	ip := net.ParseIP(query)
	if ip == nil {
		return nil, errors.New("query should be IPv6")
	}
	ip6 := ip.To16()
	if ip6 == nil {
		return nil, errors.New("query should be IPv6")
	}
	ip6 = ip6[:8]
	ipu64 := binary.BigEndian.Uint64(ip6)

	offset := db.SearchIndexV6(ipu64)
	reader := wry.NewReader(db.Data)
	reader.Parse(offset)
	return reader.Result, nil
}

func (db *ZXwry) Name() string {
	return "zxipv6wry"
}

func CheckFile(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	if string(data[:4]) != "IPDB" {
		return false
	}

	if len(data) < 24 {
		return false
	}
	header := data[:24]
	start := binary.LittleEndian.Uint64(header[16:24])
	counts := binary.LittleEndian.Uint64(header[8:16])
	end := start + counts*11
	if start >= end || uint64(len(data)) < end {
		return false
	}

	return true
}
