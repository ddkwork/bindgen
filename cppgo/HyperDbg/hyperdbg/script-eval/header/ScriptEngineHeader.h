
#pragma once
BOOLEAN SetRegValue(PGUEST_REGS GuestRegs, UINT32 RegisterId, UINT64 Value);
UINT64 GetRegValue(PGUEST_REGS GuestRegs, REGS_ENUM RegId);
BOOLEAN SetRegValueHwdbg(UINT64 *Regs, UINT32 RegisterId, UINT64 Value);
UINT64 GetRegValueHwdbg(UINT64 *Regs, UINT32 RegId);
BOOL ScriptEngineExecute(
    PGUEST_REGS GuestRegs, ACTION_BUFFER *ActionDetail,
    PSCRIPT_ENGINE_GENERAL_REGISTERS ScriptGeneralRegisters,
    SYMBOL_BUFFER *CodeBuffer, UINT64 *Indx, SYMBOL *ErrorOperator);
UINT64 GetRegValue(PGUEST_REGS GuestRegs, REGS_ENUM RegId);
VOID ScriptEngineGetOperatorName(PSYMBOL OperatorSymbol, CHAR *BufferForName);
VOID ScriptEngineGetOperatorName(PSYMBOL OperatorSymbol, CHAR *BufferForName);
