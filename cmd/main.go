package main

import (
	"log"

	"parallel.downloader.com/entity"
)

func main() {
	dcp := new(entity.DownloadCmdParam)
	err := dcp.ExecCmd()
	if err != nil {
		return
	}
	if dcp == nil {
		log.Fatal("dcp is nil")
		return
	}
	log.Printf("%v", dcp)
	d := entity.Downloader{
		DownloadCmdParam: *dcp,
	}
	log.Printf("%v", d)
}
