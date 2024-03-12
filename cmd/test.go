package main

import (
	"github.com/justseemore/ip_loc"
	"log"
)

func main() {
	log.Println(ip_loc.Parse("117.118.0.94"))
}
