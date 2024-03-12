package qqwry

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"ip_loc/pkg/wry"
)

var DownloadUrls = []string{
	"https://gh-release.zu1k.com/HMBSbige/qqwry/qqwry.dat", // redirect to HMBSbige/qqwry
	// Other repo:
	// https://github.com/HMBSbige/qqwry
	// https://github.com/metowolf/qqwry.dat
}

type QQwry struct {
	wry.IPDB[uint32]
}

// NewQQwry new database from path
func NewQQwry(filePath string) (*QQwry, error) {
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
	header := fileData[0:8]
	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])

	return &QQwry{
		IPDB: wry.IPDB[uint32]{
			Data: fileData,

			OffLen:   3,
			IPLen:    4,
			IPCnt:    (end-start)/7 + 1,
			IdxStart: start,
			IdxEnd:   end,
		},
	}, nil
}

func (db QQwry) Find(query string, params ...string) (result fmt.Stringer, err error) {
	ip := net.ParseIP(query)
	if ip == nil {
		return nil, errors.New("query should be IPv4")
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return nil, errors.New("query should be IPv4")
	}
	ip4uint := binary.BigEndian.Uint32(ip4)

	offset := db.SearchIndexV4(ip4uint)
	if offset <= 0 {
		return nil, errors.New("query not valid")
	}

	reader := wry.NewReader(db.Data)
	reader.Parse(offset + 4)
	return reader.Result.DecodeGBK(), nil
}

func (db QQwry) Name() string {
	return "qqwry"
}
