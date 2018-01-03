package main

import (
	"./utils"
	"fmt"
)


func main(){
	config := utils.ReadConfig()
	fmt.Print(config)
}