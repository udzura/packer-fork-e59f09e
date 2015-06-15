package chroot

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestCommunicator_ImplementsCommunicator(t *testing.T) {
	var raw interface{}
	raw = &Communicator{}
	if _, ok := raw.(packer.Communicator); !ok {
		t.Fatalf("Communicator should be a communicator")
	}
}
