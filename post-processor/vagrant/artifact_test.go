package vagrant

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestArtifact_ImplementsArtifact(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatalf("Artifact should be a Artifact")
	}
}

func TestArtifact_Id(t *testing.T) {
	artifact := NewArtifact("vmware", "./")
	if artifact.Id() != "vmware" {
		t.Fatalf("should return name as Id")
	}
}
