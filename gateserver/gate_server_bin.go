package main

import (
	"./src"
	"flag"
)

func main() {
	flag.Parse()
	gateserver.Start()
}