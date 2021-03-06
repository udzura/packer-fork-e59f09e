package common

import (
	"bytes"
	"github.com/mitchellh/multistep"
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
