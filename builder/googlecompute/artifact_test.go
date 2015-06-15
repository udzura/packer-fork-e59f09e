package googlecompute

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestArtifact_impl(t *testing.T) {
	var _ packer.Artifact = new(Artifact)
}
