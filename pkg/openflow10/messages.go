/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

type Hello struct {
	Header Header
}

type FeaturesRequest struct {
	Header Header
}

type SwitchFeatures struct {
	Header       Header
	DatapathID   uint64
	N_Buffers    uint32
	N_Tables     uint8
	Capabilities uint32
	Actions      uint32
	Ports        []PhyPort
}
