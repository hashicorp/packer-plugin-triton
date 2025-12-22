// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package triton

import (
	"testing"
)

func testConfig(t *testing.T) *Config {
	return &Config{
		AccessConfig:        testAccessConfig(),
		SourceMachineConfig: testSourceMachineConfig(t),
		TargetImageConfig:   testTargetImageConfig(t),
	}
}
