package main

import (
	"flag"
	"fmt"

	p2rm "github.com/timmyw/pocket2rm"
)

func main() {

	command := flag.String("command", "pull", "Command to execute")
	flag.Parse()
	fmt.Printf("%s\n", *command)

	var p *p2rm.Pocket2RM = new(p2rm.Pocket2RM)
	p.Init()

	switch *command {
		case "pull":
		case "auth":
		p.Authorise()
	}
}
