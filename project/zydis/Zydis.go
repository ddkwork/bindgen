package zydis

import (
	"fmt"
	"unsafe"
)

// Source: Zydis.h
type (
	ZyanU8 = uint8
	ZyanU16 = uint16
	ZyanU32 = uint32
	ZyanU64 = uint64
	ZyanI8 = int8
	ZyanI16 = int16
	ZyanI32 = int32
	ZyanI64 = int64
	ZyanUSize = uint64
	ZyanISize = int64
	ZyanUPointer = uint64
	ZyanIPointer = int64
	ZyanVoidPointer = *uintptr
	ZyanConstVoidPointer = *uintptr
	ZyanBool = uint8
	ZyanTernary = int8
	ZyanCharPointer = *int8
	ZyanConstCharPointer = *int8
	ZydisInstructionCategory = ZydisInstructionCategory_
	ZydisISASet = ZydisISASet_
	ZydisISAExt = ZydisISAExt_
	ZydisShortString = ZydisShortString_
	ZydisMnemonic = ZydisMnemonic_
	ZydisMachineMode = ZydisMachineMode_
	ZydisStackWidth = ZydisStackWidth_
	ZydisElementType = ZydisElementType_
	ZydisElementSize = uint16
	ZydisOperandType = ZydisOperandType_
	ZydisOperandEncoding = ZydisOperandEncoding_
	ZydisOperandVisibility = ZydisOperandVisibility_
	ZydisOperandAction = ZydisOperandAction_
	ZydisOperandActions = uint8
	ZydisInstructionEncoding = ZydisInstructionEncoding_
	ZydisOpcodeMap = ZydisOpcodeMap_
	ZydisInstructionAttributes = uint64
	ZydisRegister = ZydisRegister_
	ZydisRegisterKind = ZydisRegisterKind_
	ZydisRegisterClass = ZydisRegisterClass_
	ZydisRegisterWidth = uint16
	ZydisRegisterContext = ZydisRegisterContext_
	ZydisOperandAttributes = uint8
	ZydisMemoryOperandType = ZydisMemoryOperandType_
	ZydisDecodedOperand = ZydisDecodedOperand_
	ZydisAccessedFlagsMask = uint32
	ZydisAccessedFlags = ZydisAccessedFlags_
	ZydisBranchType = ZydisBranchType_
	ZydisExceptionClass = ZydisExceptionClass_
	ZydisMaskMode = ZydisMaskMode_
	ZydisBroadcastMode = ZydisBroadcastMode_
	ZydisRoundingMode = ZydisRoundingMode_
	ZydisSwizzleMode = ZydisSwizzleMode_
	ZydisConversionMode = ZydisConversionMode_
	ZydisDefaultFlagsValue = uint8
	ZydisSourceConditionCode = ZydisSourceConditionCode_
	ZydisPrefixType = ZydisPrefixType_
	ZydisDecoderContext = ZydisDecoderContext_
	ZyanStatus = uint32
	ZydisDecoderMode = ZydisDecoderMode_
	ZydisDecoder = ZydisDecoder_
	ZydisEncodableEncoding = ZydisEncodableEncoding_
	ZydisBranchWidth = ZydisBranchWidth_
	ZydisAddressSizeHint = ZydisAddressSizeHint_
	ZydisOperandSizeHint = ZydisOperandSizeHint_
	ZydisEncoderOperand = ZydisEncoderOperand_
	ZydisEncoderRequest = ZydisEncoderRequest_
	ZyanAllocator = ZyanAllocator_
	ZyanStringFlags = uint8
	ZydisTokenType = uint8
	ZydisFormatterToken = ZydisFormatterToken_
	ZydisFormatterTokenConst = ZydisFormatterToken_
	ZydisFormatterBuffer = ZydisFormatterBuffer_
	ZydisFormatterStyle = ZydisFormatterStyle_
	ZydisFormatterProperty = ZydisFormatterProperty_
	ZydisNumericBase = ZydisNumericBase_
	ZydisSignedness = ZydisSignedness_
	ZydisPadding = ZydisPadding_
	ZydisFormatterFunction = ZydisFormatterFunction_
	ZydisDecorator = ZydisDecorator_
	ZydisFormatter = ZydisFormatter_
	ZydisFormatterContext = ZydisFormatterContext_
	ZydisInstructionSegment = ZydisInstructionSegment_
	ZydisInstructionSegments = ZydisInstructionSegments_
	ZydisDisassembledInstruction = ZydisDisassembledInstruction_
	ZydisFeature = ZydisFeature_
)

type ZyanAllocatorAllocate func(*ZyanAllocator_, *unsafe.Pointer, uintptr, uintptr) uint32

type ZyanAllocatorDeallocate func(*ZyanAllocator_, unsafe.Pointer, uintptr, uintptr) uint32

type ZyanEqualityComparison func(unsafe.Pointer, unsafe.Pointer) uint8

type ZyanComparison func(unsafe.Pointer, unsafe.Pointer) int32

type ZyanMemberProcedure func(unsafe.Pointer)

type ZyanConstMemberProcedure func(unsafe.Pointer)

type ZyanMemberFunction func(unsafe.Pointer) uint32

type ZyanConstMemberFunction func(unsafe.Pointer) uint32

type ZydisFormatterFunc func(*ZydisFormatter, *ZydisFormatterBuffer, *ZydisFormatterContext) uint32

type ZydisFormatterRegisterFunc func(*ZydisFormatter, *ZydisFormatterBuffer, *ZydisFormatterContext, ZydisRegister) uint32

type ZydisFormatterDecoratorFunc func(*ZydisFormatter, *ZydisFormatterBuffer, *ZydisFormatterContext, ZydisDecorator) uint32

// Source: unknown.h:4 -> ZydisInstructionCategory_
type ZydisInstructionCategory_ uint32

const (
	ZydisCategoryInvalid ZydisInstructionCategory_ = 0
	ZydisCategoryAdoxAdcx ZydisInstructionCategory_ = 1
	ZydisCategoryAes ZydisInstructionCategory_ = 2
	ZydisCategoryAmd3dnow ZydisInstructionCategory_ = 3
	ZydisCategoryAmxTile ZydisInstructionCategory_ = 4
	ZydisCategoryApx ZydisInstructionCategory_ = 5
	ZydisCategoryAvx ZydisInstructionCategory_ = 6
	ZydisCategoryAvx2 ZydisInstructionCategory_ = 7
	ZydisCategoryAvx2gather ZydisInstructionCategory_ = 8
	ZydisCategoryAvx512 ZydisInstructionCategory_ = 9
	ZydisCategoryAvx5124fmaps ZydisInstructionCategory_ = 10
	ZydisCategoryAvx5124vnniw ZydisInstructionCategory_ = 11
	ZydisCategoryAvx512Bitalg ZydisInstructionCategory_ = 12
	ZydisCategoryAvx512Vbmi ZydisInstructionCategory_ = 13
	ZydisCategoryAvx512Vp2intersect ZydisInstructionCategory_ = 14
	ZydisCategoryAvxIfma ZydisInstructionCategory_ = 15
	ZydisCategoryBinary ZydisInstructionCategory_ = 16
	ZydisCategoryBitbyte ZydisInstructionCategory_ = 17
	ZydisCategoryBlend ZydisInstructionCategory_ = 18
	ZydisCategoryBmi1 ZydisInstructionCategory_ = 19
	ZydisCategoryBmi2 ZydisInstructionCategory_ = 20
	ZydisCategoryBroadcast ZydisInstructionCategory_ = 21
	ZydisCategoryCall ZydisInstructionCategory_ = 22
	ZydisCategoryCet ZydisInstructionCategory_ = 23
	ZydisCategoryCldemote ZydisInstructionCategory_ = 24
	ZydisCategoryClflushopt ZydisInstructionCategory_ = 25
	ZydisCategoryClwb ZydisInstructionCategory_ = 26
	ZydisCategoryClzero ZydisInstructionCategory_ = 27
	ZydisCategoryCmov ZydisInstructionCategory_ = 28
	ZydisCategoryCompress ZydisInstructionCategory_ = 29
	ZydisCategoryCondBr ZydisInstructionCategory_ = 30
	ZydisCategoryConflict ZydisInstructionCategory_ = 31
	ZydisCategoryConvert ZydisInstructionCategory_ = 32
	ZydisCategoryDataxfer ZydisInstructionCategory_ = 33
	ZydisCategoryDecimal ZydisInstructionCategory_ = 34
	ZydisCategoryEnqcmd ZydisInstructionCategory_ = 35
	ZydisCategoryExpand ZydisInstructionCategory_ = 36
	ZydisCategoryFcmov ZydisInstructionCategory_ = 37
	ZydisCategoryFlagop ZydisInstructionCategory_ = 38
	ZydisCategoryFma4 ZydisInstructionCategory_ = 39
	ZydisCategoryFp16 ZydisInstructionCategory_ = 40
	ZydisCategoryFred ZydisInstructionCategory_ = 41
	ZydisCategoryGather ZydisInstructionCategory_ = 42
	ZydisCategoryGfni ZydisInstructionCategory_ = 43
	ZydisCategoryHreset ZydisInstructionCategory_ = 44
	ZydisCategoryIfma ZydisInstructionCategory_ = 45
	ZydisCategoryInterrupt ZydisInstructionCategory_ = 46
	ZydisCategoryIo ZydisInstructionCategory_ = 47
	ZydisCategoryIostringop ZydisInstructionCategory_ = 48
	ZydisCategoryKeylocker ZydisInstructionCategory_ = 49
	ZydisCategoryKeylockerWide ZydisInstructionCategory_ = 50
	ZydisCategoryKmask ZydisInstructionCategory_ = 51
	ZydisCategoryKnc ZydisInstructionCategory_ = 52
	ZydisCategoryKncmask ZydisInstructionCategory_ = 53
	ZydisCategoryKncscalar ZydisInstructionCategory_ = 54
	ZydisCategoryLegacy ZydisInstructionCategory_ = 55
	ZydisCategoryLkgs ZydisInstructionCategory_ = 56
	ZydisCategoryLogical ZydisInstructionCategory_ = 57
	ZydisCategoryLogicalFp ZydisInstructionCategory_ = 58
	ZydisCategoryLzcnt ZydisInstructionCategory_ = 59
	ZydisCategoryMisc ZydisInstructionCategory_ = 60
	ZydisCategoryMmx ZydisInstructionCategory_ = 61
	ZydisCategoryMovdir ZydisInstructionCategory_ = 62
	ZydisCategoryMpx ZydisInstructionCategory_ = 63
	ZydisCategoryMsrlist ZydisInstructionCategory_ = 64
	ZydisCategoryNop ZydisInstructionCategory_ = 65
	ZydisCategoryPadlock ZydisInstructionCategory_ = 66
	ZydisCategoryPbndkb ZydisInstructionCategory_ = 67
	ZydisCategoryPclmulqdq ZydisInstructionCategory_ = 68
	ZydisCategoryPcommit ZydisInstructionCategory_ = 69
	ZydisCategoryPconfig ZydisInstructionCategory_ = 70
	ZydisCategoryPku ZydisInstructionCategory_ = 71
	ZydisCategoryPop ZydisInstructionCategory_ = 72
	ZydisCategoryPrefetch ZydisInstructionCategory_ = 73
	ZydisCategoryPrefetchwt1 ZydisInstructionCategory_ = 74
	ZydisCategoryPt ZydisInstructionCategory_ = 75
	ZydisCategoryPush ZydisInstructionCategory_ = 76
	ZydisCategoryRdpid ZydisInstructionCategory_ = 77
	ZydisCategoryRdpru ZydisInstructionCategory_ = 78
	ZydisCategoryRdrand ZydisInstructionCategory_ = 79
	ZydisCategoryRdseed ZydisInstructionCategory_ = 80
	ZydisCategoryRdwrfsgs ZydisInstructionCategory_ = 81
	ZydisCategoryRet ZydisInstructionCategory_ = 82
	ZydisCategoryRotate ZydisInstructionCategory_ = 83
	ZydisCategoryScatter ZydisInstructionCategory_ = 84
	ZydisCategorySegop ZydisInstructionCategory_ = 85
	ZydisCategorySemaphore ZydisInstructionCategory_ = 86
	ZydisCategorySerialize ZydisInstructionCategory_ = 87
	ZydisCategorySetcc ZydisInstructionCategory_ = 88
	ZydisCategorySgx ZydisInstructionCategory_ = 89
	ZydisCategorySha ZydisInstructionCategory_ = 90
	ZydisCategorySha512 ZydisInstructionCategory_ = 91
	ZydisCategoryShift ZydisInstructionCategory_ = 92
	ZydisCategorySmap ZydisInstructionCategory_ = 93
	ZydisCategorySse ZydisInstructionCategory_ = 94
	ZydisCategoryStringop ZydisInstructionCategory_ = 95
	ZydisCategorySttni ZydisInstructionCategory_ = 96
	ZydisCategorySyscall ZydisInstructionCategory_ = 97
	ZydisCategorySysret ZydisInstructionCategory_ = 98
	ZydisCategorySystem ZydisInstructionCategory_ = 99
	ZydisCategoryTbm ZydisInstructionCategory_ = 100
	ZydisCategoryTsxLdtrk ZydisInstructionCategory_ = 101
	ZydisCategoryUfma ZydisInstructionCategory_ = 102
	ZydisCategoryUintr ZydisInstructionCategory_ = 103
	ZydisCategoryUncondBr ZydisInstructionCategory_ = 104
	ZydisCategoryVaes ZydisInstructionCategory_ = 105
	ZydisCategoryVbmi2 ZydisInstructionCategory_ = 106
	ZydisCategoryVex ZydisInstructionCategory_ = 107
	ZydisCategoryVfma ZydisInstructionCategory_ = 108
	ZydisCategoryVpclmulqdq ZydisInstructionCategory_ = 109
	ZydisCategoryVtx ZydisInstructionCategory_ = 110
	ZydisCategoryWaitpkg ZydisInstructionCategory_ = 111
	ZydisCategoryWidenop ZydisInstructionCategory_ = 112
	ZydisCategoryWrmsrns ZydisInstructionCategory_ = 113
	ZydisCategoryX87Alu ZydisInstructionCategory_ = 114
	ZydisCategoryXop ZydisInstructionCategory_ = 115
	ZydisCategoryXsave ZydisInstructionCategory_ = 116
	ZydisCategoryXsaveopt ZydisInstructionCategory_ = 117
	ZydisCategoryMaxValue ZydisInstructionCategory_ = 117
	ZydisCategoryRequiredBits ZydisInstructionCategory_ = 7
)

func (z ZydisInstructionCategory_) String() string {
	switch z {
	case ZydisCategoryInvalid:
		return "Zydis Category Invalid"
	case ZydisCategoryAdoxAdcx:
		return "Zydis Category Adox Adcx"
	case ZydisCategoryAes:
		return "Zydis Category Aes"
	case ZydisCategoryAmd3dnow:
		return "Zydis Category Amd 3dnow"
	case ZydisCategoryAmxTile:
		return "Zydis Category Amx Tile"
	case ZydisCategoryApx:
		return "Zydis Category Apx"
	case ZydisCategoryAvx:
		return "Zydis Category Avx"
	case ZydisCategoryAvx2:
		return "Zydis Category Avx 2"
	case ZydisCategoryAvx2gather:
		return "Zydis Category Avx 2gather"
	case ZydisCategoryAvx512:
		return "Zydis Category Avx 512"
	case ZydisCategoryAvx5124fmaps:
		return "Zydis Category Avx 5124fmaps"
	case ZydisCategoryAvx5124vnniw:
		return "Zydis Category Avx 5124vnniw"
	case ZydisCategoryAvx512Bitalg:
		return "Zydis Category Avx 512 Bitalg"
	case ZydisCategoryAvx512Vbmi:
		return "Zydis Category Avx 512 Vbmi"
	case ZydisCategoryAvx512Vp2intersect:
		return "Zydis Category Avx 512 Vp 2intersect"
	case ZydisCategoryAvxIfma:
		return "Zydis Category Avx Ifma"
	case ZydisCategoryBinary:
		return "Zydis Category Binary"
	case ZydisCategoryBitbyte:
		return "Zydis Category Bitbyte"
	case ZydisCategoryBlend:
		return "Zydis Category Blend"
	case ZydisCategoryBmi1:
		return "Zydis Category Bmi 1"
	case ZydisCategoryBmi2:
		return "Zydis Category Bmi 2"
	case ZydisCategoryBroadcast:
		return "Zydis Category Broadcast"
	case ZydisCategoryCall:
		return "Zydis Category Call"
	case ZydisCategoryCet:
		return "Zydis Category Cet"
	case ZydisCategoryCldemote:
		return "Zydis Category Cldemote"
	case ZydisCategoryClflushopt:
		return "Zydis Category Clflushopt"
	case ZydisCategoryClwb:
		return "Zydis Category Clwb"
	case ZydisCategoryClzero:
		return "Zydis Category Clzero"
	case ZydisCategoryCmov:
		return "Zydis Category Cmov"
	case ZydisCategoryCompress:
		return "Zydis Category Compress"
	case ZydisCategoryCondBr:
		return "Zydis Category Cond Br"
	case ZydisCategoryConflict:
		return "Zydis Category Conflict"
	case ZydisCategoryConvert:
		return "Zydis Category Convert"
	case ZydisCategoryDataxfer:
		return "Zydis Category Dataxfer"
	case ZydisCategoryDecimal:
		return "Zydis Category Decimal"
	case ZydisCategoryEnqcmd:
		return "Zydis Category Enqcmd"
	case ZydisCategoryExpand:
		return "Zydis Category Expand"
	case ZydisCategoryFcmov:
		return "Zydis Category Fcmov"
	case ZydisCategoryFlagop:
		return "Zydis Category Flagop"
	case ZydisCategoryFma4:
		return "Zydis Category Fma 4"
	case ZydisCategoryFp16:
		return "Zydis Category Fp 16"
	case ZydisCategoryFred:
		return "Zydis Category Fred"
	case ZydisCategoryGather:
		return "Zydis Category Gather"
	case ZydisCategoryGfni:
		return "Zydis Category Gfni"
	case ZydisCategoryHreset:
		return "Zydis Category Hreset"
	case ZydisCategoryIfma:
		return "Zydis Category Ifma"
	case ZydisCategoryInterrupt:
		return "Zydis Category Interrupt"
	case ZydisCategoryIo:
		return "Zydis Category Io"
	case ZydisCategoryIostringop:
		return "Zydis Category Iostringop"
	case ZydisCategoryKeylocker:
		return "Zydis Category Keylocker"
	case ZydisCategoryKeylockerWide:
		return "Zydis Category Keylocker Wide"
	case ZydisCategoryKmask:
		return "Zydis Category Kmask"
	case ZydisCategoryKnc:
		return "Zydis Category Knc"
	case ZydisCategoryKncmask:
		return "Zydis Category Kncmask"
	case ZydisCategoryKncscalar:
		return "Zydis Category Kncscalar"
	case ZydisCategoryLegacy:
		return "Zydis Category Legacy"
	case ZydisCategoryLkgs:
		return "Zydis Category Lkgs"
	case ZydisCategoryLogical:
		return "Zydis Category Logical"
	case ZydisCategoryLogicalFp:
		return "Zydis Category Logical Fp"
	case ZydisCategoryLzcnt:
		return "Zydis Category Lzcnt"
	case ZydisCategoryMisc:
		return "Zydis Category Misc"
	case ZydisCategoryMmx:
		return "Zydis Category Mmx"
	case ZydisCategoryMovdir:
		return "Zydis Category Movdir"
	case ZydisCategoryMpx:
		return "Zydis Category Mpx"
	case ZydisCategoryMsrlist:
		return "Zydis Category Msrlist"
	case ZydisCategoryNop:
		return "Zydis Category Nop"
	case ZydisCategoryPadlock:
		return "Zydis Category Padlock"
	case ZydisCategoryPbndkb:
		return "Zydis Category Pbndkb"
	case ZydisCategoryPclmulqdq:
		return "Zydis Category Pclmulqdq"
	case ZydisCategoryPcommit:
		return "Zydis Category Pcommit"
	case ZydisCategoryPconfig:
		return "Zydis Category Pconfig"
	case ZydisCategoryPku:
		return "Zydis Category Pku"
	case ZydisCategoryPop:
		return "Zydis Category Pop"
	case ZydisCategoryPrefetch:
		return "Zydis Category Prefetch"
	case ZydisCategoryPrefetchwt1:
		return "Zydis Category Prefetchwt 1"
	case ZydisCategoryPt:
		return "Zydis Category Pt"
	case ZydisCategoryPush:
		return "Zydis Category Push"
	case ZydisCategoryRdpid:
		return "Zydis Category Rdpid"
	case ZydisCategoryRdpru:
		return "Zydis Category Rdpru"
	case ZydisCategoryRdrand:
		return "Zydis Category Rdrand"
	case ZydisCategoryRdseed:
		return "Zydis Category Rdseed"
	case ZydisCategoryRdwrfsgs:
		return "Zydis Category Rdwrfsgs"
	case ZydisCategoryRet:
		return "Zydis Category Ret"
	case ZydisCategoryRotate:
		return "Zydis Category Rotate"
	case ZydisCategoryScatter:
		return "Zydis Category Scatter"
	case ZydisCategorySegop:
		return "Zydis Category Segop"
	case ZydisCategorySemaphore:
		return "Zydis Category Semaphore"
	case ZydisCategorySerialize:
		return "Zydis Category Serialize"
	case ZydisCategorySetcc:
		return "Zydis Category Setcc"
	case ZydisCategorySgx:
		return "Zydis Category Sgx"
	case ZydisCategorySha:
		return "Zydis Category Sha"
	case ZydisCategorySha512:
		return "Zydis Category Sha 512"
	case ZydisCategoryShift:
		return "Zydis Category Shift"
	case ZydisCategorySmap:
		return "Zydis Category Smap"
	case ZydisCategorySse:
		return "Zydis Category Sse"
	case ZydisCategoryStringop:
		return "Zydis Category Stringop"
	case ZydisCategorySttni:
		return "Zydis Category Sttni"
	case ZydisCategorySyscall:
		return "Zydis Category Syscall"
	case ZydisCategorySysret:
		return "Zydis Category Sysret"
	case ZydisCategorySystem:
		return "Zydis Category System"
	case ZydisCategoryTbm:
		return "Zydis Category Tbm"
	case ZydisCategoryTsxLdtrk:
		return "Zydis Category Tsx Ldtrk"
	case ZydisCategoryUfma:
		return "Zydis Category Ufma"
	case ZydisCategoryUintr:
		return "Zydis Category Uintr"
	case ZydisCategoryUncondBr:
		return "Zydis Category Uncond Br"
	case ZydisCategoryVaes:
		return "Zydis Category Vaes"
	case ZydisCategoryVbmi2:
		return "Zydis Category Vbmi 2"
	case ZydisCategoryVex:
		return "Zydis Category Vex"
	case ZydisCategoryVfma:
		return "Zydis Category Vfma"
	case ZydisCategoryVpclmulqdq:
		return "Zydis Category Vpclmulqdq"
	case ZydisCategoryVtx:
		return "Zydis Category Vtx"
	case ZydisCategoryWaitpkg:
		return "Zydis Category Waitpkg"
	case ZydisCategoryWidenop:
		return "Zydis Category Widenop"
	case ZydisCategoryWrmsrns:
		return "Zydis Category Wrmsrns"
	case ZydisCategoryX87Alu:
		return "Zydis Category X87 Alu"
	case ZydisCategoryXop:
		return "Zydis Category Xop"
	case ZydisCategoryXsave:
		return "Zydis Category Xsave"
	case ZydisCategoryXsaveopt:
		return "Zydis Category Xsaveopt"
	default:
		return fmt.Sprintf("ZydisInstructionCategory_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:4 -> ZydisISASet_
type ZydisISASet_ uint32

const (
	ZydisIsaSetInvalid ZydisISASet_ = 0
	ZydisIsaSetAdoxAdcx ZydisISASet_ = 1
	ZydisIsaSetAes ZydisISASet_ = 2
	ZydisIsaSetAmd ZydisISASet_ = 3
	ZydisIsaSetAmd3dnow ZydisISASet_ = 4
	ZydisIsaSetAmdInvlpgb ZydisISASet_ = 5
	ZydisIsaSetAmxBf16 ZydisISASet_ = 6
	ZydisIsaSetAmxFp16 ZydisISASet_ = 7
	ZydisIsaSetAmxInt8 ZydisISASet_ = 8
	ZydisIsaSetAmxTile ZydisISASet_ = 9
	ZydisIsaSetApxF ZydisISASet_ = 10
	ZydisIsaSetApxFAdx ZydisISASet_ = 11
	ZydisIsaSetApxFAmx ZydisISASet_ = 12
	ZydisIsaSetApxFBmi1 ZydisISASet_ = 13
	ZydisIsaSetApxFBmi2 ZydisISASet_ = 14
	ZydisIsaSetApxFCet ZydisISASet_ = 15
	ZydisIsaSetApxFCmpccxadd ZydisISASet_ = 16
	ZydisIsaSetApxFEnqcmd ZydisISASet_ = 17
	ZydisIsaSetApxFInvpcid ZydisISASet_ = 18
	ZydisIsaSetApxFKopb ZydisISASet_ = 19
	ZydisIsaSetApxFKopd ZydisISASet_ = 20
	ZydisIsaSetApxFKopq ZydisISASet_ = 21
	ZydisIsaSetApxFKopw ZydisISASet_ = 22
	ZydisIsaSetApxFLzcnt ZydisISASet_ = 23
	ZydisIsaSetApxFMovbe ZydisISASet_ = 24
	ZydisIsaSetApxFMovdir64b ZydisISASet_ = 25
	ZydisIsaSetApxFMovdiri ZydisISASet_ = 26
	ZydisIsaSetApxFPopcnt ZydisISASet_ = 27
	ZydisIsaSetApxFRaoInt ZydisISASet_ = 28
	ZydisIsaSetApxFUserMsr ZydisISASet_ = 29
	ZydisIsaSetApxFVmx ZydisISASet_ = 30
	ZydisIsaSetAvx ZydisISASet_ = 31
	ZydisIsaSetAvx2 ZydisISASet_ = 32
	ZydisIsaSetAvx2gather ZydisISASet_ = 33
	ZydisIsaSetAvx512bw128 ZydisISASet_ = 34
	ZydisIsaSetAvx512bw128n ZydisISASet_ = 35
	ZydisIsaSetAvx512bw256 ZydisISASet_ = 36
	ZydisIsaSetAvx512bw512 ZydisISASet_ = 37
	ZydisIsaSetAvx512bwKop ZydisISASet_ = 38
	ZydisIsaSetAvx512cd128 ZydisISASet_ = 39
	ZydisIsaSetAvx512cd256 ZydisISASet_ = 40
	ZydisIsaSetAvx512cd512 ZydisISASet_ = 41
	ZydisIsaSetAvx512dq128 ZydisISASet_ = 42
	ZydisIsaSetAvx512dq128n ZydisISASet_ = 43
	ZydisIsaSetAvx512dq256 ZydisISASet_ = 44
	ZydisIsaSetAvx512dq512 ZydisISASet_ = 45
	ZydisIsaSetAvx512dqKop ZydisISASet_ = 46
	ZydisIsaSetAvx512dqScalar ZydisISASet_ = 47
	ZydisIsaSetAvx512er512 ZydisISASet_ = 48
	ZydisIsaSetAvx512erScalar ZydisISASet_ = 49
	ZydisIsaSetAvx512f128 ZydisISASet_ = 50
	ZydisIsaSetAvx512f128n ZydisISASet_ = 51
	ZydisIsaSetAvx512f256 ZydisISASet_ = 52
	ZydisIsaSetAvx512f512 ZydisISASet_ = 53
	ZydisIsaSetAvx512fKop ZydisISASet_ = 54
	ZydisIsaSetAvx512fScalar ZydisISASet_ = 55
	ZydisIsaSetAvx512pf512 ZydisISASet_ = 56
	ZydisIsaSetAvx5124fmaps512 ZydisISASet_ = 57
	ZydisIsaSetAvx5124fmapsScalar ZydisISASet_ = 58
	ZydisIsaSetAvx5124vnniw512 ZydisISASet_ = 59
	ZydisIsaSetAvx512Bf16128 ZydisISASet_ = 60
	ZydisIsaSetAvx512Bf16256 ZydisISASet_ = 61
	ZydisIsaSetAvx512Bf16512 ZydisISASet_ = 62
	ZydisIsaSetAvx512Bitalg128 ZydisISASet_ = 63
	ZydisIsaSetAvx512Bitalg256 ZydisISASet_ = 64
	ZydisIsaSetAvx512Bitalg512 ZydisISASet_ = 65
	ZydisIsaSetAvx512Fp16128 ZydisISASet_ = 66
	ZydisIsaSetAvx512Fp16128n ZydisISASet_ = 67
	ZydisIsaSetAvx512Fp16256 ZydisISASet_ = 68
	ZydisIsaSetAvx512Fp16512 ZydisISASet_ = 69
	ZydisIsaSetAvx512Fp16Scalar ZydisISASet_ = 70
	ZydisIsaSetAvx512Gfni128 ZydisISASet_ = 71
	ZydisIsaSetAvx512Gfni256 ZydisISASet_ = 72
	ZydisIsaSetAvx512Gfni512 ZydisISASet_ = 73
	ZydisIsaSetAvx512Ifma128 ZydisISASet_ = 74
	ZydisIsaSetAvx512Ifma256 ZydisISASet_ = 75
	ZydisIsaSetAvx512Ifma512 ZydisISASet_ = 76
	ZydisIsaSetAvx512Vaes128 ZydisISASet_ = 77
	ZydisIsaSetAvx512Vaes256 ZydisISASet_ = 78
	ZydisIsaSetAvx512Vaes512 ZydisISASet_ = 79
	ZydisIsaSetAvx512Vbmi2128 ZydisISASet_ = 80
	ZydisIsaSetAvx512Vbmi2256 ZydisISASet_ = 81
	ZydisIsaSetAvx512Vbmi2512 ZydisISASet_ = 82
	ZydisIsaSetAvx512Vbmi128 ZydisISASet_ = 83
	ZydisIsaSetAvx512Vbmi256 ZydisISASet_ = 84
	ZydisIsaSetAvx512Vbmi512 ZydisISASet_ = 85
	ZydisIsaSetAvx512Vnni128 ZydisISASet_ = 86
	ZydisIsaSetAvx512Vnni256 ZydisISASet_ = 87
	ZydisIsaSetAvx512Vnni512 ZydisISASet_ = 88
	ZydisIsaSetAvx512Vp2intersect128 ZydisISASet_ = 89
	ZydisIsaSetAvx512Vp2intersect256 ZydisISASet_ = 90
	ZydisIsaSetAvx512Vp2intersect512 ZydisISASet_ = 91
	ZydisIsaSetAvx512Vpclmulqdq128 ZydisISASet_ = 92
	ZydisIsaSetAvx512Vpclmulqdq256 ZydisISASet_ = 93
	ZydisIsaSetAvx512Vpclmulqdq512 ZydisISASet_ = 94
	ZydisIsaSetAvx512Vpopcntdq128 ZydisISASet_ = 95
	ZydisIsaSetAvx512Vpopcntdq256 ZydisISASet_ = 96
	ZydisIsaSetAvx512Vpopcntdq512 ZydisISASet_ = 97
	ZydisIsaSetAvxaes ZydisISASet_ = 98
	ZydisIsaSetAvxGfni ZydisISASet_ = 99
	ZydisIsaSetAvxIfma ZydisISASet_ = 100
	ZydisIsaSetAvxNeConvert ZydisISASet_ = 101
	ZydisIsaSetAvxVnni ZydisISASet_ = 102
	ZydisIsaSetAvxVnniInt16 ZydisISASet_ = 103
	ZydisIsaSetAvxVnniInt8 ZydisISASet_ = 104
	ZydisIsaSetBmi1 ZydisISASet_ = 105
	ZydisIsaSetBmi2 ZydisISASet_ = 106
	ZydisIsaSetCet ZydisISASet_ = 107
	ZydisIsaSetCldemote ZydisISASet_ = 108
	ZydisIsaSetClflushopt ZydisISASet_ = 109
	ZydisIsaSetClfsh ZydisISASet_ = 110
	ZydisIsaSetClwb ZydisISASet_ = 111
	ZydisIsaSetClzero ZydisISASet_ = 112
	ZydisIsaSetCmov ZydisISASet_ = 113
	ZydisIsaSetCmpxchg16b ZydisISASet_ = 114
	ZydisIsaSetEnqcmd ZydisISASet_ = 115
	ZydisIsaSetF16c ZydisISASet_ = 116
	ZydisIsaSetFatNop ZydisISASet_ = 117
	ZydisIsaSetFcmov ZydisISASet_ = 118
	ZydisIsaSetFcomi ZydisISASet_ = 119
	ZydisIsaSetFma ZydisISASet_ = 120
	ZydisIsaSetFma4 ZydisISASet_ = 121
	ZydisIsaSetFred ZydisISASet_ = 122
	ZydisIsaSetFxsave ZydisISASet_ = 123
	ZydisIsaSetFxsave64 ZydisISASet_ = 124
	ZydisIsaSetGfni ZydisISASet_ = 125
	ZydisIsaSetHreset ZydisISASet_ = 126
	ZydisIsaSetI186 ZydisISASet_ = 127
	ZydisIsaSetI286protected ZydisISASet_ = 128
	ZydisIsaSetI286real ZydisISASet_ = 129
	ZydisIsaSetI386 ZydisISASet_ = 130
	ZydisIsaSetI486 ZydisISASet_ = 131
	ZydisIsaSetI486real ZydisISASet_ = 132
	ZydisIsaSetI86 ZydisISASet_ = 133
	ZydisIsaSetIcachePrefetch ZydisISASet_ = 134
	ZydisIsaSetInvpcid ZydisISASet_ = 135
	ZydisIsaSetKeylocker ZydisISASet_ = 136
	ZydisIsaSetKeylockerWide ZydisISASet_ = 137
	ZydisIsaSetKnce ZydisISASet_ = 138
	ZydisIsaSetKncjkbr ZydisISASet_ = 139
	ZydisIsaSetKncstream ZydisISASet_ = 140
	ZydisIsaSetKncv ZydisISASet_ = 141
	ZydisIsaSetKncMisc ZydisISASet_ = 142
	ZydisIsaSetKncPfHint ZydisISASet_ = 143
	ZydisIsaSetLahf ZydisISASet_ = 144
	ZydisIsaSetLkgs ZydisISASet_ = 145
	ZydisIsaSetLongmode ZydisISASet_ = 146
	ZydisIsaSetLwp ZydisISASet_ = 147
	ZydisIsaSetLzcnt ZydisISASet_ = 148
	ZydisIsaSetMcommit ZydisISASet_ = 149
	ZydisIsaSetMonitor ZydisISASet_ = 150
	ZydisIsaSetMonitorx ZydisISASet_ = 151
	ZydisIsaSetMovbe ZydisISASet_ = 152
	ZydisIsaSetMovdir ZydisISASet_ = 153
	ZydisIsaSetMpx ZydisISASet_ = 154
	ZydisIsaSetMsrlist ZydisISASet_ = 155
	ZydisIsaSetPadlockAce ZydisISASet_ = 156
	ZydisIsaSetPadlockPhe ZydisISASet_ = 157
	ZydisIsaSetPadlockPmm ZydisISASet_ = 158
	ZydisIsaSetPadlockRng ZydisISASet_ = 159
	ZydisIsaSetPause ZydisISASet_ = 160
	ZydisIsaSetPbndkb ZydisISASet_ = 161
	ZydisIsaSetPclmulqdq ZydisISASet_ = 162
	ZydisIsaSetPcommit ZydisISASet_ = 163
	ZydisIsaSetPconfig ZydisISASet_ = 164
	ZydisIsaSetPentiummmx ZydisISASet_ = 165
	ZydisIsaSetPentiumreal ZydisISASet_ = 166
	ZydisIsaSetPku ZydisISASet_ = 167
	ZydisIsaSetPopcnt ZydisISASet_ = 168
	ZydisIsaSetPpro ZydisISASet_ = 169
	ZydisIsaSetPrefetchwt1 ZydisISASet_ = 170
	ZydisIsaSetPrefetchNop ZydisISASet_ = 171
	ZydisIsaSetPt ZydisISASet_ = 172
	ZydisIsaSetRaoInt ZydisISASet_ = 173
	ZydisIsaSetRdpid ZydisISASet_ = 174
	ZydisIsaSetRdpmc ZydisISASet_ = 175
	ZydisIsaSetRdpru ZydisISASet_ = 176
	ZydisIsaSetRdrand ZydisISASet_ = 177
	ZydisIsaSetRdseed ZydisISASet_ = 178
	ZydisIsaSetRdtscp ZydisISASet_ = 179
	ZydisIsaSetRdwrfsgs ZydisISASet_ = 180
	ZydisIsaSetRtm ZydisISASet_ = 181
	ZydisIsaSetSerialize ZydisISASet_ = 182
	ZydisIsaSetSgx ZydisISASet_ = 183
	ZydisIsaSetSgxEnclv ZydisISASet_ = 184
	ZydisIsaSetSha ZydisISASet_ = 185
	ZydisIsaSetSha512 ZydisISASet_ = 186
	ZydisIsaSetSm3 ZydisISASet_ = 187
	ZydisIsaSetSm4 ZydisISASet_ = 188
	ZydisIsaSetSmap ZydisISASet_ = 189
	ZydisIsaSetSmx ZydisISASet_ = 190
	ZydisIsaSetSnp ZydisISASet_ = 191
	ZydisIsaSetSse ZydisISASet_ = 192
	ZydisIsaSetSse2 ZydisISASet_ = 193
	ZydisIsaSetSse2mmx ZydisISASet_ = 194
	ZydisIsaSetSse3 ZydisISASet_ = 195
	ZydisIsaSetSse3x87 ZydisISASet_ = 196
	ZydisIsaSetSse4 ZydisISASet_ = 197
	ZydisIsaSetSse42 ZydisISASet_ = 198
	ZydisIsaSetSse4a ZydisISASet_ = 199
	ZydisIsaSetSsemxcsr ZydisISASet_ = 200
	ZydisIsaSetSsePrefetch ZydisISASet_ = 201
	ZydisIsaSetSsse3 ZydisISASet_ = 202
	ZydisIsaSetSsse3mmx ZydisISASet_ = 203
	ZydisIsaSetSvm ZydisISASet_ = 204
	ZydisIsaSetTbm ZydisISASet_ = 205
	ZydisIsaSetTdx ZydisISASet_ = 206
	ZydisIsaSetTsxLdtrk ZydisISASet_ = 207
	ZydisIsaSetUintr ZydisISASet_ = 208
	ZydisIsaSetVaes ZydisISASet_ = 209
	ZydisIsaSetVmfunc ZydisISASet_ = 210
	ZydisIsaSetVpclmulqdq ZydisISASet_ = 211
	ZydisIsaSetVtx ZydisISASet_ = 212
	ZydisIsaSetWaitpkg ZydisISASet_ = 213
	ZydisIsaSetWrmsrns ZydisISASet_ = 214
	ZydisIsaSetX87 ZydisISASet_ = 215
	ZydisIsaSetXop ZydisISASet_ = 216
	ZydisIsaSetXsave ZydisISASet_ = 217
	ZydisIsaSetXsavec ZydisISASet_ = 218
	ZydisIsaSetXsaveopt ZydisISASet_ = 219
	ZydisIsaSetXsaves ZydisISASet_ = 220
	ZydisIsaSetMaxValue ZydisISASet_ = 220
	ZydisIsaSetRequiredBits ZydisISASet_ = 8
)

func (z ZydisISASet_) String() string {
	switch z {
	case ZydisIsaSetInvalid:
		return "Zydis Isa Set Invalid"
	case ZydisIsaSetAdoxAdcx:
		return "Zydis Isa Set Adox Adcx"
	case ZydisIsaSetAes:
		return "Zydis Isa Set Aes"
	case ZydisIsaSetAmd:
		return "Zydis Isa Set Amd"
	case ZydisIsaSetAmd3dnow:
		return "Zydis Isa Set Amd 3dnow"
	case ZydisIsaSetAmdInvlpgb:
		return "Zydis Isa Set Amd Invlpgb"
	case ZydisIsaSetAmxBf16:
		return "Zydis Isa Set Amx Bf 16"
	case ZydisIsaSetAmxFp16:
		return "Zydis Isa Set Amx Fp 16"
	case ZydisIsaSetAmxInt8:
		return "Zydis Isa Set Amx Int 8"
	case ZydisIsaSetAmxTile:
		return "Zydis Isa Set Amx Tile"
	case ZydisIsaSetApxF:
		return "Zydis Isa Set Apx F"
	case ZydisIsaSetApxFAdx:
		return "Zydis Isa Set Apx F Adx"
	case ZydisIsaSetApxFAmx:
		return "Zydis Isa Set Apx F Amx"
	case ZydisIsaSetApxFBmi1:
		return "Zydis Isa Set Apx F Bmi 1"
	case ZydisIsaSetApxFBmi2:
		return "Zydis Isa Set Apx F Bmi 2"
	case ZydisIsaSetApxFCet:
		return "Zydis Isa Set Apx F Cet"
	case ZydisIsaSetApxFCmpccxadd:
		return "Zydis Isa Set Apx F Cmpccxadd"
	case ZydisIsaSetApxFEnqcmd:
		return "Zydis Isa Set Apx F Enqcmd"
	case ZydisIsaSetApxFInvpcid:
		return "Zydis Isa Set Apx F Invpcid"
	case ZydisIsaSetApxFKopb:
		return "Zydis Isa Set Apx F Kopb"
	case ZydisIsaSetApxFKopd:
		return "Zydis Isa Set Apx F Kopd"
	case ZydisIsaSetApxFKopq:
		return "Zydis Isa Set Apx F Kopq"
	case ZydisIsaSetApxFKopw:
		return "Zydis Isa Set Apx F Kopw"
	case ZydisIsaSetApxFLzcnt:
		return "Zydis Isa Set Apx F Lzcnt"
	case ZydisIsaSetApxFMovbe:
		return "Zydis Isa Set Apx F Movbe"
	case ZydisIsaSetApxFMovdir64b:
		return "Zydis Isa Set Apx F Movdir 64b"
	case ZydisIsaSetApxFMovdiri:
		return "Zydis Isa Set Apx F Movdiri"
	case ZydisIsaSetApxFPopcnt:
		return "Zydis Isa Set Apx F Popcnt"
	case ZydisIsaSetApxFRaoInt:
		return "Zydis Isa Set Apx F Rao Int"
	case ZydisIsaSetApxFUserMsr:
		return "Zydis Isa Set Apx F User Msr"
	case ZydisIsaSetApxFVmx:
		return "Zydis Isa Set Apx F Vmx"
	case ZydisIsaSetAvx:
		return "Zydis Isa Set Avx"
	case ZydisIsaSetAvx2:
		return "Zydis Isa Set Avx 2"
	case ZydisIsaSetAvx2gather:
		return "Zydis Isa Set Avx 2gather"
	case ZydisIsaSetAvx512bw128:
		return "Zydis Isa Set Avx 512bw 128"
	case ZydisIsaSetAvx512bw128n:
		return "Zydis Isa Set Avx 512bw 128n"
	case ZydisIsaSetAvx512bw256:
		return "Zydis Isa Set Avx 512bw 256"
	case ZydisIsaSetAvx512bw512:
		return "Zydis Isa Set Avx 512bw 512"
	case ZydisIsaSetAvx512bwKop:
		return "Zydis Isa Set Avx 512bw Kop"
	case ZydisIsaSetAvx512cd128:
		return "Zydis Isa Set Avx 512cd 128"
	case ZydisIsaSetAvx512cd256:
		return "Zydis Isa Set Avx 512cd 256"
	case ZydisIsaSetAvx512cd512:
		return "Zydis Isa Set Avx 512cd 512"
	case ZydisIsaSetAvx512dq128:
		return "Zydis Isa Set Avx 512dq 128"
	case ZydisIsaSetAvx512dq128n:
		return "Zydis Isa Set Avx 512dq 128n"
	case ZydisIsaSetAvx512dq256:
		return "Zydis Isa Set Avx 512dq 256"
	case ZydisIsaSetAvx512dq512:
		return "Zydis Isa Set Avx 512dq 512"
	case ZydisIsaSetAvx512dqKop:
		return "Zydis Isa Set Avx 512dq Kop"
	case ZydisIsaSetAvx512dqScalar:
		return "Zydis Isa Set Avx 512dq Scalar"
	case ZydisIsaSetAvx512er512:
		return "Zydis Isa Set Avx 512er 512"
	case ZydisIsaSetAvx512erScalar:
		return "Zydis Isa Set Avx 512er Scalar"
	case ZydisIsaSetAvx512f128:
		return "Zydis Isa Set Avx 512f 128"
	case ZydisIsaSetAvx512f128n:
		return "Zydis Isa Set Avx 512f 128n"
	case ZydisIsaSetAvx512f256:
		return "Zydis Isa Set Avx 512f 256"
	case ZydisIsaSetAvx512f512:
		return "Zydis Isa Set Avx 512f 512"
	case ZydisIsaSetAvx512fKop:
		return "Zydis Isa Set Avx 512f Kop"
	case ZydisIsaSetAvx512fScalar:
		return "Zydis Isa Set Avx 512f Scalar"
	case ZydisIsaSetAvx512pf512:
		return "Zydis Isa Set Avx 512pf 512"
	case ZydisIsaSetAvx5124fmaps512:
		return "Zydis Isa Set Avx 5124fmaps 512"
	case ZydisIsaSetAvx5124fmapsScalar:
		return "Zydis Isa Set Avx 5124fmaps Scalar"
	case ZydisIsaSetAvx5124vnniw512:
		return "Zydis Isa Set Avx 5124vnniw 512"
	case ZydisIsaSetAvx512Bf16128:
		return "Zydis Isa Set Avx 512 Bf 16128"
	case ZydisIsaSetAvx512Bf16256:
		return "Zydis Isa Set Avx 512 Bf 16256"
	case ZydisIsaSetAvx512Bf16512:
		return "Zydis Isa Set Avx 512 Bf 16512"
	case ZydisIsaSetAvx512Bitalg128:
		return "Zydis Isa Set Avx 512 Bitalg 128"
	case ZydisIsaSetAvx512Bitalg256:
		return "Zydis Isa Set Avx 512 Bitalg 256"
	case ZydisIsaSetAvx512Bitalg512:
		return "Zydis Isa Set Avx 512 Bitalg 512"
	case ZydisIsaSetAvx512Fp16128:
		return "Zydis Isa Set Avx 512 Fp 16128"
	case ZydisIsaSetAvx512Fp16128n:
		return "Zydis Isa Set Avx 512 Fp 16128n"
	case ZydisIsaSetAvx512Fp16256:
		return "Zydis Isa Set Avx 512 Fp 16256"
	case ZydisIsaSetAvx512Fp16512:
		return "Zydis Isa Set Avx 512 Fp 16512"
	case ZydisIsaSetAvx512Fp16Scalar:
		return "Zydis Isa Set Avx 512 Fp 16 Scalar"
	case ZydisIsaSetAvx512Gfni128:
		return "Zydis Isa Set Avx 512 Gfni 128"
	case ZydisIsaSetAvx512Gfni256:
		return "Zydis Isa Set Avx 512 Gfni 256"
	case ZydisIsaSetAvx512Gfni512:
		return "Zydis Isa Set Avx 512 Gfni 512"
	case ZydisIsaSetAvx512Ifma128:
		return "Zydis Isa Set Avx 512 Ifma 128"
	case ZydisIsaSetAvx512Ifma256:
		return "Zydis Isa Set Avx 512 Ifma 256"
	case ZydisIsaSetAvx512Ifma512:
		return "Zydis Isa Set Avx 512 Ifma 512"
	case ZydisIsaSetAvx512Vaes128:
		return "Zydis Isa Set Avx 512 Vaes 128"
	case ZydisIsaSetAvx512Vaes256:
		return "Zydis Isa Set Avx 512 Vaes 256"
	case ZydisIsaSetAvx512Vaes512:
		return "Zydis Isa Set Avx 512 Vaes 512"
	case ZydisIsaSetAvx512Vbmi2128:
		return "Zydis Isa Set Avx 512 Vbmi 2128"
	case ZydisIsaSetAvx512Vbmi2256:
		return "Zydis Isa Set Avx 512 Vbmi 2256"
	case ZydisIsaSetAvx512Vbmi2512:
		return "Zydis Isa Set Avx 512 Vbmi 2512"
	case ZydisIsaSetAvx512Vbmi128:
		return "Zydis Isa Set Avx 512 Vbmi 128"
	case ZydisIsaSetAvx512Vbmi256:
		return "Zydis Isa Set Avx 512 Vbmi 256"
	case ZydisIsaSetAvx512Vbmi512:
		return "Zydis Isa Set Avx 512 Vbmi 512"
	case ZydisIsaSetAvx512Vnni128:
		return "Zydis Isa Set Avx 512 Vnni 128"
	case ZydisIsaSetAvx512Vnni256:
		return "Zydis Isa Set Avx 512 Vnni 256"
	case ZydisIsaSetAvx512Vnni512:
		return "Zydis Isa Set Avx 512 Vnni 512"
	case ZydisIsaSetAvx512Vp2intersect128:
		return "Zydis Isa Set Avx 512 Vp 2intersect 128"
	case ZydisIsaSetAvx512Vp2intersect256:
		return "Zydis Isa Set Avx 512 Vp 2intersect 256"
	case ZydisIsaSetAvx512Vp2intersect512:
		return "Zydis Isa Set Avx 512 Vp 2intersect 512"
	case ZydisIsaSetAvx512Vpclmulqdq128:
		return "Zydis Isa Set Avx 512 Vpclmulqdq 128"
	case ZydisIsaSetAvx512Vpclmulqdq256:
		return "Zydis Isa Set Avx 512 Vpclmulqdq 256"
	case ZydisIsaSetAvx512Vpclmulqdq512:
		return "Zydis Isa Set Avx 512 Vpclmulqdq 512"
	case ZydisIsaSetAvx512Vpopcntdq128:
		return "Zydis Isa Set Avx 512 Vpopcntdq 128"
	case ZydisIsaSetAvx512Vpopcntdq256:
		return "Zydis Isa Set Avx 512 Vpopcntdq 256"
	case ZydisIsaSetAvx512Vpopcntdq512:
		return "Zydis Isa Set Avx 512 Vpopcntdq 512"
	case ZydisIsaSetAvxaes:
		return "Zydis Isa Set Avxaes"
	case ZydisIsaSetAvxGfni:
		return "Zydis Isa Set Avx Gfni"
	case ZydisIsaSetAvxIfma:
		return "Zydis Isa Set Avx Ifma"
	case ZydisIsaSetAvxNeConvert:
		return "Zydis Isa Set Avx Ne Convert"
	case ZydisIsaSetAvxVnni:
		return "Zydis Isa Set Avx Vnni"
	case ZydisIsaSetAvxVnniInt16:
		return "Zydis Isa Set Avx Vnni Int 16"
	case ZydisIsaSetAvxVnniInt8:
		return "Zydis Isa Set Avx Vnni Int 8"
	case ZydisIsaSetBmi1:
		return "Zydis Isa Set Bmi 1"
	case ZydisIsaSetBmi2:
		return "Zydis Isa Set Bmi 2"
	case ZydisIsaSetCet:
		return "Zydis Isa Set Cet"
	case ZydisIsaSetCldemote:
		return "Zydis Isa Set Cldemote"
	case ZydisIsaSetClflushopt:
		return "Zydis Isa Set Clflushopt"
	case ZydisIsaSetClfsh:
		return "Zydis Isa Set Clfsh"
	case ZydisIsaSetClwb:
		return "Zydis Isa Set Clwb"
	case ZydisIsaSetClzero:
		return "Zydis Isa Set Clzero"
	case ZydisIsaSetCmov:
		return "Zydis Isa Set Cmov"
	case ZydisIsaSetCmpxchg16b:
		return "Zydis Isa Set Cmpxchg 16b"
	case ZydisIsaSetEnqcmd:
		return "Zydis Isa Set Enqcmd"
	case ZydisIsaSetF16c:
		return "Zydis Isa Set F16c"
	case ZydisIsaSetFatNop:
		return "Zydis Isa Set Fat Nop"
	case ZydisIsaSetFcmov:
		return "Zydis Isa Set Fcmov"
	case ZydisIsaSetFcomi:
		return "Zydis Isa Set Fcomi"
	case ZydisIsaSetFma:
		return "Zydis Isa Set Fma"
	case ZydisIsaSetFma4:
		return "Zydis Isa Set Fma 4"
	case ZydisIsaSetFred:
		return "Zydis Isa Set Fred"
	case ZydisIsaSetFxsave:
		return "Zydis Isa Set Fxsave"
	case ZydisIsaSetFxsave64:
		return "Zydis Isa Set Fxsave 64"
	case ZydisIsaSetGfni:
		return "Zydis Isa Set Gfni"
	case ZydisIsaSetHreset:
		return "Zydis Isa Set Hreset"
	case ZydisIsaSetI186:
		return "Zydis Isa Set I186"
	case ZydisIsaSetI286protected:
		return "Zydis Isa Set I286protected"
	case ZydisIsaSetI286real:
		return "Zydis Isa Set I286real"
	case ZydisIsaSetI386:
		return "Zydis Isa Set I386"
	case ZydisIsaSetI486:
		return "Zydis Isa Set I486"
	case ZydisIsaSetI486real:
		return "Zydis Isa Set I486real"
	case ZydisIsaSetI86:
		return "Zydis Isa Set I86"
	case ZydisIsaSetIcachePrefetch:
		return "Zydis Isa Set Icache Prefetch"
	case ZydisIsaSetInvpcid:
		return "Zydis Isa Set Invpcid"
	case ZydisIsaSetKeylocker:
		return "Zydis Isa Set Keylocker"
	case ZydisIsaSetKeylockerWide:
		return "Zydis Isa Set Keylocker Wide"
	case ZydisIsaSetKnce:
		return "Zydis Isa Set Knce"
	case ZydisIsaSetKncjkbr:
		return "Zydis Isa Set Kncjkbr"
	case ZydisIsaSetKncstream:
		return "Zydis Isa Set Kncstream"
	case ZydisIsaSetKncv:
		return "Zydis Isa Set Kncv"
	case ZydisIsaSetKncMisc:
		return "Zydis Isa Set Knc Misc"
	case ZydisIsaSetKncPfHint:
		return "Zydis Isa Set Knc Pf Hint"
	case ZydisIsaSetLahf:
		return "Zydis Isa Set Lahf"
	case ZydisIsaSetLkgs:
		return "Zydis Isa Set Lkgs"
	case ZydisIsaSetLongmode:
		return "Zydis Isa Set Longmode"
	case ZydisIsaSetLwp:
		return "Zydis Isa Set Lwp"
	case ZydisIsaSetLzcnt:
		return "Zydis Isa Set Lzcnt"
	case ZydisIsaSetMcommit:
		return "Zydis Isa Set Mcommit"
	case ZydisIsaSetMonitor:
		return "Zydis Isa Set Monitor"
	case ZydisIsaSetMonitorx:
		return "Zydis Isa Set Monitorx"
	case ZydisIsaSetMovbe:
		return "Zydis Isa Set Movbe"
	case ZydisIsaSetMovdir:
		return "Zydis Isa Set Movdir"
	case ZydisIsaSetMpx:
		return "Zydis Isa Set Mpx"
	case ZydisIsaSetMsrlist:
		return "Zydis Isa Set Msrlist"
	case ZydisIsaSetPadlockAce:
		return "Zydis Isa Set Padlock Ace"
	case ZydisIsaSetPadlockPhe:
		return "Zydis Isa Set Padlock Phe"
	case ZydisIsaSetPadlockPmm:
		return "Zydis Isa Set Padlock Pmm"
	case ZydisIsaSetPadlockRng:
		return "Zydis Isa Set Padlock Rng"
	case ZydisIsaSetPause:
		return "Zydis Isa Set Pause"
	case ZydisIsaSetPbndkb:
		return "Zydis Isa Set Pbndkb"
	case ZydisIsaSetPclmulqdq:
		return "Zydis Isa Set Pclmulqdq"
	case ZydisIsaSetPcommit:
		return "Zydis Isa Set Pcommit"
	case ZydisIsaSetPconfig:
		return "Zydis Isa Set Pconfig"
	case ZydisIsaSetPentiummmx:
		return "Zydis Isa Set Pentiummmx"
	case ZydisIsaSetPentiumreal:
		return "Zydis Isa Set Pentiumreal"
	case ZydisIsaSetPku:
		return "Zydis Isa Set Pku"
	case ZydisIsaSetPopcnt:
		return "Zydis Isa Set Popcnt"
	case ZydisIsaSetPpro:
		return "Zydis Isa Set Ppro"
	case ZydisIsaSetPrefetchwt1:
		return "Zydis Isa Set Prefetchwt 1"
	case ZydisIsaSetPrefetchNop:
		return "Zydis Isa Set Prefetch Nop"
	case ZydisIsaSetPt:
		return "Zydis Isa Set Pt"
	case ZydisIsaSetRaoInt:
		return "Zydis Isa Set Rao Int"
	case ZydisIsaSetRdpid:
		return "Zydis Isa Set Rdpid"
	case ZydisIsaSetRdpmc:
		return "Zydis Isa Set Rdpmc"
	case ZydisIsaSetRdpru:
		return "Zydis Isa Set Rdpru"
	case ZydisIsaSetRdrand:
		return "Zydis Isa Set Rdrand"
	case ZydisIsaSetRdseed:
		return "Zydis Isa Set Rdseed"
	case ZydisIsaSetRdtscp:
		return "Zydis Isa Set Rdtscp"
	case ZydisIsaSetRdwrfsgs:
		return "Zydis Isa Set Rdwrfsgs"
	case ZydisIsaSetRtm:
		return "Zydis Isa Set Rtm"
	case ZydisIsaSetSerialize:
		return "Zydis Isa Set Serialize"
	case ZydisIsaSetSgx:
		return "Zydis Isa Set Sgx"
	case ZydisIsaSetSgxEnclv:
		return "Zydis Isa Set Sgx Enclv"
	case ZydisIsaSetSha:
		return "Zydis Isa Set Sha"
	case ZydisIsaSetSha512:
		return "Zydis Isa Set Sha 512"
	case ZydisIsaSetSm3:
		return "Zydis Isa Set Sm 3"
	case ZydisIsaSetSm4:
		return "Zydis Isa Set Sm 4"
	case ZydisIsaSetSmap:
		return "Zydis Isa Set Smap"
	case ZydisIsaSetSmx:
		return "Zydis Isa Set Smx"
	case ZydisIsaSetSnp:
		return "Zydis Isa Set Snp"
	case ZydisIsaSetSse:
		return "Zydis Isa Set Sse"
	case ZydisIsaSetSse2:
		return "Zydis Isa Set Sse 2"
	case ZydisIsaSetSse2mmx:
		return "Zydis Isa Set Sse 2mmx"
	case ZydisIsaSetSse3:
		return "Zydis Isa Set Sse 3"
	case ZydisIsaSetSse3x87:
		return "Zydis Isa Set Sse 3x 87"
	case ZydisIsaSetSse4:
		return "Zydis Isa Set Sse 4"
	case ZydisIsaSetSse42:
		return "Zydis Isa Set Sse 42"
	case ZydisIsaSetSse4a:
		return "Zydis Isa Set Sse 4a"
	case ZydisIsaSetSsemxcsr:
		return "Zydis Isa Set Ssemxcsr"
	case ZydisIsaSetSsePrefetch:
		return "Zydis Isa Set Sse Prefetch"
	case ZydisIsaSetSsse3:
		return "Zydis Isa Set Ssse 3"
	case ZydisIsaSetSsse3mmx:
		return "Zydis Isa Set Ssse 3mmx"
	case ZydisIsaSetSvm:
		return "Zydis Isa Set Svm"
	case ZydisIsaSetTbm:
		return "Zydis Isa Set Tbm"
	case ZydisIsaSetTdx:
		return "Zydis Isa Set Tdx"
	case ZydisIsaSetTsxLdtrk:
		return "Zydis Isa Set Tsx Ldtrk"
	case ZydisIsaSetUintr:
		return "Zydis Isa Set Uintr"
	case ZydisIsaSetVaes:
		return "Zydis Isa Set Vaes"
	case ZydisIsaSetVmfunc:
		return "Zydis Isa Set Vmfunc"
	case ZydisIsaSetVpclmulqdq:
		return "Zydis Isa Set Vpclmulqdq"
	case ZydisIsaSetVtx:
		return "Zydis Isa Set Vtx"
	case ZydisIsaSetWaitpkg:
		return "Zydis Isa Set Waitpkg"
	case ZydisIsaSetWrmsrns:
		return "Zydis Isa Set Wrmsrns"
	case ZydisIsaSetX87:
		return "Zydis Isa Set X87"
	case ZydisIsaSetXop:
		return "Zydis Isa Set Xop"
	case ZydisIsaSetXsave:
		return "Zydis Isa Set Xsave"
	case ZydisIsaSetXsavec:
		return "Zydis Isa Set Xsavec"
	case ZydisIsaSetXsaveopt:
		return "Zydis Isa Set Xsaveopt"
	case ZydisIsaSetXsaves:
		return "Zydis Isa Set Xsaves"
	default:
		return fmt.Sprintf("ZydisISASet_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:4 -> ZydisISAExt_
type ZydisISAExt_ uint32

const (
	ZydisIsaExtInvalid ZydisISAExt_ = 0
	ZydisIsaExtAdoxAdcx ZydisISAExt_ = 1
	ZydisIsaExtAes ZydisISAExt_ = 2
	ZydisIsaExtAmd3dnow ZydisISAExt_ = 3
	ZydisIsaExtAmd3dnowPrefetch ZydisISAExt_ = 4
	ZydisIsaExtAmdInvlpgb ZydisISAExt_ = 5
	ZydisIsaExtAmxBf16 ZydisISAExt_ = 6
	ZydisIsaExtAmxFp16 ZydisISAExt_ = 7
	ZydisIsaExtAmxInt8 ZydisISAExt_ = 8
	ZydisIsaExtAmxTile ZydisISAExt_ = 9
	ZydisIsaExtApxevex ZydisISAExt_ = 10
	ZydisIsaExtApxlegacy ZydisISAExt_ = 11
	ZydisIsaExtAvx ZydisISAExt_ = 12
	ZydisIsaExtAvx2 ZydisISAExt_ = 13
	ZydisIsaExtAvx2gather ZydisISAExt_ = 14
	ZydisIsaExtAvx512evex ZydisISAExt_ = 15
	ZydisIsaExtAvx512vex ZydisISAExt_ = 16
	ZydisIsaExtAvxaes ZydisISAExt_ = 17
	ZydisIsaExtAvxIfma ZydisISAExt_ = 18
	ZydisIsaExtAvxNeConvert ZydisISAExt_ = 19
	ZydisIsaExtAvxVnni ZydisISAExt_ = 20
	ZydisIsaExtAvxVnniInt16 ZydisISAExt_ = 21
	ZydisIsaExtAvxVnniInt8 ZydisISAExt_ = 22
	ZydisIsaExtBase ZydisISAExt_ = 23
	ZydisIsaExtBmi1 ZydisISAExt_ = 24
	ZydisIsaExtBmi2 ZydisISAExt_ = 25
	ZydisIsaExtCet ZydisISAExt_ = 26
	ZydisIsaExtCldemote ZydisISAExt_ = 27
	ZydisIsaExtClflushopt ZydisISAExt_ = 28
	ZydisIsaExtClfsh ZydisISAExt_ = 29
	ZydisIsaExtClwb ZydisISAExt_ = 30
	ZydisIsaExtClzero ZydisISAExt_ = 31
	ZydisIsaExtEnqcmd ZydisISAExt_ = 32
	ZydisIsaExtF16c ZydisISAExt_ = 33
	ZydisIsaExtFma ZydisISAExt_ = 34
	ZydisIsaExtFma4 ZydisISAExt_ = 35
	ZydisIsaExtFred ZydisISAExt_ = 36
	ZydisIsaExtGfni ZydisISAExt_ = 37
	ZydisIsaExtHreset ZydisISAExt_ = 38
	ZydisIsaExtIcachePrefetch ZydisISAExt_ = 39
	ZydisIsaExtInvpcid ZydisISAExt_ = 40
	ZydisIsaExtKeylocker ZydisISAExt_ = 41
	ZydisIsaExtKeylockerWide ZydisISAExt_ = 42
	ZydisIsaExtKnc ZydisISAExt_ = 43
	ZydisIsaExtKnce ZydisISAExt_ = 44
	ZydisIsaExtKncv ZydisISAExt_ = 45
	ZydisIsaExtLkgs ZydisISAExt_ = 46
	ZydisIsaExtLongmode ZydisISAExt_ = 47
	ZydisIsaExtLzcnt ZydisISAExt_ = 48
	ZydisIsaExtMcommit ZydisISAExt_ = 49
	ZydisIsaExtMmx ZydisISAExt_ = 50
	ZydisIsaExtMonitor ZydisISAExt_ = 51
	ZydisIsaExtMonitorx ZydisISAExt_ = 52
	ZydisIsaExtMovbe ZydisISAExt_ = 53
	ZydisIsaExtMovdir ZydisISAExt_ = 54
	ZydisIsaExtMpx ZydisISAExt_ = 55
	ZydisIsaExtMsrlist ZydisISAExt_ = 56
	ZydisIsaExtPadlock ZydisISAExt_ = 57
	ZydisIsaExtPause ZydisISAExt_ = 58
	ZydisIsaExtPbndkb ZydisISAExt_ = 59
	ZydisIsaExtPclmulqdq ZydisISAExt_ = 60
	ZydisIsaExtPcommit ZydisISAExt_ = 61
	ZydisIsaExtPconfig ZydisISAExt_ = 62
	ZydisIsaExtPku ZydisISAExt_ = 63
	ZydisIsaExtPrefetchwt1 ZydisISAExt_ = 64
	ZydisIsaExtPt ZydisISAExt_ = 65
	ZydisIsaExtRaoInt ZydisISAExt_ = 66
	ZydisIsaExtRdpid ZydisISAExt_ = 67
	ZydisIsaExtRdpru ZydisISAExt_ = 68
	ZydisIsaExtRdrand ZydisISAExt_ = 69
	ZydisIsaExtRdseed ZydisISAExt_ = 70
	ZydisIsaExtRdtscp ZydisISAExt_ = 71
	ZydisIsaExtRdwrfsgs ZydisISAExt_ = 72
	ZydisIsaExtRtm ZydisISAExt_ = 73
	ZydisIsaExtSerialize ZydisISAExt_ = 74
	ZydisIsaExtSgx ZydisISAExt_ = 75
	ZydisIsaExtSgxEnclv ZydisISAExt_ = 76
	ZydisIsaExtSha ZydisISAExt_ = 77
	ZydisIsaExtSha512 ZydisISAExt_ = 78
	ZydisIsaExtSm3 ZydisISAExt_ = 79
	ZydisIsaExtSm4 ZydisISAExt_ = 80
	ZydisIsaExtSmap ZydisISAExt_ = 81
	ZydisIsaExtSmx ZydisISAExt_ = 82
	ZydisIsaExtSnp ZydisISAExt_ = 83
	ZydisIsaExtSse ZydisISAExt_ = 84
	ZydisIsaExtSse2 ZydisISAExt_ = 85
	ZydisIsaExtSse3 ZydisISAExt_ = 86
	ZydisIsaExtSse4 ZydisISAExt_ = 87
	ZydisIsaExtSse4a ZydisISAExt_ = 88
	ZydisIsaExtSsse3 ZydisISAExt_ = 89
	ZydisIsaExtSvm ZydisISAExt_ = 90
	ZydisIsaExtTbm ZydisISAExt_ = 91
	ZydisIsaExtTdx ZydisISAExt_ = 92
	ZydisIsaExtTsxLdtrk ZydisISAExt_ = 93
	ZydisIsaExtUintr ZydisISAExt_ = 94
	ZydisIsaExtVaes ZydisISAExt_ = 95
	ZydisIsaExtVmfunc ZydisISAExt_ = 96
	ZydisIsaExtVpclmulqdq ZydisISAExt_ = 97
	ZydisIsaExtVtx ZydisISAExt_ = 98
	ZydisIsaExtWaitpkg ZydisISAExt_ = 99
	ZydisIsaExtWrmsrns ZydisISAExt_ = 100
	ZydisIsaExtX87 ZydisISAExt_ = 101
	ZydisIsaExtXop ZydisISAExt_ = 102
	ZydisIsaExtXsave ZydisISAExt_ = 103
	ZydisIsaExtXsavec ZydisISAExt_ = 104
	ZydisIsaExtXsaveopt ZydisISAExt_ = 105
	ZydisIsaExtXsaves ZydisISAExt_ = 106
	ZydisIsaExtMaxValue ZydisISAExt_ = 106
	ZydisIsaExtRequiredBits ZydisISAExt_ = 7
)

func (z ZydisISAExt_) String() string {
	switch z {
	case ZydisIsaExtInvalid:
		return "Zydis Isa Ext Invalid"
	case ZydisIsaExtAdoxAdcx:
		return "Zydis Isa Ext Adox Adcx"
	case ZydisIsaExtAes:
		return "Zydis Isa Ext Aes"
	case ZydisIsaExtAmd3dnow:
		return "Zydis Isa Ext Amd 3dnow"
	case ZydisIsaExtAmd3dnowPrefetch:
		return "Zydis Isa Ext Amd 3dnow Prefetch"
	case ZydisIsaExtAmdInvlpgb:
		return "Zydis Isa Ext Amd Invlpgb"
	case ZydisIsaExtAmxBf16:
		return "Zydis Isa Ext Amx Bf 16"
	case ZydisIsaExtAmxFp16:
		return "Zydis Isa Ext Amx Fp 16"
	case ZydisIsaExtAmxInt8:
		return "Zydis Isa Ext Amx Int 8"
	case ZydisIsaExtAmxTile:
		return "Zydis Isa Ext Amx Tile"
	case ZydisIsaExtApxevex:
		return "Zydis Isa Ext Apxevex"
	case ZydisIsaExtApxlegacy:
		return "Zydis Isa Ext Apxlegacy"
	case ZydisIsaExtAvx:
		return "Zydis Isa Ext Avx"
	case ZydisIsaExtAvx2:
		return "Zydis Isa Ext Avx 2"
	case ZydisIsaExtAvx2gather:
		return "Zydis Isa Ext Avx 2gather"
	case ZydisIsaExtAvx512evex:
		return "Zydis Isa Ext Avx 512evex"
	case ZydisIsaExtAvx512vex:
		return "Zydis Isa Ext Avx 512vex"
	case ZydisIsaExtAvxaes:
		return "Zydis Isa Ext Avxaes"
	case ZydisIsaExtAvxIfma:
		return "Zydis Isa Ext Avx Ifma"
	case ZydisIsaExtAvxNeConvert:
		return "Zydis Isa Ext Avx Ne Convert"
	case ZydisIsaExtAvxVnni:
		return "Zydis Isa Ext Avx Vnni"
	case ZydisIsaExtAvxVnniInt16:
		return "Zydis Isa Ext Avx Vnni Int 16"
	case ZydisIsaExtAvxVnniInt8:
		return "Zydis Isa Ext Avx Vnni Int 8"
	case ZydisIsaExtBase:
		return "Zydis Isa Ext Base"
	case ZydisIsaExtBmi1:
		return "Zydis Isa Ext Bmi 1"
	case ZydisIsaExtBmi2:
		return "Zydis Isa Ext Bmi 2"
	case ZydisIsaExtCet:
		return "Zydis Isa Ext Cet"
	case ZydisIsaExtCldemote:
		return "Zydis Isa Ext Cldemote"
	case ZydisIsaExtClflushopt:
		return "Zydis Isa Ext Clflushopt"
	case ZydisIsaExtClfsh:
		return "Zydis Isa Ext Clfsh"
	case ZydisIsaExtClwb:
		return "Zydis Isa Ext Clwb"
	case ZydisIsaExtClzero:
		return "Zydis Isa Ext Clzero"
	case ZydisIsaExtEnqcmd:
		return "Zydis Isa Ext Enqcmd"
	case ZydisIsaExtF16c:
		return "Zydis Isa Ext F16c"
	case ZydisIsaExtFma:
		return "Zydis Isa Ext Fma"
	case ZydisIsaExtFma4:
		return "Zydis Isa Ext Fma 4"
	case ZydisIsaExtFred:
		return "Zydis Isa Ext Fred"
	case ZydisIsaExtGfni:
		return "Zydis Isa Ext Gfni"
	case ZydisIsaExtHreset:
		return "Zydis Isa Ext Hreset"
	case ZydisIsaExtIcachePrefetch:
		return "Zydis Isa Ext Icache Prefetch"
	case ZydisIsaExtInvpcid:
		return "Zydis Isa Ext Invpcid"
	case ZydisIsaExtKeylocker:
		return "Zydis Isa Ext Keylocker"
	case ZydisIsaExtKeylockerWide:
		return "Zydis Isa Ext Keylocker Wide"
	case ZydisIsaExtKnc:
		return "Zydis Isa Ext Knc"
	case ZydisIsaExtKnce:
		return "Zydis Isa Ext Knce"
	case ZydisIsaExtKncv:
		return "Zydis Isa Ext Kncv"
	case ZydisIsaExtLkgs:
		return "Zydis Isa Ext Lkgs"
	case ZydisIsaExtLongmode:
		return "Zydis Isa Ext Longmode"
	case ZydisIsaExtLzcnt:
		return "Zydis Isa Ext Lzcnt"
	case ZydisIsaExtMcommit:
		return "Zydis Isa Ext Mcommit"
	case ZydisIsaExtMmx:
		return "Zydis Isa Ext Mmx"
	case ZydisIsaExtMonitor:
		return "Zydis Isa Ext Monitor"
	case ZydisIsaExtMonitorx:
		return "Zydis Isa Ext Monitorx"
	case ZydisIsaExtMovbe:
		return "Zydis Isa Ext Movbe"
	case ZydisIsaExtMovdir:
		return "Zydis Isa Ext Movdir"
	case ZydisIsaExtMpx:
		return "Zydis Isa Ext Mpx"
	case ZydisIsaExtMsrlist:
		return "Zydis Isa Ext Msrlist"
	case ZydisIsaExtPadlock:
		return "Zydis Isa Ext Padlock"
	case ZydisIsaExtPause:
		return "Zydis Isa Ext Pause"
	case ZydisIsaExtPbndkb:
		return "Zydis Isa Ext Pbndkb"
	case ZydisIsaExtPclmulqdq:
		return "Zydis Isa Ext Pclmulqdq"
	case ZydisIsaExtPcommit:
		return "Zydis Isa Ext Pcommit"
	case ZydisIsaExtPconfig:
		return "Zydis Isa Ext Pconfig"
	case ZydisIsaExtPku:
		return "Zydis Isa Ext Pku"
	case ZydisIsaExtPrefetchwt1:
		return "Zydis Isa Ext Prefetchwt 1"
	case ZydisIsaExtPt:
		return "Zydis Isa Ext Pt"
	case ZydisIsaExtRaoInt:
		return "Zydis Isa Ext Rao Int"
	case ZydisIsaExtRdpid:
		return "Zydis Isa Ext Rdpid"
	case ZydisIsaExtRdpru:
		return "Zydis Isa Ext Rdpru"
	case ZydisIsaExtRdrand:
		return "Zydis Isa Ext Rdrand"
	case ZydisIsaExtRdseed:
		return "Zydis Isa Ext Rdseed"
	case ZydisIsaExtRdtscp:
		return "Zydis Isa Ext Rdtscp"
	case ZydisIsaExtRdwrfsgs:
		return "Zydis Isa Ext Rdwrfsgs"
	case ZydisIsaExtRtm:
		return "Zydis Isa Ext Rtm"
	case ZydisIsaExtSerialize:
		return "Zydis Isa Ext Serialize"
	case ZydisIsaExtSgx:
		return "Zydis Isa Ext Sgx"
	case ZydisIsaExtSgxEnclv:
		return "Zydis Isa Ext Sgx Enclv"
	case ZydisIsaExtSha:
		return "Zydis Isa Ext Sha"
	case ZydisIsaExtSha512:
		return "Zydis Isa Ext Sha 512"
	case ZydisIsaExtSm3:
		return "Zydis Isa Ext Sm 3"
	case ZydisIsaExtSm4:
		return "Zydis Isa Ext Sm 4"
	case ZydisIsaExtSmap:
		return "Zydis Isa Ext Smap"
	case ZydisIsaExtSmx:
		return "Zydis Isa Ext Smx"
	case ZydisIsaExtSnp:
		return "Zydis Isa Ext Snp"
	case ZydisIsaExtSse:
		return "Zydis Isa Ext Sse"
	case ZydisIsaExtSse2:
		return "Zydis Isa Ext Sse 2"
	case ZydisIsaExtSse3:
		return "Zydis Isa Ext Sse 3"
	case ZydisIsaExtSse4:
		return "Zydis Isa Ext Sse 4"
	case ZydisIsaExtSse4a:
		return "Zydis Isa Ext Sse 4a"
	case ZydisIsaExtSsse3:
		return "Zydis Isa Ext Ssse 3"
	case ZydisIsaExtSvm:
		return "Zydis Isa Ext Svm"
	case ZydisIsaExtTbm:
		return "Zydis Isa Ext Tbm"
	case ZydisIsaExtTdx:
		return "Zydis Isa Ext Tdx"
	case ZydisIsaExtTsxLdtrk:
		return "Zydis Isa Ext Tsx Ldtrk"
	case ZydisIsaExtUintr:
		return "Zydis Isa Ext Uintr"
	case ZydisIsaExtVaes:
		return "Zydis Isa Ext Vaes"
	case ZydisIsaExtVmfunc:
		return "Zydis Isa Ext Vmfunc"
	case ZydisIsaExtVpclmulqdq:
		return "Zydis Isa Ext Vpclmulqdq"
	case ZydisIsaExtVtx:
		return "Zydis Isa Ext Vtx"
	case ZydisIsaExtWaitpkg:
		return "Zydis Isa Ext Waitpkg"
	case ZydisIsaExtWrmsrns:
		return "Zydis Isa Ext Wrmsrns"
	case ZydisIsaExtX87:
		return "Zydis Isa Ext X87"
	case ZydisIsaExtXop:
		return "Zydis Isa Ext Xop"
	case ZydisIsaExtXsave:
		return "Zydis Isa Ext Xsave"
	case ZydisIsaExtXsavec:
		return "Zydis Isa Ext Xsavec"
	case ZydisIsaExtXsaveopt:
		return "Zydis Isa Ext Xsaveopt"
	case ZydisIsaExtXsaves:
		return "Zydis Isa Ext Xsaves"
	default:
		return fmt.Sprintf("ZydisISAExt_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:4 -> ZydisMnemonic_
type ZydisMnemonic_ uint32

const (
	ZydisMnemonicInvalid ZydisMnemonic_ = 0
	ZydisMnemonicAaa ZydisMnemonic_ = 1
	ZydisMnemonicAad ZydisMnemonic_ = 2
	ZydisMnemonicAadd ZydisMnemonic_ = 3
	ZydisMnemonicAam ZydisMnemonic_ = 4
	ZydisMnemonicAand ZydisMnemonic_ = 5
	ZydisMnemonicAas ZydisMnemonic_ = 6
	ZydisMnemonicAdc ZydisMnemonic_ = 7
	ZydisMnemonicAdcx ZydisMnemonic_ = 8
	ZydisMnemonicAdd ZydisMnemonic_ = 9
	ZydisMnemonicAddpd ZydisMnemonic_ = 10
	ZydisMnemonicAddps ZydisMnemonic_ = 11
	ZydisMnemonicAddsd ZydisMnemonic_ = 12
	ZydisMnemonicAddss ZydisMnemonic_ = 13
	ZydisMnemonicAddsubpd ZydisMnemonic_ = 14
	ZydisMnemonicAddsubps ZydisMnemonic_ = 15
	ZydisMnemonicAdox ZydisMnemonic_ = 16
	ZydisMnemonicAesdec ZydisMnemonic_ = 17
	ZydisMnemonicAesdec128kl ZydisMnemonic_ = 18
	ZydisMnemonicAesdec256kl ZydisMnemonic_ = 19
	ZydisMnemonicAesdeclast ZydisMnemonic_ = 20
	ZydisMnemonicAesdecwide128kl ZydisMnemonic_ = 21
	ZydisMnemonicAesdecwide256kl ZydisMnemonic_ = 22
	ZydisMnemonicAesenc ZydisMnemonic_ = 23
	ZydisMnemonicAesenc128kl ZydisMnemonic_ = 24
	ZydisMnemonicAesenc256kl ZydisMnemonic_ = 25
	ZydisMnemonicAesenclast ZydisMnemonic_ = 26
	ZydisMnemonicAesencwide128kl ZydisMnemonic_ = 27
	ZydisMnemonicAesencwide256kl ZydisMnemonic_ = 28
	ZydisMnemonicAesimc ZydisMnemonic_ = 29
	ZydisMnemonicAeskeygenassist ZydisMnemonic_ = 30
	ZydisMnemonicAnd ZydisMnemonic_ = 31
	ZydisMnemonicAndn ZydisMnemonic_ = 32
	ZydisMnemonicAndnpd ZydisMnemonic_ = 33
	ZydisMnemonicAndnps ZydisMnemonic_ = 34
	ZydisMnemonicAndpd ZydisMnemonic_ = 35
	ZydisMnemonicAndps ZydisMnemonic_ = 36
	ZydisMnemonicAor ZydisMnemonic_ = 37
	ZydisMnemonicArpl ZydisMnemonic_ = 38
	ZydisMnemonicAxor ZydisMnemonic_ = 39
	ZydisMnemonicBextr ZydisMnemonic_ = 40
	ZydisMnemonicBlcfill ZydisMnemonic_ = 41
	ZydisMnemonicBlci ZydisMnemonic_ = 42
	ZydisMnemonicBlcic ZydisMnemonic_ = 43
	ZydisMnemonicBlcmsk ZydisMnemonic_ = 44
	ZydisMnemonicBlcs ZydisMnemonic_ = 45
	ZydisMnemonicBlendpd ZydisMnemonic_ = 46
	ZydisMnemonicBlendps ZydisMnemonic_ = 47
	ZydisMnemonicBlendvpd ZydisMnemonic_ = 48
	ZydisMnemonicBlendvps ZydisMnemonic_ = 49
	ZydisMnemonicBlsfill ZydisMnemonic_ = 50
	ZydisMnemonicBlsi ZydisMnemonic_ = 51
	ZydisMnemonicBlsic ZydisMnemonic_ = 52
	ZydisMnemonicBlsmsk ZydisMnemonic_ = 53
	ZydisMnemonicBlsr ZydisMnemonic_ = 54
	ZydisMnemonicBndcl ZydisMnemonic_ = 55
	ZydisMnemonicBndcn ZydisMnemonic_ = 56
	ZydisMnemonicBndcu ZydisMnemonic_ = 57
	ZydisMnemonicBndldx ZydisMnemonic_ = 58
	ZydisMnemonicBndmk ZydisMnemonic_ = 59
	ZydisMnemonicBndmov ZydisMnemonic_ = 60
	ZydisMnemonicBndstx ZydisMnemonic_ = 61
	ZydisMnemonicBound ZydisMnemonic_ = 62
	ZydisMnemonicBsf ZydisMnemonic_ = 63
	ZydisMnemonicBsr ZydisMnemonic_ = 64
	ZydisMnemonicBswap ZydisMnemonic_ = 65
	ZydisMnemonicBt ZydisMnemonic_ = 66
	ZydisMnemonicBtc ZydisMnemonic_ = 67
	ZydisMnemonicBtr ZydisMnemonic_ = 68
	ZydisMnemonicBts ZydisMnemonic_ = 69
	ZydisMnemonicBzhi ZydisMnemonic_ = 70
	ZydisMnemonicCall ZydisMnemonic_ = 71
	ZydisMnemonicCbw ZydisMnemonic_ = 72
	ZydisMnemonicCcmpb ZydisMnemonic_ = 73
	ZydisMnemonicCcmpbe ZydisMnemonic_ = 74
	ZydisMnemonicCcmpf ZydisMnemonic_ = 75
	ZydisMnemonicCcmpl ZydisMnemonic_ = 76
	ZydisMnemonicCcmple ZydisMnemonic_ = 77
	ZydisMnemonicCcmpnb ZydisMnemonic_ = 78
	ZydisMnemonicCcmpnbe ZydisMnemonic_ = 79
	ZydisMnemonicCcmpnl ZydisMnemonic_ = 80
	ZydisMnemonicCcmpnle ZydisMnemonic_ = 81
	ZydisMnemonicCcmpno ZydisMnemonic_ = 82
	ZydisMnemonicCcmpns ZydisMnemonic_ = 83
	ZydisMnemonicCcmpnz ZydisMnemonic_ = 84
	ZydisMnemonicCcmpo ZydisMnemonic_ = 85
	ZydisMnemonicCcmps ZydisMnemonic_ = 86
	ZydisMnemonicCcmpt ZydisMnemonic_ = 87
	ZydisMnemonicCcmpz ZydisMnemonic_ = 88
	ZydisMnemonicCdq ZydisMnemonic_ = 89
	ZydisMnemonicCdqe ZydisMnemonic_ = 90
	ZydisMnemonicCfcmovb ZydisMnemonic_ = 91
	ZydisMnemonicCfcmovbe ZydisMnemonic_ = 92
	ZydisMnemonicCfcmovl ZydisMnemonic_ = 93
	ZydisMnemonicCfcmovle ZydisMnemonic_ = 94
	ZydisMnemonicCfcmovnb ZydisMnemonic_ = 95
	ZydisMnemonicCfcmovnbe ZydisMnemonic_ = 96
	ZydisMnemonicCfcmovnl ZydisMnemonic_ = 97
	ZydisMnemonicCfcmovnle ZydisMnemonic_ = 98
	ZydisMnemonicCfcmovno ZydisMnemonic_ = 99
	ZydisMnemonicCfcmovnp ZydisMnemonic_ = 100
	ZydisMnemonicCfcmovns ZydisMnemonic_ = 101
	ZydisMnemonicCfcmovnz ZydisMnemonic_ = 102
	ZydisMnemonicCfcmovo ZydisMnemonic_ = 103
	ZydisMnemonicCfcmovp ZydisMnemonic_ = 104
	ZydisMnemonicCfcmovs ZydisMnemonic_ = 105
	ZydisMnemonicCfcmovz ZydisMnemonic_ = 106
	ZydisMnemonicClac ZydisMnemonic_ = 107
	ZydisMnemonicClc ZydisMnemonic_ = 108
	ZydisMnemonicCld ZydisMnemonic_ = 109
	ZydisMnemonicCldemote ZydisMnemonic_ = 110
	ZydisMnemonicClevict0 ZydisMnemonic_ = 111
	ZydisMnemonicClevict1 ZydisMnemonic_ = 112
	ZydisMnemonicClflush ZydisMnemonic_ = 113
	ZydisMnemonicClflushopt ZydisMnemonic_ = 114
	ZydisMnemonicClgi ZydisMnemonic_ = 115
	ZydisMnemonicCli ZydisMnemonic_ = 116
	ZydisMnemonicClrssbsy ZydisMnemonic_ = 117
	ZydisMnemonicClts ZydisMnemonic_ = 118
	ZydisMnemonicClui ZydisMnemonic_ = 119
	ZydisMnemonicClwb ZydisMnemonic_ = 120
	ZydisMnemonicClzero ZydisMnemonic_ = 121
	ZydisMnemonicCmc ZydisMnemonic_ = 122
	ZydisMnemonicCmovb ZydisMnemonic_ = 123
	ZydisMnemonicCmovbe ZydisMnemonic_ = 124
	ZydisMnemonicCmovl ZydisMnemonic_ = 125
	ZydisMnemonicCmovle ZydisMnemonic_ = 126
	ZydisMnemonicCmovnb ZydisMnemonic_ = 127
	ZydisMnemonicCmovnbe ZydisMnemonic_ = 128
	ZydisMnemonicCmovnl ZydisMnemonic_ = 129
	ZydisMnemonicCmovnle ZydisMnemonic_ = 130
	ZydisMnemonicCmovno ZydisMnemonic_ = 131
	ZydisMnemonicCmovnp ZydisMnemonic_ = 132
	ZydisMnemonicCmovns ZydisMnemonic_ = 133
	ZydisMnemonicCmovnz ZydisMnemonic_ = 134
	ZydisMnemonicCmovo ZydisMnemonic_ = 135
	ZydisMnemonicCmovp ZydisMnemonic_ = 136
	ZydisMnemonicCmovs ZydisMnemonic_ = 137
	ZydisMnemonicCmovz ZydisMnemonic_ = 138
	ZydisMnemonicCmp ZydisMnemonic_ = 139
	ZydisMnemonicCmpbexadd ZydisMnemonic_ = 140
	ZydisMnemonicCmpbxadd ZydisMnemonic_ = 141
	ZydisMnemonicCmplexadd ZydisMnemonic_ = 142
	ZydisMnemonicCmplxadd ZydisMnemonic_ = 143
	ZydisMnemonicCmpnbexadd ZydisMnemonic_ = 144
	ZydisMnemonicCmpnbxadd ZydisMnemonic_ = 145
	ZydisMnemonicCmpnlexadd ZydisMnemonic_ = 146
	ZydisMnemonicCmpnlxadd ZydisMnemonic_ = 147
	ZydisMnemonicCmpnoxadd ZydisMnemonic_ = 148
	ZydisMnemonicCmpnpxadd ZydisMnemonic_ = 149
	ZydisMnemonicCmpnsxadd ZydisMnemonic_ = 150
	ZydisMnemonicCmpnzxadd ZydisMnemonic_ = 151
	ZydisMnemonicCmpoxadd ZydisMnemonic_ = 152
	ZydisMnemonicCmppd ZydisMnemonic_ = 153
	ZydisMnemonicCmpps ZydisMnemonic_ = 154
	ZydisMnemonicCmppxadd ZydisMnemonic_ = 155
	ZydisMnemonicCmpsb ZydisMnemonic_ = 156
	ZydisMnemonicCmpsd ZydisMnemonic_ = 157
	ZydisMnemonicCmpsq ZydisMnemonic_ = 158
	ZydisMnemonicCmpss ZydisMnemonic_ = 159
	ZydisMnemonicCmpsw ZydisMnemonic_ = 160
	ZydisMnemonicCmpsxadd ZydisMnemonic_ = 161
	ZydisMnemonicCmpxchg ZydisMnemonic_ = 162
	ZydisMnemonicCmpxchg16b ZydisMnemonic_ = 163
	ZydisMnemonicCmpxchg8b ZydisMnemonic_ = 164
	ZydisMnemonicCmpzxadd ZydisMnemonic_ = 165
	ZydisMnemonicComisd ZydisMnemonic_ = 166
	ZydisMnemonicComiss ZydisMnemonic_ = 167
	ZydisMnemonicCpuid ZydisMnemonic_ = 168
	ZydisMnemonicCqo ZydisMnemonic_ = 169
	ZydisMnemonicCrc32 ZydisMnemonic_ = 170
	ZydisMnemonicCtestb ZydisMnemonic_ = 171
	ZydisMnemonicCtestbe ZydisMnemonic_ = 172
	ZydisMnemonicCtestf ZydisMnemonic_ = 173
	ZydisMnemonicCtestl ZydisMnemonic_ = 174
	ZydisMnemonicCtestle ZydisMnemonic_ = 175
	ZydisMnemonicCtestnb ZydisMnemonic_ = 176
	ZydisMnemonicCtestnbe ZydisMnemonic_ = 177
	ZydisMnemonicCtestnl ZydisMnemonic_ = 178
	ZydisMnemonicCtestnle ZydisMnemonic_ = 179
	ZydisMnemonicCtestno ZydisMnemonic_ = 180
	ZydisMnemonicCtestns ZydisMnemonic_ = 181
	ZydisMnemonicCtestnz ZydisMnemonic_ = 182
	ZydisMnemonicCtesto ZydisMnemonic_ = 183
	ZydisMnemonicCtests ZydisMnemonic_ = 184
	ZydisMnemonicCtestt ZydisMnemonic_ = 185
	ZydisMnemonicCtestz ZydisMnemonic_ = 186
	ZydisMnemonicCvtdq2pd ZydisMnemonic_ = 187
	ZydisMnemonicCvtdq2ps ZydisMnemonic_ = 188
	ZydisMnemonicCvtpd2dq ZydisMnemonic_ = 189
	ZydisMnemonicCvtpd2pi ZydisMnemonic_ = 190
	ZydisMnemonicCvtpd2ps ZydisMnemonic_ = 191
	ZydisMnemonicCvtpi2pd ZydisMnemonic_ = 192
	ZydisMnemonicCvtpi2ps ZydisMnemonic_ = 193
	ZydisMnemonicCvtps2dq ZydisMnemonic_ = 194
	ZydisMnemonicCvtps2pd ZydisMnemonic_ = 195
	ZydisMnemonicCvtps2pi ZydisMnemonic_ = 196
	ZydisMnemonicCvtsd2si ZydisMnemonic_ = 197
	ZydisMnemonicCvtsd2ss ZydisMnemonic_ = 198
	ZydisMnemonicCvtsi2sd ZydisMnemonic_ = 199
	ZydisMnemonicCvtsi2ss ZydisMnemonic_ = 200
	ZydisMnemonicCvtss2sd ZydisMnemonic_ = 201
	ZydisMnemonicCvtss2si ZydisMnemonic_ = 202
	ZydisMnemonicCvttpd2dq ZydisMnemonic_ = 203
	ZydisMnemonicCvttpd2pi ZydisMnemonic_ = 204
	ZydisMnemonicCvttps2dq ZydisMnemonic_ = 205
	ZydisMnemonicCvttps2pi ZydisMnemonic_ = 206
	ZydisMnemonicCvttsd2si ZydisMnemonic_ = 207
	ZydisMnemonicCvttss2si ZydisMnemonic_ = 208
	ZydisMnemonicCwd ZydisMnemonic_ = 209
	ZydisMnemonicCwde ZydisMnemonic_ = 210
	ZydisMnemonicDaa ZydisMnemonic_ = 211
	ZydisMnemonicDas ZydisMnemonic_ = 212
	ZydisMnemonicDec ZydisMnemonic_ = 213
	ZydisMnemonicDelay ZydisMnemonic_ = 214
	ZydisMnemonicDiv ZydisMnemonic_ = 215
	ZydisMnemonicDivpd ZydisMnemonic_ = 216
	ZydisMnemonicDivps ZydisMnemonic_ = 217
	ZydisMnemonicDivsd ZydisMnemonic_ = 218
	ZydisMnemonicDivss ZydisMnemonic_ = 219
	ZydisMnemonicDppd ZydisMnemonic_ = 220
	ZydisMnemonicDpps ZydisMnemonic_ = 221
	ZydisMnemonicEmms ZydisMnemonic_ = 222
	ZydisMnemonicEncls ZydisMnemonic_ = 223
	ZydisMnemonicEnclu ZydisMnemonic_ = 224
	ZydisMnemonicEnclv ZydisMnemonic_ = 225
	ZydisMnemonicEncodekey128 ZydisMnemonic_ = 226
	ZydisMnemonicEncodekey256 ZydisMnemonic_ = 227
	ZydisMnemonicEndbr32 ZydisMnemonic_ = 228
	ZydisMnemonicEndbr64 ZydisMnemonic_ = 229
	ZydisMnemonicEnqcmd ZydisMnemonic_ = 230
	ZydisMnemonicEnqcmds ZydisMnemonic_ = 231
	ZydisMnemonicEnter ZydisMnemonic_ = 232
	ZydisMnemonicErets ZydisMnemonic_ = 233
	ZydisMnemonicEretu ZydisMnemonic_ = 234
	ZydisMnemonicExtractps ZydisMnemonic_ = 235
	ZydisMnemonicExtrq ZydisMnemonic_ = 236
	ZydisMnemonicF2xm1 ZydisMnemonic_ = 237
	ZydisMnemonicFabs ZydisMnemonic_ = 238
	ZydisMnemonicFadd ZydisMnemonic_ = 239
	ZydisMnemonicFaddp ZydisMnemonic_ = 240
	ZydisMnemonicFbld ZydisMnemonic_ = 241
	ZydisMnemonicFbstp ZydisMnemonic_ = 242
	ZydisMnemonicFchs ZydisMnemonic_ = 243
	ZydisMnemonicFcmovb ZydisMnemonic_ = 244
	ZydisMnemonicFcmovbe ZydisMnemonic_ = 245
	ZydisMnemonicFcmove ZydisMnemonic_ = 246
	ZydisMnemonicFcmovnb ZydisMnemonic_ = 247
	ZydisMnemonicFcmovnbe ZydisMnemonic_ = 248
	ZydisMnemonicFcmovne ZydisMnemonic_ = 249
	ZydisMnemonicFcmovnu ZydisMnemonic_ = 250
	ZydisMnemonicFcmovu ZydisMnemonic_ = 251
	ZydisMnemonicFcom ZydisMnemonic_ = 252
	ZydisMnemonicFcomi ZydisMnemonic_ = 253
	ZydisMnemonicFcomip ZydisMnemonic_ = 254
	ZydisMnemonicFcomp ZydisMnemonic_ = 255
	ZydisMnemonicFcompp ZydisMnemonic_ = 256
	ZydisMnemonicFcos ZydisMnemonic_ = 257
	ZydisMnemonicFdecstp ZydisMnemonic_ = 258
	ZydisMnemonicFdisi8087Nop ZydisMnemonic_ = 259
	ZydisMnemonicFdiv ZydisMnemonic_ = 260
	ZydisMnemonicFdivp ZydisMnemonic_ = 261
	ZydisMnemonicFdivr ZydisMnemonic_ = 262
	ZydisMnemonicFdivrp ZydisMnemonic_ = 263
	ZydisMnemonicFemms ZydisMnemonic_ = 264
	ZydisMnemonicFeni8087Nop ZydisMnemonic_ = 265
	ZydisMnemonicFfree ZydisMnemonic_ = 266
	ZydisMnemonicFfreep ZydisMnemonic_ = 267
	ZydisMnemonicFiadd ZydisMnemonic_ = 268
	ZydisMnemonicFicom ZydisMnemonic_ = 269
	ZydisMnemonicFicomp ZydisMnemonic_ = 270
	ZydisMnemonicFidiv ZydisMnemonic_ = 271
	ZydisMnemonicFidivr ZydisMnemonic_ = 272
	ZydisMnemonicFild ZydisMnemonic_ = 273
	ZydisMnemonicFimul ZydisMnemonic_ = 274
	ZydisMnemonicFincstp ZydisMnemonic_ = 275
	ZydisMnemonicFist ZydisMnemonic_ = 276
	ZydisMnemonicFistp ZydisMnemonic_ = 277
	ZydisMnemonicFisttp ZydisMnemonic_ = 278
	ZydisMnemonicFisub ZydisMnemonic_ = 279
	ZydisMnemonicFisubr ZydisMnemonic_ = 280
	ZydisMnemonicFld ZydisMnemonic_ = 281
	ZydisMnemonicFld1 ZydisMnemonic_ = 282
	ZydisMnemonicFldcw ZydisMnemonic_ = 283
	ZydisMnemonicFldenv ZydisMnemonic_ = 284
	ZydisMnemonicFldl2e ZydisMnemonic_ = 285
	ZydisMnemonicFldl2t ZydisMnemonic_ = 286
	ZydisMnemonicFldlg2 ZydisMnemonic_ = 287
	ZydisMnemonicFldln2 ZydisMnemonic_ = 288
	ZydisMnemonicFldpi ZydisMnemonic_ = 289
	ZydisMnemonicFldz ZydisMnemonic_ = 290
	ZydisMnemonicFmul ZydisMnemonic_ = 291
	ZydisMnemonicFmulp ZydisMnemonic_ = 292
	ZydisMnemonicFnclex ZydisMnemonic_ = 293
	ZydisMnemonicFninit ZydisMnemonic_ = 294
	ZydisMnemonicFnop ZydisMnemonic_ = 295
	ZydisMnemonicFnsave ZydisMnemonic_ = 296
	ZydisMnemonicFnstcw ZydisMnemonic_ = 297
	ZydisMnemonicFnstenv ZydisMnemonic_ = 298
	ZydisMnemonicFnstsw ZydisMnemonic_ = 299
	ZydisMnemonicFpatan ZydisMnemonic_ = 300
	ZydisMnemonicFprem ZydisMnemonic_ = 301
	ZydisMnemonicFprem1 ZydisMnemonic_ = 302
	ZydisMnemonicFptan ZydisMnemonic_ = 303
	ZydisMnemonicFrndint ZydisMnemonic_ = 304
	ZydisMnemonicFrstor ZydisMnemonic_ = 305
	ZydisMnemonicFscale ZydisMnemonic_ = 306
	ZydisMnemonicFsetpm287Nop ZydisMnemonic_ = 307
	ZydisMnemonicFsin ZydisMnemonic_ = 308
	ZydisMnemonicFsincos ZydisMnemonic_ = 309
	ZydisMnemonicFsqrt ZydisMnemonic_ = 310
	ZydisMnemonicFst ZydisMnemonic_ = 311
	ZydisMnemonicFstp ZydisMnemonic_ = 312
	ZydisMnemonicFstpnce ZydisMnemonic_ = 313
	ZydisMnemonicFsub ZydisMnemonic_ = 314
	ZydisMnemonicFsubp ZydisMnemonic_ = 315
	ZydisMnemonicFsubr ZydisMnemonic_ = 316
	ZydisMnemonicFsubrp ZydisMnemonic_ = 317
	ZydisMnemonicFtst ZydisMnemonic_ = 318
	ZydisMnemonicFucom ZydisMnemonic_ = 319
	ZydisMnemonicFucomi ZydisMnemonic_ = 320
	ZydisMnemonicFucomip ZydisMnemonic_ = 321
	ZydisMnemonicFucomp ZydisMnemonic_ = 322
	ZydisMnemonicFucompp ZydisMnemonic_ = 323
	ZydisMnemonicFwait ZydisMnemonic_ = 324
	ZydisMnemonicFxam ZydisMnemonic_ = 325
	ZydisMnemonicFxch ZydisMnemonic_ = 326
	ZydisMnemonicFxrstor ZydisMnemonic_ = 327
	ZydisMnemonicFxrstor64 ZydisMnemonic_ = 328
	ZydisMnemonicFxsave ZydisMnemonic_ = 329
	ZydisMnemonicFxsave64 ZydisMnemonic_ = 330
	ZydisMnemonicFxtract ZydisMnemonic_ = 331
	ZydisMnemonicFyl2x ZydisMnemonic_ = 332
	ZydisMnemonicFyl2xp1 ZydisMnemonic_ = 333
	ZydisMnemonicGetsec ZydisMnemonic_ = 334
	ZydisMnemonicGf2p8affineinvqb ZydisMnemonic_ = 335
	ZydisMnemonicGf2p8affineqb ZydisMnemonic_ = 336
	ZydisMnemonicGf2p8mulb ZydisMnemonic_ = 337
	ZydisMnemonicHaddpd ZydisMnemonic_ = 338
	ZydisMnemonicHaddps ZydisMnemonic_ = 339
	ZydisMnemonicHlt ZydisMnemonic_ = 340
	ZydisMnemonicHreset ZydisMnemonic_ = 341
	ZydisMnemonicHsubpd ZydisMnemonic_ = 342
	ZydisMnemonicHsubps ZydisMnemonic_ = 343
	ZydisMnemonicIdiv ZydisMnemonic_ = 344
	ZydisMnemonicImul ZydisMnemonic_ = 345
	ZydisMnemonicImulzu ZydisMnemonic_ = 346
	ZydisMnemonicIn ZydisMnemonic_ = 347
	ZydisMnemonicInc ZydisMnemonic_ = 348
	ZydisMnemonicIncsspd ZydisMnemonic_ = 349
	ZydisMnemonicIncsspq ZydisMnemonic_ = 350
	ZydisMnemonicInsb ZydisMnemonic_ = 351
	ZydisMnemonicInsd ZydisMnemonic_ = 352
	ZydisMnemonicInsertps ZydisMnemonic_ = 353
	ZydisMnemonicInsertq ZydisMnemonic_ = 354
	ZydisMnemonicInsw ZydisMnemonic_ = 355
	ZydisMnemonicInt ZydisMnemonic_ = 356
	ZydisMnemonicInt1 ZydisMnemonic_ = 357
	ZydisMnemonicInt3 ZydisMnemonic_ = 358
	ZydisMnemonicInto ZydisMnemonic_ = 359
	ZydisMnemonicInvd ZydisMnemonic_ = 360
	ZydisMnemonicInvept ZydisMnemonic_ = 361
	ZydisMnemonicInvlpg ZydisMnemonic_ = 362
	ZydisMnemonicInvlpga ZydisMnemonic_ = 363
	ZydisMnemonicInvlpgb ZydisMnemonic_ = 364
	ZydisMnemonicInvpcid ZydisMnemonic_ = 365
	ZydisMnemonicInvvpid ZydisMnemonic_ = 366
	ZydisMnemonicIret ZydisMnemonic_ = 367
	ZydisMnemonicIretd ZydisMnemonic_ = 368
	ZydisMnemonicIretq ZydisMnemonic_ = 369
	ZydisMnemonicJb ZydisMnemonic_ = 370
	ZydisMnemonicJbe ZydisMnemonic_ = 371
	ZydisMnemonicJcxz ZydisMnemonic_ = 372
	ZydisMnemonicJecxz ZydisMnemonic_ = 373
	ZydisMnemonicJknzd ZydisMnemonic_ = 374
	ZydisMnemonicJkzd ZydisMnemonic_ = 375
	ZydisMnemonicJl ZydisMnemonic_ = 376
	ZydisMnemonicJle ZydisMnemonic_ = 377
	ZydisMnemonicJmp ZydisMnemonic_ = 378
	ZydisMnemonicJmpabs ZydisMnemonic_ = 379
	ZydisMnemonicJnb ZydisMnemonic_ = 380
	ZydisMnemonicJnbe ZydisMnemonic_ = 381
	ZydisMnemonicJnl ZydisMnemonic_ = 382
	ZydisMnemonicJnle ZydisMnemonic_ = 383
	ZydisMnemonicJno ZydisMnemonic_ = 384
	ZydisMnemonicJnp ZydisMnemonic_ = 385
	ZydisMnemonicJns ZydisMnemonic_ = 386
	ZydisMnemonicJnz ZydisMnemonic_ = 387
	ZydisMnemonicJo ZydisMnemonic_ = 388
	ZydisMnemonicJp ZydisMnemonic_ = 389
	ZydisMnemonicJrcxz ZydisMnemonic_ = 390
	ZydisMnemonicJs ZydisMnemonic_ = 391
	ZydisMnemonicJz ZydisMnemonic_ = 392
	ZydisMnemonicKaddb ZydisMnemonic_ = 393
	ZydisMnemonicKaddd ZydisMnemonic_ = 394
	ZydisMnemonicKaddq ZydisMnemonic_ = 395
	ZydisMnemonicKaddw ZydisMnemonic_ = 396
	ZydisMnemonicKand ZydisMnemonic_ = 397
	ZydisMnemonicKandb ZydisMnemonic_ = 398
	ZydisMnemonicKandd ZydisMnemonic_ = 399
	ZydisMnemonicKandn ZydisMnemonic_ = 400
	ZydisMnemonicKandnb ZydisMnemonic_ = 401
	ZydisMnemonicKandnd ZydisMnemonic_ = 402
	ZydisMnemonicKandnq ZydisMnemonic_ = 403
	ZydisMnemonicKandnr ZydisMnemonic_ = 404
	ZydisMnemonicKandnw ZydisMnemonic_ = 405
	ZydisMnemonicKandq ZydisMnemonic_ = 406
	ZydisMnemonicKandw ZydisMnemonic_ = 407
	ZydisMnemonicKconcath ZydisMnemonic_ = 408
	ZydisMnemonicKconcatl ZydisMnemonic_ = 409
	ZydisMnemonicKextract ZydisMnemonic_ = 410
	ZydisMnemonicKmerge2l1h ZydisMnemonic_ = 411
	ZydisMnemonicKmerge2l1l ZydisMnemonic_ = 412
	ZydisMnemonicKmov ZydisMnemonic_ = 413
	ZydisMnemonicKmovb ZydisMnemonic_ = 414
	ZydisMnemonicKmovd ZydisMnemonic_ = 415
	ZydisMnemonicKmovq ZydisMnemonic_ = 416
	ZydisMnemonicKmovw ZydisMnemonic_ = 417
	ZydisMnemonicKnot ZydisMnemonic_ = 418
	ZydisMnemonicKnotb ZydisMnemonic_ = 419
	ZydisMnemonicKnotd ZydisMnemonic_ = 420
	ZydisMnemonicKnotq ZydisMnemonic_ = 421
	ZydisMnemonicKnotw ZydisMnemonic_ = 422
	ZydisMnemonicKor ZydisMnemonic_ = 423
	ZydisMnemonicKorb ZydisMnemonic_ = 424
	ZydisMnemonicKord ZydisMnemonic_ = 425
	ZydisMnemonicKorq ZydisMnemonic_ = 426
	ZydisMnemonicKortest ZydisMnemonic_ = 427
	ZydisMnemonicKortestb ZydisMnemonic_ = 428
	ZydisMnemonicKortestd ZydisMnemonic_ = 429
	ZydisMnemonicKortestq ZydisMnemonic_ = 430
	ZydisMnemonicKortestw ZydisMnemonic_ = 431
	ZydisMnemonicKorw ZydisMnemonic_ = 432
	ZydisMnemonicKshiftlb ZydisMnemonic_ = 433
	ZydisMnemonicKshiftld ZydisMnemonic_ = 434
	ZydisMnemonicKshiftlq ZydisMnemonic_ = 435
	ZydisMnemonicKshiftlw ZydisMnemonic_ = 436
	ZydisMnemonicKshiftrb ZydisMnemonic_ = 437
	ZydisMnemonicKshiftrd ZydisMnemonic_ = 438
	ZydisMnemonicKshiftrq ZydisMnemonic_ = 439
	ZydisMnemonicKshiftrw ZydisMnemonic_ = 440
	ZydisMnemonicKtestb ZydisMnemonic_ = 441
	ZydisMnemonicKtestd ZydisMnemonic_ = 442
	ZydisMnemonicKtestq ZydisMnemonic_ = 443
	ZydisMnemonicKtestw ZydisMnemonic_ = 444
	ZydisMnemonicKunpckbw ZydisMnemonic_ = 445
	ZydisMnemonicKunpckdq ZydisMnemonic_ = 446
	ZydisMnemonicKunpckwd ZydisMnemonic_ = 447
	ZydisMnemonicKxnor ZydisMnemonic_ = 448
	ZydisMnemonicKxnorb ZydisMnemonic_ = 449
	ZydisMnemonicKxnord ZydisMnemonic_ = 450
	ZydisMnemonicKxnorq ZydisMnemonic_ = 451
	ZydisMnemonicKxnorw ZydisMnemonic_ = 452
	ZydisMnemonicKxor ZydisMnemonic_ = 453
	ZydisMnemonicKxorb ZydisMnemonic_ = 454
	ZydisMnemonicKxord ZydisMnemonic_ = 455
	ZydisMnemonicKxorq ZydisMnemonic_ = 456
	ZydisMnemonicKxorw ZydisMnemonic_ = 457
	ZydisMnemonicLahf ZydisMnemonic_ = 458
	ZydisMnemonicLar ZydisMnemonic_ = 459
	ZydisMnemonicLddqu ZydisMnemonic_ = 460
	ZydisMnemonicLdmxcsr ZydisMnemonic_ = 461
	ZydisMnemonicLds ZydisMnemonic_ = 462
	ZydisMnemonicLdtilecfg ZydisMnemonic_ = 463
	ZydisMnemonicLea ZydisMnemonic_ = 464
	ZydisMnemonicLeave ZydisMnemonic_ = 465
	ZydisMnemonicLes ZydisMnemonic_ = 466
	ZydisMnemonicLfence ZydisMnemonic_ = 467
	ZydisMnemonicLfs ZydisMnemonic_ = 468
	ZydisMnemonicLgdt ZydisMnemonic_ = 469
	ZydisMnemonicLgs ZydisMnemonic_ = 470
	ZydisMnemonicLidt ZydisMnemonic_ = 471
	ZydisMnemonicLkgs ZydisMnemonic_ = 472
	ZydisMnemonicLldt ZydisMnemonic_ = 473
	ZydisMnemonicLlwpcb ZydisMnemonic_ = 474
	ZydisMnemonicLmsw ZydisMnemonic_ = 475
	ZydisMnemonicLoadiwkey ZydisMnemonic_ = 476
	ZydisMnemonicLodsb ZydisMnemonic_ = 477
	ZydisMnemonicLodsd ZydisMnemonic_ = 478
	ZydisMnemonicLodsq ZydisMnemonic_ = 479
	ZydisMnemonicLodsw ZydisMnemonic_ = 480
	ZydisMnemonicLoop ZydisMnemonic_ = 481
	ZydisMnemonicLoope ZydisMnemonic_ = 482
	ZydisMnemonicLoopne ZydisMnemonic_ = 483
	ZydisMnemonicLsl ZydisMnemonic_ = 484
	ZydisMnemonicLss ZydisMnemonic_ = 485
	ZydisMnemonicLtr ZydisMnemonic_ = 486
	ZydisMnemonicLwpins ZydisMnemonic_ = 487
	ZydisMnemonicLwpval ZydisMnemonic_ = 488
	ZydisMnemonicLzcnt ZydisMnemonic_ = 489
	ZydisMnemonicMaskmovdqu ZydisMnemonic_ = 490
	ZydisMnemonicMaskmovq ZydisMnemonic_ = 491
	ZydisMnemonicMaxpd ZydisMnemonic_ = 492
	ZydisMnemonicMaxps ZydisMnemonic_ = 493
	ZydisMnemonicMaxsd ZydisMnemonic_ = 494
	ZydisMnemonicMaxss ZydisMnemonic_ = 495
	ZydisMnemonicMcommit ZydisMnemonic_ = 496
	ZydisMnemonicMfence ZydisMnemonic_ = 497
	ZydisMnemonicMinpd ZydisMnemonic_ = 498
	ZydisMnemonicMinps ZydisMnemonic_ = 499
	ZydisMnemonicMinsd ZydisMnemonic_ = 500
	ZydisMnemonicMinss ZydisMnemonic_ = 501
	ZydisMnemonicMonitor ZydisMnemonic_ = 502
	ZydisMnemonicMonitorx ZydisMnemonic_ = 503
	ZydisMnemonicMontmul ZydisMnemonic_ = 504
	ZydisMnemonicMov ZydisMnemonic_ = 505
	ZydisMnemonicMovapd ZydisMnemonic_ = 506
	ZydisMnemonicMovaps ZydisMnemonic_ = 507
	ZydisMnemonicMovbe ZydisMnemonic_ = 508
	ZydisMnemonicMovd ZydisMnemonic_ = 509
	ZydisMnemonicMovddup ZydisMnemonic_ = 510
	ZydisMnemonicMovdir64b ZydisMnemonic_ = 511
	ZydisMnemonicMovdiri ZydisMnemonic_ = 512
	ZydisMnemonicMovdq2q ZydisMnemonic_ = 513
	ZydisMnemonicMovdqa ZydisMnemonic_ = 514
	ZydisMnemonicMovdqu ZydisMnemonic_ = 515
	ZydisMnemonicMovhlps ZydisMnemonic_ = 516
	ZydisMnemonicMovhpd ZydisMnemonic_ = 517
	ZydisMnemonicMovhps ZydisMnemonic_ = 518
	ZydisMnemonicMovlhps ZydisMnemonic_ = 519
	ZydisMnemonicMovlpd ZydisMnemonic_ = 520
	ZydisMnemonicMovlps ZydisMnemonic_ = 521
	ZydisMnemonicMovmskpd ZydisMnemonic_ = 522
	ZydisMnemonicMovmskps ZydisMnemonic_ = 523
	ZydisMnemonicMovntdq ZydisMnemonic_ = 524
	ZydisMnemonicMovntdqa ZydisMnemonic_ = 525
	ZydisMnemonicMovnti ZydisMnemonic_ = 526
	ZydisMnemonicMovntpd ZydisMnemonic_ = 527
	ZydisMnemonicMovntps ZydisMnemonic_ = 528
	ZydisMnemonicMovntq ZydisMnemonic_ = 529
	ZydisMnemonicMovntsd ZydisMnemonic_ = 530
	ZydisMnemonicMovntss ZydisMnemonic_ = 531
	ZydisMnemonicMovq ZydisMnemonic_ = 532
	ZydisMnemonicMovq2dq ZydisMnemonic_ = 533
	ZydisMnemonicMovsb ZydisMnemonic_ = 534
	ZydisMnemonicMovsd ZydisMnemonic_ = 535
	ZydisMnemonicMovshdup ZydisMnemonic_ = 536
	ZydisMnemonicMovsldup ZydisMnemonic_ = 537
	ZydisMnemonicMovsq ZydisMnemonic_ = 538
	ZydisMnemonicMovss ZydisMnemonic_ = 539
	ZydisMnemonicMovsw ZydisMnemonic_ = 540
	ZydisMnemonicMovsx ZydisMnemonic_ = 541
	ZydisMnemonicMovsxd ZydisMnemonic_ = 542
	ZydisMnemonicMovupd ZydisMnemonic_ = 543
	ZydisMnemonicMovups ZydisMnemonic_ = 544
	ZydisMnemonicMovzx ZydisMnemonic_ = 545
	ZydisMnemonicMpsadbw ZydisMnemonic_ = 546
	ZydisMnemonicMul ZydisMnemonic_ = 547
	ZydisMnemonicMulpd ZydisMnemonic_ = 548
	ZydisMnemonicMulps ZydisMnemonic_ = 549
	ZydisMnemonicMulsd ZydisMnemonic_ = 550
	ZydisMnemonicMulss ZydisMnemonic_ = 551
	ZydisMnemonicMulx ZydisMnemonic_ = 552
	ZydisMnemonicMwait ZydisMnemonic_ = 553
	ZydisMnemonicMwaitx ZydisMnemonic_ = 554
	ZydisMnemonicNeg ZydisMnemonic_ = 555
	ZydisMnemonicNop ZydisMnemonic_ = 556
	ZydisMnemonicNot ZydisMnemonic_ = 557
	ZydisMnemonicOr ZydisMnemonic_ = 558
	ZydisMnemonicOrpd ZydisMnemonic_ = 559
	ZydisMnemonicOrps ZydisMnemonic_ = 560
	ZydisMnemonicOut ZydisMnemonic_ = 561
	ZydisMnemonicOutsb ZydisMnemonic_ = 562
	ZydisMnemonicOutsd ZydisMnemonic_ = 563
	ZydisMnemonicOutsw ZydisMnemonic_ = 564
	ZydisMnemonicPabsb ZydisMnemonic_ = 565
	ZydisMnemonicPabsd ZydisMnemonic_ = 566
	ZydisMnemonicPabsw ZydisMnemonic_ = 567
	ZydisMnemonicPackssdw ZydisMnemonic_ = 568
	ZydisMnemonicPacksswb ZydisMnemonic_ = 569
	ZydisMnemonicPackusdw ZydisMnemonic_ = 570
	ZydisMnemonicPackuswb ZydisMnemonic_ = 571
	ZydisMnemonicPaddb ZydisMnemonic_ = 572
	ZydisMnemonicPaddd ZydisMnemonic_ = 573
	ZydisMnemonicPaddq ZydisMnemonic_ = 574
	ZydisMnemonicPaddsb ZydisMnemonic_ = 575
	ZydisMnemonicPaddsw ZydisMnemonic_ = 576
	ZydisMnemonicPaddusb ZydisMnemonic_ = 577
	ZydisMnemonicPaddusw ZydisMnemonic_ = 578
	ZydisMnemonicPaddw ZydisMnemonic_ = 579
	ZydisMnemonicPalignr ZydisMnemonic_ = 580
	ZydisMnemonicPand ZydisMnemonic_ = 581
	ZydisMnemonicPandn ZydisMnemonic_ = 582
	ZydisMnemonicPause ZydisMnemonic_ = 583
	ZydisMnemonicPavgb ZydisMnemonic_ = 584
	ZydisMnemonicPavgusb ZydisMnemonic_ = 585
	ZydisMnemonicPavgw ZydisMnemonic_ = 586
	ZydisMnemonicPblendvb ZydisMnemonic_ = 587
	ZydisMnemonicPblendw ZydisMnemonic_ = 588
	ZydisMnemonicPbndkb ZydisMnemonic_ = 589
	ZydisMnemonicPclmulqdq ZydisMnemonic_ = 590
	ZydisMnemonicPcmpeqb ZydisMnemonic_ = 591
	ZydisMnemonicPcmpeqd ZydisMnemonic_ = 592
	ZydisMnemonicPcmpeqq ZydisMnemonic_ = 593
	ZydisMnemonicPcmpeqw ZydisMnemonic_ = 594
	ZydisMnemonicPcmpestri ZydisMnemonic_ = 595
	ZydisMnemonicPcmpestrm ZydisMnemonic_ = 596
	ZydisMnemonicPcmpgtb ZydisMnemonic_ = 597
	ZydisMnemonicPcmpgtd ZydisMnemonic_ = 598
	ZydisMnemonicPcmpgtq ZydisMnemonic_ = 599
	ZydisMnemonicPcmpgtw ZydisMnemonic_ = 600
	ZydisMnemonicPcmpistri ZydisMnemonic_ = 601
	ZydisMnemonicPcmpistrm ZydisMnemonic_ = 602
	ZydisMnemonicPcommit ZydisMnemonic_ = 603
	ZydisMnemonicPconfig ZydisMnemonic_ = 604
	ZydisMnemonicPdep ZydisMnemonic_ = 605
	ZydisMnemonicPext ZydisMnemonic_ = 606
	ZydisMnemonicPextrb ZydisMnemonic_ = 607
	ZydisMnemonicPextrd ZydisMnemonic_ = 608
	ZydisMnemonicPextrq ZydisMnemonic_ = 609
	ZydisMnemonicPextrw ZydisMnemonic_ = 610
	ZydisMnemonicPf2id ZydisMnemonic_ = 611
	ZydisMnemonicPf2iw ZydisMnemonic_ = 612
	ZydisMnemonicPfacc ZydisMnemonic_ = 613
	ZydisMnemonicPfadd ZydisMnemonic_ = 614
	ZydisMnemonicPfcmpeq ZydisMnemonic_ = 615
	ZydisMnemonicPfcmpge ZydisMnemonic_ = 616
	ZydisMnemonicPfcmpgt ZydisMnemonic_ = 617
	ZydisMnemonicPfcpit1 ZydisMnemonic_ = 618
	ZydisMnemonicPfmax ZydisMnemonic_ = 619
	ZydisMnemonicPfmin ZydisMnemonic_ = 620
	ZydisMnemonicPfmul ZydisMnemonic_ = 621
	ZydisMnemonicPfnacc ZydisMnemonic_ = 622
	ZydisMnemonicPfpnacc ZydisMnemonic_ = 623
	ZydisMnemonicPfrcp ZydisMnemonic_ = 624
	ZydisMnemonicPfrcpit2 ZydisMnemonic_ = 625
	ZydisMnemonicPfrsqit1 ZydisMnemonic_ = 626
	ZydisMnemonicPfsqrt ZydisMnemonic_ = 627
	ZydisMnemonicPfsub ZydisMnemonic_ = 628
	ZydisMnemonicPfsubr ZydisMnemonic_ = 629
	ZydisMnemonicPhaddd ZydisMnemonic_ = 630
	ZydisMnemonicPhaddsw ZydisMnemonic_ = 631
	ZydisMnemonicPhaddw ZydisMnemonic_ = 632
	ZydisMnemonicPhminposuw ZydisMnemonic_ = 633
	ZydisMnemonicPhsubd ZydisMnemonic_ = 634
	ZydisMnemonicPhsubsw ZydisMnemonic_ = 635
	ZydisMnemonicPhsubw ZydisMnemonic_ = 636
	ZydisMnemonicPi2fd ZydisMnemonic_ = 637
	ZydisMnemonicPi2fw ZydisMnemonic_ = 638
	ZydisMnemonicPinsrb ZydisMnemonic_ = 639
	ZydisMnemonicPinsrd ZydisMnemonic_ = 640
	ZydisMnemonicPinsrq ZydisMnemonic_ = 641
	ZydisMnemonicPinsrw ZydisMnemonic_ = 642
	ZydisMnemonicPmaddubsw ZydisMnemonic_ = 643
	ZydisMnemonicPmaddwd ZydisMnemonic_ = 644
	ZydisMnemonicPmaxsb ZydisMnemonic_ = 645
	ZydisMnemonicPmaxsd ZydisMnemonic_ = 646
	ZydisMnemonicPmaxsw ZydisMnemonic_ = 647
	ZydisMnemonicPmaxub ZydisMnemonic_ = 648
	ZydisMnemonicPmaxud ZydisMnemonic_ = 649
	ZydisMnemonicPmaxuw ZydisMnemonic_ = 650
	ZydisMnemonicPminsb ZydisMnemonic_ = 651
	ZydisMnemonicPminsd ZydisMnemonic_ = 652
	ZydisMnemonicPminsw ZydisMnemonic_ = 653
	ZydisMnemonicPminub ZydisMnemonic_ = 654
	ZydisMnemonicPminud ZydisMnemonic_ = 655
	ZydisMnemonicPminuw ZydisMnemonic_ = 656
	ZydisMnemonicPmovmskb ZydisMnemonic_ = 657
	ZydisMnemonicPmovsxbd ZydisMnemonic_ = 658
	ZydisMnemonicPmovsxbq ZydisMnemonic_ = 659
	ZydisMnemonicPmovsxbw ZydisMnemonic_ = 660
	ZydisMnemonicPmovsxdq ZydisMnemonic_ = 661
	ZydisMnemonicPmovsxwd ZydisMnemonic_ = 662
	ZydisMnemonicPmovsxwq ZydisMnemonic_ = 663
	ZydisMnemonicPmovzxbd ZydisMnemonic_ = 664
	ZydisMnemonicPmovzxbq ZydisMnemonic_ = 665
	ZydisMnemonicPmovzxbw ZydisMnemonic_ = 666
	ZydisMnemonicPmovzxdq ZydisMnemonic_ = 667
	ZydisMnemonicPmovzxwd ZydisMnemonic_ = 668
	ZydisMnemonicPmovzxwq ZydisMnemonic_ = 669
	ZydisMnemonicPmuldq ZydisMnemonic_ = 670
	ZydisMnemonicPmulhrsw ZydisMnemonic_ = 671
	ZydisMnemonicPmulhrw ZydisMnemonic_ = 672
	ZydisMnemonicPmulhuw ZydisMnemonic_ = 673
	ZydisMnemonicPmulhw ZydisMnemonic_ = 674
	ZydisMnemonicPmulld ZydisMnemonic_ = 675
	ZydisMnemonicPmullw ZydisMnemonic_ = 676
	ZydisMnemonicPmuludq ZydisMnemonic_ = 677
	ZydisMnemonicPop ZydisMnemonic_ = 678
	ZydisMnemonicPop2 ZydisMnemonic_ = 679
	ZydisMnemonicPop2p ZydisMnemonic_ = 680
	ZydisMnemonicPopa ZydisMnemonic_ = 681
	ZydisMnemonicPopad ZydisMnemonic_ = 682
	ZydisMnemonicPopcnt ZydisMnemonic_ = 683
	ZydisMnemonicPopf ZydisMnemonic_ = 684
	ZydisMnemonicPopfd ZydisMnemonic_ = 685
	ZydisMnemonicPopfq ZydisMnemonic_ = 686
	ZydisMnemonicPopp ZydisMnemonic_ = 687
	ZydisMnemonicPor ZydisMnemonic_ = 688
	ZydisMnemonicPrefetch ZydisMnemonic_ = 689
	ZydisMnemonicPrefetchit0 ZydisMnemonic_ = 690
	ZydisMnemonicPrefetchit1 ZydisMnemonic_ = 691
	ZydisMnemonicPrefetchnta ZydisMnemonic_ = 692
	ZydisMnemonicPrefetcht0 ZydisMnemonic_ = 693
	ZydisMnemonicPrefetcht1 ZydisMnemonic_ = 694
	ZydisMnemonicPrefetcht2 ZydisMnemonic_ = 695
	ZydisMnemonicPrefetchw ZydisMnemonic_ = 696
	ZydisMnemonicPrefetchwt1 ZydisMnemonic_ = 697
	ZydisMnemonicPsadbw ZydisMnemonic_ = 698
	ZydisMnemonicPshufb ZydisMnemonic_ = 699
	ZydisMnemonicPshufd ZydisMnemonic_ = 700
	ZydisMnemonicPshufhw ZydisMnemonic_ = 701
	ZydisMnemonicPshuflw ZydisMnemonic_ = 702
	ZydisMnemonicPshufw ZydisMnemonic_ = 703
	ZydisMnemonicPsignb ZydisMnemonic_ = 704
	ZydisMnemonicPsignd ZydisMnemonic_ = 705
	ZydisMnemonicPsignw ZydisMnemonic_ = 706
	ZydisMnemonicPslld ZydisMnemonic_ = 707
	ZydisMnemonicPslldq ZydisMnemonic_ = 708
	ZydisMnemonicPsllq ZydisMnemonic_ = 709
	ZydisMnemonicPsllw ZydisMnemonic_ = 710
	ZydisMnemonicPsmash ZydisMnemonic_ = 711
	ZydisMnemonicPsrad ZydisMnemonic_ = 712
	ZydisMnemonicPsraw ZydisMnemonic_ = 713
	ZydisMnemonicPsrld ZydisMnemonic_ = 714
	ZydisMnemonicPsrldq ZydisMnemonic_ = 715
	ZydisMnemonicPsrlq ZydisMnemonic_ = 716
	ZydisMnemonicPsrlw ZydisMnemonic_ = 717
	ZydisMnemonicPsubb ZydisMnemonic_ = 718
	ZydisMnemonicPsubd ZydisMnemonic_ = 719
	ZydisMnemonicPsubq ZydisMnemonic_ = 720
	ZydisMnemonicPsubsb ZydisMnemonic_ = 721
	ZydisMnemonicPsubsw ZydisMnemonic_ = 722
	ZydisMnemonicPsubusb ZydisMnemonic_ = 723
	ZydisMnemonicPsubusw ZydisMnemonic_ = 724
	ZydisMnemonicPsubw ZydisMnemonic_ = 725
	ZydisMnemonicPswapd ZydisMnemonic_ = 726
	ZydisMnemonicPtest ZydisMnemonic_ = 727
	ZydisMnemonicPtwrite ZydisMnemonic_ = 728
	ZydisMnemonicPunpckhbw ZydisMnemonic_ = 729
	ZydisMnemonicPunpckhdq ZydisMnemonic_ = 730
	ZydisMnemonicPunpckhqdq ZydisMnemonic_ = 731
	ZydisMnemonicPunpckhwd ZydisMnemonic_ = 732
	ZydisMnemonicPunpcklbw ZydisMnemonic_ = 733
	ZydisMnemonicPunpckldq ZydisMnemonic_ = 734
	ZydisMnemonicPunpcklqdq ZydisMnemonic_ = 735
	ZydisMnemonicPunpcklwd ZydisMnemonic_ = 736
	ZydisMnemonicPush ZydisMnemonic_ = 737
	ZydisMnemonicPush2 ZydisMnemonic_ = 738
	ZydisMnemonicPush2p ZydisMnemonic_ = 739
	ZydisMnemonicPusha ZydisMnemonic_ = 740
	ZydisMnemonicPushad ZydisMnemonic_ = 741
	ZydisMnemonicPushf ZydisMnemonic_ = 742
	ZydisMnemonicPushfd ZydisMnemonic_ = 743
	ZydisMnemonicPushfq ZydisMnemonic_ = 744
	ZydisMnemonicPushp ZydisMnemonic_ = 745
	ZydisMnemonicPvalidate ZydisMnemonic_ = 746
	ZydisMnemonicPxor ZydisMnemonic_ = 747
	ZydisMnemonicRcl ZydisMnemonic_ = 748
	ZydisMnemonicRcpps ZydisMnemonic_ = 749
	ZydisMnemonicRcpss ZydisMnemonic_ = 750
	ZydisMnemonicRcr ZydisMnemonic_ = 751
	ZydisMnemonicRdfsbase ZydisMnemonic_ = 752
	ZydisMnemonicRdgsbase ZydisMnemonic_ = 753
	ZydisMnemonicRdmsr ZydisMnemonic_ = 754
	ZydisMnemonicRdmsrlist ZydisMnemonic_ = 755
	ZydisMnemonicRdpid ZydisMnemonic_ = 756
	ZydisMnemonicRdpkru ZydisMnemonic_ = 757
	ZydisMnemonicRdpmc ZydisMnemonic_ = 758
	ZydisMnemonicRdpru ZydisMnemonic_ = 759
	ZydisMnemonicRdrand ZydisMnemonic_ = 760
	ZydisMnemonicRdseed ZydisMnemonic_ = 761
	ZydisMnemonicRdsspd ZydisMnemonic_ = 762
	ZydisMnemonicRdsspq ZydisMnemonic_ = 763
	ZydisMnemonicRdtsc ZydisMnemonic_ = 764
	ZydisMnemonicRdtscp ZydisMnemonic_ = 765
	ZydisMnemonicRet ZydisMnemonic_ = 766
	ZydisMnemonicRmpadjust ZydisMnemonic_ = 767
	ZydisMnemonicRmpupdate ZydisMnemonic_ = 768
	ZydisMnemonicRol ZydisMnemonic_ = 769
	ZydisMnemonicRor ZydisMnemonic_ = 770
	ZydisMnemonicRorx ZydisMnemonic_ = 771
	ZydisMnemonicRoundpd ZydisMnemonic_ = 772
	ZydisMnemonicRoundps ZydisMnemonic_ = 773
	ZydisMnemonicRoundsd ZydisMnemonic_ = 774
	ZydisMnemonicRoundss ZydisMnemonic_ = 775
	ZydisMnemonicRsm ZydisMnemonic_ = 776
	ZydisMnemonicRsqrtps ZydisMnemonic_ = 777
	ZydisMnemonicRsqrtss ZydisMnemonic_ = 778
	ZydisMnemonicRstorssp ZydisMnemonic_ = 779
	ZydisMnemonicSahf ZydisMnemonic_ = 780
	ZydisMnemonicSalc ZydisMnemonic_ = 781
	ZydisMnemonicSar ZydisMnemonic_ = 782
	ZydisMnemonicSarx ZydisMnemonic_ = 783
	ZydisMnemonicSaveprevssp ZydisMnemonic_ = 784
	ZydisMnemonicSbb ZydisMnemonic_ = 785
	ZydisMnemonicScasb ZydisMnemonic_ = 786
	ZydisMnemonicScasd ZydisMnemonic_ = 787
	ZydisMnemonicScasq ZydisMnemonic_ = 788
	ZydisMnemonicScasw ZydisMnemonic_ = 789
	ZydisMnemonicSeamcall ZydisMnemonic_ = 790
	ZydisMnemonicSeamops ZydisMnemonic_ = 791
	ZydisMnemonicSeamret ZydisMnemonic_ = 792
	ZydisMnemonicSenduipi ZydisMnemonic_ = 793
	ZydisMnemonicSerialize ZydisMnemonic_ = 794
	ZydisMnemonicSetb ZydisMnemonic_ = 795
	ZydisMnemonicSetbe ZydisMnemonic_ = 796
	ZydisMnemonicSetl ZydisMnemonic_ = 797
	ZydisMnemonicSetle ZydisMnemonic_ = 798
	ZydisMnemonicSetnb ZydisMnemonic_ = 799
	ZydisMnemonicSetnbe ZydisMnemonic_ = 800
	ZydisMnemonicSetnl ZydisMnemonic_ = 801
	ZydisMnemonicSetnle ZydisMnemonic_ = 802
	ZydisMnemonicSetno ZydisMnemonic_ = 803
	ZydisMnemonicSetnp ZydisMnemonic_ = 804
	ZydisMnemonicSetns ZydisMnemonic_ = 805
	ZydisMnemonicSetnz ZydisMnemonic_ = 806
	ZydisMnemonicSeto ZydisMnemonic_ = 807
	ZydisMnemonicSetp ZydisMnemonic_ = 808
	ZydisMnemonicSets ZydisMnemonic_ = 809
	ZydisMnemonicSetssbsy ZydisMnemonic_ = 810
	ZydisMnemonicSetz ZydisMnemonic_ = 811
	ZydisMnemonicSetzub ZydisMnemonic_ = 812
	ZydisMnemonicSetzube ZydisMnemonic_ = 813
	ZydisMnemonicSetzul ZydisMnemonic_ = 814
	ZydisMnemonicSetzule ZydisMnemonic_ = 815
	ZydisMnemonicSetzunb ZydisMnemonic_ = 816
	ZydisMnemonicSetzunbe ZydisMnemonic_ = 817
	ZydisMnemonicSetzunl ZydisMnemonic_ = 818
	ZydisMnemonicSetzunle ZydisMnemonic_ = 819
	ZydisMnemonicSetzuno ZydisMnemonic_ = 820
	ZydisMnemonicSetzunp ZydisMnemonic_ = 821
	ZydisMnemonicSetzuns ZydisMnemonic_ = 822
	ZydisMnemonicSetzunz ZydisMnemonic_ = 823
	ZydisMnemonicSetzuo ZydisMnemonic_ = 824
	ZydisMnemonicSetzup ZydisMnemonic_ = 825
	ZydisMnemonicSetzus ZydisMnemonic_ = 826
	ZydisMnemonicSetzuz ZydisMnemonic_ = 827
	ZydisMnemonicSfence ZydisMnemonic_ = 828
	ZydisMnemonicSgdt ZydisMnemonic_ = 829
	ZydisMnemonicSha1msg1 ZydisMnemonic_ = 830
	ZydisMnemonicSha1msg2 ZydisMnemonic_ = 831
	ZydisMnemonicSha1nexte ZydisMnemonic_ = 832
	ZydisMnemonicSha1rnds4 ZydisMnemonic_ = 833
	ZydisMnemonicSha256msg1 ZydisMnemonic_ = 834
	ZydisMnemonicSha256msg2 ZydisMnemonic_ = 835
	ZydisMnemonicSha256rnds2 ZydisMnemonic_ = 836
	ZydisMnemonicShl ZydisMnemonic_ = 837
	ZydisMnemonicShld ZydisMnemonic_ = 838
	ZydisMnemonicShlx ZydisMnemonic_ = 839
	ZydisMnemonicShr ZydisMnemonic_ = 840
	ZydisMnemonicShrd ZydisMnemonic_ = 841
	ZydisMnemonicShrx ZydisMnemonic_ = 842
	ZydisMnemonicShufpd ZydisMnemonic_ = 843
	ZydisMnemonicShufps ZydisMnemonic_ = 844
	ZydisMnemonicSidt ZydisMnemonic_ = 845
	ZydisMnemonicSkinit ZydisMnemonic_ = 846
	ZydisMnemonicSldt ZydisMnemonic_ = 847
	ZydisMnemonicSlwpcb ZydisMnemonic_ = 848
	ZydisMnemonicSmsw ZydisMnemonic_ = 849
	ZydisMnemonicSpflt ZydisMnemonic_ = 850
	ZydisMnemonicSqrtpd ZydisMnemonic_ = 851
	ZydisMnemonicSqrtps ZydisMnemonic_ = 852
	ZydisMnemonicSqrtsd ZydisMnemonic_ = 853
	ZydisMnemonicSqrtss ZydisMnemonic_ = 854
	ZydisMnemonicStac ZydisMnemonic_ = 855
	ZydisMnemonicStc ZydisMnemonic_ = 856
	ZydisMnemonicStd ZydisMnemonic_ = 857
	ZydisMnemonicStgi ZydisMnemonic_ = 858
	ZydisMnemonicSti ZydisMnemonic_ = 859
	ZydisMnemonicStmxcsr ZydisMnemonic_ = 860
	ZydisMnemonicStosb ZydisMnemonic_ = 861
	ZydisMnemonicStosd ZydisMnemonic_ = 862
	ZydisMnemonicStosq ZydisMnemonic_ = 863
	ZydisMnemonicStosw ZydisMnemonic_ = 864
	ZydisMnemonicStr ZydisMnemonic_ = 865
	ZydisMnemonicSttilecfg ZydisMnemonic_ = 866
	ZydisMnemonicStui ZydisMnemonic_ = 867
	ZydisMnemonicSub ZydisMnemonic_ = 868
	ZydisMnemonicSubpd ZydisMnemonic_ = 869
	ZydisMnemonicSubps ZydisMnemonic_ = 870
	ZydisMnemonicSubsd ZydisMnemonic_ = 871
	ZydisMnemonicSubss ZydisMnemonic_ = 872
	ZydisMnemonicSwapgs ZydisMnemonic_ = 873
	ZydisMnemonicSyscall ZydisMnemonic_ = 874
	ZydisMnemonicSysenter ZydisMnemonic_ = 875
	ZydisMnemonicSysexit ZydisMnemonic_ = 876
	ZydisMnemonicSysret ZydisMnemonic_ = 877
	ZydisMnemonicT1mskc ZydisMnemonic_ = 878
	ZydisMnemonicTdcall ZydisMnemonic_ = 879
	ZydisMnemonicTdpbf16ps ZydisMnemonic_ = 880
	ZydisMnemonicTdpbssd ZydisMnemonic_ = 881
	ZydisMnemonicTdpbsud ZydisMnemonic_ = 882
	ZydisMnemonicTdpbusd ZydisMnemonic_ = 883
	ZydisMnemonicTdpbuud ZydisMnemonic_ = 884
	ZydisMnemonicTdpfp16ps ZydisMnemonic_ = 885
	ZydisMnemonicTest ZydisMnemonic_ = 886
	ZydisMnemonicTestui ZydisMnemonic_ = 887
	ZydisMnemonicTileloadd ZydisMnemonic_ = 888
	ZydisMnemonicTileloaddt1 ZydisMnemonic_ = 889
	ZydisMnemonicTilerelease ZydisMnemonic_ = 890
	ZydisMnemonicTilestored ZydisMnemonic_ = 891
	ZydisMnemonicTilezero ZydisMnemonic_ = 892
	ZydisMnemonicTlbsync ZydisMnemonic_ = 893
	ZydisMnemonicTpause ZydisMnemonic_ = 894
	ZydisMnemonicTzcnt ZydisMnemonic_ = 895
	ZydisMnemonicTzcnti ZydisMnemonic_ = 896
	ZydisMnemonicTzmsk ZydisMnemonic_ = 897
	ZydisMnemonicUcomisd ZydisMnemonic_ = 898
	ZydisMnemonicUcomiss ZydisMnemonic_ = 899
	ZydisMnemonicUd0 ZydisMnemonic_ = 900
	ZydisMnemonicUd1 ZydisMnemonic_ = 901
	ZydisMnemonicUd2 ZydisMnemonic_ = 902
	ZydisMnemonicUiret ZydisMnemonic_ = 903
	ZydisMnemonicUmonitor ZydisMnemonic_ = 904
	ZydisMnemonicUmwait ZydisMnemonic_ = 905
	ZydisMnemonicUnpckhpd ZydisMnemonic_ = 906
	ZydisMnemonicUnpckhps ZydisMnemonic_ = 907
	ZydisMnemonicUnpcklpd ZydisMnemonic_ = 908
	ZydisMnemonicUnpcklps ZydisMnemonic_ = 909
	ZydisMnemonicUrdmsr ZydisMnemonic_ = 910
	ZydisMnemonicUwrmsr ZydisMnemonic_ = 911
	ZydisMnemonicV4fmaddps ZydisMnemonic_ = 912
	ZydisMnemonicV4fmaddss ZydisMnemonic_ = 913
	ZydisMnemonicV4fnmaddps ZydisMnemonic_ = 914
	ZydisMnemonicV4fnmaddss ZydisMnemonic_ = 915
	ZydisMnemonicVaddnpd ZydisMnemonic_ = 916
	ZydisMnemonicVaddnps ZydisMnemonic_ = 917
	ZydisMnemonicVaddpd ZydisMnemonic_ = 918
	ZydisMnemonicVaddph ZydisMnemonic_ = 919
	ZydisMnemonicVaddps ZydisMnemonic_ = 920
	ZydisMnemonicVaddsd ZydisMnemonic_ = 921
	ZydisMnemonicVaddsetsps ZydisMnemonic_ = 922
	ZydisMnemonicVaddsh ZydisMnemonic_ = 923
	ZydisMnemonicVaddss ZydisMnemonic_ = 924
	ZydisMnemonicVaddsubpd ZydisMnemonic_ = 925
	ZydisMnemonicVaddsubps ZydisMnemonic_ = 926
	ZydisMnemonicVaesdec ZydisMnemonic_ = 927
	ZydisMnemonicVaesdeclast ZydisMnemonic_ = 928
	ZydisMnemonicVaesenc ZydisMnemonic_ = 929
	ZydisMnemonicVaesenclast ZydisMnemonic_ = 930
	ZydisMnemonicVaesimc ZydisMnemonic_ = 931
	ZydisMnemonicVaeskeygenassist ZydisMnemonic_ = 932
	ZydisMnemonicValignd ZydisMnemonic_ = 933
	ZydisMnemonicValignq ZydisMnemonic_ = 934
	ZydisMnemonicVandnpd ZydisMnemonic_ = 935
	ZydisMnemonicVandnps ZydisMnemonic_ = 936
	ZydisMnemonicVandpd ZydisMnemonic_ = 937
	ZydisMnemonicVandps ZydisMnemonic_ = 938
	ZydisMnemonicVbcstnebf162ps ZydisMnemonic_ = 939
	ZydisMnemonicVbcstnesh2ps ZydisMnemonic_ = 940
	ZydisMnemonicVblendmpd ZydisMnemonic_ = 941
	ZydisMnemonicVblendmps ZydisMnemonic_ = 942
	ZydisMnemonicVblendpd ZydisMnemonic_ = 943
	ZydisMnemonicVblendps ZydisMnemonic_ = 944
	ZydisMnemonicVblendvpd ZydisMnemonic_ = 945
	ZydisMnemonicVblendvps ZydisMnemonic_ = 946
	ZydisMnemonicVbroadcastf128 ZydisMnemonic_ = 947
	ZydisMnemonicVbroadcastf32x2 ZydisMnemonic_ = 948
	ZydisMnemonicVbroadcastf32x4 ZydisMnemonic_ = 949
	ZydisMnemonicVbroadcastf32x8 ZydisMnemonic_ = 950
	ZydisMnemonicVbroadcastf64x2 ZydisMnemonic_ = 951
	ZydisMnemonicVbroadcastf64x4 ZydisMnemonic_ = 952
	ZydisMnemonicVbroadcasti128 ZydisMnemonic_ = 953
	ZydisMnemonicVbroadcasti32x2 ZydisMnemonic_ = 954
	ZydisMnemonicVbroadcasti32x4 ZydisMnemonic_ = 955
	ZydisMnemonicVbroadcasti32x8 ZydisMnemonic_ = 956
	ZydisMnemonicVbroadcasti64x2 ZydisMnemonic_ = 957
	ZydisMnemonicVbroadcasti64x4 ZydisMnemonic_ = 958
	ZydisMnemonicVbroadcastsd ZydisMnemonic_ = 959
	ZydisMnemonicVbroadcastss ZydisMnemonic_ = 960
	ZydisMnemonicVcmppd ZydisMnemonic_ = 961
	ZydisMnemonicVcmpph ZydisMnemonic_ = 962
	ZydisMnemonicVcmpps ZydisMnemonic_ = 963
	ZydisMnemonicVcmpsd ZydisMnemonic_ = 964
	ZydisMnemonicVcmpsh ZydisMnemonic_ = 965
	ZydisMnemonicVcmpss ZydisMnemonic_ = 966
	ZydisMnemonicVcomisd ZydisMnemonic_ = 967
	ZydisMnemonicVcomish ZydisMnemonic_ = 968
	ZydisMnemonicVcomiss ZydisMnemonic_ = 969
	ZydisMnemonicVcompresspd ZydisMnemonic_ = 970
	ZydisMnemonicVcompressps ZydisMnemonic_ = 971
	ZydisMnemonicVcvtdq2pd ZydisMnemonic_ = 972
	ZydisMnemonicVcvtdq2ph ZydisMnemonic_ = 973
	ZydisMnemonicVcvtdq2ps ZydisMnemonic_ = 974
	ZydisMnemonicVcvtfxpntdq2ps ZydisMnemonic_ = 975
	ZydisMnemonicVcvtfxpntpd2dq ZydisMnemonic_ = 976
	ZydisMnemonicVcvtfxpntpd2udq ZydisMnemonic_ = 977
	ZydisMnemonicVcvtfxpntps2dq ZydisMnemonic_ = 978
	ZydisMnemonicVcvtfxpntps2udq ZydisMnemonic_ = 979
	ZydisMnemonicVcvtfxpntudq2ps ZydisMnemonic_ = 980
	ZydisMnemonicVcvtne2ps2bf16 ZydisMnemonic_ = 981
	ZydisMnemonicVcvtneebf162ps ZydisMnemonic_ = 982
	ZydisMnemonicVcvtneeph2ps ZydisMnemonic_ = 983
	ZydisMnemonicVcvtneobf162ps ZydisMnemonic_ = 984
	ZydisMnemonicVcvtneoph2ps ZydisMnemonic_ = 985
	ZydisMnemonicVcvtneps2bf16 ZydisMnemonic_ = 986
	ZydisMnemonicVcvtpd2dq ZydisMnemonic_ = 987
	ZydisMnemonicVcvtpd2ph ZydisMnemonic_ = 988
	ZydisMnemonicVcvtpd2ps ZydisMnemonic_ = 989
	ZydisMnemonicVcvtpd2qq ZydisMnemonic_ = 990
	ZydisMnemonicVcvtpd2udq ZydisMnemonic_ = 991
	ZydisMnemonicVcvtpd2uqq ZydisMnemonic_ = 992
	ZydisMnemonicVcvtph2dq ZydisMnemonic_ = 993
	ZydisMnemonicVcvtph2pd ZydisMnemonic_ = 994
	ZydisMnemonicVcvtph2ps ZydisMnemonic_ = 995
	ZydisMnemonicVcvtph2psx ZydisMnemonic_ = 996
	ZydisMnemonicVcvtph2qq ZydisMnemonic_ = 997
	ZydisMnemonicVcvtph2udq ZydisMnemonic_ = 998
	ZydisMnemonicVcvtph2uqq ZydisMnemonic_ = 999
	ZydisMnemonicVcvtph2uw ZydisMnemonic_ = 1000
	ZydisMnemonicVcvtph2w ZydisMnemonic_ = 1001
	ZydisMnemonicVcvtps2dq ZydisMnemonic_ = 1002
	ZydisMnemonicVcvtps2pd ZydisMnemonic_ = 1003
	ZydisMnemonicVcvtps2ph ZydisMnemonic_ = 1004
	ZydisMnemonicVcvtps2phx ZydisMnemonic_ = 1005
	ZydisMnemonicVcvtps2qq ZydisMnemonic_ = 1006
	ZydisMnemonicVcvtps2udq ZydisMnemonic_ = 1007
	ZydisMnemonicVcvtps2uqq ZydisMnemonic_ = 1008
	ZydisMnemonicVcvtqq2pd ZydisMnemonic_ = 1009
	ZydisMnemonicVcvtqq2ph ZydisMnemonic_ = 1010
	ZydisMnemonicVcvtqq2ps ZydisMnemonic_ = 1011
	ZydisMnemonicVcvtsd2sh ZydisMnemonic_ = 1012
	ZydisMnemonicVcvtsd2si ZydisMnemonic_ = 1013
	ZydisMnemonicVcvtsd2ss ZydisMnemonic_ = 1014
	ZydisMnemonicVcvtsd2usi ZydisMnemonic_ = 1015
	ZydisMnemonicVcvtsh2sd ZydisMnemonic_ = 1016
	ZydisMnemonicVcvtsh2si ZydisMnemonic_ = 1017
	ZydisMnemonicVcvtsh2ss ZydisMnemonic_ = 1018
	ZydisMnemonicVcvtsh2usi ZydisMnemonic_ = 1019
	ZydisMnemonicVcvtsi2sd ZydisMnemonic_ = 1020
	ZydisMnemonicVcvtsi2sh ZydisMnemonic_ = 1021
	ZydisMnemonicVcvtsi2ss ZydisMnemonic_ = 1022
	ZydisMnemonicVcvtss2sd ZydisMnemonic_ = 1023
	ZydisMnemonicVcvtss2sh ZydisMnemonic_ = 1024
	ZydisMnemonicVcvtss2si ZydisMnemonic_ = 1025
	ZydisMnemonicVcvtss2usi ZydisMnemonic_ = 1026
	ZydisMnemonicVcvttpd2dq ZydisMnemonic_ = 1027
	ZydisMnemonicVcvttpd2qq ZydisMnemonic_ = 1028
	ZydisMnemonicVcvttpd2udq ZydisMnemonic_ = 1029
	ZydisMnemonicVcvttpd2uqq ZydisMnemonic_ = 1030
	ZydisMnemonicVcvttph2dq ZydisMnemonic_ = 1031
	ZydisMnemonicVcvttph2qq ZydisMnemonic_ = 1032
	ZydisMnemonicVcvttph2udq ZydisMnemonic_ = 1033
	ZydisMnemonicVcvttph2uqq ZydisMnemonic_ = 1034
	ZydisMnemonicVcvttph2uw ZydisMnemonic_ = 1035
	ZydisMnemonicVcvttph2w ZydisMnemonic_ = 1036
	ZydisMnemonicVcvttps2dq ZydisMnemonic_ = 1037
	ZydisMnemonicVcvttps2qq ZydisMnemonic_ = 1038
	ZydisMnemonicVcvttps2udq ZydisMnemonic_ = 1039
	ZydisMnemonicVcvttps2uqq ZydisMnemonic_ = 1040
	ZydisMnemonicVcvttsd2si ZydisMnemonic_ = 1041
	ZydisMnemonicVcvttsd2usi ZydisMnemonic_ = 1042
	ZydisMnemonicVcvttsh2si ZydisMnemonic_ = 1043
	ZydisMnemonicVcvttsh2usi ZydisMnemonic_ = 1044
	ZydisMnemonicVcvttss2si ZydisMnemonic_ = 1045
	ZydisMnemonicVcvttss2usi ZydisMnemonic_ = 1046
	ZydisMnemonicVcvtudq2pd ZydisMnemonic_ = 1047
	ZydisMnemonicVcvtudq2ph ZydisMnemonic_ = 1048
	ZydisMnemonicVcvtudq2ps ZydisMnemonic_ = 1049
	ZydisMnemonicVcvtuqq2pd ZydisMnemonic_ = 1050
	ZydisMnemonicVcvtuqq2ph ZydisMnemonic_ = 1051
	ZydisMnemonicVcvtuqq2ps ZydisMnemonic_ = 1052
	ZydisMnemonicVcvtusi2sd ZydisMnemonic_ = 1053
	ZydisMnemonicVcvtusi2sh ZydisMnemonic_ = 1054
	ZydisMnemonicVcvtusi2ss ZydisMnemonic_ = 1055
	ZydisMnemonicVcvtuw2ph ZydisMnemonic_ = 1056
	ZydisMnemonicVcvtw2ph ZydisMnemonic_ = 1057
	ZydisMnemonicVdbpsadbw ZydisMnemonic_ = 1058
	ZydisMnemonicVdivpd ZydisMnemonic_ = 1059
	ZydisMnemonicVdivph ZydisMnemonic_ = 1060
	ZydisMnemonicVdivps ZydisMnemonic_ = 1061
	ZydisMnemonicVdivsd ZydisMnemonic_ = 1062
	ZydisMnemonicVdivsh ZydisMnemonic_ = 1063
	ZydisMnemonicVdivss ZydisMnemonic_ = 1064
	ZydisMnemonicVdpbf16ps ZydisMnemonic_ = 1065
	ZydisMnemonicVdppd ZydisMnemonic_ = 1066
	ZydisMnemonicVdpps ZydisMnemonic_ = 1067
	ZydisMnemonicVerr ZydisMnemonic_ = 1068
	ZydisMnemonicVerw ZydisMnemonic_ = 1069
	ZydisMnemonicVexp223ps ZydisMnemonic_ = 1070
	ZydisMnemonicVexp2pd ZydisMnemonic_ = 1071
	ZydisMnemonicVexp2ps ZydisMnemonic_ = 1072
	ZydisMnemonicVexpandpd ZydisMnemonic_ = 1073
	ZydisMnemonicVexpandps ZydisMnemonic_ = 1074
	ZydisMnemonicVextractf128 ZydisMnemonic_ = 1075
	ZydisMnemonicVextractf32x4 ZydisMnemonic_ = 1076
	ZydisMnemonicVextractf32x8 ZydisMnemonic_ = 1077
	ZydisMnemonicVextractf64x2 ZydisMnemonic_ = 1078
	ZydisMnemonicVextractf64x4 ZydisMnemonic_ = 1079
	ZydisMnemonicVextracti128 ZydisMnemonic_ = 1080
	ZydisMnemonicVextracti32x4 ZydisMnemonic_ = 1081
	ZydisMnemonicVextracti32x8 ZydisMnemonic_ = 1082
	ZydisMnemonicVextracti64x2 ZydisMnemonic_ = 1083
	ZydisMnemonicVextracti64x4 ZydisMnemonic_ = 1084
	ZydisMnemonicVextractps ZydisMnemonic_ = 1085
	ZydisMnemonicVfcmaddcph ZydisMnemonic_ = 1086
	ZydisMnemonicVfcmaddcsh ZydisMnemonic_ = 1087
	ZydisMnemonicVfcmulcph ZydisMnemonic_ = 1088
	ZydisMnemonicVfcmulcsh ZydisMnemonic_ = 1089
	ZydisMnemonicVfixupimmpd ZydisMnemonic_ = 1090
	ZydisMnemonicVfixupimmps ZydisMnemonic_ = 1091
	ZydisMnemonicVfixupimmsd ZydisMnemonic_ = 1092
	ZydisMnemonicVfixupimmss ZydisMnemonic_ = 1093
	ZydisMnemonicVfixupnanpd ZydisMnemonic_ = 1094
	ZydisMnemonicVfixupnanps ZydisMnemonic_ = 1095
	ZydisMnemonicVfmadd132pd ZydisMnemonic_ = 1096
	ZydisMnemonicVfmadd132ph ZydisMnemonic_ = 1097
	ZydisMnemonicVfmadd132ps ZydisMnemonic_ = 1098
	ZydisMnemonicVfmadd132sd ZydisMnemonic_ = 1099
	ZydisMnemonicVfmadd132sh ZydisMnemonic_ = 1100
	ZydisMnemonicVfmadd132ss ZydisMnemonic_ = 1101
	ZydisMnemonicVfmadd213pd ZydisMnemonic_ = 1102
	ZydisMnemonicVfmadd213ph ZydisMnemonic_ = 1103
	ZydisMnemonicVfmadd213ps ZydisMnemonic_ = 1104
	ZydisMnemonicVfmadd213sd ZydisMnemonic_ = 1105
	ZydisMnemonicVfmadd213sh ZydisMnemonic_ = 1106
	ZydisMnemonicVfmadd213ss ZydisMnemonic_ = 1107
	ZydisMnemonicVfmadd231pd ZydisMnemonic_ = 1108
	ZydisMnemonicVfmadd231ph ZydisMnemonic_ = 1109
	ZydisMnemonicVfmadd231ps ZydisMnemonic_ = 1110
	ZydisMnemonicVfmadd231sd ZydisMnemonic_ = 1111
	ZydisMnemonicVfmadd231sh ZydisMnemonic_ = 1112
	ZydisMnemonicVfmadd231ss ZydisMnemonic_ = 1113
	ZydisMnemonicVfmadd233ps ZydisMnemonic_ = 1114
	ZydisMnemonicVfmaddcph ZydisMnemonic_ = 1115
	ZydisMnemonicVfmaddcsh ZydisMnemonic_ = 1116
	ZydisMnemonicVfmaddpd ZydisMnemonic_ = 1117
	ZydisMnemonicVfmaddps ZydisMnemonic_ = 1118
	ZydisMnemonicVfmaddsd ZydisMnemonic_ = 1119
	ZydisMnemonicVfmaddss ZydisMnemonic_ = 1120
	ZydisMnemonicVfmaddsub132pd ZydisMnemonic_ = 1121
	ZydisMnemonicVfmaddsub132ph ZydisMnemonic_ = 1122
	ZydisMnemonicVfmaddsub132ps ZydisMnemonic_ = 1123
	ZydisMnemonicVfmaddsub213pd ZydisMnemonic_ = 1124
	ZydisMnemonicVfmaddsub213ph ZydisMnemonic_ = 1125
	ZydisMnemonicVfmaddsub213ps ZydisMnemonic_ = 1126
	ZydisMnemonicVfmaddsub231pd ZydisMnemonic_ = 1127
	ZydisMnemonicVfmaddsub231ph ZydisMnemonic_ = 1128
	ZydisMnemonicVfmaddsub231ps ZydisMnemonic_ = 1129
	ZydisMnemonicVfmaddsubpd ZydisMnemonic_ = 1130
	ZydisMnemonicVfmaddsubps ZydisMnemonic_ = 1131
	ZydisMnemonicVfmsub132pd ZydisMnemonic_ = 1132
	ZydisMnemonicVfmsub132ph ZydisMnemonic_ = 1133
	ZydisMnemonicVfmsub132ps ZydisMnemonic_ = 1134
	ZydisMnemonicVfmsub132sd ZydisMnemonic_ = 1135
	ZydisMnemonicVfmsub132sh ZydisMnemonic_ = 1136
	ZydisMnemonicVfmsub132ss ZydisMnemonic_ = 1137
	ZydisMnemonicVfmsub213pd ZydisMnemonic_ = 1138
	ZydisMnemonicVfmsub213ph ZydisMnemonic_ = 1139
	ZydisMnemonicVfmsub213ps ZydisMnemonic_ = 1140
	ZydisMnemonicVfmsub213sd ZydisMnemonic_ = 1141
	ZydisMnemonicVfmsub213sh ZydisMnemonic_ = 1142
	ZydisMnemonicVfmsub213ss ZydisMnemonic_ = 1143
	ZydisMnemonicVfmsub231pd ZydisMnemonic_ = 1144
	ZydisMnemonicVfmsub231ph ZydisMnemonic_ = 1145
	ZydisMnemonicVfmsub231ps ZydisMnemonic_ = 1146
	ZydisMnemonicVfmsub231sd ZydisMnemonic_ = 1147
	ZydisMnemonicVfmsub231sh ZydisMnemonic_ = 1148
	ZydisMnemonicVfmsub231ss ZydisMnemonic_ = 1149
	ZydisMnemonicVfmsubadd132pd ZydisMnemonic_ = 1150
	ZydisMnemonicVfmsubadd132ph ZydisMnemonic_ = 1151
	ZydisMnemonicVfmsubadd132ps ZydisMnemonic_ = 1152
	ZydisMnemonicVfmsubadd213pd ZydisMnemonic_ = 1153
	ZydisMnemonicVfmsubadd213ph ZydisMnemonic_ = 1154
	ZydisMnemonicVfmsubadd213ps ZydisMnemonic_ = 1155
	ZydisMnemonicVfmsubadd231pd ZydisMnemonic_ = 1156
	ZydisMnemonicVfmsubadd231ph ZydisMnemonic_ = 1157
	ZydisMnemonicVfmsubadd231ps ZydisMnemonic_ = 1158
	ZydisMnemonicVfmsubaddpd ZydisMnemonic_ = 1159
	ZydisMnemonicVfmsubaddps ZydisMnemonic_ = 1160
	ZydisMnemonicVfmsubpd ZydisMnemonic_ = 1161
	ZydisMnemonicVfmsubps ZydisMnemonic_ = 1162
	ZydisMnemonicVfmsubsd ZydisMnemonic_ = 1163
	ZydisMnemonicVfmsubss ZydisMnemonic_ = 1164
	ZydisMnemonicVfmulcph ZydisMnemonic_ = 1165
	ZydisMnemonicVfmulcsh ZydisMnemonic_ = 1166
	ZydisMnemonicVfnmadd132pd ZydisMnemonic_ = 1167
	ZydisMnemonicVfnmadd132ph ZydisMnemonic_ = 1168
	ZydisMnemonicVfnmadd132ps ZydisMnemonic_ = 1169
	ZydisMnemonicVfnmadd132sd ZydisMnemonic_ = 1170
	ZydisMnemonicVfnmadd132sh ZydisMnemonic_ = 1171
	ZydisMnemonicVfnmadd132ss ZydisMnemonic_ = 1172
	ZydisMnemonicVfnmadd213pd ZydisMnemonic_ = 1173
	ZydisMnemonicVfnmadd213ph ZydisMnemonic_ = 1174
	ZydisMnemonicVfnmadd213ps ZydisMnemonic_ = 1175
	ZydisMnemonicVfnmadd213sd ZydisMnemonic_ = 1176
	ZydisMnemonicVfnmadd213sh ZydisMnemonic_ = 1177
	ZydisMnemonicVfnmadd213ss ZydisMnemonic_ = 1178
	ZydisMnemonicVfnmadd231pd ZydisMnemonic_ = 1179
	ZydisMnemonicVfnmadd231ph ZydisMnemonic_ = 1180
	ZydisMnemonicVfnmadd231ps ZydisMnemonic_ = 1181
	ZydisMnemonicVfnmadd231sd ZydisMnemonic_ = 1182
	ZydisMnemonicVfnmadd231sh ZydisMnemonic_ = 1183
	ZydisMnemonicVfnmadd231ss ZydisMnemonic_ = 1184
	ZydisMnemonicVfnmaddpd ZydisMnemonic_ = 1185
	ZydisMnemonicVfnmaddps ZydisMnemonic_ = 1186
	ZydisMnemonicVfnmaddsd ZydisMnemonic_ = 1187
	ZydisMnemonicVfnmaddss ZydisMnemonic_ = 1188
	ZydisMnemonicVfnmsub132pd ZydisMnemonic_ = 1189
	ZydisMnemonicVfnmsub132ph ZydisMnemonic_ = 1190
	ZydisMnemonicVfnmsub132ps ZydisMnemonic_ = 1191
	ZydisMnemonicVfnmsub132sd ZydisMnemonic_ = 1192
	ZydisMnemonicVfnmsub132sh ZydisMnemonic_ = 1193
	ZydisMnemonicVfnmsub132ss ZydisMnemonic_ = 1194
	ZydisMnemonicVfnmsub213pd ZydisMnemonic_ = 1195
	ZydisMnemonicVfnmsub213ph ZydisMnemonic_ = 1196
	ZydisMnemonicVfnmsub213ps ZydisMnemonic_ = 1197
	ZydisMnemonicVfnmsub213sd ZydisMnemonic_ = 1198
	ZydisMnemonicVfnmsub213sh ZydisMnemonic_ = 1199
	ZydisMnemonicVfnmsub213ss ZydisMnemonic_ = 1200
	ZydisMnemonicVfnmsub231pd ZydisMnemonic_ = 1201
	ZydisMnemonicVfnmsub231ph ZydisMnemonic_ = 1202
	ZydisMnemonicVfnmsub231ps ZydisMnemonic_ = 1203
	ZydisMnemonicVfnmsub231sd ZydisMnemonic_ = 1204
	ZydisMnemonicVfnmsub231sh ZydisMnemonic_ = 1205
	ZydisMnemonicVfnmsub231ss ZydisMnemonic_ = 1206
	ZydisMnemonicVfnmsubpd ZydisMnemonic_ = 1207
	ZydisMnemonicVfnmsubps ZydisMnemonic_ = 1208
	ZydisMnemonicVfnmsubsd ZydisMnemonic_ = 1209
	ZydisMnemonicVfnmsubss ZydisMnemonic_ = 1210
	ZydisMnemonicVfpclasspd ZydisMnemonic_ = 1211
	ZydisMnemonicVfpclassph ZydisMnemonic_ = 1212
	ZydisMnemonicVfpclassps ZydisMnemonic_ = 1213
	ZydisMnemonicVfpclasssd ZydisMnemonic_ = 1214
	ZydisMnemonicVfpclasssh ZydisMnemonic_ = 1215
	ZydisMnemonicVfpclassss ZydisMnemonic_ = 1216
	ZydisMnemonicVfrczpd ZydisMnemonic_ = 1217
	ZydisMnemonicVfrczps ZydisMnemonic_ = 1218
	ZydisMnemonicVfrczsd ZydisMnemonic_ = 1219
	ZydisMnemonicVfrczss ZydisMnemonic_ = 1220
	ZydisMnemonicVgatherdpd ZydisMnemonic_ = 1221
	ZydisMnemonicVgatherdps ZydisMnemonic_ = 1222
	ZydisMnemonicVgatherpf0dpd ZydisMnemonic_ = 1223
	ZydisMnemonicVgatherpf0dps ZydisMnemonic_ = 1224
	ZydisMnemonicVgatherpf0hintdpd ZydisMnemonic_ = 1225
	ZydisMnemonicVgatherpf0hintdps ZydisMnemonic_ = 1226
	ZydisMnemonicVgatherpf0qpd ZydisMnemonic_ = 1227
	ZydisMnemonicVgatherpf0qps ZydisMnemonic_ = 1228
	ZydisMnemonicVgatherpf1dpd ZydisMnemonic_ = 1229
	ZydisMnemonicVgatherpf1dps ZydisMnemonic_ = 1230
	ZydisMnemonicVgatherpf1qpd ZydisMnemonic_ = 1231
	ZydisMnemonicVgatherpf1qps ZydisMnemonic_ = 1232
	ZydisMnemonicVgatherqpd ZydisMnemonic_ = 1233
	ZydisMnemonicVgatherqps ZydisMnemonic_ = 1234
	ZydisMnemonicVgetexppd ZydisMnemonic_ = 1235
	ZydisMnemonicVgetexpph ZydisMnemonic_ = 1236
	ZydisMnemonicVgetexpps ZydisMnemonic_ = 1237
	ZydisMnemonicVgetexpsd ZydisMnemonic_ = 1238
	ZydisMnemonicVgetexpsh ZydisMnemonic_ = 1239
	ZydisMnemonicVgetexpss ZydisMnemonic_ = 1240
	ZydisMnemonicVgetmantpd ZydisMnemonic_ = 1241
	ZydisMnemonicVgetmantph ZydisMnemonic_ = 1242
	ZydisMnemonicVgetmantps ZydisMnemonic_ = 1243
	ZydisMnemonicVgetmantsd ZydisMnemonic_ = 1244
	ZydisMnemonicVgetmantsh ZydisMnemonic_ = 1245
	ZydisMnemonicVgetmantss ZydisMnemonic_ = 1246
	ZydisMnemonicVgf2p8affineinvqb ZydisMnemonic_ = 1247
	ZydisMnemonicVgf2p8affineqb ZydisMnemonic_ = 1248
	ZydisMnemonicVgf2p8mulb ZydisMnemonic_ = 1249
	ZydisMnemonicVgmaxabsps ZydisMnemonic_ = 1250
	ZydisMnemonicVgmaxpd ZydisMnemonic_ = 1251
	ZydisMnemonicVgmaxps ZydisMnemonic_ = 1252
	ZydisMnemonicVgminpd ZydisMnemonic_ = 1253
	ZydisMnemonicVgminps ZydisMnemonic_ = 1254
	ZydisMnemonicVhaddpd ZydisMnemonic_ = 1255
	ZydisMnemonicVhaddps ZydisMnemonic_ = 1256
	ZydisMnemonicVhsubpd ZydisMnemonic_ = 1257
	ZydisMnemonicVhsubps ZydisMnemonic_ = 1258
	ZydisMnemonicVinsertf128 ZydisMnemonic_ = 1259
	ZydisMnemonicVinsertf32x4 ZydisMnemonic_ = 1260
	ZydisMnemonicVinsertf32x8 ZydisMnemonic_ = 1261
	ZydisMnemonicVinsertf64x2 ZydisMnemonic_ = 1262
	ZydisMnemonicVinsertf64x4 ZydisMnemonic_ = 1263
	ZydisMnemonicVinserti128 ZydisMnemonic_ = 1264
	ZydisMnemonicVinserti32x4 ZydisMnemonic_ = 1265
	ZydisMnemonicVinserti32x8 ZydisMnemonic_ = 1266
	ZydisMnemonicVinserti64x2 ZydisMnemonic_ = 1267
	ZydisMnemonicVinserti64x4 ZydisMnemonic_ = 1268
	ZydisMnemonicVinsertps ZydisMnemonic_ = 1269
	ZydisMnemonicVlddqu ZydisMnemonic_ = 1270
	ZydisMnemonicVldmxcsr ZydisMnemonic_ = 1271
	ZydisMnemonicVloadunpackhd ZydisMnemonic_ = 1272
	ZydisMnemonicVloadunpackhpd ZydisMnemonic_ = 1273
	ZydisMnemonicVloadunpackhps ZydisMnemonic_ = 1274
	ZydisMnemonicVloadunpackhq ZydisMnemonic_ = 1275
	ZydisMnemonicVloadunpackld ZydisMnemonic_ = 1276
	ZydisMnemonicVloadunpacklpd ZydisMnemonic_ = 1277
	ZydisMnemonicVloadunpacklps ZydisMnemonic_ = 1278
	ZydisMnemonicVloadunpacklq ZydisMnemonic_ = 1279
	ZydisMnemonicVlog2ps ZydisMnemonic_ = 1280
	ZydisMnemonicVmaskmovdqu ZydisMnemonic_ = 1281
	ZydisMnemonicVmaskmovpd ZydisMnemonic_ = 1282
	ZydisMnemonicVmaskmovps ZydisMnemonic_ = 1283
	ZydisMnemonicVmaxpd ZydisMnemonic_ = 1284
	ZydisMnemonicVmaxph ZydisMnemonic_ = 1285
	ZydisMnemonicVmaxps ZydisMnemonic_ = 1286
	ZydisMnemonicVmaxsd ZydisMnemonic_ = 1287
	ZydisMnemonicVmaxsh ZydisMnemonic_ = 1288
	ZydisMnemonicVmaxss ZydisMnemonic_ = 1289
	ZydisMnemonicVmcall ZydisMnemonic_ = 1290
	ZydisMnemonicVmclear ZydisMnemonic_ = 1291
	ZydisMnemonicVmfunc ZydisMnemonic_ = 1292
	ZydisMnemonicVminpd ZydisMnemonic_ = 1293
	ZydisMnemonicVminph ZydisMnemonic_ = 1294
	ZydisMnemonicVminps ZydisMnemonic_ = 1295
	ZydisMnemonicVminsd ZydisMnemonic_ = 1296
	ZydisMnemonicVminsh ZydisMnemonic_ = 1297
	ZydisMnemonicVminss ZydisMnemonic_ = 1298
	ZydisMnemonicVmlaunch ZydisMnemonic_ = 1299
	ZydisMnemonicVmload ZydisMnemonic_ = 1300
	ZydisMnemonicVmmcall ZydisMnemonic_ = 1301
	ZydisMnemonicVmovapd ZydisMnemonic_ = 1302
	ZydisMnemonicVmovaps ZydisMnemonic_ = 1303
	ZydisMnemonicVmovd ZydisMnemonic_ = 1304
	ZydisMnemonicVmovddup ZydisMnemonic_ = 1305
	ZydisMnemonicVmovdqa ZydisMnemonic_ = 1306
	ZydisMnemonicVmovdqa32 ZydisMnemonic_ = 1307
	ZydisMnemonicVmovdqa64 ZydisMnemonic_ = 1308
	ZydisMnemonicVmovdqu ZydisMnemonic_ = 1309
	ZydisMnemonicVmovdqu16 ZydisMnemonic_ = 1310
	ZydisMnemonicVmovdqu32 ZydisMnemonic_ = 1311
	ZydisMnemonicVmovdqu64 ZydisMnemonic_ = 1312
	ZydisMnemonicVmovdqu8 ZydisMnemonic_ = 1313
	ZydisMnemonicVmovhlps ZydisMnemonic_ = 1314
	ZydisMnemonicVmovhpd ZydisMnemonic_ = 1315
	ZydisMnemonicVmovhps ZydisMnemonic_ = 1316
	ZydisMnemonicVmovlhps ZydisMnemonic_ = 1317
	ZydisMnemonicVmovlpd ZydisMnemonic_ = 1318
	ZydisMnemonicVmovlps ZydisMnemonic_ = 1319
	ZydisMnemonicVmovmskpd ZydisMnemonic_ = 1320
	ZydisMnemonicVmovmskps ZydisMnemonic_ = 1321
	ZydisMnemonicVmovnrapd ZydisMnemonic_ = 1322
	ZydisMnemonicVmovnraps ZydisMnemonic_ = 1323
	ZydisMnemonicVmovnrngoapd ZydisMnemonic_ = 1324
	ZydisMnemonicVmovnrngoaps ZydisMnemonic_ = 1325
	ZydisMnemonicVmovntdq ZydisMnemonic_ = 1326
	ZydisMnemonicVmovntdqa ZydisMnemonic_ = 1327
	ZydisMnemonicVmovntpd ZydisMnemonic_ = 1328
	ZydisMnemonicVmovntps ZydisMnemonic_ = 1329
	ZydisMnemonicVmovq ZydisMnemonic_ = 1330
	ZydisMnemonicVmovsd ZydisMnemonic_ = 1331
	ZydisMnemonicVmovsh ZydisMnemonic_ = 1332
	ZydisMnemonicVmovshdup ZydisMnemonic_ = 1333
	ZydisMnemonicVmovsldup ZydisMnemonic_ = 1334
	ZydisMnemonicVmovss ZydisMnemonic_ = 1335
	ZydisMnemonicVmovupd ZydisMnemonic_ = 1336
	ZydisMnemonicVmovups ZydisMnemonic_ = 1337
	ZydisMnemonicVmovw ZydisMnemonic_ = 1338
	ZydisMnemonicVmpsadbw ZydisMnemonic_ = 1339
	ZydisMnemonicVmptrld ZydisMnemonic_ = 1340
	ZydisMnemonicVmptrst ZydisMnemonic_ = 1341
	ZydisMnemonicVmread ZydisMnemonic_ = 1342
	ZydisMnemonicVmresume ZydisMnemonic_ = 1343
	ZydisMnemonicVmrun ZydisMnemonic_ = 1344
	ZydisMnemonicVmsave ZydisMnemonic_ = 1345
	ZydisMnemonicVmulpd ZydisMnemonic_ = 1346
	ZydisMnemonicVmulph ZydisMnemonic_ = 1347
	ZydisMnemonicVmulps ZydisMnemonic_ = 1348
	ZydisMnemonicVmulsd ZydisMnemonic_ = 1349
	ZydisMnemonicVmulsh ZydisMnemonic_ = 1350
	ZydisMnemonicVmulss ZydisMnemonic_ = 1351
	ZydisMnemonicVmwrite ZydisMnemonic_ = 1352
	ZydisMnemonicVmxoff ZydisMnemonic_ = 1353
	ZydisMnemonicVmxon ZydisMnemonic_ = 1354
	ZydisMnemonicVorpd ZydisMnemonic_ = 1355
	ZydisMnemonicVorps ZydisMnemonic_ = 1356
	ZydisMnemonicVp2intersectd ZydisMnemonic_ = 1357
	ZydisMnemonicVp2intersectq ZydisMnemonic_ = 1358
	ZydisMnemonicVp4dpwssd ZydisMnemonic_ = 1359
	ZydisMnemonicVp4dpwssds ZydisMnemonic_ = 1360
	ZydisMnemonicVpabsb ZydisMnemonic_ = 1361
	ZydisMnemonicVpabsd ZydisMnemonic_ = 1362
	ZydisMnemonicVpabsq ZydisMnemonic_ = 1363
	ZydisMnemonicVpabsw ZydisMnemonic_ = 1364
	ZydisMnemonicVpackssdw ZydisMnemonic_ = 1365
	ZydisMnemonicVpacksswb ZydisMnemonic_ = 1366
	ZydisMnemonicVpackstorehd ZydisMnemonic_ = 1367
	ZydisMnemonicVpackstorehpd ZydisMnemonic_ = 1368
	ZydisMnemonicVpackstorehps ZydisMnemonic_ = 1369
	ZydisMnemonicVpackstorehq ZydisMnemonic_ = 1370
	ZydisMnemonicVpackstoreld ZydisMnemonic_ = 1371
	ZydisMnemonicVpackstorelpd ZydisMnemonic_ = 1372
	ZydisMnemonicVpackstorelps ZydisMnemonic_ = 1373
	ZydisMnemonicVpackstorelq ZydisMnemonic_ = 1374
	ZydisMnemonicVpackusdw ZydisMnemonic_ = 1375
	ZydisMnemonicVpackuswb ZydisMnemonic_ = 1376
	ZydisMnemonicVpadcd ZydisMnemonic_ = 1377
	ZydisMnemonicVpaddb ZydisMnemonic_ = 1378
	ZydisMnemonicVpaddd ZydisMnemonic_ = 1379
	ZydisMnemonicVpaddq ZydisMnemonic_ = 1380
	ZydisMnemonicVpaddsb ZydisMnemonic_ = 1381
	ZydisMnemonicVpaddsetcd ZydisMnemonic_ = 1382
	ZydisMnemonicVpaddsetsd ZydisMnemonic_ = 1383
	ZydisMnemonicVpaddsw ZydisMnemonic_ = 1384
	ZydisMnemonicVpaddusb ZydisMnemonic_ = 1385
	ZydisMnemonicVpaddusw ZydisMnemonic_ = 1386
	ZydisMnemonicVpaddw ZydisMnemonic_ = 1387
	ZydisMnemonicVpalignr ZydisMnemonic_ = 1388
	ZydisMnemonicVpand ZydisMnemonic_ = 1389
	ZydisMnemonicVpandd ZydisMnemonic_ = 1390
	ZydisMnemonicVpandn ZydisMnemonic_ = 1391
	ZydisMnemonicVpandnd ZydisMnemonic_ = 1392
	ZydisMnemonicVpandnq ZydisMnemonic_ = 1393
	ZydisMnemonicVpandq ZydisMnemonic_ = 1394
	ZydisMnemonicVpavgb ZydisMnemonic_ = 1395
	ZydisMnemonicVpavgw ZydisMnemonic_ = 1396
	ZydisMnemonicVpblendd ZydisMnemonic_ = 1397
	ZydisMnemonicVpblendmb ZydisMnemonic_ = 1398
	ZydisMnemonicVpblendmd ZydisMnemonic_ = 1399
	ZydisMnemonicVpblendmq ZydisMnemonic_ = 1400
	ZydisMnemonicVpblendmw ZydisMnemonic_ = 1401
	ZydisMnemonicVpblendvb ZydisMnemonic_ = 1402
	ZydisMnemonicVpblendw ZydisMnemonic_ = 1403
	ZydisMnemonicVpbroadcastb ZydisMnemonic_ = 1404
	ZydisMnemonicVpbroadcastd ZydisMnemonic_ = 1405
	ZydisMnemonicVpbroadcastmb2q ZydisMnemonic_ = 1406
	ZydisMnemonicVpbroadcastmw2d ZydisMnemonic_ = 1407
	ZydisMnemonicVpbroadcastq ZydisMnemonic_ = 1408
	ZydisMnemonicVpbroadcastw ZydisMnemonic_ = 1409
	ZydisMnemonicVpclmulqdq ZydisMnemonic_ = 1410
	ZydisMnemonicVpcmov ZydisMnemonic_ = 1411
	ZydisMnemonicVpcmpb ZydisMnemonic_ = 1412
	ZydisMnemonicVpcmpd ZydisMnemonic_ = 1413
	ZydisMnemonicVpcmpeqb ZydisMnemonic_ = 1414
	ZydisMnemonicVpcmpeqd ZydisMnemonic_ = 1415
	ZydisMnemonicVpcmpeqq ZydisMnemonic_ = 1416
	ZydisMnemonicVpcmpeqw ZydisMnemonic_ = 1417
	ZydisMnemonicVpcmpestri ZydisMnemonic_ = 1418
	ZydisMnemonicVpcmpestrm ZydisMnemonic_ = 1419
	ZydisMnemonicVpcmpgtb ZydisMnemonic_ = 1420
	ZydisMnemonicVpcmpgtd ZydisMnemonic_ = 1421
	ZydisMnemonicVpcmpgtq ZydisMnemonic_ = 1422
	ZydisMnemonicVpcmpgtw ZydisMnemonic_ = 1423
	ZydisMnemonicVpcmpistri ZydisMnemonic_ = 1424
	ZydisMnemonicVpcmpistrm ZydisMnemonic_ = 1425
	ZydisMnemonicVpcmpltd ZydisMnemonic_ = 1426
	ZydisMnemonicVpcmpq ZydisMnemonic_ = 1427
	ZydisMnemonicVpcmpub ZydisMnemonic_ = 1428
	ZydisMnemonicVpcmpud ZydisMnemonic_ = 1429
	ZydisMnemonicVpcmpuq ZydisMnemonic_ = 1430
	ZydisMnemonicVpcmpuw ZydisMnemonic_ = 1431
	ZydisMnemonicVpcmpw ZydisMnemonic_ = 1432
	ZydisMnemonicVpcomb ZydisMnemonic_ = 1433
	ZydisMnemonicVpcomd ZydisMnemonic_ = 1434
	ZydisMnemonicVpcompressb ZydisMnemonic_ = 1435
	ZydisMnemonicVpcompressd ZydisMnemonic_ = 1436
	ZydisMnemonicVpcompressq ZydisMnemonic_ = 1437
	ZydisMnemonicVpcompressw ZydisMnemonic_ = 1438
	ZydisMnemonicVpcomq ZydisMnemonic_ = 1439
	ZydisMnemonicVpcomub ZydisMnemonic_ = 1440
	ZydisMnemonicVpcomud ZydisMnemonic_ = 1441
	ZydisMnemonicVpcomuq ZydisMnemonic_ = 1442
	ZydisMnemonicVpcomuw ZydisMnemonic_ = 1443
	ZydisMnemonicVpcomw ZydisMnemonic_ = 1444
	ZydisMnemonicVpconflictd ZydisMnemonic_ = 1445
	ZydisMnemonicVpconflictq ZydisMnemonic_ = 1446
	ZydisMnemonicVpdpbssd ZydisMnemonic_ = 1447
	ZydisMnemonicVpdpbssds ZydisMnemonic_ = 1448
	ZydisMnemonicVpdpbsud ZydisMnemonic_ = 1449
	ZydisMnemonicVpdpbsuds ZydisMnemonic_ = 1450
	ZydisMnemonicVpdpbusd ZydisMnemonic_ = 1451
	ZydisMnemonicVpdpbusds ZydisMnemonic_ = 1452
	ZydisMnemonicVpdpbuud ZydisMnemonic_ = 1453
	ZydisMnemonicVpdpbuuds ZydisMnemonic_ = 1454
	ZydisMnemonicVpdpwssd ZydisMnemonic_ = 1455
	ZydisMnemonicVpdpwssds ZydisMnemonic_ = 1456
	ZydisMnemonicVpdpwsud ZydisMnemonic_ = 1457
	ZydisMnemonicVpdpwsuds ZydisMnemonic_ = 1458
	ZydisMnemonicVpdpwusd ZydisMnemonic_ = 1459
	ZydisMnemonicVpdpwusds ZydisMnemonic_ = 1460
	ZydisMnemonicVpdpwuud ZydisMnemonic_ = 1461
	ZydisMnemonicVpdpwuuds ZydisMnemonic_ = 1462
	ZydisMnemonicVperm2f128 ZydisMnemonic_ = 1463
	ZydisMnemonicVperm2i128 ZydisMnemonic_ = 1464
	ZydisMnemonicVpermb ZydisMnemonic_ = 1465
	ZydisMnemonicVpermd ZydisMnemonic_ = 1466
	ZydisMnemonicVpermf32x4 ZydisMnemonic_ = 1467
	ZydisMnemonicVpermi2b ZydisMnemonic_ = 1468
	ZydisMnemonicVpermi2d ZydisMnemonic_ = 1469
	ZydisMnemonicVpermi2pd ZydisMnemonic_ = 1470
	ZydisMnemonicVpermi2ps ZydisMnemonic_ = 1471
	ZydisMnemonicVpermi2q ZydisMnemonic_ = 1472
	ZydisMnemonicVpermi2w ZydisMnemonic_ = 1473
	ZydisMnemonicVpermil2pd ZydisMnemonic_ = 1474
	ZydisMnemonicVpermil2ps ZydisMnemonic_ = 1475
	ZydisMnemonicVpermilpd ZydisMnemonic_ = 1476
	ZydisMnemonicVpermilps ZydisMnemonic_ = 1477
	ZydisMnemonicVpermpd ZydisMnemonic_ = 1478
	ZydisMnemonicVpermps ZydisMnemonic_ = 1479
	ZydisMnemonicVpermq ZydisMnemonic_ = 1480
	ZydisMnemonicVpermt2b ZydisMnemonic_ = 1481
	ZydisMnemonicVpermt2d ZydisMnemonic_ = 1482
	ZydisMnemonicVpermt2pd ZydisMnemonic_ = 1483
	ZydisMnemonicVpermt2ps ZydisMnemonic_ = 1484
	ZydisMnemonicVpermt2q ZydisMnemonic_ = 1485
	ZydisMnemonicVpermt2w ZydisMnemonic_ = 1486
	ZydisMnemonicVpermw ZydisMnemonic_ = 1487
	ZydisMnemonicVpexpandb ZydisMnemonic_ = 1488
	ZydisMnemonicVpexpandd ZydisMnemonic_ = 1489
	ZydisMnemonicVpexpandq ZydisMnemonic_ = 1490
	ZydisMnemonicVpexpandw ZydisMnemonic_ = 1491
	ZydisMnemonicVpextrb ZydisMnemonic_ = 1492
	ZydisMnemonicVpextrd ZydisMnemonic_ = 1493
	ZydisMnemonicVpextrq ZydisMnemonic_ = 1494
	ZydisMnemonicVpextrw ZydisMnemonic_ = 1495
	ZydisMnemonicVpgatherdd ZydisMnemonic_ = 1496
	ZydisMnemonicVpgatherdq ZydisMnemonic_ = 1497
	ZydisMnemonicVpgatherqd ZydisMnemonic_ = 1498
	ZydisMnemonicVpgatherqq ZydisMnemonic_ = 1499
	ZydisMnemonicVphaddbd ZydisMnemonic_ = 1500
	ZydisMnemonicVphaddbq ZydisMnemonic_ = 1501
	ZydisMnemonicVphaddbw ZydisMnemonic_ = 1502
	ZydisMnemonicVphaddd ZydisMnemonic_ = 1503
	ZydisMnemonicVphadddq ZydisMnemonic_ = 1504
	ZydisMnemonicVphaddsw ZydisMnemonic_ = 1505
	ZydisMnemonicVphaddubd ZydisMnemonic_ = 1506
	ZydisMnemonicVphaddubq ZydisMnemonic_ = 1507
	ZydisMnemonicVphaddubw ZydisMnemonic_ = 1508
	ZydisMnemonicVphaddudq ZydisMnemonic_ = 1509
	ZydisMnemonicVphadduwd ZydisMnemonic_ = 1510
	ZydisMnemonicVphadduwq ZydisMnemonic_ = 1511
	ZydisMnemonicVphaddw ZydisMnemonic_ = 1512
	ZydisMnemonicVphaddwd ZydisMnemonic_ = 1513
	ZydisMnemonicVphaddwq ZydisMnemonic_ = 1514
	ZydisMnemonicVphminposuw ZydisMnemonic_ = 1515
	ZydisMnemonicVphsubbw ZydisMnemonic_ = 1516
	ZydisMnemonicVphsubd ZydisMnemonic_ = 1517
	ZydisMnemonicVphsubdq ZydisMnemonic_ = 1518
	ZydisMnemonicVphsubsw ZydisMnemonic_ = 1519
	ZydisMnemonicVphsubw ZydisMnemonic_ = 1520
	ZydisMnemonicVphsubwd ZydisMnemonic_ = 1521
	ZydisMnemonicVpinsrb ZydisMnemonic_ = 1522
	ZydisMnemonicVpinsrd ZydisMnemonic_ = 1523
	ZydisMnemonicVpinsrq ZydisMnemonic_ = 1524
	ZydisMnemonicVpinsrw ZydisMnemonic_ = 1525
	ZydisMnemonicVplzcntd ZydisMnemonic_ = 1526
	ZydisMnemonicVplzcntq ZydisMnemonic_ = 1527
	ZydisMnemonicVpmacsdd ZydisMnemonic_ = 1528
	ZydisMnemonicVpmacsdqh ZydisMnemonic_ = 1529
	ZydisMnemonicVpmacsdql ZydisMnemonic_ = 1530
	ZydisMnemonicVpmacssdd ZydisMnemonic_ = 1531
	ZydisMnemonicVpmacssdqh ZydisMnemonic_ = 1532
	ZydisMnemonicVpmacssdql ZydisMnemonic_ = 1533
	ZydisMnemonicVpmacsswd ZydisMnemonic_ = 1534
	ZydisMnemonicVpmacssww ZydisMnemonic_ = 1535
	ZydisMnemonicVpmacswd ZydisMnemonic_ = 1536
	ZydisMnemonicVpmacsww ZydisMnemonic_ = 1537
	ZydisMnemonicVpmadcsswd ZydisMnemonic_ = 1538
	ZydisMnemonicVpmadcswd ZydisMnemonic_ = 1539
	ZydisMnemonicVpmadd231d ZydisMnemonic_ = 1540
	ZydisMnemonicVpmadd233d ZydisMnemonic_ = 1541
	ZydisMnemonicVpmadd52huq ZydisMnemonic_ = 1542
	ZydisMnemonicVpmadd52luq ZydisMnemonic_ = 1543
	ZydisMnemonicVpmaddubsw ZydisMnemonic_ = 1544
	ZydisMnemonicVpmaddwd ZydisMnemonic_ = 1545
	ZydisMnemonicVpmaskmovd ZydisMnemonic_ = 1546
	ZydisMnemonicVpmaskmovq ZydisMnemonic_ = 1547
	ZydisMnemonicVpmaxsb ZydisMnemonic_ = 1548
	ZydisMnemonicVpmaxsd ZydisMnemonic_ = 1549
	ZydisMnemonicVpmaxsq ZydisMnemonic_ = 1550
	ZydisMnemonicVpmaxsw ZydisMnemonic_ = 1551
	ZydisMnemonicVpmaxub ZydisMnemonic_ = 1552
	ZydisMnemonicVpmaxud ZydisMnemonic_ = 1553
	ZydisMnemonicVpmaxuq ZydisMnemonic_ = 1554
	ZydisMnemonicVpmaxuw ZydisMnemonic_ = 1555
	ZydisMnemonicVpminsb ZydisMnemonic_ = 1556
	ZydisMnemonicVpminsd ZydisMnemonic_ = 1557
	ZydisMnemonicVpminsq ZydisMnemonic_ = 1558
	ZydisMnemonicVpminsw ZydisMnemonic_ = 1559
	ZydisMnemonicVpminub ZydisMnemonic_ = 1560
	ZydisMnemonicVpminud ZydisMnemonic_ = 1561
	ZydisMnemonicVpminuq ZydisMnemonic_ = 1562
	ZydisMnemonicVpminuw ZydisMnemonic_ = 1563
	ZydisMnemonicVpmovb2m ZydisMnemonic_ = 1564
	ZydisMnemonicVpmovd2m ZydisMnemonic_ = 1565
	ZydisMnemonicVpmovdb ZydisMnemonic_ = 1566
	ZydisMnemonicVpmovdw ZydisMnemonic_ = 1567
	ZydisMnemonicVpmovm2b ZydisMnemonic_ = 1568
	ZydisMnemonicVpmovm2d ZydisMnemonic_ = 1569
	ZydisMnemonicVpmovm2q ZydisMnemonic_ = 1570
	ZydisMnemonicVpmovm2w ZydisMnemonic_ = 1571
	ZydisMnemonicVpmovmskb ZydisMnemonic_ = 1572
	ZydisMnemonicVpmovq2m ZydisMnemonic_ = 1573
	ZydisMnemonicVpmovqb ZydisMnemonic_ = 1574
	ZydisMnemonicVpmovqd ZydisMnemonic_ = 1575
	ZydisMnemonicVpmovqw ZydisMnemonic_ = 1576
	ZydisMnemonicVpmovsdb ZydisMnemonic_ = 1577
	ZydisMnemonicVpmovsdw ZydisMnemonic_ = 1578
	ZydisMnemonicVpmovsqb ZydisMnemonic_ = 1579
	ZydisMnemonicVpmovsqd ZydisMnemonic_ = 1580
	ZydisMnemonicVpmovsqw ZydisMnemonic_ = 1581
	ZydisMnemonicVpmovswb ZydisMnemonic_ = 1582
	ZydisMnemonicVpmovsxbd ZydisMnemonic_ = 1583
	ZydisMnemonicVpmovsxbq ZydisMnemonic_ = 1584
	ZydisMnemonicVpmovsxbw ZydisMnemonic_ = 1585
	ZydisMnemonicVpmovsxdq ZydisMnemonic_ = 1586
	ZydisMnemonicVpmovsxwd ZydisMnemonic_ = 1587
	ZydisMnemonicVpmovsxwq ZydisMnemonic_ = 1588
	ZydisMnemonicVpmovusdb ZydisMnemonic_ = 1589
	ZydisMnemonicVpmovusdw ZydisMnemonic_ = 1590
	ZydisMnemonicVpmovusqb ZydisMnemonic_ = 1591
	ZydisMnemonicVpmovusqd ZydisMnemonic_ = 1592
	ZydisMnemonicVpmovusqw ZydisMnemonic_ = 1593
	ZydisMnemonicVpmovuswb ZydisMnemonic_ = 1594
	ZydisMnemonicVpmovw2m ZydisMnemonic_ = 1595
	ZydisMnemonicVpmovwb ZydisMnemonic_ = 1596
	ZydisMnemonicVpmovzxbd ZydisMnemonic_ = 1597
	ZydisMnemonicVpmovzxbq ZydisMnemonic_ = 1598
	ZydisMnemonicVpmovzxbw ZydisMnemonic_ = 1599
	ZydisMnemonicVpmovzxdq ZydisMnemonic_ = 1600
	ZydisMnemonicVpmovzxwd ZydisMnemonic_ = 1601
	ZydisMnemonicVpmovzxwq ZydisMnemonic_ = 1602
	ZydisMnemonicVpmuldq ZydisMnemonic_ = 1603
	ZydisMnemonicVpmulhd ZydisMnemonic_ = 1604
	ZydisMnemonicVpmulhrsw ZydisMnemonic_ = 1605
	ZydisMnemonicVpmulhud ZydisMnemonic_ = 1606
	ZydisMnemonicVpmulhuw ZydisMnemonic_ = 1607
	ZydisMnemonicVpmulhw ZydisMnemonic_ = 1608
	ZydisMnemonicVpmulld ZydisMnemonic_ = 1609
	ZydisMnemonicVpmullq ZydisMnemonic_ = 1610
	ZydisMnemonicVpmullw ZydisMnemonic_ = 1611
	ZydisMnemonicVpmultishiftqb ZydisMnemonic_ = 1612
	ZydisMnemonicVpmuludq ZydisMnemonic_ = 1613
	ZydisMnemonicVpopcntb ZydisMnemonic_ = 1614
	ZydisMnemonicVpopcntd ZydisMnemonic_ = 1615
	ZydisMnemonicVpopcntq ZydisMnemonic_ = 1616
	ZydisMnemonicVpopcntw ZydisMnemonic_ = 1617
	ZydisMnemonicVpor ZydisMnemonic_ = 1618
	ZydisMnemonicVpord ZydisMnemonic_ = 1619
	ZydisMnemonicVporq ZydisMnemonic_ = 1620
	ZydisMnemonicVpperm ZydisMnemonic_ = 1621
	ZydisMnemonicVprefetch0 ZydisMnemonic_ = 1622
	ZydisMnemonicVprefetch1 ZydisMnemonic_ = 1623
	ZydisMnemonicVprefetch2 ZydisMnemonic_ = 1624
	ZydisMnemonicVprefetche0 ZydisMnemonic_ = 1625
	ZydisMnemonicVprefetche1 ZydisMnemonic_ = 1626
	ZydisMnemonicVprefetche2 ZydisMnemonic_ = 1627
	ZydisMnemonicVprefetchenta ZydisMnemonic_ = 1628
	ZydisMnemonicVprefetchnta ZydisMnemonic_ = 1629
	ZydisMnemonicVprold ZydisMnemonic_ = 1630
	ZydisMnemonicVprolq ZydisMnemonic_ = 1631
	ZydisMnemonicVprolvd ZydisMnemonic_ = 1632
	ZydisMnemonicVprolvq ZydisMnemonic_ = 1633
	ZydisMnemonicVprord ZydisMnemonic_ = 1634
	ZydisMnemonicVprorq ZydisMnemonic_ = 1635
	ZydisMnemonicVprorvd ZydisMnemonic_ = 1636
	ZydisMnemonicVprorvq ZydisMnemonic_ = 1637
	ZydisMnemonicVprotb ZydisMnemonic_ = 1638
	ZydisMnemonicVprotd ZydisMnemonic_ = 1639
	ZydisMnemonicVprotq ZydisMnemonic_ = 1640
	ZydisMnemonicVprotw ZydisMnemonic_ = 1641
	ZydisMnemonicVpsadbw ZydisMnemonic_ = 1642
	ZydisMnemonicVpsbbd ZydisMnemonic_ = 1643
	ZydisMnemonicVpsbbrd ZydisMnemonic_ = 1644
	ZydisMnemonicVpscatterdd ZydisMnemonic_ = 1645
	ZydisMnemonicVpscatterdq ZydisMnemonic_ = 1646
	ZydisMnemonicVpscatterqd ZydisMnemonic_ = 1647
	ZydisMnemonicVpscatterqq ZydisMnemonic_ = 1648
	ZydisMnemonicVpshab ZydisMnemonic_ = 1649
	ZydisMnemonicVpshad ZydisMnemonic_ = 1650
	ZydisMnemonicVpshaq ZydisMnemonic_ = 1651
	ZydisMnemonicVpshaw ZydisMnemonic_ = 1652
	ZydisMnemonicVpshlb ZydisMnemonic_ = 1653
	ZydisMnemonicVpshld ZydisMnemonic_ = 1654
	ZydisMnemonicVpshldd ZydisMnemonic_ = 1655
	ZydisMnemonicVpshldq ZydisMnemonic_ = 1656
	ZydisMnemonicVpshldvd ZydisMnemonic_ = 1657
	ZydisMnemonicVpshldvq ZydisMnemonic_ = 1658
	ZydisMnemonicVpshldvw ZydisMnemonic_ = 1659
	ZydisMnemonicVpshldw ZydisMnemonic_ = 1660
	ZydisMnemonicVpshlq ZydisMnemonic_ = 1661
	ZydisMnemonicVpshlw ZydisMnemonic_ = 1662
	ZydisMnemonicVpshrdd ZydisMnemonic_ = 1663
	ZydisMnemonicVpshrdq ZydisMnemonic_ = 1664
	ZydisMnemonicVpshrdvd ZydisMnemonic_ = 1665
	ZydisMnemonicVpshrdvq ZydisMnemonic_ = 1666
	ZydisMnemonicVpshrdvw ZydisMnemonic_ = 1667
	ZydisMnemonicVpshrdw ZydisMnemonic_ = 1668
	ZydisMnemonicVpshufb ZydisMnemonic_ = 1669
	ZydisMnemonicVpshufbitqmb ZydisMnemonic_ = 1670
	ZydisMnemonicVpshufd ZydisMnemonic_ = 1671
	ZydisMnemonicVpshufhw ZydisMnemonic_ = 1672
	ZydisMnemonicVpshuflw ZydisMnemonic_ = 1673
	ZydisMnemonicVpsignb ZydisMnemonic_ = 1674
	ZydisMnemonicVpsignd ZydisMnemonic_ = 1675
	ZydisMnemonicVpsignw ZydisMnemonic_ = 1676
	ZydisMnemonicVpslld ZydisMnemonic_ = 1677
	ZydisMnemonicVpslldq ZydisMnemonic_ = 1678
	ZydisMnemonicVpsllq ZydisMnemonic_ = 1679
	ZydisMnemonicVpsllvd ZydisMnemonic_ = 1680
	ZydisMnemonicVpsllvq ZydisMnemonic_ = 1681
	ZydisMnemonicVpsllvw ZydisMnemonic_ = 1682
	ZydisMnemonicVpsllw ZydisMnemonic_ = 1683
	ZydisMnemonicVpsrad ZydisMnemonic_ = 1684
	ZydisMnemonicVpsraq ZydisMnemonic_ = 1685
	ZydisMnemonicVpsravd ZydisMnemonic_ = 1686
	ZydisMnemonicVpsravq ZydisMnemonic_ = 1687
	ZydisMnemonicVpsravw ZydisMnemonic_ = 1688
	ZydisMnemonicVpsraw ZydisMnemonic_ = 1689
	ZydisMnemonicVpsrld ZydisMnemonic_ = 1690
	ZydisMnemonicVpsrldq ZydisMnemonic_ = 1691
	ZydisMnemonicVpsrlq ZydisMnemonic_ = 1692
	ZydisMnemonicVpsrlvd ZydisMnemonic_ = 1693
	ZydisMnemonicVpsrlvq ZydisMnemonic_ = 1694
	ZydisMnemonicVpsrlvw ZydisMnemonic_ = 1695
	ZydisMnemonicVpsrlw ZydisMnemonic_ = 1696
	ZydisMnemonicVpsubb ZydisMnemonic_ = 1697
	ZydisMnemonicVpsubd ZydisMnemonic_ = 1698
	ZydisMnemonicVpsubq ZydisMnemonic_ = 1699
	ZydisMnemonicVpsubrd ZydisMnemonic_ = 1700
	ZydisMnemonicVpsubrsetbd ZydisMnemonic_ = 1701
	ZydisMnemonicVpsubsb ZydisMnemonic_ = 1702
	ZydisMnemonicVpsubsetbd ZydisMnemonic_ = 1703
	ZydisMnemonicVpsubsw ZydisMnemonic_ = 1704
	ZydisMnemonicVpsubusb ZydisMnemonic_ = 1705
	ZydisMnemonicVpsubusw ZydisMnemonic_ = 1706
	ZydisMnemonicVpsubw ZydisMnemonic_ = 1707
	ZydisMnemonicVpternlogd ZydisMnemonic_ = 1708
	ZydisMnemonicVpternlogq ZydisMnemonic_ = 1709
	ZydisMnemonicVptest ZydisMnemonic_ = 1710
	ZydisMnemonicVptestmb ZydisMnemonic_ = 1711
	ZydisMnemonicVptestmd ZydisMnemonic_ = 1712
	ZydisMnemonicVptestmq ZydisMnemonic_ = 1713
	ZydisMnemonicVptestmw ZydisMnemonic_ = 1714
	ZydisMnemonicVptestnmb ZydisMnemonic_ = 1715
	ZydisMnemonicVptestnmd ZydisMnemonic_ = 1716
	ZydisMnemonicVptestnmq ZydisMnemonic_ = 1717
	ZydisMnemonicVptestnmw ZydisMnemonic_ = 1718
	ZydisMnemonicVpunpckhbw ZydisMnemonic_ = 1719
	ZydisMnemonicVpunpckhdq ZydisMnemonic_ = 1720
	ZydisMnemonicVpunpckhqdq ZydisMnemonic_ = 1721
	ZydisMnemonicVpunpckhwd ZydisMnemonic_ = 1722
	ZydisMnemonicVpunpcklbw ZydisMnemonic_ = 1723
	ZydisMnemonicVpunpckldq ZydisMnemonic_ = 1724
	ZydisMnemonicVpunpcklqdq ZydisMnemonic_ = 1725
	ZydisMnemonicVpunpcklwd ZydisMnemonic_ = 1726
	ZydisMnemonicVpxor ZydisMnemonic_ = 1727
	ZydisMnemonicVpxord ZydisMnemonic_ = 1728
	ZydisMnemonicVpxorq ZydisMnemonic_ = 1729
	ZydisMnemonicVrangepd ZydisMnemonic_ = 1730
	ZydisMnemonicVrangeps ZydisMnemonic_ = 1731
	ZydisMnemonicVrangesd ZydisMnemonic_ = 1732
	ZydisMnemonicVrangess ZydisMnemonic_ = 1733
	ZydisMnemonicVrcp14pd ZydisMnemonic_ = 1734
	ZydisMnemonicVrcp14ps ZydisMnemonic_ = 1735
	ZydisMnemonicVrcp14sd ZydisMnemonic_ = 1736
	ZydisMnemonicVrcp14ss ZydisMnemonic_ = 1737
	ZydisMnemonicVrcp23ps ZydisMnemonic_ = 1738
	ZydisMnemonicVrcp28pd ZydisMnemonic_ = 1739
	ZydisMnemonicVrcp28ps ZydisMnemonic_ = 1740
	ZydisMnemonicVrcp28sd ZydisMnemonic_ = 1741
	ZydisMnemonicVrcp28ss ZydisMnemonic_ = 1742
	ZydisMnemonicVrcpph ZydisMnemonic_ = 1743
	ZydisMnemonicVrcpps ZydisMnemonic_ = 1744
	ZydisMnemonicVrcpsh ZydisMnemonic_ = 1745
	ZydisMnemonicVrcpss ZydisMnemonic_ = 1746
	ZydisMnemonicVreducepd ZydisMnemonic_ = 1747
	ZydisMnemonicVreduceph ZydisMnemonic_ = 1748
	ZydisMnemonicVreduceps ZydisMnemonic_ = 1749
	ZydisMnemonicVreducesd ZydisMnemonic_ = 1750
	ZydisMnemonicVreducesh ZydisMnemonic_ = 1751
	ZydisMnemonicVreducess ZydisMnemonic_ = 1752
	ZydisMnemonicVrndfxpntpd ZydisMnemonic_ = 1753
	ZydisMnemonicVrndfxpntps ZydisMnemonic_ = 1754
	ZydisMnemonicVrndscalepd ZydisMnemonic_ = 1755
	ZydisMnemonicVrndscaleph ZydisMnemonic_ = 1756
	ZydisMnemonicVrndscaleps ZydisMnemonic_ = 1757
	ZydisMnemonicVrndscalesd ZydisMnemonic_ = 1758
	ZydisMnemonicVrndscalesh ZydisMnemonic_ = 1759
	ZydisMnemonicVrndscaless ZydisMnemonic_ = 1760
	ZydisMnemonicVroundpd ZydisMnemonic_ = 1761
	ZydisMnemonicVroundps ZydisMnemonic_ = 1762
	ZydisMnemonicVroundsd ZydisMnemonic_ = 1763
	ZydisMnemonicVroundss ZydisMnemonic_ = 1764
	ZydisMnemonicVrsqrt14pd ZydisMnemonic_ = 1765
	ZydisMnemonicVrsqrt14ps ZydisMnemonic_ = 1766
	ZydisMnemonicVrsqrt14sd ZydisMnemonic_ = 1767
	ZydisMnemonicVrsqrt14ss ZydisMnemonic_ = 1768
	ZydisMnemonicVrsqrt23ps ZydisMnemonic_ = 1769
	ZydisMnemonicVrsqrt28pd ZydisMnemonic_ = 1770
	ZydisMnemonicVrsqrt28ps ZydisMnemonic_ = 1771
	ZydisMnemonicVrsqrt28sd ZydisMnemonic_ = 1772
	ZydisMnemonicVrsqrt28ss ZydisMnemonic_ = 1773
	ZydisMnemonicVrsqrtph ZydisMnemonic_ = 1774
	ZydisMnemonicVrsqrtps ZydisMnemonic_ = 1775
	ZydisMnemonicVrsqrtsh ZydisMnemonic_ = 1776
	ZydisMnemonicVrsqrtss ZydisMnemonic_ = 1777
	ZydisMnemonicVscalefpd ZydisMnemonic_ = 1778
	ZydisMnemonicVscalefph ZydisMnemonic_ = 1779
	ZydisMnemonicVscalefps ZydisMnemonic_ = 1780
	ZydisMnemonicVscalefsd ZydisMnemonic_ = 1781
	ZydisMnemonicVscalefsh ZydisMnemonic_ = 1782
	ZydisMnemonicVscalefss ZydisMnemonic_ = 1783
	ZydisMnemonicVscaleps ZydisMnemonic_ = 1784
	ZydisMnemonicVscatterdpd ZydisMnemonic_ = 1785
	ZydisMnemonicVscatterdps ZydisMnemonic_ = 1786
	ZydisMnemonicVscatterpf0dpd ZydisMnemonic_ = 1787
	ZydisMnemonicVscatterpf0dps ZydisMnemonic_ = 1788
	ZydisMnemonicVscatterpf0hintdpd ZydisMnemonic_ = 1789
	ZydisMnemonicVscatterpf0hintdps ZydisMnemonic_ = 1790
	ZydisMnemonicVscatterpf0qpd ZydisMnemonic_ = 1791
	ZydisMnemonicVscatterpf0qps ZydisMnemonic_ = 1792
	ZydisMnemonicVscatterpf1dpd ZydisMnemonic_ = 1793
	ZydisMnemonicVscatterpf1dps ZydisMnemonic_ = 1794
	ZydisMnemonicVscatterpf1qpd ZydisMnemonic_ = 1795
	ZydisMnemonicVscatterpf1qps ZydisMnemonic_ = 1796
	ZydisMnemonicVscatterqpd ZydisMnemonic_ = 1797
	ZydisMnemonicVscatterqps ZydisMnemonic_ = 1798
	ZydisMnemonicVsha512msg1 ZydisMnemonic_ = 1799
	ZydisMnemonicVsha512msg2 ZydisMnemonic_ = 1800
	ZydisMnemonicVsha512rnds2 ZydisMnemonic_ = 1801
	ZydisMnemonicVshuff32x4 ZydisMnemonic_ = 1802
	ZydisMnemonicVshuff64x2 ZydisMnemonic_ = 1803
	ZydisMnemonicVshufi32x4 ZydisMnemonic_ = 1804
	ZydisMnemonicVshufi64x2 ZydisMnemonic_ = 1805
	ZydisMnemonicVshufpd ZydisMnemonic_ = 1806
	ZydisMnemonicVshufps ZydisMnemonic_ = 1807
	ZydisMnemonicVsm3msg1 ZydisMnemonic_ = 1808
	ZydisMnemonicVsm3msg2 ZydisMnemonic_ = 1809
	ZydisMnemonicVsm3rnds2 ZydisMnemonic_ = 1810
	ZydisMnemonicVsm4key4 ZydisMnemonic_ = 1811
	ZydisMnemonicVsm4rnds4 ZydisMnemonic_ = 1812
	ZydisMnemonicVsqrtpd ZydisMnemonic_ = 1813
	ZydisMnemonicVsqrtph ZydisMnemonic_ = 1814
	ZydisMnemonicVsqrtps ZydisMnemonic_ = 1815
	ZydisMnemonicVsqrtsd ZydisMnemonic_ = 1816
	ZydisMnemonicVsqrtsh ZydisMnemonic_ = 1817
	ZydisMnemonicVsqrtss ZydisMnemonic_ = 1818
	ZydisMnemonicVstmxcsr ZydisMnemonic_ = 1819
	ZydisMnemonicVsubpd ZydisMnemonic_ = 1820
	ZydisMnemonicVsubph ZydisMnemonic_ = 1821
	ZydisMnemonicVsubps ZydisMnemonic_ = 1822
	ZydisMnemonicVsubrpd ZydisMnemonic_ = 1823
	ZydisMnemonicVsubrps ZydisMnemonic_ = 1824
	ZydisMnemonicVsubsd ZydisMnemonic_ = 1825
	ZydisMnemonicVsubsh ZydisMnemonic_ = 1826
	ZydisMnemonicVsubss ZydisMnemonic_ = 1827
	ZydisMnemonicVtestpd ZydisMnemonic_ = 1828
	ZydisMnemonicVtestps ZydisMnemonic_ = 1829
	ZydisMnemonicVucomisd ZydisMnemonic_ = 1830
	ZydisMnemonicVucomish ZydisMnemonic_ = 1831
	ZydisMnemonicVucomiss ZydisMnemonic_ = 1832
	ZydisMnemonicVunpckhpd ZydisMnemonic_ = 1833
	ZydisMnemonicVunpckhps ZydisMnemonic_ = 1834
	ZydisMnemonicVunpcklpd ZydisMnemonic_ = 1835
	ZydisMnemonicVunpcklps ZydisMnemonic_ = 1836
	ZydisMnemonicVxorpd ZydisMnemonic_ = 1837
	ZydisMnemonicVxorps ZydisMnemonic_ = 1838
	ZydisMnemonicVzeroall ZydisMnemonic_ = 1839
	ZydisMnemonicVzeroupper ZydisMnemonic_ = 1840
	ZydisMnemonicWbinvd ZydisMnemonic_ = 1841
	ZydisMnemonicWrfsbase ZydisMnemonic_ = 1842
	ZydisMnemonicWrgsbase ZydisMnemonic_ = 1843
	ZydisMnemonicWrmsr ZydisMnemonic_ = 1844
	ZydisMnemonicWrmsrlist ZydisMnemonic_ = 1845
	ZydisMnemonicWrmsrns ZydisMnemonic_ = 1846
	ZydisMnemonicWrpkru ZydisMnemonic_ = 1847
	ZydisMnemonicWrssd ZydisMnemonic_ = 1848
	ZydisMnemonicWrssq ZydisMnemonic_ = 1849
	ZydisMnemonicWrussd ZydisMnemonic_ = 1850
	ZydisMnemonicWrussq ZydisMnemonic_ = 1851
	ZydisMnemonicXabort ZydisMnemonic_ = 1852
	ZydisMnemonicXadd ZydisMnemonic_ = 1853
	ZydisMnemonicXbegin ZydisMnemonic_ = 1854
	ZydisMnemonicXchg ZydisMnemonic_ = 1855
	ZydisMnemonicXcryptCbc ZydisMnemonic_ = 1856
	ZydisMnemonicXcryptCfb ZydisMnemonic_ = 1857
	ZydisMnemonicXcryptCtr ZydisMnemonic_ = 1858
	ZydisMnemonicXcryptEcb ZydisMnemonic_ = 1859
	ZydisMnemonicXcryptOfb ZydisMnemonic_ = 1860
	ZydisMnemonicXend ZydisMnemonic_ = 1861
	ZydisMnemonicXgetbv ZydisMnemonic_ = 1862
	ZydisMnemonicXlat ZydisMnemonic_ = 1863
	ZydisMnemonicXor ZydisMnemonic_ = 1864
	ZydisMnemonicXorpd ZydisMnemonic_ = 1865
	ZydisMnemonicXorps ZydisMnemonic_ = 1866
	ZydisMnemonicXresldtrk ZydisMnemonic_ = 1867
	ZydisMnemonicXrstor ZydisMnemonic_ = 1868
	ZydisMnemonicXrstor64 ZydisMnemonic_ = 1869
	ZydisMnemonicXrstors ZydisMnemonic_ = 1870
	ZydisMnemonicXrstors64 ZydisMnemonic_ = 1871
	ZydisMnemonicXsave ZydisMnemonic_ = 1872
	ZydisMnemonicXsave64 ZydisMnemonic_ = 1873
	ZydisMnemonicXsavec ZydisMnemonic_ = 1874
	ZydisMnemonicXsavec64 ZydisMnemonic_ = 1875
	ZydisMnemonicXsaveopt ZydisMnemonic_ = 1876
	ZydisMnemonicXsaveopt64 ZydisMnemonic_ = 1877
	ZydisMnemonicXsaves ZydisMnemonic_ = 1878
	ZydisMnemonicXsaves64 ZydisMnemonic_ = 1879
	ZydisMnemonicXsetbv ZydisMnemonic_ = 1880
	ZydisMnemonicXsha1 ZydisMnemonic_ = 1881
	ZydisMnemonicXsha256 ZydisMnemonic_ = 1882
	ZydisMnemonicXstore ZydisMnemonic_ = 1883
	ZydisMnemonicXsusldtrk ZydisMnemonic_ = 1884
	ZydisMnemonicXtest ZydisMnemonic_ = 1885
	ZydisMnemonicMaxValue ZydisMnemonic_ = 1885
	ZydisMnemonicRequiredBits ZydisMnemonic_ = 11
)

func (z ZydisMnemonic_) String() string {
	switch z {
	case ZydisMnemonicInvalid:
		return "Zydis Mnemonic Invalid"
	case ZydisMnemonicAaa:
		return "Zydis Mnemonic Aaa"
	case ZydisMnemonicAad:
		return "Zydis Mnemonic Aad"
	case ZydisMnemonicAadd:
		return "Zydis Mnemonic Aadd"
	case ZydisMnemonicAam:
		return "Zydis Mnemonic Aam"
	case ZydisMnemonicAand:
		return "Zydis Mnemonic Aand"
	case ZydisMnemonicAas:
		return "Zydis Mnemonic Aas"
	case ZydisMnemonicAdc:
		return "Zydis Mnemonic Adc"
	case ZydisMnemonicAdcx:
		return "Zydis Mnemonic Adcx"
	case ZydisMnemonicAdd:
		return "Zydis Mnemonic Add"
	case ZydisMnemonicAddpd:
		return "Zydis Mnemonic Addpd"
	case ZydisMnemonicAddps:
		return "Zydis Mnemonic Addps"
	case ZydisMnemonicAddsd:
		return "Zydis Mnemonic Addsd"
	case ZydisMnemonicAddss:
		return "Zydis Mnemonic Addss"
	case ZydisMnemonicAddsubpd:
		return "Zydis Mnemonic Addsubpd"
	case ZydisMnemonicAddsubps:
		return "Zydis Mnemonic Addsubps"
	case ZydisMnemonicAdox:
		return "Zydis Mnemonic Adox"
	case ZydisMnemonicAesdec:
		return "Zydis Mnemonic Aesdec"
	case ZydisMnemonicAesdec128kl:
		return "Zydis Mnemonic Aesdec 128kl"
	case ZydisMnemonicAesdec256kl:
		return "Zydis Mnemonic Aesdec 256kl"
	case ZydisMnemonicAesdeclast:
		return "Zydis Mnemonic Aesdeclast"
	case ZydisMnemonicAesdecwide128kl:
		return "Zydis Mnemonic Aesdecwide 128kl"
	case ZydisMnemonicAesdecwide256kl:
		return "Zydis Mnemonic Aesdecwide 256kl"
	case ZydisMnemonicAesenc:
		return "Zydis Mnemonic Aesenc"
	case ZydisMnemonicAesenc128kl:
		return "Zydis Mnemonic Aesenc 128kl"
	case ZydisMnemonicAesenc256kl:
		return "Zydis Mnemonic Aesenc 256kl"
	case ZydisMnemonicAesenclast:
		return "Zydis Mnemonic Aesenclast"
	case ZydisMnemonicAesencwide128kl:
		return "Zydis Mnemonic Aesencwide 128kl"
	case ZydisMnemonicAesencwide256kl:
		return "Zydis Mnemonic Aesencwide 256kl"
	case ZydisMnemonicAesimc:
		return "Zydis Mnemonic Aesimc"
	case ZydisMnemonicAeskeygenassist:
		return "Zydis Mnemonic Aeskeygenassist"
	case ZydisMnemonicAnd:
		return "Zydis Mnemonic And"
	case ZydisMnemonicAndn:
		return "Zydis Mnemonic Andn"
	case ZydisMnemonicAndnpd:
		return "Zydis Mnemonic Andnpd"
	case ZydisMnemonicAndnps:
		return "Zydis Mnemonic Andnps"
	case ZydisMnemonicAndpd:
		return "Zydis Mnemonic Andpd"
	case ZydisMnemonicAndps:
		return "Zydis Mnemonic Andps"
	case ZydisMnemonicAor:
		return "Zydis Mnemonic Aor"
	case ZydisMnemonicArpl:
		return "Zydis Mnemonic Arpl"
	case ZydisMnemonicAxor:
		return "Zydis Mnemonic Axor"
	case ZydisMnemonicBextr:
		return "Zydis Mnemonic Bextr"
	case ZydisMnemonicBlcfill:
		return "Zydis Mnemonic Blcfill"
	case ZydisMnemonicBlci:
		return "Zydis Mnemonic Blci"
	case ZydisMnemonicBlcic:
		return "Zydis Mnemonic Blcic"
	case ZydisMnemonicBlcmsk:
		return "Zydis Mnemonic Blcmsk"
	case ZydisMnemonicBlcs:
		return "Zydis Mnemonic Blcs"
	case ZydisMnemonicBlendpd:
		return "Zydis Mnemonic Blendpd"
	case ZydisMnemonicBlendps:
		return "Zydis Mnemonic Blendps"
	case ZydisMnemonicBlendvpd:
		return "Zydis Mnemonic Blendvpd"
	case ZydisMnemonicBlendvps:
		return "Zydis Mnemonic Blendvps"
	case ZydisMnemonicBlsfill:
		return "Zydis Mnemonic Blsfill"
	case ZydisMnemonicBlsi:
		return "Zydis Mnemonic Blsi"
	case ZydisMnemonicBlsic:
		return "Zydis Mnemonic Blsic"
	case ZydisMnemonicBlsmsk:
		return "Zydis Mnemonic Blsmsk"
	case ZydisMnemonicBlsr:
		return "Zydis Mnemonic Blsr"
	case ZydisMnemonicBndcl:
		return "Zydis Mnemonic Bndcl"
	case ZydisMnemonicBndcn:
		return "Zydis Mnemonic Bndcn"
	case ZydisMnemonicBndcu:
		return "Zydis Mnemonic Bndcu"
	case ZydisMnemonicBndldx:
		return "Zydis Mnemonic Bndldx"
	case ZydisMnemonicBndmk:
		return "Zydis Mnemonic Bndmk"
	case ZydisMnemonicBndmov:
		return "Zydis Mnemonic Bndmov"
	case ZydisMnemonicBndstx:
		return "Zydis Mnemonic Bndstx"
	case ZydisMnemonicBound:
		return "Zydis Mnemonic Bound"
	case ZydisMnemonicBsf:
		return "Zydis Mnemonic Bsf"
	case ZydisMnemonicBsr:
		return "Zydis Mnemonic Bsr"
	case ZydisMnemonicBswap:
		return "Zydis Mnemonic Bswap"
	case ZydisMnemonicBt:
		return "Zydis Mnemonic Bt"
	case ZydisMnemonicBtc:
		return "Zydis Mnemonic Btc"
	case ZydisMnemonicBtr:
		return "Zydis Mnemonic Btr"
	case ZydisMnemonicBts:
		return "Zydis Mnemonic Bts"
	case ZydisMnemonicBzhi:
		return "Zydis Mnemonic Bzhi"
	case ZydisMnemonicCall:
		return "Zydis Mnemonic Call"
	case ZydisMnemonicCbw:
		return "Zydis Mnemonic Cbw"
	case ZydisMnemonicCcmpb:
		return "Zydis Mnemonic Ccmpb"
	case ZydisMnemonicCcmpbe:
		return "Zydis Mnemonic Ccmpbe"
	case ZydisMnemonicCcmpf:
		return "Zydis Mnemonic Ccmpf"
	case ZydisMnemonicCcmpl:
		return "Zydis Mnemonic Ccmpl"
	case ZydisMnemonicCcmple:
		return "Zydis Mnemonic Ccmple"
	case ZydisMnemonicCcmpnb:
		return "Zydis Mnemonic Ccmpnb"
	case ZydisMnemonicCcmpnbe:
		return "Zydis Mnemonic Ccmpnbe"
	case ZydisMnemonicCcmpnl:
		return "Zydis Mnemonic Ccmpnl"
	case ZydisMnemonicCcmpnle:
		return "Zydis Mnemonic Ccmpnle"
	case ZydisMnemonicCcmpno:
		return "Zydis Mnemonic Ccmpno"
	case ZydisMnemonicCcmpns:
		return "Zydis Mnemonic Ccmpns"
	case ZydisMnemonicCcmpnz:
		return "Zydis Mnemonic Ccmpnz"
	case ZydisMnemonicCcmpo:
		return "Zydis Mnemonic Ccmpo"
	case ZydisMnemonicCcmps:
		return "Zydis Mnemonic Ccmps"
	case ZydisMnemonicCcmpt:
		return "Zydis Mnemonic Ccmpt"
	case ZydisMnemonicCcmpz:
		return "Zydis Mnemonic Ccmpz"
	case ZydisMnemonicCdq:
		return "Zydis Mnemonic Cdq"
	case ZydisMnemonicCdqe:
		return "Zydis Mnemonic Cdqe"
	case ZydisMnemonicCfcmovb:
		return "Zydis Mnemonic Cfcmovb"
	case ZydisMnemonicCfcmovbe:
		return "Zydis Mnemonic Cfcmovbe"
	case ZydisMnemonicCfcmovl:
		return "Zydis Mnemonic Cfcmovl"
	case ZydisMnemonicCfcmovle:
		return "Zydis Mnemonic Cfcmovle"
	case ZydisMnemonicCfcmovnb:
		return "Zydis Mnemonic Cfcmovnb"
	case ZydisMnemonicCfcmovnbe:
		return "Zydis Mnemonic Cfcmovnbe"
	case ZydisMnemonicCfcmovnl:
		return "Zydis Mnemonic Cfcmovnl"
	case ZydisMnemonicCfcmovnle:
		return "Zydis Mnemonic Cfcmovnle"
	case ZydisMnemonicCfcmovno:
		return "Zydis Mnemonic Cfcmovno"
	case ZydisMnemonicCfcmovnp:
		return "Zydis Mnemonic Cfcmovnp"
	case ZydisMnemonicCfcmovns:
		return "Zydis Mnemonic Cfcmovns"
	case ZydisMnemonicCfcmovnz:
		return "Zydis Mnemonic Cfcmovnz"
	case ZydisMnemonicCfcmovo:
		return "Zydis Mnemonic Cfcmovo"
	case ZydisMnemonicCfcmovp:
		return "Zydis Mnemonic Cfcmovp"
	case ZydisMnemonicCfcmovs:
		return "Zydis Mnemonic Cfcmovs"
	case ZydisMnemonicCfcmovz:
		return "Zydis Mnemonic Cfcmovz"
	case ZydisMnemonicClac:
		return "Zydis Mnemonic Clac"
	case ZydisMnemonicClc:
		return "Zydis Mnemonic Clc"
	case ZydisMnemonicCld:
		return "Zydis Mnemonic Cld"
	case ZydisMnemonicCldemote:
		return "Zydis Mnemonic Cldemote"
	case ZydisMnemonicClevict0:
		return "Zydis Mnemonic Clevict 0"
	case ZydisMnemonicClevict1:
		return "Zydis Mnemonic Clevict 1"
	case ZydisMnemonicClflush:
		return "Zydis Mnemonic Clflush"
	case ZydisMnemonicClflushopt:
		return "Zydis Mnemonic Clflushopt"
	case ZydisMnemonicClgi:
		return "Zydis Mnemonic Clgi"
	case ZydisMnemonicCli:
		return "Zydis Mnemonic Cli"
	case ZydisMnemonicClrssbsy:
		return "Zydis Mnemonic Clrssbsy"
	case ZydisMnemonicClts:
		return "Zydis Mnemonic Clts"
	case ZydisMnemonicClui:
		return "Zydis Mnemonic Clui"
	case ZydisMnemonicClwb:
		return "Zydis Mnemonic Clwb"
	case ZydisMnemonicClzero:
		return "Zydis Mnemonic Clzero"
	case ZydisMnemonicCmc:
		return "Zydis Mnemonic Cmc"
	case ZydisMnemonicCmovb:
		return "Zydis Mnemonic Cmovb"
	case ZydisMnemonicCmovbe:
		return "Zydis Mnemonic Cmovbe"
	case ZydisMnemonicCmovl:
		return "Zydis Mnemonic Cmovl"
	case ZydisMnemonicCmovle:
		return "Zydis Mnemonic Cmovle"
	case ZydisMnemonicCmovnb:
		return "Zydis Mnemonic Cmovnb"
	case ZydisMnemonicCmovnbe:
		return "Zydis Mnemonic Cmovnbe"
	case ZydisMnemonicCmovnl:
		return "Zydis Mnemonic Cmovnl"
	case ZydisMnemonicCmovnle:
		return "Zydis Mnemonic Cmovnle"
	case ZydisMnemonicCmovno:
		return "Zydis Mnemonic Cmovno"
	case ZydisMnemonicCmovnp:
		return "Zydis Mnemonic Cmovnp"
	case ZydisMnemonicCmovns:
		return "Zydis Mnemonic Cmovns"
	case ZydisMnemonicCmovnz:
		return "Zydis Mnemonic Cmovnz"
	case ZydisMnemonicCmovo:
		return "Zydis Mnemonic Cmovo"
	case ZydisMnemonicCmovp:
		return "Zydis Mnemonic Cmovp"
	case ZydisMnemonicCmovs:
		return "Zydis Mnemonic Cmovs"
	case ZydisMnemonicCmovz:
		return "Zydis Mnemonic Cmovz"
	case ZydisMnemonicCmp:
		return "Zydis Mnemonic Cmp"
	case ZydisMnemonicCmpbexadd:
		return "Zydis Mnemonic Cmpbexadd"
	case ZydisMnemonicCmpbxadd:
		return "Zydis Mnemonic Cmpbxadd"
	case ZydisMnemonicCmplexadd:
		return "Zydis Mnemonic Cmplexadd"
	case ZydisMnemonicCmplxadd:
		return "Zydis Mnemonic Cmplxadd"
	case ZydisMnemonicCmpnbexadd:
		return "Zydis Mnemonic Cmpnbexadd"
	case ZydisMnemonicCmpnbxadd:
		return "Zydis Mnemonic Cmpnbxadd"
	case ZydisMnemonicCmpnlexadd:
		return "Zydis Mnemonic Cmpnlexadd"
	case ZydisMnemonicCmpnlxadd:
		return "Zydis Mnemonic Cmpnlxadd"
	case ZydisMnemonicCmpnoxadd:
		return "Zydis Mnemonic Cmpnoxadd"
	case ZydisMnemonicCmpnpxadd:
		return "Zydis Mnemonic Cmpnpxadd"
	case ZydisMnemonicCmpnsxadd:
		return "Zydis Mnemonic Cmpnsxadd"
	case ZydisMnemonicCmpnzxadd:
		return "Zydis Mnemonic Cmpnzxadd"
	case ZydisMnemonicCmpoxadd:
		return "Zydis Mnemonic Cmpoxadd"
	case ZydisMnemonicCmppd:
		return "Zydis Mnemonic Cmppd"
	case ZydisMnemonicCmpps:
		return "Zydis Mnemonic Cmpps"
	case ZydisMnemonicCmppxadd:
		return "Zydis Mnemonic Cmppxadd"
	case ZydisMnemonicCmpsb:
		return "Zydis Mnemonic Cmpsb"
	case ZydisMnemonicCmpsd:
		return "Zydis Mnemonic Cmpsd"
	case ZydisMnemonicCmpsq:
		return "Zydis Mnemonic Cmpsq"
	case ZydisMnemonicCmpss:
		return "Zydis Mnemonic Cmpss"
	case ZydisMnemonicCmpsw:
		return "Zydis Mnemonic Cmpsw"
	case ZydisMnemonicCmpsxadd:
		return "Zydis Mnemonic Cmpsxadd"
	case ZydisMnemonicCmpxchg:
		return "Zydis Mnemonic Cmpxchg"
	case ZydisMnemonicCmpxchg16b:
		return "Zydis Mnemonic Cmpxchg 16b"
	case ZydisMnemonicCmpxchg8b:
		return "Zydis Mnemonic Cmpxchg 8b"
	case ZydisMnemonicCmpzxadd:
		return "Zydis Mnemonic Cmpzxadd"
	case ZydisMnemonicComisd:
		return "Zydis Mnemonic Comisd"
	case ZydisMnemonicComiss:
		return "Zydis Mnemonic Comiss"
	case ZydisMnemonicCpuid:
		return "Zydis Mnemonic Cpuid"
	case ZydisMnemonicCqo:
		return "Zydis Mnemonic Cqo"
	case ZydisMnemonicCrc32:
		return "Zydis Mnemonic Crc 32"
	case ZydisMnemonicCtestb:
		return "Zydis Mnemonic Ctestb"
	case ZydisMnemonicCtestbe:
		return "Zydis Mnemonic Ctestbe"
	case ZydisMnemonicCtestf:
		return "Zydis Mnemonic Ctestf"
	case ZydisMnemonicCtestl:
		return "Zydis Mnemonic Ctestl"
	case ZydisMnemonicCtestle:
		return "Zydis Mnemonic Ctestle"
	case ZydisMnemonicCtestnb:
		return "Zydis Mnemonic Ctestnb"
	case ZydisMnemonicCtestnbe:
		return "Zydis Mnemonic Ctestnbe"
	case ZydisMnemonicCtestnl:
		return "Zydis Mnemonic Ctestnl"
	case ZydisMnemonicCtestnle:
		return "Zydis Mnemonic Ctestnle"
	case ZydisMnemonicCtestno:
		return "Zydis Mnemonic Ctestno"
	case ZydisMnemonicCtestns:
		return "Zydis Mnemonic Ctestns"
	case ZydisMnemonicCtestnz:
		return "Zydis Mnemonic Ctestnz"
	case ZydisMnemonicCtesto:
		return "Zydis Mnemonic Ctesto"
	case ZydisMnemonicCtests:
		return "Zydis Mnemonic Ctests"
	case ZydisMnemonicCtestt:
		return "Zydis Mnemonic Ctestt"
	case ZydisMnemonicCtestz:
		return "Zydis Mnemonic Ctestz"
	case ZydisMnemonicCvtdq2pd:
		return "Zydis Mnemonic Cvtdq 2pd"
	case ZydisMnemonicCvtdq2ps:
		return "Zydis Mnemonic Cvtdq 2ps"
	case ZydisMnemonicCvtpd2dq:
		return "Zydis Mnemonic Cvtpd 2dq"
	case ZydisMnemonicCvtpd2pi:
		return "Zydis Mnemonic Cvtpd 2pi"
	case ZydisMnemonicCvtpd2ps:
		return "Zydis Mnemonic Cvtpd 2ps"
	case ZydisMnemonicCvtpi2pd:
		return "Zydis Mnemonic Cvtpi 2pd"
	case ZydisMnemonicCvtpi2ps:
		return "Zydis Mnemonic Cvtpi 2ps"
	case ZydisMnemonicCvtps2dq:
		return "Zydis Mnemonic Cvtps 2dq"
	case ZydisMnemonicCvtps2pd:
		return "Zydis Mnemonic Cvtps 2pd"
	case ZydisMnemonicCvtps2pi:
		return "Zydis Mnemonic Cvtps 2pi"
	case ZydisMnemonicCvtsd2si:
		return "Zydis Mnemonic Cvtsd 2si"
	case ZydisMnemonicCvtsd2ss:
		return "Zydis Mnemonic Cvtsd 2ss"
	case ZydisMnemonicCvtsi2sd:
		return "Zydis Mnemonic Cvtsi 2sd"
	case ZydisMnemonicCvtsi2ss:
		return "Zydis Mnemonic Cvtsi 2ss"
	case ZydisMnemonicCvtss2sd:
		return "Zydis Mnemonic Cvtss 2sd"
	case ZydisMnemonicCvtss2si:
		return "Zydis Mnemonic Cvtss 2si"
	case ZydisMnemonicCvttpd2dq:
		return "Zydis Mnemonic Cvttpd 2dq"
	case ZydisMnemonicCvttpd2pi:
		return "Zydis Mnemonic Cvttpd 2pi"
	case ZydisMnemonicCvttps2dq:
		return "Zydis Mnemonic Cvttps 2dq"
	case ZydisMnemonicCvttps2pi:
		return "Zydis Mnemonic Cvttps 2pi"
	case ZydisMnemonicCvttsd2si:
		return "Zydis Mnemonic Cvttsd 2si"
	case ZydisMnemonicCvttss2si:
		return "Zydis Mnemonic Cvttss 2si"
	case ZydisMnemonicCwd:
		return "Zydis Mnemonic Cwd"
	case ZydisMnemonicCwde:
		return "Zydis Mnemonic Cwde"
	case ZydisMnemonicDaa:
		return "Zydis Mnemonic Daa"
	case ZydisMnemonicDas:
		return "Zydis Mnemonic Das"
	case ZydisMnemonicDec:
		return "Zydis Mnemonic Dec"
	case ZydisMnemonicDelay:
		return "Zydis Mnemonic Delay"
	case ZydisMnemonicDiv:
		return "Zydis Mnemonic Div"
	case ZydisMnemonicDivpd:
		return "Zydis Mnemonic Divpd"
	case ZydisMnemonicDivps:
		return "Zydis Mnemonic Divps"
	case ZydisMnemonicDivsd:
		return "Zydis Mnemonic Divsd"
	case ZydisMnemonicDivss:
		return "Zydis Mnemonic Divss"
	case ZydisMnemonicDppd:
		return "Zydis Mnemonic Dppd"
	case ZydisMnemonicDpps:
		return "Zydis Mnemonic Dpps"
	case ZydisMnemonicEmms:
		return "Zydis Mnemonic Emms"
	case ZydisMnemonicEncls:
		return "Zydis Mnemonic Encls"
	case ZydisMnemonicEnclu:
		return "Zydis Mnemonic Enclu"
	case ZydisMnemonicEnclv:
		return "Zydis Mnemonic Enclv"
	case ZydisMnemonicEncodekey128:
		return "Zydis Mnemonic Encodekey 128"
	case ZydisMnemonicEncodekey256:
		return "Zydis Mnemonic Encodekey 256"
	case ZydisMnemonicEndbr32:
		return "Zydis Mnemonic Endbr 32"
	case ZydisMnemonicEndbr64:
		return "Zydis Mnemonic Endbr 64"
	case ZydisMnemonicEnqcmd:
		return "Zydis Mnemonic Enqcmd"
	case ZydisMnemonicEnqcmds:
		return "Zydis Mnemonic Enqcmds"
	case ZydisMnemonicEnter:
		return "Zydis Mnemonic Enter"
	case ZydisMnemonicErets:
		return "Zydis Mnemonic Erets"
	case ZydisMnemonicEretu:
		return "Zydis Mnemonic Eretu"
	case ZydisMnemonicExtractps:
		return "Zydis Mnemonic Extractps"
	case ZydisMnemonicExtrq:
		return "Zydis Mnemonic Extrq"
	case ZydisMnemonicF2xm1:
		return "Zydis Mnemonic F2xm 1"
	case ZydisMnemonicFabs:
		return "Zydis Mnemonic Fabs"
	case ZydisMnemonicFadd:
		return "Zydis Mnemonic Fadd"
	case ZydisMnemonicFaddp:
		return "Zydis Mnemonic Faddp"
	case ZydisMnemonicFbld:
		return "Zydis Mnemonic Fbld"
	case ZydisMnemonicFbstp:
		return "Zydis Mnemonic Fbstp"
	case ZydisMnemonicFchs:
		return "Zydis Mnemonic Fchs"
	case ZydisMnemonicFcmovb:
		return "Zydis Mnemonic Fcmovb"
	case ZydisMnemonicFcmovbe:
		return "Zydis Mnemonic Fcmovbe"
	case ZydisMnemonicFcmove:
		return "Zydis Mnemonic Fcmove"
	case ZydisMnemonicFcmovnb:
		return "Zydis Mnemonic Fcmovnb"
	case ZydisMnemonicFcmovnbe:
		return "Zydis Mnemonic Fcmovnbe"
	case ZydisMnemonicFcmovne:
		return "Zydis Mnemonic Fcmovne"
	case ZydisMnemonicFcmovnu:
		return "Zydis Mnemonic Fcmovnu"
	case ZydisMnemonicFcmovu:
		return "Zydis Mnemonic Fcmovu"
	case ZydisMnemonicFcom:
		return "Zydis Mnemonic Fcom"
	case ZydisMnemonicFcomi:
		return "Zydis Mnemonic Fcomi"
	case ZydisMnemonicFcomip:
		return "Zydis Mnemonic Fcomip"
	case ZydisMnemonicFcomp:
		return "Zydis Mnemonic Fcomp"
	case ZydisMnemonicFcompp:
		return "Zydis Mnemonic Fcompp"
	case ZydisMnemonicFcos:
		return "Zydis Mnemonic Fcos"
	case ZydisMnemonicFdecstp:
		return "Zydis Mnemonic Fdecstp"
	case ZydisMnemonicFdisi8087Nop:
		return "Zydis Mnemonic Fdisi 8087 Nop"
	case ZydisMnemonicFdiv:
		return "Zydis Mnemonic Fdiv"
	case ZydisMnemonicFdivp:
		return "Zydis Mnemonic Fdivp"
	case ZydisMnemonicFdivr:
		return "Zydis Mnemonic Fdivr"
	case ZydisMnemonicFdivrp:
		return "Zydis Mnemonic Fdivrp"
	case ZydisMnemonicFemms:
		return "Zydis Mnemonic Femms"
	case ZydisMnemonicFeni8087Nop:
		return "Zydis Mnemonic Feni 8087 Nop"
	case ZydisMnemonicFfree:
		return "Zydis Mnemonic Ffree"
	case ZydisMnemonicFfreep:
		return "Zydis Mnemonic Ffreep"
	case ZydisMnemonicFiadd:
		return "Zydis Mnemonic Fiadd"
	case ZydisMnemonicFicom:
		return "Zydis Mnemonic Ficom"
	case ZydisMnemonicFicomp:
		return "Zydis Mnemonic Ficomp"
	case ZydisMnemonicFidiv:
		return "Zydis Mnemonic Fidiv"
	case ZydisMnemonicFidivr:
		return "Zydis Mnemonic Fidivr"
	case ZydisMnemonicFild:
		return "Zydis Mnemonic Fild"
	case ZydisMnemonicFimul:
		return "Zydis Mnemonic Fimul"
	case ZydisMnemonicFincstp:
		return "Zydis Mnemonic Fincstp"
	case ZydisMnemonicFist:
		return "Zydis Mnemonic Fist"
	case ZydisMnemonicFistp:
		return "Zydis Mnemonic Fistp"
	case ZydisMnemonicFisttp:
		return "Zydis Mnemonic Fisttp"
	case ZydisMnemonicFisub:
		return "Zydis Mnemonic Fisub"
	case ZydisMnemonicFisubr:
		return "Zydis Mnemonic Fisubr"
	case ZydisMnemonicFld:
		return "Zydis Mnemonic Fld"
	case ZydisMnemonicFld1:
		return "Zydis Mnemonic Fld 1"
	case ZydisMnemonicFldcw:
		return "Zydis Mnemonic Fldcw"
	case ZydisMnemonicFldenv:
		return "Zydis Mnemonic Fldenv"
	case ZydisMnemonicFldl2e:
		return "Zydis Mnemonic Fldl 2e"
	case ZydisMnemonicFldl2t:
		return "Zydis Mnemonic Fldl 2t"
	case ZydisMnemonicFldlg2:
		return "Zydis Mnemonic Fldlg 2"
	case ZydisMnemonicFldln2:
		return "Zydis Mnemonic Fldln 2"
	case ZydisMnemonicFldpi:
		return "Zydis Mnemonic Fldpi"
	case ZydisMnemonicFldz:
		return "Zydis Mnemonic Fldz"
	case ZydisMnemonicFmul:
		return "Zydis Mnemonic Fmul"
	case ZydisMnemonicFmulp:
		return "Zydis Mnemonic Fmulp"
	case ZydisMnemonicFnclex:
		return "Zydis Mnemonic Fnclex"
	case ZydisMnemonicFninit:
		return "Zydis Mnemonic Fninit"
	case ZydisMnemonicFnop:
		return "Zydis Mnemonic Fnop"
	case ZydisMnemonicFnsave:
		return "Zydis Mnemonic Fnsave"
	case ZydisMnemonicFnstcw:
		return "Zydis Mnemonic Fnstcw"
	case ZydisMnemonicFnstenv:
		return "Zydis Mnemonic Fnstenv"
	case ZydisMnemonicFnstsw:
		return "Zydis Mnemonic Fnstsw"
	case ZydisMnemonicFpatan:
		return "Zydis Mnemonic Fpatan"
	case ZydisMnemonicFprem:
		return "Zydis Mnemonic Fprem"
	case ZydisMnemonicFprem1:
		return "Zydis Mnemonic Fprem 1"
	case ZydisMnemonicFptan:
		return "Zydis Mnemonic Fptan"
	case ZydisMnemonicFrndint:
		return "Zydis Mnemonic Frndint"
	case ZydisMnemonicFrstor:
		return "Zydis Mnemonic Frstor"
	case ZydisMnemonicFscale:
		return "Zydis Mnemonic Fscale"
	case ZydisMnemonicFsetpm287Nop:
		return "Zydis Mnemonic Fsetpm 287 Nop"
	case ZydisMnemonicFsin:
		return "Zydis Mnemonic Fsin"
	case ZydisMnemonicFsincos:
		return "Zydis Mnemonic Fsincos"
	case ZydisMnemonicFsqrt:
		return "Zydis Mnemonic Fsqrt"
	case ZydisMnemonicFst:
		return "Zydis Mnemonic Fst"
	case ZydisMnemonicFstp:
		return "Zydis Mnemonic Fstp"
	case ZydisMnemonicFstpnce:
		return "Zydis Mnemonic Fstpnce"
	case ZydisMnemonicFsub:
		return "Zydis Mnemonic Fsub"
	case ZydisMnemonicFsubp:
		return "Zydis Mnemonic Fsubp"
	case ZydisMnemonicFsubr:
		return "Zydis Mnemonic Fsubr"
	case ZydisMnemonicFsubrp:
		return "Zydis Mnemonic Fsubrp"
	case ZydisMnemonicFtst:
		return "Zydis Mnemonic Ftst"
	case ZydisMnemonicFucom:
		return "Zydis Mnemonic Fucom"
	case ZydisMnemonicFucomi:
		return "Zydis Mnemonic Fucomi"
	case ZydisMnemonicFucomip:
		return "Zydis Mnemonic Fucomip"
	case ZydisMnemonicFucomp:
		return "Zydis Mnemonic Fucomp"
	case ZydisMnemonicFucompp:
		return "Zydis Mnemonic Fucompp"
	case ZydisMnemonicFwait:
		return "Zydis Mnemonic Fwait"
	case ZydisMnemonicFxam:
		return "Zydis Mnemonic Fxam"
	case ZydisMnemonicFxch:
		return "Zydis Mnemonic Fxch"
	case ZydisMnemonicFxrstor:
		return "Zydis Mnemonic Fxrstor"
	case ZydisMnemonicFxrstor64:
		return "Zydis Mnemonic Fxrstor 64"
	case ZydisMnemonicFxsave:
		return "Zydis Mnemonic Fxsave"
	case ZydisMnemonicFxsave64:
		return "Zydis Mnemonic Fxsave 64"
	case ZydisMnemonicFxtract:
		return "Zydis Mnemonic Fxtract"
	case ZydisMnemonicFyl2x:
		return "Zydis Mnemonic Fyl 2x"
	case ZydisMnemonicFyl2xp1:
		return "Zydis Mnemonic Fyl 2xp 1"
	case ZydisMnemonicGetsec:
		return "Zydis Mnemonic Getsec"
	case ZydisMnemonicGf2p8affineinvqb:
		return "Zydis Mnemonic Gf 2p 8affineinvqb"
	case ZydisMnemonicGf2p8affineqb:
		return "Zydis Mnemonic Gf 2p 8affineqb"
	case ZydisMnemonicGf2p8mulb:
		return "Zydis Mnemonic Gf 2p 8mulb"
	case ZydisMnemonicHaddpd:
		return "Zydis Mnemonic Haddpd"
	case ZydisMnemonicHaddps:
		return "Zydis Mnemonic Haddps"
	case ZydisMnemonicHlt:
		return "Zydis Mnemonic Hlt"
	case ZydisMnemonicHreset:
		return "Zydis Mnemonic Hreset"
	case ZydisMnemonicHsubpd:
		return "Zydis Mnemonic Hsubpd"
	case ZydisMnemonicHsubps:
		return "Zydis Mnemonic Hsubps"
	case ZydisMnemonicIdiv:
		return "Zydis Mnemonic Idiv"
	case ZydisMnemonicImul:
		return "Zydis Mnemonic Imul"
	case ZydisMnemonicImulzu:
		return "Zydis Mnemonic Imulzu"
	case ZydisMnemonicIn:
		return "Zydis Mnemonic In"
	case ZydisMnemonicInc:
		return "Zydis Mnemonic Inc"
	case ZydisMnemonicIncsspd:
		return "Zydis Mnemonic Incsspd"
	case ZydisMnemonicIncsspq:
		return "Zydis Mnemonic Incsspq"
	case ZydisMnemonicInsb:
		return "Zydis Mnemonic Insb"
	case ZydisMnemonicInsd:
		return "Zydis Mnemonic Insd"
	case ZydisMnemonicInsertps:
		return "Zydis Mnemonic Insertps"
	case ZydisMnemonicInsertq:
		return "Zydis Mnemonic Insertq"
	case ZydisMnemonicInsw:
		return "Zydis Mnemonic Insw"
	case ZydisMnemonicInt:
		return "Zydis Mnemonic Int"
	case ZydisMnemonicInt1:
		return "Zydis Mnemonic Int 1"
	case ZydisMnemonicInt3:
		return "Zydis Mnemonic Int 3"
	case ZydisMnemonicInto:
		return "Zydis Mnemonic Into"
	case ZydisMnemonicInvd:
		return "Zydis Mnemonic Invd"
	case ZydisMnemonicInvept:
		return "Zydis Mnemonic Invept"
	case ZydisMnemonicInvlpg:
		return "Zydis Mnemonic Invlpg"
	case ZydisMnemonicInvlpga:
		return "Zydis Mnemonic Invlpga"
	case ZydisMnemonicInvlpgb:
		return "Zydis Mnemonic Invlpgb"
	case ZydisMnemonicInvpcid:
		return "Zydis Mnemonic Invpcid"
	case ZydisMnemonicInvvpid:
		return "Zydis Mnemonic Invvpid"
	case ZydisMnemonicIret:
		return "Zydis Mnemonic Iret"
	case ZydisMnemonicIretd:
		return "Zydis Mnemonic Iretd"
	case ZydisMnemonicIretq:
		return "Zydis Mnemonic Iretq"
	case ZydisMnemonicJb:
		return "Zydis Mnemonic Jb"
	case ZydisMnemonicJbe:
		return "Zydis Mnemonic Jbe"
	case ZydisMnemonicJcxz:
		return "Zydis Mnemonic Jcxz"
	case ZydisMnemonicJecxz:
		return "Zydis Mnemonic Jecxz"
	case ZydisMnemonicJknzd:
		return "Zydis Mnemonic Jknzd"
	case ZydisMnemonicJkzd:
		return "Zydis Mnemonic Jkzd"
	case ZydisMnemonicJl:
		return "Zydis Mnemonic Jl"
	case ZydisMnemonicJle:
		return "Zydis Mnemonic Jle"
	case ZydisMnemonicJmp:
		return "Zydis Mnemonic Jmp"
	case ZydisMnemonicJmpabs:
		return "Zydis Mnemonic Jmpabs"
	case ZydisMnemonicJnb:
		return "Zydis Mnemonic Jnb"
	case ZydisMnemonicJnbe:
		return "Zydis Mnemonic Jnbe"
	case ZydisMnemonicJnl:
		return "Zydis Mnemonic Jnl"
	case ZydisMnemonicJnle:
		return "Zydis Mnemonic Jnle"
	case ZydisMnemonicJno:
		return "Zydis Mnemonic Jno"
	case ZydisMnemonicJnp:
		return "Zydis Mnemonic Jnp"
	case ZydisMnemonicJns:
		return "Zydis Mnemonic Jns"
	case ZydisMnemonicJnz:
		return "Zydis Mnemonic Jnz"
	case ZydisMnemonicJo:
		return "Zydis Mnemonic Jo"
	case ZydisMnemonicJp:
		return "Zydis Mnemonic Jp"
	case ZydisMnemonicJrcxz:
		return "Zydis Mnemonic Jrcxz"
	case ZydisMnemonicJs:
		return "Zydis Mnemonic Js"
	case ZydisMnemonicJz:
		return "Zydis Mnemonic Jz"
	case ZydisMnemonicKaddb:
		return "Zydis Mnemonic Kaddb"
	case ZydisMnemonicKaddd:
		return "Zydis Mnemonic Kaddd"
	case ZydisMnemonicKaddq:
		return "Zydis Mnemonic Kaddq"
	case ZydisMnemonicKaddw:
		return "Zydis Mnemonic Kaddw"
	case ZydisMnemonicKand:
		return "Zydis Mnemonic Kand"
	case ZydisMnemonicKandb:
		return "Zydis Mnemonic Kandb"
	case ZydisMnemonicKandd:
		return "Zydis Mnemonic Kandd"
	case ZydisMnemonicKandn:
		return "Zydis Mnemonic Kandn"
	case ZydisMnemonicKandnb:
		return "Zydis Mnemonic Kandnb"
	case ZydisMnemonicKandnd:
		return "Zydis Mnemonic Kandnd"
	case ZydisMnemonicKandnq:
		return "Zydis Mnemonic Kandnq"
	case ZydisMnemonicKandnr:
		return "Zydis Mnemonic Kandnr"
	case ZydisMnemonicKandnw:
		return "Zydis Mnemonic Kandnw"
	case ZydisMnemonicKandq:
		return "Zydis Mnemonic Kandq"
	case ZydisMnemonicKandw:
		return "Zydis Mnemonic Kandw"
	case ZydisMnemonicKconcath:
		return "Zydis Mnemonic Kconcath"
	case ZydisMnemonicKconcatl:
		return "Zydis Mnemonic Kconcatl"
	case ZydisMnemonicKextract:
		return "Zydis Mnemonic Kextract"
	case ZydisMnemonicKmerge2l1h:
		return "Zydis Mnemonic Kmerge 2l 1h"
	case ZydisMnemonicKmerge2l1l:
		return "Zydis Mnemonic Kmerge 2l 1l"
	case ZydisMnemonicKmov:
		return "Zydis Mnemonic Kmov"
	case ZydisMnemonicKmovb:
		return "Zydis Mnemonic Kmovb"
	case ZydisMnemonicKmovd:
		return "Zydis Mnemonic Kmovd"
	case ZydisMnemonicKmovq:
		return "Zydis Mnemonic Kmovq"
	case ZydisMnemonicKmovw:
		return "Zydis Mnemonic Kmovw"
	case ZydisMnemonicKnot:
		return "Zydis Mnemonic Knot"
	case ZydisMnemonicKnotb:
		return "Zydis Mnemonic Knotb"
	case ZydisMnemonicKnotd:
		return "Zydis Mnemonic Knotd"
	case ZydisMnemonicKnotq:
		return "Zydis Mnemonic Knotq"
	case ZydisMnemonicKnotw:
		return "Zydis Mnemonic Knotw"
	case ZydisMnemonicKor:
		return "Zydis Mnemonic Kor"
	case ZydisMnemonicKorb:
		return "Zydis Mnemonic Korb"
	case ZydisMnemonicKord:
		return "Zydis Mnemonic Kord"
	case ZydisMnemonicKorq:
		return "Zydis Mnemonic Korq"
	case ZydisMnemonicKortest:
		return "Zydis Mnemonic Kortest"
	case ZydisMnemonicKortestb:
		return "Zydis Mnemonic Kortestb"
	case ZydisMnemonicKortestd:
		return "Zydis Mnemonic Kortestd"
	case ZydisMnemonicKortestq:
		return "Zydis Mnemonic Kortestq"
	case ZydisMnemonicKortestw:
		return "Zydis Mnemonic Kortestw"
	case ZydisMnemonicKorw:
		return "Zydis Mnemonic Korw"
	case ZydisMnemonicKshiftlb:
		return "Zydis Mnemonic Kshiftlb"
	case ZydisMnemonicKshiftld:
		return "Zydis Mnemonic Kshiftld"
	case ZydisMnemonicKshiftlq:
		return "Zydis Mnemonic Kshiftlq"
	case ZydisMnemonicKshiftlw:
		return "Zydis Mnemonic Kshiftlw"
	case ZydisMnemonicKshiftrb:
		return "Zydis Mnemonic Kshiftrb"
	case ZydisMnemonicKshiftrd:
		return "Zydis Mnemonic Kshiftrd"
	case ZydisMnemonicKshiftrq:
		return "Zydis Mnemonic Kshiftrq"
	case ZydisMnemonicKshiftrw:
		return "Zydis Mnemonic Kshiftrw"
	case ZydisMnemonicKtestb:
		return "Zydis Mnemonic Ktestb"
	case ZydisMnemonicKtestd:
		return "Zydis Mnemonic Ktestd"
	case ZydisMnemonicKtestq:
		return "Zydis Mnemonic Ktestq"
	case ZydisMnemonicKtestw:
		return "Zydis Mnemonic Ktestw"
	case ZydisMnemonicKunpckbw:
		return "Zydis Mnemonic Kunpckbw"
	case ZydisMnemonicKunpckdq:
		return "Zydis Mnemonic Kunpckdq"
	case ZydisMnemonicKunpckwd:
		return "Zydis Mnemonic Kunpckwd"
	case ZydisMnemonicKxnor:
		return "Zydis Mnemonic Kxnor"
	case ZydisMnemonicKxnorb:
		return "Zydis Mnemonic Kxnorb"
	case ZydisMnemonicKxnord:
		return "Zydis Mnemonic Kxnord"
	case ZydisMnemonicKxnorq:
		return "Zydis Mnemonic Kxnorq"
	case ZydisMnemonicKxnorw:
		return "Zydis Mnemonic Kxnorw"
	case ZydisMnemonicKxor:
		return "Zydis Mnemonic Kxor"
	case ZydisMnemonicKxorb:
		return "Zydis Mnemonic Kxorb"
	case ZydisMnemonicKxord:
		return "Zydis Mnemonic Kxord"
	case ZydisMnemonicKxorq:
		return "Zydis Mnemonic Kxorq"
	case ZydisMnemonicKxorw:
		return "Zydis Mnemonic Kxorw"
	case ZydisMnemonicLahf:
		return "Zydis Mnemonic Lahf"
	case ZydisMnemonicLar:
		return "Zydis Mnemonic Lar"
	case ZydisMnemonicLddqu:
		return "Zydis Mnemonic Lddqu"
	case ZydisMnemonicLdmxcsr:
		return "Zydis Mnemonic Ldmxcsr"
	case ZydisMnemonicLds:
		return "Zydis Mnemonic Lds"
	case ZydisMnemonicLdtilecfg:
		return "Zydis Mnemonic Ldtilecfg"
	case ZydisMnemonicLea:
		return "Zydis Mnemonic Lea"
	case ZydisMnemonicLeave:
		return "Zydis Mnemonic Leave"
	case ZydisMnemonicLes:
		return "Zydis Mnemonic Les"
	case ZydisMnemonicLfence:
		return "Zydis Mnemonic Lfence"
	case ZydisMnemonicLfs:
		return "Zydis Mnemonic Lfs"
	case ZydisMnemonicLgdt:
		return "Zydis Mnemonic Lgdt"
	case ZydisMnemonicLgs:
		return "Zydis Mnemonic Lgs"
	case ZydisMnemonicLidt:
		return "Zydis Mnemonic Lidt"
	case ZydisMnemonicLkgs:
		return "Zydis Mnemonic Lkgs"
	case ZydisMnemonicLldt:
		return "Zydis Mnemonic Lldt"
	case ZydisMnemonicLlwpcb:
		return "Zydis Mnemonic Llwpcb"
	case ZydisMnemonicLmsw:
		return "Zydis Mnemonic Lmsw"
	case ZydisMnemonicLoadiwkey:
		return "Zydis Mnemonic Loadiwkey"
	case ZydisMnemonicLodsb:
		return "Zydis Mnemonic Lodsb"
	case ZydisMnemonicLodsd:
		return "Zydis Mnemonic Lodsd"
	case ZydisMnemonicLodsq:
		return "Zydis Mnemonic Lodsq"
	case ZydisMnemonicLodsw:
		return "Zydis Mnemonic Lodsw"
	case ZydisMnemonicLoop:
		return "Zydis Mnemonic Loop"
	case ZydisMnemonicLoope:
		return "Zydis Mnemonic Loope"
	case ZydisMnemonicLoopne:
		return "Zydis Mnemonic Loopne"
	case ZydisMnemonicLsl:
		return "Zydis Mnemonic Lsl"
	case ZydisMnemonicLss:
		return "Zydis Mnemonic Lss"
	case ZydisMnemonicLtr:
		return "Zydis Mnemonic Ltr"
	case ZydisMnemonicLwpins:
		return "Zydis Mnemonic Lwpins"
	case ZydisMnemonicLwpval:
		return "Zydis Mnemonic Lwpval"
	case ZydisMnemonicLzcnt:
		return "Zydis Mnemonic Lzcnt"
	case ZydisMnemonicMaskmovdqu:
		return "Zydis Mnemonic Maskmovdqu"
	case ZydisMnemonicMaskmovq:
		return "Zydis Mnemonic Maskmovq"
	case ZydisMnemonicMaxpd:
		return "Zydis Mnemonic Maxpd"
	case ZydisMnemonicMaxps:
		return "Zydis Mnemonic Maxps"
	case ZydisMnemonicMaxsd:
		return "Zydis Mnemonic Maxsd"
	case ZydisMnemonicMaxss:
		return "Zydis Mnemonic Maxss"
	case ZydisMnemonicMcommit:
		return "Zydis Mnemonic Mcommit"
	case ZydisMnemonicMfence:
		return "Zydis Mnemonic Mfence"
	case ZydisMnemonicMinpd:
		return "Zydis Mnemonic Minpd"
	case ZydisMnemonicMinps:
		return "Zydis Mnemonic Minps"
	case ZydisMnemonicMinsd:
		return "Zydis Mnemonic Minsd"
	case ZydisMnemonicMinss:
		return "Zydis Mnemonic Minss"
	case ZydisMnemonicMonitor:
		return "Zydis Mnemonic Monitor"
	case ZydisMnemonicMonitorx:
		return "Zydis Mnemonic Monitorx"
	case ZydisMnemonicMontmul:
		return "Zydis Mnemonic Montmul"
	case ZydisMnemonicMov:
		return "Zydis Mnemonic Mov"
	case ZydisMnemonicMovapd:
		return "Zydis Mnemonic Movapd"
	case ZydisMnemonicMovaps:
		return "Zydis Mnemonic Movaps"
	case ZydisMnemonicMovbe:
		return "Zydis Mnemonic Movbe"
	case ZydisMnemonicMovd:
		return "Zydis Mnemonic Movd"
	case ZydisMnemonicMovddup:
		return "Zydis Mnemonic Movddup"
	case ZydisMnemonicMovdir64b:
		return "Zydis Mnemonic Movdir 64b"
	case ZydisMnemonicMovdiri:
		return "Zydis Mnemonic Movdiri"
	case ZydisMnemonicMovdq2q:
		return "Zydis Mnemonic Movdq 2q"
	case ZydisMnemonicMovdqa:
		return "Zydis Mnemonic Movdqa"
	case ZydisMnemonicMovdqu:
		return "Zydis Mnemonic Movdqu"
	case ZydisMnemonicMovhlps:
		return "Zydis Mnemonic Movhlps"
	case ZydisMnemonicMovhpd:
		return "Zydis Mnemonic Movhpd"
	case ZydisMnemonicMovhps:
		return "Zydis Mnemonic Movhps"
	case ZydisMnemonicMovlhps:
		return "Zydis Mnemonic Movlhps"
	case ZydisMnemonicMovlpd:
		return "Zydis Mnemonic Movlpd"
	case ZydisMnemonicMovlps:
		return "Zydis Mnemonic Movlps"
	case ZydisMnemonicMovmskpd:
		return "Zydis Mnemonic Movmskpd"
	case ZydisMnemonicMovmskps:
		return "Zydis Mnemonic Movmskps"
	case ZydisMnemonicMovntdq:
		return "Zydis Mnemonic Movntdq"
	case ZydisMnemonicMovntdqa:
		return "Zydis Mnemonic Movntdqa"
	case ZydisMnemonicMovnti:
		return "Zydis Mnemonic Movnti"
	case ZydisMnemonicMovntpd:
		return "Zydis Mnemonic Movntpd"
	case ZydisMnemonicMovntps:
		return "Zydis Mnemonic Movntps"
	case ZydisMnemonicMovntq:
		return "Zydis Mnemonic Movntq"
	case ZydisMnemonicMovntsd:
		return "Zydis Mnemonic Movntsd"
	case ZydisMnemonicMovntss:
		return "Zydis Mnemonic Movntss"
	case ZydisMnemonicMovq:
		return "Zydis Mnemonic Movq"
	case ZydisMnemonicMovq2dq:
		return "Zydis Mnemonic Movq 2dq"
	case ZydisMnemonicMovsb:
		return "Zydis Mnemonic Movsb"
	case ZydisMnemonicMovsd:
		return "Zydis Mnemonic Movsd"
	case ZydisMnemonicMovshdup:
		return "Zydis Mnemonic Movshdup"
	case ZydisMnemonicMovsldup:
		return "Zydis Mnemonic Movsldup"
	case ZydisMnemonicMovsq:
		return "Zydis Mnemonic Movsq"
	case ZydisMnemonicMovss:
		return "Zydis Mnemonic Movss"
	case ZydisMnemonicMovsw:
		return "Zydis Mnemonic Movsw"
	case ZydisMnemonicMovsx:
		return "Zydis Mnemonic Movsx"
	case ZydisMnemonicMovsxd:
		return "Zydis Mnemonic Movsxd"
	case ZydisMnemonicMovupd:
		return "Zydis Mnemonic Movupd"
	case ZydisMnemonicMovups:
		return "Zydis Mnemonic Movups"
	case ZydisMnemonicMovzx:
		return "Zydis Mnemonic Movzx"
	case ZydisMnemonicMpsadbw:
		return "Zydis Mnemonic Mpsadbw"
	case ZydisMnemonicMul:
		return "Zydis Mnemonic Mul"
	case ZydisMnemonicMulpd:
		return "Zydis Mnemonic Mulpd"
	case ZydisMnemonicMulps:
		return "Zydis Mnemonic Mulps"
	case ZydisMnemonicMulsd:
		return "Zydis Mnemonic Mulsd"
	case ZydisMnemonicMulss:
		return "Zydis Mnemonic Mulss"
	case ZydisMnemonicMulx:
		return "Zydis Mnemonic Mulx"
	case ZydisMnemonicMwait:
		return "Zydis Mnemonic Mwait"
	case ZydisMnemonicMwaitx:
		return "Zydis Mnemonic Mwaitx"
	case ZydisMnemonicNeg:
		return "Zydis Mnemonic Neg"
	case ZydisMnemonicNop:
		return "Zydis Mnemonic Nop"
	case ZydisMnemonicNot:
		return "Zydis Mnemonic Not"
	case ZydisMnemonicOr:
		return "Zydis Mnemonic Or"
	case ZydisMnemonicOrpd:
		return "Zydis Mnemonic Orpd"
	case ZydisMnemonicOrps:
		return "Zydis Mnemonic Orps"
	case ZydisMnemonicOut:
		return "Zydis Mnemonic Out"
	case ZydisMnemonicOutsb:
		return "Zydis Mnemonic Outsb"
	case ZydisMnemonicOutsd:
		return "Zydis Mnemonic Outsd"
	case ZydisMnemonicOutsw:
		return "Zydis Mnemonic Outsw"
	case ZydisMnemonicPabsb:
		return "Zydis Mnemonic Pabsb"
	case ZydisMnemonicPabsd:
		return "Zydis Mnemonic Pabsd"
	case ZydisMnemonicPabsw:
		return "Zydis Mnemonic Pabsw"
	case ZydisMnemonicPackssdw:
		return "Zydis Mnemonic Packssdw"
	case ZydisMnemonicPacksswb:
		return "Zydis Mnemonic Packsswb"
	case ZydisMnemonicPackusdw:
		return "Zydis Mnemonic Packusdw"
	case ZydisMnemonicPackuswb:
		return "Zydis Mnemonic Packuswb"
	case ZydisMnemonicPaddb:
		return "Zydis Mnemonic Paddb"
	case ZydisMnemonicPaddd:
		return "Zydis Mnemonic Paddd"
	case ZydisMnemonicPaddq:
		return "Zydis Mnemonic Paddq"
	case ZydisMnemonicPaddsb:
		return "Zydis Mnemonic Paddsb"
	case ZydisMnemonicPaddsw:
		return "Zydis Mnemonic Paddsw"
	case ZydisMnemonicPaddusb:
		return "Zydis Mnemonic Paddusb"
	case ZydisMnemonicPaddusw:
		return "Zydis Mnemonic Paddusw"
	case ZydisMnemonicPaddw:
		return "Zydis Mnemonic Paddw"
	case ZydisMnemonicPalignr:
		return "Zydis Mnemonic Palignr"
	case ZydisMnemonicPand:
		return "Zydis Mnemonic Pand"
	case ZydisMnemonicPandn:
		return "Zydis Mnemonic Pandn"
	case ZydisMnemonicPause:
		return "Zydis Mnemonic Pause"
	case ZydisMnemonicPavgb:
		return "Zydis Mnemonic Pavgb"
	case ZydisMnemonicPavgusb:
		return "Zydis Mnemonic Pavgusb"
	case ZydisMnemonicPavgw:
		return "Zydis Mnemonic Pavgw"
	case ZydisMnemonicPblendvb:
		return "Zydis Mnemonic Pblendvb"
	case ZydisMnemonicPblendw:
		return "Zydis Mnemonic Pblendw"
	case ZydisMnemonicPbndkb:
		return "Zydis Mnemonic Pbndkb"
	case ZydisMnemonicPclmulqdq:
		return "Zydis Mnemonic Pclmulqdq"
	case ZydisMnemonicPcmpeqb:
		return "Zydis Mnemonic Pcmpeqb"
	case ZydisMnemonicPcmpeqd:
		return "Zydis Mnemonic Pcmpeqd"
	case ZydisMnemonicPcmpeqq:
		return "Zydis Mnemonic Pcmpeqq"
	case ZydisMnemonicPcmpeqw:
		return "Zydis Mnemonic Pcmpeqw"
	case ZydisMnemonicPcmpestri:
		return "Zydis Mnemonic Pcmpestri"
	case ZydisMnemonicPcmpestrm:
		return "Zydis Mnemonic Pcmpestrm"
	case ZydisMnemonicPcmpgtb:
		return "Zydis Mnemonic Pcmpgtb"
	case ZydisMnemonicPcmpgtd:
		return "Zydis Mnemonic Pcmpgtd"
	case ZydisMnemonicPcmpgtq:
		return "Zydis Mnemonic Pcmpgtq"
	case ZydisMnemonicPcmpgtw:
		return "Zydis Mnemonic Pcmpgtw"
	case ZydisMnemonicPcmpistri:
		return "Zydis Mnemonic Pcmpistri"
	case ZydisMnemonicPcmpistrm:
		return "Zydis Mnemonic Pcmpistrm"
	case ZydisMnemonicPcommit:
		return "Zydis Mnemonic Pcommit"
	case ZydisMnemonicPconfig:
		return "Zydis Mnemonic Pconfig"
	case ZydisMnemonicPdep:
		return "Zydis Mnemonic Pdep"
	case ZydisMnemonicPext:
		return "Zydis Mnemonic Pext"
	case ZydisMnemonicPextrb:
		return "Zydis Mnemonic Pextrb"
	case ZydisMnemonicPextrd:
		return "Zydis Mnemonic Pextrd"
	case ZydisMnemonicPextrq:
		return "Zydis Mnemonic Pextrq"
	case ZydisMnemonicPextrw:
		return "Zydis Mnemonic Pextrw"
	case ZydisMnemonicPf2id:
		return "Zydis Mnemonic Pf 2id"
	case ZydisMnemonicPf2iw:
		return "Zydis Mnemonic Pf 2iw"
	case ZydisMnemonicPfacc:
		return "Zydis Mnemonic Pfacc"
	case ZydisMnemonicPfadd:
		return "Zydis Mnemonic Pfadd"
	case ZydisMnemonicPfcmpeq:
		return "Zydis Mnemonic Pfcmpeq"
	case ZydisMnemonicPfcmpge:
		return "Zydis Mnemonic Pfcmpge"
	case ZydisMnemonicPfcmpgt:
		return "Zydis Mnemonic Pfcmpgt"
	case ZydisMnemonicPfcpit1:
		return "Zydis Mnemonic Pfcpit 1"
	case ZydisMnemonicPfmax:
		return "Zydis Mnemonic Pfmax"
	case ZydisMnemonicPfmin:
		return "Zydis Mnemonic Pfmin"
	case ZydisMnemonicPfmul:
		return "Zydis Mnemonic Pfmul"
	case ZydisMnemonicPfnacc:
		return "Zydis Mnemonic Pfnacc"
	case ZydisMnemonicPfpnacc:
		return "Zydis Mnemonic Pfpnacc"
	case ZydisMnemonicPfrcp:
		return "Zydis Mnemonic Pfrcp"
	case ZydisMnemonicPfrcpit2:
		return "Zydis Mnemonic Pfrcpit 2"
	case ZydisMnemonicPfrsqit1:
		return "Zydis Mnemonic Pfrsqit 1"
	case ZydisMnemonicPfsqrt:
		return "Zydis Mnemonic Pfsqrt"
	case ZydisMnemonicPfsub:
		return "Zydis Mnemonic Pfsub"
	case ZydisMnemonicPfsubr:
		return "Zydis Mnemonic Pfsubr"
	case ZydisMnemonicPhaddd:
		return "Zydis Mnemonic Phaddd"
	case ZydisMnemonicPhaddsw:
		return "Zydis Mnemonic Phaddsw"
	case ZydisMnemonicPhaddw:
		return "Zydis Mnemonic Phaddw"
	case ZydisMnemonicPhminposuw:
		return "Zydis Mnemonic Phminposuw"
	case ZydisMnemonicPhsubd:
		return "Zydis Mnemonic Phsubd"
	case ZydisMnemonicPhsubsw:
		return "Zydis Mnemonic Phsubsw"
	case ZydisMnemonicPhsubw:
		return "Zydis Mnemonic Phsubw"
	case ZydisMnemonicPi2fd:
		return "Zydis Mnemonic Pi 2fd"
	case ZydisMnemonicPi2fw:
		return "Zydis Mnemonic Pi 2fw"
	case ZydisMnemonicPinsrb:
		return "Zydis Mnemonic Pinsrb"
	case ZydisMnemonicPinsrd:
		return "Zydis Mnemonic Pinsrd"
	case ZydisMnemonicPinsrq:
		return "Zydis Mnemonic Pinsrq"
	case ZydisMnemonicPinsrw:
		return "Zydis Mnemonic Pinsrw"
	case ZydisMnemonicPmaddubsw:
		return "Zydis Mnemonic Pmaddubsw"
	case ZydisMnemonicPmaddwd:
		return "Zydis Mnemonic Pmaddwd"
	case ZydisMnemonicPmaxsb:
		return "Zydis Mnemonic Pmaxsb"
	case ZydisMnemonicPmaxsd:
		return "Zydis Mnemonic Pmaxsd"
	case ZydisMnemonicPmaxsw:
		return "Zydis Mnemonic Pmaxsw"
	case ZydisMnemonicPmaxub:
		return "Zydis Mnemonic Pmaxub"
	case ZydisMnemonicPmaxud:
		return "Zydis Mnemonic Pmaxud"
	case ZydisMnemonicPmaxuw:
		return "Zydis Mnemonic Pmaxuw"
	case ZydisMnemonicPminsb:
		return "Zydis Mnemonic Pminsb"
	case ZydisMnemonicPminsd:
		return "Zydis Mnemonic Pminsd"
	case ZydisMnemonicPminsw:
		return "Zydis Mnemonic Pminsw"
	case ZydisMnemonicPminub:
		return "Zydis Mnemonic Pminub"
	case ZydisMnemonicPminud:
		return "Zydis Mnemonic Pminud"
	case ZydisMnemonicPminuw:
		return "Zydis Mnemonic Pminuw"
	case ZydisMnemonicPmovmskb:
		return "Zydis Mnemonic Pmovmskb"
	case ZydisMnemonicPmovsxbd:
		return "Zydis Mnemonic Pmovsxbd"
	case ZydisMnemonicPmovsxbq:
		return "Zydis Mnemonic Pmovsxbq"
	case ZydisMnemonicPmovsxbw:
		return "Zydis Mnemonic Pmovsxbw"
	case ZydisMnemonicPmovsxdq:
		return "Zydis Mnemonic Pmovsxdq"
	case ZydisMnemonicPmovsxwd:
		return "Zydis Mnemonic Pmovsxwd"
	case ZydisMnemonicPmovsxwq:
		return "Zydis Mnemonic Pmovsxwq"
	case ZydisMnemonicPmovzxbd:
		return "Zydis Mnemonic Pmovzxbd"
	case ZydisMnemonicPmovzxbq:
		return "Zydis Mnemonic Pmovzxbq"
	case ZydisMnemonicPmovzxbw:
		return "Zydis Mnemonic Pmovzxbw"
	case ZydisMnemonicPmovzxdq:
		return "Zydis Mnemonic Pmovzxdq"
	case ZydisMnemonicPmovzxwd:
		return "Zydis Mnemonic Pmovzxwd"
	case ZydisMnemonicPmovzxwq:
		return "Zydis Mnemonic Pmovzxwq"
	case ZydisMnemonicPmuldq:
		return "Zydis Mnemonic Pmuldq"
	case ZydisMnemonicPmulhrsw:
		return "Zydis Mnemonic Pmulhrsw"
	case ZydisMnemonicPmulhrw:
		return "Zydis Mnemonic Pmulhrw"
	case ZydisMnemonicPmulhuw:
		return "Zydis Mnemonic Pmulhuw"
	case ZydisMnemonicPmulhw:
		return "Zydis Mnemonic Pmulhw"
	case ZydisMnemonicPmulld:
		return "Zydis Mnemonic Pmulld"
	case ZydisMnemonicPmullw:
		return "Zydis Mnemonic Pmullw"
	case ZydisMnemonicPmuludq:
		return "Zydis Mnemonic Pmuludq"
	case ZydisMnemonicPop:
		return "Zydis Mnemonic Pop"
	case ZydisMnemonicPop2:
		return "Zydis Mnemonic Pop 2"
	case ZydisMnemonicPop2p:
		return "Zydis Mnemonic Pop 2p"
	case ZydisMnemonicPopa:
		return "Zydis Mnemonic Popa"
	case ZydisMnemonicPopad:
		return "Zydis Mnemonic Popad"
	case ZydisMnemonicPopcnt:
		return "Zydis Mnemonic Popcnt"
	case ZydisMnemonicPopf:
		return "Zydis Mnemonic Popf"
	case ZydisMnemonicPopfd:
		return "Zydis Mnemonic Popfd"
	case ZydisMnemonicPopfq:
		return "Zydis Mnemonic Popfq"
	case ZydisMnemonicPopp:
		return "Zydis Mnemonic Popp"
	case ZydisMnemonicPor:
		return "Zydis Mnemonic Por"
	case ZydisMnemonicPrefetch:
		return "Zydis Mnemonic Prefetch"
	case ZydisMnemonicPrefetchit0:
		return "Zydis Mnemonic Prefetchit 0"
	case ZydisMnemonicPrefetchit1:
		return "Zydis Mnemonic Prefetchit 1"
	case ZydisMnemonicPrefetchnta:
		return "Zydis Mnemonic Prefetchnta"
	case ZydisMnemonicPrefetcht0:
		return "Zydis Mnemonic Prefetcht 0"
	case ZydisMnemonicPrefetcht1:
		return "Zydis Mnemonic Prefetcht 1"
	case ZydisMnemonicPrefetcht2:
		return "Zydis Mnemonic Prefetcht 2"
	case ZydisMnemonicPrefetchw:
		return "Zydis Mnemonic Prefetchw"
	case ZydisMnemonicPrefetchwt1:
		return "Zydis Mnemonic Prefetchwt 1"
	case ZydisMnemonicPsadbw:
		return "Zydis Mnemonic Psadbw"
	case ZydisMnemonicPshufb:
		return "Zydis Mnemonic Pshufb"
	case ZydisMnemonicPshufd:
		return "Zydis Mnemonic Pshufd"
	case ZydisMnemonicPshufhw:
		return "Zydis Mnemonic Pshufhw"
	case ZydisMnemonicPshuflw:
		return "Zydis Mnemonic Pshuflw"
	case ZydisMnemonicPshufw:
		return "Zydis Mnemonic Pshufw"
	case ZydisMnemonicPsignb:
		return "Zydis Mnemonic Psignb"
	case ZydisMnemonicPsignd:
		return "Zydis Mnemonic Psignd"
	case ZydisMnemonicPsignw:
		return "Zydis Mnemonic Psignw"
	case ZydisMnemonicPslld:
		return "Zydis Mnemonic Pslld"
	case ZydisMnemonicPslldq:
		return "Zydis Mnemonic Pslldq"
	case ZydisMnemonicPsllq:
		return "Zydis Mnemonic Psllq"
	case ZydisMnemonicPsllw:
		return "Zydis Mnemonic Psllw"
	case ZydisMnemonicPsmash:
		return "Zydis Mnemonic Psmash"
	case ZydisMnemonicPsrad:
		return "Zydis Mnemonic Psrad"
	case ZydisMnemonicPsraw:
		return "Zydis Mnemonic Psraw"
	case ZydisMnemonicPsrld:
		return "Zydis Mnemonic Psrld"
	case ZydisMnemonicPsrldq:
		return "Zydis Mnemonic Psrldq"
	case ZydisMnemonicPsrlq:
		return "Zydis Mnemonic Psrlq"
	case ZydisMnemonicPsrlw:
		return "Zydis Mnemonic Psrlw"
	case ZydisMnemonicPsubb:
		return "Zydis Mnemonic Psubb"
	case ZydisMnemonicPsubd:
		return "Zydis Mnemonic Psubd"
	case ZydisMnemonicPsubq:
		return "Zydis Mnemonic Psubq"
	case ZydisMnemonicPsubsb:
		return "Zydis Mnemonic Psubsb"
	case ZydisMnemonicPsubsw:
		return "Zydis Mnemonic Psubsw"
	case ZydisMnemonicPsubusb:
		return "Zydis Mnemonic Psubusb"
	case ZydisMnemonicPsubusw:
		return "Zydis Mnemonic Psubusw"
	case ZydisMnemonicPsubw:
		return "Zydis Mnemonic Psubw"
	case ZydisMnemonicPswapd:
		return "Zydis Mnemonic Pswapd"
	case ZydisMnemonicPtest:
		return "Zydis Mnemonic Ptest"
	case ZydisMnemonicPtwrite:
		return "Zydis Mnemonic Ptwrite"
	case ZydisMnemonicPunpckhbw:
		return "Zydis Mnemonic Punpckhbw"
	case ZydisMnemonicPunpckhdq:
		return "Zydis Mnemonic Punpckhdq"
	case ZydisMnemonicPunpckhqdq:
		return "Zydis Mnemonic Punpckhqdq"
	case ZydisMnemonicPunpckhwd:
		return "Zydis Mnemonic Punpckhwd"
	case ZydisMnemonicPunpcklbw:
		return "Zydis Mnemonic Punpcklbw"
	case ZydisMnemonicPunpckldq:
		return "Zydis Mnemonic Punpckldq"
	case ZydisMnemonicPunpcklqdq:
		return "Zydis Mnemonic Punpcklqdq"
	case ZydisMnemonicPunpcklwd:
		return "Zydis Mnemonic Punpcklwd"
	case ZydisMnemonicPush:
		return "Zydis Mnemonic Push"
	case ZydisMnemonicPush2:
		return "Zydis Mnemonic Push 2"
	case ZydisMnemonicPush2p:
		return "Zydis Mnemonic Push 2p"
	case ZydisMnemonicPusha:
		return "Zydis Mnemonic Pusha"
	case ZydisMnemonicPushad:
		return "Zydis Mnemonic Pushad"
	case ZydisMnemonicPushf:
		return "Zydis Mnemonic Pushf"
	case ZydisMnemonicPushfd:
		return "Zydis Mnemonic Pushfd"
	case ZydisMnemonicPushfq:
		return "Zydis Mnemonic Pushfq"
	case ZydisMnemonicPushp:
		return "Zydis Mnemonic Pushp"
	case ZydisMnemonicPvalidate:
		return "Zydis Mnemonic Pvalidate"
	case ZydisMnemonicPxor:
		return "Zydis Mnemonic Pxor"
	case ZydisMnemonicRcl:
		return "Zydis Mnemonic Rcl"
	case ZydisMnemonicRcpps:
		return "Zydis Mnemonic Rcpps"
	case ZydisMnemonicRcpss:
		return "Zydis Mnemonic Rcpss"
	case ZydisMnemonicRcr:
		return "Zydis Mnemonic Rcr"
	case ZydisMnemonicRdfsbase:
		return "Zydis Mnemonic Rdfsbase"
	case ZydisMnemonicRdgsbase:
		return "Zydis Mnemonic Rdgsbase"
	case ZydisMnemonicRdmsr:
		return "Zydis Mnemonic Rdmsr"
	case ZydisMnemonicRdmsrlist:
		return "Zydis Mnemonic Rdmsrlist"
	case ZydisMnemonicRdpid:
		return "Zydis Mnemonic Rdpid"
	case ZydisMnemonicRdpkru:
		return "Zydis Mnemonic Rdpkru"
	case ZydisMnemonicRdpmc:
		return "Zydis Mnemonic Rdpmc"
	case ZydisMnemonicRdpru:
		return "Zydis Mnemonic Rdpru"
	case ZydisMnemonicRdrand:
		return "Zydis Mnemonic Rdrand"
	case ZydisMnemonicRdseed:
		return "Zydis Mnemonic Rdseed"
	case ZydisMnemonicRdsspd:
		return "Zydis Mnemonic Rdsspd"
	case ZydisMnemonicRdsspq:
		return "Zydis Mnemonic Rdsspq"
	case ZydisMnemonicRdtsc:
		return "Zydis Mnemonic Rdtsc"
	case ZydisMnemonicRdtscp:
		return "Zydis Mnemonic Rdtscp"
	case ZydisMnemonicRet:
		return "Zydis Mnemonic Ret"
	case ZydisMnemonicRmpadjust:
		return "Zydis Mnemonic Rmpadjust"
	case ZydisMnemonicRmpupdate:
		return "Zydis Mnemonic Rmpupdate"
	case ZydisMnemonicRol:
		return "Zydis Mnemonic Rol"
	case ZydisMnemonicRor:
		return "Zydis Mnemonic Ror"
	case ZydisMnemonicRorx:
		return "Zydis Mnemonic Rorx"
	case ZydisMnemonicRoundpd:
		return "Zydis Mnemonic Roundpd"
	case ZydisMnemonicRoundps:
		return "Zydis Mnemonic Roundps"
	case ZydisMnemonicRoundsd:
		return "Zydis Mnemonic Roundsd"
	case ZydisMnemonicRoundss:
		return "Zydis Mnemonic Roundss"
	case ZydisMnemonicRsm:
		return "Zydis Mnemonic Rsm"
	case ZydisMnemonicRsqrtps:
		return "Zydis Mnemonic Rsqrtps"
	case ZydisMnemonicRsqrtss:
		return "Zydis Mnemonic Rsqrtss"
	case ZydisMnemonicRstorssp:
		return "Zydis Mnemonic Rstorssp"
	case ZydisMnemonicSahf:
		return "Zydis Mnemonic Sahf"
	case ZydisMnemonicSalc:
		return "Zydis Mnemonic Salc"
	case ZydisMnemonicSar:
		return "Zydis Mnemonic Sar"
	case ZydisMnemonicSarx:
		return "Zydis Mnemonic Sarx"
	case ZydisMnemonicSaveprevssp:
		return "Zydis Mnemonic Saveprevssp"
	case ZydisMnemonicSbb:
		return "Zydis Mnemonic Sbb"
	case ZydisMnemonicScasb:
		return "Zydis Mnemonic Scasb"
	case ZydisMnemonicScasd:
		return "Zydis Mnemonic Scasd"
	case ZydisMnemonicScasq:
		return "Zydis Mnemonic Scasq"
	case ZydisMnemonicScasw:
		return "Zydis Mnemonic Scasw"
	case ZydisMnemonicSeamcall:
		return "Zydis Mnemonic Seamcall"
	case ZydisMnemonicSeamops:
		return "Zydis Mnemonic Seamops"
	case ZydisMnemonicSeamret:
		return "Zydis Mnemonic Seamret"
	case ZydisMnemonicSenduipi:
		return "Zydis Mnemonic Senduipi"
	case ZydisMnemonicSerialize:
		return "Zydis Mnemonic Serialize"
	case ZydisMnemonicSetb:
		return "Zydis Mnemonic Setb"
	case ZydisMnemonicSetbe:
		return "Zydis Mnemonic Setbe"
	case ZydisMnemonicSetl:
		return "Zydis Mnemonic Setl"
	case ZydisMnemonicSetle:
		return "Zydis Mnemonic Setle"
	case ZydisMnemonicSetnb:
		return "Zydis Mnemonic Setnb"
	case ZydisMnemonicSetnbe:
		return "Zydis Mnemonic Setnbe"
	case ZydisMnemonicSetnl:
		return "Zydis Mnemonic Setnl"
	case ZydisMnemonicSetnle:
		return "Zydis Mnemonic Setnle"
	case ZydisMnemonicSetno:
		return "Zydis Mnemonic Setno"
	case ZydisMnemonicSetnp:
		return "Zydis Mnemonic Setnp"
	case ZydisMnemonicSetns:
		return "Zydis Mnemonic Setns"
	case ZydisMnemonicSetnz:
		return "Zydis Mnemonic Setnz"
	case ZydisMnemonicSeto:
		return "Zydis Mnemonic Seto"
	case ZydisMnemonicSetp:
		return "Zydis Mnemonic Setp"
	case ZydisMnemonicSets:
		return "Zydis Mnemonic Sets"
	case ZydisMnemonicSetssbsy:
		return "Zydis Mnemonic Setssbsy"
	case ZydisMnemonicSetz:
		return "Zydis Mnemonic Setz"
	case ZydisMnemonicSetzub:
		return "Zydis Mnemonic Setzub"
	case ZydisMnemonicSetzube:
		return "Zydis Mnemonic Setzube"
	case ZydisMnemonicSetzul:
		return "Zydis Mnemonic Setzul"
	case ZydisMnemonicSetzule:
		return "Zydis Mnemonic Setzule"
	case ZydisMnemonicSetzunb:
		return "Zydis Mnemonic Setzunb"
	case ZydisMnemonicSetzunbe:
		return "Zydis Mnemonic Setzunbe"
	case ZydisMnemonicSetzunl:
		return "Zydis Mnemonic Setzunl"
	case ZydisMnemonicSetzunle:
		return "Zydis Mnemonic Setzunle"
	case ZydisMnemonicSetzuno:
		return "Zydis Mnemonic Setzuno"
	case ZydisMnemonicSetzunp:
		return "Zydis Mnemonic Setzunp"
	case ZydisMnemonicSetzuns:
		return "Zydis Mnemonic Setzuns"
	case ZydisMnemonicSetzunz:
		return "Zydis Mnemonic Setzunz"
	case ZydisMnemonicSetzuo:
		return "Zydis Mnemonic Setzuo"
	case ZydisMnemonicSetzup:
		return "Zydis Mnemonic Setzup"
	case ZydisMnemonicSetzus:
		return "Zydis Mnemonic Setzus"
	case ZydisMnemonicSetzuz:
		return "Zydis Mnemonic Setzuz"
	case ZydisMnemonicSfence:
		return "Zydis Mnemonic Sfence"
	case ZydisMnemonicSgdt:
		return "Zydis Mnemonic Sgdt"
	case ZydisMnemonicSha1msg1:
		return "Zydis Mnemonic Sha 1msg 1"
	case ZydisMnemonicSha1msg2:
		return "Zydis Mnemonic Sha 1msg 2"
	case ZydisMnemonicSha1nexte:
		return "Zydis Mnemonic Sha 1nexte"
	case ZydisMnemonicSha1rnds4:
		return "Zydis Mnemonic Sha 1rnds 4"
	case ZydisMnemonicSha256msg1:
		return "Zydis Mnemonic Sha 256msg 1"
	case ZydisMnemonicSha256msg2:
		return "Zydis Mnemonic Sha 256msg 2"
	case ZydisMnemonicSha256rnds2:
		return "Zydis Mnemonic Sha 256rnds 2"
	case ZydisMnemonicShl:
		return "Zydis Mnemonic Shl"
	case ZydisMnemonicShld:
		return "Zydis Mnemonic Shld"
	case ZydisMnemonicShlx:
		return "Zydis Mnemonic Shlx"
	case ZydisMnemonicShr:
		return "Zydis Mnemonic Shr"
	case ZydisMnemonicShrd:
		return "Zydis Mnemonic Shrd"
	case ZydisMnemonicShrx:
		return "Zydis Mnemonic Shrx"
	case ZydisMnemonicShufpd:
		return "Zydis Mnemonic Shufpd"
	case ZydisMnemonicShufps:
		return "Zydis Mnemonic Shufps"
	case ZydisMnemonicSidt:
		return "Zydis Mnemonic Sidt"
	case ZydisMnemonicSkinit:
		return "Zydis Mnemonic Skinit"
	case ZydisMnemonicSldt:
		return "Zydis Mnemonic Sldt"
	case ZydisMnemonicSlwpcb:
		return "Zydis Mnemonic Slwpcb"
	case ZydisMnemonicSmsw:
		return "Zydis Mnemonic Smsw"
	case ZydisMnemonicSpflt:
		return "Zydis Mnemonic Spflt"
	case ZydisMnemonicSqrtpd:
		return "Zydis Mnemonic Sqrtpd"
	case ZydisMnemonicSqrtps:
		return "Zydis Mnemonic Sqrtps"
	case ZydisMnemonicSqrtsd:
		return "Zydis Mnemonic Sqrtsd"
	case ZydisMnemonicSqrtss:
		return "Zydis Mnemonic Sqrtss"
	case ZydisMnemonicStac:
		return "Zydis Mnemonic Stac"
	case ZydisMnemonicStc:
		return "Zydis Mnemonic Stc"
	case ZydisMnemonicStd:
		return "Zydis Mnemonic Std"
	case ZydisMnemonicStgi:
		return "Zydis Mnemonic Stgi"
	case ZydisMnemonicSti:
		return "Zydis Mnemonic Sti"
	case ZydisMnemonicStmxcsr:
		return "Zydis Mnemonic Stmxcsr"
	case ZydisMnemonicStosb:
		return "Zydis Mnemonic Stosb"
	case ZydisMnemonicStosd:
		return "Zydis Mnemonic Stosd"
	case ZydisMnemonicStosq:
		return "Zydis Mnemonic Stosq"
	case ZydisMnemonicStosw:
		return "Zydis Mnemonic Stosw"
	case ZydisMnemonicStr:
		return "Zydis Mnemonic Str"
	case ZydisMnemonicSttilecfg:
		return "Zydis Mnemonic Sttilecfg"
	case ZydisMnemonicStui:
		return "Zydis Mnemonic Stui"
	case ZydisMnemonicSub:
		return "Zydis Mnemonic Sub"
	case ZydisMnemonicSubpd:
		return "Zydis Mnemonic Subpd"
	case ZydisMnemonicSubps:
		return "Zydis Mnemonic Subps"
	case ZydisMnemonicSubsd:
		return "Zydis Mnemonic Subsd"
	case ZydisMnemonicSubss:
		return "Zydis Mnemonic Subss"
	case ZydisMnemonicSwapgs:
		return "Zydis Mnemonic Swapgs"
	case ZydisMnemonicSyscall:
		return "Zydis Mnemonic Syscall"
	case ZydisMnemonicSysenter:
		return "Zydis Mnemonic Sysenter"
	case ZydisMnemonicSysexit:
		return "Zydis Mnemonic Sysexit"
	case ZydisMnemonicSysret:
		return "Zydis Mnemonic Sysret"
	case ZydisMnemonicT1mskc:
		return "Zydis Mnemonic T1mskc"
	case ZydisMnemonicTdcall:
		return "Zydis Mnemonic Tdcall"
	case ZydisMnemonicTdpbf16ps:
		return "Zydis Mnemonic Tdpbf 16ps"
	case ZydisMnemonicTdpbssd:
		return "Zydis Mnemonic Tdpbssd"
	case ZydisMnemonicTdpbsud:
		return "Zydis Mnemonic Tdpbsud"
	case ZydisMnemonicTdpbusd:
		return "Zydis Mnemonic Tdpbusd"
	case ZydisMnemonicTdpbuud:
		return "Zydis Mnemonic Tdpbuud"
	case ZydisMnemonicTdpfp16ps:
		return "Zydis Mnemonic Tdpfp 16ps"
	case ZydisMnemonicTest:
		return "Zydis Mnemonic Test"
	case ZydisMnemonicTestui:
		return "Zydis Mnemonic Testui"
	case ZydisMnemonicTileloadd:
		return "Zydis Mnemonic Tileloadd"
	case ZydisMnemonicTileloaddt1:
		return "Zydis Mnemonic Tileloaddt 1"
	case ZydisMnemonicTilerelease:
		return "Zydis Mnemonic Tilerelease"
	case ZydisMnemonicTilestored:
		return "Zydis Mnemonic Tilestored"
	case ZydisMnemonicTilezero:
		return "Zydis Mnemonic Tilezero"
	case ZydisMnemonicTlbsync:
		return "Zydis Mnemonic Tlbsync"
	case ZydisMnemonicTpause:
		return "Zydis Mnemonic Tpause"
	case ZydisMnemonicTzcnt:
		return "Zydis Mnemonic Tzcnt"
	case ZydisMnemonicTzcnti:
		return "Zydis Mnemonic Tzcnti"
	case ZydisMnemonicTzmsk:
		return "Zydis Mnemonic Tzmsk"
	case ZydisMnemonicUcomisd:
		return "Zydis Mnemonic Ucomisd"
	case ZydisMnemonicUcomiss:
		return "Zydis Mnemonic Ucomiss"
	case ZydisMnemonicUd0:
		return "Zydis Mnemonic Ud 0"
	case ZydisMnemonicUd1:
		return "Zydis Mnemonic Ud 1"
	case ZydisMnemonicUd2:
		return "Zydis Mnemonic Ud 2"
	case ZydisMnemonicUiret:
		return "Zydis Mnemonic Uiret"
	case ZydisMnemonicUmonitor:
		return "Zydis Mnemonic Umonitor"
	case ZydisMnemonicUmwait:
		return "Zydis Mnemonic Umwait"
	case ZydisMnemonicUnpckhpd:
		return "Zydis Mnemonic Unpckhpd"
	case ZydisMnemonicUnpckhps:
		return "Zydis Mnemonic Unpckhps"
	case ZydisMnemonicUnpcklpd:
		return "Zydis Mnemonic Unpcklpd"
	case ZydisMnemonicUnpcklps:
		return "Zydis Mnemonic Unpcklps"
	case ZydisMnemonicUrdmsr:
		return "Zydis Mnemonic Urdmsr"
	case ZydisMnemonicUwrmsr:
		return "Zydis Mnemonic Uwrmsr"
	case ZydisMnemonicV4fmaddps:
		return "Zydis Mnemonic V4fmaddps"
	case ZydisMnemonicV4fmaddss:
		return "Zydis Mnemonic V4fmaddss"
	case ZydisMnemonicV4fnmaddps:
		return "Zydis Mnemonic V4fnmaddps"
	case ZydisMnemonicV4fnmaddss:
		return "Zydis Mnemonic V4fnmaddss"
	case ZydisMnemonicVaddnpd:
		return "Zydis Mnemonic Vaddnpd"
	case ZydisMnemonicVaddnps:
		return "Zydis Mnemonic Vaddnps"
	case ZydisMnemonicVaddpd:
		return "Zydis Mnemonic Vaddpd"
	case ZydisMnemonicVaddph:
		return "Zydis Mnemonic Vaddph"
	case ZydisMnemonicVaddps:
		return "Zydis Mnemonic Vaddps"
	case ZydisMnemonicVaddsd:
		return "Zydis Mnemonic Vaddsd"
	case ZydisMnemonicVaddsetsps:
		return "Zydis Mnemonic Vaddsetsps"
	case ZydisMnemonicVaddsh:
		return "Zydis Mnemonic Vaddsh"
	case ZydisMnemonicVaddss:
		return "Zydis Mnemonic Vaddss"
	case ZydisMnemonicVaddsubpd:
		return "Zydis Mnemonic Vaddsubpd"
	case ZydisMnemonicVaddsubps:
		return "Zydis Mnemonic Vaddsubps"
	case ZydisMnemonicVaesdec:
		return "Zydis Mnemonic Vaesdec"
	case ZydisMnemonicVaesdeclast:
		return "Zydis Mnemonic Vaesdeclast"
	case ZydisMnemonicVaesenc:
		return "Zydis Mnemonic Vaesenc"
	case ZydisMnemonicVaesenclast:
		return "Zydis Mnemonic Vaesenclast"
	case ZydisMnemonicVaesimc:
		return "Zydis Mnemonic Vaesimc"
	case ZydisMnemonicVaeskeygenassist:
		return "Zydis Mnemonic Vaeskeygenassist"
	case ZydisMnemonicValignd:
		return "Zydis Mnemonic Valignd"
	case ZydisMnemonicValignq:
		return "Zydis Mnemonic Valignq"
	case ZydisMnemonicVandnpd:
		return "Zydis Mnemonic Vandnpd"
	case ZydisMnemonicVandnps:
		return "Zydis Mnemonic Vandnps"
	case ZydisMnemonicVandpd:
		return "Zydis Mnemonic Vandpd"
	case ZydisMnemonicVandps:
		return "Zydis Mnemonic Vandps"
	case ZydisMnemonicVbcstnebf162ps:
		return "Zydis Mnemonic Vbcstnebf 162ps"
	case ZydisMnemonicVbcstnesh2ps:
		return "Zydis Mnemonic Vbcstnesh 2ps"
	case ZydisMnemonicVblendmpd:
		return "Zydis Mnemonic Vblendmpd"
	case ZydisMnemonicVblendmps:
		return "Zydis Mnemonic Vblendmps"
	case ZydisMnemonicVblendpd:
		return "Zydis Mnemonic Vblendpd"
	case ZydisMnemonicVblendps:
		return "Zydis Mnemonic Vblendps"
	case ZydisMnemonicVblendvpd:
		return "Zydis Mnemonic Vblendvpd"
	case ZydisMnemonicVblendvps:
		return "Zydis Mnemonic Vblendvps"
	case ZydisMnemonicVbroadcastf128:
		return "Zydis Mnemonic Vbroadcastf 128"
	case ZydisMnemonicVbroadcastf32x2:
		return "Zydis Mnemonic Vbroadcastf 32x 2"
	case ZydisMnemonicVbroadcastf32x4:
		return "Zydis Mnemonic Vbroadcastf 32x 4"
	case ZydisMnemonicVbroadcastf32x8:
		return "Zydis Mnemonic Vbroadcastf 32x 8"
	case ZydisMnemonicVbroadcastf64x2:
		return "Zydis Mnemonic Vbroadcastf 64x 2"
	case ZydisMnemonicVbroadcastf64x4:
		return "Zydis Mnemonic Vbroadcastf 64x 4"
	case ZydisMnemonicVbroadcasti128:
		return "Zydis Mnemonic Vbroadcasti 128"
	case ZydisMnemonicVbroadcasti32x2:
		return "Zydis Mnemonic Vbroadcasti 32x 2"
	case ZydisMnemonicVbroadcasti32x4:
		return "Zydis Mnemonic Vbroadcasti 32x 4"
	case ZydisMnemonicVbroadcasti32x8:
		return "Zydis Mnemonic Vbroadcasti 32x 8"
	case ZydisMnemonicVbroadcasti64x2:
		return "Zydis Mnemonic Vbroadcasti 64x 2"
	case ZydisMnemonicVbroadcasti64x4:
		return "Zydis Mnemonic Vbroadcasti 64x 4"
	case ZydisMnemonicVbroadcastsd:
		return "Zydis Mnemonic Vbroadcastsd"
	case ZydisMnemonicVbroadcastss:
		return "Zydis Mnemonic Vbroadcastss"
	case ZydisMnemonicVcmppd:
		return "Zydis Mnemonic Vcmppd"
	case ZydisMnemonicVcmpph:
		return "Zydis Mnemonic Vcmpph"
	case ZydisMnemonicVcmpps:
		return "Zydis Mnemonic Vcmpps"
	case ZydisMnemonicVcmpsd:
		return "Zydis Mnemonic Vcmpsd"
	case ZydisMnemonicVcmpsh:
		return "Zydis Mnemonic Vcmpsh"
	case ZydisMnemonicVcmpss:
		return "Zydis Mnemonic Vcmpss"
	case ZydisMnemonicVcomisd:
		return "Zydis Mnemonic Vcomisd"
	case ZydisMnemonicVcomish:
		return "Zydis Mnemonic Vcomish"
	case ZydisMnemonicVcomiss:
		return "Zydis Mnemonic Vcomiss"
	case ZydisMnemonicVcompresspd:
		return "Zydis Mnemonic Vcompresspd"
	case ZydisMnemonicVcompressps:
		return "Zydis Mnemonic Vcompressps"
	case ZydisMnemonicVcvtdq2pd:
		return "Zydis Mnemonic Vcvtdq 2pd"
	case ZydisMnemonicVcvtdq2ph:
		return "Zydis Mnemonic Vcvtdq 2ph"
	case ZydisMnemonicVcvtdq2ps:
		return "Zydis Mnemonic Vcvtdq 2ps"
	case ZydisMnemonicVcvtfxpntdq2ps:
		return "Zydis Mnemonic Vcvtfxpntdq 2ps"
	case ZydisMnemonicVcvtfxpntpd2dq:
		return "Zydis Mnemonic Vcvtfxpntpd 2dq"
	case ZydisMnemonicVcvtfxpntpd2udq:
		return "Zydis Mnemonic Vcvtfxpntpd 2udq"
	case ZydisMnemonicVcvtfxpntps2dq:
		return "Zydis Mnemonic Vcvtfxpntps 2dq"
	case ZydisMnemonicVcvtfxpntps2udq:
		return "Zydis Mnemonic Vcvtfxpntps 2udq"
	case ZydisMnemonicVcvtfxpntudq2ps:
		return "Zydis Mnemonic Vcvtfxpntudq 2ps"
	case ZydisMnemonicVcvtne2ps2bf16:
		return "Zydis Mnemonic Vcvtne 2ps 2bf 16"
	case ZydisMnemonicVcvtneebf162ps:
		return "Zydis Mnemonic Vcvtneebf 162ps"
	case ZydisMnemonicVcvtneeph2ps:
		return "Zydis Mnemonic Vcvtneeph 2ps"
	case ZydisMnemonicVcvtneobf162ps:
		return "Zydis Mnemonic Vcvtneobf 162ps"
	case ZydisMnemonicVcvtneoph2ps:
		return "Zydis Mnemonic Vcvtneoph 2ps"
	case ZydisMnemonicVcvtneps2bf16:
		return "Zydis Mnemonic Vcvtneps 2bf 16"
	case ZydisMnemonicVcvtpd2dq:
		return "Zydis Mnemonic Vcvtpd 2dq"
	case ZydisMnemonicVcvtpd2ph:
		return "Zydis Mnemonic Vcvtpd 2ph"
	case ZydisMnemonicVcvtpd2ps:
		return "Zydis Mnemonic Vcvtpd 2ps"
	case ZydisMnemonicVcvtpd2qq:
		return "Zydis Mnemonic Vcvtpd 2qq"
	case ZydisMnemonicVcvtpd2udq:
		return "Zydis Mnemonic Vcvtpd 2udq"
	case ZydisMnemonicVcvtpd2uqq:
		return "Zydis Mnemonic Vcvtpd 2uqq"
	case ZydisMnemonicVcvtph2dq:
		return "Zydis Mnemonic Vcvtph 2dq"
	case ZydisMnemonicVcvtph2pd:
		return "Zydis Mnemonic Vcvtph 2pd"
	case ZydisMnemonicVcvtph2ps:
		return "Zydis Mnemonic Vcvtph 2ps"
	case ZydisMnemonicVcvtph2psx:
		return "Zydis Mnemonic Vcvtph 2psx"
	case ZydisMnemonicVcvtph2qq:
		return "Zydis Mnemonic Vcvtph 2qq"
	case ZydisMnemonicVcvtph2udq:
		return "Zydis Mnemonic Vcvtph 2udq"
	case ZydisMnemonicVcvtph2uqq:
		return "Zydis Mnemonic Vcvtph 2uqq"
	case ZydisMnemonicVcvtph2uw:
		return "Zydis Mnemonic Vcvtph 2uw"
	case ZydisMnemonicVcvtph2w:
		return "Zydis Mnemonic Vcvtph 2w"
	case ZydisMnemonicVcvtps2dq:
		return "Zydis Mnemonic Vcvtps 2dq"
	case ZydisMnemonicVcvtps2pd:
		return "Zydis Mnemonic Vcvtps 2pd"
	case ZydisMnemonicVcvtps2ph:
		return "Zydis Mnemonic Vcvtps 2ph"
	case ZydisMnemonicVcvtps2phx:
		return "Zydis Mnemonic Vcvtps 2phx"
	case ZydisMnemonicVcvtps2qq:
		return "Zydis Mnemonic Vcvtps 2qq"
	case ZydisMnemonicVcvtps2udq:
		return "Zydis Mnemonic Vcvtps 2udq"
	case ZydisMnemonicVcvtps2uqq:
		return "Zydis Mnemonic Vcvtps 2uqq"
	case ZydisMnemonicVcvtqq2pd:
		return "Zydis Mnemonic Vcvtqq 2pd"
	case ZydisMnemonicVcvtqq2ph:
		return "Zydis Mnemonic Vcvtqq 2ph"
	case ZydisMnemonicVcvtqq2ps:
		return "Zydis Mnemonic Vcvtqq 2ps"
	case ZydisMnemonicVcvtsd2sh:
		return "Zydis Mnemonic Vcvtsd 2sh"
	case ZydisMnemonicVcvtsd2si:
		return "Zydis Mnemonic Vcvtsd 2si"
	case ZydisMnemonicVcvtsd2ss:
		return "Zydis Mnemonic Vcvtsd 2ss"
	case ZydisMnemonicVcvtsd2usi:
		return "Zydis Mnemonic Vcvtsd 2usi"
	case ZydisMnemonicVcvtsh2sd:
		return "Zydis Mnemonic Vcvtsh 2sd"
	case ZydisMnemonicVcvtsh2si:
		return "Zydis Mnemonic Vcvtsh 2si"
	case ZydisMnemonicVcvtsh2ss:
		return "Zydis Mnemonic Vcvtsh 2ss"
	case ZydisMnemonicVcvtsh2usi:
		return "Zydis Mnemonic Vcvtsh 2usi"
	case ZydisMnemonicVcvtsi2sd:
		return "Zydis Mnemonic Vcvtsi 2sd"
	case ZydisMnemonicVcvtsi2sh:
		return "Zydis Mnemonic Vcvtsi 2sh"
	case ZydisMnemonicVcvtsi2ss:
		return "Zydis Mnemonic Vcvtsi 2ss"
	case ZydisMnemonicVcvtss2sd:
		return "Zydis Mnemonic Vcvtss 2sd"
	case ZydisMnemonicVcvtss2sh:
		return "Zydis Mnemonic Vcvtss 2sh"
	case ZydisMnemonicVcvtss2si:
		return "Zydis Mnemonic Vcvtss 2si"
	case ZydisMnemonicVcvtss2usi:
		return "Zydis Mnemonic Vcvtss 2usi"
	case ZydisMnemonicVcvttpd2dq:
		return "Zydis Mnemonic Vcvttpd 2dq"
	case ZydisMnemonicVcvttpd2qq:
		return "Zydis Mnemonic Vcvttpd 2qq"
	case ZydisMnemonicVcvttpd2udq:
		return "Zydis Mnemonic Vcvttpd 2udq"
	case ZydisMnemonicVcvttpd2uqq:
		return "Zydis Mnemonic Vcvttpd 2uqq"
	case ZydisMnemonicVcvttph2dq:
		return "Zydis Mnemonic Vcvttph 2dq"
	case ZydisMnemonicVcvttph2qq:
		return "Zydis Mnemonic Vcvttph 2qq"
	case ZydisMnemonicVcvttph2udq:
		return "Zydis Mnemonic Vcvttph 2udq"
	case ZydisMnemonicVcvttph2uqq:
		return "Zydis Mnemonic Vcvttph 2uqq"
	case ZydisMnemonicVcvttph2uw:
		return "Zydis Mnemonic Vcvttph 2uw"
	case ZydisMnemonicVcvttph2w:
		return "Zydis Mnemonic Vcvttph 2w"
	case ZydisMnemonicVcvttps2dq:
		return "Zydis Mnemonic Vcvttps 2dq"
	case ZydisMnemonicVcvttps2qq:
		return "Zydis Mnemonic Vcvttps 2qq"
	case ZydisMnemonicVcvttps2udq:
		return "Zydis Mnemonic Vcvttps 2udq"
	case ZydisMnemonicVcvttps2uqq:
		return "Zydis Mnemonic Vcvttps 2uqq"
	case ZydisMnemonicVcvttsd2si:
		return "Zydis Mnemonic Vcvttsd 2si"
	case ZydisMnemonicVcvttsd2usi:
		return "Zydis Mnemonic Vcvttsd 2usi"
	case ZydisMnemonicVcvttsh2si:
		return "Zydis Mnemonic Vcvttsh 2si"
	case ZydisMnemonicVcvttsh2usi:
		return "Zydis Mnemonic Vcvttsh 2usi"
	case ZydisMnemonicVcvttss2si:
		return "Zydis Mnemonic Vcvttss 2si"
	case ZydisMnemonicVcvttss2usi:
		return "Zydis Mnemonic Vcvttss 2usi"
	case ZydisMnemonicVcvtudq2pd:
		return "Zydis Mnemonic Vcvtudq 2pd"
	case ZydisMnemonicVcvtudq2ph:
		return "Zydis Mnemonic Vcvtudq 2ph"
	case ZydisMnemonicVcvtudq2ps:
		return "Zydis Mnemonic Vcvtudq 2ps"
	case ZydisMnemonicVcvtuqq2pd:
		return "Zydis Mnemonic Vcvtuqq 2pd"
	case ZydisMnemonicVcvtuqq2ph:
		return "Zydis Mnemonic Vcvtuqq 2ph"
	case ZydisMnemonicVcvtuqq2ps:
		return "Zydis Mnemonic Vcvtuqq 2ps"
	case ZydisMnemonicVcvtusi2sd:
		return "Zydis Mnemonic Vcvtusi 2sd"
	case ZydisMnemonicVcvtusi2sh:
		return "Zydis Mnemonic Vcvtusi 2sh"
	case ZydisMnemonicVcvtusi2ss:
		return "Zydis Mnemonic Vcvtusi 2ss"
	case ZydisMnemonicVcvtuw2ph:
		return "Zydis Mnemonic Vcvtuw 2ph"
	case ZydisMnemonicVcvtw2ph:
		return "Zydis Mnemonic Vcvtw 2ph"
	case ZydisMnemonicVdbpsadbw:
		return "Zydis Mnemonic Vdbpsadbw"
	case ZydisMnemonicVdivpd:
		return "Zydis Mnemonic Vdivpd"
	case ZydisMnemonicVdivph:
		return "Zydis Mnemonic Vdivph"
	case ZydisMnemonicVdivps:
		return "Zydis Mnemonic Vdivps"
	case ZydisMnemonicVdivsd:
		return "Zydis Mnemonic Vdivsd"
	case ZydisMnemonicVdivsh:
		return "Zydis Mnemonic Vdivsh"
	case ZydisMnemonicVdivss:
		return "Zydis Mnemonic Vdivss"
	case ZydisMnemonicVdpbf16ps:
		return "Zydis Mnemonic Vdpbf 16ps"
	case ZydisMnemonicVdppd:
		return "Zydis Mnemonic Vdppd"
	case ZydisMnemonicVdpps:
		return "Zydis Mnemonic Vdpps"
	case ZydisMnemonicVerr:
		return "Zydis Mnemonic Verr"
	case ZydisMnemonicVerw:
		return "Zydis Mnemonic Verw"
	case ZydisMnemonicVexp223ps:
		return "Zydis Mnemonic Vexp 223ps"
	case ZydisMnemonicVexp2pd:
		return "Zydis Mnemonic Vexp 2pd"
	case ZydisMnemonicVexp2ps:
		return "Zydis Mnemonic Vexp 2ps"
	case ZydisMnemonicVexpandpd:
		return "Zydis Mnemonic Vexpandpd"
	case ZydisMnemonicVexpandps:
		return "Zydis Mnemonic Vexpandps"
	case ZydisMnemonicVextractf128:
		return "Zydis Mnemonic Vextractf 128"
	case ZydisMnemonicVextractf32x4:
		return "Zydis Mnemonic Vextractf 32x 4"
	case ZydisMnemonicVextractf32x8:
		return "Zydis Mnemonic Vextractf 32x 8"
	case ZydisMnemonicVextractf64x2:
		return "Zydis Mnemonic Vextractf 64x 2"
	case ZydisMnemonicVextractf64x4:
		return "Zydis Mnemonic Vextractf 64x 4"
	case ZydisMnemonicVextracti128:
		return "Zydis Mnemonic Vextracti 128"
	case ZydisMnemonicVextracti32x4:
		return "Zydis Mnemonic Vextracti 32x 4"
	case ZydisMnemonicVextracti32x8:
		return "Zydis Mnemonic Vextracti 32x 8"
	case ZydisMnemonicVextracti64x2:
		return "Zydis Mnemonic Vextracti 64x 2"
	case ZydisMnemonicVextracti64x4:
		return "Zydis Mnemonic Vextracti 64x 4"
	case ZydisMnemonicVextractps:
		return "Zydis Mnemonic Vextractps"
	case ZydisMnemonicVfcmaddcph:
		return "Zydis Mnemonic Vfcmaddcph"
	case ZydisMnemonicVfcmaddcsh:
		return "Zydis Mnemonic Vfcmaddcsh"
	case ZydisMnemonicVfcmulcph:
		return "Zydis Mnemonic Vfcmulcph"
	case ZydisMnemonicVfcmulcsh:
		return "Zydis Mnemonic Vfcmulcsh"
	case ZydisMnemonicVfixupimmpd:
		return "Zydis Mnemonic Vfixupimmpd"
	case ZydisMnemonicVfixupimmps:
		return "Zydis Mnemonic Vfixupimmps"
	case ZydisMnemonicVfixupimmsd:
		return "Zydis Mnemonic Vfixupimmsd"
	case ZydisMnemonicVfixupimmss:
		return "Zydis Mnemonic Vfixupimmss"
	case ZydisMnemonicVfixupnanpd:
		return "Zydis Mnemonic Vfixupnanpd"
	case ZydisMnemonicVfixupnanps:
		return "Zydis Mnemonic Vfixupnanps"
	case ZydisMnemonicVfmadd132pd:
		return "Zydis Mnemonic Vfmadd 132pd"
	case ZydisMnemonicVfmadd132ph:
		return "Zydis Mnemonic Vfmadd 132ph"
	case ZydisMnemonicVfmadd132ps:
		return "Zydis Mnemonic Vfmadd 132ps"
	case ZydisMnemonicVfmadd132sd:
		return "Zydis Mnemonic Vfmadd 132sd"
	case ZydisMnemonicVfmadd132sh:
		return "Zydis Mnemonic Vfmadd 132sh"
	case ZydisMnemonicVfmadd132ss:
		return "Zydis Mnemonic Vfmadd 132ss"
	case ZydisMnemonicVfmadd213pd:
		return "Zydis Mnemonic Vfmadd 213pd"
	case ZydisMnemonicVfmadd213ph:
		return "Zydis Mnemonic Vfmadd 213ph"
	case ZydisMnemonicVfmadd213ps:
		return "Zydis Mnemonic Vfmadd 213ps"
	case ZydisMnemonicVfmadd213sd:
		return "Zydis Mnemonic Vfmadd 213sd"
	case ZydisMnemonicVfmadd213sh:
		return "Zydis Mnemonic Vfmadd 213sh"
	case ZydisMnemonicVfmadd213ss:
		return "Zydis Mnemonic Vfmadd 213ss"
	case ZydisMnemonicVfmadd231pd:
		return "Zydis Mnemonic Vfmadd 231pd"
	case ZydisMnemonicVfmadd231ph:
		return "Zydis Mnemonic Vfmadd 231ph"
	case ZydisMnemonicVfmadd231ps:
		return "Zydis Mnemonic Vfmadd 231ps"
	case ZydisMnemonicVfmadd231sd:
		return "Zydis Mnemonic Vfmadd 231sd"
	case ZydisMnemonicVfmadd231sh:
		return "Zydis Mnemonic Vfmadd 231sh"
	case ZydisMnemonicVfmadd231ss:
		return "Zydis Mnemonic Vfmadd 231ss"
	case ZydisMnemonicVfmadd233ps:
		return "Zydis Mnemonic Vfmadd 233ps"
	case ZydisMnemonicVfmaddcph:
		return "Zydis Mnemonic Vfmaddcph"
	case ZydisMnemonicVfmaddcsh:
		return "Zydis Mnemonic Vfmaddcsh"
	case ZydisMnemonicVfmaddpd:
		return "Zydis Mnemonic Vfmaddpd"
	case ZydisMnemonicVfmaddps:
		return "Zydis Mnemonic Vfmaddps"
	case ZydisMnemonicVfmaddsd:
		return "Zydis Mnemonic Vfmaddsd"
	case ZydisMnemonicVfmaddss:
		return "Zydis Mnemonic Vfmaddss"
	case ZydisMnemonicVfmaddsub132pd:
		return "Zydis Mnemonic Vfmaddsub 132pd"
	case ZydisMnemonicVfmaddsub132ph:
		return "Zydis Mnemonic Vfmaddsub 132ph"
	case ZydisMnemonicVfmaddsub132ps:
		return "Zydis Mnemonic Vfmaddsub 132ps"
	case ZydisMnemonicVfmaddsub213pd:
		return "Zydis Mnemonic Vfmaddsub 213pd"
	case ZydisMnemonicVfmaddsub213ph:
		return "Zydis Mnemonic Vfmaddsub 213ph"
	case ZydisMnemonicVfmaddsub213ps:
		return "Zydis Mnemonic Vfmaddsub 213ps"
	case ZydisMnemonicVfmaddsub231pd:
		return "Zydis Mnemonic Vfmaddsub 231pd"
	case ZydisMnemonicVfmaddsub231ph:
		return "Zydis Mnemonic Vfmaddsub 231ph"
	case ZydisMnemonicVfmaddsub231ps:
		return "Zydis Mnemonic Vfmaddsub 231ps"
	case ZydisMnemonicVfmaddsubpd:
		return "Zydis Mnemonic Vfmaddsubpd"
	case ZydisMnemonicVfmaddsubps:
		return "Zydis Mnemonic Vfmaddsubps"
	case ZydisMnemonicVfmsub132pd:
		return "Zydis Mnemonic Vfmsub 132pd"
	case ZydisMnemonicVfmsub132ph:
		return "Zydis Mnemonic Vfmsub 132ph"
	case ZydisMnemonicVfmsub132ps:
		return "Zydis Mnemonic Vfmsub 132ps"
	case ZydisMnemonicVfmsub132sd:
		return "Zydis Mnemonic Vfmsub 132sd"
	case ZydisMnemonicVfmsub132sh:
		return "Zydis Mnemonic Vfmsub 132sh"
	case ZydisMnemonicVfmsub132ss:
		return "Zydis Mnemonic Vfmsub 132ss"
	case ZydisMnemonicVfmsub213pd:
		return "Zydis Mnemonic Vfmsub 213pd"
	case ZydisMnemonicVfmsub213ph:
		return "Zydis Mnemonic Vfmsub 213ph"
	case ZydisMnemonicVfmsub213ps:
		return "Zydis Mnemonic Vfmsub 213ps"
	case ZydisMnemonicVfmsub213sd:
		return "Zydis Mnemonic Vfmsub 213sd"
	case ZydisMnemonicVfmsub213sh:
		return "Zydis Mnemonic Vfmsub 213sh"
	case ZydisMnemonicVfmsub213ss:
		return "Zydis Mnemonic Vfmsub 213ss"
	case ZydisMnemonicVfmsub231pd:
		return "Zydis Mnemonic Vfmsub 231pd"
	case ZydisMnemonicVfmsub231ph:
		return "Zydis Mnemonic Vfmsub 231ph"
	case ZydisMnemonicVfmsub231ps:
		return "Zydis Mnemonic Vfmsub 231ps"
	case ZydisMnemonicVfmsub231sd:
		return "Zydis Mnemonic Vfmsub 231sd"
	case ZydisMnemonicVfmsub231sh:
		return "Zydis Mnemonic Vfmsub 231sh"
	case ZydisMnemonicVfmsub231ss:
		return "Zydis Mnemonic Vfmsub 231ss"
	case ZydisMnemonicVfmsubadd132pd:
		return "Zydis Mnemonic Vfmsubadd 132pd"
	case ZydisMnemonicVfmsubadd132ph:
		return "Zydis Mnemonic Vfmsubadd 132ph"
	case ZydisMnemonicVfmsubadd132ps:
		return "Zydis Mnemonic Vfmsubadd 132ps"
	case ZydisMnemonicVfmsubadd213pd:
		return "Zydis Mnemonic Vfmsubadd 213pd"
	case ZydisMnemonicVfmsubadd213ph:
		return "Zydis Mnemonic Vfmsubadd 213ph"
	case ZydisMnemonicVfmsubadd213ps:
		return "Zydis Mnemonic Vfmsubadd 213ps"
	case ZydisMnemonicVfmsubadd231pd:
		return "Zydis Mnemonic Vfmsubadd 231pd"
	case ZydisMnemonicVfmsubadd231ph:
		return "Zydis Mnemonic Vfmsubadd 231ph"
	case ZydisMnemonicVfmsubadd231ps:
		return "Zydis Mnemonic Vfmsubadd 231ps"
	case ZydisMnemonicVfmsubaddpd:
		return "Zydis Mnemonic Vfmsubaddpd"
	case ZydisMnemonicVfmsubaddps:
		return "Zydis Mnemonic Vfmsubaddps"
	case ZydisMnemonicVfmsubpd:
		return "Zydis Mnemonic Vfmsubpd"
	case ZydisMnemonicVfmsubps:
		return "Zydis Mnemonic Vfmsubps"
	case ZydisMnemonicVfmsubsd:
		return "Zydis Mnemonic Vfmsubsd"
	case ZydisMnemonicVfmsubss:
		return "Zydis Mnemonic Vfmsubss"
	case ZydisMnemonicVfmulcph:
		return "Zydis Mnemonic Vfmulcph"
	case ZydisMnemonicVfmulcsh:
		return "Zydis Mnemonic Vfmulcsh"
	case ZydisMnemonicVfnmadd132pd:
		return "Zydis Mnemonic Vfnmadd 132pd"
	case ZydisMnemonicVfnmadd132ph:
		return "Zydis Mnemonic Vfnmadd 132ph"
	case ZydisMnemonicVfnmadd132ps:
		return "Zydis Mnemonic Vfnmadd 132ps"
	case ZydisMnemonicVfnmadd132sd:
		return "Zydis Mnemonic Vfnmadd 132sd"
	case ZydisMnemonicVfnmadd132sh:
		return "Zydis Mnemonic Vfnmadd 132sh"
	case ZydisMnemonicVfnmadd132ss:
		return "Zydis Mnemonic Vfnmadd 132ss"
	case ZydisMnemonicVfnmadd213pd:
		return "Zydis Mnemonic Vfnmadd 213pd"
	case ZydisMnemonicVfnmadd213ph:
		return "Zydis Mnemonic Vfnmadd 213ph"
	case ZydisMnemonicVfnmadd213ps:
		return "Zydis Mnemonic Vfnmadd 213ps"
	case ZydisMnemonicVfnmadd213sd:
		return "Zydis Mnemonic Vfnmadd 213sd"
	case ZydisMnemonicVfnmadd213sh:
		return "Zydis Mnemonic Vfnmadd 213sh"
	case ZydisMnemonicVfnmadd213ss:
		return "Zydis Mnemonic Vfnmadd 213ss"
	case ZydisMnemonicVfnmadd231pd:
		return "Zydis Mnemonic Vfnmadd 231pd"
	case ZydisMnemonicVfnmadd231ph:
		return "Zydis Mnemonic Vfnmadd 231ph"
	case ZydisMnemonicVfnmadd231ps:
		return "Zydis Mnemonic Vfnmadd 231ps"
	case ZydisMnemonicVfnmadd231sd:
		return "Zydis Mnemonic Vfnmadd 231sd"
	case ZydisMnemonicVfnmadd231sh:
		return "Zydis Mnemonic Vfnmadd 231sh"
	case ZydisMnemonicVfnmadd231ss:
		return "Zydis Mnemonic Vfnmadd 231ss"
	case ZydisMnemonicVfnmaddpd:
		return "Zydis Mnemonic Vfnmaddpd"
	case ZydisMnemonicVfnmaddps:
		return "Zydis Mnemonic Vfnmaddps"
	case ZydisMnemonicVfnmaddsd:
		return "Zydis Mnemonic Vfnmaddsd"
	case ZydisMnemonicVfnmaddss:
		return "Zydis Mnemonic Vfnmaddss"
	case ZydisMnemonicVfnmsub132pd:
		return "Zydis Mnemonic Vfnmsub 132pd"
	case ZydisMnemonicVfnmsub132ph:
		return "Zydis Mnemonic Vfnmsub 132ph"
	case ZydisMnemonicVfnmsub132ps:
		return "Zydis Mnemonic Vfnmsub 132ps"
	case ZydisMnemonicVfnmsub132sd:
		return "Zydis Mnemonic Vfnmsub 132sd"
	case ZydisMnemonicVfnmsub132sh:
		return "Zydis Mnemonic Vfnmsub 132sh"
	case ZydisMnemonicVfnmsub132ss:
		return "Zydis Mnemonic Vfnmsub 132ss"
	case ZydisMnemonicVfnmsub213pd:
		return "Zydis Mnemonic Vfnmsub 213pd"
	case ZydisMnemonicVfnmsub213ph:
		return "Zydis Mnemonic Vfnmsub 213ph"
	case ZydisMnemonicVfnmsub213ps:
		return "Zydis Mnemonic Vfnmsub 213ps"
	case ZydisMnemonicVfnmsub213sd:
		return "Zydis Mnemonic Vfnmsub 213sd"
	case ZydisMnemonicVfnmsub213sh:
		return "Zydis Mnemonic Vfnmsub 213sh"
	case ZydisMnemonicVfnmsub213ss:
		return "Zydis Mnemonic Vfnmsub 213ss"
	case ZydisMnemonicVfnmsub231pd:
		return "Zydis Mnemonic Vfnmsub 231pd"
	case ZydisMnemonicVfnmsub231ph:
		return "Zydis Mnemonic Vfnmsub 231ph"
	case ZydisMnemonicVfnmsub231ps:
		return "Zydis Mnemonic Vfnmsub 231ps"
	case ZydisMnemonicVfnmsub231sd:
		return "Zydis Mnemonic Vfnmsub 231sd"
	case ZydisMnemonicVfnmsub231sh:
		return "Zydis Mnemonic Vfnmsub 231sh"
	case ZydisMnemonicVfnmsub231ss:
		return "Zydis Mnemonic Vfnmsub 231ss"
	case ZydisMnemonicVfnmsubpd:
		return "Zydis Mnemonic Vfnmsubpd"
	case ZydisMnemonicVfnmsubps:
		return "Zydis Mnemonic Vfnmsubps"
	case ZydisMnemonicVfnmsubsd:
		return "Zydis Mnemonic Vfnmsubsd"
	case ZydisMnemonicVfnmsubss:
		return "Zydis Mnemonic Vfnmsubss"
	case ZydisMnemonicVfpclasspd:
		return "Zydis Mnemonic Vfpclasspd"
	case ZydisMnemonicVfpclassph:
		return "Zydis Mnemonic Vfpclassph"
	case ZydisMnemonicVfpclassps:
		return "Zydis Mnemonic Vfpclassps"
	case ZydisMnemonicVfpclasssd:
		return "Zydis Mnemonic Vfpclasssd"
	case ZydisMnemonicVfpclasssh:
		return "Zydis Mnemonic Vfpclasssh"
	case ZydisMnemonicVfpclassss:
		return "Zydis Mnemonic Vfpclassss"
	case ZydisMnemonicVfrczpd:
		return "Zydis Mnemonic Vfrczpd"
	case ZydisMnemonicVfrczps:
		return "Zydis Mnemonic Vfrczps"
	case ZydisMnemonicVfrczsd:
		return "Zydis Mnemonic Vfrczsd"
	case ZydisMnemonicVfrczss:
		return "Zydis Mnemonic Vfrczss"
	case ZydisMnemonicVgatherdpd:
		return "Zydis Mnemonic Vgatherdpd"
	case ZydisMnemonicVgatherdps:
		return "Zydis Mnemonic Vgatherdps"
	case ZydisMnemonicVgatherpf0dpd:
		return "Zydis Mnemonic Vgatherpf 0dpd"
	case ZydisMnemonicVgatherpf0dps:
		return "Zydis Mnemonic Vgatherpf 0dps"
	case ZydisMnemonicVgatherpf0hintdpd:
		return "Zydis Mnemonic Vgatherpf 0hintdpd"
	case ZydisMnemonicVgatherpf0hintdps:
		return "Zydis Mnemonic Vgatherpf 0hintdps"
	case ZydisMnemonicVgatherpf0qpd:
		return "Zydis Mnemonic Vgatherpf 0qpd"
	case ZydisMnemonicVgatherpf0qps:
		return "Zydis Mnemonic Vgatherpf 0qps"
	case ZydisMnemonicVgatherpf1dpd:
		return "Zydis Mnemonic Vgatherpf 1dpd"
	case ZydisMnemonicVgatherpf1dps:
		return "Zydis Mnemonic Vgatherpf 1dps"
	case ZydisMnemonicVgatherpf1qpd:
		return "Zydis Mnemonic Vgatherpf 1qpd"
	case ZydisMnemonicVgatherpf1qps:
		return "Zydis Mnemonic Vgatherpf 1qps"
	case ZydisMnemonicVgatherqpd:
		return "Zydis Mnemonic Vgatherqpd"
	case ZydisMnemonicVgatherqps:
		return "Zydis Mnemonic Vgatherqps"
	case ZydisMnemonicVgetexppd:
		return "Zydis Mnemonic Vgetexppd"
	case ZydisMnemonicVgetexpph:
		return "Zydis Mnemonic Vgetexpph"
	case ZydisMnemonicVgetexpps:
		return "Zydis Mnemonic Vgetexpps"
	case ZydisMnemonicVgetexpsd:
		return "Zydis Mnemonic Vgetexpsd"
	case ZydisMnemonicVgetexpsh:
		return "Zydis Mnemonic Vgetexpsh"
	case ZydisMnemonicVgetexpss:
		return "Zydis Mnemonic Vgetexpss"
	case ZydisMnemonicVgetmantpd:
		return "Zydis Mnemonic Vgetmantpd"
	case ZydisMnemonicVgetmantph:
		return "Zydis Mnemonic Vgetmantph"
	case ZydisMnemonicVgetmantps:
		return "Zydis Mnemonic Vgetmantps"
	case ZydisMnemonicVgetmantsd:
		return "Zydis Mnemonic Vgetmantsd"
	case ZydisMnemonicVgetmantsh:
		return "Zydis Mnemonic Vgetmantsh"
	case ZydisMnemonicVgetmantss:
		return "Zydis Mnemonic Vgetmantss"
	case ZydisMnemonicVgf2p8affineinvqb:
		return "Zydis Mnemonic Vgf 2p 8affineinvqb"
	case ZydisMnemonicVgf2p8affineqb:
		return "Zydis Mnemonic Vgf 2p 8affineqb"
	case ZydisMnemonicVgf2p8mulb:
		return "Zydis Mnemonic Vgf 2p 8mulb"
	case ZydisMnemonicVgmaxabsps:
		return "Zydis Mnemonic Vgmaxabsps"
	case ZydisMnemonicVgmaxpd:
		return "Zydis Mnemonic Vgmaxpd"
	case ZydisMnemonicVgmaxps:
		return "Zydis Mnemonic Vgmaxps"
	case ZydisMnemonicVgminpd:
		return "Zydis Mnemonic Vgminpd"
	case ZydisMnemonicVgminps:
		return "Zydis Mnemonic Vgminps"
	case ZydisMnemonicVhaddpd:
		return "Zydis Mnemonic Vhaddpd"
	case ZydisMnemonicVhaddps:
		return "Zydis Mnemonic Vhaddps"
	case ZydisMnemonicVhsubpd:
		return "Zydis Mnemonic Vhsubpd"
	case ZydisMnemonicVhsubps:
		return "Zydis Mnemonic Vhsubps"
	case ZydisMnemonicVinsertf128:
		return "Zydis Mnemonic Vinsertf 128"
	case ZydisMnemonicVinsertf32x4:
		return "Zydis Mnemonic Vinsertf 32x 4"
	case ZydisMnemonicVinsertf32x8:
		return "Zydis Mnemonic Vinsertf 32x 8"
	case ZydisMnemonicVinsertf64x2:
		return "Zydis Mnemonic Vinsertf 64x 2"
	case ZydisMnemonicVinsertf64x4:
		return "Zydis Mnemonic Vinsertf 64x 4"
	case ZydisMnemonicVinserti128:
		return "Zydis Mnemonic Vinserti 128"
	case ZydisMnemonicVinserti32x4:
		return "Zydis Mnemonic Vinserti 32x 4"
	case ZydisMnemonicVinserti32x8:
		return "Zydis Mnemonic Vinserti 32x 8"
	case ZydisMnemonicVinserti64x2:
		return "Zydis Mnemonic Vinserti 64x 2"
	case ZydisMnemonicVinserti64x4:
		return "Zydis Mnemonic Vinserti 64x 4"
	case ZydisMnemonicVinsertps:
		return "Zydis Mnemonic Vinsertps"
	case ZydisMnemonicVlddqu:
		return "Zydis Mnemonic Vlddqu"
	case ZydisMnemonicVldmxcsr:
		return "Zydis Mnemonic Vldmxcsr"
	case ZydisMnemonicVloadunpackhd:
		return "Zydis Mnemonic Vloadunpackhd"
	case ZydisMnemonicVloadunpackhpd:
		return "Zydis Mnemonic Vloadunpackhpd"
	case ZydisMnemonicVloadunpackhps:
		return "Zydis Mnemonic Vloadunpackhps"
	case ZydisMnemonicVloadunpackhq:
		return "Zydis Mnemonic Vloadunpackhq"
	case ZydisMnemonicVloadunpackld:
		return "Zydis Mnemonic Vloadunpackld"
	case ZydisMnemonicVloadunpacklpd:
		return "Zydis Mnemonic Vloadunpacklpd"
	case ZydisMnemonicVloadunpacklps:
		return "Zydis Mnemonic Vloadunpacklps"
	case ZydisMnemonicVloadunpacklq:
		return "Zydis Mnemonic Vloadunpacklq"
	case ZydisMnemonicVlog2ps:
		return "Zydis Mnemonic Vlog 2ps"
	case ZydisMnemonicVmaskmovdqu:
		return "Zydis Mnemonic Vmaskmovdqu"
	case ZydisMnemonicVmaskmovpd:
		return "Zydis Mnemonic Vmaskmovpd"
	case ZydisMnemonicVmaskmovps:
		return "Zydis Mnemonic Vmaskmovps"
	case ZydisMnemonicVmaxpd:
		return "Zydis Mnemonic Vmaxpd"
	case ZydisMnemonicVmaxph:
		return "Zydis Mnemonic Vmaxph"
	case ZydisMnemonicVmaxps:
		return "Zydis Mnemonic Vmaxps"
	case ZydisMnemonicVmaxsd:
		return "Zydis Mnemonic Vmaxsd"
	case ZydisMnemonicVmaxsh:
		return "Zydis Mnemonic Vmaxsh"
	case ZydisMnemonicVmaxss:
		return "Zydis Mnemonic Vmaxss"
	case ZydisMnemonicVmcall:
		return "Zydis Mnemonic Vmcall"
	case ZydisMnemonicVmclear:
		return "Zydis Mnemonic Vmclear"
	case ZydisMnemonicVmfunc:
		return "Zydis Mnemonic Vmfunc"
	case ZydisMnemonicVminpd:
		return "Zydis Mnemonic Vminpd"
	case ZydisMnemonicVminph:
		return "Zydis Mnemonic Vminph"
	case ZydisMnemonicVminps:
		return "Zydis Mnemonic Vminps"
	case ZydisMnemonicVminsd:
		return "Zydis Mnemonic Vminsd"
	case ZydisMnemonicVminsh:
		return "Zydis Mnemonic Vminsh"
	case ZydisMnemonicVminss:
		return "Zydis Mnemonic Vminss"
	case ZydisMnemonicVmlaunch:
		return "Zydis Mnemonic Vmlaunch"
	case ZydisMnemonicVmload:
		return "Zydis Mnemonic Vmload"
	case ZydisMnemonicVmmcall:
		return "Zydis Mnemonic Vmmcall"
	case ZydisMnemonicVmovapd:
		return "Zydis Mnemonic Vmovapd"
	case ZydisMnemonicVmovaps:
		return "Zydis Mnemonic Vmovaps"
	case ZydisMnemonicVmovd:
		return "Zydis Mnemonic Vmovd"
	case ZydisMnemonicVmovddup:
		return "Zydis Mnemonic Vmovddup"
	case ZydisMnemonicVmovdqa:
		return "Zydis Mnemonic Vmovdqa"
	case ZydisMnemonicVmovdqa32:
		return "Zydis Mnemonic Vmovdqa 32"
	case ZydisMnemonicVmovdqa64:
		return "Zydis Mnemonic Vmovdqa 64"
	case ZydisMnemonicVmovdqu:
		return "Zydis Mnemonic Vmovdqu"
	case ZydisMnemonicVmovdqu16:
		return "Zydis Mnemonic Vmovdqu 16"
	case ZydisMnemonicVmovdqu32:
		return "Zydis Mnemonic Vmovdqu 32"
	case ZydisMnemonicVmovdqu64:
		return "Zydis Mnemonic Vmovdqu 64"
	case ZydisMnemonicVmovdqu8:
		return "Zydis Mnemonic Vmovdqu 8"
	case ZydisMnemonicVmovhlps:
		return "Zydis Mnemonic Vmovhlps"
	case ZydisMnemonicVmovhpd:
		return "Zydis Mnemonic Vmovhpd"
	case ZydisMnemonicVmovhps:
		return "Zydis Mnemonic Vmovhps"
	case ZydisMnemonicVmovlhps:
		return "Zydis Mnemonic Vmovlhps"
	case ZydisMnemonicVmovlpd:
		return "Zydis Mnemonic Vmovlpd"
	case ZydisMnemonicVmovlps:
		return "Zydis Mnemonic Vmovlps"
	case ZydisMnemonicVmovmskpd:
		return "Zydis Mnemonic Vmovmskpd"
	case ZydisMnemonicVmovmskps:
		return "Zydis Mnemonic Vmovmskps"
	case ZydisMnemonicVmovnrapd:
		return "Zydis Mnemonic Vmovnrapd"
	case ZydisMnemonicVmovnraps:
		return "Zydis Mnemonic Vmovnraps"
	case ZydisMnemonicVmovnrngoapd:
		return "Zydis Mnemonic Vmovnrngoapd"
	case ZydisMnemonicVmovnrngoaps:
		return "Zydis Mnemonic Vmovnrngoaps"
	case ZydisMnemonicVmovntdq:
		return "Zydis Mnemonic Vmovntdq"
	case ZydisMnemonicVmovntdqa:
		return "Zydis Mnemonic Vmovntdqa"
	case ZydisMnemonicVmovntpd:
		return "Zydis Mnemonic Vmovntpd"
	case ZydisMnemonicVmovntps:
		return "Zydis Mnemonic Vmovntps"
	case ZydisMnemonicVmovq:
		return "Zydis Mnemonic Vmovq"
	case ZydisMnemonicVmovsd:
		return "Zydis Mnemonic Vmovsd"
	case ZydisMnemonicVmovsh:
		return "Zydis Mnemonic Vmovsh"
	case ZydisMnemonicVmovshdup:
		return "Zydis Mnemonic Vmovshdup"
	case ZydisMnemonicVmovsldup:
		return "Zydis Mnemonic Vmovsldup"
	case ZydisMnemonicVmovss:
		return "Zydis Mnemonic Vmovss"
	case ZydisMnemonicVmovupd:
		return "Zydis Mnemonic Vmovupd"
	case ZydisMnemonicVmovups:
		return "Zydis Mnemonic Vmovups"
	case ZydisMnemonicVmovw:
		return "Zydis Mnemonic Vmovw"
	case ZydisMnemonicVmpsadbw:
		return "Zydis Mnemonic Vmpsadbw"
	case ZydisMnemonicVmptrld:
		return "Zydis Mnemonic Vmptrld"
	case ZydisMnemonicVmptrst:
		return "Zydis Mnemonic Vmptrst"
	case ZydisMnemonicVmread:
		return "Zydis Mnemonic Vmread"
	case ZydisMnemonicVmresume:
		return "Zydis Mnemonic Vmresume"
	case ZydisMnemonicVmrun:
		return "Zydis Mnemonic Vmrun"
	case ZydisMnemonicVmsave:
		return "Zydis Mnemonic Vmsave"
	case ZydisMnemonicVmulpd:
		return "Zydis Mnemonic Vmulpd"
	case ZydisMnemonicVmulph:
		return "Zydis Mnemonic Vmulph"
	case ZydisMnemonicVmulps:
		return "Zydis Mnemonic Vmulps"
	case ZydisMnemonicVmulsd:
		return "Zydis Mnemonic Vmulsd"
	case ZydisMnemonicVmulsh:
		return "Zydis Mnemonic Vmulsh"
	case ZydisMnemonicVmulss:
		return "Zydis Mnemonic Vmulss"
	case ZydisMnemonicVmwrite:
		return "Zydis Mnemonic Vmwrite"
	case ZydisMnemonicVmxoff:
		return "Zydis Mnemonic Vmxoff"
	case ZydisMnemonicVmxon:
		return "Zydis Mnemonic Vmxon"
	case ZydisMnemonicVorpd:
		return "Zydis Mnemonic Vorpd"
	case ZydisMnemonicVorps:
		return "Zydis Mnemonic Vorps"
	case ZydisMnemonicVp2intersectd:
		return "Zydis Mnemonic Vp 2intersectd"
	case ZydisMnemonicVp2intersectq:
		return "Zydis Mnemonic Vp 2intersectq"
	case ZydisMnemonicVp4dpwssd:
		return "Zydis Mnemonic Vp 4dpwssd"
	case ZydisMnemonicVp4dpwssds:
		return "Zydis Mnemonic Vp 4dpwssds"
	case ZydisMnemonicVpabsb:
		return "Zydis Mnemonic Vpabsb"
	case ZydisMnemonicVpabsd:
		return "Zydis Mnemonic Vpabsd"
	case ZydisMnemonicVpabsq:
		return "Zydis Mnemonic Vpabsq"
	case ZydisMnemonicVpabsw:
		return "Zydis Mnemonic Vpabsw"
	case ZydisMnemonicVpackssdw:
		return "Zydis Mnemonic Vpackssdw"
	case ZydisMnemonicVpacksswb:
		return "Zydis Mnemonic Vpacksswb"
	case ZydisMnemonicVpackstorehd:
		return "Zydis Mnemonic Vpackstorehd"
	case ZydisMnemonicVpackstorehpd:
		return "Zydis Mnemonic Vpackstorehpd"
	case ZydisMnemonicVpackstorehps:
		return "Zydis Mnemonic Vpackstorehps"
	case ZydisMnemonicVpackstorehq:
		return "Zydis Mnemonic Vpackstorehq"
	case ZydisMnemonicVpackstoreld:
		return "Zydis Mnemonic Vpackstoreld"
	case ZydisMnemonicVpackstorelpd:
		return "Zydis Mnemonic Vpackstorelpd"
	case ZydisMnemonicVpackstorelps:
		return "Zydis Mnemonic Vpackstorelps"
	case ZydisMnemonicVpackstorelq:
		return "Zydis Mnemonic Vpackstorelq"
	case ZydisMnemonicVpackusdw:
		return "Zydis Mnemonic Vpackusdw"
	case ZydisMnemonicVpackuswb:
		return "Zydis Mnemonic Vpackuswb"
	case ZydisMnemonicVpadcd:
		return "Zydis Mnemonic Vpadcd"
	case ZydisMnemonicVpaddb:
		return "Zydis Mnemonic Vpaddb"
	case ZydisMnemonicVpaddd:
		return "Zydis Mnemonic Vpaddd"
	case ZydisMnemonicVpaddq:
		return "Zydis Mnemonic Vpaddq"
	case ZydisMnemonicVpaddsb:
		return "Zydis Mnemonic Vpaddsb"
	case ZydisMnemonicVpaddsetcd:
		return "Zydis Mnemonic Vpaddsetcd"
	case ZydisMnemonicVpaddsetsd:
		return "Zydis Mnemonic Vpaddsetsd"
	case ZydisMnemonicVpaddsw:
		return "Zydis Mnemonic Vpaddsw"
	case ZydisMnemonicVpaddusb:
		return "Zydis Mnemonic Vpaddusb"
	case ZydisMnemonicVpaddusw:
		return "Zydis Mnemonic Vpaddusw"
	case ZydisMnemonicVpaddw:
		return "Zydis Mnemonic Vpaddw"
	case ZydisMnemonicVpalignr:
		return "Zydis Mnemonic Vpalignr"
	case ZydisMnemonicVpand:
		return "Zydis Mnemonic Vpand"
	case ZydisMnemonicVpandd:
		return "Zydis Mnemonic Vpandd"
	case ZydisMnemonicVpandn:
		return "Zydis Mnemonic Vpandn"
	case ZydisMnemonicVpandnd:
		return "Zydis Mnemonic Vpandnd"
	case ZydisMnemonicVpandnq:
		return "Zydis Mnemonic Vpandnq"
	case ZydisMnemonicVpandq:
		return "Zydis Mnemonic Vpandq"
	case ZydisMnemonicVpavgb:
		return "Zydis Mnemonic Vpavgb"
	case ZydisMnemonicVpavgw:
		return "Zydis Mnemonic Vpavgw"
	case ZydisMnemonicVpblendd:
		return "Zydis Mnemonic Vpblendd"
	case ZydisMnemonicVpblendmb:
		return "Zydis Mnemonic Vpblendmb"
	case ZydisMnemonicVpblendmd:
		return "Zydis Mnemonic Vpblendmd"
	case ZydisMnemonicVpblendmq:
		return "Zydis Mnemonic Vpblendmq"
	case ZydisMnemonicVpblendmw:
		return "Zydis Mnemonic Vpblendmw"
	case ZydisMnemonicVpblendvb:
		return "Zydis Mnemonic Vpblendvb"
	case ZydisMnemonicVpblendw:
		return "Zydis Mnemonic Vpblendw"
	case ZydisMnemonicVpbroadcastb:
		return "Zydis Mnemonic Vpbroadcastb"
	case ZydisMnemonicVpbroadcastd:
		return "Zydis Mnemonic Vpbroadcastd"
	case ZydisMnemonicVpbroadcastmb2q:
		return "Zydis Mnemonic Vpbroadcastmb 2q"
	case ZydisMnemonicVpbroadcastmw2d:
		return "Zydis Mnemonic Vpbroadcastmw 2d"
	case ZydisMnemonicVpbroadcastq:
		return "Zydis Mnemonic Vpbroadcastq"
	case ZydisMnemonicVpbroadcastw:
		return "Zydis Mnemonic Vpbroadcastw"
	case ZydisMnemonicVpclmulqdq:
		return "Zydis Mnemonic Vpclmulqdq"
	case ZydisMnemonicVpcmov:
		return "Zydis Mnemonic Vpcmov"
	case ZydisMnemonicVpcmpb:
		return "Zydis Mnemonic Vpcmpb"
	case ZydisMnemonicVpcmpd:
		return "Zydis Mnemonic Vpcmpd"
	case ZydisMnemonicVpcmpeqb:
		return "Zydis Mnemonic Vpcmpeqb"
	case ZydisMnemonicVpcmpeqd:
		return "Zydis Mnemonic Vpcmpeqd"
	case ZydisMnemonicVpcmpeqq:
		return "Zydis Mnemonic Vpcmpeqq"
	case ZydisMnemonicVpcmpeqw:
		return "Zydis Mnemonic Vpcmpeqw"
	case ZydisMnemonicVpcmpestri:
		return "Zydis Mnemonic Vpcmpestri"
	case ZydisMnemonicVpcmpestrm:
		return "Zydis Mnemonic Vpcmpestrm"
	case ZydisMnemonicVpcmpgtb:
		return "Zydis Mnemonic Vpcmpgtb"
	case ZydisMnemonicVpcmpgtd:
		return "Zydis Mnemonic Vpcmpgtd"
	case ZydisMnemonicVpcmpgtq:
		return "Zydis Mnemonic Vpcmpgtq"
	case ZydisMnemonicVpcmpgtw:
		return "Zydis Mnemonic Vpcmpgtw"
	case ZydisMnemonicVpcmpistri:
		return "Zydis Mnemonic Vpcmpistri"
	case ZydisMnemonicVpcmpistrm:
		return "Zydis Mnemonic Vpcmpistrm"
	case ZydisMnemonicVpcmpltd:
		return "Zydis Mnemonic Vpcmpltd"
	case ZydisMnemonicVpcmpq:
		return "Zydis Mnemonic Vpcmpq"
	case ZydisMnemonicVpcmpub:
		return "Zydis Mnemonic Vpcmpub"
	case ZydisMnemonicVpcmpud:
		return "Zydis Mnemonic Vpcmpud"
	case ZydisMnemonicVpcmpuq:
		return "Zydis Mnemonic Vpcmpuq"
	case ZydisMnemonicVpcmpuw:
		return "Zydis Mnemonic Vpcmpuw"
	case ZydisMnemonicVpcmpw:
		return "Zydis Mnemonic Vpcmpw"
	case ZydisMnemonicVpcomb:
		return "Zydis Mnemonic Vpcomb"
	case ZydisMnemonicVpcomd:
		return "Zydis Mnemonic Vpcomd"
	case ZydisMnemonicVpcompressb:
		return "Zydis Mnemonic Vpcompressb"
	case ZydisMnemonicVpcompressd:
		return "Zydis Mnemonic Vpcompressd"
	case ZydisMnemonicVpcompressq:
		return "Zydis Mnemonic Vpcompressq"
	case ZydisMnemonicVpcompressw:
		return "Zydis Mnemonic Vpcompressw"
	case ZydisMnemonicVpcomq:
		return "Zydis Mnemonic Vpcomq"
	case ZydisMnemonicVpcomub:
		return "Zydis Mnemonic Vpcomub"
	case ZydisMnemonicVpcomud:
		return "Zydis Mnemonic Vpcomud"
	case ZydisMnemonicVpcomuq:
		return "Zydis Mnemonic Vpcomuq"
	case ZydisMnemonicVpcomuw:
		return "Zydis Mnemonic Vpcomuw"
	case ZydisMnemonicVpcomw:
		return "Zydis Mnemonic Vpcomw"
	case ZydisMnemonicVpconflictd:
		return "Zydis Mnemonic Vpconflictd"
	case ZydisMnemonicVpconflictq:
		return "Zydis Mnemonic Vpconflictq"
	case ZydisMnemonicVpdpbssd:
		return "Zydis Mnemonic Vpdpbssd"
	case ZydisMnemonicVpdpbssds:
		return "Zydis Mnemonic Vpdpbssds"
	case ZydisMnemonicVpdpbsud:
		return "Zydis Mnemonic Vpdpbsud"
	case ZydisMnemonicVpdpbsuds:
		return "Zydis Mnemonic Vpdpbsuds"
	case ZydisMnemonicVpdpbusd:
		return "Zydis Mnemonic Vpdpbusd"
	case ZydisMnemonicVpdpbusds:
		return "Zydis Mnemonic Vpdpbusds"
	case ZydisMnemonicVpdpbuud:
		return "Zydis Mnemonic Vpdpbuud"
	case ZydisMnemonicVpdpbuuds:
		return "Zydis Mnemonic Vpdpbuuds"
	case ZydisMnemonicVpdpwssd:
		return "Zydis Mnemonic Vpdpwssd"
	case ZydisMnemonicVpdpwssds:
		return "Zydis Mnemonic Vpdpwssds"
	case ZydisMnemonicVpdpwsud:
		return "Zydis Mnemonic Vpdpwsud"
	case ZydisMnemonicVpdpwsuds:
		return "Zydis Mnemonic Vpdpwsuds"
	case ZydisMnemonicVpdpwusd:
		return "Zydis Mnemonic Vpdpwusd"
	case ZydisMnemonicVpdpwusds:
		return "Zydis Mnemonic Vpdpwusds"
	case ZydisMnemonicVpdpwuud:
		return "Zydis Mnemonic Vpdpwuud"
	case ZydisMnemonicVpdpwuuds:
		return "Zydis Mnemonic Vpdpwuuds"
	case ZydisMnemonicVperm2f128:
		return "Zydis Mnemonic Vperm 2f 128"
	case ZydisMnemonicVperm2i128:
		return "Zydis Mnemonic Vperm 2i 128"
	case ZydisMnemonicVpermb:
		return "Zydis Mnemonic Vpermb"
	case ZydisMnemonicVpermd:
		return "Zydis Mnemonic Vpermd"
	case ZydisMnemonicVpermf32x4:
		return "Zydis Mnemonic Vpermf 32x 4"
	case ZydisMnemonicVpermi2b:
		return "Zydis Mnemonic Vpermi 2b"
	case ZydisMnemonicVpermi2d:
		return "Zydis Mnemonic Vpermi 2d"
	case ZydisMnemonicVpermi2pd:
		return "Zydis Mnemonic Vpermi 2pd"
	case ZydisMnemonicVpermi2ps:
		return "Zydis Mnemonic Vpermi 2ps"
	case ZydisMnemonicVpermi2q:
		return "Zydis Mnemonic Vpermi 2q"
	case ZydisMnemonicVpermi2w:
		return "Zydis Mnemonic Vpermi 2w"
	case ZydisMnemonicVpermil2pd:
		return "Zydis Mnemonic Vpermil 2pd"
	case ZydisMnemonicVpermil2ps:
		return "Zydis Mnemonic Vpermil 2ps"
	case ZydisMnemonicVpermilpd:
		return "Zydis Mnemonic Vpermilpd"
	case ZydisMnemonicVpermilps:
		return "Zydis Mnemonic Vpermilps"
	case ZydisMnemonicVpermpd:
		return "Zydis Mnemonic Vpermpd"
	case ZydisMnemonicVpermps:
		return "Zydis Mnemonic Vpermps"
	case ZydisMnemonicVpermq:
		return "Zydis Mnemonic Vpermq"
	case ZydisMnemonicVpermt2b:
		return "Zydis Mnemonic Vpermt 2b"
	case ZydisMnemonicVpermt2d:
		return "Zydis Mnemonic Vpermt 2d"
	case ZydisMnemonicVpermt2pd:
		return "Zydis Mnemonic Vpermt 2pd"
	case ZydisMnemonicVpermt2ps:
		return "Zydis Mnemonic Vpermt 2ps"
	case ZydisMnemonicVpermt2q:
		return "Zydis Mnemonic Vpermt 2q"
	case ZydisMnemonicVpermt2w:
		return "Zydis Mnemonic Vpermt 2w"
	case ZydisMnemonicVpermw:
		return "Zydis Mnemonic Vpermw"
	case ZydisMnemonicVpexpandb:
		return "Zydis Mnemonic Vpexpandb"
	case ZydisMnemonicVpexpandd:
		return "Zydis Mnemonic Vpexpandd"
	case ZydisMnemonicVpexpandq:
		return "Zydis Mnemonic Vpexpandq"
	case ZydisMnemonicVpexpandw:
		return "Zydis Mnemonic Vpexpandw"
	case ZydisMnemonicVpextrb:
		return "Zydis Mnemonic Vpextrb"
	case ZydisMnemonicVpextrd:
		return "Zydis Mnemonic Vpextrd"
	case ZydisMnemonicVpextrq:
		return "Zydis Mnemonic Vpextrq"
	case ZydisMnemonicVpextrw:
		return "Zydis Mnemonic Vpextrw"
	case ZydisMnemonicVpgatherdd:
		return "Zydis Mnemonic Vpgatherdd"
	case ZydisMnemonicVpgatherdq:
		return "Zydis Mnemonic Vpgatherdq"
	case ZydisMnemonicVpgatherqd:
		return "Zydis Mnemonic Vpgatherqd"
	case ZydisMnemonicVpgatherqq:
		return "Zydis Mnemonic Vpgatherqq"
	case ZydisMnemonicVphaddbd:
		return "Zydis Mnemonic Vphaddbd"
	case ZydisMnemonicVphaddbq:
		return "Zydis Mnemonic Vphaddbq"
	case ZydisMnemonicVphaddbw:
		return "Zydis Mnemonic Vphaddbw"
	case ZydisMnemonicVphaddd:
		return "Zydis Mnemonic Vphaddd"
	case ZydisMnemonicVphadddq:
		return "Zydis Mnemonic Vphadddq"
	case ZydisMnemonicVphaddsw:
		return "Zydis Mnemonic Vphaddsw"
	case ZydisMnemonicVphaddubd:
		return "Zydis Mnemonic Vphaddubd"
	case ZydisMnemonicVphaddubq:
		return "Zydis Mnemonic Vphaddubq"
	case ZydisMnemonicVphaddubw:
		return "Zydis Mnemonic Vphaddubw"
	case ZydisMnemonicVphaddudq:
		return "Zydis Mnemonic Vphaddudq"
	case ZydisMnemonicVphadduwd:
		return "Zydis Mnemonic Vphadduwd"
	case ZydisMnemonicVphadduwq:
		return "Zydis Mnemonic Vphadduwq"
	case ZydisMnemonicVphaddw:
		return "Zydis Mnemonic Vphaddw"
	case ZydisMnemonicVphaddwd:
		return "Zydis Mnemonic Vphaddwd"
	case ZydisMnemonicVphaddwq:
		return "Zydis Mnemonic Vphaddwq"
	case ZydisMnemonicVphminposuw:
		return "Zydis Mnemonic Vphminposuw"
	case ZydisMnemonicVphsubbw:
		return "Zydis Mnemonic Vphsubbw"
	case ZydisMnemonicVphsubd:
		return "Zydis Mnemonic Vphsubd"
	case ZydisMnemonicVphsubdq:
		return "Zydis Mnemonic Vphsubdq"
	case ZydisMnemonicVphsubsw:
		return "Zydis Mnemonic Vphsubsw"
	case ZydisMnemonicVphsubw:
		return "Zydis Mnemonic Vphsubw"
	case ZydisMnemonicVphsubwd:
		return "Zydis Mnemonic Vphsubwd"
	case ZydisMnemonicVpinsrb:
		return "Zydis Mnemonic Vpinsrb"
	case ZydisMnemonicVpinsrd:
		return "Zydis Mnemonic Vpinsrd"
	case ZydisMnemonicVpinsrq:
		return "Zydis Mnemonic Vpinsrq"
	case ZydisMnemonicVpinsrw:
		return "Zydis Mnemonic Vpinsrw"
	case ZydisMnemonicVplzcntd:
		return "Zydis Mnemonic Vplzcntd"
	case ZydisMnemonicVplzcntq:
		return "Zydis Mnemonic Vplzcntq"
	case ZydisMnemonicVpmacsdd:
		return "Zydis Mnemonic Vpmacsdd"
	case ZydisMnemonicVpmacsdqh:
		return "Zydis Mnemonic Vpmacsdqh"
	case ZydisMnemonicVpmacsdql:
		return "Zydis Mnemonic Vpmacsdql"
	case ZydisMnemonicVpmacssdd:
		return "Zydis Mnemonic Vpmacssdd"
	case ZydisMnemonicVpmacssdqh:
		return "Zydis Mnemonic Vpmacssdqh"
	case ZydisMnemonicVpmacssdql:
		return "Zydis Mnemonic Vpmacssdql"
	case ZydisMnemonicVpmacsswd:
		return "Zydis Mnemonic Vpmacsswd"
	case ZydisMnemonicVpmacssww:
		return "Zydis Mnemonic Vpmacssww"
	case ZydisMnemonicVpmacswd:
		return "Zydis Mnemonic Vpmacswd"
	case ZydisMnemonicVpmacsww:
		return "Zydis Mnemonic Vpmacsww"
	case ZydisMnemonicVpmadcsswd:
		return "Zydis Mnemonic Vpmadcsswd"
	case ZydisMnemonicVpmadcswd:
		return "Zydis Mnemonic Vpmadcswd"
	case ZydisMnemonicVpmadd231d:
		return "Zydis Mnemonic Vpmadd 231d"
	case ZydisMnemonicVpmadd233d:
		return "Zydis Mnemonic Vpmadd 233d"
	case ZydisMnemonicVpmadd52huq:
		return "Zydis Mnemonic Vpmadd 52huq"
	case ZydisMnemonicVpmadd52luq:
		return "Zydis Mnemonic Vpmadd 52luq"
	case ZydisMnemonicVpmaddubsw:
		return "Zydis Mnemonic Vpmaddubsw"
	case ZydisMnemonicVpmaddwd:
		return "Zydis Mnemonic Vpmaddwd"
	case ZydisMnemonicVpmaskmovd:
		return "Zydis Mnemonic Vpmaskmovd"
	case ZydisMnemonicVpmaskmovq:
		return "Zydis Mnemonic Vpmaskmovq"
	case ZydisMnemonicVpmaxsb:
		return "Zydis Mnemonic Vpmaxsb"
	case ZydisMnemonicVpmaxsd:
		return "Zydis Mnemonic Vpmaxsd"
	case ZydisMnemonicVpmaxsq:
		return "Zydis Mnemonic Vpmaxsq"
	case ZydisMnemonicVpmaxsw:
		return "Zydis Mnemonic Vpmaxsw"
	case ZydisMnemonicVpmaxub:
		return "Zydis Mnemonic Vpmaxub"
	case ZydisMnemonicVpmaxud:
		return "Zydis Mnemonic Vpmaxud"
	case ZydisMnemonicVpmaxuq:
		return "Zydis Mnemonic Vpmaxuq"
	case ZydisMnemonicVpmaxuw:
		return "Zydis Mnemonic Vpmaxuw"
	case ZydisMnemonicVpminsb:
		return "Zydis Mnemonic Vpminsb"
	case ZydisMnemonicVpminsd:
		return "Zydis Mnemonic Vpminsd"
	case ZydisMnemonicVpminsq:
		return "Zydis Mnemonic Vpminsq"
	case ZydisMnemonicVpminsw:
		return "Zydis Mnemonic Vpminsw"
	case ZydisMnemonicVpminub:
		return "Zydis Mnemonic Vpminub"
	case ZydisMnemonicVpminud:
		return "Zydis Mnemonic Vpminud"
	case ZydisMnemonicVpminuq:
		return "Zydis Mnemonic Vpminuq"
	case ZydisMnemonicVpminuw:
		return "Zydis Mnemonic Vpminuw"
	case ZydisMnemonicVpmovb2m:
		return "Zydis Mnemonic Vpmovb 2m"
	case ZydisMnemonicVpmovd2m:
		return "Zydis Mnemonic Vpmovd 2m"
	case ZydisMnemonicVpmovdb:
		return "Zydis Mnemonic Vpmovdb"
	case ZydisMnemonicVpmovdw:
		return "Zydis Mnemonic Vpmovdw"
	case ZydisMnemonicVpmovm2b:
		return "Zydis Mnemonic Vpmovm 2b"
	case ZydisMnemonicVpmovm2d:
		return "Zydis Mnemonic Vpmovm 2d"
	case ZydisMnemonicVpmovm2q:
		return "Zydis Mnemonic Vpmovm 2q"
	case ZydisMnemonicVpmovm2w:
		return "Zydis Mnemonic Vpmovm 2w"
	case ZydisMnemonicVpmovmskb:
		return "Zydis Mnemonic Vpmovmskb"
	case ZydisMnemonicVpmovq2m:
		return "Zydis Mnemonic Vpmovq 2m"
	case ZydisMnemonicVpmovqb:
		return "Zydis Mnemonic Vpmovqb"
	case ZydisMnemonicVpmovqd:
		return "Zydis Mnemonic Vpmovqd"
	case ZydisMnemonicVpmovqw:
		return "Zydis Mnemonic Vpmovqw"
	case ZydisMnemonicVpmovsdb:
		return "Zydis Mnemonic Vpmovsdb"
	case ZydisMnemonicVpmovsdw:
		return "Zydis Mnemonic Vpmovsdw"
	case ZydisMnemonicVpmovsqb:
		return "Zydis Mnemonic Vpmovsqb"
	case ZydisMnemonicVpmovsqd:
		return "Zydis Mnemonic Vpmovsqd"
	case ZydisMnemonicVpmovsqw:
		return "Zydis Mnemonic Vpmovsqw"
	case ZydisMnemonicVpmovswb:
		return "Zydis Mnemonic Vpmovswb"
	case ZydisMnemonicVpmovsxbd:
		return "Zydis Mnemonic Vpmovsxbd"
	case ZydisMnemonicVpmovsxbq:
		return "Zydis Mnemonic Vpmovsxbq"
	case ZydisMnemonicVpmovsxbw:
		return "Zydis Mnemonic Vpmovsxbw"
	case ZydisMnemonicVpmovsxdq:
		return "Zydis Mnemonic Vpmovsxdq"
	case ZydisMnemonicVpmovsxwd:
		return "Zydis Mnemonic Vpmovsxwd"
	case ZydisMnemonicVpmovsxwq:
		return "Zydis Mnemonic Vpmovsxwq"
	case ZydisMnemonicVpmovusdb:
		return "Zydis Mnemonic Vpmovusdb"
	case ZydisMnemonicVpmovusdw:
		return "Zydis Mnemonic Vpmovusdw"
	case ZydisMnemonicVpmovusqb:
		return "Zydis Mnemonic Vpmovusqb"
	case ZydisMnemonicVpmovusqd:
		return "Zydis Mnemonic Vpmovusqd"
	case ZydisMnemonicVpmovusqw:
		return "Zydis Mnemonic Vpmovusqw"
	case ZydisMnemonicVpmovuswb:
		return "Zydis Mnemonic Vpmovuswb"
	case ZydisMnemonicVpmovw2m:
		return "Zydis Mnemonic Vpmovw 2m"
	case ZydisMnemonicVpmovwb:
		return "Zydis Mnemonic Vpmovwb"
	case ZydisMnemonicVpmovzxbd:
		return "Zydis Mnemonic Vpmovzxbd"
	case ZydisMnemonicVpmovzxbq:
		return "Zydis Mnemonic Vpmovzxbq"
	case ZydisMnemonicVpmovzxbw:
		return "Zydis Mnemonic Vpmovzxbw"
	case ZydisMnemonicVpmovzxdq:
		return "Zydis Mnemonic Vpmovzxdq"
	case ZydisMnemonicVpmovzxwd:
		return "Zydis Mnemonic Vpmovzxwd"
	case ZydisMnemonicVpmovzxwq:
		return "Zydis Mnemonic Vpmovzxwq"
	case ZydisMnemonicVpmuldq:
		return "Zydis Mnemonic Vpmuldq"
	case ZydisMnemonicVpmulhd:
		return "Zydis Mnemonic Vpmulhd"
	case ZydisMnemonicVpmulhrsw:
		return "Zydis Mnemonic Vpmulhrsw"
	case ZydisMnemonicVpmulhud:
		return "Zydis Mnemonic Vpmulhud"
	case ZydisMnemonicVpmulhuw:
		return "Zydis Mnemonic Vpmulhuw"
	case ZydisMnemonicVpmulhw:
		return "Zydis Mnemonic Vpmulhw"
	case ZydisMnemonicVpmulld:
		return "Zydis Mnemonic Vpmulld"
	case ZydisMnemonicVpmullq:
		return "Zydis Mnemonic Vpmullq"
	case ZydisMnemonicVpmullw:
		return "Zydis Mnemonic Vpmullw"
	case ZydisMnemonicVpmultishiftqb:
		return "Zydis Mnemonic Vpmultishiftqb"
	case ZydisMnemonicVpmuludq:
		return "Zydis Mnemonic Vpmuludq"
	case ZydisMnemonicVpopcntb:
		return "Zydis Mnemonic Vpopcntb"
	case ZydisMnemonicVpopcntd:
		return "Zydis Mnemonic Vpopcntd"
	case ZydisMnemonicVpopcntq:
		return "Zydis Mnemonic Vpopcntq"
	case ZydisMnemonicVpopcntw:
		return "Zydis Mnemonic Vpopcntw"
	case ZydisMnemonicVpor:
		return "Zydis Mnemonic Vpor"
	case ZydisMnemonicVpord:
		return "Zydis Mnemonic Vpord"
	case ZydisMnemonicVporq:
		return "Zydis Mnemonic Vporq"
	case ZydisMnemonicVpperm:
		return "Zydis Mnemonic Vpperm"
	case ZydisMnemonicVprefetch0:
		return "Zydis Mnemonic Vprefetch 0"
	case ZydisMnemonicVprefetch1:
		return "Zydis Mnemonic Vprefetch 1"
	case ZydisMnemonicVprefetch2:
		return "Zydis Mnemonic Vprefetch 2"
	case ZydisMnemonicVprefetche0:
		return "Zydis Mnemonic Vprefetche 0"
	case ZydisMnemonicVprefetche1:
		return "Zydis Mnemonic Vprefetche 1"
	case ZydisMnemonicVprefetche2:
		return "Zydis Mnemonic Vprefetche 2"
	case ZydisMnemonicVprefetchenta:
		return "Zydis Mnemonic Vprefetchenta"
	case ZydisMnemonicVprefetchnta:
		return "Zydis Mnemonic Vprefetchnta"
	case ZydisMnemonicVprold:
		return "Zydis Mnemonic Vprold"
	case ZydisMnemonicVprolq:
		return "Zydis Mnemonic Vprolq"
	case ZydisMnemonicVprolvd:
		return "Zydis Mnemonic Vprolvd"
	case ZydisMnemonicVprolvq:
		return "Zydis Mnemonic Vprolvq"
	case ZydisMnemonicVprord:
		return "Zydis Mnemonic Vprord"
	case ZydisMnemonicVprorq:
		return "Zydis Mnemonic Vprorq"
	case ZydisMnemonicVprorvd:
		return "Zydis Mnemonic Vprorvd"
	case ZydisMnemonicVprorvq:
		return "Zydis Mnemonic Vprorvq"
	case ZydisMnemonicVprotb:
		return "Zydis Mnemonic Vprotb"
	case ZydisMnemonicVprotd:
		return "Zydis Mnemonic Vprotd"
	case ZydisMnemonicVprotq:
		return "Zydis Mnemonic Vprotq"
	case ZydisMnemonicVprotw:
		return "Zydis Mnemonic Vprotw"
	case ZydisMnemonicVpsadbw:
		return "Zydis Mnemonic Vpsadbw"
	case ZydisMnemonicVpsbbd:
		return "Zydis Mnemonic Vpsbbd"
	case ZydisMnemonicVpsbbrd:
		return "Zydis Mnemonic Vpsbbrd"
	case ZydisMnemonicVpscatterdd:
		return "Zydis Mnemonic Vpscatterdd"
	case ZydisMnemonicVpscatterdq:
		return "Zydis Mnemonic Vpscatterdq"
	case ZydisMnemonicVpscatterqd:
		return "Zydis Mnemonic Vpscatterqd"
	case ZydisMnemonicVpscatterqq:
		return "Zydis Mnemonic Vpscatterqq"
	case ZydisMnemonicVpshab:
		return "Zydis Mnemonic Vpshab"
	case ZydisMnemonicVpshad:
		return "Zydis Mnemonic Vpshad"
	case ZydisMnemonicVpshaq:
		return "Zydis Mnemonic Vpshaq"
	case ZydisMnemonicVpshaw:
		return "Zydis Mnemonic Vpshaw"
	case ZydisMnemonicVpshlb:
		return "Zydis Mnemonic Vpshlb"
	case ZydisMnemonicVpshld:
		return "Zydis Mnemonic Vpshld"
	case ZydisMnemonicVpshldd:
		return "Zydis Mnemonic Vpshldd"
	case ZydisMnemonicVpshldq:
		return "Zydis Mnemonic Vpshldq"
	case ZydisMnemonicVpshldvd:
		return "Zydis Mnemonic Vpshldvd"
	case ZydisMnemonicVpshldvq:
		return "Zydis Mnemonic Vpshldvq"
	case ZydisMnemonicVpshldvw:
		return "Zydis Mnemonic Vpshldvw"
	case ZydisMnemonicVpshldw:
		return "Zydis Mnemonic Vpshldw"
	case ZydisMnemonicVpshlq:
		return "Zydis Mnemonic Vpshlq"
	case ZydisMnemonicVpshlw:
		return "Zydis Mnemonic Vpshlw"
	case ZydisMnemonicVpshrdd:
		return "Zydis Mnemonic Vpshrdd"
	case ZydisMnemonicVpshrdq:
		return "Zydis Mnemonic Vpshrdq"
	case ZydisMnemonicVpshrdvd:
		return "Zydis Mnemonic Vpshrdvd"
	case ZydisMnemonicVpshrdvq:
		return "Zydis Mnemonic Vpshrdvq"
	case ZydisMnemonicVpshrdvw:
		return "Zydis Mnemonic Vpshrdvw"
	case ZydisMnemonicVpshrdw:
		return "Zydis Mnemonic Vpshrdw"
	case ZydisMnemonicVpshufb:
		return "Zydis Mnemonic Vpshufb"
	case ZydisMnemonicVpshufbitqmb:
		return "Zydis Mnemonic Vpshufbitqmb"
	case ZydisMnemonicVpshufd:
		return "Zydis Mnemonic Vpshufd"
	case ZydisMnemonicVpshufhw:
		return "Zydis Mnemonic Vpshufhw"
	case ZydisMnemonicVpshuflw:
		return "Zydis Mnemonic Vpshuflw"
	case ZydisMnemonicVpsignb:
		return "Zydis Mnemonic Vpsignb"
	case ZydisMnemonicVpsignd:
		return "Zydis Mnemonic Vpsignd"
	case ZydisMnemonicVpsignw:
		return "Zydis Mnemonic Vpsignw"
	case ZydisMnemonicVpslld:
		return "Zydis Mnemonic Vpslld"
	case ZydisMnemonicVpslldq:
		return "Zydis Mnemonic Vpslldq"
	case ZydisMnemonicVpsllq:
		return "Zydis Mnemonic Vpsllq"
	case ZydisMnemonicVpsllvd:
		return "Zydis Mnemonic Vpsllvd"
	case ZydisMnemonicVpsllvq:
		return "Zydis Mnemonic Vpsllvq"
	case ZydisMnemonicVpsllvw:
		return "Zydis Mnemonic Vpsllvw"
	case ZydisMnemonicVpsllw:
		return "Zydis Mnemonic Vpsllw"
	case ZydisMnemonicVpsrad:
		return "Zydis Mnemonic Vpsrad"
	case ZydisMnemonicVpsraq:
		return "Zydis Mnemonic Vpsraq"
	case ZydisMnemonicVpsravd:
		return "Zydis Mnemonic Vpsravd"
	case ZydisMnemonicVpsravq:
		return "Zydis Mnemonic Vpsravq"
	case ZydisMnemonicVpsravw:
		return "Zydis Mnemonic Vpsravw"
	case ZydisMnemonicVpsraw:
		return "Zydis Mnemonic Vpsraw"
	case ZydisMnemonicVpsrld:
		return "Zydis Mnemonic Vpsrld"
	case ZydisMnemonicVpsrldq:
		return "Zydis Mnemonic Vpsrldq"
	case ZydisMnemonicVpsrlq:
		return "Zydis Mnemonic Vpsrlq"
	case ZydisMnemonicVpsrlvd:
		return "Zydis Mnemonic Vpsrlvd"
	case ZydisMnemonicVpsrlvq:
		return "Zydis Mnemonic Vpsrlvq"
	case ZydisMnemonicVpsrlvw:
		return "Zydis Mnemonic Vpsrlvw"
	case ZydisMnemonicVpsrlw:
		return "Zydis Mnemonic Vpsrlw"
	case ZydisMnemonicVpsubb:
		return "Zydis Mnemonic Vpsubb"
	case ZydisMnemonicVpsubd:
		return "Zydis Mnemonic Vpsubd"
	case ZydisMnemonicVpsubq:
		return "Zydis Mnemonic Vpsubq"
	case ZydisMnemonicVpsubrd:
		return "Zydis Mnemonic Vpsubrd"
	case ZydisMnemonicVpsubrsetbd:
		return "Zydis Mnemonic Vpsubrsetbd"
	case ZydisMnemonicVpsubsb:
		return "Zydis Mnemonic Vpsubsb"
	case ZydisMnemonicVpsubsetbd:
		return "Zydis Mnemonic Vpsubsetbd"
	case ZydisMnemonicVpsubsw:
		return "Zydis Mnemonic Vpsubsw"
	case ZydisMnemonicVpsubusb:
		return "Zydis Mnemonic Vpsubusb"
	case ZydisMnemonicVpsubusw:
		return "Zydis Mnemonic Vpsubusw"
	case ZydisMnemonicVpsubw:
		return "Zydis Mnemonic Vpsubw"
	case ZydisMnemonicVpternlogd:
		return "Zydis Mnemonic Vpternlogd"
	case ZydisMnemonicVpternlogq:
		return "Zydis Mnemonic Vpternlogq"
	case ZydisMnemonicVptest:
		return "Zydis Mnemonic Vptest"
	case ZydisMnemonicVptestmb:
		return "Zydis Mnemonic Vptestmb"
	case ZydisMnemonicVptestmd:
		return "Zydis Mnemonic Vptestmd"
	case ZydisMnemonicVptestmq:
		return "Zydis Mnemonic Vptestmq"
	case ZydisMnemonicVptestmw:
		return "Zydis Mnemonic Vptestmw"
	case ZydisMnemonicVptestnmb:
		return "Zydis Mnemonic Vptestnmb"
	case ZydisMnemonicVptestnmd:
		return "Zydis Mnemonic Vptestnmd"
	case ZydisMnemonicVptestnmq:
		return "Zydis Mnemonic Vptestnmq"
	case ZydisMnemonicVptestnmw:
		return "Zydis Mnemonic Vptestnmw"
	case ZydisMnemonicVpunpckhbw:
		return "Zydis Mnemonic Vpunpckhbw"
	case ZydisMnemonicVpunpckhdq:
		return "Zydis Mnemonic Vpunpckhdq"
	case ZydisMnemonicVpunpckhqdq:
		return "Zydis Mnemonic Vpunpckhqdq"
	case ZydisMnemonicVpunpckhwd:
		return "Zydis Mnemonic Vpunpckhwd"
	case ZydisMnemonicVpunpcklbw:
		return "Zydis Mnemonic Vpunpcklbw"
	case ZydisMnemonicVpunpckldq:
		return "Zydis Mnemonic Vpunpckldq"
	case ZydisMnemonicVpunpcklqdq:
		return "Zydis Mnemonic Vpunpcklqdq"
	case ZydisMnemonicVpunpcklwd:
		return "Zydis Mnemonic Vpunpcklwd"
	case ZydisMnemonicVpxor:
		return "Zydis Mnemonic Vpxor"
	case ZydisMnemonicVpxord:
		return "Zydis Mnemonic Vpxord"
	case ZydisMnemonicVpxorq:
		return "Zydis Mnemonic Vpxorq"
	case ZydisMnemonicVrangepd:
		return "Zydis Mnemonic Vrangepd"
	case ZydisMnemonicVrangeps:
		return "Zydis Mnemonic Vrangeps"
	case ZydisMnemonicVrangesd:
		return "Zydis Mnemonic Vrangesd"
	case ZydisMnemonicVrangess:
		return "Zydis Mnemonic Vrangess"
	case ZydisMnemonicVrcp14pd:
		return "Zydis Mnemonic Vrcp 14pd"
	case ZydisMnemonicVrcp14ps:
		return "Zydis Mnemonic Vrcp 14ps"
	case ZydisMnemonicVrcp14sd:
		return "Zydis Mnemonic Vrcp 14sd"
	case ZydisMnemonicVrcp14ss:
		return "Zydis Mnemonic Vrcp 14ss"
	case ZydisMnemonicVrcp23ps:
		return "Zydis Mnemonic Vrcp 23ps"
	case ZydisMnemonicVrcp28pd:
		return "Zydis Mnemonic Vrcp 28pd"
	case ZydisMnemonicVrcp28ps:
		return "Zydis Mnemonic Vrcp 28ps"
	case ZydisMnemonicVrcp28sd:
		return "Zydis Mnemonic Vrcp 28sd"
	case ZydisMnemonicVrcp28ss:
		return "Zydis Mnemonic Vrcp 28ss"
	case ZydisMnemonicVrcpph:
		return "Zydis Mnemonic Vrcpph"
	case ZydisMnemonicVrcpps:
		return "Zydis Mnemonic Vrcpps"
	case ZydisMnemonicVrcpsh:
		return "Zydis Mnemonic Vrcpsh"
	case ZydisMnemonicVrcpss:
		return "Zydis Mnemonic Vrcpss"
	case ZydisMnemonicVreducepd:
		return "Zydis Mnemonic Vreducepd"
	case ZydisMnemonicVreduceph:
		return "Zydis Mnemonic Vreduceph"
	case ZydisMnemonicVreduceps:
		return "Zydis Mnemonic Vreduceps"
	case ZydisMnemonicVreducesd:
		return "Zydis Mnemonic Vreducesd"
	case ZydisMnemonicVreducesh:
		return "Zydis Mnemonic Vreducesh"
	case ZydisMnemonicVreducess:
		return "Zydis Mnemonic Vreducess"
	case ZydisMnemonicVrndfxpntpd:
		return "Zydis Mnemonic Vrndfxpntpd"
	case ZydisMnemonicVrndfxpntps:
		return "Zydis Mnemonic Vrndfxpntps"
	case ZydisMnemonicVrndscalepd:
		return "Zydis Mnemonic Vrndscalepd"
	case ZydisMnemonicVrndscaleph:
		return "Zydis Mnemonic Vrndscaleph"
	case ZydisMnemonicVrndscaleps:
		return "Zydis Mnemonic Vrndscaleps"
	case ZydisMnemonicVrndscalesd:
		return "Zydis Mnemonic Vrndscalesd"
	case ZydisMnemonicVrndscalesh:
		return "Zydis Mnemonic Vrndscalesh"
	case ZydisMnemonicVrndscaless:
		return "Zydis Mnemonic Vrndscaless"
	case ZydisMnemonicVroundpd:
		return "Zydis Mnemonic Vroundpd"
	case ZydisMnemonicVroundps:
		return "Zydis Mnemonic Vroundps"
	case ZydisMnemonicVroundsd:
		return "Zydis Mnemonic Vroundsd"
	case ZydisMnemonicVroundss:
		return "Zydis Mnemonic Vroundss"
	case ZydisMnemonicVrsqrt14pd:
		return "Zydis Mnemonic Vrsqrt 14pd"
	case ZydisMnemonicVrsqrt14ps:
		return "Zydis Mnemonic Vrsqrt 14ps"
	case ZydisMnemonicVrsqrt14sd:
		return "Zydis Mnemonic Vrsqrt 14sd"
	case ZydisMnemonicVrsqrt14ss:
		return "Zydis Mnemonic Vrsqrt 14ss"
	case ZydisMnemonicVrsqrt23ps:
		return "Zydis Mnemonic Vrsqrt 23ps"
	case ZydisMnemonicVrsqrt28pd:
		return "Zydis Mnemonic Vrsqrt 28pd"
	case ZydisMnemonicVrsqrt28ps:
		return "Zydis Mnemonic Vrsqrt 28ps"
	case ZydisMnemonicVrsqrt28sd:
		return "Zydis Mnemonic Vrsqrt 28sd"
	case ZydisMnemonicVrsqrt28ss:
		return "Zydis Mnemonic Vrsqrt 28ss"
	case ZydisMnemonicVrsqrtph:
		return "Zydis Mnemonic Vrsqrtph"
	case ZydisMnemonicVrsqrtps:
		return "Zydis Mnemonic Vrsqrtps"
	case ZydisMnemonicVrsqrtsh:
		return "Zydis Mnemonic Vrsqrtsh"
	case ZydisMnemonicVrsqrtss:
		return "Zydis Mnemonic Vrsqrtss"
	case ZydisMnemonicVscalefpd:
		return "Zydis Mnemonic Vscalefpd"
	case ZydisMnemonicVscalefph:
		return "Zydis Mnemonic Vscalefph"
	case ZydisMnemonicVscalefps:
		return "Zydis Mnemonic Vscalefps"
	case ZydisMnemonicVscalefsd:
		return "Zydis Mnemonic Vscalefsd"
	case ZydisMnemonicVscalefsh:
		return "Zydis Mnemonic Vscalefsh"
	case ZydisMnemonicVscalefss:
		return "Zydis Mnemonic Vscalefss"
	case ZydisMnemonicVscaleps:
		return "Zydis Mnemonic Vscaleps"
	case ZydisMnemonicVscatterdpd:
		return "Zydis Mnemonic Vscatterdpd"
	case ZydisMnemonicVscatterdps:
		return "Zydis Mnemonic Vscatterdps"
	case ZydisMnemonicVscatterpf0dpd:
		return "Zydis Mnemonic Vscatterpf 0dpd"
	case ZydisMnemonicVscatterpf0dps:
		return "Zydis Mnemonic Vscatterpf 0dps"
	case ZydisMnemonicVscatterpf0hintdpd:
		return "Zydis Mnemonic Vscatterpf 0hintdpd"
	case ZydisMnemonicVscatterpf0hintdps:
		return "Zydis Mnemonic Vscatterpf 0hintdps"
	case ZydisMnemonicVscatterpf0qpd:
		return "Zydis Mnemonic Vscatterpf 0qpd"
	case ZydisMnemonicVscatterpf0qps:
		return "Zydis Mnemonic Vscatterpf 0qps"
	case ZydisMnemonicVscatterpf1dpd:
		return "Zydis Mnemonic Vscatterpf 1dpd"
	case ZydisMnemonicVscatterpf1dps:
		return "Zydis Mnemonic Vscatterpf 1dps"
	case ZydisMnemonicVscatterpf1qpd:
		return "Zydis Mnemonic Vscatterpf 1qpd"
	case ZydisMnemonicVscatterpf1qps:
		return "Zydis Mnemonic Vscatterpf 1qps"
	case ZydisMnemonicVscatterqpd:
		return "Zydis Mnemonic Vscatterqpd"
	case ZydisMnemonicVscatterqps:
		return "Zydis Mnemonic Vscatterqps"
	case ZydisMnemonicVsha512msg1:
		return "Zydis Mnemonic Vsha 512msg 1"
	case ZydisMnemonicVsha512msg2:
		return "Zydis Mnemonic Vsha 512msg 2"
	case ZydisMnemonicVsha512rnds2:
		return "Zydis Mnemonic Vsha 512rnds 2"
	case ZydisMnemonicVshuff32x4:
		return "Zydis Mnemonic Vshuff 32x 4"
	case ZydisMnemonicVshuff64x2:
		return "Zydis Mnemonic Vshuff 64x 2"
	case ZydisMnemonicVshufi32x4:
		return "Zydis Mnemonic Vshufi 32x 4"
	case ZydisMnemonicVshufi64x2:
		return "Zydis Mnemonic Vshufi 64x 2"
	case ZydisMnemonicVshufpd:
		return "Zydis Mnemonic Vshufpd"
	case ZydisMnemonicVshufps:
		return "Zydis Mnemonic Vshufps"
	case ZydisMnemonicVsm3msg1:
		return "Zydis Mnemonic Vsm 3msg 1"
	case ZydisMnemonicVsm3msg2:
		return "Zydis Mnemonic Vsm 3msg 2"
	case ZydisMnemonicVsm3rnds2:
		return "Zydis Mnemonic Vsm 3rnds 2"
	case ZydisMnemonicVsm4key4:
		return "Zydis Mnemonic Vsm 4key 4"
	case ZydisMnemonicVsm4rnds4:
		return "Zydis Mnemonic Vsm 4rnds 4"
	case ZydisMnemonicVsqrtpd:
		return "Zydis Mnemonic Vsqrtpd"
	case ZydisMnemonicVsqrtph:
		return "Zydis Mnemonic Vsqrtph"
	case ZydisMnemonicVsqrtps:
		return "Zydis Mnemonic Vsqrtps"
	case ZydisMnemonicVsqrtsd:
		return "Zydis Mnemonic Vsqrtsd"
	case ZydisMnemonicVsqrtsh:
		return "Zydis Mnemonic Vsqrtsh"
	case ZydisMnemonicVsqrtss:
		return "Zydis Mnemonic Vsqrtss"
	case ZydisMnemonicVstmxcsr:
		return "Zydis Mnemonic Vstmxcsr"
	case ZydisMnemonicVsubpd:
		return "Zydis Mnemonic Vsubpd"
	case ZydisMnemonicVsubph:
		return "Zydis Mnemonic Vsubph"
	case ZydisMnemonicVsubps:
		return "Zydis Mnemonic Vsubps"
	case ZydisMnemonicVsubrpd:
		return "Zydis Mnemonic Vsubrpd"
	case ZydisMnemonicVsubrps:
		return "Zydis Mnemonic Vsubrps"
	case ZydisMnemonicVsubsd:
		return "Zydis Mnemonic Vsubsd"
	case ZydisMnemonicVsubsh:
		return "Zydis Mnemonic Vsubsh"
	case ZydisMnemonicVsubss:
		return "Zydis Mnemonic Vsubss"
	case ZydisMnemonicVtestpd:
		return "Zydis Mnemonic Vtestpd"
	case ZydisMnemonicVtestps:
		return "Zydis Mnemonic Vtestps"
	case ZydisMnemonicVucomisd:
		return "Zydis Mnemonic Vucomisd"
	case ZydisMnemonicVucomish:
		return "Zydis Mnemonic Vucomish"
	case ZydisMnemonicVucomiss:
		return "Zydis Mnemonic Vucomiss"
	case ZydisMnemonicVunpckhpd:
		return "Zydis Mnemonic Vunpckhpd"
	case ZydisMnemonicVunpckhps:
		return "Zydis Mnemonic Vunpckhps"
	case ZydisMnemonicVunpcklpd:
		return "Zydis Mnemonic Vunpcklpd"
	case ZydisMnemonicVunpcklps:
		return "Zydis Mnemonic Vunpcklps"
	case ZydisMnemonicVxorpd:
		return "Zydis Mnemonic Vxorpd"
	case ZydisMnemonicVxorps:
		return "Zydis Mnemonic Vxorps"
	case ZydisMnemonicVzeroall:
		return "Zydis Mnemonic Vzeroall"
	case ZydisMnemonicVzeroupper:
		return "Zydis Mnemonic Vzeroupper"
	case ZydisMnemonicWbinvd:
		return "Zydis Mnemonic Wbinvd"
	case ZydisMnemonicWrfsbase:
		return "Zydis Mnemonic Wrfsbase"
	case ZydisMnemonicWrgsbase:
		return "Zydis Mnemonic Wrgsbase"
	case ZydisMnemonicWrmsr:
		return "Zydis Mnemonic Wrmsr"
	case ZydisMnemonicWrmsrlist:
		return "Zydis Mnemonic Wrmsrlist"
	case ZydisMnemonicWrmsrns:
		return "Zydis Mnemonic Wrmsrns"
	case ZydisMnemonicWrpkru:
		return "Zydis Mnemonic Wrpkru"
	case ZydisMnemonicWrssd:
		return "Zydis Mnemonic Wrssd"
	case ZydisMnemonicWrssq:
		return "Zydis Mnemonic Wrssq"
	case ZydisMnemonicWrussd:
		return "Zydis Mnemonic Wrussd"
	case ZydisMnemonicWrussq:
		return "Zydis Mnemonic Wrussq"
	case ZydisMnemonicXabort:
		return "Zydis Mnemonic Xabort"
	case ZydisMnemonicXadd:
		return "Zydis Mnemonic Xadd"
	case ZydisMnemonicXbegin:
		return "Zydis Mnemonic Xbegin"
	case ZydisMnemonicXchg:
		return "Zydis Mnemonic Xchg"
	case ZydisMnemonicXcryptCbc:
		return "Zydis Mnemonic Xcrypt Cbc"
	case ZydisMnemonicXcryptCfb:
		return "Zydis Mnemonic Xcrypt Cfb"
	case ZydisMnemonicXcryptCtr:
		return "Zydis Mnemonic Xcrypt Ctr"
	case ZydisMnemonicXcryptEcb:
		return "Zydis Mnemonic Xcrypt Ecb"
	case ZydisMnemonicXcryptOfb:
		return "Zydis Mnemonic Xcrypt Ofb"
	case ZydisMnemonicXend:
		return "Zydis Mnemonic Xend"
	case ZydisMnemonicXgetbv:
		return "Zydis Mnemonic Xgetbv"
	case ZydisMnemonicXlat:
		return "Zydis Mnemonic Xlat"
	case ZydisMnemonicXor:
		return "Zydis Mnemonic Xor"
	case ZydisMnemonicXorpd:
		return "Zydis Mnemonic Xorpd"
	case ZydisMnemonicXorps:
		return "Zydis Mnemonic Xorps"
	case ZydisMnemonicXresldtrk:
		return "Zydis Mnemonic Xresldtrk"
	case ZydisMnemonicXrstor:
		return "Zydis Mnemonic Xrstor"
	case ZydisMnemonicXrstor64:
		return "Zydis Mnemonic Xrstor 64"
	case ZydisMnemonicXrstors:
		return "Zydis Mnemonic Xrstors"
	case ZydisMnemonicXrstors64:
		return "Zydis Mnemonic Xrstors 64"
	case ZydisMnemonicXsave:
		return "Zydis Mnemonic Xsave"
	case ZydisMnemonicXsave64:
		return "Zydis Mnemonic Xsave 64"
	case ZydisMnemonicXsavec:
		return "Zydis Mnemonic Xsavec"
	case ZydisMnemonicXsavec64:
		return "Zydis Mnemonic Xsavec 64"
	case ZydisMnemonicXsaveopt:
		return "Zydis Mnemonic Xsaveopt"
	case ZydisMnemonicXsaveopt64:
		return "Zydis Mnemonic Xsaveopt 64"
	case ZydisMnemonicXsaves:
		return "Zydis Mnemonic Xsaves"
	case ZydisMnemonicXsaves64:
		return "Zydis Mnemonic Xsaves 64"
	case ZydisMnemonicXsetbv:
		return "Zydis Mnemonic Xsetbv"
	case ZydisMnemonicXsha1:
		return "Zydis Mnemonic Xsha 1"
	case ZydisMnemonicXsha256:
		return "Zydis Mnemonic Xsha 256"
	case ZydisMnemonicXstore:
		return "Zydis Mnemonic Xstore"
	case ZydisMnemonicXsusldtrk:
		return "Zydis Mnemonic Xsusldtrk"
	case ZydisMnemonicXtest:
		return "Zydis Mnemonic Xtest"
	default:
		return fmt.Sprintf("ZydisMnemonic_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:66 -> ZydisMachineMode_
type ZydisMachineMode_ uint32

const (
	ZydisMachineModeLong64 ZydisMachineMode_ = 0
	ZydisMachineModeLongCompat32 ZydisMachineMode_ = 1
	ZydisMachineModeLongCompat16 ZydisMachineMode_ = 2
	ZydisMachineModeLegacy32 ZydisMachineMode_ = 3
	ZydisMachineModeLegacy16 ZydisMachineMode_ = 4
	ZydisMachineModeReal16 ZydisMachineMode_ = 5
	ZydisMachineModeMaxValue ZydisMachineMode_ = 5
	ZydisMachineModeRequiredBits ZydisMachineMode_ = 3
)

func (z ZydisMachineMode_) String() string {
	switch z {
	case ZydisMachineModeLong64:
		return "Zydis Machine Mode Long 64"
	case ZydisMachineModeLongCompat32:
		return "Zydis Machine Mode Long Compat 32"
	case ZydisMachineModeLongCompat16:
		return "Zydis Machine Mode Long Compat 16"
	case ZydisMachineModeLegacy32:
		return "Zydis Machine Mode Legacy 32"
	case ZydisMachineModeLegacy16:
		return "Zydis Machine Mode Legacy 16"
	case ZydisMachineModeReal16:
		return "Zydis Machine Mode Real 16"
	default:
		return fmt.Sprintf("ZydisMachineMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:110 -> ZydisStackWidth_
type ZydisStackWidth_ uint32

const (
	ZydisStackWidth16 ZydisStackWidth_ = 0
	ZydisStackWidth32 ZydisStackWidth_ = 1
	ZydisStackWidth64 ZydisStackWidth_ = 2
	ZydisStackWidthMaxValue ZydisStackWidth_ = 2
	ZydisStackWidthRequiredBits ZydisStackWidth_ = 2
)

func (z ZydisStackWidth_) String() string {
	switch z {
	case ZydisStackWidth16:
		return "Zydis Stack Width 16"
	case ZydisStackWidth32:
		return "Zydis Stack Width 32"
	case ZydisStackWidth64:
		return "Zydis Stack Width 64"
	default:
		return fmt.Sprintf("ZydisStackWidth_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:133 -> ZydisElementType_
type ZydisElementType_ uint32

const (
	ZydisElementTypeInvalid ZydisElementType_ = 0
	ZydisElementTypeStruct ZydisElementType_ = 1
	ZydisElementTypeUint ZydisElementType_ = 2
	ZydisElementTypeInt ZydisElementType_ = 3
	ZydisElementTypeFloat16 ZydisElementType_ = 4
	ZydisElementTypeFloat32 ZydisElementType_ = 5
	ZydisElementTypeFloat64 ZydisElementType_ = 6
	ZydisElementTypeFloat80 ZydisElementType_ = 7
	ZydisElementTypeBfloat16 ZydisElementType_ = 8
	ZydisElementTypeLongbcd ZydisElementType_ = 9
	ZydisElementTypeCc ZydisElementType_ = 10
	ZydisElementTypeMaxValue ZydisElementType_ = 10
	ZydisElementTypeRequiredBits ZydisElementType_ = 4
)

func (z ZydisElementType_) String() string {
	switch z {
	case ZydisElementTypeInvalid:
		return "Zydis Element Type Invalid"
	case ZydisElementTypeStruct:
		return "Zydis Element Type Struct"
	case ZydisElementTypeUint:
		return "Zydis Element Type Uint"
	case ZydisElementTypeInt:
		return "Zydis Element Type Int"
	case ZydisElementTypeFloat16:
		return "Zydis Element Type Float 16"
	case ZydisElementTypeFloat32:
		return "Zydis Element Type Float 32"
	case ZydisElementTypeFloat64:
		return "Zydis Element Type Float 64"
	case ZydisElementTypeFloat80:
		return "Zydis Element Type Float 80"
	case ZydisElementTypeBfloat16:
		return "Zydis Element Type Bfloat 16"
	case ZydisElementTypeLongbcd:
		return "Zydis Element Type Longbcd"
	case ZydisElementTypeCc:
		return "Zydis Element Type Cc"
	default:
		return fmt.Sprintf("ZydisElementType_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:203 -> ZydisOperandType_
type ZydisOperandType_ uint32

const (
	ZydisOperandTypeUnused ZydisOperandType_ = 0
	ZydisOperandTypeRegister ZydisOperandType_ = 1
	ZydisOperandTypeMemory ZydisOperandType_ = 2
	ZydisOperandTypePointer ZydisOperandType_ = 3
	ZydisOperandTypeImmediate ZydisOperandType_ = 4
	ZydisOperandTypeMaxValue ZydisOperandType_ = 4
	ZydisOperandTypeRequiredBits ZydisOperandType_ = 3
)

func (z ZydisOperandType_) String() string {
	switch z {
	case ZydisOperandTypeUnused:
		return "Zydis Operand Type Unused"
	case ZydisOperandTypeRegister:
		return "Zydis Operand Type Register"
	case ZydisOperandTypeMemory:
		return "Zydis Operand Type Memory"
	case ZydisOperandTypePointer:
		return "Zydis Operand Type Pointer"
	case ZydisOperandTypeImmediate:
		return "Zydis Operand Type Immediate"
	default:
		return fmt.Sprintf("ZydisOperandType_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:247 -> ZydisOperandEncoding_
type ZydisOperandEncoding_ uint32

const (
	ZydisOperandEncodingNone ZydisOperandEncoding_ = 0
	ZydisOperandEncodingModrmReg ZydisOperandEncoding_ = 1
	ZydisOperandEncodingModrmRm ZydisOperandEncoding_ = 2
	ZydisOperandEncodingOpcode ZydisOperandEncoding_ = 3
	ZydisOperandEncodingNdsndd ZydisOperandEncoding_ = 4
	ZydisOperandEncodingIs4 ZydisOperandEncoding_ = 5
	ZydisOperandEncodingMask ZydisOperandEncoding_ = 6
	ZydisOperandEncodingDisp8 ZydisOperandEncoding_ = 7
	ZydisOperandEncodingDisp16 ZydisOperandEncoding_ = 8
	ZydisOperandEncodingDisp32 ZydisOperandEncoding_ = 9
	ZydisOperandEncodingDisp64 ZydisOperandEncoding_ = 10
	ZydisOperandEncodingDisp163264 ZydisOperandEncoding_ = 11
	ZydisOperandEncodingDisp323264 ZydisOperandEncoding_ = 12
	ZydisOperandEncodingDisp163232 ZydisOperandEncoding_ = 13
	ZydisOperandEncodingUimm8 ZydisOperandEncoding_ = 14
	ZydisOperandEncodingUimm16 ZydisOperandEncoding_ = 15
	ZydisOperandEncodingUimm32 ZydisOperandEncoding_ = 16
	ZydisOperandEncodingUimm64 ZydisOperandEncoding_ = 17
	ZydisOperandEncodingUimm163264 ZydisOperandEncoding_ = 18
	ZydisOperandEncodingUimm323264 ZydisOperandEncoding_ = 19
	ZydisOperandEncodingUimm163232 ZydisOperandEncoding_ = 20
	ZydisOperandEncodingSimm8 ZydisOperandEncoding_ = 21
	ZydisOperandEncodingSimm16 ZydisOperandEncoding_ = 22
	ZydisOperandEncodingSimm32 ZydisOperandEncoding_ = 23
	ZydisOperandEncodingSimm64 ZydisOperandEncoding_ = 24
	ZydisOperandEncodingSimm163264 ZydisOperandEncoding_ = 25
	ZydisOperandEncodingSimm323264 ZydisOperandEncoding_ = 26
	ZydisOperandEncodingSimm163232 ZydisOperandEncoding_ = 27
	ZydisOperandEncodingJimm8 ZydisOperandEncoding_ = 28
	ZydisOperandEncodingJimm16 ZydisOperandEncoding_ = 29
	ZydisOperandEncodingJimm32 ZydisOperandEncoding_ = 30
	ZydisOperandEncodingJimm64 ZydisOperandEncoding_ = 31
	ZydisOperandEncodingJimm163264 ZydisOperandEncoding_ = 32
	ZydisOperandEncodingJimm323264 ZydisOperandEncoding_ = 33
	ZydisOperandEncodingJimm163232 ZydisOperandEncoding_ = 34
	ZydisOperandEncodingMaxValue ZydisOperandEncoding_ = 34
	ZydisOperandEncodingRequiredBits ZydisOperandEncoding_ = 6
)

func (z ZydisOperandEncoding_) String() string {
	switch z {
	case ZydisOperandEncodingNone:
		return "Zydis Operand Encoding None"
	case ZydisOperandEncodingModrmReg:
		return "Zydis Operand Encoding Modrm Reg"
	case ZydisOperandEncodingModrmRm:
		return "Zydis Operand Encoding Modrm Rm"
	case ZydisOperandEncodingOpcode:
		return "Zydis Operand Encoding Opcode"
	case ZydisOperandEncodingNdsndd:
		return "Zydis Operand Encoding Ndsndd"
	case ZydisOperandEncodingIs4:
		return "Zydis Operand Encoding Is 4"
	case ZydisOperandEncodingMask:
		return "Zydis Operand Encoding Mask"
	case ZydisOperandEncodingDisp8:
		return "Zydis Operand Encoding Disp 8"
	case ZydisOperandEncodingDisp16:
		return "Zydis Operand Encoding Disp 16"
	case ZydisOperandEncodingDisp32:
		return "Zydis Operand Encoding Disp 32"
	case ZydisOperandEncodingDisp64:
		return "Zydis Operand Encoding Disp 64"
	case ZydisOperandEncodingDisp163264:
		return "Zydis Operand Encoding Disp 163264"
	case ZydisOperandEncodingDisp323264:
		return "Zydis Operand Encoding Disp 323264"
	case ZydisOperandEncodingDisp163232:
		return "Zydis Operand Encoding Disp 163232"
	case ZydisOperandEncodingUimm8:
		return "Zydis Operand Encoding Uimm 8"
	case ZydisOperandEncodingUimm16:
		return "Zydis Operand Encoding Uimm 16"
	case ZydisOperandEncodingUimm32:
		return "Zydis Operand Encoding Uimm 32"
	case ZydisOperandEncodingUimm64:
		return "Zydis Operand Encoding Uimm 64"
	case ZydisOperandEncodingUimm163264:
		return "Zydis Operand Encoding Uimm 163264"
	case ZydisOperandEncodingUimm323264:
		return "Zydis Operand Encoding Uimm 323264"
	case ZydisOperandEncodingUimm163232:
		return "Zydis Operand Encoding Uimm 163232"
	case ZydisOperandEncodingSimm8:
		return "Zydis Operand Encoding Simm 8"
	case ZydisOperandEncodingSimm16:
		return "Zydis Operand Encoding Simm 16"
	case ZydisOperandEncodingSimm32:
		return "Zydis Operand Encoding Simm 32"
	case ZydisOperandEncodingSimm64:
		return "Zydis Operand Encoding Simm 64"
	case ZydisOperandEncodingSimm163264:
		return "Zydis Operand Encoding Simm 163264"
	case ZydisOperandEncodingSimm323264:
		return "Zydis Operand Encoding Simm 323264"
	case ZydisOperandEncodingSimm163232:
		return "Zydis Operand Encoding Simm 163232"
	case ZydisOperandEncodingJimm8:
		return "Zydis Operand Encoding Jimm 8"
	case ZydisOperandEncodingJimm16:
		return "Zydis Operand Encoding Jimm 16"
	case ZydisOperandEncodingJimm32:
		return "Zydis Operand Encoding Jimm 32"
	case ZydisOperandEncodingJimm64:
		return "Zydis Operand Encoding Jimm 64"
	case ZydisOperandEncodingJimm163264:
		return "Zydis Operand Encoding Jimm 163264"
	case ZydisOperandEncodingJimm323264:
		return "Zydis Operand Encoding Jimm 323264"
	case ZydisOperandEncodingJimm163232:
		return "Zydis Operand Encoding Jimm 163232"
	default:
		return fmt.Sprintf("ZydisOperandEncoding_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:302 -> ZydisOperandVisibility_
type ZydisOperandVisibility_ uint32

const (
	ZydisOperandVisibilityInvalid ZydisOperandVisibility_ = 0
	ZydisOperandVisibilityExplicit ZydisOperandVisibility_ = 1
	ZydisOperandVisibilityImplicit ZydisOperandVisibility_ = 2
	ZydisOperandVisibilityHidden ZydisOperandVisibility_ = 3
	ZydisOperandVisibilityMaxValue ZydisOperandVisibility_ = 3
	ZydisOperandVisibilityRequiredBits ZydisOperandVisibility_ = 2
)

func (z ZydisOperandVisibility_) String() string {
	switch z {
	case ZydisOperandVisibilityInvalid:
		return "Zydis Operand Visibility Invalid"
	case ZydisOperandVisibilityExplicit:
		return "Zydis Operand Visibility Explicit"
	case ZydisOperandVisibilityImplicit:
		return "Zydis Operand Visibility Implicit"
	case ZydisOperandVisibilityHidden:
		return "Zydis Operand Visibility Hidden"
	default:
		return fmt.Sprintf("ZydisOperandVisibility_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:336 -> ZydisOperandAction_
type ZydisOperandAction_ uint32

const (
	ZydisOperandActionRead ZydisOperandAction_ = 1
	ZydisOperandActionWrite ZydisOperandAction_ = 2
	ZydisOperandActionCondread ZydisOperandAction_ = 4
	ZydisOperandActionCondwrite ZydisOperandAction_ = 8
	ZydisOperandActionReadwrite ZydisOperandAction_ = 3
	ZydisOperandActionCondreadCondwrite ZydisOperandAction_ = 12
	ZydisOperandActionReadCondwrite ZydisOperandAction_ = 9
	ZydisOperandActionCondreadWrite ZydisOperandAction_ = 6
	ZydisOperandActionMaskRead ZydisOperandAction_ = 5
	ZydisOperandActionMaskWrite ZydisOperandAction_ = 10
	ZydisOperandActionRequiredBits ZydisOperandAction_ = 4
)

func (z ZydisOperandAction_) String() string {
	switch z {
	case ZydisOperandActionRead:
		return "Zydis Operand Action Read"
	case ZydisOperandActionWrite:
		return "Zydis Operand Action Write"
	case ZydisOperandActionCondread:
		return "Zydis Operand Action Condread"
	case ZydisOperandActionCondwrite:
		return "Zydis Operand Action Condwrite"
	case ZydisOperandActionReadwrite:
		return "Zydis Operand Action Readwrite"
	case ZydisOperandActionCondreadCondwrite:
		return "Zydis Operand Action Condread Condwrite"
	case ZydisOperandActionReadCondwrite:
		return "Zydis Operand Action Read Condwrite"
	case ZydisOperandActionCondreadWrite:
		return "Zydis Operand Action Condread Write"
	case ZydisOperandActionMaskRead:
		return "Zydis Operand Action Mask Read"
	case ZydisOperandActionMaskWrite:
		return "Zydis Operand Action Mask Write"
	default:
		return fmt.Sprintf("ZydisOperandAction_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:415 -> ZydisInstructionEncoding_
type ZydisInstructionEncoding_ uint32

const (
	ZydisInstructionEncodingLegacy ZydisInstructionEncoding_ = 0
	ZydisInstructionEncoding3dnow ZydisInstructionEncoding_ = 1
	ZydisInstructionEncodingXop ZydisInstructionEncoding_ = 2
	ZydisInstructionEncodingVex ZydisInstructionEncoding_ = 3
	ZydisInstructionEncodingEvex ZydisInstructionEncoding_ = 4
	ZydisInstructionEncodingMvex ZydisInstructionEncoding_ = 5
	ZydisInstructionEncodingRex2 ZydisInstructionEncoding_ = 6
	ZydisInstructionEncodingMaxValue ZydisInstructionEncoding_ = 6
	ZydisInstructionEncodingRequiredBits ZydisInstructionEncoding_ = 3
)

func (z ZydisInstructionEncoding_) String() string {
	switch z {
	case ZydisInstructionEncodingLegacy:
		return "Zydis Instruction Encoding Legacy"
	case ZydisInstructionEncoding3dnow:
		return "Zydis Instruction Encoding 3dnow"
	case ZydisInstructionEncodingXop:
		return "Zydis Instruction Encoding Xop"
	case ZydisInstructionEncodingVex:
		return "Zydis Instruction Encoding Vex"
	case ZydisInstructionEncodingEvex:
		return "Zydis Instruction Encoding Evex"
	case ZydisInstructionEncodingMvex:
		return "Zydis Instruction Encoding Mvex"
	case ZydisInstructionEncodingRex2:
		return "Zydis Instruction Encoding Rex 2"
	default:
		return fmt.Sprintf("ZydisInstructionEncoding_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:464 -> ZydisOpcodeMap_
type ZydisOpcodeMap_ uint32

const (
	ZydisOpcodeMapDefault ZydisOpcodeMap_ = 0
	ZydisOpcodeMap0f ZydisOpcodeMap_ = 1
	ZydisOpcodeMap0f38 ZydisOpcodeMap_ = 2
	ZydisOpcodeMap0f3a ZydisOpcodeMap_ = 3
	ZydisOpcodeMapMap4 ZydisOpcodeMap_ = 4
	ZydisOpcodeMapMap5 ZydisOpcodeMap_ = 5
	ZydisOpcodeMapMap6 ZydisOpcodeMap_ = 6
	ZydisOpcodeMapMap7 ZydisOpcodeMap_ = 7
	ZydisOpcodeMap0f0f ZydisOpcodeMap_ = 8
	ZydisOpcodeMapXop8 ZydisOpcodeMap_ = 9
	ZydisOpcodeMapXop9 ZydisOpcodeMap_ = 10
	ZydisOpcodeMapXopa ZydisOpcodeMap_ = 11
	ZydisOpcodeMapMinValue ZydisOpcodeMap_ = 0
	ZydisOpcodeMapMaxValue ZydisOpcodeMap_ = 11
	ZydisOpcodeMapRequiredBits ZydisOpcodeMap_ = 4
)

func (z ZydisOpcodeMap_) String() string {
	switch z {
	case ZydisOpcodeMapDefault:
		return "Zydis Opcode Map Default"
	case ZydisOpcodeMap0f:
		return "Zydis Opcode Map 0f"
	case ZydisOpcodeMap0f38:
		return "Zydis Opcode Map 0f 38"
	case ZydisOpcodeMap0f3a:
		return "Zydis Opcode Map 0f 3a"
	case ZydisOpcodeMapMap4:
		return "Zydis Opcode Map Map 4"
	case ZydisOpcodeMapMap5:
		return "Zydis Opcode Map Map 5"
	case ZydisOpcodeMapMap6:
		return "Zydis Opcode Map Map 6"
	case ZydisOpcodeMapMap7:
		return "Zydis Opcode Map Map 7"
	case ZydisOpcodeMap0f0f:
		return "Zydis Opcode Map 0f 0f"
	case ZydisOpcodeMapXop8:
		return "Zydis Opcode Map Xop 8"
	case ZydisOpcodeMapXop9:
		return "Zydis Opcode Map Xop 9"
	case ZydisOpcodeMapXopa:
		return "Zydis Opcode Map Xopa"
	default:
		return fmt.Sprintf("ZydisOpcodeMap_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:4 -> ZydisRegister_
type ZydisRegister_ uint32

const (
	ZydisRegisterNone ZydisRegister_ = 0
	ZydisRegisterAl ZydisRegister_ = 1
	ZydisRegisterCl ZydisRegister_ = 2
	ZydisRegisterDl ZydisRegister_ = 3
	ZydisRegisterBl ZydisRegister_ = 4
	ZydisRegisterAh ZydisRegister_ = 5
	ZydisRegisterCh ZydisRegister_ = 6
	ZydisRegisterDh ZydisRegister_ = 7
	ZydisRegisterBh ZydisRegister_ = 8
	ZydisRegisterSpl ZydisRegister_ = 9
	ZydisRegisterBpl ZydisRegister_ = 10
	ZydisRegisterSil ZydisRegister_ = 11
	ZydisRegisterDil ZydisRegister_ = 12
	ZydisRegisterR8b ZydisRegister_ = 13
	ZydisRegisterR9b ZydisRegister_ = 14
	ZydisRegisterR10b ZydisRegister_ = 15
	ZydisRegisterR11b ZydisRegister_ = 16
	ZydisRegisterR12b ZydisRegister_ = 17
	ZydisRegisterR13b ZydisRegister_ = 18
	ZydisRegisterR14b ZydisRegister_ = 19
	ZydisRegisterR15b ZydisRegister_ = 20
	ZydisRegisterR16b ZydisRegister_ = 21
	ZydisRegisterR17b ZydisRegister_ = 22
	ZydisRegisterR18b ZydisRegister_ = 23
	ZydisRegisterR19b ZydisRegister_ = 24
	ZydisRegisterR20b ZydisRegister_ = 25
	ZydisRegisterR21b ZydisRegister_ = 26
	ZydisRegisterR22b ZydisRegister_ = 27
	ZydisRegisterR23b ZydisRegister_ = 28
	ZydisRegisterR24b ZydisRegister_ = 29
	ZydisRegisterR25b ZydisRegister_ = 30
	ZydisRegisterR26b ZydisRegister_ = 31
	ZydisRegisterR27b ZydisRegister_ = 32
	ZydisRegisterR28b ZydisRegister_ = 33
	ZydisRegisterR29b ZydisRegister_ = 34
	ZydisRegisterR30b ZydisRegister_ = 35
	ZydisRegisterR31b ZydisRegister_ = 36
	ZydisRegisterAx ZydisRegister_ = 37
	ZydisRegisterCx ZydisRegister_ = 38
	ZydisRegisterDx ZydisRegister_ = 39
	ZydisRegisterBx ZydisRegister_ = 40
	ZydisRegisterSp ZydisRegister_ = 41
	ZydisRegisterBp ZydisRegister_ = 42
	ZydisRegisterSi ZydisRegister_ = 43
	ZydisRegisterDi ZydisRegister_ = 44
	ZydisRegisterR8w ZydisRegister_ = 45
	ZydisRegisterR9w ZydisRegister_ = 46
	ZydisRegisterR10w ZydisRegister_ = 47
	ZydisRegisterR11w ZydisRegister_ = 48
	ZydisRegisterR12w ZydisRegister_ = 49
	ZydisRegisterR13w ZydisRegister_ = 50
	ZydisRegisterR14w ZydisRegister_ = 51
	ZydisRegisterR15w ZydisRegister_ = 52
	ZydisRegisterR16w ZydisRegister_ = 53
	ZydisRegisterR17w ZydisRegister_ = 54
	ZydisRegisterR18w ZydisRegister_ = 55
	ZydisRegisterR19w ZydisRegister_ = 56
	ZydisRegisterR20w ZydisRegister_ = 57
	ZydisRegisterR21w ZydisRegister_ = 58
	ZydisRegisterR22w ZydisRegister_ = 59
	ZydisRegisterR23w ZydisRegister_ = 60
	ZydisRegisterR24w ZydisRegister_ = 61
	ZydisRegisterR25w ZydisRegister_ = 62
	ZydisRegisterR26w ZydisRegister_ = 63
	ZydisRegisterR27w ZydisRegister_ = 64
	ZydisRegisterR28w ZydisRegister_ = 65
	ZydisRegisterR29w ZydisRegister_ = 66
	ZydisRegisterR30w ZydisRegister_ = 67
	ZydisRegisterR31w ZydisRegister_ = 68
	ZydisRegisterEax ZydisRegister_ = 69
	ZydisRegisterEcx ZydisRegister_ = 70
	ZydisRegisterEdx ZydisRegister_ = 71
	ZydisRegisterEbx ZydisRegister_ = 72
	ZydisRegisterEsp ZydisRegister_ = 73
	ZydisRegisterEbp ZydisRegister_ = 74
	ZydisRegisterEsi ZydisRegister_ = 75
	ZydisRegisterEdi ZydisRegister_ = 76
	ZydisRegisterR8d ZydisRegister_ = 77
	ZydisRegisterR9d ZydisRegister_ = 78
	ZydisRegisterR10d ZydisRegister_ = 79
	ZydisRegisterR11d ZydisRegister_ = 80
	ZydisRegisterR12d ZydisRegister_ = 81
	ZydisRegisterR13d ZydisRegister_ = 82
	ZydisRegisterR14d ZydisRegister_ = 83
	ZydisRegisterR15d ZydisRegister_ = 84
	ZydisRegisterR16d ZydisRegister_ = 85
	ZydisRegisterR17d ZydisRegister_ = 86
	ZydisRegisterR18d ZydisRegister_ = 87
	ZydisRegisterR19d ZydisRegister_ = 88
	ZydisRegisterR20d ZydisRegister_ = 89
	ZydisRegisterR21d ZydisRegister_ = 90
	ZydisRegisterR22d ZydisRegister_ = 91
	ZydisRegisterR23d ZydisRegister_ = 92
	ZydisRegisterR24d ZydisRegister_ = 93
	ZydisRegisterR25d ZydisRegister_ = 94
	ZydisRegisterR26d ZydisRegister_ = 95
	ZydisRegisterR27d ZydisRegister_ = 96
	ZydisRegisterR28d ZydisRegister_ = 97
	ZydisRegisterR29d ZydisRegister_ = 98
	ZydisRegisterR30d ZydisRegister_ = 99
	ZydisRegisterR31d ZydisRegister_ = 100
	ZydisRegisterRax ZydisRegister_ = 101
	ZydisRegisterRcx ZydisRegister_ = 102
	ZydisRegisterRdx ZydisRegister_ = 103
	ZydisRegisterRbx ZydisRegister_ = 104
	ZydisRegisterRsp ZydisRegister_ = 105
	ZydisRegisterRbp ZydisRegister_ = 106
	ZydisRegisterRsi ZydisRegister_ = 107
	ZydisRegisterRdi ZydisRegister_ = 108
	ZydisRegisterR8 ZydisRegister_ = 109
	ZydisRegisterR9 ZydisRegister_ = 110
	ZydisRegisterR10 ZydisRegister_ = 111
	ZydisRegisterR11 ZydisRegister_ = 112
	ZydisRegisterR12 ZydisRegister_ = 113
	ZydisRegisterR13 ZydisRegister_ = 114
	ZydisRegisterR14 ZydisRegister_ = 115
	ZydisRegisterR15 ZydisRegister_ = 116
	ZydisRegisterR16 ZydisRegister_ = 117
	ZydisRegisterR17 ZydisRegister_ = 118
	ZydisRegisterR18 ZydisRegister_ = 119
	ZydisRegisterR19 ZydisRegister_ = 120
	ZydisRegisterR20 ZydisRegister_ = 121
	ZydisRegisterR21 ZydisRegister_ = 122
	ZydisRegisterR22 ZydisRegister_ = 123
	ZydisRegisterR23 ZydisRegister_ = 124
	ZydisRegisterR24 ZydisRegister_ = 125
	ZydisRegisterR25 ZydisRegister_ = 126
	ZydisRegisterR26 ZydisRegister_ = 127
	ZydisRegisterR27 ZydisRegister_ = 128
	ZydisRegisterR28 ZydisRegister_ = 129
	ZydisRegisterR29 ZydisRegister_ = 130
	ZydisRegisterR30 ZydisRegister_ = 131
	ZydisRegisterR31 ZydisRegister_ = 132
	ZydisRegisterSt0 ZydisRegister_ = 133
	ZydisRegisterSt1 ZydisRegister_ = 134
	ZydisRegisterSt2 ZydisRegister_ = 135
	ZydisRegisterSt3 ZydisRegister_ = 136
	ZydisRegisterSt4 ZydisRegister_ = 137
	ZydisRegisterSt5 ZydisRegister_ = 138
	ZydisRegisterSt6 ZydisRegister_ = 139
	ZydisRegisterSt7 ZydisRegister_ = 140
	ZydisRegisterX87control ZydisRegister_ = 141
	ZydisRegisterX87status ZydisRegister_ = 142
	ZydisRegisterX87tag ZydisRegister_ = 143
	ZydisRegisterMm0 ZydisRegister_ = 144
	ZydisRegisterMm1 ZydisRegister_ = 145
	ZydisRegisterMm2 ZydisRegister_ = 146
	ZydisRegisterMm3 ZydisRegister_ = 147
	ZydisRegisterMm4 ZydisRegister_ = 148
	ZydisRegisterMm5 ZydisRegister_ = 149
	ZydisRegisterMm6 ZydisRegister_ = 150
	ZydisRegisterMm7 ZydisRegister_ = 151
	ZydisRegisterXmm0 ZydisRegister_ = 152
	ZydisRegisterXmm1 ZydisRegister_ = 153
	ZydisRegisterXmm2 ZydisRegister_ = 154
	ZydisRegisterXmm3 ZydisRegister_ = 155
	ZydisRegisterXmm4 ZydisRegister_ = 156
	ZydisRegisterXmm5 ZydisRegister_ = 157
	ZydisRegisterXmm6 ZydisRegister_ = 158
	ZydisRegisterXmm7 ZydisRegister_ = 159
	ZydisRegisterXmm8 ZydisRegister_ = 160
	ZydisRegisterXmm9 ZydisRegister_ = 161
	ZydisRegisterXmm10 ZydisRegister_ = 162
	ZydisRegisterXmm11 ZydisRegister_ = 163
	ZydisRegisterXmm12 ZydisRegister_ = 164
	ZydisRegisterXmm13 ZydisRegister_ = 165
	ZydisRegisterXmm14 ZydisRegister_ = 166
	ZydisRegisterXmm15 ZydisRegister_ = 167
	ZydisRegisterXmm16 ZydisRegister_ = 168
	ZydisRegisterXmm17 ZydisRegister_ = 169
	ZydisRegisterXmm18 ZydisRegister_ = 170
	ZydisRegisterXmm19 ZydisRegister_ = 171
	ZydisRegisterXmm20 ZydisRegister_ = 172
	ZydisRegisterXmm21 ZydisRegister_ = 173
	ZydisRegisterXmm22 ZydisRegister_ = 174
	ZydisRegisterXmm23 ZydisRegister_ = 175
	ZydisRegisterXmm24 ZydisRegister_ = 176
	ZydisRegisterXmm25 ZydisRegister_ = 177
	ZydisRegisterXmm26 ZydisRegister_ = 178
	ZydisRegisterXmm27 ZydisRegister_ = 179
	ZydisRegisterXmm28 ZydisRegister_ = 180
	ZydisRegisterXmm29 ZydisRegister_ = 181
	ZydisRegisterXmm30 ZydisRegister_ = 182
	ZydisRegisterXmm31 ZydisRegister_ = 183
	ZydisRegisterYmm0 ZydisRegister_ = 184
	ZydisRegisterYmm1 ZydisRegister_ = 185
	ZydisRegisterYmm2 ZydisRegister_ = 186
	ZydisRegisterYmm3 ZydisRegister_ = 187
	ZydisRegisterYmm4 ZydisRegister_ = 188
	ZydisRegisterYmm5 ZydisRegister_ = 189
	ZydisRegisterYmm6 ZydisRegister_ = 190
	ZydisRegisterYmm7 ZydisRegister_ = 191
	ZydisRegisterYmm8 ZydisRegister_ = 192
	ZydisRegisterYmm9 ZydisRegister_ = 193
	ZydisRegisterYmm10 ZydisRegister_ = 194
	ZydisRegisterYmm11 ZydisRegister_ = 195
	ZydisRegisterYmm12 ZydisRegister_ = 196
	ZydisRegisterYmm13 ZydisRegister_ = 197
	ZydisRegisterYmm14 ZydisRegister_ = 198
	ZydisRegisterYmm15 ZydisRegister_ = 199
	ZydisRegisterYmm16 ZydisRegister_ = 200
	ZydisRegisterYmm17 ZydisRegister_ = 201
	ZydisRegisterYmm18 ZydisRegister_ = 202
	ZydisRegisterYmm19 ZydisRegister_ = 203
	ZydisRegisterYmm20 ZydisRegister_ = 204
	ZydisRegisterYmm21 ZydisRegister_ = 205
	ZydisRegisterYmm22 ZydisRegister_ = 206
	ZydisRegisterYmm23 ZydisRegister_ = 207
	ZydisRegisterYmm24 ZydisRegister_ = 208
	ZydisRegisterYmm25 ZydisRegister_ = 209
	ZydisRegisterYmm26 ZydisRegister_ = 210
	ZydisRegisterYmm27 ZydisRegister_ = 211
	ZydisRegisterYmm28 ZydisRegister_ = 212
	ZydisRegisterYmm29 ZydisRegister_ = 213
	ZydisRegisterYmm30 ZydisRegister_ = 214
	ZydisRegisterYmm31 ZydisRegister_ = 215
	ZydisRegisterZmm0 ZydisRegister_ = 216
	ZydisRegisterZmm1 ZydisRegister_ = 217
	ZydisRegisterZmm2 ZydisRegister_ = 218
	ZydisRegisterZmm3 ZydisRegister_ = 219
	ZydisRegisterZmm4 ZydisRegister_ = 220
	ZydisRegisterZmm5 ZydisRegister_ = 221
	ZydisRegisterZmm6 ZydisRegister_ = 222
	ZydisRegisterZmm7 ZydisRegister_ = 223
	ZydisRegisterZmm8 ZydisRegister_ = 224
	ZydisRegisterZmm9 ZydisRegister_ = 225
	ZydisRegisterZmm10 ZydisRegister_ = 226
	ZydisRegisterZmm11 ZydisRegister_ = 227
	ZydisRegisterZmm12 ZydisRegister_ = 228
	ZydisRegisterZmm13 ZydisRegister_ = 229
	ZydisRegisterZmm14 ZydisRegister_ = 230
	ZydisRegisterZmm15 ZydisRegister_ = 231
	ZydisRegisterZmm16 ZydisRegister_ = 232
	ZydisRegisterZmm17 ZydisRegister_ = 233
	ZydisRegisterZmm18 ZydisRegister_ = 234
	ZydisRegisterZmm19 ZydisRegister_ = 235
	ZydisRegisterZmm20 ZydisRegister_ = 236
	ZydisRegisterZmm21 ZydisRegister_ = 237
	ZydisRegisterZmm22 ZydisRegister_ = 238
	ZydisRegisterZmm23 ZydisRegister_ = 239
	ZydisRegisterZmm24 ZydisRegister_ = 240
	ZydisRegisterZmm25 ZydisRegister_ = 241
	ZydisRegisterZmm26 ZydisRegister_ = 242
	ZydisRegisterZmm27 ZydisRegister_ = 243
	ZydisRegisterZmm28 ZydisRegister_ = 244
	ZydisRegisterZmm29 ZydisRegister_ = 245
	ZydisRegisterZmm30 ZydisRegister_ = 246
	ZydisRegisterZmm31 ZydisRegister_ = 247
	ZydisRegisterTmm0 ZydisRegister_ = 248
	ZydisRegisterTmm1 ZydisRegister_ = 249
	ZydisRegisterTmm2 ZydisRegister_ = 250
	ZydisRegisterTmm3 ZydisRegister_ = 251
	ZydisRegisterTmm4 ZydisRegister_ = 252
	ZydisRegisterTmm5 ZydisRegister_ = 253
	ZydisRegisterTmm6 ZydisRegister_ = 254
	ZydisRegisterTmm7 ZydisRegister_ = 255
	ZydisRegisterFlags ZydisRegister_ = 256
	ZydisRegisterEflags ZydisRegister_ = 257
	ZydisRegisterRflags ZydisRegister_ = 258
	ZydisRegisterIp ZydisRegister_ = 259
	ZydisRegisterEip ZydisRegister_ = 260
	ZydisRegisterRip ZydisRegister_ = 261
	ZydisRegisterEs ZydisRegister_ = 262
	ZydisRegisterCs ZydisRegister_ = 263
	ZydisRegisterSs ZydisRegister_ = 264
	ZydisRegisterDs ZydisRegister_ = 265
	ZydisRegisterFs ZydisRegister_ = 266
	ZydisRegisterGs ZydisRegister_ = 267
	ZydisRegisterGdtr ZydisRegister_ = 268
	ZydisRegisterLdtr ZydisRegister_ = 269
	ZydisRegisterIdtr ZydisRegister_ = 270
	ZydisRegisterTr ZydisRegister_ = 271
	ZydisRegisterTr0 ZydisRegister_ = 272
	ZydisRegisterTr1 ZydisRegister_ = 273
	ZydisRegisterTr2 ZydisRegister_ = 274
	ZydisRegisterTr3 ZydisRegister_ = 275
	ZydisRegisterTr4 ZydisRegister_ = 276
	ZydisRegisterTr5 ZydisRegister_ = 277
	ZydisRegisterTr6 ZydisRegister_ = 278
	ZydisRegisterTr7 ZydisRegister_ = 279
	ZydisRegisterCr0 ZydisRegister_ = 280
	ZydisRegisterCr1 ZydisRegister_ = 281
	ZydisRegisterCr2 ZydisRegister_ = 282
	ZydisRegisterCr3 ZydisRegister_ = 283
	ZydisRegisterCr4 ZydisRegister_ = 284
	ZydisRegisterCr5 ZydisRegister_ = 285
	ZydisRegisterCr6 ZydisRegister_ = 286
	ZydisRegisterCr7 ZydisRegister_ = 287
	ZydisRegisterCr8 ZydisRegister_ = 288
	ZydisRegisterCr9 ZydisRegister_ = 289
	ZydisRegisterCr10 ZydisRegister_ = 290
	ZydisRegisterCr11 ZydisRegister_ = 291
	ZydisRegisterCr12 ZydisRegister_ = 292
	ZydisRegisterCr13 ZydisRegister_ = 293
	ZydisRegisterCr14 ZydisRegister_ = 294
	ZydisRegisterCr15 ZydisRegister_ = 295
	ZydisRegisterDr0 ZydisRegister_ = 296
	ZydisRegisterDr1 ZydisRegister_ = 297
	ZydisRegisterDr2 ZydisRegister_ = 298
	ZydisRegisterDr3 ZydisRegister_ = 299
	ZydisRegisterDr4 ZydisRegister_ = 300
	ZydisRegisterDr5 ZydisRegister_ = 301
	ZydisRegisterDr6 ZydisRegister_ = 302
	ZydisRegisterDr7 ZydisRegister_ = 303
	ZydisRegisterDr8 ZydisRegister_ = 304
	ZydisRegisterDr9 ZydisRegister_ = 305
	ZydisRegisterDr10 ZydisRegister_ = 306
	ZydisRegisterDr11 ZydisRegister_ = 307
	ZydisRegisterDr12 ZydisRegister_ = 308
	ZydisRegisterDr13 ZydisRegister_ = 309
	ZydisRegisterDr14 ZydisRegister_ = 310
	ZydisRegisterDr15 ZydisRegister_ = 311
	ZydisRegisterK0 ZydisRegister_ = 312
	ZydisRegisterK1 ZydisRegister_ = 313
	ZydisRegisterK2 ZydisRegister_ = 314
	ZydisRegisterK3 ZydisRegister_ = 315
	ZydisRegisterK4 ZydisRegister_ = 316
	ZydisRegisterK5 ZydisRegister_ = 317
	ZydisRegisterK6 ZydisRegister_ = 318
	ZydisRegisterK7 ZydisRegister_ = 319
	ZydisRegisterBnd0 ZydisRegister_ = 320
	ZydisRegisterBnd1 ZydisRegister_ = 321
	ZydisRegisterBnd2 ZydisRegister_ = 322
	ZydisRegisterBnd3 ZydisRegister_ = 323
	ZydisRegisterBndcfg ZydisRegister_ = 324
	ZydisRegisterBndstatus ZydisRegister_ = 325
	ZydisRegisterMxcsr ZydisRegister_ = 326
	ZydisRegisterPkru ZydisRegister_ = 327
	ZydisRegisterXcr0 ZydisRegister_ = 328
	ZydisRegisterUif ZydisRegister_ = 329
	ZydisRegisterIa32KernelGsBase ZydisRegister_ = 330
	ZydisRegisterMaxValue ZydisRegister_ = 330
	ZydisRegisterRequiredBits ZydisRegister_ = 9
)

func (z ZydisRegister_) String() string {
	switch z {
	case ZydisRegisterNone:
		return "Zydis Register None"
	case ZydisRegisterAl:
		return "Zydis Register Al"
	case ZydisRegisterCl:
		return "Zydis Register Cl"
	case ZydisRegisterDl:
		return "Zydis Register Dl"
	case ZydisRegisterBl:
		return "Zydis Register Bl"
	case ZydisRegisterAh:
		return "Zydis Register Ah"
	case ZydisRegisterCh:
		return "Zydis Register Ch"
	case ZydisRegisterDh:
		return "Zydis Register Dh"
	case ZydisRegisterBh:
		return "Zydis Register Bh"
	case ZydisRegisterSpl:
		return "Zydis Register Spl"
	case ZydisRegisterBpl:
		return "Zydis Register Bpl"
	case ZydisRegisterSil:
		return "Zydis Register Sil"
	case ZydisRegisterDil:
		return "Zydis Register Dil"
	case ZydisRegisterR8b:
		return "Zydis Register R8b"
	case ZydisRegisterR9b:
		return "Zydis Register R9b"
	case ZydisRegisterR10b:
		return "Zydis Register R10b"
	case ZydisRegisterR11b:
		return "Zydis Register R11b"
	case ZydisRegisterR12b:
		return "Zydis Register R12b"
	case ZydisRegisterR13b:
		return "Zydis Register R13b"
	case ZydisRegisterR14b:
		return "Zydis Register R14b"
	case ZydisRegisterR15b:
		return "Zydis Register R15b"
	case ZydisRegisterR16b:
		return "Zydis Register R16b"
	case ZydisRegisterR17b:
		return "Zydis Register R17b"
	case ZydisRegisterR18b:
		return "Zydis Register R18b"
	case ZydisRegisterR19b:
		return "Zydis Register R19b"
	case ZydisRegisterR20b:
		return "Zydis Register R20b"
	case ZydisRegisterR21b:
		return "Zydis Register R21b"
	case ZydisRegisterR22b:
		return "Zydis Register R22b"
	case ZydisRegisterR23b:
		return "Zydis Register R23b"
	case ZydisRegisterR24b:
		return "Zydis Register R24b"
	case ZydisRegisterR25b:
		return "Zydis Register R25b"
	case ZydisRegisterR26b:
		return "Zydis Register R26b"
	case ZydisRegisterR27b:
		return "Zydis Register R27b"
	case ZydisRegisterR28b:
		return "Zydis Register R28b"
	case ZydisRegisterR29b:
		return "Zydis Register R29b"
	case ZydisRegisterR30b:
		return "Zydis Register R30b"
	case ZydisRegisterR31b:
		return "Zydis Register R31b"
	case ZydisRegisterAx:
		return "Zydis Register Ax"
	case ZydisRegisterCx:
		return "Zydis Register Cx"
	case ZydisRegisterDx:
		return "Zydis Register Dx"
	case ZydisRegisterBx:
		return "Zydis Register Bx"
	case ZydisRegisterSp:
		return "Zydis Register Sp"
	case ZydisRegisterBp:
		return "Zydis Register Bp"
	case ZydisRegisterSi:
		return "Zydis Register Si"
	case ZydisRegisterDi:
		return "Zydis Register Di"
	case ZydisRegisterR8w:
		return "Zydis Register R8w"
	case ZydisRegisterR9w:
		return "Zydis Register R9w"
	case ZydisRegisterR10w:
		return "Zydis Register R10w"
	case ZydisRegisterR11w:
		return "Zydis Register R11w"
	case ZydisRegisterR12w:
		return "Zydis Register R12w"
	case ZydisRegisterR13w:
		return "Zydis Register R13w"
	case ZydisRegisterR14w:
		return "Zydis Register R14w"
	case ZydisRegisterR15w:
		return "Zydis Register R15w"
	case ZydisRegisterR16w:
		return "Zydis Register R16w"
	case ZydisRegisterR17w:
		return "Zydis Register R17w"
	case ZydisRegisterR18w:
		return "Zydis Register R18w"
	case ZydisRegisterR19w:
		return "Zydis Register R19w"
	case ZydisRegisterR20w:
		return "Zydis Register R20w"
	case ZydisRegisterR21w:
		return "Zydis Register R21w"
	case ZydisRegisterR22w:
		return "Zydis Register R22w"
	case ZydisRegisterR23w:
		return "Zydis Register R23w"
	case ZydisRegisterR24w:
		return "Zydis Register R24w"
	case ZydisRegisterR25w:
		return "Zydis Register R25w"
	case ZydisRegisterR26w:
		return "Zydis Register R26w"
	case ZydisRegisterR27w:
		return "Zydis Register R27w"
	case ZydisRegisterR28w:
		return "Zydis Register R28w"
	case ZydisRegisterR29w:
		return "Zydis Register R29w"
	case ZydisRegisterR30w:
		return "Zydis Register R30w"
	case ZydisRegisterR31w:
		return "Zydis Register R31w"
	case ZydisRegisterEax:
		return "Zydis Register Eax"
	case ZydisRegisterEcx:
		return "Zydis Register Ecx"
	case ZydisRegisterEdx:
		return "Zydis Register Edx"
	case ZydisRegisterEbx:
		return "Zydis Register Ebx"
	case ZydisRegisterEsp:
		return "Zydis Register Esp"
	case ZydisRegisterEbp:
		return "Zydis Register Ebp"
	case ZydisRegisterEsi:
		return "Zydis Register Esi"
	case ZydisRegisterEdi:
		return "Zydis Register Edi"
	case ZydisRegisterR8d:
		return "Zydis Register R8d"
	case ZydisRegisterR9d:
		return "Zydis Register R9d"
	case ZydisRegisterR10d:
		return "Zydis Register R10d"
	case ZydisRegisterR11d:
		return "Zydis Register R11d"
	case ZydisRegisterR12d:
		return "Zydis Register R12d"
	case ZydisRegisterR13d:
		return "Zydis Register R13d"
	case ZydisRegisterR14d:
		return "Zydis Register R14d"
	case ZydisRegisterR15d:
		return "Zydis Register R15d"
	case ZydisRegisterR16d:
		return "Zydis Register R16d"
	case ZydisRegisterR17d:
		return "Zydis Register R17d"
	case ZydisRegisterR18d:
		return "Zydis Register R18d"
	case ZydisRegisterR19d:
		return "Zydis Register R19d"
	case ZydisRegisterR20d:
		return "Zydis Register R20d"
	case ZydisRegisterR21d:
		return "Zydis Register R21d"
	case ZydisRegisterR22d:
		return "Zydis Register R22d"
	case ZydisRegisterR23d:
		return "Zydis Register R23d"
	case ZydisRegisterR24d:
		return "Zydis Register R24d"
	case ZydisRegisterR25d:
		return "Zydis Register R25d"
	case ZydisRegisterR26d:
		return "Zydis Register R26d"
	case ZydisRegisterR27d:
		return "Zydis Register R27d"
	case ZydisRegisterR28d:
		return "Zydis Register R28d"
	case ZydisRegisterR29d:
		return "Zydis Register R29d"
	case ZydisRegisterR30d:
		return "Zydis Register R30d"
	case ZydisRegisterR31d:
		return "Zydis Register R31d"
	case ZydisRegisterRax:
		return "Zydis Register Rax"
	case ZydisRegisterRcx:
		return "Zydis Register Rcx"
	case ZydisRegisterRdx:
		return "Zydis Register Rdx"
	case ZydisRegisterRbx:
		return "Zydis Register Rbx"
	case ZydisRegisterRsp:
		return "Zydis Register Rsp"
	case ZydisRegisterRbp:
		return "Zydis Register Rbp"
	case ZydisRegisterRsi:
		return "Zydis Register Rsi"
	case ZydisRegisterRdi:
		return "Zydis Register Rdi"
	case ZydisRegisterR8:
		return "Zydis Register R8"
	case ZydisRegisterR9:
		return "Zydis Register R9"
	case ZydisRegisterR10:
		return "Zydis Register R10"
	case ZydisRegisterR11:
		return "Zydis Register R11"
	case ZydisRegisterR12:
		return "Zydis Register R12"
	case ZydisRegisterR13:
		return "Zydis Register R13"
	case ZydisRegisterR14:
		return "Zydis Register R14"
	case ZydisRegisterR15:
		return "Zydis Register R15"
	case ZydisRegisterR16:
		return "Zydis Register R16"
	case ZydisRegisterR17:
		return "Zydis Register R17"
	case ZydisRegisterR18:
		return "Zydis Register R18"
	case ZydisRegisterR19:
		return "Zydis Register R19"
	case ZydisRegisterR20:
		return "Zydis Register R20"
	case ZydisRegisterR21:
		return "Zydis Register R21"
	case ZydisRegisterR22:
		return "Zydis Register R22"
	case ZydisRegisterR23:
		return "Zydis Register R23"
	case ZydisRegisterR24:
		return "Zydis Register R24"
	case ZydisRegisterR25:
		return "Zydis Register R25"
	case ZydisRegisterR26:
		return "Zydis Register R26"
	case ZydisRegisterR27:
		return "Zydis Register R27"
	case ZydisRegisterR28:
		return "Zydis Register R28"
	case ZydisRegisterR29:
		return "Zydis Register R29"
	case ZydisRegisterR30:
		return "Zydis Register R30"
	case ZydisRegisterR31:
		return "Zydis Register R31"
	case ZydisRegisterSt0:
		return "Zydis Register St 0"
	case ZydisRegisterSt1:
		return "Zydis Register St 1"
	case ZydisRegisterSt2:
		return "Zydis Register St 2"
	case ZydisRegisterSt3:
		return "Zydis Register St 3"
	case ZydisRegisterSt4:
		return "Zydis Register St 4"
	case ZydisRegisterSt5:
		return "Zydis Register St 5"
	case ZydisRegisterSt6:
		return "Zydis Register St 6"
	case ZydisRegisterSt7:
		return "Zydis Register St 7"
	case ZydisRegisterX87control:
		return "Zydis Register X87control"
	case ZydisRegisterX87status:
		return "Zydis Register X87status"
	case ZydisRegisterX87tag:
		return "Zydis Register X87tag"
	case ZydisRegisterMm0:
		return "Zydis Register Mm 0"
	case ZydisRegisterMm1:
		return "Zydis Register Mm 1"
	case ZydisRegisterMm2:
		return "Zydis Register Mm 2"
	case ZydisRegisterMm3:
		return "Zydis Register Mm 3"
	case ZydisRegisterMm4:
		return "Zydis Register Mm 4"
	case ZydisRegisterMm5:
		return "Zydis Register Mm 5"
	case ZydisRegisterMm6:
		return "Zydis Register Mm 6"
	case ZydisRegisterMm7:
		return "Zydis Register Mm 7"
	case ZydisRegisterXmm0:
		return "Zydis Register Xmm 0"
	case ZydisRegisterXmm1:
		return "Zydis Register Xmm 1"
	case ZydisRegisterXmm2:
		return "Zydis Register Xmm 2"
	case ZydisRegisterXmm3:
		return "Zydis Register Xmm 3"
	case ZydisRegisterXmm4:
		return "Zydis Register Xmm 4"
	case ZydisRegisterXmm5:
		return "Zydis Register Xmm 5"
	case ZydisRegisterXmm6:
		return "Zydis Register Xmm 6"
	case ZydisRegisterXmm7:
		return "Zydis Register Xmm 7"
	case ZydisRegisterXmm8:
		return "Zydis Register Xmm 8"
	case ZydisRegisterXmm9:
		return "Zydis Register Xmm 9"
	case ZydisRegisterXmm10:
		return "Zydis Register Xmm 10"
	case ZydisRegisterXmm11:
		return "Zydis Register Xmm 11"
	case ZydisRegisterXmm12:
		return "Zydis Register Xmm 12"
	case ZydisRegisterXmm13:
		return "Zydis Register Xmm 13"
	case ZydisRegisterXmm14:
		return "Zydis Register Xmm 14"
	case ZydisRegisterXmm15:
		return "Zydis Register Xmm 15"
	case ZydisRegisterXmm16:
		return "Zydis Register Xmm 16"
	case ZydisRegisterXmm17:
		return "Zydis Register Xmm 17"
	case ZydisRegisterXmm18:
		return "Zydis Register Xmm 18"
	case ZydisRegisterXmm19:
		return "Zydis Register Xmm 19"
	case ZydisRegisterXmm20:
		return "Zydis Register Xmm 20"
	case ZydisRegisterXmm21:
		return "Zydis Register Xmm 21"
	case ZydisRegisterXmm22:
		return "Zydis Register Xmm 22"
	case ZydisRegisterXmm23:
		return "Zydis Register Xmm 23"
	case ZydisRegisterXmm24:
		return "Zydis Register Xmm 24"
	case ZydisRegisterXmm25:
		return "Zydis Register Xmm 25"
	case ZydisRegisterXmm26:
		return "Zydis Register Xmm 26"
	case ZydisRegisterXmm27:
		return "Zydis Register Xmm 27"
	case ZydisRegisterXmm28:
		return "Zydis Register Xmm 28"
	case ZydisRegisterXmm29:
		return "Zydis Register Xmm 29"
	case ZydisRegisterXmm30:
		return "Zydis Register Xmm 30"
	case ZydisRegisterXmm31:
		return "Zydis Register Xmm 31"
	case ZydisRegisterYmm0:
		return "Zydis Register Ymm 0"
	case ZydisRegisterYmm1:
		return "Zydis Register Ymm 1"
	case ZydisRegisterYmm2:
		return "Zydis Register Ymm 2"
	case ZydisRegisterYmm3:
		return "Zydis Register Ymm 3"
	case ZydisRegisterYmm4:
		return "Zydis Register Ymm 4"
	case ZydisRegisterYmm5:
		return "Zydis Register Ymm 5"
	case ZydisRegisterYmm6:
		return "Zydis Register Ymm 6"
	case ZydisRegisterYmm7:
		return "Zydis Register Ymm 7"
	case ZydisRegisterYmm8:
		return "Zydis Register Ymm 8"
	case ZydisRegisterYmm9:
		return "Zydis Register Ymm 9"
	case ZydisRegisterYmm10:
		return "Zydis Register Ymm 10"
	case ZydisRegisterYmm11:
		return "Zydis Register Ymm 11"
	case ZydisRegisterYmm12:
		return "Zydis Register Ymm 12"
	case ZydisRegisterYmm13:
		return "Zydis Register Ymm 13"
	case ZydisRegisterYmm14:
		return "Zydis Register Ymm 14"
	case ZydisRegisterYmm15:
		return "Zydis Register Ymm 15"
	case ZydisRegisterYmm16:
		return "Zydis Register Ymm 16"
	case ZydisRegisterYmm17:
		return "Zydis Register Ymm 17"
	case ZydisRegisterYmm18:
		return "Zydis Register Ymm 18"
	case ZydisRegisterYmm19:
		return "Zydis Register Ymm 19"
	case ZydisRegisterYmm20:
		return "Zydis Register Ymm 20"
	case ZydisRegisterYmm21:
		return "Zydis Register Ymm 21"
	case ZydisRegisterYmm22:
		return "Zydis Register Ymm 22"
	case ZydisRegisterYmm23:
		return "Zydis Register Ymm 23"
	case ZydisRegisterYmm24:
		return "Zydis Register Ymm 24"
	case ZydisRegisterYmm25:
		return "Zydis Register Ymm 25"
	case ZydisRegisterYmm26:
		return "Zydis Register Ymm 26"
	case ZydisRegisterYmm27:
		return "Zydis Register Ymm 27"
	case ZydisRegisterYmm28:
		return "Zydis Register Ymm 28"
	case ZydisRegisterYmm29:
		return "Zydis Register Ymm 29"
	case ZydisRegisterYmm30:
		return "Zydis Register Ymm 30"
	case ZydisRegisterYmm31:
		return "Zydis Register Ymm 31"
	case ZydisRegisterZmm0:
		return "Zydis Register Zmm 0"
	case ZydisRegisterZmm1:
		return "Zydis Register Zmm 1"
	case ZydisRegisterZmm2:
		return "Zydis Register Zmm 2"
	case ZydisRegisterZmm3:
		return "Zydis Register Zmm 3"
	case ZydisRegisterZmm4:
		return "Zydis Register Zmm 4"
	case ZydisRegisterZmm5:
		return "Zydis Register Zmm 5"
	case ZydisRegisterZmm6:
		return "Zydis Register Zmm 6"
	case ZydisRegisterZmm7:
		return "Zydis Register Zmm 7"
	case ZydisRegisterZmm8:
		return "Zydis Register Zmm 8"
	case ZydisRegisterZmm9:
		return "Zydis Register Zmm 9"
	case ZydisRegisterZmm10:
		return "Zydis Register Zmm 10"
	case ZydisRegisterZmm11:
		return "Zydis Register Zmm 11"
	case ZydisRegisterZmm12:
		return "Zydis Register Zmm 12"
	case ZydisRegisterZmm13:
		return "Zydis Register Zmm 13"
	case ZydisRegisterZmm14:
		return "Zydis Register Zmm 14"
	case ZydisRegisterZmm15:
		return "Zydis Register Zmm 15"
	case ZydisRegisterZmm16:
		return "Zydis Register Zmm 16"
	case ZydisRegisterZmm17:
		return "Zydis Register Zmm 17"
	case ZydisRegisterZmm18:
		return "Zydis Register Zmm 18"
	case ZydisRegisterZmm19:
		return "Zydis Register Zmm 19"
	case ZydisRegisterZmm20:
		return "Zydis Register Zmm 20"
	case ZydisRegisterZmm21:
		return "Zydis Register Zmm 21"
	case ZydisRegisterZmm22:
		return "Zydis Register Zmm 22"
	case ZydisRegisterZmm23:
		return "Zydis Register Zmm 23"
	case ZydisRegisterZmm24:
		return "Zydis Register Zmm 24"
	case ZydisRegisterZmm25:
		return "Zydis Register Zmm 25"
	case ZydisRegisterZmm26:
		return "Zydis Register Zmm 26"
	case ZydisRegisterZmm27:
		return "Zydis Register Zmm 27"
	case ZydisRegisterZmm28:
		return "Zydis Register Zmm 28"
	case ZydisRegisterZmm29:
		return "Zydis Register Zmm 29"
	case ZydisRegisterZmm30:
		return "Zydis Register Zmm 30"
	case ZydisRegisterZmm31:
		return "Zydis Register Zmm 31"
	case ZydisRegisterTmm0:
		return "Zydis Register Tmm 0"
	case ZydisRegisterTmm1:
		return "Zydis Register Tmm 1"
	case ZydisRegisterTmm2:
		return "Zydis Register Tmm 2"
	case ZydisRegisterTmm3:
		return "Zydis Register Tmm 3"
	case ZydisRegisterTmm4:
		return "Zydis Register Tmm 4"
	case ZydisRegisterTmm5:
		return "Zydis Register Tmm 5"
	case ZydisRegisterTmm6:
		return "Zydis Register Tmm 6"
	case ZydisRegisterTmm7:
		return "Zydis Register Tmm 7"
	case ZydisRegisterFlags:
		return "Zydis Register Flags"
	case ZydisRegisterEflags:
		return "Zydis Register Eflags"
	case ZydisRegisterRflags:
		return "Zydis Register Rflags"
	case ZydisRegisterIp:
		return "Zydis Register Ip"
	case ZydisRegisterEip:
		return "Zydis Register Eip"
	case ZydisRegisterRip:
		return "Zydis Register Rip"
	case ZydisRegisterEs:
		return "Zydis Register Es"
	case ZydisRegisterCs:
		return "Zydis Register Cs"
	case ZydisRegisterSs:
		return "Zydis Register Ss"
	case ZydisRegisterDs:
		return "Zydis Register Ds"
	case ZydisRegisterFs:
		return "Zydis Register Fs"
	case ZydisRegisterGs:
		return "Zydis Register Gs"
	case ZydisRegisterGdtr:
		return "Zydis Register Gdtr"
	case ZydisRegisterLdtr:
		return "Zydis Register Ldtr"
	case ZydisRegisterIdtr:
		return "Zydis Register Idtr"
	case ZydisRegisterTr:
		return "Zydis Register Tr"
	case ZydisRegisterTr0:
		return "Zydis Register Tr 0"
	case ZydisRegisterTr1:
		return "Zydis Register Tr 1"
	case ZydisRegisterTr2:
		return "Zydis Register Tr 2"
	case ZydisRegisterTr3:
		return "Zydis Register Tr 3"
	case ZydisRegisterTr4:
		return "Zydis Register Tr 4"
	case ZydisRegisterTr5:
		return "Zydis Register Tr 5"
	case ZydisRegisterTr6:
		return "Zydis Register Tr 6"
	case ZydisRegisterTr7:
		return "Zydis Register Tr 7"
	case ZydisRegisterCr0:
		return "Zydis Register Cr 0"
	case ZydisRegisterCr1:
		return "Zydis Register Cr 1"
	case ZydisRegisterCr2:
		return "Zydis Register Cr 2"
	case ZydisRegisterCr3:
		return "Zydis Register Cr 3"
	case ZydisRegisterCr4:
		return "Zydis Register Cr 4"
	case ZydisRegisterCr5:
		return "Zydis Register Cr 5"
	case ZydisRegisterCr6:
		return "Zydis Register Cr 6"
	case ZydisRegisterCr7:
		return "Zydis Register Cr 7"
	case ZydisRegisterCr8:
		return "Zydis Register Cr 8"
	case ZydisRegisterCr9:
		return "Zydis Register Cr 9"
	case ZydisRegisterCr10:
		return "Zydis Register Cr 10"
	case ZydisRegisterCr11:
		return "Zydis Register Cr 11"
	case ZydisRegisterCr12:
		return "Zydis Register Cr 12"
	case ZydisRegisterCr13:
		return "Zydis Register Cr 13"
	case ZydisRegisterCr14:
		return "Zydis Register Cr 14"
	case ZydisRegisterCr15:
		return "Zydis Register Cr 15"
	case ZydisRegisterDr0:
		return "Zydis Register Dr 0"
	case ZydisRegisterDr1:
		return "Zydis Register Dr 1"
	case ZydisRegisterDr2:
		return "Zydis Register Dr 2"
	case ZydisRegisterDr3:
		return "Zydis Register Dr 3"
	case ZydisRegisterDr4:
		return "Zydis Register Dr 4"
	case ZydisRegisterDr5:
		return "Zydis Register Dr 5"
	case ZydisRegisterDr6:
		return "Zydis Register Dr 6"
	case ZydisRegisterDr7:
		return "Zydis Register Dr 7"
	case ZydisRegisterDr8:
		return "Zydis Register Dr 8"
	case ZydisRegisterDr9:
		return "Zydis Register Dr 9"
	case ZydisRegisterDr10:
		return "Zydis Register Dr 10"
	case ZydisRegisterDr11:
		return "Zydis Register Dr 11"
	case ZydisRegisterDr12:
		return "Zydis Register Dr 12"
	case ZydisRegisterDr13:
		return "Zydis Register Dr 13"
	case ZydisRegisterDr14:
		return "Zydis Register Dr 14"
	case ZydisRegisterDr15:
		return "Zydis Register Dr 15"
	case ZydisRegisterK0:
		return "Zydis Register K0"
	case ZydisRegisterK1:
		return "Zydis Register K1"
	case ZydisRegisterK2:
		return "Zydis Register K2"
	case ZydisRegisterK3:
		return "Zydis Register K3"
	case ZydisRegisterK4:
		return "Zydis Register K4"
	case ZydisRegisterK5:
		return "Zydis Register K5"
	case ZydisRegisterK6:
		return "Zydis Register K6"
	case ZydisRegisterK7:
		return "Zydis Register K7"
	case ZydisRegisterBnd0:
		return "Zydis Register Bnd 0"
	case ZydisRegisterBnd1:
		return "Zydis Register Bnd 1"
	case ZydisRegisterBnd2:
		return "Zydis Register Bnd 2"
	case ZydisRegisterBnd3:
		return "Zydis Register Bnd 3"
	case ZydisRegisterBndcfg:
		return "Zydis Register Bndcfg"
	case ZydisRegisterBndstatus:
		return "Zydis Register Bndstatus"
	case ZydisRegisterMxcsr:
		return "Zydis Register Mxcsr"
	case ZydisRegisterPkru:
		return "Zydis Register Pkru"
	case ZydisRegisterXcr0:
		return "Zydis Register Xcr 0"
	case ZydisRegisterUif:
		return "Zydis Register Uif"
	case ZydisRegisterIa32KernelGsBase:
		return "Zydis Register Ia 32 Kernel Gs Base"
	default:
		return fmt.Sprintf("ZydisRegister_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:69 -> ZydisRegisterKind_
type ZydisRegisterKind_ uint32

const (
	ZydisRegkindInvalid ZydisRegisterKind_ = 0
	ZydisRegkindGpr ZydisRegisterKind_ = 1
	ZydisRegkindX87 ZydisRegisterKind_ = 2
	ZydisRegkindMmx ZydisRegisterKind_ = 3
	ZydisRegkindVr ZydisRegisterKind_ = 4
	ZydisRegkindTmm ZydisRegisterKind_ = 5
	ZydisRegkindSegment ZydisRegisterKind_ = 6
	ZydisRegkindTest ZydisRegisterKind_ = 7
	ZydisRegkindControl ZydisRegisterKind_ = 8
	ZydisRegkindDebug ZydisRegisterKind_ = 9
	ZydisRegkindMask ZydisRegisterKind_ = 10
	ZydisRegkindBound ZydisRegisterKind_ = 11
	ZydisRegkindMaxValue ZydisRegisterKind_ = 11
	ZydisRegkindRequiredBits ZydisRegisterKind_ = 4
)

func (z ZydisRegisterKind_) String() string {
	switch z {
	case ZydisRegkindInvalid:
		return "Zydis Regkind Invalid"
	case ZydisRegkindGpr:
		return "Zydis Regkind Gpr"
	case ZydisRegkindX87:
		return "Zydis Regkind X87"
	case ZydisRegkindMmx:
		return "Zydis Regkind Mmx"
	case ZydisRegkindVr:
		return "Zydis Regkind Vr"
	case ZydisRegkindTmm:
		return "Zydis Regkind Tmm"
	case ZydisRegkindSegment:
		return "Zydis Regkind Segment"
	case ZydisRegkindTest:
		return "Zydis Regkind Test"
	case ZydisRegkindControl:
		return "Zydis Regkind Control"
	case ZydisRegkindDebug:
		return "Zydis Regkind Debug"
	case ZydisRegkindMask:
		return "Zydis Regkind Mask"
	case ZydisRegkindBound:
		return "Zydis Regkind Bound"
	default:
		return fmt.Sprintf("ZydisRegisterKind_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:108 -> ZydisRegisterClass_
type ZydisRegisterClass_ uint32

const (
	ZydisRegclassInvalid ZydisRegisterClass_ = 0
	ZydisRegclassGpr8 ZydisRegisterClass_ = 1
	ZydisRegclassGpr16 ZydisRegisterClass_ = 2
	ZydisRegclassGpr32 ZydisRegisterClass_ = 3
	ZydisRegclassGpr64 ZydisRegisterClass_ = 4
	ZydisRegclassX87 ZydisRegisterClass_ = 5
	ZydisRegclassMmx ZydisRegisterClass_ = 6
	ZydisRegclassXmm ZydisRegisterClass_ = 7
	ZydisRegclassYmm ZydisRegisterClass_ = 8
	ZydisRegclassZmm ZydisRegisterClass_ = 9
	ZydisRegclassTmm ZydisRegisterClass_ = 10
	ZydisRegclassFlags ZydisRegisterClass_ = 11
	ZydisRegclassIp ZydisRegisterClass_ = 12
	ZydisRegclassSegment ZydisRegisterClass_ = 13
	ZydisRegclassTable ZydisRegisterClass_ = 14
	ZydisRegclassTest ZydisRegisterClass_ = 15
	ZydisRegclassControl ZydisRegisterClass_ = 16
	ZydisRegclassDebug ZydisRegisterClass_ = 17
	ZydisRegclassMask ZydisRegisterClass_ = 18
	ZydisRegclassBound ZydisRegisterClass_ = 19
	ZydisRegclassMaxValue ZydisRegisterClass_ = 19
	ZydisRegclassRequiredBits ZydisRegisterClass_ = 5
)

func (z ZydisRegisterClass_) String() string {
	switch z {
	case ZydisRegclassInvalid:
		return "Zydis Regclass Invalid"
	case ZydisRegclassGpr8:
		return "Zydis Regclass Gpr 8"
	case ZydisRegclassGpr16:
		return "Zydis Regclass Gpr 16"
	case ZydisRegclassGpr32:
		return "Zydis Regclass Gpr 32"
	case ZydisRegclassGpr64:
		return "Zydis Regclass Gpr 64"
	case ZydisRegclassX87:
		return "Zydis Regclass X87"
	case ZydisRegclassMmx:
		return "Zydis Regclass Mmx"
	case ZydisRegclassXmm:
		return "Zydis Regclass Xmm"
	case ZydisRegclassYmm:
		return "Zydis Regclass Ymm"
	case ZydisRegclassZmm:
		return "Zydis Regclass Zmm"
	case ZydisRegclassTmm:
		return "Zydis Regclass Tmm"
	case ZydisRegclassFlags:
		return "Zydis Regclass Flags"
	case ZydisRegclassIp:
		return "Zydis Regclass Ip"
	case ZydisRegclassSegment:
		return "Zydis Regclass Segment"
	case ZydisRegclassTable:
		return "Zydis Regclass Table"
	case ZydisRegclassTest:
		return "Zydis Regclass Test"
	case ZydisRegclassControl:
		return "Zydis Regclass Control"
	case ZydisRegclassDebug:
		return "Zydis Regclass Debug"
	case ZydisRegclassMask:
		return "Zydis Regclass Mask"
	case ZydisRegclassBound:
		return "Zydis Regclass Bound"
	default:
		return fmt.Sprintf("ZydisRegisterClass_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:75 -> ZydisMemoryOperandType_
type ZydisMemoryOperandType_ uint32

const (
	ZydisMemopTypeInvalid ZydisMemoryOperandType_ = 0
	ZydisMemopTypeMem ZydisMemoryOperandType_ = 1
	ZydisMemopTypeAgen ZydisMemoryOperandType_ = 2
	ZydisMemopTypeMib ZydisMemoryOperandType_ = 3
	ZydisMemopTypeVsib ZydisMemoryOperandType_ = 4
	ZydisMemopTypeMaxValue ZydisMemoryOperandType_ = 4
	ZydisMemopTypeRequiredBits ZydisMemoryOperandType_ = 3
)

func (z ZydisMemoryOperandType_) String() string {
	switch z {
	case ZydisMemopTypeInvalid:
		return "Zydis Memop Type Invalid"
	case ZydisMemopTypeMem:
		return "Zydis Memop Type Mem"
	case ZydisMemopTypeAgen:
		return "Zydis Memop Type Agen"
	case ZydisMemopTypeMib:
		return "Zydis Memop Type Mib"
	case ZydisMemopTypeVsib:
		return "Zydis Memop Type Vsib"
	default:
		return fmt.Sprintf("ZydisMemoryOperandType_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:434 -> ZydisBranchType_
type ZydisBranchType_ uint32

const (
	ZydisBranchTypeNone ZydisBranchType_ = 0
	ZydisBranchTypeShort ZydisBranchType_ = 1
	ZydisBranchTypeNear ZydisBranchType_ = 2
	ZydisBranchTypeFar ZydisBranchType_ = 3
	ZydisBranchTypeAbsolute ZydisBranchType_ = 4
	ZydisBranchTypeMaxValue ZydisBranchType_ = 4
	ZydisBranchTypeRequiredBits ZydisBranchType_ = 3
)

func (z ZydisBranchType_) String() string {
	switch z {
	case ZydisBranchTypeNone:
		return "Zydis Branch Type None"
	case ZydisBranchTypeShort:
		return "Zydis Branch Type Short"
	case ZydisBranchTypeNear:
		return "Zydis Branch Type Near"
	case ZydisBranchTypeFar:
		return "Zydis Branch Type Far"
	case ZydisBranchTypeAbsolute:
		return "Zydis Branch Type Absolute"
	default:
		return fmt.Sprintf("ZydisBranchType_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:474 -> ZydisExceptionClass_
type ZydisExceptionClass_ uint32

const (
	ZydisExceptionClassNone ZydisExceptionClass_ = 0
	ZydisExceptionClassSse1 ZydisExceptionClass_ = 1
	ZydisExceptionClassSse2 ZydisExceptionClass_ = 2
	ZydisExceptionClassSse3 ZydisExceptionClass_ = 3
	ZydisExceptionClassSse4 ZydisExceptionClass_ = 4
	ZydisExceptionClassSse5 ZydisExceptionClass_ = 5
	ZydisExceptionClassSse7 ZydisExceptionClass_ = 6
	ZydisExceptionClassAvx1 ZydisExceptionClass_ = 7
	ZydisExceptionClassAvx2 ZydisExceptionClass_ = 8
	ZydisExceptionClassAvx3 ZydisExceptionClass_ = 9
	ZydisExceptionClassAvx4 ZydisExceptionClass_ = 10
	ZydisExceptionClassAvx5 ZydisExceptionClass_ = 11
	ZydisExceptionClassAvx6 ZydisExceptionClass_ = 12
	ZydisExceptionClassAvx7 ZydisExceptionClass_ = 13
	ZydisExceptionClassAvx8 ZydisExceptionClass_ = 14
	ZydisExceptionClassAvx11 ZydisExceptionClass_ = 15
	ZydisExceptionClassAvx12 ZydisExceptionClass_ = 16
	ZydisExceptionClassAvx14 ZydisExceptionClass_ = 17
	ZydisExceptionClassE1 ZydisExceptionClass_ = 18
	ZydisExceptionClassE1nf ZydisExceptionClass_ = 19
	ZydisExceptionClassE2 ZydisExceptionClass_ = 20
	ZydisExceptionClassE2nf ZydisExceptionClass_ = 21
	ZydisExceptionClassE3 ZydisExceptionClass_ = 22
	ZydisExceptionClassE3nf ZydisExceptionClass_ = 23
	ZydisExceptionClassE4 ZydisExceptionClass_ = 24
	ZydisExceptionClassE4nf ZydisExceptionClass_ = 25
	ZydisExceptionClassE5 ZydisExceptionClass_ = 26
	ZydisExceptionClassE5nf ZydisExceptionClass_ = 27
	ZydisExceptionClassE6 ZydisExceptionClass_ = 28
	ZydisExceptionClassE6nf ZydisExceptionClass_ = 29
	ZydisExceptionClassE7nm ZydisExceptionClass_ = 30
	ZydisExceptionClassE7nm128 ZydisExceptionClass_ = 31
	ZydisExceptionClassE9nf ZydisExceptionClass_ = 32
	ZydisExceptionClassE10 ZydisExceptionClass_ = 33
	ZydisExceptionClassE10nf ZydisExceptionClass_ = 34
	ZydisExceptionClassE11 ZydisExceptionClass_ = 35
	ZydisExceptionClassE11nf ZydisExceptionClass_ = 36
	ZydisExceptionClassE12 ZydisExceptionClass_ = 37
	ZydisExceptionClassE12np ZydisExceptionClass_ = 38
	ZydisExceptionClassK20 ZydisExceptionClass_ = 39
	ZydisExceptionClassK21 ZydisExceptionClass_ = 40
	ZydisExceptionClassAmxe1 ZydisExceptionClass_ = 41
	ZydisExceptionClassAmxe2 ZydisExceptionClass_ = 42
	ZydisExceptionClassAmxe3 ZydisExceptionClass_ = 43
	ZydisExceptionClassAmxe4 ZydisExceptionClass_ = 44
	ZydisExceptionClassAmxe5 ZydisExceptionClass_ = 45
	ZydisExceptionClassAmxe6 ZydisExceptionClass_ = 46
	ZydisExceptionClassAmxe1Evex ZydisExceptionClass_ = 47
	ZydisExceptionClassAmxe2Evex ZydisExceptionClass_ = 48
	ZydisExceptionClassAmxe3Evex ZydisExceptionClass_ = 49
	ZydisExceptionClassAmxe4Evex ZydisExceptionClass_ = 50
	ZydisExceptionClassAmxe5Evex ZydisExceptionClass_ = 51
	ZydisExceptionClassAmxe6Evex ZydisExceptionClass_ = 52
	ZydisExceptionClassApxEvexInt ZydisExceptionClass_ = 53
	ZydisExceptionClassApxEvexKeylocker ZydisExceptionClass_ = 54
	ZydisExceptionClassApxEvexBmi ZydisExceptionClass_ = 55
	ZydisExceptionClassApxEvexCcmp ZydisExceptionClass_ = 56
	ZydisExceptionClassApxEvexCfcmov ZydisExceptionClass_ = 57
	ZydisExceptionClassApxEvexCmpccxadd ZydisExceptionClass_ = 58
	ZydisExceptionClassApxEvexEnqcmd ZydisExceptionClass_ = 59
	ZydisExceptionClassApxEvexInvept ZydisExceptionClass_ = 60
	ZydisExceptionClassApxEvexInvpcid ZydisExceptionClass_ = 61
	ZydisExceptionClassApxEvexInvvpid ZydisExceptionClass_ = 62
	ZydisExceptionClassApxEvexKmov ZydisExceptionClass_ = 63
	ZydisExceptionClassApxEvexPp2 ZydisExceptionClass_ = 64
	ZydisExceptionClassApxEvexSha ZydisExceptionClass_ = 65
	ZydisExceptionClassApxEvexCetWrss ZydisExceptionClass_ = 66
	ZydisExceptionClassApxEvexCetWruss ZydisExceptionClass_ = 67
	ZydisExceptionClassApxLegacyJmpabs ZydisExceptionClass_ = 68
	ZydisExceptionClassApxEvexRaoInt ZydisExceptionClass_ = 69
	ZydisExceptionClassUserMsrEvex ZydisExceptionClass_ = 70
	ZydisExceptionClassLegacyRaoInt ZydisExceptionClass_ = 71
	ZydisExceptionClassMaxValue ZydisExceptionClass_ = 71
	ZydisExceptionClassRequiredBits ZydisExceptionClass_ = 7
)

func (z ZydisExceptionClass_) String() string {
	switch z {
	case ZydisExceptionClassNone:
		return "Zydis Exception Class None"
	case ZydisExceptionClassSse1:
		return "Zydis Exception Class Sse 1"
	case ZydisExceptionClassSse2:
		return "Zydis Exception Class Sse 2"
	case ZydisExceptionClassSse3:
		return "Zydis Exception Class Sse 3"
	case ZydisExceptionClassSse4:
		return "Zydis Exception Class Sse 4"
	case ZydisExceptionClassSse5:
		return "Zydis Exception Class Sse 5"
	case ZydisExceptionClassSse7:
		return "Zydis Exception Class Sse 7"
	case ZydisExceptionClassAvx1:
		return "Zydis Exception Class Avx 1"
	case ZydisExceptionClassAvx2:
		return "Zydis Exception Class Avx 2"
	case ZydisExceptionClassAvx3:
		return "Zydis Exception Class Avx 3"
	case ZydisExceptionClassAvx4:
		return "Zydis Exception Class Avx 4"
	case ZydisExceptionClassAvx5:
		return "Zydis Exception Class Avx 5"
	case ZydisExceptionClassAvx6:
		return "Zydis Exception Class Avx 6"
	case ZydisExceptionClassAvx7:
		return "Zydis Exception Class Avx 7"
	case ZydisExceptionClassAvx8:
		return "Zydis Exception Class Avx 8"
	case ZydisExceptionClassAvx11:
		return "Zydis Exception Class Avx 11"
	case ZydisExceptionClassAvx12:
		return "Zydis Exception Class Avx 12"
	case ZydisExceptionClassAvx14:
		return "Zydis Exception Class Avx 14"
	case ZydisExceptionClassE1:
		return "Zydis Exception Class E1"
	case ZydisExceptionClassE1nf:
		return "Zydis Exception Class E1nf"
	case ZydisExceptionClassE2:
		return "Zydis Exception Class E2"
	case ZydisExceptionClassE2nf:
		return "Zydis Exception Class E2nf"
	case ZydisExceptionClassE3:
		return "Zydis Exception Class E3"
	case ZydisExceptionClassE3nf:
		return "Zydis Exception Class E3nf"
	case ZydisExceptionClassE4:
		return "Zydis Exception Class E4"
	case ZydisExceptionClassE4nf:
		return "Zydis Exception Class E4nf"
	case ZydisExceptionClassE5:
		return "Zydis Exception Class E5"
	case ZydisExceptionClassE5nf:
		return "Zydis Exception Class E5nf"
	case ZydisExceptionClassE6:
		return "Zydis Exception Class E6"
	case ZydisExceptionClassE6nf:
		return "Zydis Exception Class E6nf"
	case ZydisExceptionClassE7nm:
		return "Zydis Exception Class E7nm"
	case ZydisExceptionClassE7nm128:
		return "Zydis Exception Class E7nm 128"
	case ZydisExceptionClassE9nf:
		return "Zydis Exception Class E9nf"
	case ZydisExceptionClassE10:
		return "Zydis Exception Class E10"
	case ZydisExceptionClassE10nf:
		return "Zydis Exception Class E10nf"
	case ZydisExceptionClassE11:
		return "Zydis Exception Class E11"
	case ZydisExceptionClassE11nf:
		return "Zydis Exception Class E11nf"
	case ZydisExceptionClassE12:
		return "Zydis Exception Class E12"
	case ZydisExceptionClassE12np:
		return "Zydis Exception Class E12np"
	case ZydisExceptionClassK20:
		return "Zydis Exception Class K20"
	case ZydisExceptionClassK21:
		return "Zydis Exception Class K21"
	case ZydisExceptionClassAmxe1:
		return "Zydis Exception Class Amxe 1"
	case ZydisExceptionClassAmxe2:
		return "Zydis Exception Class Amxe 2"
	case ZydisExceptionClassAmxe3:
		return "Zydis Exception Class Amxe 3"
	case ZydisExceptionClassAmxe4:
		return "Zydis Exception Class Amxe 4"
	case ZydisExceptionClassAmxe5:
		return "Zydis Exception Class Amxe 5"
	case ZydisExceptionClassAmxe6:
		return "Zydis Exception Class Amxe 6"
	case ZydisExceptionClassAmxe1Evex:
		return "Zydis Exception Class Amxe 1 Evex"
	case ZydisExceptionClassAmxe2Evex:
		return "Zydis Exception Class Amxe 2 Evex"
	case ZydisExceptionClassAmxe3Evex:
		return "Zydis Exception Class Amxe 3 Evex"
	case ZydisExceptionClassAmxe4Evex:
		return "Zydis Exception Class Amxe 4 Evex"
	case ZydisExceptionClassAmxe5Evex:
		return "Zydis Exception Class Amxe 5 Evex"
	case ZydisExceptionClassAmxe6Evex:
		return "Zydis Exception Class Amxe 6 Evex"
	case ZydisExceptionClassApxEvexInt:
		return "Zydis Exception Class Apx Evex Int"
	case ZydisExceptionClassApxEvexKeylocker:
		return "Zydis Exception Class Apx Evex Keylocker"
	case ZydisExceptionClassApxEvexBmi:
		return "Zydis Exception Class Apx Evex Bmi"
	case ZydisExceptionClassApxEvexCcmp:
		return "Zydis Exception Class Apx Evex Ccmp"
	case ZydisExceptionClassApxEvexCfcmov:
		return "Zydis Exception Class Apx Evex Cfcmov"
	case ZydisExceptionClassApxEvexCmpccxadd:
		return "Zydis Exception Class Apx Evex Cmpccxadd"
	case ZydisExceptionClassApxEvexEnqcmd:
		return "Zydis Exception Class Apx Evex Enqcmd"
	case ZydisExceptionClassApxEvexInvept:
		return "Zydis Exception Class Apx Evex Invept"
	case ZydisExceptionClassApxEvexInvpcid:
		return "Zydis Exception Class Apx Evex Invpcid"
	case ZydisExceptionClassApxEvexInvvpid:
		return "Zydis Exception Class Apx Evex Invvpid"
	case ZydisExceptionClassApxEvexKmov:
		return "Zydis Exception Class Apx Evex Kmov"
	case ZydisExceptionClassApxEvexPp2:
		return "Zydis Exception Class Apx Evex Pp 2"
	case ZydisExceptionClassApxEvexSha:
		return "Zydis Exception Class Apx Evex Sha"
	case ZydisExceptionClassApxEvexCetWrss:
		return "Zydis Exception Class Apx Evex Cet Wrss"
	case ZydisExceptionClassApxEvexCetWruss:
		return "Zydis Exception Class Apx Evex Cet Wruss"
	case ZydisExceptionClassApxLegacyJmpabs:
		return "Zydis Exception Class Apx Legacy Jmpabs"
	case ZydisExceptionClassApxEvexRaoInt:
		return "Zydis Exception Class Apx Evex Rao Int"
	case ZydisExceptionClassUserMsrEvex:
		return "Zydis Exception Class User Msr Evex"
	case ZydisExceptionClassLegacyRaoInt:
		return "Zydis Exception Class Legacy Rao Int"
	default:
		return fmt.Sprintf("ZydisExceptionClass_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:567 -> ZydisMaskMode_
type ZydisMaskMode_ uint32

const (
	ZydisMaskModeNone ZydisMaskMode_ = 0
	ZydisMaskModeDisabled ZydisMaskMode_ = 1
	ZydisMaskModeMerging ZydisMaskMode_ = 2
	ZydisMaskModeZeroing ZydisMaskMode_ = 3
	ZydisMaskModeControl ZydisMaskMode_ = 4
	ZydisMaskModeControlZeroing ZydisMaskMode_ = 5
	ZydisMaskModeMaxValue ZydisMaskMode_ = 5
	ZydisMaskModeRequiredBits ZydisMaskMode_ = 3
)

func (z ZydisMaskMode_) String() string {
	switch z {
	case ZydisMaskModeNone:
		return "Zydis Mask Mode None"
	case ZydisMaskModeDisabled:
		return "Zydis Mask Mode Disabled"
	case ZydisMaskModeMerging:
		return "Zydis Mask Mode Merging"
	case ZydisMaskModeZeroing:
		return "Zydis Mask Mode Zeroing"
	case ZydisMaskModeControl:
		return "Zydis Mask Mode Control"
	case ZydisMaskModeControlZeroing:
		return "Zydis Mask Mode Control Zeroing"
	default:
		return fmt.Sprintf("ZydisMaskMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:608 -> ZydisBroadcastMode_
type ZydisBroadcastMode_ uint32

const (
	ZydisBroadcastModeNone ZydisBroadcastMode_ = 0
	ZydisBroadcastMode1To2 ZydisBroadcastMode_ = 1
	ZydisBroadcastMode1To4 ZydisBroadcastMode_ = 2
	ZydisBroadcastMode1To8 ZydisBroadcastMode_ = 3
	ZydisBroadcastMode1To16 ZydisBroadcastMode_ = 4
	ZydisBroadcastMode1To32 ZydisBroadcastMode_ = 5
	ZydisBroadcastMode1To64 ZydisBroadcastMode_ = 6
	ZydisBroadcastMode2To4 ZydisBroadcastMode_ = 7
	ZydisBroadcastMode2To8 ZydisBroadcastMode_ = 8
	ZydisBroadcastMode2To16 ZydisBroadcastMode_ = 9
	ZydisBroadcastMode4To8 ZydisBroadcastMode_ = 10
	ZydisBroadcastMode4To16 ZydisBroadcastMode_ = 11
	ZydisBroadcastMode8To16 ZydisBroadcastMode_ = 12
	ZydisBroadcastModeMaxValue ZydisBroadcastMode_ = 12
	ZydisBroadcastModeRequiredBits ZydisBroadcastMode_ = 4
)

func (z ZydisBroadcastMode_) String() string {
	switch z {
	case ZydisBroadcastModeNone:
		return "Zydis Broadcast Mode None"
	case ZydisBroadcastMode1To2:
		return "Zydis Broadcast Mode 1 To 2"
	case ZydisBroadcastMode1To4:
		return "Zydis Broadcast Mode 1 To 4"
	case ZydisBroadcastMode1To8:
		return "Zydis Broadcast Mode 1 To 8"
	case ZydisBroadcastMode1To16:
		return "Zydis Broadcast Mode 1 To 16"
	case ZydisBroadcastMode1To32:
		return "Zydis Broadcast Mode 1 To 32"
	case ZydisBroadcastMode1To64:
		return "Zydis Broadcast Mode 1 To 64"
	case ZydisBroadcastMode2To4:
		return "Zydis Broadcast Mode 2 To 4"
	case ZydisBroadcastMode2To8:
		return "Zydis Broadcast Mode 2 To 8"
	case ZydisBroadcastMode2To16:
		return "Zydis Broadcast Mode 2 To 16"
	case ZydisBroadcastMode4To8:
		return "Zydis Broadcast Mode 4 To 8"
	case ZydisBroadcastMode4To16:
		return "Zydis Broadcast Mode 4 To 16"
	case ZydisBroadcastMode8To16:
		return "Zydis Broadcast Mode 8 To 16"
	default:
		return fmt.Sprintf("ZydisBroadcastMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:641 -> ZydisRoundingMode_
type ZydisRoundingMode_ uint32

const (
	ZydisRoundingModeNone ZydisRoundingMode_ = 0
	ZydisRoundingModeRn ZydisRoundingMode_ = 1
	ZydisRoundingModeRd ZydisRoundingMode_ = 2
	ZydisRoundingModeRu ZydisRoundingMode_ = 3
	ZydisRoundingModeRz ZydisRoundingMode_ = 4
	ZydisRoundingModeMaxValue ZydisRoundingMode_ = 4
	ZydisRoundingModeRequiredBits ZydisRoundingMode_ = 3
)

func (z ZydisRoundingMode_) String() string {
	switch z {
	case ZydisRoundingModeNone:
		return "Zydis Rounding Mode None"
	case ZydisRoundingModeRn:
		return "Zydis Rounding Mode Rn"
	case ZydisRoundingModeRd:
		return "Zydis Rounding Mode Rd"
	case ZydisRoundingModeRu:
		return "Zydis Rounding Mode Ru"
	case ZydisRoundingModeRz:
		return "Zydis Rounding Mode Rz"
	default:
		return fmt.Sprintf("ZydisRoundingMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:678 -> ZydisSwizzleMode_
type ZydisSwizzleMode_ uint32

const (
	ZydisSwizzleModeNone ZydisSwizzleMode_ = 0
	ZydisSwizzleModeDcba ZydisSwizzleMode_ = 1
	ZydisSwizzleModeCdab ZydisSwizzleMode_ = 2
	ZydisSwizzleModeBadc ZydisSwizzleMode_ = 3
	ZydisSwizzleModeDacb ZydisSwizzleMode_ = 4
	ZydisSwizzleModeAaaa ZydisSwizzleMode_ = 5
	ZydisSwizzleModeBbbb ZydisSwizzleMode_ = 6
	ZydisSwizzleModeCccc ZydisSwizzleMode_ = 7
	ZydisSwizzleModeDddd ZydisSwizzleMode_ = 8
	ZydisSwizzleModeMaxValue ZydisSwizzleMode_ = 8
	ZydisSwizzleModeRequiredBits ZydisSwizzleMode_ = 4
)

func (z ZydisSwizzleMode_) String() string {
	switch z {
	case ZydisSwizzleModeNone:
		return "Zydis Swizzle Mode None"
	case ZydisSwizzleModeDcba:
		return "Zydis Swizzle Mode Dcba"
	case ZydisSwizzleModeCdab:
		return "Zydis Swizzle Mode Cdab"
	case ZydisSwizzleModeBadc:
		return "Zydis Swizzle Mode Badc"
	case ZydisSwizzleModeDacb:
		return "Zydis Swizzle Mode Dacb"
	case ZydisSwizzleModeAaaa:
		return "Zydis Swizzle Mode Aaaa"
	case ZydisSwizzleModeBbbb:
		return "Zydis Swizzle Mode Bbbb"
	case ZydisSwizzleModeCccc:
		return "Zydis Swizzle Mode Cccc"
	case ZydisSwizzleModeDddd:
		return "Zydis Swizzle Mode Dddd"
	default:
		return fmt.Sprintf("ZydisSwizzleMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:707 -> ZydisConversionMode_
type ZydisConversionMode_ uint32

const (
	ZydisConversionModeNone ZydisConversionMode_ = 0
	ZydisConversionModeFloat16 ZydisConversionMode_ = 1
	ZydisConversionModeSint8 ZydisConversionMode_ = 2
	ZydisConversionModeUint8 ZydisConversionMode_ = 3
	ZydisConversionModeSint16 ZydisConversionMode_ = 4
	ZydisConversionModeUint16 ZydisConversionMode_ = 5
	ZydisConversionModeMaxValue ZydisConversionMode_ = 5
	ZydisConversionModeRequiredBits ZydisConversionMode_ = 3
)

func (z ZydisConversionMode_) String() string {
	switch z {
	case ZydisConversionModeNone:
		return "Zydis Conversion Mode None"
	case ZydisConversionModeFloat16:
		return "Zydis Conversion Mode Float 16"
	case ZydisConversionModeSint8:
		return "Zydis Conversion Mode Sint 8"
	case ZydisConversionModeUint8:
		return "Zydis Conversion Mode Uint 8"
	case ZydisConversionModeSint16:
		return "Zydis Conversion Mode Sint 16"
	case ZydisConversionModeUint16:
		return "Zydis Conversion Mode Uint 16"
	default:
		return fmt.Sprintf("ZydisConversionMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:782 -> ZydisSourceConditionCode_
type ZydisSourceConditionCode_ uint32

const (
	ZydisSccNone ZydisSourceConditionCode_ = 0
	ZydisSccO ZydisSourceConditionCode_ = 1
	ZydisSccNo ZydisSourceConditionCode_ = 2
	ZydisSccB ZydisSourceConditionCode_ = 3
	ZydisSccNb ZydisSourceConditionCode_ = 4
	ZydisSccZ ZydisSourceConditionCode_ = 5
	ZydisSccNz ZydisSourceConditionCode_ = 6
	ZydisSccBe ZydisSourceConditionCode_ = 7
	ZydisSccNbe ZydisSourceConditionCode_ = 8
	ZydisSccS ZydisSourceConditionCode_ = 9
	ZydisSccNs ZydisSourceConditionCode_ = 10
	ZydisSccTrue ZydisSourceConditionCode_ = 11
	ZydisSccFalse ZydisSourceConditionCode_ = 12
	ZydisSccL ZydisSourceConditionCode_ = 13
	ZydisSccNl ZydisSourceConditionCode_ = 14
	ZydisSccLe ZydisSourceConditionCode_ = 15
	ZydisSccNle ZydisSourceConditionCode_ = 16
	ZydisSccMaxValue ZydisSourceConditionCode_ = 16
	ZydisSccRequiredBits ZydisSourceConditionCode_ = 5
)

func (z ZydisSourceConditionCode_) String() string {
	switch z {
	case ZydisSccNone:
		return "Zydis Scc None"
	case ZydisSccO:
		return "Zydis Scc O"
	case ZydisSccNo:
		return "Zydis Scc No"
	case ZydisSccB:
		return "Zydis Scc B"
	case ZydisSccNb:
		return "Zydis Scc Nb"
	case ZydisSccZ:
		return "Zydis Scc Z"
	case ZydisSccNz:
		return "Zydis Scc Nz"
	case ZydisSccBe:
		return "Zydis Scc Be"
	case ZydisSccNbe:
		return "Zydis Scc Nbe"
	case ZydisSccS:
		return "Zydis Scc S"
	case ZydisSccNs:
		return "Zydis Scc Ns"
	case ZydisSccTrue:
		return "Zydis Scc True"
	case ZydisSccFalse:
		return "Zydis Scc False"
	case ZydisSccL:
		return "Zydis Scc L"
	case ZydisSccNl:
		return "Zydis Scc Nl"
	case ZydisSccLe:
		return "Zydis Scc Le"
	case ZydisSccNle:
		return "Zydis Scc Nle"
	default:
		return fmt.Sprintf("ZydisSourceConditionCode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:819 -> ZydisPrefixType_
type ZydisPrefixType_ uint32

const (
	ZydisPrefixTypeIgnored ZydisPrefixType_ = 0
	ZydisPrefixTypeEffective ZydisPrefixType_ = 1
	ZydisPrefixTypeMandatory ZydisPrefixType_ = 2
	ZydisPrefixTypeMaxValue ZydisPrefixType_ = 2
	ZydisPrefixTypeRequiredBits ZydisPrefixType_ = 2
)

func (z ZydisPrefixType_) String() string {
	switch z {
	case ZydisPrefixTypeIgnored:
		return "Zydis Prefix Type Ignored"
	case ZydisPrefixTypeEffective:
		return "Zydis Prefix Type Effective"
	case ZydisPrefixTypeMandatory:
		return "Zydis Prefix Type Mandatory"
	default:
		return fmt.Sprintf("ZydisPrefixType_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:55 -> ZydisDecoderMode_
type ZydisDecoderMode_ uint32

const (
	ZydisDecoderModeMinimal ZydisDecoderMode_ = 0
	ZydisDecoderModeAmdBranches ZydisDecoderMode_ = 1
	ZydisDecoderModeKnc ZydisDecoderMode_ = 2
	ZydisDecoderModeMpx ZydisDecoderMode_ = 3
	ZydisDecoderModeCet ZydisDecoderMode_ = 4
	ZydisDecoderModeLzcnt ZydisDecoderMode_ = 5
	ZydisDecoderModeTzcnt ZydisDecoderMode_ = 6
	ZydisDecoderModeWbnoinvd ZydisDecoderMode_ = 7
	ZydisDecoderModeCldemote ZydisDecoderMode_ = 8
	ZydisDecoderModeIprefetch ZydisDecoderMode_ = 9
	ZydisDecoderModeUd0Compat ZydisDecoderMode_ = 10
	ZydisDecoderModeApx ZydisDecoderMode_ = 11
	ZydisDecoderModeMaxValue ZydisDecoderMode_ = 11
	ZydisDecoderModeRequiredBits ZydisDecoderMode_ = 4
)

func (z ZydisDecoderMode_) String() string {
	switch z {
	case ZydisDecoderModeMinimal:
		return "Zydis Decoder Mode Minimal"
	case ZydisDecoderModeAmdBranches:
		return "Zydis Decoder Mode Amd Branches"
	case ZydisDecoderModeKnc:
		return "Zydis Decoder Mode Knc"
	case ZydisDecoderModeMpx:
		return "Zydis Decoder Mode Mpx"
	case ZydisDecoderModeCet:
		return "Zydis Decoder Mode Cet"
	case ZydisDecoderModeLzcnt:
		return "Zydis Decoder Mode Lzcnt"
	case ZydisDecoderModeTzcnt:
		return "Zydis Decoder Mode Tzcnt"
	case ZydisDecoderModeWbnoinvd:
		return "Zydis Decoder Mode Wbnoinvd"
	case ZydisDecoderModeCldemote:
		return "Zydis Decoder Mode Cldemote"
	case ZydisDecoderModeIprefetch:
		return "Zydis Decoder Mode Iprefetch"
	case ZydisDecoderModeUd0Compat:
		return "Zydis Decoder Mode Ud 0 Compat"
	case ZydisDecoderModeApx:
		return "Zydis Decoder Mode Apx"
	default:
		return fmt.Sprintf("ZydisDecoderMode_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:92 -> ZydisEncodableEncoding_
type ZydisEncodableEncoding_ uint32

const (
	ZydisEncodableEncodingDefault ZydisEncodableEncoding_ = 0
	ZydisEncodableEncodingLegacy ZydisEncodableEncoding_ = 1
	ZydisEncodableEncoding3dnow ZydisEncodableEncoding_ = 2
	ZydisEncodableEncodingXop ZydisEncodableEncoding_ = 4
	ZydisEncodableEncodingVex ZydisEncodableEncoding_ = 8
	ZydisEncodableEncodingEvex ZydisEncodableEncoding_ = 16
	ZydisEncodableEncodingMvex ZydisEncodableEncoding_ = 32
	ZydisEncodableEncodingMaxValue ZydisEncodableEncoding_ = 63
	ZydisEncodableEncodingRequiredBits ZydisEncodableEncoding_ = 6
)

func (z ZydisEncodableEncoding_) String() string {
	switch z {
	case ZydisEncodableEncodingDefault:
		return "Zydis Encodable Encoding Default"
	case ZydisEncodableEncodingLegacy:
		return "Zydis Encodable Encoding Legacy"
	case ZydisEncodableEncoding3dnow:
		return "Zydis Encodable Encoding 3dnow"
	case ZydisEncodableEncodingXop:
		return "Zydis Encodable Encoding Xop"
	case ZydisEncodableEncodingVex:
		return "Zydis Encodable Encoding Vex"
	case ZydisEncodableEncodingEvex:
		return "Zydis Encodable Encoding Evex"
	case ZydisEncodableEncodingMvex:
		return "Zydis Encodable Encoding Mvex"
	case ZydisEncodableEncodingMaxValue:
		return "Zydis Encodable Encoding Max Value"
	case ZydisEncodableEncodingRequiredBits:
		return "Zydis Encodable Encoding Required Bits"
	default:
		return fmt.Sprintf("ZydisEncodableEncoding_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:118 -> ZydisBranchWidth_
type ZydisBranchWidth_ uint32

const (
	ZydisBranchWidthNone ZydisBranchWidth_ = 0
	ZydisBranchWidth8 ZydisBranchWidth_ = 1
	ZydisBranchWidth16 ZydisBranchWidth_ = 2
	ZydisBranchWidth32 ZydisBranchWidth_ = 3
	ZydisBranchWidth64 ZydisBranchWidth_ = 4
	ZydisBranchWidthMaxValue ZydisBranchWidth_ = 4
	ZydisBranchWidthRequiredBits ZydisBranchWidth_ = 3
)

func (z ZydisBranchWidth_) String() string {
	switch z {
	case ZydisBranchWidthNone:
		return "Zydis Branch Width None"
	case ZydisBranchWidth8:
		return "Zydis Branch Width 8"
	case ZydisBranchWidth16:
		return "Zydis Branch Width 16"
	case ZydisBranchWidth32:
		return "Zydis Branch Width 32"
	case ZydisBranchWidth64:
		return "Zydis Branch Width 64"
	default:
		return fmt.Sprintf("ZydisBranchWidth_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:140 -> ZydisAddressSizeHint_
type ZydisAddressSizeHint_ uint32

const (
	ZydisAddressSizeHintNone ZydisAddressSizeHint_ = 0
	ZydisAddressSizeHint16 ZydisAddressSizeHint_ = 1
	ZydisAddressSizeHint32 ZydisAddressSizeHint_ = 2
	ZydisAddressSizeHint64 ZydisAddressSizeHint_ = 3
	ZydisAddressSizeHintMaxValue ZydisAddressSizeHint_ = 3
	ZydisAddressSizeHintRequiredBits ZydisAddressSizeHint_ = 2
)

func (z ZydisAddressSizeHint_) String() string {
	switch z {
	case ZydisAddressSizeHintNone:
		return "Zydis Address Size Hint None"
	case ZydisAddressSizeHint16:
		return "Zydis Address Size Hint 16"
	case ZydisAddressSizeHint32:
		return "Zydis Address Size Hint 32"
	case ZydisAddressSizeHint64:
		return "Zydis Address Size Hint 64"
	default:
		return fmt.Sprintf("ZydisAddressSizeHint_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:162 -> ZydisOperandSizeHint_
type ZydisOperandSizeHint_ uint32

const (
	ZydisOperandSizeHintNone ZydisOperandSizeHint_ = 0
	ZydisOperandSizeHint8 ZydisOperandSizeHint_ = 1
	ZydisOperandSizeHint16 ZydisOperandSizeHint_ = 2
	ZydisOperandSizeHint32 ZydisOperandSizeHint_ = 3
	ZydisOperandSizeHint64 ZydisOperandSizeHint_ = 4
	ZydisOperandSizeHintMaxValue ZydisOperandSizeHint_ = 4
	ZydisOperandSizeHintRequiredBits ZydisOperandSizeHint_ = 3
)

func (z ZydisOperandSizeHint_) String() string {
	switch z {
	case ZydisOperandSizeHintNone:
		return "Zydis Operand Size Hint None"
	case ZydisOperandSizeHint8:
		return "Zydis Operand Size Hint 8"
	case ZydisOperandSizeHint16:
		return "Zydis Operand Size Hint 16"
	case ZydisOperandSizeHint32:
		return "Zydis Operand Size Hint 32"
	case ZydisOperandSizeHint64:
		return "Zydis Operand Size Hint 64"
	default:
		return fmt.Sprintf("ZydisOperandSizeHint_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:66 -> ZydisFormatterStyle_
type ZydisFormatterStyle_ uint32

const (
	ZydisFormatterStyleAtt ZydisFormatterStyle_ = 0
	ZydisFormatterStyleIntel ZydisFormatterStyle_ = 1
	ZydisFormatterStyleIntelMasm ZydisFormatterStyle_ = 2
	ZydisFormatterStyleMaxValue ZydisFormatterStyle_ = 2
	ZydisFormatterStyleRequiredBits ZydisFormatterStyle_ = 2
)

func (z ZydisFormatterStyle_) String() string {
	switch z {
	case ZydisFormatterStyleAtt:
		return "Zydis Formatter Style Att"
	case ZydisFormatterStyleIntel:
		return "Zydis Formatter Style Intel"
	case ZydisFormatterStyleIntelMasm:
		return "Zydis Formatter Style Intel Masm"
	default:
		return fmt.Sprintf("ZydisFormatterStyle_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:101 -> ZydisFormatterProperty_
type ZydisFormatterProperty_ uint32

const (
	ZydisFormatterPropForceSize ZydisFormatterProperty_ = 0
	ZydisFormatterPropForceSegment ZydisFormatterProperty_ = 1
	ZydisFormatterPropForceScaleOne ZydisFormatterProperty_ = 2
	ZydisFormatterPropForceRelativeBranches ZydisFormatterProperty_ = 3
	ZydisFormatterPropForceRelativeRiprel ZydisFormatterProperty_ = 4
	ZydisFormatterPropPrintBranchSize ZydisFormatterProperty_ = 5
	ZydisFormatterPropDetailedPrefixes ZydisFormatterProperty_ = 6
	ZydisFormatterPropAddrBase ZydisFormatterProperty_ = 7
	ZydisFormatterPropAddrSignedness ZydisFormatterProperty_ = 8
	ZydisFormatterPropAddrPaddingAbsolute ZydisFormatterProperty_ = 9
	ZydisFormatterPropAddrPaddingRelative ZydisFormatterProperty_ = 10
	ZydisFormatterPropDispBase ZydisFormatterProperty_ = 11
	ZydisFormatterPropDispSignedness ZydisFormatterProperty_ = 12
	ZydisFormatterPropDispPadding ZydisFormatterProperty_ = 13
	ZydisFormatterPropImmBase ZydisFormatterProperty_ = 14
	ZydisFormatterPropImmSignedness ZydisFormatterProperty_ = 15
	ZydisFormatterPropImmPadding ZydisFormatterProperty_ = 16
	ZydisFormatterPropUppercasePrefixes ZydisFormatterProperty_ = 17
	ZydisFormatterPropUppercaseMnemonic ZydisFormatterProperty_ = 18
	ZydisFormatterPropUppercaseRegisters ZydisFormatterProperty_ = 19
	ZydisFormatterPropUppercaseTypecasts ZydisFormatterProperty_ = 20
	ZydisFormatterPropUppercaseDecorators ZydisFormatterProperty_ = 21
	ZydisFormatterPropDecPrefix ZydisFormatterProperty_ = 22
	ZydisFormatterPropDecSuffix ZydisFormatterProperty_ = 23
	ZydisFormatterPropHexUppercase ZydisFormatterProperty_ = 24
	ZydisFormatterPropHexForceLeadingNumber ZydisFormatterProperty_ = 25
	ZydisFormatterPropHexPrefix ZydisFormatterProperty_ = 26
	ZydisFormatterPropHexSuffix ZydisFormatterProperty_ = 27
	ZydisFormatterPropDecoApxNfUseSuffix ZydisFormatterProperty_ = 28
	ZydisFormatterPropDecoApxDfvUseImmediate ZydisFormatterProperty_ = 29
	ZydisFormatterPropMaxValue ZydisFormatterProperty_ = 29
	ZydisFormatterPropRequiredBits ZydisFormatterProperty_ = 5
)

func (z ZydisFormatterProperty_) String() string {
	switch z {
	case ZydisFormatterPropForceSize:
		return "Zydis Formatter Prop Force Size"
	case ZydisFormatterPropForceSegment:
		return "Zydis Formatter Prop Force Segment"
	case ZydisFormatterPropForceScaleOne:
		return "Zydis Formatter Prop Force Scale One"
	case ZydisFormatterPropForceRelativeBranches:
		return "Zydis Formatter Prop Force Relative Branches"
	case ZydisFormatterPropForceRelativeRiprel:
		return "Zydis Formatter Prop Force Relative Riprel"
	case ZydisFormatterPropPrintBranchSize:
		return "Zydis Formatter Prop Print Branch Size"
	case ZydisFormatterPropDetailedPrefixes:
		return "Zydis Formatter Prop Detailed Prefixes"
	case ZydisFormatterPropAddrBase:
		return "Zydis Formatter Prop Addr Base"
	case ZydisFormatterPropAddrSignedness:
		return "Zydis Formatter Prop Addr Signedness"
	case ZydisFormatterPropAddrPaddingAbsolute:
		return "Zydis Formatter Prop Addr Padding Absolute"
	case ZydisFormatterPropAddrPaddingRelative:
		return "Zydis Formatter Prop Addr Padding Relative"
	case ZydisFormatterPropDispBase:
		return "Zydis Formatter Prop Disp Base"
	case ZydisFormatterPropDispSignedness:
		return "Zydis Formatter Prop Disp Signedness"
	case ZydisFormatterPropDispPadding:
		return "Zydis Formatter Prop Disp Padding"
	case ZydisFormatterPropImmBase:
		return "Zydis Formatter Prop Imm Base"
	case ZydisFormatterPropImmSignedness:
		return "Zydis Formatter Prop Imm Signedness"
	case ZydisFormatterPropImmPadding:
		return "Zydis Formatter Prop Imm Padding"
	case ZydisFormatterPropUppercasePrefixes:
		return "Zydis Formatter Prop Uppercase Prefixes"
	case ZydisFormatterPropUppercaseMnemonic:
		return "Zydis Formatter Prop Uppercase Mnemonic"
	case ZydisFormatterPropUppercaseRegisters:
		return "Zydis Formatter Prop Uppercase Registers"
	case ZydisFormatterPropUppercaseTypecasts:
		return "Zydis Formatter Prop Uppercase Typecasts"
	case ZydisFormatterPropUppercaseDecorators:
		return "Zydis Formatter Prop Uppercase Decorators"
	case ZydisFormatterPropDecPrefix:
		return "Zydis Formatter Prop Dec Prefix"
	case ZydisFormatterPropDecSuffix:
		return "Zydis Formatter Prop Dec Suffix"
	case ZydisFormatterPropHexUppercase:
		return "Zydis Formatter Prop Hex Uppercase"
	case ZydisFormatterPropHexForceLeadingNumber:
		return "Zydis Formatter Prop Hex Force Leading Number"
	case ZydisFormatterPropHexPrefix:
		return "Zydis Formatter Prop Hex Prefix"
	case ZydisFormatterPropHexSuffix:
		return "Zydis Formatter Prop Hex Suffix"
	case ZydisFormatterPropDecoApxNfUseSuffix:
		return "Zydis Formatter Prop Deco Apx Nf Use Suffix"
	case ZydisFormatterPropDecoApxDfvUseImmediate:
		return "Zydis Formatter Prop Deco Apx Dfv Use Immediate"
	default:
		return fmt.Sprintf("ZydisFormatterProperty_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:376 -> ZydisNumericBase_
type ZydisNumericBase_ uint32

const (
	ZydisNumericBaseDec ZydisNumericBase_ = 0
	ZydisNumericBaseHex ZydisNumericBase_ = 1
	ZydisNumericBaseMaxValue ZydisNumericBase_ = 1
	ZydisNumericBaseRequiredBits ZydisNumericBase_ = 1
)

func (z ZydisNumericBase_) String() string {
	switch z {
	case ZydisNumericBaseDec:
		return "Zydis Numeric Base Dec"
	case ZydisNumericBaseHex:
		return "Zydis Numeric Base Hex"
	default:
		return fmt.Sprintf("ZydisNumericBase_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:402 -> ZydisSignedness_
type ZydisSignedness_ uint32

const (
	ZydisSignednessAuto ZydisSignedness_ = 0
	ZydisSignednessSigned ZydisSignedness_ = 1
	ZydisSignednessUnsigned ZydisSignedness_ = 2
	ZydisSignednessMaxValue ZydisSignedness_ = 2
	ZydisSignednessRequiredBits ZydisSignedness_ = 2
)

func (z ZydisSignedness_) String() string {
	switch z {
	case ZydisSignednessAuto:
		return "Zydis Signedness Auto"
	case ZydisSignednessSigned:
		return "Zydis Signedness Signed"
	case ZydisSignednessUnsigned:
		return "Zydis Signedness Unsigned"
	default:
		return fmt.Sprintf("ZydisSignedness_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:434 -> ZydisPadding_
type ZydisPadding_ int32

const (
	ZydisPaddingDisabled ZydisPadding_ = 0
	ZydisPaddingAuto ZydisPadding_ = -1
	ZydisPaddingMaxValue ZydisPadding_ = -1
	ZydisPaddingRequiredBits ZydisPadding_ = 32
)

func (z ZydisPadding_) String() string {
	switch z {
	case ZydisPaddingDisabled:
		return "Zydis Padding Disabled"
	case ZydisPaddingAuto:
		return "Zydis Padding Auto"
	case ZydisPaddingRequiredBits:
		return "Zydis Padding Required Bits"
	default:
		return fmt.Sprintf("ZydisPadding_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:466 -> ZydisFormatterFunction_
type ZydisFormatterFunction_ uint32

const (
	ZydisFormatterFuncPreInstruction ZydisFormatterFunction_ = 0
	ZydisFormatterFuncPostInstruction ZydisFormatterFunction_ = 1
	ZydisFormatterFuncFormatInstruction ZydisFormatterFunction_ = 2
	ZydisFormatterFuncPreOperand ZydisFormatterFunction_ = 3
	ZydisFormatterFuncPostOperand ZydisFormatterFunction_ = 4
	ZydisFormatterFuncFormatOperandReg ZydisFormatterFunction_ = 5
	ZydisFormatterFuncFormatOperandMem ZydisFormatterFunction_ = 6
	ZydisFormatterFuncFormatOperandPtr ZydisFormatterFunction_ = 7
	ZydisFormatterFuncFormatOperandImm ZydisFormatterFunction_ = 8
	ZydisFormatterFuncPrintMnemonic ZydisFormatterFunction_ = 9
	ZydisFormatterFuncPrintRegister ZydisFormatterFunction_ = 10
	ZydisFormatterFuncPrintAddressAbs ZydisFormatterFunction_ = 11
	ZydisFormatterFuncPrintAddressRel ZydisFormatterFunction_ = 12
	ZydisFormatterFuncPrintDisp ZydisFormatterFunction_ = 13
	ZydisFormatterFuncPrintImm ZydisFormatterFunction_ = 14
	ZydisFormatterFuncPrintTypecast ZydisFormatterFunction_ = 15
	ZydisFormatterFuncPrintSegment ZydisFormatterFunction_ = 16
	ZydisFormatterFuncPrintPrefixes ZydisFormatterFunction_ = 17
	ZydisFormatterFuncPrintDecorator ZydisFormatterFunction_ = 18
	ZydisFormatterFuncMaxValue ZydisFormatterFunction_ = 18
	ZydisFormatterFuncRequiredBits ZydisFormatterFunction_ = 5
)

func (z ZydisFormatterFunction_) String() string {
	switch z {
	case ZydisFormatterFuncPreInstruction:
		return "Zydis Formatter Func Pre Instruction"
	case ZydisFormatterFuncPostInstruction:
		return "Zydis Formatter Func Post Instruction"
	case ZydisFormatterFuncFormatInstruction:
		return "Zydis Formatter Func Format Instruction"
	case ZydisFormatterFuncPreOperand:
		return "Zydis Formatter Func Pre Operand"
	case ZydisFormatterFuncPostOperand:
		return "Zydis Formatter Func Post Operand"
	case ZydisFormatterFuncFormatOperandReg:
		return "Zydis Formatter Func Format Operand Reg"
	case ZydisFormatterFuncFormatOperandMem:
		return "Zydis Formatter Func Format Operand Mem"
	case ZydisFormatterFuncFormatOperandPtr:
		return "Zydis Formatter Func Format Operand Ptr"
	case ZydisFormatterFuncFormatOperandImm:
		return "Zydis Formatter Func Format Operand Imm"
	case ZydisFormatterFuncPrintMnemonic:
		return "Zydis Formatter Func Print Mnemonic"
	case ZydisFormatterFuncPrintRegister:
		return "Zydis Formatter Func Print Register"
	case ZydisFormatterFuncPrintAddressAbs:
		return "Zydis Formatter Func Print Address Abs"
	case ZydisFormatterFuncPrintAddressRel:
		return "Zydis Formatter Func Print Address Rel"
	case ZydisFormatterFuncPrintDisp:
		return "Zydis Formatter Func Print Disp"
	case ZydisFormatterFuncPrintImm:
		return "Zydis Formatter Func Print Imm"
	case ZydisFormatterFuncPrintTypecast:
		return "Zydis Formatter Func Print Typecast"
	case ZydisFormatterFuncPrintSegment:
		return "Zydis Formatter Func Print Segment"
	case ZydisFormatterFuncPrintPrefixes:
		return "Zydis Formatter Func Print Prefixes"
	case ZydisFormatterFuncPrintDecorator:
		return "Zydis Formatter Func Print Decorator"
	default:
		return fmt.Sprintf("ZydisFormatterFunction_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:627 -> ZydisDecorator_
type ZydisDecorator_ uint32

const (
	ZydisDecoratorInvalid ZydisDecorator_ = 0
	ZydisDecoratorMask ZydisDecorator_ = 1
	ZydisDecoratorBc ZydisDecorator_ = 2
	ZydisDecoratorRc ZydisDecorator_ = 3
	ZydisDecoratorSae ZydisDecorator_ = 4
	ZydisDecoratorSwizzle ZydisDecorator_ = 5
	ZydisDecoratorConversion ZydisDecorator_ = 6
	ZydisDecoratorEh ZydisDecorator_ = 7
	ZydisDecoratorApxNf ZydisDecorator_ = 8
	ZydisDecoratorApxDfv ZydisDecorator_ = 9
	ZydisDecoratorMaxValue ZydisDecorator_ = 9
	ZydisDecoratorRequiredBits ZydisDecorator_ = 4
)

func (z ZydisDecorator_) String() string {
	switch z {
	case ZydisDecoratorInvalid:
		return "Zydis Decorator Invalid"
	case ZydisDecoratorMask:
		return "Zydis Decorator Mask"
	case ZydisDecoratorBc:
		return "Zydis Decorator Bc"
	case ZydisDecoratorRc:
		return "Zydis Decorator Rc"
	case ZydisDecoratorSae:
		return "Zydis Decorator Sae"
	case ZydisDecoratorSwizzle:
		return "Zydis Decorator Swizzle"
	case ZydisDecoratorConversion:
		return "Zydis Decorator Conversion"
	case ZydisDecoratorEh:
		return "Zydis Decorator Eh"
	case ZydisDecoratorApxNf:
		return "Zydis Decorator Apx Nf"
	case ZydisDecoratorApxDfv:
		return "Zydis Decorator Apx Dfv"
	default:
		return fmt.Sprintf("ZydisDecorator_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:68 -> ZydisInstructionSegment_
type ZydisInstructionSegment_ uint32

const (
	ZydisInstrSegmentNone ZydisInstructionSegment_ = 0
	ZydisInstrSegmentPrefixes ZydisInstructionSegment_ = 1
	ZydisInstrSegmentRex ZydisInstructionSegment_ = 2
	ZydisInstrSegmentRex2 ZydisInstructionSegment_ = 3
	ZydisInstrSegmentXop ZydisInstructionSegment_ = 4
	ZydisInstrSegmentVex ZydisInstructionSegment_ = 5
	ZydisInstrSegmentEvex ZydisInstructionSegment_ = 6
	ZydisInstrSegmentMvex ZydisInstructionSegment_ = 7
	ZydisInstrSegmentOpcode ZydisInstructionSegment_ = 8
	ZydisInstrSegmentModrm ZydisInstructionSegment_ = 9
	ZydisInstrSegmentSib ZydisInstructionSegment_ = 10
	ZydisInstrSegmentDisplacement ZydisInstructionSegment_ = 11
	ZydisInstrSegmentImmediate ZydisInstructionSegment_ = 12
	ZydisInstrSegmentMaxValue ZydisInstructionSegment_ = 12
	ZydisInstrSegmentRequiredBits ZydisInstructionSegment_ = 4
)

func (z ZydisInstructionSegment_) String() string {
	switch z {
	case ZydisInstrSegmentNone:
		return "Zydis Instr Segment None"
	case ZydisInstrSegmentPrefixes:
		return "Zydis Instr Segment Prefixes"
	case ZydisInstrSegmentRex:
		return "Zydis Instr Segment Rex"
	case ZydisInstrSegmentRex2:
		return "Zydis Instr Segment Rex 2"
	case ZydisInstrSegmentXop:
		return "Zydis Instr Segment Xop"
	case ZydisInstrSegmentVex:
		return "Zydis Instr Segment Vex"
	case ZydisInstrSegmentEvex:
		return "Zydis Instr Segment Evex"
	case ZydisInstrSegmentMvex:
		return "Zydis Instr Segment Mvex"
	case ZydisInstrSegmentOpcode:
		return "Zydis Instr Segment Opcode"
	case ZydisInstrSegmentModrm:
		return "Zydis Instr Segment Modrm"
	case ZydisInstrSegmentSib:
		return "Zydis Instr Segment Sib"
	case ZydisInstrSegmentDisplacement:
		return "Zydis Instr Segment Displacement"
	case ZydisInstrSegmentImmediate:
		return "Zydis Instr Segment Immediate"
	default:
		return fmt.Sprintf("ZydisInstructionSegment_(0x%X)", uint32(z))
	}
}

// Source: Zydis.h:132 -> ZydisFeature_
type ZydisFeature_ uint32

const (
	ZydisFeatureDecoder ZydisFeature_ = 0
	ZydisFeatureEncoder ZydisFeature_ = 1
	ZydisFeatureFormatter ZydisFeature_ = 2
	ZydisFeatureAvx512 ZydisFeature_ = 3
	ZydisFeatureKnc ZydisFeature_ = 4
	ZydisFeatureSegment ZydisFeature_ = 5
	ZydisFeatureMaxValue ZydisFeature_ = 5
	ZydisFeatureRequiredBits ZydisFeature_ = 3
)

func (z ZydisFeature_) String() string {
	switch z {
	case ZydisFeatureDecoder:
		return "Zydis Feature Decoder"
	case ZydisFeatureEncoder:
		return "Zydis Feature Encoder"
	case ZydisFeatureFormatter:
		return "Zydis Feature Formatter"
	case ZydisFeatureAvx512:
		return "Zydis Feature Avx 512"
	case ZydisFeatureKnc:
		return "Zydis Feature Knc"
	case ZydisFeatureSegment:
		return "Zydis Feature Segment"
	default:
		return fmt.Sprintf("ZydisFeature_(0x%X)", uint32(z))
	}
}

// Source: unknown.h:59 -> ZydisShortString_
type ZydisShortString_ struct {
	Data *int8
	Size uint8
	_  [7]byte
}

// Source: unknown.h:214 -> ZydisRegisterContext_
type ZydisRegisterContext_ struct {
	Values [331]uint64
}

// Source: unknown.h:114 -> ZydisDecodedOperandReg_
type ZydisDecodedOperandReg_ struct {
	Value ZydisRegister
}

// Source: unknown.h:125 -> ZydisDecodedOperandMem_
type ZydisDecodedOperandMem_ struct {
	Type ZydisMemoryOperandType
	Segment ZydisRegister
	Base ZydisRegister
	Index ZydisRegister
	Scale uint8
	_  [7]byte
	Disp ZydisDecodedOperandMemDisp_
}

// Source: unknown.h:0 -> ZydisDecodedOperandMemDisp_
type ZydisDecodedOperandMemDisp_ struct {
	Value int64
	Offset uint8
	Size uint8
	_  [6]byte
}

// Source: unknown.h:171 -> ZydisDecodedOperandPtr_
type ZydisDecodedOperandPtr_ struct {
	Segment uint16
	_  [2]byte
	Offset uint32
}

// Source: unknown.h:180 -> ZydisDecodedOperandImm_
type ZydisDecodedOperandImm_ struct {
	IsSigned uint8
	IsAddress uint8
	IsRelative uint8
	_  [5]byte
	Value ZydisDecodedOperandImmValue_
	Offset uint8
	Size uint8
	_  [6]byte
}

// Source: unknown.h:0 -> ZydisDecodedOperandImmValue_
type ZydisDecodedOperandImmValue__ struct {
	U uint64
	S int64
}

// Source: unknown.h:0 -> ZydisDecodedOperandImmValue_
type ZydisDecodedOperandImmValue_ struct {
	Data uint64
}

// Source: unknown.h:217 -> ZydisDecodedOperand_
type ZydisDecodedOperand_ struct {
	Id uint8
	_  [3]byte
	Visibility ZydisOperandVisibility
	Actions ZydisOperandActions
	_  [3]byte
	Encoding ZydisOperandEncoding
	Size uint16
	_  [2]byte
	ElementType ZydisElementType
	ElementSize ZydisElementSize
	ElementCount uint16
	Attributes ZydisOperandAttributes
	_  [3]byte
	Type ZydisOperandType
	_  [4]byte
	Anon1 ZydisDecodedOperand__Anon1Union
}

// Source: unknown.h:0 -> ZydisDecodedOperandReg_
type ZydisDecodedOperandReg struct {
	Value ZydisRegister
}

// Source: unknown.h:0 -> ZydisDecodedOperandMem_
type ZydisDecodedOperandMem struct {
	Type ZydisMemoryOperandType
	Segment ZydisRegister
	Base ZydisRegister
	Index ZydisRegister
	Scale uint8
	_  [7]byte
	Disp ZydisDecodedOperandMemDisp_
}

// Source: unknown.h:0 -> ZydisDecodedOperandPtr_
type ZydisDecodedOperandPtr struct {
	Segment uint16
	_  [2]byte
	Offset uint32
}

// Source: unknown.h:0 -> ZydisDecodedOperandImm_
type ZydisDecodedOperandImm struct {
	IsSigned uint8
	IsAddress uint8
	IsRelative uint8
	_  [5]byte
	Value ZydisDecodedOperandImmValue_
	Offset uint8
	Size uint8
	_  [6]byte
}

// Source: unknown.h:0 -> ZydisDecodedOperand__Anon1Union
type ZydisDecodedOperand__Anon1Union_ struct {
	Reg ZydisDecodedOperandReg
	Mem ZydisDecodedOperandMem
	Ptr ZydisDecodedOperandPtr
	Imm ZydisDecodedOperandImm
}

// Source: unknown.h:0 -> ZydisDecodedOperand__Anon1Union
type ZydisDecodedOperand__Anon1Union struct {
	Data ZydisDecodedOperandMem
}

// Source: unknown.h:403 -> ZydisAccessedFlags_
type ZydisAccessedFlags_ struct {
	Tested ZydisAccessedFlagsMask
	Modified ZydisAccessedFlagsMask
	Set0 ZydisAccessedFlagsMask
	Set1 ZydisAccessedFlagsMask
	Undefined ZydisAccessedFlagsMask
}

// Source: unknown.h:859 -> ZydisDecodedInstructionRawRex_
type ZydisDecodedInstructionRawRex_ struct {
	W uint8
	R uint8
	X uint8
	B uint8
	Offset uint8
}

// Source: unknown.h:894 -> ZydisDecodedInstructionRawRex2_
type ZydisDecodedInstructionRawRex2_ struct {
	M0 uint8
	R4 uint8
	X4 uint8
	B4 uint8
	W uint8
	R3 uint8
	X3 uint8
	B3 uint8
	Offset uint8
}

// Source: unknown.h:942 -> ZydisDecodedInstructionRawXop_
type ZydisDecodedInstructionRawXop_ struct {
	R uint8
	X uint8
	B uint8
	MMmmm uint8
	W uint8
	Vvvv uint8
	L uint8
	Pp uint8
	Offset uint8
}

// Source: unknown.h:987 -> ZydisDecodedInstructionRawVex_
type ZydisDecodedInstructionRawVex_ struct {
	R uint8
	X uint8
	B uint8
	MMmmm uint8
	W uint8
	Vvvv uint8
	L uint8
	Pp uint8
	Offset uint8
	Size uint8
}

// Source: unknown.h:1036 -> ZydisDecodedInstructionRawEvex_
type ZydisDecodedInstructionRawEvex_ struct {
	R3 uint8
	X3 uint8
	B3 uint8
	R4 uint8
	B4 uint8
	Mmm uint8
	W uint8
	Vvvv uint8
	U uint8
	X4 uint8
	Pp uint8
	Z uint8
	L2 uint8
	L uint8
	B uint8
	V4 uint8
	Aaa uint8
	ND uint8
	NF uint8
	SCC uint8
	Offset uint8
}

// Source: unknown.h:1122 -> ZydisDecodedInstructionRawMvex_
type ZydisDecodedInstructionRawMvex_ struct {
	R uint8
	X uint8
	B uint8
	R2 uint8
	Mmmm uint8
	W uint8
	Vvvv uint8
	Pp uint8
	E uint8
	SSS uint8
	V2 uint8
	Kkk uint8
	Offset uint8
}

// Source: unknown.h:1183 -> ZydisDecodedInstructionAvx_
type ZydisDecodedInstructionAvx_ struct {
	VectorLength uint16
	_  [2]byte
	Mask ZydisDecodedInstructionAvxMask_
	Broadcast ZydisDecodedInstructionAvxBroadcast_
	Rounding ZydisDecodedInstructionAvxRounding_
	Swizzle ZydisDecodedInstructionAvxSwizzle_
	Conversion ZydisDecodedInstructionAvxConversion_
	HasSae uint8
	HasEvictionHint uint8
	_  [2]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvxMask_
type ZydisDecodedInstructionAvxMask_ struct {
	Mode ZydisMaskMode
	Reg ZydisRegister
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvxBroadcast_
type ZydisDecodedInstructionAvxBroadcast_ struct {
	IsStatic uint8
	_  [3]byte
	Mode ZydisBroadcastMode
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvxRounding_
type ZydisDecodedInstructionAvxRounding_ struct {
	Mode ZydisRoundingMode
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvxSwizzle_
type ZydisDecodedInstructionAvxSwizzle_ struct {
	Mode ZydisSwizzleMode
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvxConversion_
type ZydisDecodedInstructionAvxConversion_ struct {
	Mode ZydisConversionMode
}

// Source: unknown.h:1265 -> ZydisDecodedInstructionApx_
type ZydisDecodedInstructionApx_ struct {
	UsesEgpr uint8
	HasNf uint8
	HasZu uint8
	HasPpx uint8
	Scc ZydisSourceConditionCode
	DefaultFlags ZydisDefaultFlagsValue
	_  [3]byte
}

// Source: unknown.h:1301 -> ZydisDecodedInstructionMeta_
type ZydisDecodedInstructionMeta_ struct {
	Category ZydisInstructionCategory
	IsaSet ZydisISASet
	IsaExt ZydisISAExt
	BranchType ZydisBranchType
	ExceptionClass ZydisExceptionClass
}

// Source: unknown.h:1329 -> ZydisDecodedInstructionRaw_
type ZydisDecodedInstructionRaw_ struct {
	PrefixCount uint8
	_  [3]byte
	Prefixes [15]ZydisDecodedInstructionRawPrefixes_
	Encoding2 ZydisInstructionEncoding
	Anon1 ZydisDecodedInstructionRaw__Anon1Union
	Modrm ZydisDecodedInstructionModRm_
	Sib ZydisDecodedInstructionRawSib_
	_  [3]byte
	Disp ZydisDecodedInstructionRawDisp_
	Imm [2]ZydisDecodedInstructionRawImm_
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawPrefixes_
type ZydisDecodedInstructionRawPrefixes_ struct {
	Type ZydisPrefixType
	Value uint8
	_  [3]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawRex_
type ZydisDecodedInstructionRawRex struct {
	W uint8
	R uint8
	X uint8
	B uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawRex2_
type ZydisDecodedInstructionRawRex2 struct {
	M0 uint8
	R4 uint8
	X4 uint8
	B4 uint8
	W uint8
	R3 uint8
	X3 uint8
	B3 uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawXop_
type ZydisDecodedInstructionRawXop struct {
	R uint8
	X uint8
	B uint8
	MMmmm uint8
	W uint8
	Vvvv uint8
	L uint8
	Pp uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawVex_
type ZydisDecodedInstructionRawVex struct {
	R uint8
	X uint8
	B uint8
	MMmmm uint8
	W uint8
	Vvvv uint8
	L uint8
	Pp uint8
	Offset uint8
	Size uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawEvex_
type ZydisDecodedInstructionRawEvex struct {
	R3 uint8
	X3 uint8
	B3 uint8
	R4 uint8
	B4 uint8
	Mmm uint8
	W uint8
	Vvvv uint8
	U uint8
	X4 uint8
	Pp uint8
	Z uint8
	L2 uint8
	L uint8
	B uint8
	V4 uint8
	Aaa uint8
	ND uint8
	NF uint8
	SCC uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawMvex_
type ZydisDecodedInstructionRawMvex struct {
	R uint8
	X uint8
	B uint8
	R2 uint8
	Mmmm uint8
	W uint8
	Vvvv uint8
	Pp uint8
	E uint8
	SSS uint8
	V2 uint8
	Kkk uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRaw__Anon1Union
type ZydisDecodedInstructionRaw__Anon1Union_ struct {
	Rex ZydisDecodedInstructionRawRex
	Rex2 ZydisDecodedInstructionRawRex2
	Xop ZydisDecodedInstructionRawXop
	Vex ZydisDecodedInstructionRawVex
	Evex ZydisDecodedInstructionRawEvex
	Mvex ZydisDecodedInstructionRawMvex
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRaw__Anon1Union
type ZydisDecodedInstructionRaw__Anon1Union struct {
	Data [21]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionModRm_
type ZydisDecodedInstructionModRm_ struct {
	Mod uint8
	Reg uint8
	Rm uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawSib_
type ZydisDecodedInstructionRawSib_ struct {
	Scale uint8
	Index uint8
	Base uint8
	Offset uint8
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawDisp_
type ZydisDecodedInstructionRawDisp_ struct {
	Value int64
	Size uint8
	Offset uint8
	_  [6]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawImm_
type ZydisDecodedInstructionRawImm_ struct {
	IsSigned uint8
	IsAddress uint8
	IsRelative uint8
	_  [5]byte
	Value ZydisDecodedInstructionRawImmValue_
	Size uint8
	Offset uint8
	_  [6]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawImmValue_
type ZydisDecodedInstructionRawImmValue__ struct {
	U uint64
	S int64
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRawImmValue_
type ZydisDecodedInstructionRawImmValue_ struct {
	Data uint64
}

// Source: unknown.h:1477 -> ZydisDecodedInstruction_
type ZydisDecodedInstruction_ struct {
	MachineMode ZydisMachineMode
	Mnemonic ZydisMnemonic
	Length uint8
	_  [3]byte
	Encoding ZydisInstructionEncoding
	OpcodeMap ZydisOpcodeMap
	Opcode uint8
	StackWidth uint8
	OperandWidth uint8
	AddressWidth uint8
	OperandCount uint8
	OperandCountVisible uint8
	_  [6]byte
	Attributes ZydisInstructionAttributes
	CpuFlags *ZydisAccessedFlags
	FpuFlags *ZydisAccessedFlags
	Avx ZydisDecodedInstructionAvx
	Apx ZydisDecodedInstructionApx
	Meta ZydisDecodedInstructionMeta
	_  [4]byte
	Raw ZydisDecodedInstructionRaw
}

// Source: unknown.h:0 -> ZydisDecodedInstructionAvx_
type ZydisDecodedInstructionAvx struct {
	VectorLength uint16
	_  [2]byte
	Mask ZydisDecodedInstructionAvxMask_
	Broadcast ZydisDecodedInstructionAvxBroadcast_
	Rounding ZydisDecodedInstructionAvxRounding_
	Swizzle ZydisDecodedInstructionAvxSwizzle_
	Conversion ZydisDecodedInstructionAvxConversion_
	HasSae uint8
	HasEvictionHint uint8
	_  [2]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionApx_
type ZydisDecodedInstructionApx struct {
	UsesEgpr uint8
	HasNf uint8
	HasZu uint8
	HasPpx uint8
	Scc ZydisSourceConditionCode
	DefaultFlags ZydisDefaultFlagsValue
	_  [3]byte
}

// Source: unknown.h:0 -> ZydisDecodedInstructionMeta_
type ZydisDecodedInstructionMeta struct {
	Category ZydisInstructionCategory
	IsaSet ZydisISASet
	IsaExt ZydisISAExt
	BranchType ZydisBranchType
	ExceptionClass ZydisExceptionClass
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRaw_
type ZydisDecodedInstructionRaw struct {
	PrefixCount uint8
	_  [3]byte
	Prefixes [15]ZydisDecodedInstructionRawPrefixes_
	Encoding2 ZydisInstructionEncoding
	Anon1 ZydisDecodedInstructionRaw_Anon1Union
	Modrm ZydisDecodedInstructionModRm_
	Sib ZydisDecodedInstructionRawSib_
	_  [3]byte
	Disp ZydisDecodedInstructionRawDisp_
	Imm [2]ZydisDecodedInstructionRawImm_
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRaw_Anon1Union
type ZydisDecodedInstructionRaw_Anon1Union_ struct {
	Rex ZydisDecodedInstructionRawRex
	Rex2 ZydisDecodedInstructionRawRex2
	Xop ZydisDecodedInstructionRawXop
	Vex ZydisDecodedInstructionRawVex
	Evex ZydisDecodedInstructionRawEvex
	Mvex ZydisDecodedInstructionRawMvex
}

// Source: unknown.h:0 -> ZydisDecodedInstructionRaw_Anon1Union
type ZydisDecodedInstructionRaw_Anon1Union struct {
	Data [21]byte
}

// Source: unknown.h:1582 -> ZydisDecoderContext_
type ZydisDecoderContext_ struct {
	Definition unsafe.Pointer
	EoszIndex uint8
	EaszIndex uint8
	VectorUnified ZydisDecoderContext__Anon1
	RegInfo ZydisDecoderContext__Anon2
	Evex ZydisDecoderContext__Anon3
	Mvex ZydisDecoderContext__Anon4
	Cd8Scale uint8
}

// Source: unknown.h:0 -> ZydisDecoderContext__Anon1
type ZydisDecoderContext__Anon1 struct {
	W uint8
	R3 uint8
	R4 uint8
	X3 uint8
	X4 uint8
	B3 uint8
	B4 uint8
	L uint8
	LL uint8
	V4 uint8
	Vvvv uint8
	Mask uint8
}

// Source: unknown.h:0 -> ZydisDecoderContext__Anon2
type ZydisDecoderContext__Anon2 struct {
	IsModReg uint8
	IdReg uint8
	IdRm uint8
	IdNdsndd uint8
	IdBase uint8
	IdIndex uint8
}

// Source: unknown.h:0 -> ZydisDecoderContext__Anon3
type ZydisDecoderContext__Anon3 struct {
	TupleType uint8
	ElementSize uint8
}

// Source: unknown.h:0 -> ZydisDecoderContext__Anon4
type ZydisDecoderContext__Anon4 struct {
	Functionality uint8
}

// Source: unknown.h:190 -> ZydisDecoder_
type ZydisDecoder_ struct {
	MachineMode ZydisMachineMode
	StackWidth ZydisStackWidth
	DecoderMode uint32
}

// Source: unknown.h:184 -> ZydisEncoderOperand_
type ZydisEncoderOperand_ struct {
	Type ZydisOperandType
	Reg ZydisEncoderOperandReg_
	_  [4]byte
	Mem ZydisEncoderOperandMem_
	Ptr ZydisEncoderOperandPtr_
	Imm ZydisEncoderOperandImm_
}

// Source: unknown.h:0 -> ZydisEncoderOperandReg_
type ZydisEncoderOperandReg_ struct {
	Value ZydisRegister
	Is4 uint8
	_  [3]byte
}

// Source: unknown.h:0 -> ZydisEncoderOperandMem_
type ZydisEncoderOperandMem_ struct {
	Base ZydisRegister
	Index ZydisRegister
	Scale uint8
	_  [7]byte
	Displacement int64
	Size uint16
	_  [6]byte
}

// Source: unknown.h:0 -> ZydisEncoderOperandPtr_
type ZydisEncoderOperandPtr_ struct {
	Segment uint16
	_  [2]byte
	Offset uint32
}

// Source: unknown.h:0 -> ZydisEncoderOperandImm_
type ZydisEncoderOperandImm__ struct {
	U uint64
	S int64
}

// Source: unknown.h:0 -> ZydisEncoderOperandImm_
type ZydisEncoderOperandImm_ struct {
	Data uint64
}

// Source: unknown.h:269 -> ZydisEncoderRequest_
type ZydisEncoderRequest_ struct {
	MachineMode ZydisMachineMode
	AllowedEncodings ZydisEncodableEncoding
	Mnemonic ZydisMnemonic
	_  [4]byte
	Prefixes ZydisInstructionAttributes
	BranchType ZydisBranchType
	BranchWidth ZydisBranchWidth
	AddressSizeHint ZydisAddressSizeHint
	OperandSizeHint ZydisOperandSizeHint
	OperandCount uint8
	_  [7]byte
	Operands [5]ZydisEncoderOperand
	Evex ZydisEncoderRequestEvexFeatures_
	Mvex ZydisEncoderRequestMvexFeatures_
}

// Source: unknown.h:0 -> ZydisEncoderRequestEvexFeatures_
type ZydisEncoderRequestEvexFeatures_ struct {
	Broadcast ZydisBroadcastMode
	Rounding ZydisRoundingMode
	Sae uint8
	ZeroingMask uint8
	NoFlags uint8
	DefaultFlags ZydisDefaultFlagsValue
}

// Source: unknown.h:0 -> ZydisEncoderRequestMvexFeatures_
type ZydisEncoderRequestMvexFeatures_ struct {
	Broadcast ZydisBroadcastMode
	Conversion ZydisConversionMode
	Rounding ZydisRoundingMode
	Swizzle ZydisSwizzleMode
	Sae uint8
	EvictionHint uint8
	_  [2]byte
}

// Source: unknown.h:46 -> ZyanAllocator_
type ZyanAllocator_ struct {
	Allocate ZyanAllocatorAllocate
	Reallocate ZyanAllocatorAllocate
	Deallocate ZyanAllocatorDeallocate
}

// Source: unknown.h:75 -> ZyanVector_
type ZyanVector_ struct {
	Allocator *ZyanAllocator
	GrowthFactor uint8
	ShrinkThreshold uint8
	_  [6]byte
	Size uintptr
	Capacity uintptr
	ElementSize uintptr
	Destructor ZyanMemberProcedure
	Data unsafe.Pointer
}

// Source: unknown.h:97 -> ZyanString_
type ZyanString_ struct {
	Flags ZyanStringFlags
	_  [7]byte
	Vector ZyanVector
}

// Source: unknown.h:0 -> ZyanVector_
type ZyanVector struct {
	Allocator *ZyanAllocator
	GrowthFactor uint8
	ShrinkThreshold uint8
	_  [6]byte
	Size uintptr
	Capacity uintptr
	ElementSize uintptr
	Destructor ZyanMemberProcedure
	Data unsafe.Pointer
}

// Source: unknown.h:131 -> ZyanStringView_
type ZyanStringView_ struct {
	String ZyanString
}

// Source: unknown.h:0 -> ZyanString
type ZyanString struct {
	Flags ZyanStringFlags
	_  [7]byte
	Vector ZyanVector
}

// Source: unknown.h:138 -> ZydisFormatterToken_
type ZydisFormatterToken_ struct {
	Type ZydisTokenType
	Next uint8
}

// Source: unknown.h:167 -> ZydisFormatterBuffer_
type ZydisFormatterBuffer_ struct {
	IsTokenList uint8
	_  [7]byte
	Capacity uintptr
	String ZyanString
}

// Source: unknown.h:681 -> ZydisFormatter_
type ZydisFormatter_ struct {
	Style ZydisFormatterStyle
	ForceMemorySize uint8
	ForceMemorySegment uint8
	ForceMemoryScale uint8
	ForceRelativeBranches uint8
	ForceRelativeRiprel uint8
	PrintBranchSize uint8
	DetailedPrefixes uint8
	_  [1]byte
	AddrBase ZydisNumericBase
	AddrSignedness ZydisSignedness
	AddrPaddingAbsolute ZydisPadding
	AddrPaddingRelative ZydisPadding
	DispBase ZydisNumericBase
	DispSignedness ZydisSignedness
	DispPadding ZydisPadding
	ImmBase ZydisNumericBase
	ImmSignedness ZydisSignedness
	ImmPadding ZydisPadding
	CasePrefixes int32
	CaseMnemonic int32
	CaseRegisters int32
	CaseTypecasts int32
	CaseDecorators int32
	HexUppercase uint8
	HexForceLeadingNumber uint8
	_  [6]byte
	NumberFormat ZydisFormatter__Anon1Elem
	DecoApxNfUseSuffix uint8
	DecoApxDfvUseImmediate uint8
	_  [6]byte
	FuncPreInstruction ZydisFormatterFunc
	FuncPostInstruction ZydisFormatterFunc
	FuncFormatInstruction ZydisFormatterFunc
	FuncPreOperand ZydisFormatterFunc
	FuncPostOperand ZydisFormatterFunc
	FuncFormatOperandReg ZydisFormatterFunc
	FuncFormatOperandMem ZydisFormatterFunc
	FuncFormatOperandPtr ZydisFormatterFunc
	FuncFormatOperandImm ZydisFormatterFunc
	FuncPrintMnemonic ZydisFormatterFunc
	FuncPrintRegister ZydisFormatterRegisterFunc
	FuncPrintAddressAbs ZydisFormatterFunc
	FuncPrintAddressRel ZydisFormatterFunc
	FuncPrintDisp ZydisFormatterFunc
	FuncPrintImm ZydisFormatterFunc
	FuncPrintTypecast ZydisFormatterFunc
	FuncPrintSegment ZydisFormatterFunc
	FuncPrintPrefixes ZydisFormatterFunc
	FuncPrintDecorator ZydisFormatterDecoratorFunc
}

// Source: unknown.h:0 -> ZydisFormatter__Anon1Elem
type ZydisFormatter__Anon1Elem struct {
	String *ZyanStringView
	StringData ZyanStringView
	Buffer [11]int8
	_  [5]byte
}

// Source: unknown.h:0 -> ZyanStringView_
type ZyanStringView struct {
	String ZyanString
}

// Source: unknown.h:686 -> ZydisFormatterContext_
type ZydisFormatterContext_ struct {
	Instruction *ZydisDecodedInstruction
	Operands *ZydisDecodedOperand
	Operand *ZydisDecodedOperand
	RuntimeAddress uint64
	UserData unsafe.Pointer
}

// Source: unknown.h:133 -> ZydisInstructionSegments_
type ZydisInstructionSegments_ struct {
	Count uint8
	_  [3]byte
	Segments ZydisInstructionSegments__Anon1Elem
}

// Source: unknown.h:0 -> ZydisInstructionSegments__Anon1Elem
type ZydisInstructionSegments__Anon1Elem struct {
	Type ZydisInstructionSegment
	Offset uint8
	Size uint8
	_  [2]byte
}

// Source: unknown.h:51 -> ZydisDisassembledInstruction_
type ZydisDisassembledInstruction_ struct {
	RuntimeAddress uint64
	Info ZydisDecodedInstruction
	Operands [10]ZydisDecodedOperand
	Text [96]int8
}

// Source: unknown.h:0 -> ZydisDecodedInstruction_
type ZydisDecodedInstruction struct {
	MachineMode ZydisMachineMode
	Mnemonic ZydisMnemonic
	Length uint8
	_  [3]byte
	Encoding ZydisInstructionEncoding
	OpcodeMap ZydisOpcodeMap
	Opcode uint8
	StackWidth uint8
	OperandWidth uint8
	AddressWidth uint8
	OperandCount uint8
	OperandCountVisible uint8
	_  [6]byte
	Attributes ZydisInstructionAttributes
	CpuFlags *ZydisAccessedFlags
	FpuFlags *ZydisAccessedFlags
	Avx ZydisDecodedInstructionAvx
	Apx ZydisDecodedInstructionApx
	Meta ZydisDecodedInstructionMeta
	_  [4]byte
	Raw ZydisDecodedInstructionRaw
}

