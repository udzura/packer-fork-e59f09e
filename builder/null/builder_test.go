package null

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func TestBuilder_implBuilder(t *testing.T) {
	var _ packer.Builder = new(Builder)
}
