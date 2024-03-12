package main

import (
	"github.com/justseemore/ip_loc"
	"log"
)

func main() {
	log.Println(ip_loc.Parse("240e:441:9a1a:386f:24ed:84ff:fea7:d28a"))
}
