package sdk

type assembler struct{}

func (a *assembler) Assemble(addr uint, dest *byte, size *int, instruction string) {
	response := safeGet("_scriptapi_assembler.h/Assemble", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *assembler) AssembleEx(addr uint, dest *byte, size *int, instruction string, error *int8) {
	response := safeGet("_scriptapi_assembler.h/AssembleEx", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *assembler) AssembleMem(addr uint, instruction string) {
	response := safeGet("_scriptapi_assembler.h/AssembleMem", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *assembler) AssembleMemEx(addr uint, instruction string, size *int, error *int8, fillnop bool) {
	response := safeGet("_scriptapi_assembler.h/AssembleMemEx", map[string]string{})
	if len(response) == 0 {
		return
	}
}
