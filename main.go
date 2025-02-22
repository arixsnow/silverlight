/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"net"

	"github.com/Arka-Mondal/silverlight/pkg/openflow10"
)

func main() {

	sf := &openflow10.SwitchFeatures{
		Header: openflow10.Header{
			Version: openflow10.OFP_VERSION,
			Type:    openflow10.OFPT_FEATURES_REPLY,
			Xid:     123,
		},
		DatapathID:   0x0000000000000001,
		N_Buffers:    256,
		N_Tables:     3,
		Capabilities: 0x00000001,
		Actions:      0x0000ffff,
		Ports: []openflow10.PhyPort{
			openflow10.PhyPort{
				PortNo: 1,
				HWAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
				Name:   "eth0",
			},
		},
	}

	// Serialize
	data := sf.Serialize()

	// Deserialize
	decoded, err := openflow10.DeserializeSwitchFeatures(data)
	if err != nil {
		fmt.Printf("Deserialization failed: %v", err)
	} else if decoded.DatapathID != sf.DatapathID {
		fmt.Printf("DatapathID mismatch: got %#x, want %#x", decoded.DatapathID, sf.DatapathID)
	} else {
		fmt.Printf("Decoded successfully...\ndecoded: %v\nactual: %v\n", decoded, sf)
	}
}
