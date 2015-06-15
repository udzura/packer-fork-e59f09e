package main

import (
	"github.com/udzura/packer-fork-e59f09e/builder/vmware/vmx"
	"github.com/udzura/packer-fork-e59f09e/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(vmx.Builder))
	server.Serve()
}
