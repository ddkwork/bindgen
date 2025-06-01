
#pragma once
#pragma pack(1)
#define BUS_BIT_WIDTH 8
#define DEVICE_BIT_WIDTH 5
#define FUNCTION_BIT_WIDTH 3
#define DOMAIN_MAX_NUM 0
#define BUS_MAX_NUM 255
#define DEVICE_MAX_NUM 32
#define FUNCTION_MAX_NUM 8
#define DEV_MAX_NUM 255
#define CAM_CONFIG_SPACE_LENGTH 255

typedef struct _PORTABLE_PCI_COMMON_HEADER {
  UINT16 VendorId;
  UINT16 DeviceId;
  UINT16 Command;
  UINT16 Status;
  UINT8 RevisionId;
  UINT8 ClassCode[3];
  UINT8 CacheLineSize;
  UINT8 PrimaryLatencyTimer;
  UINT8 HeaderType;
  UINT8 Bist;
} PORTABLE_PCI_COMMON_HEADER, *PPORTABLE_PCI_COMMON_HEADER;

typedef union _PORTABLE_PCI_DEVICE_HEADER {
  struct _PORTABLE_PCI_EP_HEADER {
    UINT32 Bar[6];
    UINT32 CardBusCISPtr;
    UINT16 SubVendorId;
    UINT16 SubSystemId;
    UINT32 ROMBar;
    UINT8 CapabilitiesPtr;
    UINT8 Reserved[3];
    UINT32 Reserved1;
    UINT8 InterruptLine;
    UINT8 InterruptPin;
    UINT8 MinGnt;
    UINT8 MaxLat;
  } ConfigSpaceEp;

  struct _PORTABLE_PCI_BRIDGE_HEADER {
    UINT32 Bar[2];
    UINT8 PrimaryBusNumber;
    UINT8 SecondaryBusNumber;
    UINT8 SubordinateBusNumber;
    UINT8 SecondaryLatencyTimer;
    UINT8 IoBase;
    UINT8 IoLimit;
    UINT16 SecondaryStatus;
    UINT16 MemoryBase;
    UINT16 MemoryLimit;
    UINT16 PrefetchableMemoryBase;
    UINT16 PrefetchableMemoryLimit;
    UINT32 PrefetchableBaseUpper32b;
    UINT32 PrefetchableLimitUpper32b;
    UINT16 IoBaseUpper16b;
    UINT16 IoLimitUpper16b;
    UINT8 CapabilityPtr;
    UINT8 Reserved[3];
    UINT32 ROMBar;
    UINT8 InterruptLine;
    UINT8 InterruptPin;
    UINT16 BridgeControl;
  } ConfigSpacePtpBridge;
} PORTABLE_PCI_DEVICE_HEADER, *PPORTABLE_PCI_DEVICE_HEADER;

typedef struct _PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL {
  UINT16 VendorId;
  UINT16 DeviceId;
  UINT8 ClassCode[3];
} PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL,
    *PPORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL;

typedef struct _PCI_DEV_MINIMAL {
  UINT8 Bus : BUS_BIT_WIDTH;
  UINT8 Device : DEVICE_BIT_WIDTH;
  UINT8 Function : FUNCTION_BIT_WIDTH;
  PORTABLE_PCI_CONFIG_SPACE_HEADER_MINIMAL ConfigSpace;
} PCI_DEV_MINIMAL, *PPCI_DEV_MINIMAL;

typedef struct _PCI_DEV_MMIOBAR_INFO {
  BOOL Is64Bit;
  BOOL IsEnabled;
  UINT64 BarOffsetEnd;
  UINT64 BarSize;
} PCI_DEV_MMIOBAR_INFO, *PPCI_DEV_MMIOBAR_INFO;

typedef struct _PORTABLE_PCI_CONFIG_SPACE_HEADER {
  PORTABLE_PCI_COMMON_HEADER CommonHeader;
  PORTABLE_PCI_DEVICE_HEADER DeviceHeader;
} PORTABLE_PCI_CONFIG_SPACE_HEADER, *PPORTABLE_PCI_CONFIG_SPACE_HEADER;

typedef struct _PCI_DEV {
  UINT8 Bus : BUS_BIT_WIDTH;
  UINT8 Device : DEVICE_BIT_WIDTH;
  UINT8 Function : FUNCTION_BIT_WIDTH;
  PORTABLE_PCI_CONFIG_SPACE_HEADER ConfigSpace;
  BYTE ConfigSpaceAdditional[CAM_CONFIG_SPACE_LENGTH -
                             sizeof(PORTABLE_PCI_CONFIG_SPACE_HEADER)];
  PCI_DEV_MMIOBAR_INFO MmioBarInfo[6];
} PCI_DEV, *PPCI_DEV;

#pragma pack()
