package main

import (
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
	"github.com/udzura/packer-fork-e59f09e/provisioner/file"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(file.Provisioner))
	server.Serve()
}
