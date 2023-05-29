package main

import (
	"github.com/tugasmeilyanto/go-trial-class/cli"
	"github.com/tugasmeilyanto/go-trial-class/config"
)

func main() {
	// fmt.Println("Hello World")
	config.DBConnect()
	// cli.MainMenu()
	cli.LoginMenu()
}
