package main

import (
	"fmt"
	"github.com/hypebeast/go-config/example/config"
)

func main() {
	fmt.Println("####")
	fmt.Println("## This example program shows how to use go-config")
	fmt.Println("## You can set GOENV to ci, stage or prod to load different options accordingly to the set environment")
	fmt.Println("## Example: go build; GOENV=stage ./example")
	fmt.Println("####\n")

	fmt.Println("## First, get the basic config and read the options:")
	baseConfig := config.BaseConf()
	fmt.Printf("%#v\n", baseConfig)

	fmt.Println("\n## Get the MongoDB options:")
	mongoConfig := config.MongoConf()
	fmt.Printf("%#v\n", mongoConfig)
}
