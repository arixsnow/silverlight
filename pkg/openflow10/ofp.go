/*
 * Copyright (c) 2025, Arka Mondal. All rights reserved.
 * Use of this source code is governed by a BSD-style license that
 * can be found in the LICENSE file.
 */

package openflow10

import "net"

// OpenFlow v1.0 version
const OFP_VERSION = 0x01

// OpenFlow v1.0 types (OFPT_ constants)
const (
	// Immutable messages
	OFPT_HELLO        = iota // Semmetric message
	OFPT_ERROR               // Semmetric message
	OFPT_ECHO_REQUEST        // Semmetric message
	OFPT_ECHO_REPLY          // Semmetric message
	OFPT_VENDOR              // Semmetric message

	// Switch configuration messages
	OFPT_FEATURES_REQUEST   // Controller/switch message
	OFPT_FEATURES_REPLY     // Controller/switch message
	OFPT_GET_CONFIG_REQUEST // Controller/switch message
	OFPT_GET_CONFIG_REPLY   // Controller/switch message
	OFPT_SET_CONFIG         // Controller/switch message

	// Asynchronous messages
	OFPT_PACKET_IN    // Async message
	OFPT_FLOW_REMOVED // Async message
	OFPT_PORT_STATUS  // Async message

	// Controller command messages
	OFPT_PACKET_OUT // Controller/switch message
	OFPT_FLOW_MOD   // Controller/switch message
	OFPT_PORT_MOD   // Controller/switch message

	// Statistics messages
	OFPT_STATS_REQUEST // Controller/switch message
	OFPT_STATS_REPLY   // Controller/switch message

	// Barrier messages
	OFPT_BARRIER_REQUEST // Controller/switch message
	OFPT_BARRIER_REPLY   // Controller/switch message

	// Queue Configuration messages
	OFPT_QUEUE_GET_CONFIG_REQUEST // Controller/switch message
	OFPT_QUEUE_GET_CONFIG_REPLY   // Controller/switch message
)

// OpenFlow v1.0 constants
const OFP_MAX_ETH_ALEN = 6
const OFP_MAX_PORT_NAME_LEN = 16

// OpenFlow v1.0 header
type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	Xid     uint32
}

// OpenFlow v1.0 description of a physical port
type PhyPort struct {
	PortNo uint16
	HWAddr net.HardwareAddr
	Name   string

	Config uint32 // Bitmap of OFPPC_* flags
	State  uint32 // Bitmap of OFPPS_* flags

	/* Bitmaps of OFPPF_* that describe features. All bits zeroed if
	 * unsupported or unavailable. */
	Curr       uint32 // Current features
	Advertised uint32 // Features being advertised by the port
	Supported  uint32 // Features supported by the port
	Peer       uint32 // Features advertised by peer
}

/* OpenFlow v1.0 Flags to indicate behavior of the physical port.
 * These flags are used in PhyPort to describe the current
 * configuration. They are used in the PortMod message to configure the
 * port's behavior.
 */

type PortConfig uint32

const (
	OFPPC_PORT_DOWN    PortConfig = 1 << iota // Port is administratively down
	OFPPC_NO_STP                              // Disable 802.1D spanning tree on port
	OFPPC_NO_RECV                             // Drop all packets except 802.1D spanning tree packets
	OFPPC_NO_RECV_STP                         // Drop received 802.1D STP packets
	OFPPC_NO_FLOOD                            // Do not include this port when flooding
	OFPPC_NO_FWD                              // Drop packets forwarded to port
	OFPPC_NO_PACKET_IN                        // Do not send packet-in msgs for port
)

/* OpenFlow v1.0 Current state of the physical port. These are not
 * configurable from the controller.
 */

type PortState uint32

const (
	OFPPS_LINK_DOWN   PortState = 1 << 0 // No physical link present
	OFPPS_STP_LISTEN  PortState = 0 << 8 // Not learning or relaying frames
	OFPPS_STP_LEARN   PortState = 1 << 8 // Learning but not relaying frames
	OFPPS_STP_FORWARD PortState = 2 << 8 // Learning and relaying frames
	OFPPS_STP_BLOCK   PortState = 3 << 8 // Not part of spanning tree
	OFPPS_STP_MASK    PortState = 3 << 8 // Bit mask for OFPPS_STP_* values
)

/* OpenFlow v1.0 Port numbering. Physical ports are numbered starting from 1 */

type PortType uint16

const (
	// Maximum number of physical switch ports
	OFPP_MAX PortType = 0xff00

	// Fake output "ports"

	/*
		   * Send the packet out` the input port. This
			 * virtual port must be explicitly used
			 * in order to send back out the input port
	*/
	OFPP_IN_PORT PortType = 0xfff8

	/*
	 * Perform actions in flow table.
	 * NB: This can only be the destination port for packet-out messages
	 */
	OFPP_TABLE PortType = 0xfff9

	/* Process with normal L2/L3 swtiching */
	OFPP_NORMAL PortType = 0xfffa

	/* All physical ports except input port and those disabled by STP */
	OFPP_FLOOD PortType = 0xfffb

	/* All physical ports except input port */
	OFPP_ALL PortType = 0xfffc

	/* Send to controller */
	OFPP_CONTROLLER PortType = 0xfffd

	/* Local openflow "port" */
	OFPP_LOCAL PortType = 0xfffe

	/* Not associated with a physical port */
	OFPP_NONE PortType = 0xffff
)
