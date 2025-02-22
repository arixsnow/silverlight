/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

import (
	"bytes"
	"errors"
  "encoding/binary"
	"fmt"
	"net"
)

func decodePortName(name []byte) string {
	nullIndex := bytes.IndexByte(name, 0x00)

	if nullIndex == -1 {
		return ""
	}

	return string(name[:nullIndex])
}

func DeserializeHeader(data []byte) (Header, error) {
	if len(data) != 8 {
		return Header{}, errors.New("Header size mismatch")
	}

	return Header{
		Version: data[0],
		Type:    data[1],
		Length:  binary.BigEndian.Uint16(data[2:4]),
		Xid:     binary.BigEndian.Uint32(data[4:8]),
	}, nil
}

func DeserializePhyPort(data []byte) (*PhyPort, error) {
	if len(data) < 48 {
		return nil, errors.New("PhyPort size mismatch")
	}

	port := &PhyPort{}

	port.PortNo = binary.BigEndian.Uint16(data[:2])
	port.HWAddr = net.HardwareAddr(data[2:8])
	port.Name = decodePortName(data[8:24])
	port.Config = binary.BigEndian.Uint32(data[24:28])
	port.State = binary.BigEndian.Uint32(data[28:32])
	port.Curr = binary.BigEndian.Uint32(data[32:36])
	port.Advertised = binary.BigEndian.Uint32(data[36:40])
	port.Supported = binary.BigEndian.Uint32(data[40:44])
	port.Peer = binary.BigEndian.Uint32(data[44:48])

	return port, nil
}

func DeserializeSwitchFeatures(data []byte) (*SwitchFeatures, error) {
	if len(data) < 32 {
		return nil, errors.New("SwitchFeatures size mismatch")
	}

	sf := &SwitchFeatures{}

	header, err := DeserializeHeader(data[:8])
	if err != nil {
		return nil, err
	}

	sf.Header = header
	sf.DatapathID = binary.BigEndian.Uint64(data[8:16])
	sf.N_Buffers = binary.BigEndian.Uint32(data[16:20])
	sf.N_Tables = data[20]
	sf.Capabilities = binary.BigEndian.Uint32(data[24:28])
	sf.Actions = binary.BigEndian.Uint32(data[28:32])

	// Ports [offset 32 onwards, 48 bytes each]
	portsData := data[32:]
	numPorts := len(portsData) / 48

	for i := 0; i < numPorts; i++ {
		start := i * 48
		end := start + 48

		if end > len(portsData) {
			break
		}

		port, err := DeserializePhyPort(portsData[start:end])
		if err != nil {
			return nil, fmt.Errorf("Error in parsing port %d %v", i, err)
		}

		sf.Ports = append(sf.Ports, *port)
	}

	return sf, nil
}
