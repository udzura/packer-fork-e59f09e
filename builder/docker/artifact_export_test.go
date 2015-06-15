package docker

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestExportArtifact_impl(t *testing.T) {
	var _ packer.Artifact = new(ExportArtifact)
}
