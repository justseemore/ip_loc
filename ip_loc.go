package ip_loc

import (
	"ip_loc/internal/db"
	"ip_loc/pkg/dbif"
	"ip_loc/pkg/entity"
)

func Parse(line string) entity.Entities {
	return entity.ParseLine(line)
}

func InitDb() {
	db.GetDB(dbif.TypeIPv4)
	db.GetDB(dbif.TypeIPv6)
}
