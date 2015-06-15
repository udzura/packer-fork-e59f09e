package docker

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestCommunicator_impl(t *testing.T) {
	var _ packer.Communicator = new(Communicator)
}
