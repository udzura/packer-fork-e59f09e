package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/udzura/packer-fork-e59f09e/packer"
	"strings"
)

type commandTemplate struct {
	Name string
}

// This step executes additional VBoxManage commands as specified by the
// template.
//
// Uses:
//   driver Driver
//   ui packer.Ui
//   vmName string
//
// Produces:
type StepVBoxManage struct {
	Commands [][]string
	Tpl      *packer.ConfigTemplate
}

func (s *StepVBoxManage) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	vmName := state.Get("vmName").(string)

	if len(s.Commands) > 0 {
		ui.Say("Executing custom VBoxManage commands...")
	}

	tplData := &commandTemplate{
		Name: vmName,
	}

	for _, originalCommand := range s.Commands {
		command := make([]string, len(originalCommand))
		copy(command, originalCommand)

		for i, arg := range command {
			var err error
			command[i], err = s.Tpl.Process(arg, tplData)
			if err != nil {
				err := fmt.Errorf("Error preparing vboxmanage command: %s", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

		ui.Message(fmt.Sprintf("Executing: %s", strings.Join(command, " ")))
		if err := driver.VBoxManage(command...); err != nil {
			err := fmt.Errorf("Error executing command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *StepVBoxManage) Cleanup(state multistep.StateBag) {}
