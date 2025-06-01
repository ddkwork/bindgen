
#pragma once

typedef enum _LOG_TYPE { LOG_INFO, LOG_WARNING, LOG_ERROR } LOG_TYPE;

#if UseDbgPrintInsteadOfUsermodeMessageTracking

#define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 ##__VA_ARGS__)
#define LogWarning(format, ...)                                      \
  DbgPrint("[-] Warning (%s:%d) | " format "\n", __func__, __LINE__, \
           ##__VA_ARGS__)
#define LogError(format, ...)                                      \
  DbgPrint("[!] Error (%s:%d) | " format "\n", __func__, __LINE__, \
           ##__VA_ARGS__);                                         \
  DbgBreakPoint()

#define Log(format, ...) DbgPrint(format, ##__VA_ARGS__)
#else

#define LogInfo(format, ...)                                        \
  LogCallbackPrepareAndSendMessageToQueue(                          \
      OPERATION_LOG_INFO_MESSAGE, UseImmediateMessaging,            \
      ShowSystemTimeOnDebugMessages, FALSE,                         \
      "[+] Information (%s:%d) | " format "\n", __func__, __LINE__, \
      ##__VA_ARGS__)

#define LogInfoPriority(format, ...)                                         \
  LogCallbackPrepareAndSendMessageToQueue(                                   \
      OPERATION_LOG_INFO_MESSAGE, TRUE, ShowSystemTimeOnDebugMessages, TRUE, \
      "[+] Information (%s:%d) | " format "\n", __func__, __LINE__,          \
      ##__VA_ARGS__)

#define LogWarning(format, ...)                             \
  LogCallbackPrepareAndSendMessageToQueue(                  \
      OPERATION_LOG_WARNING_MESSAGE, UseImmediateMessaging, \
      ShowSystemTimeOnDebugMessages, TRUE,                  \
      "[-] Warning (%s:%d) | " format "\n", __func__, __LINE__, ##__VA_ARGS__)

#define LogError(format, ...)                                                  \
  LogCallbackPrepareAndSendMessageToQueue(                                     \
      OPERATION_LOG_ERROR_MESSAGE, UseImmediateMessaging,                      \
      ShowSystemTimeOnDebugMessages, TRUE, "[!] Error (%s:%d) | " format "\n", \
      __func__, __LINE__, ##__VA_ARGS__);                                      \
  if (DebugMode) DbgBreakPoint()

#define Log(format, ...)                                                    \
  LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, TRUE, \
                                          FALSE, FALSE, format, ##__VA_ARGS__)

#define LogSimpleWithTag(tag, isimmdte, buffer, len) \
  LogCallbackSendMessageToQueue(tag, isimmdte, buffer, len, FALSE)
#endif

#define LogDebugInfo(format, ...)                                   \
  if (DebugMode)                                                    \
  LogCallbackPrepareAndSendMessageToQueue(                          \
      OPERATION_LOG_INFO_MESSAGE, UseImmediateMessaging,            \
      ShowSystemTimeOnDebugMessages, FALSE,                         \
      "[+] Information (%s:%d) | " format "\n", __func__, __LINE__, \
      ##__VA_ARGS__)
