// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package triton

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepStopMachine(t *testing.T) {
	state := testState(t)
	step := new(StepStopMachine)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*DriverMock)

	machineId := "test-machine-id"
	state.Put("machine", machineId)

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	step.Cleanup(state)

	if driver.StopMachineId != machineId {
		t.Fatalf("should've stopped machine (%s != %s)", driver.StopMachineId, machineId)
	}
}

func TestStepStopMachine_StopMachineError(t *testing.T) {
	state := testState(t)
	step := new(StepStopMachine)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*DriverMock)

	machineId := "test-machine-id"
	state.Put("machine", machineId)

	driver.StopMachineErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}
}

func TestStepStopMachine_WaitForMachineStoppedError(t *testing.T) {
	state := testState(t)
	step := new(StepStopMachine)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*DriverMock)

	machineId := "test-machine-id"
	state.Put("machine", machineId)

	driver.WaitForMachineStateErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}
}
