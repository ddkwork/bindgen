
#include "pch.h"

UINT64 IoApicRead(volatile IO_APIC_ENT *IoApicBaseVa, UINT32 Reg) {
  UINT32 High = 0, Low;
  IoApicBaseVa->Reg = Reg;
  Low = IoApicBaseVa->Data;
  if (Reg >= IOAPIC_REDTBL(0) && Reg < IOAPIC_REDTBL(IOAPIC_REDTBL_MAX)) {
    Reg++;
    IoApicBaseVa->Reg = Reg;
    High = IoApicBaseVa->Data;
    return IOAPIC_APPEND_QWORD(High, Low);
  } else {
    return Low;
  }
}

VOID IoApicWrite(volatile IO_APIC_ENT *IoApicBaseVa, UINT32 Reg, UINT64 Data) {
  IoApicBaseVa->Reg = Reg;
  IoApicBaseVa->Data = IOAPIC_LOW_DWORD(Data);
  if (Reg >= IOAPIC_REDTBL(0) && Reg < IOAPIC_REDTBL(IOAPIC_REDTBL_MAX)) {
    Reg++;
    IoApicBaseVa->Reg = Reg;
    IoApicBaseVa->Data = IOAPIC_HIGH_DWORD(Data);
  }
}

VOID ApicDumpIoApic(IO_APIC_ENTRY_PACKETS *IoApicPackets) {
  UINT32 Index = 0;
  UINT32 Max;
  UINT64 ll, lh;
  UINT64 ApicBasePa = IO_APIC_DEFAULT_BASE_ADDR;
  ll = IoApicRead(g_IoApicBase, IO_VERS_REGISTER), Max = (ll >> 16) & 0xff;
  IoApicPackets->ApicBasePa = (UINT32)ApicBasePa;
  IoApicPackets->ApicBaseVa = (UINT64)g_IoApicBase;
  IoApicPackets->IoIdReg = (UINT32)IoApicRead(g_IoApicBase, IO_ID_REGISTER);
  IoApicPackets->IoLl = (UINT32)ll;
  IoApicPackets->IoArbIdReg =
      (UINT32)IoApicRead(g_IoApicBase, IO_ARB_ID_REGISTER);
  Max *= 2;
  for (Index = 0; Index <= Max; Index += 2) {
    if (Index >= MAX_NUMBER_OF_IO_APIC_ENTRIES) {
      return;
    }
    ll = IoApicRead(g_IoApicBase, IO_REDIR_BASE + Index + 0);
    lh = IoApicRead(g_IoApicBase, IO_REDIR_BASE + Index + 1);
    IoApicPackets->LlLhData[Index] = ll;
    IoApicPackets->LlLhData[Index + 1] = lh;
  }
}

VOID XApicIcrWrite(UINT32 Low, UINT32 High) {
  *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset + 0x10) = High;
  *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset) = Low;
}

VOID X2ApicIcrWrite(UINT32 Low, UINT32 High) {
  __writemsr(X2_MSR_BASE + TO_X2(ICROffset), ((UINT64)High << 32) | Low);
}

UINT64 X2ApicRead(UINT32 Offset) {
  return __readmsr(X2_MSR_BASE + TO_X2(Offset));
}

BOOLEAN ApicStoretLocalApicInXApicMode(PLAPIC_PAGE LApicBuffer) {
  volatile LAPIC_PAGE *LocalApicVa = g_ApicBase;
  if (!g_ApicBase) {
    return FALSE;
  }
  LApicBuffer->Id = LocalApicVa->Id;
  LApicBuffer->Version = LocalApicVa->Version;
  LApicBuffer->SpuriousInterruptVector = LocalApicVa->SpuriousInterruptVector;
  LApicBuffer->TPR = LocalApicVa->TPR;
  LApicBuffer->ProcessorPriority = LocalApicVa->ProcessorPriority;
  LApicBuffer->LogicalDestination = LocalApicVa->LogicalDestination;
  LApicBuffer->ErrorStatus = LocalApicVa->ErrorStatus;
  LApicBuffer->LvtLINT0 = LocalApicVa->LvtLINT0;
  LApicBuffer->LvtLINT1 = LocalApicVa->LvtLINT1;
  LApicBuffer->LvtCmci = LocalApicVa->LvtCmci;
  LApicBuffer->LvtPerfMonCounters = LocalApicVa->LvtPerfMonCounters;
  LApicBuffer->LvtTimer = LocalApicVa->LvtTimer;
  LApicBuffer->LvtThermalSensor = LocalApicVa->LvtThermalSensor;
  LApicBuffer->LvtError = LocalApicVa->LvtError;
  LApicBuffer->InitialCount = LocalApicVa->InitialCount;
  LApicBuffer->CurrentCount = LocalApicVa->CurrentCount;
  LApicBuffer->DivideConfiguration = LocalApicVa->DivideConfiguration;
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->ISR[i * 4] = LocalApicVa->ISR[i * 4];
  }
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->TMR[i * 4] = LocalApicVa->TMR[i * 4];
  }
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->IRR[i * 4] = LocalApicVa->IRR[i * 4];
  }
  return TRUE;
}

BOOLEAN ApicStoreLocalApicInX2ApicMode(PLAPIC_PAGE LApicBuffer) {
  LApicBuffer->Id = (UINT32)X2ApicRead(APIC_ID);
  LApicBuffer->Version = (UINT32)X2ApicRead(APIC_VERSION);
  LApicBuffer->SpuriousInterruptVector =
      (UINT32)X2ApicRead(APIC_SPURIOUS_INTERRUPT_VECTOR);
  LApicBuffer->TPR = (UINT32)X2ApicRead(APIC_TASK_PRIORITY);
  LApicBuffer->ProcessorPriority = (UINT32)X2ApicRead(APIC_PROCESSOR_PRIORITY);
  LApicBuffer->LogicalDestination =
      (UINT32)X2ApicRead(APIC_LOGICAL_DESTINATION);
  LApicBuffer->ErrorStatus = (UINT32)X2ApicRead(APIC_ERROR_STATUS);
  LApicBuffer->LvtLINT0 = (UINT32)X2ApicRead(APIC_LVT_LINT0);
  LApicBuffer->LvtLINT1 = (UINT32)X2ApicRead(APIC_LVT_LINT1);
  LApicBuffer->LvtCmci =
      (UINT32)X2ApicRead(APIC_LVT_CORRECTED_MACHINE_CHECK_INTERRUPT);
  LApicBuffer->LvtPerfMonCounters =
      (UINT32)X2ApicRead(APIC_LVT_PERFORMANCE_MONITORING_COUNTERS);
  LApicBuffer->LvtTimer = (UINT32)X2ApicRead(APIC_LVT_TIMER);
  LApicBuffer->LvtThermalSensor = (UINT32)X2ApicRead(APIC_LVT_THERMAL_SENSOR);
  LApicBuffer->LvtError = (UINT32)X2ApicRead(APIC_LVT_ERROR);
  LApicBuffer->InitialCount = (UINT32)X2ApicRead(APIC_INITIAL_COUNT);
  LApicBuffer->CurrentCount = (UINT32)X2ApicRead(APIC_CURRENT_COUNT);
  LApicBuffer->DivideConfiguration =
      (UINT32)X2ApicRead(APIC_DIVIDE_CONFIGURATION);
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->ISR[i * 4] =
        (UINT32)X2ApicRead(APIC_IN_SERVICE_BITS_31_0 + (0x10 * i));
  }
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->TMR[i * 4] =
        (UINT32)X2ApicRead(APIC_TRIGGER_MODE_BITS_31_0 + (0x10 * i));
  }
  for (UINT32 i = 0; i < 8; i++) {
    LApicBuffer->IRR[i * 4] =
        (UINT32)X2ApicRead(APIC_INTERRUPT_REQUEST_BITS_31_0 + (0x10 * i));
  }
  return TRUE;
}

VOID ApicTriggerGenericNmi() {
  if (g_CompatibilityCheck.IsX2Apic) {
    X2ApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
  } else {
    XApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
  }
}

VOID ApicTriggerGenericSmi() {
  LogInfo("Generating SMIs");
  if (g_CompatibilityCheck.IsX2Apic) {
    X2ApicIcrWrite((2 << 8) | (1 << 14) | (3 << 18), 0);
  } else {
    XApicIcrWrite((2 << 8) | (1 << 14) | (3 << 18), 0);
  }
}

BOOLEAN ApicStoreLocalApicFields(PLAPIC_PAGE LApicBuffer,
                                 PBOOLEAN IsUsingX2APIC) {
  if (g_CompatibilityCheck.IsX2Apic) {
    *IsUsingX2APIC = TRUE;
    return ApicStoreLocalApicInX2ApicMode(LApicBuffer);
  } else {
    *IsUsingX2APIC = FALSE;
    return ApicStoretLocalApicInXApicMode(LApicBuffer);
  }
}

BOOLEAN ApicStoreIoApicFields(IO_APIC_ENTRY_PACKETS *IoApicPackets) {
  ApicDumpIoApic(IoApicPackets);
  return TRUE;
}

BOOLEAN ApicInitialize() {
  UINT64 ApicBaseMSR;
  PHYSICAL_ADDRESS PaApicBase;
  PHYSICAL_ADDRESS PaIoApicBase;
  ApicBaseMSR = __readmsr(IA32_APIC_BASE);
  if (!(ApicBaseMSR & (1 << 11))) {
    return FALSE;
  }
  PaIoApicBase.QuadPart = IO_APIC_DEFAULT_BASE_ADDR & 0xFFFFFF000;
  g_IoApicBase = MmMapIoSpace(PaIoApicBase, 0x1000, MmNonCached);
  if (!g_IoApicBase) {
  }
  if (ApicBaseMSR & (1 << 10)) {
    g_CompatibilityCheck.IsX2Apic = TRUE;
    return FALSE;
  } else {
    PaApicBase.QuadPart = ApicBaseMSR & 0xFFFFFF000;
    g_ApicBase = MmMapIoSpace(PaApicBase, 0x1000, MmNonCached);
    if (!g_ApicBase) {
      return FALSE;
    }
    g_CompatibilityCheck.IsX2Apic = FALSE;
  }
  return TRUE;
}

VOID ApicUninitialize() {
  if (g_ApicBase) {
    MmUnmapIoSpace(g_ApicBase, 0x1000);
  }
  if (g_IoApicBase) {
    MmUnmapIoSpace(g_IoApicBase, 0x1000);
  }
}

VOID ApicSelfIpi(UINT32 Vector) {
  if (g_CompatibilityCheck.IsX2Apic) {
    X2ApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector,
                   0);
  } else {
    XApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector,
                  0);
  }
}
