package main

import (
	"dstuhack/internal/app"

	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "", "config path")
}

func main() {
	flag.Parse()

	application := app.NewApplication()
	application.ReadConfig(configPath)
	application.Run()
}
