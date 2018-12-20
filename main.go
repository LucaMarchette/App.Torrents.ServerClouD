package main

import (
	"log"

	"github.com/servercd/App.Torrents.ServerClouD/server"
	"github.com/servercd/App.Torrents.ServerClouD/opts"
)

var VERSION = "0.0.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Torrents.ServerClouD",
		Port:       5000,
		ConfigPath: "torrents-svrcd.json",
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
