
#pragma once
VOID MsrHandleRdmsrVmexit(VIRTUAL_MACHINE_STATE *VCpu);
VOID MsrHandleWrmsrVmexit(VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN MsrHandleSetMsrBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Msr,
                              BOOLEAN ReadDetection, BOOLEAN WriteDetection);
BOOLEAN MsrHandleUnSetMsrBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Msr,
                                BOOLEAN ReadDetection, BOOLEAN WriteDetection);
VOID MsrHandlePerformMsrBitmapReadChange(VIRTUAL_MACHINE_STATE *VCpu,
                                         UINT32 MsrMask);
VOID MsrHandlePerformMsrBitmapReadReset(VIRTUAL_MACHINE_STATE *VCpu);
VOID MsrHandlePerformMsrBitmapWriteChange(VIRTUAL_MACHINE_STATE *VCpu,
                                          UINT32 MsrMask);
VOID MsrHandlePerformMsrBitmapWriteReset(VIRTUAL_MACHINE_STATE *VCpu);
