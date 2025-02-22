/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

import (
	"bytes"
	"testing"
)

func TestHeaderSerialize(t *testing.T) {
	tests := []struct {
		name     string
		header   Header
		expected []byte
	}{
		{
			name: "basic header",
			header: Header{
				Version: 0x01,
				Type:    0x00,
				Length:  8,
				Xid:     1,
			},
			expected: []byte{
				0x01,       // Version
				0x00,       // Type
				0x00, 0x08, // Length
				0x00, 0x00, 0x00, 0x01, // XID
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.header.Serialize()
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("Header.Serialize() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPhyPortSerialize(t *testing.T) {
	tests := []struct {
		name     string
		port     PhyPort
		expected []byte
	}{
		{
			name: "basic physical port",
			port: PhyPort{
				PortNo:     uint16(1),
				HWAddr:     []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
				Name:       "eth0",
				Config:     0,
				State:      0,
				Curr:       0,
				Advertised: 0,
				Supported:  0,
				Peer:       0,
			},
			expected: append(
				append(
					append(
						[]byte{
							0x00, 0x01, // PortNo as uint16
						},
						[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}..., // HWAddr
					),
					append(
						[]byte("eth0"),
						make([]byte, OFP_MAX_PORT_NAME_LEN-4)..., // Pad name to max length
					)...,
				),
				make([]byte, 24)..., // Config, State, Curr, Advertised, Supported, Peer (all 0)
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.port.Serialize()
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("PhyPort.Serialize() = %v, want %v", got, tt.expected)
			}
		})
	}
}
