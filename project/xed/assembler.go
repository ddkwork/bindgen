package xed

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type Assembler struct {
	x           *Xed
	iclassMap   map[string]Xed_iclass_enum_t
	regMap      map[string]Xed_reg_enum_t
	regSet      map[string]bool
	brdispTable []Xed_encoder_operand_type_t
}

func NewAssembler(x *Xed) *Assembler {
	a := &Assembler{x: x}
	a.buildIclassMap()
	a.buildRegMap()
	a.buildBrdispTable()
	return a
}

func cstring(s *int8) string {
	if s == nil {
		return "<nil>"
	}
	return goStringInt8((*[256]int8)(unsafe.Pointer(s))[:])
}

func goStringInt8(buf []int8) string {
	for i, b := range buf {
		if b == 0 {
			return string((*[256]byte)(unsafe.Pointer(&buf[0]))[:i])
		}
	}
	return string((*[256]byte)(unsafe.Pointer(&buf[0]))[:])
}

func goString(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf)
}

func xedState(mode Xed_machine_mode_enum_t) Xed_state_t {
	s := Xed_state_t{}
	switch mode {
	case XedMachineModeLegacy16:
		s.Mmode = mode
		s.StackAddrWidth = XedAddressWidth16b
	case XedMachineModeLegacy32:
		s.Mmode = mode
		s.StackAddrWidth = XedAddressWidth32b
	case XedMachineModeLong64:
		s.Mmode = mode
		s.StackAddrWidth = XedAddressWidth64b
	default:
		s.Mmode = mode
		s.StackAddrWidth = XedAddressWidth32b
	}
	return s
}

func (a *Assembler) buildIclassMap() {
	a.iclassMap = make(map[string]Xed_iclass_enum_t, 2000)
	last := a.x.IclassEnumTLast()
	for i := Xed_iclass_enum_t(1); i < last; i++ {
		s := a.x.IclassEnumT2str(i)
		if s != nil {
			name := strings.ToUpper(cstring(s))
			a.iclassMap[name] = i
		}
	}
}

func (a *Assembler) buildRegMap() {
	a.regMap = make(map[string]Xed_reg_enum_t, 400)
	a.regSet = make(map[string]bool, 400)
	for i := Xed_reg_enum_t(1); i < 366; i++ {
		s := a.x.RegEnumT2str(i)
		if s != nil {
			name := strings.ToUpper(cstring(s))
			a.regMap[name] = i
			a.regSet[name] = true
		}
	}
}

func (a *Assembler) buildBrdispTable() {
	a.brdispTable = make([]Xed_encoder_operand_type_t, XedIclassLast)
	last := a.x.IclassEnumTLast()
	for ic := Xed_iclass_enum_t(1); ic < last; ic++ {
		s := a.x.IclassEnumT2str(ic)
		if s == nil {
			continue
		}
		name := strings.ToUpper(cstring(s))
		if isRelbrIclass(name) {
			a.brdispTable[ic] = XedEncoderOperandTypeRelBrdisp
		} else if isAbsbrIclass(name) {
			a.brdispTable[ic] = XedEncoderOperandTypeAbsBrdisp
		}
	}
}

func isRelbrIclass(name string) bool {
	switch name {
	case "JMP", "CALL_NEAR", "LOOP", "LOOPE", "LOOPNE",
		"JE", "JZ", "JNE", "JNZ", "JA", "JNBE", "JAE", "JNB", "JB", "JNAE", "JBE", "JNA",
		"JG", "JNLE", "JGE", "JNL", "JL", "JNGE", "JLE", "JNG",
		"JC", "JNC", "JO", "JNO", "JS", "JNS", "JP", "JPE", "JNP", "JPO",
		"JECXZ", "JCXZ", "JRCXZ",
		"XBEGIN":
		return true
	}
	return false
}

func isAbsbrIclass(name string) bool {
	switch name {
	case "JMP_FAR", "CALL_FAR", "INT", "INTO", "INT1", "INT3",
		"SYSCALL", "SYSRET", "SYSENTER", "SYSEXIT",
		"IRET", "IRETD", "IRETQ":
		return true
	}
	return false
}

func (a *Assembler) lookupIclass(name string) Xed_iclass_enum_t {
	if v, ok := a.iclassMap[strings.ToUpper(name)]; ok {
		return v
	}
	return XedIclassInvalid
}

func (a *Assembler) lookupReg(name string) Xed_reg_enum_t {
	if v, ok := a.regMap[strings.ToUpper(name)]; ok {
		return v
	}
	return XedRegInvalid
}

func (a *Assembler) isReg(s string) bool {
	return a.regSet[strings.ToUpper(s)]
}

func (a *Assembler) hasRelbr(iclass Xed_iclass_enum_t) bool {
	if int(iclass) >= len(a.brdispTable) {
		return false
	}
	return a.brdispTable[iclass] == XedEncoderOperandTypeRelBrdisp
}

func (a *Assembler) hasAbsbr(iclass Xed_iclass_enum_t) bool {
	if int(iclass) >= len(a.brdispTable) {
		return false
	}
	return a.brdispTable[iclass] == XedEncoderOperandTypeAbsBrdisp
}

type parsedOperand struct {
	typ opndType
	reg Xed_reg_enum_t
	imm int64
	mem memOperand
	raw string
}

type opndType int

const (
	opndReg opndType = iota
	opndImm
	opndMem
)

type memOperand struct {
	seg   Xed_reg_enum_t
	base  Xed_reg_enum_t
	index Xed_reg_enum_t
	scale uint32
	disp  int64
	width uint32
}

type parsedLine struct {
	iclass    Xed_iclass_enum_t
	prefixes  []string
	operands  []parsedOperand
	mode      int
	seenLock  bool
	seenRep   bool
	seenRepne bool
}

func (a *Assembler) Assemble(line string, mode int) ([]byte, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, fmt.Errorf("empty input")
	}
	p, err := a.parseLine(line, mode)
	if err != nil {
		return nil, err
	}
	return a.encodeLine(p)
}

func (a *Assembler) parseLine(line string, mode int) (*parsedLine, error) {
	p := &parsedLine{mode: mode}
	tokens := tokenize(line)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens in line")
	}
	idx := 0
	for idx < len(tokens) {
		upperTok := strings.ToUpper(tokens[idx])
		if upperTok == "LOCK" {
			p.seenLock = true
			p.prefixes = append(p.prefixes, "LOCK")
			idx++
		} else if upperTok == "REP" || upperTok == "REPE" {
			p.seenRep = true
			p.prefixes = append(p.prefixes, upperTok)
			idx++
		} else if upperTok == "REPNE" {
			p.seenRepne = true
			p.prefixes = append(p.prefixes, "REPNE")
			idx++
		} else {
			break
		}
	}
	if idx >= len(tokens) {
		return nil, fmt.Errorf("no instruction mnemonic")
	}
	mnemonic := tokens[idx]
	idx++
	a.reviseMnemonic(p, mnemonic)
	if p.iclass == XedIclassInvalid {
		return nil, fmt.Errorf("unknown instruction: %s", mnemonic)
	}
	if idx < len(tokens) {
		rest := strings.Join(tokens[idx:], " ")
		opnds, err := a.parseOperands(rest)
		if err != nil {
			return nil, err
		}
		p.operands = opnds
	}
	return p, nil
}

var mnemonicAliases = map[string]string{
	"RET":    "RET_NEAR",
	"RETF":   "RET_FAR",
	"CALL":   "CALL_NEAR",
	"CALLF":  "CALL_FAR",
	"JZ":     "JE",
	"JNZ":    "JNE",
	"JB":     "JNAE",
	"JNB":    "JAE",
	"JA":     "JNBE",
	"JNA":    "JBE",
	"JG":     "JNLE",
	"JGE":    "JNL",
	"JL":     "JNGE",
	"JLE":    "JNG",
	"CMOVZ":  "CMOVE",
	"CMOVNZ": "CMOVNE",
	"SETZ":   "SETE",
	"SETNZ":  "SETNE",
}

func (a *Assembler) reviseMnemonic(p *parsedLine, mnemonic string) {
	upper := strings.ToUpper(mnemonic)
	if alias, ok := mnemonicAliases[upper]; ok {
		upper = alias
	}
	if p.seenLock {
		if v := a.lookupIclass(upper + "_LOCK"); v != XedIclassInvalid {
			p.iclass = v
			return
		}
	}
	if p.seenRepne {
		if v := a.lookupIclass("REPNE_" + upper); v != XedIclassInvalid {
			p.iclass = v
			return
		}
	}
	if p.seenRep {
		if v := a.lookupIclass("REPE_" + upper); v != XedIclassInvalid {
			p.iclass = v
			return
		}
		if v := a.lookupIclass("REP_" + upper); v != XedIclassInvalid {
			p.iclass = v
			return
		}
	}
	p.iclass = a.lookupIclass(upper)
}

func (a *Assembler) parseOperands(s string) ([]parsedOperand, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	parts := splitOperands(s)
	var ops []parsedOperand
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		op, err := a.parseSingleOperand(part)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func splitOperands(s string) []string {
	var result []string
	depth := 0
	current := strings.Builder{}
	for _, ch := range s {
		if ch == '[' {
			depth++
			current.WriteRune(ch)
		} else if ch == ']' {
			depth--
			current.WriteRune(ch)
		} else if ch == ',' && depth == 0 {
			result = append(result, current.String())
			current.Reset()
		} else {
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		result = append(result, current.String())
	}
	return result
}

func (a *Assembler) parseSingleOperand(s string) (parsedOperand, error) {
	s = strings.TrimSpace(s)
	upper := strings.ToUpper(s)
	if strings.HasPrefix(upper, "[") || strings.Contains(upper, "PTR") {
		return a.parseMemOperand(s)
	}
	if reg := a.lookupReg(upper); reg != XedRegInvalid {
		return parsedOperand{typ: opndReg, reg: reg, raw: s}, nil
	}
	imm, err := parseImmediate(s)
	if err != nil {
		return parsedOperand{}, fmt.Errorf("cannot parse operand %q: %w", s, err)
	}
	return parsedOperand{typ: opndImm, imm: imm, raw: s}, nil
}

func (a *Assembler) parseMemOperand(s string) (parsedOperand, error) {
	s = strings.TrimSpace(s)
	upper := strings.ToUpper(s)
	var width uint32
	ptrIdx := strings.Index(upper, "PTR")
	if ptrIdx > 0 {
		widthStr := strings.TrimSpace(upper[:ptrIdx])
		switch widthStr {
		case "BYTE":
			width = 8
		case "WORD":
			width = 16
		case "DWORD":
			width = 32
		case "QWORD":
			width = 64
		case "FWORD":
			width = 48
		case "TBYTE":
			width = 80
		case "XMMWORD", "OWORD":
			width = 128
		case "YMMWORD":
			width = 256
		case "ZMMWORD":
			width = 512
		}
		s = s[ptrIdx+3:]
		s = strings.TrimSpace(s)
	}
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimSpace(s)
	mem := memOperand{}
	parts := strings.Split(s, "+")
	var regParts []string
	var dispStr string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if a.isReg(part) {
			regParts = append(regParts, part)
		} else if strings.Contains(part, "*") {
			scaleParts := strings.SplitN(part, "*", 2)
			idxStr := strings.TrimSpace(scaleParts[0])
			scaleStr := strings.TrimSpace(scaleParts[1])
			if reg := a.lookupReg(idxStr); reg != XedRegInvalid {
				mem.index = reg
			}
			if sc, err := strconv.ParseUint(scaleStr, 10, 32); err == nil {
				mem.scale = uint32(sc)
			}
		} else {
			dispStr = part
		}
	}
	for _, rp := range regParts {
		reg := a.lookupReg(rp)
		if reg == XedRegInvalid {
			continue
		}
		rc := a.x.RegClass(reg)
		if rc == XedRegClassSr {
			mem.seg = reg
		} else if mem.base == XedRegInvalid {
			mem.base = reg
		} else {
			mem.index = reg
		}
	}
	if dispStr != "" {
		disp, err := parseImmediate(dispStr)
		if err != nil {
			return parsedOperand{}, fmt.Errorf("cannot parse displacement %q: %w", dispStr, err)
		}
		mem.disp = disp
	}
	if width == 0 {
		width = a.guessMemWidth(mem.base, mem.index)
	}
	mem.width = width
	return parsedOperand{typ: opndMem, mem: mem, raw: s}, nil
}

func (a *Assembler) guessMemWidth(base, index Xed_reg_enum_t) uint32 {
	var rc Xed_reg_class_enum_t
	if base != XedRegInvalid {
		rc = a.x.RegClass(base)
	} else if index != XedRegInvalid {
		rc = a.x.RegClass(index)
	}
	grc := Xed_reg_class_enum_t(0)
	if base != XedRegInvalid {
		grc = a.x.GprRegClass(base)
	} else if index != XedRegInvalid {
		grc = a.x.GprRegClass(index)
	}
	switch grc {
	case XedRegClassGpr64:
		return 64
	case XedRegClassGpr32:
		return 32
	case XedRegClassGpr16:
		return 16
	}
	_ = rc
	return 32
}

func parseImmediate(s string) (int64, error) {
	s = strings.TrimSpace(s)
	negative := false
	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	} else if strings.HasPrefix(s, "+") {
		s = s[1:]
	}
	s = strings.TrimSpace(s)
	var val uint64
	var err error
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		val, err = strconv.ParseUint(s[2:], 16, 64)
	} else if strings.HasPrefix(s, "0b") || strings.HasPrefix(s, "0B") {
		val, err = strconv.ParseUint(s[2:], 2, 64)
	} else {
		val, err = strconv.ParseUint(s, 10, 64)
	}
	if err != nil {
		return 0, fmt.Errorf("invalid immediate: %q", s)
	}
	if negative {
		return -int64(val), nil
	}
	return int64(val), nil
}

func (a *Assembler) encodeLine(p *parsedLine) ([]byte, error) {
	dstate := xedStateForMode(p.mode)
	var eosz uint32
	var ops []encOp
	hasImm0 := false
	for _, opnd := range p.operands {
		switch opnd.typ {
		case opndReg:
			if opnd.reg == XedRegInvalid {
				return nil, fmt.Errorf("invalid register")
			}
			ops = append(ops, xedEncOpReg(opnd.reg))
			eosz = updateEosz(eosz, opnd.reg, a.x)
		case opndImm:
			nbits := a.getConstantWidth(opnd.raw, opnd.imm)
			if nbits == 0 {
				nbits = 32
			}
			if a.hasRelbr(p.iclass) {
				ops = append(ops, xedEncOpRelBr(uint64(opnd.imm), nbits))
			} else if a.hasAbsbr(p.iclass) {
				ops = append(ops, xedEncOpAbsBr(uint64(opnd.imm), nbits))
			} else if !hasImm0 {
				ops = append(ops, xedEncOpImm0(uint64(opnd.imm), nbits))
				hasImm0 = true
			} else {
				ops = append(ops, xedEncOpImm1(uint8(opnd.imm)))
			}
		case opndMem:
			dispBits := getDispBits(opnd.mem.disp)
			ops = append(ops, xedEncOpMem(opnd.mem.seg, opnd.mem.base, opnd.mem.index, opnd.mem.scale, opnd.mem.disp, dispBits, opnd.mem.width))
		}
	}
	if eosz == 0 {
		eosz = defaultEosz(p.mode)
	}
	inst := xedBuildInst(dstate, p.iclass, eosz, ops...)
	if p.seenRep {
		xedSetRep(&inst)
	}
	if p.seenRepne {
		xedSetRepne(&inst)
	}
	result, err := encodeInstruction(a.x, inst)
	if err != nil && eosz != 64 && p.mode == 64 {
		altEosz := uint32(64)
		altInst := xedBuildInst(dstate, p.iclass, altEosz, ops...)
		if p.seenRep {
			xedSetRep(&altInst)
		}
		if p.seenRepne {
			xedSetRepne(&altInst)
		}
		if altResult, altErr := encodeInstruction(a.x, altInst); altErr == nil {
			return altResult, nil
		}
	}
	return result, err
}

func xedStateForMode(mode int) Xed_state_t {
	switch mode {
	case 16:
		return Xed_state_t{Mmode: XedMachineModeLegacy16, StackAddrWidth: XedAddressWidth16b}
	case 32:
		return Xed_state_t{Mmode: XedMachineModeLegacy32, StackAddrWidth: XedAddressWidth32b}
	case 64:
		return Xed_state_t{Mmode: XedMachineModeLong64, StackAddrWidth: XedAddressWidth64b}
	default:
		return Xed_state_t{Mmode: XedMachineModeLong64, StackAddrWidth: XedAddressWidth64b}
	}
}

func updateEosz(current uint32, reg Xed_reg_enum_t, x *Xed) uint32 {
	grc := x.GprRegClass(reg)
	switch grc {
	case XedRegClassGpr16:
		if current < 16 {
			return 16
		}
	case XedRegClassGpr32:
		if current < 32 {
			return 32
		}
	case XedRegClassGpr64:
		return 64
	}
	return current
}

func defaultEosz(mode int) uint32 {
	switch mode {
	case 16:
		return 16
	default:
		return 32
	}
}

func (a *Assembler) getConstantWidth(text string, val int64) uint32 {
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "-") || strings.HasPrefix(text, "+") {
		return a.getNbitsSigned(val)
	}
	if hasPaddingZeroes(text) {
		return 4 * uint32(countNibbles(text))
	}
	return a.getNbitsUnsigned(uint64(val))
}

func hasPaddingZeroes(s string) bool {
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}
	if len(s) > 2 && (s[0] == '0' && (s[1] == 'x' || s[1] == 'X')) {
		s = s[2:]
	}
	return len(s) > 0 && s[0] == '0'
}

func countNibbles(s string) int {
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}
	if len(s) > 2 && (s[0] == '0' && (s[1] == 'x' || s[1] == 'X')) {
		s = s[2:]
	}
	return len(s)
}

func (a *Assembler) getNbitsSigned(val int64) uint32 {
	legalWidths := Xed_uint8_t(1 | 2 | 4 | 8)
	nbytes := a.x.ShortestWidthSigned(Xed_int64_t(val), legalWidths)
	return uint32(nbytes) * 8
}

func (a *Assembler) getNbitsUnsigned(val uint64) uint32 {
	legalWidths := Xed_uint8_t(1 | 2 | 4 | 8)
	nbytes := a.x.ShortestWidthUnsigned(Xed_uint64_t(val), legalWidths)
	return uint32(nbytes) * 8
}

func getDispBits(disp int64) uint32 {
	if disp == 0 {
		return 0
	}
	if disp >= -128 && disp <= 127 {
		return 8
	}
	return 32
}

func encodeInstruction(x *Xed, inst Xed_encoder_instruction_t) ([]byte, error) {
	var encReq Xed_encoder_request_t
	x.EncoderRequestZeroSetMode(&encReq, &inst.Mode)
	ok := x.ConvertToEncoderRequest(&encReq, &inst)
	if ok == 0 {
		return nil, fmt.Errorf("conversion to encoder request failed")
	}
	itext := make([]byte, XedMaxInstructionBytes)
	ilen := uint32(len(itext))
	var olen uint32
	err := x.Encode(&encReq, &itext[0], ilen, &olen)
	if err != XedErrorNone {
		return nil, fmt.Errorf("encode error: %s (%d)", cstring(x.ErrorEnumT2str(err)), err)
	}
	if olen == 0 || olen > XedMaxInstructionBytes {
		return nil, fmt.Errorf("bad encoded length: %d", olen)
	}
	return itext[:olen], nil
}

func tokenize(s string) []string {
	var tokens []string
	current := strings.Builder{}
	inBracket := false
	for _, ch := range s {
		if ch == '[' {
			inBracket = true
			current.WriteRune(ch)
		} else if ch == ']' {
			inBracket = false
			current.WriteRune(ch)
		} else if !inBracket && (ch == ' ' || ch == '\t') {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		} else {
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

func (a *Assembler) Disassemble(bytes []byte, mode int) (string, error) {
	var xedd Xed_decoded_inst_t
	dstate := xedStateForMode(mode)
	a.x.DecodedInstZeroSetMode(&xedd, &dstate)
	err := a.x.Decode(&xedd, &bytes[0], uint32(len(bytes)))
	if err != XedErrorNone {
		return "", fmt.Errorf("decode error: %s (%d)", cstring(a.x.ErrorEnumT2str(err)), err)
	}
	buf := make([]int8, 256)
	fmtOk := a.x.FormatContext(XedSyntaxIntel, &xedd, &buf[0], int32(len(buf)), 0, unsafe.Pointer(uintptr(0)), Xed_disassembly_callback_fn_t(nil))
	if fmtOk == 0 {
		return "", fmt.Errorf("format error")
	}
	return goStringInt8(buf), nil
}

type encOp struct {
	typ       Xed_encoder_operand_type_t
	u         uint64
	mem       Xed_memop_t
	widthBits uint32
}

func xedEncOpReg(reg Xed_reg_enum_t) encOp {
	return encOp{typ: XedEncoderOperandTypeReg, u: uint64(reg)}
}

func xedEncOpImm0(val uint64, bits uint32) encOp {
	return encOp{typ: XedEncoderOperandTypeImm0, u: val, widthBits: bits}
}

func xedEncOpSimm0(val int64, bits uint32) encOp {
	v := uint64(val)
	if val < 0 {
		v = uint64(val)
	}
	return encOp{typ: XedEncoderOperandTypeSimm0, u: v, widthBits: bits}
}

func xedEncOpRelBr(val uint64, bits uint32) encOp {
	return encOp{typ: XedEncoderOperandTypeRelBrdisp, u: val, widthBits: bits}
}

func xedEncOpAbsBr(val uint64, bits uint32) encOp {
	return encOp{typ: XedEncoderOperandTypeAbsBrdisp, u: val, widthBits: bits}
}

func xedEncOpImm1(val uint8) encOp {
	return encOp{typ: XedEncoderOperandTypeImm1, u: uint64(val)}
}

func xedEncOpMem(seg, base, index Xed_reg_enum_t, scale uint32, dispVal int64, dispBits uint32, widthBits uint32) encOp {
	m := encOp{typ: XedEncoderOperandTypeMem, widthBits: widthBits}
	m.mem.Seg = seg
	m.mem.Base = base
	m.mem.Index = index
	m.mem.Scale = scale
	m.mem.Disp.Displacement = dispVal
	m.mem.Disp.DisplacementBits = dispBits
	return m
}

func xedEncOpMemBd(base Xed_reg_enum_t, dispVal int64, dispBits uint32, widthBits uint32) encOp {
	return xedEncOpMem(XedRegInvalid, base, XedRegInvalid, 0, dispVal, dispBits, widthBits)
}

func xedEncOpPtr(val uint64, bits uint32) encOp {
	return encOp{typ: XedEncoderOperandTypePtr, u: val, widthBits: bits}
}

func xedEncOpOther(name Xed_operand_enum_t, val uint32) encOp {
	return encOp{typ: XedEncoderOperandTypeOther, u: uint64(uint64(name)<<32 | uint64(val))}
}

func xedBuildInst(dstate Xed_state_t, iclass Xed_iclass_enum_t, eosz uint32, ops ...encOp) Xed_encoder_instruction_t {
	inst := Xed_encoder_instruction_t{}
	inst.Mode = dstate
	inst.Iclass = iclass
	inst.EffectiveOperandWidth = eosz
	nops := len(ops)
	if nops > 15 {
		nops = 15
	}
	inst.Noperands = uint32(nops)
	opArray := (*[8]Xed_encoder_instruction_t_Anon1Elem)(unsafe.Pointer(&inst.Operands))
	for i, op := range ops {
		if i >= nops {
			break
		}
		opArray[i].Type = op.typ
		if op.typ == XedEncoderOperandTypeMem {
			opArray[i].U.Data = op.mem
		} else {
			*(*uint64)(unsafe.Pointer(&opArray[i].U.Data)) = op.u
		}
		opArray[i].WidthBits = op.widthBits
	}
	return inst
}

func xedSetAddr(inst *Xed_encoder_instruction_t, easz uint32) {
	inst.EffectiveAddressWidth = easz
}

func xedSetRep(inst *Xed_encoder_instruction_t) {
	inst.Prefixes.SetRep(1)
}

func xedSetRepne(inst *Xed_encoder_instruction_t) {
	inst.Prefixes.SetRepne(1)
}
