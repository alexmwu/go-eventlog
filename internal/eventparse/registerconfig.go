// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package eventparse

import (
	"github.com/google/go-eventlog/tcg"
)

// RegisterConfig contains the measurement register technology-specific indexes
// expected to contain the events corresponding to various states, like EFI
// and Secure Boot states.
// This uses the event log-encoded index, e.g., PCR or CC MR (not RTMR).
type RegisterConfig struct {
	Name                          string
	FirmwareDriverIdx             uint32
	SecureBootIdx                 uint32
	EFIAppIdx                     uint32
	ExitBootServicesIdx           uint32
	GRUBCmdIdx                    uint32
	GRUBFileIdx                   uint32
	AdditionalSecureBootIdxEvents map[tcg.EventType]bool
}

var TPMRegisterConfig = RegisterConfig{
	Name:                "PCR",
	FirmwareDriverIdx:   2,
	SecureBootIdx:       7,
	EFIAppIdx:           4,
	ExitBootServicesIdx: 5,
	GRUBCmdIdx:          8,
	GRUBFileIdx:         9,
	// AdditionalSecureBootIdxEvents is empty since
	// eventparse.ParseSecurebootState encodes all the current allowable types
	// for PCR 7.
}

var RTMRRegisterConfig = RegisterConfig{
	Name: "RTMR",
	// CCMR2=RTMR[1]=PCR[2]
	FirmwareDriverIdx: 2,
	// CCMR1=RTMR[0]=PCR[7]
	SecureBootIdx: 1,
	// CCMR2=RTMR[1]=PCR[4]
	EFIAppIdx: 2,
	/// CCMR2=RTMR[1]=PCR[5]
	ExitBootServicesIdx: 2,
	// CCMR3=RTMR[2]=PCR[8]
	GRUBCmdIdx: 3,
	// CCMR3=RTMR[2]=PCR[9]
	GRUBFileIdx: 3,
	// RTMR[0] maps to both PCR[1] and PCR[7].
	// Pulled from "Table 27 Events" in
	// "TCG PC Client Platform Firmware Profile Specification"
	AdditionalSecureBootIdxEvents: map[tcg.EventType]bool{
		tcg.CpuMicrocode:            true,
		tcg.PlatformConfigFlags:     true,
		tcg.TableOfDevices:          true,
		tcg.NonhostConfig:           true,
		tcg.EFIVariableDriverConfig: true,
		tcg.EFIVariableBoot:         true,
		tcg.EFIAction:               true,
		tcg.EFIHandoffTables2:       true,
		//tcg.EFIVariableBoot2:        true,
		// https://github.com/tianocore/edk2/blob/a29a9cce5f9afa32560d966e501247246ec96ef6/OvmfPkg/IntelTdx/TdxHelperLib/TdxMeasurementHob.c#L245
		// The following is not spec-compliant and we should remove it.
		tcg.EFIPlatformFirmwareBlob2: true,
	},
}