package main

import (
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
	"github.com/udzura/packer-fork-e59f09e/post-processor/vagrant"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(vagrant.PostProcessor))
	server.Serve()
}
