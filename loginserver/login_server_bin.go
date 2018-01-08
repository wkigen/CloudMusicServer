package main

import (
	"./src"
	"../utils"
	"flag"
)

func main() {
	flag.Parse()
	go utils.SigalLister(loginserver.Stop)
	loginserver.Start()
}