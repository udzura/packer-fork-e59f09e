package chroot

import (
	"fmt"

	"github.com/mitchellh/goamz/ec2"
	"github.com/mitchellh/multistep"
	awscommon "github.com/udzura/packer-fork-e59f09e/builder/amazon/common"
	"github.com/udzura/packer-fork-e59f09e/packer"
)

// StepRegisterAMI creates the AMI.
type StepRegisterAMI struct{}

func (s *StepRegisterAMI) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ec2conn := state.Get("ec2").(*ec2.EC2)
	image := state.Get("source_image").(*ec2.Image)
	snapshotId := state.Get("snapshot_id").(string)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Registering the AMI...")
	blockDevices := make([]ec2.BlockDeviceMapping, len(image.BlockDevices))
	for i, device := range image.BlockDevices {
		newDevice := device
		if newDevice.DeviceName == image.RootDeviceName {
			newDevice.SnapshotId = snapshotId
		}

		blockDevices[i] = newDevice
	}

	registerOpts := buildRegisterOpts(config, image, blockDevices)

	// Set SriovNetSupport to "simple". See http://goo.gl/icuXh5
	if config.AMIEnhancedNetworking {
		registerOpts.SriovNetSupport = "simple"
	}

	registerResp, err := ec2conn.RegisterImage(registerOpts)
	if err != nil {
		state.Put("error", fmt.Errorf("Error registering AMI: %s", err))
		ui.Error(state.Get("error").(error).Error())
		return multistep.ActionHalt
	}

	// Set the AMI ID in the state
	ui.Say(fmt.Sprintf("AMI: %s", registerResp.ImageId))
	amis := make(map[string]string)
	amis[ec2conn.Region.Name] = registerResp.ImageId
	state.Put("amis", amis)

	// Wait for the image to become ready
	stateChange := awscommon.StateChangeConf{
		Pending:   []string{"pending"},
		Target:    "available",
		Refresh:   awscommon.AMIStateRefreshFunc(ec2conn, registerResp.ImageId),
		StepState: state,
	}

	ui.Say("Waiting for AMI to become ready...")
	if _, err := awscommon.WaitForState(&stateChange); err != nil {
		err := fmt.Errorf("Error waiting for AMI: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepRegisterAMI) Cleanup(state multistep.StateBag) {}

func buildRegisterOpts(config *Config, image *ec2.Image, blockDevices []ec2.BlockDeviceMapping) *ec2.RegisterImage {
	registerOpts := &ec2.RegisterImage{
		Name:           config.AMIName,
		Architecture:   image.Architecture,
		RootDeviceName: image.RootDeviceName,
		BlockDevices:   blockDevices,
		VirtType:       config.AMIVirtType,
	}

	if config.AMIVirtType != "hvm" {
		registerOpts.KernelId = image.KernelId
		registerOpts.RamdiskId = image.RamdiskId
	}

	return registerOpts
}
