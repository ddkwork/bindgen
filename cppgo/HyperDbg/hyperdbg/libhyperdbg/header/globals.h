
#pragma once

BOOLEAN g_RtmSupport = FALSE;

UINT32 g_VirtualAddressWidth = 0;

BOOLEAN g_IsInterpreterOnString = FALSE;

BOOLEAN g_IsInterpreterPreviousCharacterABackSlash = FALSE;

UINT32 g_InterpreterCountOfOpenCurlyBrackets = 0;

BYTE g_EndOfBufferCheckTcp[TCP_END_OF_BUFFER_CHARS_COUNT] = {
    TCP_END_OF_BUFFER_CHAR_1, TCP_END_OF_BUFFER_CHAR_2,
    TCP_END_OF_BUFFER_CHAR_3, TCP_END_OF_BUFFER_CHAR_4};

BOOLEAN g_IsConnectedToHyperDbgLocally = FALSE;

BOOLEAN g_IsConnectedToRemoteDebuggee = FALSE;

BOOLEAN g_IsConnectedToRemoteDebugger = FALSE;

SOCKET g_ClientConnectSocket = {0};

SOCKET g_SeverSocket = {0};

SOCKET g_ServerListenSocket = {0};

string g_ServerPort = "";

string g_ServerIp = "";

HANDLE g_RemoteDebuggeeListeningThread = NULL;

HANDLE g_IsDriverLoadedSuccessfully = NULL;

HANDLE g_EndOfMessageReceivedEvent = NULL;

BOOLEAN g_IsEndOfMessageReceived = FALSE;

BOOLEAN g_SerialConnectionAlreadyClosed = FALSE;

BOOLEAN g_IgnorePauseRequests = FALSE;

BOOLEAN g_IsUserDebuggerInitialized = FALSE;

DEBUGGER_SYNCRONIZATION_EVENTS_STATE g_UserSyncronizationObjectsHandleTable
    [DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS] = {0};

BYTE g_EndOfBufferCheckSerial[SERIAL_END_OF_BUFFER_CHARS_COUNT] = {
    SERIAL_END_OF_BUFFER_CHAR_1, SERIAL_END_OF_BUFFER_CHAR_2,
    SERIAL_END_OF_BUFFER_CHAR_3, SERIAL_END_OF_BUFFER_CHAR_4};

DEBUGGER_SYNCRONIZATION_EVENTS_STATE g_KernelSyncronizationObjectsHandleTable
    [DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS] = {0};

BYTE g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE] = {0};

BOOLEAN g_IsRunningInstruction32Bit = FALSE;

HANDLE g_SerialListeningThreadHandle = NULL;

HANDLE g_SerialRemoteComPortHandle = NULL;

BOOLEAN g_IsSerialConnectedToRemoteDebuggee = FALSE;

BOOLEAN g_IsSerialConnectedToRemoteDebugger = FALSE;

BOOLEAN g_IsDebuggeeInHandshakingPhase = FALSE;

BOOLEAN g_IsDebuggeeRunning = FALSE;

BOOLEAN g_IgnoreNewLoggingMessages = FALSE;

ULONG g_CurrentRemoteCore = DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE;

BOOLEAN g_IsDebuggerConntectedToNamedPipe = FALSE;

HANDLE g_DebuggeeStopCommandEventHandle = NULL;

DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfRegisteringEvent = {0};

DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfAddingActionsToEvent = {0};

OVERLAPPED g_OverlappedIoStructureForReadDebugger = {0};
OVERLAPPED g_OverlappedIoStructureForWriteDebugger = {0};
OVERLAPPED g_OverlappedIoStructureForReadDebuggee = {0};

BOOLEAN g_SharedEventStatus = FALSE;

BOOLEAN g_ShouldPreviousCommandBeContinued;

CommandType g_CommandsList;

UINT64 *g_ScriptGlobalVariables;

UINT64 *g_ScriptStackBuffer;

BOOLEAN g_IsCommandListInitialized = FALSE;

BOOLEAN g_IsDebuggerModulesLoaded = FALSE;

ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState = {0};

UINT32 g_ProcessIdOfLatestStartingProcess = NULL;

UINT64 g_EventTag = DebuggerEventTagStartSeed;

UINT64 g_OutputSourceTag = DebuggerOutputSourceTagStartSeed;

BOOLEAN g_EventTraceInitialized = FALSE;

LIST_ENTRY g_EventTrace = {0};

BOOLEAN g_OutputSourcesInitialized = FALSE;

LIST_ENTRY g_OutputSources = {0};

TCHAR g_DriverLocation[MAX_PATH] = {0};

TCHAR g_DriverName[MAX_PATH] = {0};

BOOLEAN g_UseCustomDriverLocation = FALSE;

TCHAR g_TestLocation[MAX_PATH] = {0};

PVOID g_MessageHandler = 0;

PVOID g_MessageHandlerSharedBuffer = 0;

BOOLEAN g_IsVmxOffProcessStart;

HANDLE g_DeviceHandle;

BOOLEAN g_LogOpened = FALSE;

ofstream g_LogOpenFile;

BOOLEAN g_ExecutingScript = FALSE;

BOOLEAN g_BreakPrintingOutput = FALSE;

BOOLEAN g_IsExecutingSymbolLoadingRoutines = FALSE;

std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION> g_DisassemblerSymbolMap;

BOOLEAN g_TransparentResultsMeasured = FALSE;

UINT64 g_CpuidAverage = 0;

UINT64 g_CpuidStandardDeviation = 0;

UINT64 g_CpuidMedian = 0;

UINT64 g_RdtscAverage = 0;

UINT64 g_RdtscStandardDeviation = 0;

UINT64 g_RdtscMedian = 0;

BOOLEAN g_IsInstrumentingInstructions = FALSE;

UINT64 g_KernelBaseAddress;

BOOLEAN g_PrivilegesAlreadyAdjusted = FALSE;

BOOLEAN g_AutoUnpause = TRUE;

BOOLEAN g_AddressConversion = TRUE;

BOOLEAN g_AutoFlush = FALSE;

UINT32 g_DisassemblerSyntax = 1;

PMODULE_SYMBOL_DETAIL g_SymbolTable = NULL;

UINT32 g_SymbolTableSize = NULL;

UINT32 g_SymbolTableCurrentIndex = NULL;

UINT64 g_ResultOfEvaluatedExpression = NULL;

UINT32 g_ErrorStateOfResultOfEvaluatedExpression = NULL;

std::wstring g_StartCommandPath = L"";

std::wstring g_StartCommandPathAndArguments = L"";

UINT64 g_CurrentExprEvalResult;

BOOLEAN g_CurrentExprEvalResultHasError;

HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;

BOOLEAN g_HwdbgInstanceInfoIsValid;

std::vector<UINT32> g_HwdbgPortConfiguration;

UINT64 *g_HwdbgPinsStatus;
