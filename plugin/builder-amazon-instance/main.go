package main

import (
	"github.com/udzura/packer-fork-e59f09e/builder/amazon/instance"
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(instance.Builder))
	server.Serve()
}
