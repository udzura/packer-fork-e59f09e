package iso

import (
	"testing"

	vmwcommon "github.com/udzura/packer-fork-e59f09e/builder/vmware/common"
)

func TestRemoteDriverMock_impl(t *testing.T) {
	var _ vmwcommon.Driver = new(RemoteDriverMock)
	var _ RemoteDriver = new(RemoteDriverMock)
}
