package db

import (
	"github.com/justseemore/ip_loc/pkg/dbif"
)

var (
	dbNameCache = make(map[string]dbif.DB)
	dbTypeCache = make(map[dbif.QueryType]dbif.DB)
)

var (
	NameDBMap = make(NameMap)
	TypeDBMap = make(TypeMap)
)
