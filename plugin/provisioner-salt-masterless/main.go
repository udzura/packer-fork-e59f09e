package main

import (
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
	"github.com/udzura/packer-fork-e59f09e/provisioner/salt-masterless"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(saltmasterless.Provisioner))
	server.Serve()
}
