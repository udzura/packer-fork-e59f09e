package null

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestNullArtifact(t *testing.T) {
	var _ packer.Artifact = new(NullArtifact)
}
