package main

import (
	"github.com/udzura/packer-fork-e59f09e/builder/digitalocean"
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(digitalocean.Builder))
	server.Serve()
}
