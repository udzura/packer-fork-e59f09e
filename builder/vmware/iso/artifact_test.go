package iso

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestArtifact_Impl(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatal("Artifact must be a proper artifact")
	}
}
