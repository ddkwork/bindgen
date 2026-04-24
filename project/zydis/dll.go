package zydis

import (
	_ "embed"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Zydis struct{}

//go:embed ZydisDll.dll
var dllBytes []byte

var (
	dll            *windows.LazyDLL
	proc_ZydisCategoryGetString      *windows.LazyProc
	proc_ZydisISASetGetString      *windows.LazyProc
	proc_ZydisISAExtGetString      *windows.LazyProc
	proc_ZydisMnemonicGetString      *windows.LazyProc
	proc_ZydisMnemonicGetStringWrapped      *windows.LazyProc
	proc_ZydisRegisterEncode      *windows.LazyProc
	proc_ZydisRegisterGetId      *windows.LazyProc
	proc_ZydisRegisterGetClass      *windows.LazyProc
	proc_ZydisRegisterGetWidth      *windows.LazyProc
	proc_ZydisRegisterGetLargestEnclosing      *windows.LazyProc
	proc_ZydisRegisterGetString      *windows.LazyProc
	proc_ZydisRegisterGetStringWrapped      *windows.LazyProc
	proc_ZydisRegisterClassGetWidth      *windows.LazyProc
	proc_ZydisDecoderInit      *windows.LazyProc
	proc_ZydisDecoderEnableMode      *windows.LazyProc
	proc_ZydisDecoderDecodeFull      *windows.LazyProc
	proc_ZydisDecoderDecodeInstruction      *windows.LazyProc
	proc_ZydisDecoderDecodeOperands      *windows.LazyProc
	proc_ZydisEncoderEncodeInstruction      *windows.LazyProc
	proc_ZydisEncoderEncodeInstructionAbsolute      *windows.LazyProc
	proc_ZydisEncoderDecodedInstructionToEncoderRequest      *windows.LazyProc
	proc_ZydisEncoderNopFill      *windows.LazyProc
	proc_ZyanAllocatorInit      *windows.LazyProc
	proc_ZyanAllocatorDefault      *windows.LazyProc
	proc_ZyanVectorInit      *windows.LazyProc
	proc_ZyanVectorInitEx      *windows.LazyProc
	proc_ZyanVectorInitCustomBuffer      *windows.LazyProc
	proc_ZyanVectorDestroy      *windows.LazyProc
	proc_ZyanVectorDuplicate      *windows.LazyProc
	proc_ZyanVectorDuplicateEx      *windows.LazyProc
	proc_ZyanVectorDuplicateCustomBuffer      *windows.LazyProc
	proc_ZyanVectorGet      *windows.LazyProc
	proc_ZyanVectorGetMutable      *windows.LazyProc
	proc_ZyanVectorGetPointer      *windows.LazyProc
	proc_ZyanVectorGetPointerMutable      *windows.LazyProc
	proc_ZyanVectorSet      *windows.LazyProc
	proc_ZyanVectorPushBack      *windows.LazyProc
	proc_ZyanVectorInsert      *windows.LazyProc
	proc_ZyanVectorInsertRange      *windows.LazyProc
	proc_ZyanVectorEmplace      *windows.LazyProc
	proc_ZyanVectorEmplaceEx      *windows.LazyProc
	proc_ZyanVectorSwapElements      *windows.LazyProc
	proc_ZyanVectorDelete      *windows.LazyProc
	proc_ZyanVectorDeleteRange      *windows.LazyProc
	proc_ZyanVectorPopBack      *windows.LazyProc
	proc_ZyanVectorClear      *windows.LazyProc
	proc_ZyanVectorFind      *windows.LazyProc
	proc_ZyanVectorFindEx      *windows.LazyProc
	proc_ZyanVectorBinarySearch      *windows.LazyProc
	proc_ZyanVectorBinarySearchEx      *windows.LazyProc
	proc_ZyanVectorResize      *windows.LazyProc
	proc_ZyanVectorResizeEx      *windows.LazyProc
	proc_ZyanVectorReserve      *windows.LazyProc
	proc_ZyanVectorShrinkToFit      *windows.LazyProc
	proc_ZyanVectorGetCapacity      *windows.LazyProc
	proc_ZyanVectorGetSize      *windows.LazyProc
	proc_ZyanStringInit      *windows.LazyProc
	proc_ZyanStringInitEx      *windows.LazyProc
	proc_ZyanStringInitCustomBuffer      *windows.LazyProc
	proc_ZyanStringDestroy      *windows.LazyProc
	proc_ZyanStringDuplicate      *windows.LazyProc
	proc_ZyanStringDuplicateEx      *windows.LazyProc
	proc_ZyanStringDuplicateCustomBuffer      *windows.LazyProc
	proc_ZyanStringConcat      *windows.LazyProc
	proc_ZyanStringConcatEx      *windows.LazyProc
	proc_ZyanStringConcatCustomBuffer      *windows.LazyProc
	proc_ZyanStringViewInsideView      *windows.LazyProc
	proc_ZyanStringViewInsideViewEx      *windows.LazyProc
	proc_ZyanStringViewInsideBuffer      *windows.LazyProc
	proc_ZyanStringViewInsideBufferEx      *windows.LazyProc
	proc_ZyanStringViewGetSize      *windows.LazyProc
	proc_ZyanStringViewGetData      *windows.LazyProc
	proc_ZyanStringGetChar      *windows.LazyProc
	proc_ZyanStringGetCharMutable      *windows.LazyProc
	proc_ZyanStringSetChar      *windows.LazyProc
	proc_ZyanStringInsert      *windows.LazyProc
	proc_ZyanStringInsertEx      *windows.LazyProc
	proc_ZyanStringAppend      *windows.LazyProc
	proc_ZyanStringAppendEx      *windows.LazyProc
	proc_ZyanStringDelete      *windows.LazyProc
	proc_ZyanStringTruncate      *windows.LazyProc
	proc_ZyanStringClear      *windows.LazyProc
	proc_ZyanStringLPos      *windows.LazyProc
	proc_ZyanStringLPosEx      *windows.LazyProc
	proc_ZyanStringLPosI      *windows.LazyProc
	proc_ZyanStringLPosIEx      *windows.LazyProc
	proc_ZyanStringRPos      *windows.LazyProc
	proc_ZyanStringRPosEx      *windows.LazyProc
	proc_ZyanStringRPosI      *windows.LazyProc
	proc_ZyanStringRPosIEx      *windows.LazyProc
	proc_ZyanStringCompare      *windows.LazyProc
	proc_ZyanStringCompareI      *windows.LazyProc
	proc_ZyanStringToLowerCase      *windows.LazyProc
	proc_ZyanStringToLowerCaseEx      *windows.LazyProc
	proc_ZyanStringToUpperCase      *windows.LazyProc
	proc_ZyanStringToUpperCaseEx      *windows.LazyProc
	proc_ZyanStringResize      *windows.LazyProc
	proc_ZyanStringReserve      *windows.LazyProc
	proc_ZyanStringShrinkToFit      *windows.LazyProc
	proc_ZyanStringGetCapacity      *windows.LazyProc
	proc_ZyanStringGetSize      *windows.LazyProc
	proc_ZyanStringGetData      *windows.LazyProc
	proc_ZydisFormatterTokenGetValue      *windows.LazyProc
	proc_ZydisFormatterTokenNext      *windows.LazyProc
	proc_ZydisFormatterBufferGetToken      *windows.LazyProc
	proc_ZydisFormatterBufferGetString      *windows.LazyProc
	proc_ZydisFormatterBufferAppend      *windows.LazyProc
	proc_ZydisFormatterBufferRemember      *windows.LazyProc
	proc_ZydisFormatterBufferRestore      *windows.LazyProc
	proc_ZydisFormatterInit      *windows.LazyProc
	proc_ZydisFormatterSetProperty      *windows.LazyProc
	proc_ZydisFormatterSetHook      *windows.LazyProc
	proc_ZydisFormatterFormatInstruction      *windows.LazyProc
	proc_ZydisFormatterFormatOperand      *windows.LazyProc
	proc_ZydisFormatterTokenizeInstruction      *windows.LazyProc
	proc_ZydisFormatterTokenizeOperand      *windows.LazyProc
	proc_ZydisGetInstructionSegments      *windows.LazyProc
	proc_ZydisDisassembleIntel      *windows.LazyProc
	proc_ZydisDisassembleATT      *windows.LazyProc
	proc_ZydisCalcAbsoluteAddress      *windows.LazyProc
	proc_ZydisCalcAbsoluteAddressEx      *windows.LazyProc
	proc_ZydisGetVersion      *windows.LazyProc
	proc_ZydisIsFeatureEnabled      *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "ZydisDll.dll"))
	proc_ZydisCategoryGetString = dll.NewProc("ZydisCategoryGetString")
	proc_ZydisISASetGetString = dll.NewProc("ZydisISASetGetString")
	proc_ZydisISAExtGetString = dll.NewProc("ZydisISAExtGetString")
	proc_ZydisMnemonicGetString = dll.NewProc("ZydisMnemonicGetString")
	proc_ZydisMnemonicGetStringWrapped = dll.NewProc("ZydisMnemonicGetStringWrapped")
	proc_ZydisRegisterEncode = dll.NewProc("ZydisRegisterEncode")
	proc_ZydisRegisterGetId = dll.NewProc("ZydisRegisterGetId")
	proc_ZydisRegisterGetClass = dll.NewProc("ZydisRegisterGetClass")
	proc_ZydisRegisterGetWidth = dll.NewProc("ZydisRegisterGetWidth")
	proc_ZydisRegisterGetLargestEnclosing = dll.NewProc("ZydisRegisterGetLargestEnclosing")
	proc_ZydisRegisterGetString = dll.NewProc("ZydisRegisterGetString")
	proc_ZydisRegisterGetStringWrapped = dll.NewProc("ZydisRegisterGetStringWrapped")
	proc_ZydisRegisterClassGetWidth = dll.NewProc("ZydisRegisterClassGetWidth")
	proc_ZydisDecoderInit = dll.NewProc("ZydisDecoderInit")
	proc_ZydisDecoderEnableMode = dll.NewProc("ZydisDecoderEnableMode")
	proc_ZydisDecoderDecodeFull = dll.NewProc("ZydisDecoderDecodeFull")
	proc_ZydisDecoderDecodeInstruction = dll.NewProc("ZydisDecoderDecodeInstruction")
	proc_ZydisDecoderDecodeOperands = dll.NewProc("ZydisDecoderDecodeOperands")
	proc_ZydisEncoderEncodeInstruction = dll.NewProc("ZydisEncoderEncodeInstruction")
	proc_ZydisEncoderEncodeInstructionAbsolute = dll.NewProc("ZydisEncoderEncodeInstructionAbsolute")
	proc_ZydisEncoderDecodedInstructionToEncoderRequest = dll.NewProc("ZydisEncoderDecodedInstructionToEncoderRequest")
	proc_ZydisEncoderNopFill = dll.NewProc("ZydisEncoderNopFill")
	proc_ZyanAllocatorInit = dll.NewProc("ZyanAllocatorInit")
	proc_ZyanAllocatorDefault = dll.NewProc("ZyanAllocatorDefault")
	proc_ZyanVectorInit = dll.NewProc("ZyanVectorInit")
	proc_ZyanVectorInitEx = dll.NewProc("ZyanVectorInitEx")
	proc_ZyanVectorInitCustomBuffer = dll.NewProc("ZyanVectorInitCustomBuffer")
	proc_ZyanVectorDestroy = dll.NewProc("ZyanVectorDestroy")
	proc_ZyanVectorDuplicate = dll.NewProc("ZyanVectorDuplicate")
	proc_ZyanVectorDuplicateEx = dll.NewProc("ZyanVectorDuplicateEx")
	proc_ZyanVectorDuplicateCustomBuffer = dll.NewProc("ZyanVectorDuplicateCustomBuffer")
	proc_ZyanVectorGet = dll.NewProc("ZyanVectorGet")
	proc_ZyanVectorGetMutable = dll.NewProc("ZyanVectorGetMutable")
	proc_ZyanVectorGetPointer = dll.NewProc("ZyanVectorGetPointer")
	proc_ZyanVectorGetPointerMutable = dll.NewProc("ZyanVectorGetPointerMutable")
	proc_ZyanVectorSet = dll.NewProc("ZyanVectorSet")
	proc_ZyanVectorPushBack = dll.NewProc("ZyanVectorPushBack")
	proc_ZyanVectorInsert = dll.NewProc("ZyanVectorInsert")
	proc_ZyanVectorInsertRange = dll.NewProc("ZyanVectorInsertRange")
	proc_ZyanVectorEmplace = dll.NewProc("ZyanVectorEmplace")
	proc_ZyanVectorEmplaceEx = dll.NewProc("ZyanVectorEmplaceEx")
	proc_ZyanVectorSwapElements = dll.NewProc("ZyanVectorSwapElements")
	proc_ZyanVectorDelete = dll.NewProc("ZyanVectorDelete")
	proc_ZyanVectorDeleteRange = dll.NewProc("ZyanVectorDeleteRange")
	proc_ZyanVectorPopBack = dll.NewProc("ZyanVectorPopBack")
	proc_ZyanVectorClear = dll.NewProc("ZyanVectorClear")
	proc_ZyanVectorFind = dll.NewProc("ZyanVectorFind")
	proc_ZyanVectorFindEx = dll.NewProc("ZyanVectorFindEx")
	proc_ZyanVectorBinarySearch = dll.NewProc("ZyanVectorBinarySearch")
	proc_ZyanVectorBinarySearchEx = dll.NewProc("ZyanVectorBinarySearchEx")
	proc_ZyanVectorResize = dll.NewProc("ZyanVectorResize")
	proc_ZyanVectorResizeEx = dll.NewProc("ZyanVectorResizeEx")
	proc_ZyanVectorReserve = dll.NewProc("ZyanVectorReserve")
	proc_ZyanVectorShrinkToFit = dll.NewProc("ZyanVectorShrinkToFit")
	proc_ZyanVectorGetCapacity = dll.NewProc("ZyanVectorGetCapacity")
	proc_ZyanVectorGetSize = dll.NewProc("ZyanVectorGetSize")
	proc_ZyanStringInit = dll.NewProc("ZyanStringInit")
	proc_ZyanStringInitEx = dll.NewProc("ZyanStringInitEx")
	proc_ZyanStringInitCustomBuffer = dll.NewProc("ZyanStringInitCustomBuffer")
	proc_ZyanStringDestroy = dll.NewProc("ZyanStringDestroy")
	proc_ZyanStringDuplicate = dll.NewProc("ZyanStringDuplicate")
	proc_ZyanStringDuplicateEx = dll.NewProc("ZyanStringDuplicateEx")
	proc_ZyanStringDuplicateCustomBuffer = dll.NewProc("ZyanStringDuplicateCustomBuffer")
	proc_ZyanStringConcat = dll.NewProc("ZyanStringConcat")
	proc_ZyanStringConcatEx = dll.NewProc("ZyanStringConcatEx")
	proc_ZyanStringConcatCustomBuffer = dll.NewProc("ZyanStringConcatCustomBuffer")
	proc_ZyanStringViewInsideView = dll.NewProc("ZyanStringViewInsideView")
	proc_ZyanStringViewInsideViewEx = dll.NewProc("ZyanStringViewInsideViewEx")
	proc_ZyanStringViewInsideBuffer = dll.NewProc("ZyanStringViewInsideBuffer")
	proc_ZyanStringViewInsideBufferEx = dll.NewProc("ZyanStringViewInsideBufferEx")
	proc_ZyanStringViewGetSize = dll.NewProc("ZyanStringViewGetSize")
	proc_ZyanStringViewGetData = dll.NewProc("ZyanStringViewGetData")
	proc_ZyanStringGetChar = dll.NewProc("ZyanStringGetChar")
	proc_ZyanStringGetCharMutable = dll.NewProc("ZyanStringGetCharMutable")
	proc_ZyanStringSetChar = dll.NewProc("ZyanStringSetChar")
	proc_ZyanStringInsert = dll.NewProc("ZyanStringInsert")
	proc_ZyanStringInsertEx = dll.NewProc("ZyanStringInsertEx")
	proc_ZyanStringAppend = dll.NewProc("ZyanStringAppend")
	proc_ZyanStringAppendEx = dll.NewProc("ZyanStringAppendEx")
	proc_ZyanStringDelete = dll.NewProc("ZyanStringDelete")
	proc_ZyanStringTruncate = dll.NewProc("ZyanStringTruncate")
	proc_ZyanStringClear = dll.NewProc("ZyanStringClear")
	proc_ZyanStringLPos = dll.NewProc("ZyanStringLPos")
	proc_ZyanStringLPosEx = dll.NewProc("ZyanStringLPosEx")
	proc_ZyanStringLPosI = dll.NewProc("ZyanStringLPosI")
	proc_ZyanStringLPosIEx = dll.NewProc("ZyanStringLPosIEx")
	proc_ZyanStringRPos = dll.NewProc("ZyanStringRPos")
	proc_ZyanStringRPosEx = dll.NewProc("ZyanStringRPosEx")
	proc_ZyanStringRPosI = dll.NewProc("ZyanStringRPosI")
	proc_ZyanStringRPosIEx = dll.NewProc("ZyanStringRPosIEx")
	proc_ZyanStringCompare = dll.NewProc("ZyanStringCompare")
	proc_ZyanStringCompareI = dll.NewProc("ZyanStringCompareI")
	proc_ZyanStringToLowerCase = dll.NewProc("ZyanStringToLowerCase")
	proc_ZyanStringToLowerCaseEx = dll.NewProc("ZyanStringToLowerCaseEx")
	proc_ZyanStringToUpperCase = dll.NewProc("ZyanStringToUpperCase")
	proc_ZyanStringToUpperCaseEx = dll.NewProc("ZyanStringToUpperCaseEx")
	proc_ZyanStringResize = dll.NewProc("ZyanStringResize")
	proc_ZyanStringReserve = dll.NewProc("ZyanStringReserve")
	proc_ZyanStringShrinkToFit = dll.NewProc("ZyanStringShrinkToFit")
	proc_ZyanStringGetCapacity = dll.NewProc("ZyanStringGetCapacity")
	proc_ZyanStringGetSize = dll.NewProc("ZyanStringGetSize")
	proc_ZyanStringGetData = dll.NewProc("ZyanStringGetData")
	proc_ZydisFormatterTokenGetValue = dll.NewProc("ZydisFormatterTokenGetValue")
	proc_ZydisFormatterTokenNext = dll.NewProc("ZydisFormatterTokenNext")
	proc_ZydisFormatterBufferGetToken = dll.NewProc("ZydisFormatterBufferGetToken")
	proc_ZydisFormatterBufferGetString = dll.NewProc("ZydisFormatterBufferGetString")
	proc_ZydisFormatterBufferAppend = dll.NewProc("ZydisFormatterBufferAppend")
	proc_ZydisFormatterBufferRemember = dll.NewProc("ZydisFormatterBufferRemember")
	proc_ZydisFormatterBufferRestore = dll.NewProc("ZydisFormatterBufferRestore")
	proc_ZydisFormatterInit = dll.NewProc("ZydisFormatterInit")
	proc_ZydisFormatterSetProperty = dll.NewProc("ZydisFormatterSetProperty")
	proc_ZydisFormatterSetHook = dll.NewProc("ZydisFormatterSetHook")
	proc_ZydisFormatterFormatInstruction = dll.NewProc("ZydisFormatterFormatInstruction")
	proc_ZydisFormatterFormatOperand = dll.NewProc("ZydisFormatterFormatOperand")
	proc_ZydisFormatterTokenizeInstruction = dll.NewProc("ZydisFormatterTokenizeInstruction")
	proc_ZydisFormatterTokenizeOperand = dll.NewProc("ZydisFormatterTokenizeOperand")
	proc_ZydisGetInstructionSegments = dll.NewProc("ZydisGetInstructionSegments")
	proc_ZydisDisassembleIntel = dll.NewProc("ZydisDisassembleIntel")
	proc_ZydisDisassembleATT = dll.NewProc("ZydisDisassembleATT")
	proc_ZydisCalcAbsoluteAddress = dll.NewProc("ZydisCalcAbsoluteAddress")
	proc_ZydisCalcAbsoluteAddressEx = dll.NewProc("ZydisCalcAbsoluteAddressEx")
	proc_ZydisGetVersion = dll.NewProc("ZydisGetVersion")
	proc_ZydisIsFeatureEnabled = dll.NewProc("ZydisIsFeatureEnabled")
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

func (z *Zydis) CategoryGetString(Category ZydisInstructionCategory) *int8 {
	r1, _, _ := proc_ZydisCategoryGetString.Call(uintptr(unsafe.Pointer(&Category)))
	return (*int8)(unsafe.Pointer(r1))
}

func (z *Zydis) ISASetGetString(Isa_set ZydisISASet) *int8 {
	r1, _, _ := proc_ZydisISASetGetString.Call(uintptr(unsafe.Pointer(&Isa_set)))
	return (*int8)(unsafe.Pointer(r1))
}

func (z *Zydis) ISAExtGetString(Isa_ext ZydisISAExt) *int8 {
	r1, _, _ := proc_ZydisISAExtGetString.Call(uintptr(unsafe.Pointer(&Isa_ext)))
	return (*int8)(unsafe.Pointer(r1))
}

func (z *Zydis) MnemonicGetString(Mnemonic ZydisMnemonic) *int8 {
	r1, _, _ := proc_ZydisMnemonicGetString.Call(uintptr(unsafe.Pointer(&Mnemonic)))
	return (*int8)(unsafe.Pointer(r1))
}

func (z *Zydis) MnemonicGetStringWrapped(Mnemonic ZydisMnemonic) *ZydisShortString {
	r1, _, _ := proc_ZydisMnemonicGetStringWrapped.Call(uintptr(unsafe.Pointer(&Mnemonic)))
	return (*ZydisShortString)(unsafe.Pointer(r1))
}

func (z *Zydis) RegisterEncode(Register_class ZydisRegisterClass, Id uint8) ZydisRegister {
	r1, _, _ := proc_ZydisRegisterEncode.Call(uintptr(unsafe.Pointer(&Register_class)), uintptr(Id))
	return ZydisRegister(r1)
}

func (z *Zydis) RegisterGetId(Reg ZydisRegister) int8 {
	r1, _, _ := proc_ZydisRegisterGetId.Call(uintptr(unsafe.Pointer(&Reg)))
	return int8(r1)
}

func (z *Zydis) RegisterGetClass(Reg ZydisRegister) ZydisRegisterClass {
	r1, _, _ := proc_ZydisRegisterGetClass.Call(uintptr(unsafe.Pointer(&Reg)))
	return ZydisRegisterClass(r1)
}

func (z *Zydis) RegisterGetWidth(Mode ZydisMachineMode, Reg ZydisRegister) ZydisRegisterWidth {
	r1, _, _ := proc_ZydisRegisterGetWidth.Call(uintptr(unsafe.Pointer(&Mode)), uintptr(unsafe.Pointer(&Reg)))
	return ZydisRegisterWidth(r1)
}

func (z *Zydis) RegisterGetLargestEnclosing(Mode ZydisMachineMode, Reg ZydisRegister) ZydisRegister {
	r1, _, _ := proc_ZydisRegisterGetLargestEnclosing.Call(uintptr(unsafe.Pointer(&Mode)), uintptr(unsafe.Pointer(&Reg)))
	return ZydisRegister(r1)
}

func (z *Zydis) RegisterGetString(Reg ZydisRegister) *int8 {
	r1, _, _ := proc_ZydisRegisterGetString.Call(uintptr(unsafe.Pointer(&Reg)))
	return (*int8)(unsafe.Pointer(r1))
}

func (z *Zydis) RegisterGetStringWrapped(Reg ZydisRegister) *ZydisShortString {
	r1, _, _ := proc_ZydisRegisterGetStringWrapped.Call(uintptr(unsafe.Pointer(&Reg)))
	return (*ZydisShortString)(unsafe.Pointer(r1))
}

func (z *Zydis) RegisterClassGetWidth(Mode ZydisMachineMode, Register_class ZydisRegisterClass) ZydisRegisterWidth {
	r1, _, _ := proc_ZydisRegisterClassGetWidth.Call(uintptr(unsafe.Pointer(&Mode)), uintptr(unsafe.Pointer(&Register_class)))
	return ZydisRegisterWidth(r1)
}

func (z *Zydis) DecoderInit(Decoder *ZydisDecoder, Machine_mode ZydisMachineMode, Stack_width ZydisStackWidth) uint32 {
	r1, _, _ := proc_ZydisDecoderInit.Call(uintptr(unsafe.Pointer(Decoder)), uintptr(unsafe.Pointer(&Machine_mode)), uintptr(unsafe.Pointer(&Stack_width)))
	return uint32(r1)
}

func (z *Zydis) DecoderEnableMode(Decoder *ZydisDecoder, Mode ZydisDecoderMode, Enabled uint8) uint32 {
	r1, _, _ := proc_ZydisDecoderEnableMode.Call(uintptr(unsafe.Pointer(Decoder)), uintptr(unsafe.Pointer(&Mode)), uintptr(Enabled))
	return uint32(r1)
}

func (z *Zydis) DecoderDecodeFull(Decoder *ZydisDecoder, Buffer unsafe.Pointer, Length uintptr, Instruction *ZydisDecodedInstruction, Operands *ZydisDecodedOperand) uint32 {
	r1, _, _ := proc_ZydisDecoderDecodeFull.Call(uintptr(unsafe.Pointer(Decoder)), uintptr(Buffer), Length, uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operands)))
	return uint32(r1)
}

func (z *Zydis) DecoderDecodeInstruction(Decoder *ZydisDecoder, Context *ZydisDecoderContext, Buffer unsafe.Pointer, Length uintptr, Instruction *ZydisDecodedInstruction) uint32 {
	r1, _, _ := proc_ZydisDecoderDecodeInstruction.Call(uintptr(unsafe.Pointer(Decoder)), uintptr(unsafe.Pointer(Context)), uintptr(Buffer), Length, uintptr(unsafe.Pointer(Instruction)))
	return uint32(r1)
}

func (z *Zydis) DecoderDecodeOperands(Decoder *ZydisDecoder, Context *ZydisDecoderContext, Instruction *ZydisDecodedInstruction, Operands *ZydisDecodedOperand, Operand_count uint8) uint32 {
	r1, _, _ := proc_ZydisDecoderDecodeOperands.Call(uintptr(unsafe.Pointer(Decoder)), uintptr(unsafe.Pointer(Context)), uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operands)), uintptr(Operand_count))
	return uint32(r1)
}

func (z *Zydis) EncoderEncodeInstruction(Request *ZydisEncoderRequest, Buffer unsafe.Pointer, Length *uintptr) uint32 {
	r1, _, _ := proc_ZydisEncoderEncodeInstruction.Call(uintptr(unsafe.Pointer(Request)), uintptr(Buffer), uintptr(unsafe.Pointer(Length)))
	return uint32(r1)
}

func (z *Zydis) EncoderEncodeInstructionAbsolute(Request *ZydisEncoderRequest, Buffer unsafe.Pointer, Length *uintptr, Runtime_address uint64) uint32 {
	r1, _, _ := proc_ZydisEncoderEncodeInstructionAbsolute.Call(uintptr(unsafe.Pointer(Request)), uintptr(Buffer), uintptr(unsafe.Pointer(Length)), *(*uintptr)(unsafe.Pointer(&Runtime_address)))
	return uint32(r1)
}

func (z *Zydis) EncoderDecodedInstructionToEncoderRequest(Instruction *ZydisDecodedInstruction, Operands *ZydisDecodedOperand, Operand_count_visible uint8, Request *ZydisEncoderRequest) uint32 {
	r1, _, _ := proc_ZydisEncoderDecodedInstructionToEncoderRequest.Call(uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operands)), uintptr(Operand_count_visible), uintptr(unsafe.Pointer(Request)))
	return uint32(r1)
}

func (z *Zydis) EncoderNopFill(Buffer unsafe.Pointer, Length uintptr) uint32 {
	r1, _, _ := proc_ZydisEncoderNopFill.Call(uintptr(Buffer), Length)
	return uint32(r1)
}

func (z *Zydis) ZyanAllocatorInit(Allocator *ZyanAllocator, Allocate ZyanAllocatorAllocate, Reallocate ZyanAllocatorAllocate, Deallocate ZyanAllocatorDeallocate) uint32 {
	r1, _, _ := proc_ZyanAllocatorInit.Call(uintptr(unsafe.Pointer(Allocator)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Allocate))), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Reallocate))), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Deallocate))))
	return uint32(r1)
}

func (z *Zydis) ZyanAllocatorDefault() *ZyanAllocator {
	r1, _, _ := proc_ZyanAllocatorDefault.Call()
	return (*ZyanAllocator)(unsafe.Pointer(r1))
}

func (z *Zydis) ZyanVectorInit(Vector *ZyanVector, Element_size uintptr, Capacity uintptr, Destructor ZyanMemberProcedure) uint32 {
	r1, _, _ := proc_ZyanVectorInit.Call(uintptr(unsafe.Pointer(Vector)), Element_size, Capacity, uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Destructor))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorInitEx(Vector *ZyanVector, Element_size uintptr, Capacity uintptr, Destructor ZyanMemberProcedure, Allocator *ZyanAllocator, Growth_factor uint8, Shrink_threshold uint8) uint32 {
	r1, _, _ := proc_ZyanVectorInitEx.Call(uintptr(unsafe.Pointer(Vector)), Element_size, Capacity, uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Destructor))), uintptr(unsafe.Pointer(Allocator)), uintptr(Growth_factor), uintptr(Shrink_threshold))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorInitCustomBuffer(Vector *ZyanVector, Element_size uintptr, Buffer unsafe.Pointer, Capacity uintptr, Destructor ZyanMemberProcedure) uint32 {
	r1, _, _ := proc_ZyanVectorInitCustomBuffer.Call(uintptr(unsafe.Pointer(Vector)), Element_size, uintptr(Buffer), Capacity, uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Destructor))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDestroy(Vector *ZyanVector) uint32 {
	r1, _, _ := proc_ZyanVectorDestroy.Call(uintptr(unsafe.Pointer(Vector)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDuplicate(Destination *ZyanVector, Source *ZyanVector, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorDuplicate.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDuplicateEx(Destination *ZyanVector, Source *ZyanVector, Capacity uintptr, Allocator *ZyanAllocator, Growth_factor uint8, Shrink_threshold uint8) uint32 {
	r1, _, _ := proc_ZyanVectorDuplicateEx.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), Capacity, uintptr(unsafe.Pointer(Allocator)), uintptr(Growth_factor), uintptr(Shrink_threshold))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDuplicateCustomBuffer(Destination *ZyanVector, Source *ZyanVector, Buffer unsafe.Pointer, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorDuplicateCustomBuffer.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), uintptr(Buffer), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorGet(Vector *ZyanVector, Index uintptr) unsafe.Pointer {
	r1, _, _ := proc_ZyanVectorGet.Call(uintptr(unsafe.Pointer(Vector)), Index)
	return unsafe.Pointer(r1)
}

func (z *Zydis) ZyanVectorGetMutable(Vector *ZyanVector, Index uintptr) unsafe.Pointer {
	r1, _, _ := proc_ZyanVectorGetMutable.Call(uintptr(unsafe.Pointer(Vector)), Index)
	return unsafe.Pointer(r1)
}

func (z *Zydis) ZyanVectorGetPointer(Vector *ZyanVector, Index uintptr, Value *unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorGetPointer.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorGetPointerMutable(Vector *ZyanVector, Index uintptr, Value *unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorGetPointerMutable.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorSet(Vector *ZyanVector, Index uintptr, Value unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorSet.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(Value))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorPushBack(Vector *ZyanVector, Element unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorPushBack.Call(uintptr(unsafe.Pointer(Vector)), uintptr(Element))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorInsert(Vector *ZyanVector, Index uintptr, Element unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorInsert.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(Element))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorInsertRange(Vector *ZyanVector, Index uintptr, Elements unsafe.Pointer, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorInsertRange.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(Elements), Count)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorEmplace(Vector *ZyanVector, Element *unsafe.Pointer, Constructor ZyanMemberFunction) uint32 {
	r1, _, _ := proc_ZyanVectorEmplace.Call(uintptr(unsafe.Pointer(Vector)), uintptr(unsafe.Pointer(Element)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Constructor))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorEmplaceEx(Vector *ZyanVector, Index uintptr, Element *unsafe.Pointer, Constructor ZyanMemberFunction) uint32 {
	r1, _, _ := proc_ZyanVectorEmplaceEx.Call(uintptr(unsafe.Pointer(Vector)), Index, uintptr(unsafe.Pointer(Element)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Constructor))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorSwapElements(Vector *ZyanVector, Index_first uintptr, Index_second uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorSwapElements.Call(uintptr(unsafe.Pointer(Vector)), Index_first, Index_second)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDelete(Vector *ZyanVector, Index uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorDelete.Call(uintptr(unsafe.Pointer(Vector)), Index)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorDeleteRange(Vector *ZyanVector, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorDeleteRange.Call(uintptr(unsafe.Pointer(Vector)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorPopBack(Vector *ZyanVector) uint32 {
	r1, _, _ := proc_ZyanVectorPopBack.Call(uintptr(unsafe.Pointer(Vector)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorClear(Vector *ZyanVector) uint32 {
	r1, _, _ := proc_ZyanVectorClear.Call(uintptr(unsafe.Pointer(Vector)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorFind(Vector *ZyanVector, Element unsafe.Pointer, Found_index *uintptr, Comparison ZyanEqualityComparison) uint32 {
	r1, _, _ := proc_ZyanVectorFind.Call(uintptr(unsafe.Pointer(Vector)), uintptr(Element), uintptr(unsafe.Pointer(Found_index)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Comparison))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorFindEx(Vector *ZyanVector, Element unsafe.Pointer, Found_index *uintptr, Comparison ZyanEqualityComparison, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorFindEx.Call(uintptr(unsafe.Pointer(Vector)), uintptr(Element), uintptr(unsafe.Pointer(Found_index)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Comparison))), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorBinarySearch(Vector *ZyanVector, Element unsafe.Pointer, Found_index *uintptr, Comparison ZyanComparison) uint32 {
	r1, _, _ := proc_ZyanVectorBinarySearch.Call(uintptr(unsafe.Pointer(Vector)), uintptr(Element), uintptr(unsafe.Pointer(Found_index)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Comparison))))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorBinarySearchEx(Vector *ZyanVector, Element unsafe.Pointer, Found_index *uintptr, Comparison ZyanComparison, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorBinarySearchEx.Call(uintptr(unsafe.Pointer(Vector)), uintptr(Element), uintptr(unsafe.Pointer(Found_index)), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&Comparison))), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorResize(Vector *ZyanVector, Size uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorResize.Call(uintptr(unsafe.Pointer(Vector)), Size)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorResizeEx(Vector *ZyanVector, Size uintptr, Initializer unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZyanVectorResizeEx.Call(uintptr(unsafe.Pointer(Vector)), Size, uintptr(Initializer))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorReserve(Vector *ZyanVector, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorReserve.Call(uintptr(unsafe.Pointer(Vector)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanVectorShrinkToFit(Vector *ZyanVector) uint32 {
	r1, _, _ := proc_ZyanVectorShrinkToFit.Call(uintptr(unsafe.Pointer(Vector)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorGetCapacity(Vector *ZyanVector, Capacity *uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorGetCapacity.Call(uintptr(unsafe.Pointer(Vector)), uintptr(unsafe.Pointer(Capacity)))
	return uint32(r1)
}

func (z *Zydis) ZyanVectorGetSize(Vector *ZyanVector, Size *uintptr) uint32 {
	r1, _, _ := proc_ZyanVectorGetSize.Call(uintptr(unsafe.Pointer(Vector)), uintptr(unsafe.Pointer(Size)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringInit(String *ZyanString, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringInit.Call(uintptr(unsafe.Pointer(String)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringInitEx(String *ZyanString, Capacity uintptr, Allocator *ZyanAllocator, Growth_factor uint8, Shrink_threshold uint8) uint32 {
	r1, _, _ := proc_ZyanStringInitEx.Call(uintptr(unsafe.Pointer(String)), Capacity, uintptr(unsafe.Pointer(Allocator)), uintptr(Growth_factor), uintptr(Shrink_threshold))
	return uint32(r1)
}

func (z *Zydis) ZyanStringInitCustomBuffer(String *ZyanString, Buffer *int8, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringInitCustomBuffer.Call(uintptr(unsafe.Pointer(String)), uintptr(unsafe.Pointer(Buffer)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringDestroy(String *ZyanString) uint32 {
	r1, _, _ := proc_ZyanStringDestroy.Call(uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringDuplicate(Destination *ZyanString, Source *ZyanStringView, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringDuplicate.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringDuplicateEx(Destination *ZyanString, Source *ZyanStringView, Capacity uintptr, Allocator *ZyanAllocator, Growth_factor uint8, Shrink_threshold uint8) uint32 {
	r1, _, _ := proc_ZyanStringDuplicateEx.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), Capacity, uintptr(unsafe.Pointer(Allocator)), uintptr(Growth_factor), uintptr(Shrink_threshold))
	return uint32(r1)
}

func (z *Zydis) ZyanStringDuplicateCustomBuffer(Destination *ZyanString, Source *ZyanStringView, Buffer *int8, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringDuplicateCustomBuffer.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), uintptr(unsafe.Pointer(Buffer)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringConcat(Destination *ZyanString, S1 *ZyanStringView, S2 *ZyanStringView, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringConcat.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringConcatEx(Destination *ZyanString, S1 *ZyanStringView, S2 *ZyanStringView, Capacity uintptr, Allocator *ZyanAllocator, Growth_factor uint8, Shrink_threshold uint8) uint32 {
	r1, _, _ := proc_ZyanStringConcatEx.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), Capacity, uintptr(unsafe.Pointer(Allocator)), uintptr(Growth_factor), uintptr(Shrink_threshold))
	return uint32(r1)
}

func (z *Zydis) ZyanStringConcatCustomBuffer(Destination *ZyanString, S1 *ZyanStringView, S2 *ZyanStringView, Buffer *int8, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringConcatCustomBuffer.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), uintptr(unsafe.Pointer(Buffer)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewInsideView(View *ZyanStringView, Source *ZyanStringView) uint32 {
	r1, _, _ := proc_ZyanStringViewInsideView.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(Source)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewInsideViewEx(View *ZyanStringView, Source *ZyanStringView, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringViewInsideViewEx.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(Source)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewInsideBuffer(View *ZyanStringView, String *int8) uint32 {
	r1, _, _ := proc_ZyanStringViewInsideBuffer.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewInsideBufferEx(View *ZyanStringView, Buffer *int8, Length uintptr) uint32 {
	r1, _, _ := proc_ZyanStringViewInsideBufferEx.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(Buffer)), Length)
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewGetSize(View *ZyanStringView, Size *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringViewGetSize.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(Size)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringViewGetData(View *ZyanStringView, Buffer **int8) uint32 {
	r1, _, _ := proc_ZyanStringViewGetData.Call(uintptr(unsafe.Pointer(View)), uintptr(unsafe.Pointer(Buffer)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringGetChar(String *ZyanStringView, Index uintptr, Value *int8) uint32 {
	r1, _, _ := proc_ZyanStringGetChar.Call(uintptr(unsafe.Pointer(String)), Index, uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringGetCharMutable(String *ZyanString, Index uintptr, Value **int8) uint32 {
	r1, _, _ := proc_ZyanStringGetCharMutable.Call(uintptr(unsafe.Pointer(String)), Index, uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringSetChar(String *ZyanString, Index uintptr, Value int8) uint32 {
	r1, _, _ := proc_ZyanStringSetChar.Call(uintptr(unsafe.Pointer(String)), Index, uintptr(Value))
	return uint32(r1)
}

func (z *Zydis) ZyanStringInsert(Destination *ZyanString, Index uintptr, Source *ZyanStringView) uint32 {
	r1, _, _ := proc_ZyanStringInsert.Call(uintptr(unsafe.Pointer(Destination)), Index, uintptr(unsafe.Pointer(Source)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringInsertEx(Destination *ZyanString, Destination_index uintptr, Source *ZyanStringView, Source_index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringInsertEx.Call(uintptr(unsafe.Pointer(Destination)), Destination_index, uintptr(unsafe.Pointer(Source)), Source_index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringAppend(Destination *ZyanString, Source *ZyanStringView) uint32 {
	r1, _, _ := proc_ZyanStringAppend.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringAppendEx(Destination *ZyanString, Source *ZyanStringView, Source_index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringAppendEx.Call(uintptr(unsafe.Pointer(Destination)), uintptr(unsafe.Pointer(Source)), Source_index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringDelete(String *ZyanString, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringDelete.Call(uintptr(unsafe.Pointer(String)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringTruncate(String *ZyanString, Index uintptr) uint32 {
	r1, _, _ := proc_ZyanStringTruncate.Call(uintptr(unsafe.Pointer(String)), Index)
	return uint32(r1)
}

func (z *Zydis) ZyanStringClear(String *ZyanString) uint32 {
	r1, _, _ := proc_ZyanStringClear.Call(uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringLPos(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringLPos.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringLPosEx(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringLPosEx.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringLPosI(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringLPosI.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringLPosIEx(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringLPosIEx.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringRPos(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringRPos.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringRPosEx(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringRPosEx.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringRPosI(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringRPosI.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringRPosIEx(Haystack *ZyanStringView, Needle *ZyanStringView, Found_index *uintptr, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringRPosIEx.Call(uintptr(unsafe.Pointer(Haystack)), uintptr(unsafe.Pointer(Needle)), uintptr(unsafe.Pointer(Found_index)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringCompare(S1 *ZyanStringView, S2 *ZyanStringView, Result *int32) uint32 {
	r1, _, _ := proc_ZyanStringCompare.Call(uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), uintptr(unsafe.Pointer(Result)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringCompareI(S1 *ZyanStringView, S2 *ZyanStringView, Result *int32) uint32 {
	r1, _, _ := proc_ZyanStringCompareI.Call(uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), uintptr(unsafe.Pointer(Result)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringToLowerCase(String *ZyanString) uint32 {
	r1, _, _ := proc_ZyanStringToLowerCase.Call(uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringToLowerCaseEx(String *ZyanString, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringToLowerCaseEx.Call(uintptr(unsafe.Pointer(String)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringToUpperCase(String *ZyanString) uint32 {
	r1, _, _ := proc_ZyanStringToUpperCase.Call(uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringToUpperCaseEx(String *ZyanString, Index uintptr, Count uintptr) uint32 {
	r1, _, _ := proc_ZyanStringToUpperCaseEx.Call(uintptr(unsafe.Pointer(String)), Index, Count)
	return uint32(r1)
}

func (z *Zydis) ZyanStringResize(String *ZyanString, Size uintptr) uint32 {
	r1, _, _ := proc_ZyanStringResize.Call(uintptr(unsafe.Pointer(String)), Size)
	return uint32(r1)
}

func (z *Zydis) ZyanStringReserve(String *ZyanString, Capacity uintptr) uint32 {
	r1, _, _ := proc_ZyanStringReserve.Call(uintptr(unsafe.Pointer(String)), Capacity)
	return uint32(r1)
}

func (z *Zydis) ZyanStringShrinkToFit(String *ZyanString) uint32 {
	r1, _, _ := proc_ZyanStringShrinkToFit.Call(uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringGetCapacity(String *ZyanString, Capacity *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringGetCapacity.Call(uintptr(unsafe.Pointer(String)), uintptr(unsafe.Pointer(Capacity)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringGetSize(String *ZyanString, Size *uintptr) uint32 {
	r1, _, _ := proc_ZyanStringGetSize.Call(uintptr(unsafe.Pointer(String)), uintptr(unsafe.Pointer(Size)))
	return uint32(r1)
}

func (z *Zydis) ZyanStringGetData(String *ZyanString, Value **int8) uint32 {
	r1, _, _ := proc_ZyanStringGetData.Call(uintptr(unsafe.Pointer(String)), uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) FormatterTokenGetValue(Token *ZydisFormatterToken, Type *ZydisTokenType, Value *ZyanConstCharPointer) uint32 {
	r1, _, _ := proc_ZydisFormatterTokenGetValue.Call(uintptr(unsafe.Pointer(Token)), uintptr(unsafe.Pointer(Type)), uintptr(unsafe.Pointer(Value)))
	return uint32(r1)
}

func (z *Zydis) FormatterTokenNext(Token **ZydisFormatterTokenConst) uint32 {
	r1, _, _ := proc_ZydisFormatterTokenNext.Call(uintptr(unsafe.Pointer(Token)))
	return uint32(r1)
}

func (z *Zydis) FormatterBufferGetToken(Buffer *ZydisFormatterBuffer, Token **ZydisFormatterTokenConst) uint32 {
	r1, _, _ := proc_ZydisFormatterBufferGetToken.Call(uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(Token)))
	return uint32(r1)
}

func (z *Zydis) FormatterBufferGetString(Buffer *ZydisFormatterBuffer, String **ZyanString) uint32 {
	r1, _, _ := proc_ZydisFormatterBufferGetString.Call(uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(String)))
	return uint32(r1)
}

func (z *Zydis) FormatterBufferAppend(Buffer *ZydisFormatterBuffer, Type ZydisTokenType) uint32 {
	r1, _, _ := proc_ZydisFormatterBufferAppend.Call(uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(&Type)))
	return uint32(r1)
}

func (z *Zydis) FormatterBufferRemember(Buffer *ZydisFormatterBuffer, State *uintptr) uint32 {
	r1, _, _ := proc_ZydisFormatterBufferRemember.Call(uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(State)))
	return uint32(r1)
}

func (z *Zydis) FormatterBufferRestore(Buffer *ZydisFormatterBuffer, State uintptr) uint32 {
	r1, _, _ := proc_ZydisFormatterBufferRestore.Call(uintptr(unsafe.Pointer(Buffer)), State)
	return uint32(r1)
}

func (z *Zydis) FormatterInit(Formatter *ZydisFormatter, Style ZydisFormatterStyle) uint32 {
	r1, _, _ := proc_ZydisFormatterInit.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(&Style)))
	return uint32(r1)
}

func (z *Zydis) FormatterSetProperty(Formatter *ZydisFormatter, Property ZydisFormatterProperty, Value uintptr) uint32 {
	r1, _, _ := proc_ZydisFormatterSetProperty.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(&Property)), Value)
	return uint32(r1)
}

func (z *Zydis) FormatterSetHook(Formatter *ZydisFormatter, Type ZydisFormatterFunction, Callback *unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZydisFormatterSetHook.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(&Type)), uintptr(unsafe.Pointer(Callback)))
	return uint32(r1)
}

func (z *Zydis) FormatterFormatInstruction(Formatter *ZydisFormatter, Instruction *ZydisDecodedInstruction, Operands *ZydisDecodedOperand, Operand_count uint8, Buffer *int8, Length uintptr, Runtime_address uint64, User_data unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZydisFormatterFormatInstruction.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operands)), uintptr(Operand_count), uintptr(unsafe.Pointer(Buffer)), Length, *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(User_data))
	return uint32(r1)
}

func (z *Zydis) FormatterFormatOperand(Formatter *ZydisFormatter, Instruction *ZydisDecodedInstruction, Operand *ZydisDecodedOperand, Buffer *int8, Length uintptr, Runtime_address uint64, User_data unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZydisFormatterFormatOperand.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operand)), uintptr(unsafe.Pointer(Buffer)), Length, *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(User_data))
	return uint32(r1)
}

func (z *Zydis) FormatterTokenizeInstruction(Formatter *ZydisFormatter, Instruction *ZydisDecodedInstruction, Operands *ZydisDecodedOperand, Operand_count uint8, Buffer unsafe.Pointer, Length uintptr, Runtime_address uint64, Token **ZydisFormatterTokenConst, User_data unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZydisFormatterTokenizeInstruction.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operands)), uintptr(Operand_count), uintptr(Buffer), Length, *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(unsafe.Pointer(Token)), uintptr(User_data))
	return uint32(r1)
}

func (z *Zydis) FormatterTokenizeOperand(Formatter *ZydisFormatter, Instruction *ZydisDecodedInstruction, Operand *ZydisDecodedOperand, Buffer unsafe.Pointer, Length uintptr, Runtime_address uint64, Token **ZydisFormatterTokenConst, User_data unsafe.Pointer) uint32 {
	r1, _, _ := proc_ZydisFormatterTokenizeOperand.Call(uintptr(unsafe.Pointer(Formatter)), uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operand)), uintptr(Buffer), Length, *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(unsafe.Pointer(Token)), uintptr(User_data))
	return uint32(r1)
}

func (z *Zydis) GetInstructionSegments(Instruction *ZydisDecodedInstruction, Segments *ZydisInstructionSegments) uint32 {
	r1, _, _ := proc_ZydisGetInstructionSegments.Call(uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Segments)))
	return uint32(r1)
}

func (z *Zydis) DisassembleIntel(Machine_mode ZydisMachineMode, Runtime_address uint64, Buffer unsafe.Pointer, Length uintptr, Instruction *ZydisDisassembledInstruction) uint32 {
	r1, _, _ := proc_ZydisDisassembleIntel.Call(uintptr(unsafe.Pointer(&Machine_mode)), *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(Buffer), Length, uintptr(unsafe.Pointer(Instruction)))
	return uint32(r1)
}

func (z *Zydis) DisassembleATT(Machine_mode ZydisMachineMode, Runtime_address uint64, Buffer unsafe.Pointer, Length uintptr, Instruction *ZydisDisassembledInstruction) uint32 {
	r1, _, _ := proc_ZydisDisassembleATT.Call(uintptr(unsafe.Pointer(&Machine_mode)), *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(Buffer), Length, uintptr(unsafe.Pointer(Instruction)))
	return uint32(r1)
}

func (z *Zydis) CalcAbsoluteAddress(Instruction *ZydisDecodedInstruction, Operand *ZydisDecodedOperand, Runtime_address uint64, Result_address *uint64) uint32 {
	r1, _, _ := proc_ZydisCalcAbsoluteAddress.Call(uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operand)), *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(unsafe.Pointer(Result_address)))
	return uint32(r1)
}

func (z *Zydis) CalcAbsoluteAddressEx(Instruction *ZydisDecodedInstruction, Operand *ZydisDecodedOperand, Runtime_address uint64, Register_context *ZydisRegisterContext, Result_address *uint64) uint32 {
	r1, _, _ := proc_ZydisCalcAbsoluteAddressEx.Call(uintptr(unsafe.Pointer(Instruction)), uintptr(unsafe.Pointer(Operand)), *(*uintptr)(unsafe.Pointer(&Runtime_address)), uintptr(unsafe.Pointer(Register_context)), uintptr(unsafe.Pointer(Result_address)))
	return uint32(r1)
}

func (z *Zydis) GetVersion() uint64 {
	r1, _, _ := proc_ZydisGetVersion.Call()
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (z *Zydis) IsFeatureEnabled(Feature ZydisFeature) uint32 {
	r1, _, _ := proc_ZydisIsFeatureEnabled.Call(uintptr(unsafe.Pointer(&Feature)))
	return uint32(r1)
}

