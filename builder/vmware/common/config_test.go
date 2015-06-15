package common

import (
	"github.com/udzura/packer-fork-e59f09e/packer"
	"testing"
)

func testConfigTemplate(t *testing.T) *packer.ConfigTemplate {
	result, err := packer.NewConfigTemplate()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return result
}
