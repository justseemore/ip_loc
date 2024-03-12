package db

import (
	"log"

	"github.com/justseemore/ip_loc/pkg/cdn"
	"github.com/justseemore/ip_loc/pkg/dbif"
	"github.com/justseemore/ip_loc/pkg/qqwry"
	"github.com/justseemore/ip_loc/pkg/zxipv6wry"
)

func GetDB(typ dbif.QueryType) (db dbif.DB) {
	if db, found := dbTypeCache[typ]; found {
		return db
	}
	var err error
	switch typ {
	case dbif.TypeIPv4:
		db, err = qqwry.NewQQwry(getDbByName("qqwry").File)
	case dbif.TypeIPv6:
		db, err = zxipv6wry.NewZXwry(getDbByName("zxipv6wry").File)
	case dbif.TypeDomain:
		db, err = cdn.NewCDN(getDbByName("cdn").File)
	default:
		panic("Query type not supported!")
	}

	if err != nil || db == nil {
		log.Fatalln("Database init failed:", err)
	}

	dbTypeCache[typ] = db
	return
}

func Find(typ dbif.QueryType, query string) *Result {
	db := GetDB(typ)
	result, err := db.Find(query)
	if err != nil {
		return nil
	}
	res := &Result{db.Name(), result}
	return res
}
