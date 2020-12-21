package main

import (
	"ticket/api"
	log "github.com/jeanphorn/log4go"
)

func main() {
	log.LoadConfiguration("./log.json")
	api.Run()
}
