package main

import (
	"github.com/guionardo/gs-ops/src/host"
)

func main() {
	server, logger := host.GetServer(":8080")
	host.RunServer(server, *logger)

}
