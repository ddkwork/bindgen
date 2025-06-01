
#pragma once

extern void inline AsmEnableVmxOperation();

extern void inline AsmRestoreToVmxOffState();

extern NTSTATUS inline AsmVmxVmcall(unsigned long long VmcallNumber,
                                    unsigned long long OptionalParam1,
                                    unsigned long long OptionalParam2,
                                    long long OptionalParam3);

extern void inline AsmHypervVmcall(unsigned long long GuestRegisters);

extern unsigned long long inline AsmVmfunc(unsigned long EptpIndex,
                                           unsigned long Function);

extern void AsmVmxSaveState();

extern void AsmVmxRestoreState();

extern void AsmVmexitHandler();

extern void inline AsmSaveVmxOffState();

extern unsigned char inline AsmInvept(unsigned long Type, void *Descriptors);

extern unsigned char inline AsmInvvpid(unsigned long Type, void *Descriptors);

extern unsigned short AsmGetCs();

extern unsigned short AsmGetDs();
extern void AsmSetDs(unsigned short DsSelector);

extern unsigned short AsmGetEs();
extern void AsmSetEs(unsigned short EsSelector);

extern unsigned short AsmGetSs();
extern void AsmSetSs(unsigned short SsSelector);

extern unsigned short AsmGetFs();
extern void AsmSetFs(unsigned short FsSelector);

extern unsigned short AsmGetGs();

extern unsigned short AsmGetLdtr();

extern unsigned short AsmGetTr();

extern unsigned long long inline AsmGetGdtBase();

extern unsigned short AsmGetGdtLimit();

extern unsigned long long inline AsmGetIdtBase();

extern unsigned short AsmGetIdtLimit();
extern UINT32 AsmGetAccessRights(unsigned short Selector);

extern unsigned short AsmGetRflags();

extern void inline AsmCliInstruction();

extern void inline AsmStiInstruction();

extern void AsmReloadGdtr(void *GdtBase, unsigned long GdtLimit);

extern void AsmReloadIdtr(void *GdtBase, unsigned long GdtLimit);

extern void AsmGeneralDetourHook(void);

extern unsigned long long AsmTestWrapperWithTestTags(unsigned long long Param1,
                                                     unsigned long long Param2,
                                                     unsigned long long Param3,
                                                     unsigned long long Param4);

extern void InterruptHandler0();

extern void InterruptHandler1();

extern void InterruptHandler2();

extern void InterruptHandler3();

extern void InterruptHandler4();

extern void InterruptHandler5();

extern void InterruptHandler6();

extern void InterruptHandler7();

extern void InterruptHandler8();

extern void InterruptHandler9();

extern void InterruptHandler10();

extern void InterruptHandler11();

extern void InterruptHandler12();

extern void InterruptHandler13();

extern void InterruptHandler14();

extern void InterruptHandler15();

extern void InterruptHandler16();

extern void InterruptHandler17();

extern void InterruptHandler18();

extern void InterruptHandler19();

extern void InterruptHandler20();

extern void InterruptHandler21();

extern void InterruptHandler22();

extern void InterruptHandler23();

extern void InterruptHandler24();

extern void InterruptHandler25();

extern void InterruptHandler26();

extern void InterruptHandler27();

extern void InterruptHandler28();

extern void InterruptHandler29();

extern void InterruptHandler30();
