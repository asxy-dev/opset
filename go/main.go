package main

import (
	"flag"
	"fmt"
)

func main() {
	cfg := flag.String("config", "data/servers.json", "path to servers config")
	flag.Parse()
	fmt.Println("OpSet CLI - starter placeholder")
	fmt.Println("Using config:", *cfg)
}
