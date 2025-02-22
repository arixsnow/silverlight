/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

import (
	"testing"
)

func TestDecodePortName(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "normal port name",
			input:    []byte("eth0\x00rest"),
			expected: "eth0",
		},
		{
			name:     "empty string with null",
			input:    []byte("\x00"),
			expected: "",
		},
		{
			name:     "no null terminator",
			input:    []byte("eth0"),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodePortName(tt.input)
			if got != tt.expected {
				t.Errorf("decodePortName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestDeserializeHeader(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		wantHeader  Header
		wantErr     bool
		errContains string
	}{
		{
			name:        "invalid size",
			input:       []byte{0x01, 0x02, 0x03}, // less than 8 bytes
			wantHeader:  Header{},
			wantErr:     true,
			errContains: "Header size mismatch",
		},
		{
			name: "valid header",
			input: []byte{
				0x01,       // Version
				0x00,       // Type
				0x00, 0x08, // Length
				0x00, 0x00, 0x00, 0x01, // XID
			},
			wantHeader: Header{
				Version: 0x01,
				Type:    0x00,
				Length:  8,
				Xid:     1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeHeader(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializeHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil || err.Error() != tt.errContains {
					t.Errorf("DeserializeHeader() error = %v, want error containing %q", err, tt.errContains)
				}
				return
			}
			if got != tt.wantHeader {
				t.Errorf("DeserializeHeader() = %v, want %v", got, tt.wantHeader)
			}
		})
	}
}
