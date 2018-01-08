package main

import (
	"./src"
	"../utils"
	"flag"
)

func main() {
	flag.Parse()
	go utils.SigalLister(dataserver.Stop)
	dataserver.Start()
}