
#pragma once
#include "pch.h"
#define CFGADR 0xCF8
#define CFGDAT 0xCFC
QWORD PciReadCam(WORD Bus, WORD Device, WORD Function, BYTE Offset,
                 UINT8 Width);
BOOLEAN PciWriteCam(WORD Bus, WORD Device, WORD Function, BYTE Offset,
                    UINT8 Width, QWORD Value);
