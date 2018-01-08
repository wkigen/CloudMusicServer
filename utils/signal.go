package utils

import (
	"fmt"
	"os"
	"os/signal"
)

func SigalLister(f func())  {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("get signal:", s)
	f()
    os.Exit(1) 
}