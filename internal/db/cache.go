package db

import (
	"sync"

	"ip_loc/pkg/dbif"
)

var (
	dbNameCache = make(map[string]dbif.DB)
	dbTypeCache = make(map[dbif.QueryType]dbif.DB)
	queryCache  = sync.Map{}
)

var (
	NameDBMap = make(NameMap)
	TypeDBMap = make(TypeMap)
)
