package main

import (
	"flag"
	"fmt"
)

func main() {
	command := flag.String("command", "pull", "Command to execute")
	flag.Parse()
	fmt.Printf("%s\n", *command)
}
