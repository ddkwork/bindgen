package windivert

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Windivert struct{}

//go:embed WinDivert64.dll
var dllBytes []byte

var (
	dll            *windows.LazyDLL
	proc_WinDivertOpen      *windows.LazyProc
	proc_WinDivertRecv      *windows.LazyProc
	proc_WinDivertRecvEx      *windows.LazyProc
	proc_WinDivertSend      *windows.LazyProc
	proc_WinDivertSendEx      *windows.LazyProc
	proc_WinDivertShutdown      *windows.LazyProc
	proc_WinDivertClose      *windows.LazyProc
	proc_WinDivertSetParam      *windows.LazyProc
	proc_WinDivertGetParam      *windows.LazyProc
	proc_WinDivertHelperHashPacket      *windows.LazyProc
	proc_WinDivertHelperParsePacket      *windows.LazyProc
	proc_WinDivertHelperParseIPv4Address      *windows.LazyProc
	proc_WinDivertHelperParseIPv6Address      *windows.LazyProc
	proc_WinDivertHelperFormatIPv4Address      *windows.LazyProc
	proc_WinDivertHelperFormatIPv6Address      *windows.LazyProc
	proc_WinDivertHelperCalcChecksums      *windows.LazyProc
	proc_WinDivertHelperDecrementTTL      *windows.LazyProc
	proc_WinDivertHelperCompileFilter      *windows.LazyProc
	proc_WinDivertHelperEvalFilter      *windows.LazyProc
	proc_WinDivertHelperFormatFilter      *windows.LazyProc
	proc_WinDivertHelperNtohs      *windows.LazyProc
	proc_WinDivertHelperHtons      *windows.LazyProc
	proc_WinDivertHelperNtohl      *windows.LazyProc
	proc_WinDivertHelperHtonl      *windows.LazyProc
	proc_WinDivertHelperNtohll      *windows.LazyProc
	proc_WinDivertHelperHtonll      *windows.LazyProc
	proc_WinDivertHelperNtohIPv6Address      *windows.LazyProc
	proc_WinDivertHelperHtonIPv6Address      *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "WinDivert64.dll"))
	proc_WinDivertOpen = dll.NewProc("WinDivertOpen")
	proc_WinDivertRecv = dll.NewProc("WinDivertRecv")
	proc_WinDivertRecvEx = dll.NewProc("WinDivertRecvEx")
	proc_WinDivertSend = dll.NewProc("WinDivertSend")
	proc_WinDivertSendEx = dll.NewProc("WinDivertSendEx")
	proc_WinDivertShutdown = dll.NewProc("WinDivertShutdown")
	proc_WinDivertClose = dll.NewProc("WinDivertClose")
	proc_WinDivertSetParam = dll.NewProc("WinDivertSetParam")
	proc_WinDivertGetParam = dll.NewProc("WinDivertGetParam")
	proc_WinDivertHelperHashPacket = dll.NewProc("WinDivertHelperHashPacket")
	proc_WinDivertHelperParsePacket = dll.NewProc("WinDivertHelperParsePacket")
	proc_WinDivertHelperParseIPv4Address = dll.NewProc("WinDivertHelperParseIPv4Address")
	proc_WinDivertHelperParseIPv6Address = dll.NewProc("WinDivertHelperParseIPv6Address")
	proc_WinDivertHelperFormatIPv4Address = dll.NewProc("WinDivertHelperFormatIPv4Address")
	proc_WinDivertHelperFormatIPv6Address = dll.NewProc("WinDivertHelperFormatIPv6Address")
	proc_WinDivertHelperCalcChecksums = dll.NewProc("WinDivertHelperCalcChecksums")
	proc_WinDivertHelperDecrementTTL = dll.NewProc("WinDivertHelperDecrementTTL")
	proc_WinDivertHelperCompileFilter = dll.NewProc("WinDivertHelperCompileFilter")
	proc_WinDivertHelperEvalFilter = dll.NewProc("WinDivertHelperEvalFilter")
	proc_WinDivertHelperFormatFilter = dll.NewProc("WinDivertHelperFormatFilter")
	proc_WinDivertHelperNtohs = dll.NewProc("WinDivertHelperNtohs")
	proc_WinDivertHelperHtons = dll.NewProc("WinDivertHelperHtons")
	proc_WinDivertHelperNtohl = dll.NewProc("WinDivertHelperNtohl")
	proc_WinDivertHelperHtonl = dll.NewProc("WinDivertHelperHtonl")
	proc_WinDivertHelperNtohll = dll.NewProc("WinDivertHelperNtohll")
	proc_WinDivertHelperHtonll = dll.NewProc("WinDivertHelperHtonll")
	proc_WinDivertHelperNtohIPv6Address = dll.NewProc("WinDivertHelperNtohIPv6Address")
	proc_WinDivertHelperHtonIPv6Address = dll.NewProc("WinDivertHelperHtonIPv6Address")
}

func saveEmbeddedDLL(data []byte, name string) string {
	tmpDir := os.TempDir()
	p := filepath.Join(tmpDir, name)
	if _, err := os.Stat(p); err == nil {
		return p
	}
	os.WriteFile(p, data, 0644)
	return p
}

func (w *Windivert) WinDivertOpen(Filter *int8, Layer WINDIVERT_LAYER, Priority int16, Flags uint64) uintptr {
	r1, _, _ := proc_WinDivertOpen.Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(Priority), uintptr(Flags))
	return uintptr(r1)
}

func (w *Windivert) WinDivertRecv(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PRecvLen *uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := proc_WinDivertRecv.Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PRecvLen)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertRecvEx(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PRecvLen *uint32, Flags uint64, PAddr *WINDIVERT_ADDRESS, PAddrLen *uint32, LpOverlapped unsafe.Pointer) int32 {
	r1, _, _ := proc_WinDivertRecvEx.Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PRecvLen)), uintptr(Flags), uintptr(unsafe.Pointer(PAddr)), uintptr(unsafe.Pointer(PAddrLen)), uintptr(LpOverlapped))
	return int32(r1)
}

func (w *Windivert) WinDivertSend(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PSendLen *uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := proc_WinDivertSend.Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PSendLen)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertSendEx(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PSendLen *uint32, Flags uint64, PAddr *WINDIVERT_ADDRESS, AddrLen uint32, LpOverlapped unsafe.Pointer) int32 {
	r1, _, _ := proc_WinDivertSendEx.Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PSendLen)), uintptr(Flags), uintptr(unsafe.Pointer(PAddr)), uintptr(AddrLen), uintptr(LpOverlapped))
	return int32(r1)
}

func (w *Windivert) WinDivertShutdown(Handle uintptr, How WINDIVERT_SHUTDOWN) int32 {
	r1, _, _ := proc_WinDivertShutdown.Call(Handle, uintptr(How))
	return int32(r1)
}

func (w *Windivert) WinDivertClose(Handle uintptr) int32 {
	r1, _, _ := proc_WinDivertClose.Call(Handle)
	return int32(r1)
}

func (w *Windivert) WinDivertSetParam(Handle uintptr, Param WINDIVERT_PARAM, Value uint64) int32 {
	r1, _, _ := proc_WinDivertSetParam.Call(Handle, uintptr(Param), uintptr(Value))
	return int32(r1)
}

func (w *Windivert) WinDivertGetParam(Handle uintptr, Param WINDIVERT_PARAM, PValue *uint64) int32 {
	r1, _, _ := proc_WinDivertGetParam.Call(Handle, uintptr(Param), uintptr(unsafe.Pointer(PValue)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperHashPacket(PPacket unsafe.Pointer, PacketLen uint32, Seed uint64) uint64 {
	r1, _, _ := proc_WinDivertHelperHashPacket.Call(uintptr(PPacket), uintptr(PacketLen), uintptr(Seed))
	return uint64(r1)
}

func (w *Windivert) WinDivertHelperParsePacket(PPacket unsafe.Pointer, PacketLen uint32, PpIpHdr *PWINDIVERT_IPHDR, PpIpv6Hdr *PWINDIVERT_IPV6HDR, PProtocol *uint8, PpIcmpHdr *PWINDIVERT_ICMPHDR, PpIcmpv6Hdr *PWINDIVERT_ICMPV6HDR, PpTcpHdr *PWINDIVERT_TCPHDR, PpUdpHdr *PWINDIVERT_UDPHDR, PpData *uintptr, PDataLen *uint32, PpNext *uintptr, PNextLen *uint32) int32 {
	r1, _, _ := proc_WinDivertHelperParsePacket.Call(uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PpIpHdr)), uintptr(unsafe.Pointer(PpIpv6Hdr)), uintptr(unsafe.Pointer(PProtocol)), uintptr(unsafe.Pointer(PpIcmpHdr)), uintptr(unsafe.Pointer(PpIcmpv6Hdr)), uintptr(unsafe.Pointer(PpTcpHdr)), uintptr(unsafe.Pointer(PpUdpHdr)), uintptr(unsafe.Pointer(PpData)), uintptr(unsafe.Pointer(PDataLen)), uintptr(unsafe.Pointer(PpNext)), uintptr(unsafe.Pointer(PNextLen)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperParseIPv4Address(AddrStr *int8, PAddr *uint32) int32 {
	r1, _, _ := proc_WinDivertHelperParseIPv4Address.Call(uintptr(unsafe.Pointer(AddrStr)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperParseIPv6Address(AddrStr *int8, PAddr *uint32) int32 {
	r1, _, _ := proc_WinDivertHelperParseIPv6Address.Call(uintptr(unsafe.Pointer(AddrStr)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatIPv4Address(Addr uint32, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := proc_WinDivertHelperFormatIPv4Address.Call(uintptr(Addr), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatIPv6Address(PAddr *uint32, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := proc_WinDivertHelperFormatIPv6Address.Call(uintptr(unsafe.Pointer(PAddr)), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperCalcChecksums(PPacket unsafe.Pointer, PacketLen uint32, PAddr *WINDIVERT_ADDRESS, Flags uint64) int32 {
	r1, _, _ := proc_WinDivertHelperCalcChecksums.Call(uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PAddr)), uintptr(Flags))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperDecrementTTL(PPacket unsafe.Pointer, PacketLen uint32) int32 {
	r1, _, _ := proc_WinDivertHelperDecrementTTL.Call(uintptr(PPacket), uintptr(PacketLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperCompileFilter(Filter *int8, Layer WINDIVERT_LAYER, Object *int8, ObjLen uint32, ErrorStr **int8, ErrorPos *uint32) int32 {
	r1, _, _ := proc_WinDivertHelperCompileFilter.Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(unsafe.Pointer(Object)), uintptr(ObjLen), uintptr(unsafe.Pointer(ErrorStr)), uintptr(unsafe.Pointer(ErrorPos)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperEvalFilter(Filter *int8, PPacket unsafe.Pointer, PacketLen uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := proc_WinDivertHelperEvalFilter.Call(uintptr(unsafe.Pointer(Filter)), uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatFilter(Filter *int8, Layer WINDIVERT_LAYER, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := proc_WinDivertHelperFormatFilter.Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperNtohs(X uint16) uint16 {
	r1, _, _ := proc_WinDivertHelperNtohs.Call(uintptr(X))
	return uint16(r1)
}

func (w *Windivert) WinDivertHelperHtons(X uint16) uint16 {
	r1, _, _ := proc_WinDivertHelperHtons.Call(uintptr(X))
	return uint16(r1)
}

func (w *Windivert) WinDivertHelperNtohl(X uint32) uint32 {
	r1, _, _ := proc_WinDivertHelperNtohl.Call(uintptr(X))
	return uint32(r1)
}

func (w *Windivert) WinDivertHelperHtonl(X uint32) uint32 {
	r1, _, _ := proc_WinDivertHelperHtonl.Call(uintptr(X))
	return uint32(r1)
}

func (w *Windivert) WinDivertHelperNtohll(X uint64) uint64 {
	r1, _, _ := proc_WinDivertHelperNtohll.Call(uintptr(X))
	return uint64(r1)
}

func (w *Windivert) WinDivertHelperHtonll(X uint64) uint64 {
	r1, _, _ := proc_WinDivertHelperHtonll.Call(uintptr(X))
	return uint64(r1)
}

func (w *Windivert) WinDivertHelperNtohIPv6Address(InAddr *uint32, OutAddr *uint32) {
	proc_WinDivertHelperNtohIPv6Address.Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}

func (w *Windivert) WinDivertHelperHtonIPv6Address(InAddr *uint32, OutAddr *uint32) {
	proc_WinDivertHelperHtonIPv6Address.Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}

