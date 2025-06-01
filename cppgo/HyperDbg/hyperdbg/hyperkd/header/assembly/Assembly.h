
#pragma once

extern unsigned long long AsmTestWrapperWithTestTags(unsigned long long Param1,
                                                     unsigned long long Param2,
                                                     unsigned long long Param3,
                                                     unsigned long long Param4);

extern void AsmDebuggerCustomCodeHandler(unsigned long long Param1,
                                         unsigned long long Param2,
                                         unsigned long long Param3,
                                         unsigned long long Param4);

extern unsigned long long AsmDebuggerConditionCodeHandler(
    unsigned long long Param1, unsigned long long Param2,
    unsigned long long Param3);

extern void AsmDebuggerSpinOnThread();
