
#pragma once
UINT64 ScriptEngineWrapperGetInstructionPointer();
UINT64 ScriptEngineWrapperGetAddressOfReservedBuffer(
    PDEBUGGER_EVENT_ACTION Action);
VOID ScriptEngineUpdateTargetCoreDateTime(PROCESSOR_DEBUGGING_STATE *DbgState);
UINT64 ScriptEngineGetTargetCoreTime();
UINT64 ScriptEngineGetTargetCoreDate();
