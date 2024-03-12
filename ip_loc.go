package main

import (
	"github.com/justseemore/ip_loc/internal/db"
	"github.com/justseemore/ip_loc/pkg/dbif"
	"github.com/justseemore/ip_loc/pkg/entity"
)

func Parse(line string) entity.Entities {
	return entity.ParseLine(line)
}

func InitDb() {
	db.GetDB(dbif.TypeIPv4)
	db.GetDB(dbif.TypeIPv6)
}
