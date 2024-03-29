// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package triton

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepCreateImageFromMachine(t *testing.T) {
	state := testState(t)
	step := new(StepCreateImageFromMachine)
	defer step.Cleanup(state)

	state.Put("machine", "test-machine-id")

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	_, ok := state.GetOk("image")
	if !ok {
		t.Fatalf("should have image")
	}

	step.Cleanup(state)
}

func TestStepCreateImageFromMachine_CreateImageFromMachineError(t *testing.T) {
	state := testState(t)
	step := new(StepCreateImageFromMachine)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*DriverMock)
	state.Put("machine", "test-machine-id")

	driver.CreateImageFromMachineErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}

	if _, ok := state.GetOk("image"); ok {
		t.Fatalf("should NOT have image")
	}
}

func TestStepCreateImageFromMachine_WaitForImageCreationError(t *testing.T) {
	state := testState(t)
	step := new(StepCreateImageFromMachine)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*DriverMock)
	state.Put("machine", "test-machine-id")

	driver.WaitForImageCreationErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}

	if _, ok := state.GetOk("image"); ok {
		t.Fatalf("should NOT have image")
	}
}
