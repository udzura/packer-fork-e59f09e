package vmx

import (
	"bytes"
	"testing"

	"github.com/mitchellh/multistep"
	vmwcommon "github.com/udzura/packer-fork-e59f09e/builder/vmware/common"
	"github.com/udzura/packer-fork-e59f09e/packer"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(vmwcommon.DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
