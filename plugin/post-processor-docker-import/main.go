package main

import (
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
	"github.com/udzura/packer-fork-e59f09e/post-processor/docker-import"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockerimport.PostProcessor))
	server.Serve()
}
