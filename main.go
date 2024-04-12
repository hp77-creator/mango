package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	_help = flag.String("help", "", "Usage: mango <subcommand>")
)

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Please specify a subcommand: list or install")
	}

	fmt.Println("Trying tool for the first time")
}
