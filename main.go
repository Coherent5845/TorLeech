package main

import (
	"log"

	"github.com/ipadtechspot/iTS-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.0.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "iPad TechSpot Torrent Leecher",
		Port:       3000,
		ConfigPath: "iTS-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
