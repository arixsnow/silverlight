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
	if err := binary.Write(buf, binary.BigEndian, h.Length); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, h.Xid); err != nil {
		return nil
	}

	return buf.Bytes()
}

func (p *PhyPort) Serialize() []byte {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, p.PortNo); err != nil {
		return nil
	}
	buf.Write(p.HWAddr)

	nameBytes := make([]byte, OFP_MAX_PORT_NAME_LEN)
	copy(nameBytes, p.Name)
	buf.Write(nameBytes)
	if err := binary.Write(buf, binary.BigEndian, p.Config); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, p.State); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, p.Curr); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, p.Advertised); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, p.Supported); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, p.Peer); err != nil {
		return nil
	}

	return buf.Bytes()
}

func (sf *SwitchFeatures) Serialize() []byte {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, sf.Header); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, sf.DatapathID); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, sf.N_Buffers); err != nil {
		return nil
	}

	buf.WriteByte(sf.N_Tables)
	buf.Write([]byte{0, 0, 0})

	if err := binary.Write(buf, binary.BigEndian, sf.Capabilities); err != nil {
		return nil
	}

	if err := binary.Write(buf, binary.BigEndian, sf.Actions); err != nil {
		return nil
	}

	for _, port := range sf.Ports {
		if _, err := buf.Write(port.Serialize()); err != nil {
			return nil
		}
	}

	sf.Header.Length = uint16(buf.Len())
	headerBytes := sf.Header.Serialize()
	copy(buf.Bytes()[:8], headerBytes)

	return buf.Bytes()
}
