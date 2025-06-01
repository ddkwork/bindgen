
#include "pch.h"

UINT64 ScriptEngineKeywordPoi(PUINT64 Address, BOOL *HasError) {
  UINT64 Result = (UINT64)NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(UINT64))) {
    *HasError = TRUE;
    return (UINT64)NULL;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = *Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(UINT64));
#endif
  return Result;
}

WORD ScriptEngineKeywordHi(PUINT64 Address, BOOL *HasError) {
  QWORD Result = NULL64_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(UINT64))) {
    *HasError = TRUE;
    return NULL64_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = *Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(UINT64));
#endif
  return HIWORD(Result);
}

WORD ScriptEngineKeywordLow(PUINT64 Address, BOOL *HasError) {
  QWORD Result = NULL64_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(UINT64))) {
    *HasError = TRUE;
    return NULL64_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = *Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(UINT64));
#endif
  return LOWORD(Result);
}

BYTE ScriptEngineKeywordDb(PUINT64 Address, BOOL *HasError) {
  BYTE Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(BYTE))) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = (BYTE)*Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(BYTE));
#endif
  return Result;
}

DWORD ScriptEngineKeywordDd(PUINT64 Address, BOOL *HasError) {
  DWORD Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(DWORD))) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = (DWORD)*Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(DWORD));
#endif
  return Result;
}

WORD ScriptEngineKeywordDw(PUINT64 Address, BOOL *HasError) {
  WORD Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(WORD))) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = (WORD)*Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(WORD));
#endif
  return Result;
}

QWORD ScriptEngineKeywordDq(PUINT64 Address, BOOL *HasError) {
  QWORD Result = (QWORD)NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Address, sizeof(DWORD))) {
    *HasError = TRUE;
    return (QWORD)NULL;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = *Address;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, &Result,
                                            sizeof(QWORD));
#endif
  return Result;
}

UINT64 ScriptEngineKeywordPoiPa(PUINT64 Address, BOOL *HasError) {
  UINT64 Result = (UINT64)NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return (UINT64)NULL;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (poi_pa) is not possible in "
      "user-mode\n");
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(UINT64));
#endif
  return Result;
}

WORD ScriptEngineKeywordHiPa(PUINT64 Address, BOOL *HasError) {
  QWORD Result = NULL64_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return NULL64_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (hi_pa) is not possible in "
      "user-mode\n");
  return (WORD)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(UINT64));
#endif
  return HIWORD(Result);
}

WORD ScriptEngineKeywordLowPa(PUINT64 Address, BOOL *HasError) {
  QWORD Result = NULL64_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return NULL64_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (low_pa) is not possible in "
      "user-mode\n");
  return (WORD)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(UINT64));
#endif
  return LOWORD(Result);
}

BYTE ScriptEngineKeywordDbPa(PUINT64 Address, BOOL *HasError) {
  BYTE Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (db_pa) is not possible in "
      "user-mode\n");
  return (BYTE)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(BYTE));
#endif
  return Result;
}

DWORD ScriptEngineKeywordDdPa(PUINT64 Address, BOOL *HasError) {
  DWORD Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (dd_pa) is not possible in "
      "user-mode\n");
  return (DWORD)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(DWORD));
#endif
  return Result;
}

WORD ScriptEngineKeywordDwPa(PUINT64 Address, BOOL *HasError) {
  WORD Result = NULL_ZERO;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return NULL_ZERO;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (dw_pa) is not possible in "
      "user-mode\n");
  return (WORD)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(WORD));
#endif
  return Result;
}

QWORD ScriptEngineKeywordDqPa(PUINT64 Address, BOOL *HasError) {
  QWORD Result = (QWORD)NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical((UINT64)Address)) {
    *HasError = TRUE;
    return (QWORD)NULL;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address keywords (dq_pa) is not possible in "
      "user-mode\n");
  return (QWORD)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperReadMemorySafeByPhysicalAddress((UINT64)Address, (UINT64)&Result,
                                              sizeof(QWORD));
#endif
  return Result;
}
