/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

import (
	"bytes"
	"encoding/binary"
)

func (h *Header) Serialize() []byte {
	buf := new(bytes.Buffer)

	buf.WriteByte(h.Version)
	buf.WriteByte(h.Type)
	binary.Write(buf, binary.BigEndian, h.Length)
	binary.Write(buf, binary.BigEndian, h.Xid)

	return buf.Bytes()
}

func (p *PhyPort) Serialize() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, p.PortNo)
	buf.Write(p.HWAddr)

	nameBytes := make([]byte, OFP_MAX_PORT_NAME_LEN)
	copy(nameBytes, p.Name)
	buf.Write(nameBytes)
	binary.Write(buf, binary.BigEndian, p.Config)
	binary.Write(buf, binary.BigEndian, p.State)
	binary.Write(buf, binary.BigEndian, p.Curr)
	binary.Write(buf, binary.BigEndian, p.Advertised)
	binary.Write(buf, binary.BigEndian, p.Supported)
	binary.Write(buf, binary.BigEndian, p.Peer)

	return buf.Bytes()
}

func (sf *SwitchFeatures) Serialize() []byte {
	buf := new(bytes.Buffer)

	buf.Write(sf.Header.Serialize())
	binary.Write(buf, binary.BigEndian, sf.DatapathID)
	binary.Write(buf, binary.BigEndian, sf.N_Buffers)

	buf.WriteByte(sf.N_Tables)
	buf.Write([]byte{0, 0, 0})

	binary.Write(buf, binary.BigEndian, sf.Capabilities)

	binary.Write(buf, binary.BigEndian, sf.Actions)

	for _, port := range sf.Ports {
		buf.Write(port.Serialize())
	}

	sf.Header.Length = uint16(buf.Len())
	headerBytes := sf.Header.Serialize()
	copy(buf.Bytes()[:8], headerBytes)

	return buf.Bytes()
}
