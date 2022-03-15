package main

import (
	"gitlab.com/trustify/core/config"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})
}
