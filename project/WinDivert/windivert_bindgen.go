package windivert

import (
	"fmt"
)

// Source: windivert_bindgen.h
type (
	UINT = uint32
	PWINDIVERT_LAYER = *uint32
	PWINDIVERT_DATA_NETWORK = *struct{}
	PWINDIVERT_DATA_FLOW = *struct{}
	PWINDIVERT_DATA_SOCKET = *struct{}
	PWINDIVERT_DATA_REFLECT = *struct{}
	PWINDIVERT_ADDRESS = *struct{}
	PWINDIVERT_EVENT = *uint32
	PWINDIVERT_PARAM = *uint32
	PWINDIVERT_SHUTDOWN = *uint32
	PWINDIVERT_IPHDR = *struct{}
	PWINDIVERT_IPV6HDR = *struct{}
	PWINDIVERT_ICMPHDR = *struct{}
	PWINDIVERT_ICMPV6HDR = *struct{}
	PWINDIVERT_TCPHDR = *struct{}
	PWINDIVERT_UDPHDR = *struct{}
)

// Source: windivert_bindgen.h:0 -> WINDIVERT_LAYER
type WINDIVERT_LAYER uint32

const (
	WindivertLayerNetwork WINDIVERT_LAYER = 0
	WindivertLayerNetworkForward WINDIVERT_LAYER = 1
	WindivertLayerFlow WINDIVERT_LAYER = 2
	WindivertLayerSocket WINDIVERT_LAYER = 3
	WindivertLayerReflect WINDIVERT_LAYER = 4
)

func (w WINDIVERT_LAYER) String() string {
	switch w {
	case WindivertLayerNetwork:
		return "Windivert Layer Network"
	case WindivertLayerNetworkForward:
		return "Windivert Layer Network Forward"
	case WindivertLayerFlow:
		return "Windivert Layer Flow"
	case WindivertLayerSocket:
		return "Windivert Layer Socket"
	case WindivertLayerReflect:
		return "Windivert Layer Reflect"
	default:
		return fmt.Sprintf("WINDIVERT_LAYER(0x%X)", uint32(w))
	}
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_EVENT
type WINDIVERT_EVENT uint32

const (
	WindivertEventNetworkPacket WINDIVERT_EVENT = 0
	WindivertEventFlowEstablished WINDIVERT_EVENT = 1
	WindivertEventFlowDeleted WINDIVERT_EVENT = 2
	WindivertEventSocketBind WINDIVERT_EVENT = 3
	WindivertEventSocketConnect WINDIVERT_EVENT = 4
	WindivertEventSocketListen WINDIVERT_EVENT = 5
	WindivertEventSocketAccept WINDIVERT_EVENT = 6
	WindivertEventSocketClose WINDIVERT_EVENT = 7
	WindivertEventReflectOpen WINDIVERT_EVENT = 8
	WindivertEventReflectClose WINDIVERT_EVENT = 9
)

func (w WINDIVERT_EVENT) String() string {
	switch w {
	case WindivertEventNetworkPacket:
		return "Windivert Event Network Packet"
	case WindivertEventFlowEstablished:
		return "Windivert Event Flow Established"
	case WindivertEventFlowDeleted:
		return "Windivert Event Flow Deleted"
	case WindivertEventSocketBind:
		return "Windivert Event Socket Bind"
	case WindivertEventSocketConnect:
		return "Windivert Event Socket Connect"
	case WindivertEventSocketListen:
		return "Windivert Event Socket Listen"
	case WindivertEventSocketAccept:
		return "Windivert Event Socket Accept"
	case WindivertEventSocketClose:
		return "Windivert Event Socket Close"
	case WindivertEventReflectOpen:
		return "Windivert Event Reflect Open"
	case WindivertEventReflectClose:
		return "Windivert Event Reflect Close"
	default:
		return fmt.Sprintf("WINDIVERT_EVENT(0x%X)", uint32(w))
	}
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_PARAM
type WINDIVERT_PARAM uint32

const (
	WindivertParamQueueLength WINDIVERT_PARAM = 0
	WindivertParamQueueTime WINDIVERT_PARAM = 1
	WindivertParamQueueSize WINDIVERT_PARAM = 2
	WindivertParamVersionMajor WINDIVERT_PARAM = 3
	WindivertParamVersionMinor WINDIVERT_PARAM = 4
)

func (w WINDIVERT_PARAM) String() string {
	switch w {
	case WindivertParamQueueLength:
		return "Windivert Param Queue Length"
	case WindivertParamQueueTime:
		return "Windivert Param Queue Time"
	case WindivertParamQueueSize:
		return "Windivert Param Queue Size"
	case WindivertParamVersionMajor:
		return "Windivert Param Version Major"
	case WindivertParamVersionMinor:
		return "Windivert Param Version Minor"
	default:
		return fmt.Sprintf("WINDIVERT_PARAM(0x%X)", uint32(w))
	}
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_SHUTDOWN
type WINDIVERT_SHUTDOWN uint32

const (
	WindivertShutdownRecv WINDIVERT_SHUTDOWN = 1
	WindivertShutdownSend WINDIVERT_SHUTDOWN = 2
	WindivertShutdownBoth WINDIVERT_SHUTDOWN = 3
)

func (w WINDIVERT_SHUTDOWN) String() string {
	switch w {
	case WindivertShutdownRecv:
		return "Windivert Shutdown Recv"
	case WindivertShutdownSend:
		return "Windivert Shutdown Send"
	case WindivertShutdownBoth:
		return "Windivert Shutdown Both"
	default:
		return fmt.Sprintf("WINDIVERT_SHUTDOWN(0x%X)", uint32(w))
	}
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_DATA_NETWORK
type WINDIVERT_DATA_NETWORK struct {
	IfIdx uint32
	SubIfIdx uint32
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_DATA_FLOW
type WINDIVERT_DATA_FLOW struct {
	EndpointId uint64
	ParentEndpointId uint64
	ProcessId uint32
	LocalAddr [4]uint32
	RemoteAddr [4]uint32
	LocalPort uint16
	RemotePort uint16
	Protocol uint8
	_  [7]byte
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_DATA_SOCKET
type WINDIVERT_DATA_SOCKET struct {
	EndpointId uint64
	ParentEndpointId uint64
	ProcessId uint32
	LocalAddr [4]uint32
	RemoteAddr [4]uint32
	LocalPort uint16
	RemotePort uint16
	Protocol uint8
	_  [7]byte
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_DATA_REFLECT
type WINDIVERT_DATA_REFLECT struct {
	Timestamp int64
	ProcessId uint32
	Layer WINDIVERT_LAYER
	Flags uint64
	Priority int16
	_  [6]byte
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_ADDRESS
type WINDIVERT_ADDRESS struct {
	Timestamp int64
	Layer uint32
	Event uint32
	Sniffed uint32
	Outbound uint32
	Loopback uint32
	Impostor uint32
	IPv6 uint32
	IPChecksum uint32
	TCPChecksum uint32
	UDPChecksum uint32
	Reserved1 uint32
	Reserved2 uint32
	U WINDIVERT_ADDRESS_Anon1Union
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_ADDRESS_Anon1Union
type WINDIVERT_ADDRESS_Anon1Union struct {
	_ WINDIVERT_DATA_FLOW
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_IPHDR
type WINDIVERT_IPHDR struct {
	HdrLength uint8
	Version uint8
	TOS uint8
	_  [1]byte
	Length uint16
	Id uint16
	FragOff0 uint16
	TTL uint8
	Protocol uint8
	Checksum uint16
	_  [2]byte
	SrcAddr uint32
	DstAddr uint32
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_IPV6HDR
type WINDIVERT_IPV6HDR struct {
	TrafficClass0 uint8
	Version uint8
	FlowLabel0 uint8
	TrafficClass1 uint8
	FlowLabel1 uint16
	Length uint16
	NextHdr uint8
	HopLimit uint8
	_  [2]byte
	SrcAddr [4]uint32
	DstAddr [4]uint32
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_ICMPHDR
type WINDIVERT_ICMPHDR struct {
	Type uint8
	Code uint8
	Checksum uint16
	Body uint32
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_ICMPV6HDR
type WINDIVERT_ICMPV6HDR struct {
	Type uint8
	Code uint8
	Checksum uint16
	Body uint32
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_TCPHDR
type WINDIVERT_TCPHDR struct {
	SrcPort uint16
	DstPort uint16
	SeqNum uint32
	AckNum uint32
	Reserved1 uint16
	HdrLength uint16
	Fin uint16
	Syn uint16
	Rst uint16
	Psh uint16
	Ack uint16
	Urg uint16
	Reserved2 uint16
	Window uint16
	Checksum uint16
	UrgPtr uint16
}

// Source: windivert_bindgen.h:0 -> WINDIVERT_UDPHDR
type WINDIVERT_UDPHDR struct {
	SrcPort uint16
	DstPort uint16
	Length uint16
	Checksum uint16
}

