
#include "pch.h"

VOID PeHexDump(CHAR *Ptr, int Size, int SecAddress) {
  int i = 1, Temp = 0;
  wchar_t Buf[18];
  ShowMessages("\n\n%x: |", SecAddress);
  Buf[Temp] = ' ';
  Buf[Temp + 16] = ' ';
  Buf[Temp + 17] = 0;
  Temp++;
  for (; i <= Size; i++, Ptr++, Temp++) {
    Buf[Temp] = !iswcntrl((*Ptr) & 0xff) ? (*Ptr) & 0xff : '.';
    ShowMessages("%-3.2x", (*Ptr) & 0xff);
    if (i % 16 == 0) {
      _putws(Buf);
      if (i + 1 <= Size) ShowMessages("%x: ", SecAddress += 16);
      Temp = 0;
    }
    if (i % 4 == 0) ShowMessages("| ");
  }
  if (i % 16 != 0) {
    Buf[Temp] = 0;
    for (; i % 16 != 0; i++) ShowMessages("%-3.2c", ' ');
    _putws(Buf);
  }
}

BOOLEAN PeShowSectionInformationAndDump(const WCHAR *AddressOfFile,
                                        const CHAR *SectionToShow,
                                        BOOLEAN Is32Bit) {
  BOOLEAN Result = FALSE;
  HANDLE MapObjectHandle, FileHandle;
  UINT32 NumberOfSections;
  LPVOID BaseAddr;
  PIMAGE_DOS_HEADER DosHeader;
  PIMAGE_NT_HEADERS32 NtHeader32 = NULL;
  PIMAGE_NT_HEADERS64 NtHeader64 = NULL;
  IMAGE_FILE_HEADER Header;
  IMAGE_OPTIONAL_HEADER32 OpHeader32;
  IMAGE_OPTIONAL_HEADER64 OpHeader64;
  PIMAGE_SECTION_HEADER SecHeader;
  FileHandle = CreateFileW(AddressOfFile, GENERIC_READ, FILE_SHARE_READ, NULL,
                           OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL);
  if (FileHandle == INVALID_HANDLE_VALUE) {
    ShowMessages("err, could not open the file specified\n");
    return FALSE;
  };
  MapObjectHandle =
      CreateFileMapping(FileHandle, NULL, PAGE_READONLY, 0, 0, NULL);
  if (MapObjectHandle == NULL) {
    CloseHandle(FileHandle);
    return FALSE;
  }
  BaseAddr = MapViewOfFile(MapObjectHandle, FILE_MAP_READ, 0, 0, 0);
  if (BaseAddr == NULL) {
    CloseHandle(FileHandle);
    return FALSE;
  }
  DosHeader = (PIMAGE_DOS_HEADER)BaseAddr;
  if (DosHeader->e_magic == IMAGE_DOS_SIGNATURE) {
    ShowMessages("\nValid Dos Exe File\n------------------\n");
    ShowMessages("\nDumping DOS Header Info....\n---------------------------");
    ShowMessages("\n%-36s%s ",
                 "Magic number : ", DosHeader->e_magic == 0x5a4d ? "MZ" : "-");
    ShowMessages("\n%-36s%#x",
                 "Bytes on last page of file :", DosHeader->e_cblp);
    ShowMessages("\n%-36s%#x", "Pages in file : ", DosHeader->e_cp);
    ShowMessages("\n%-36s%#x", "Relocation : ", DosHeader->e_crlc);
    ShowMessages("\n%-36s%#x",
                 "Size of header in paragraphs : ", DosHeader->e_cparhdr);
    ShowMessages("\n%-36s%#x",
                 "Minimum extra paragraphs needed : ", DosHeader->e_minalloc);
    ShowMessages("\n%-36s%#x",
                 "Maximum extra paragraphs needed : ", DosHeader->e_maxalloc);
    ShowMessages("\n%-36s%#x",
                 "Initial (relative) SS value : ", DosHeader->e_ss);
    ShowMessages("\n%-36s%#x", "Initial SP value : ", DosHeader->e_sp);
    ShowMessages("\n%-36s%#x", "Checksum : ", DosHeader->e_csum);
    ShowMessages("\n%-36s%#x", "Initial IP value : ", DosHeader->e_ip);
    ShowMessages("\n%-36s%#x",
                 "Initial (relative) CS value : ", DosHeader->e_cs);
    ShowMessages("\n%-36s%#x",
                 "File address of relocation table : ", DosHeader->e_lfarlc);
    ShowMessages("\n%-36s%#x", "Overlay number : ", DosHeader->e_ovno);
    ShowMessages("\n%-36s%#x", "OEM identifier : ", DosHeader->e_oemid);
    ShowMessages("\n%-36s%#x",
                 "OEM information(e_oemid specific) :", DosHeader->e_oeminfo);
    ShowMessages("\n%-36s%#x",
                 "RVA address of PE header : ", DosHeader->e_lfanew);
    ShowMessages(
        "\n==============================================================="
        "================\n");
  } else {
    ShowMessages("\nGiven File is not a valid DOS file\n");
    Result = FALSE;
    goto Finished;
  }
  if (Is32Bit) {
    NtHeader32 =
        (PIMAGE_NT_HEADERS32)((UINT64)(DosHeader) + (DosHeader->e_lfanew));
  } else {
    NtHeader64 =
        (PIMAGE_NT_HEADERS64)((UINT64)(DosHeader) + (DosHeader->e_lfanew));
  }
  if (Is32Bit && NtHeader32->Signature == IMAGE_NT_SIGNATURE) {
    ShowMessages("\nValid PE32 file \n-------------\n");
  } else if (!Is32Bit && NtHeader64->Signature == IMAGE_NT_SIGNATURE) {
    ShowMessages("\nValid PE64 file \n-------------\n");
  } else {
    ShowMessages("err, the specified file is not a valid PE file");
    Result = FALSE;
    goto Finished;
  }
  ShowMessages(
      "\nDumping COFF/PE Header "
      "Info....\n--------------------------------");
  ShowMessages("\n%-36s%s", "Signature :", "PE");
  if (Is32Bit) {
    Header = NtHeader32->FileHeader;
  } else {
    Header = NtHeader64->FileHeader;
  }
  ShowMessages("\n%-36s", "Machine Architecture :");
  switch (Header.Machine) {
    case 0x0:
      ShowMessages("All ");
      break;
    case 0x14d:
      ShowMessages("Intel i860");
      break;
    case 0x14c:
      ShowMessages("Intel i386, i486, i586");
      break;
    case 0x200:
      ShowMessages("Intel Itanium processor");
      break;
    case 0x8664:
      ShowMessages("AMD x64");
      break;
    case 0x162:
      ShowMessages("MIPS R3000");
      break;
    case 0x166:
      ShowMessages("MIPS R4000");
      break;
    case 0x183:
      ShowMessages("DEC Alpha AXP");
      break;
    default:
      ShowMessages("Not Found");
      break;
  }
  ShowMessages("\n%-36s", "Characteristics : ");
  if ((Header.Characteristics & 0x0002) == 0x0002)
    ShowMessages("Executable Image, ");
  if ((Header.Characteristics & 0x0020) == 0x0020)
    ShowMessages("Application can address > 2GB, ");
  if ((Header.Characteristics & 0x1000) == 0x1000)
    ShowMessages("System file (Kernel Mode Driver(I think)), ");
  if ((Header.Characteristics & 0x2000) == 0x2000) ShowMessages("Dll file, ");
  if ((Header.Characteristics & 0x4000) == 0x4000)
    ShowMessages("Application runs only in Uniprocessor, ");
  ShowMessages("\n%-36s%s",
               "Time Stamp :", ctime((const time_t *)&(Header.TimeDateStamp)));
  ShowMessages("%-36s%d", "No.sections(size) :", Header.NumberOfSections);
  ShowMessages("\n%-36s%d",
               "No.entries in symbol table :", Header.NumberOfSymbols);
  ShowMessages("\n%-36s%d",
               "Size of optional header :", Header.SizeOfOptionalHeader);
  ShowMessages(
      "\n\nDumping PE Optional Header "
      "Info....\n-----------------------------------");
  if (Is32Bit) {
    OpHeader32 = NtHeader32->OptionalHeader;
    ShowMessages("\n\nInfo of optional Header\n-----------------------");
    ShowMessages("\n%-36s%#x",
                 "Address of Entry Point : ", OpHeader32.AddressOfEntryPoint);
    ShowMessages("\n%-36s%#llx",
                 "Base Address of the Image : ", OpHeader32.ImageBase);
    ShowMessages("\n%-36s%s", "SubSystem type : ",
                 OpHeader32.Subsystem == 1
                     ? "Device Driver(Native windows Process)"
                 : OpHeader32.Subsystem == 2 ? "Windows GUI"
                 : OpHeader32.Subsystem == 3 ? "Windows CLI"
                 : OpHeader32.Subsystem == 3 ? "Windows CLI"
                 : OpHeader32.Subsystem == 9 ? "Windows CE GUI"
                                             : "Unknown");
    ShowMessages("\n%-36s%s", "Given file is a : ",
                 OpHeader32.Magic == 0x20b ? "PE32+(64)" : "PE32");
    ShowMessages("\n%-36s%d",
                 "Size of code segment(.text) : ", OpHeader32.SizeOfCode);
    ShowMessages("\n%-36s%#x",
                 "Base address of code segment(RVA) :", OpHeader32.BaseOfCode);
    ShowMessages("\n%-36s%d", "Size of Initialized data : ",
                 OpHeader32.SizeOfInitializedData);
    ShowMessages("\n%-36s%#x",
                 "Base address of data segment(RVA) :", OpHeader32.BaseOfData);
    ShowMessages("\n%-36s%#x",
                 "Section Alignment :", OpHeader32.SectionAlignment);
    ShowMessages("\n%-36s%d",
                 "Major Linker Version : ", OpHeader32.MajorLinkerVersion);
    ShowMessages("\n%-36s%d",
                 "Minor Linker Version : ", OpHeader32.MinorLinkerVersion);
  } else {
    OpHeader64 = NtHeader64->OptionalHeader;
    ShowMessages("\n\nInfo of optional Header\n-----------------------");
    ShowMessages("\n%-36s%#x",
                 "Address of Entry Point : ", OpHeader64.AddressOfEntryPoint);
    ShowMessages("\n%-36s%#llx",
                 "Base Address of the Image : ", OpHeader64.ImageBase);
    ShowMessages("\n%-36s%s", "SubSystem type : ",
                 OpHeader64.Subsystem == 1
                     ? "Device Driver(Native windows Process)"
                 : OpHeader64.Subsystem == 2 ? "Windows GUI"
                 : OpHeader64.Subsystem == 3 ? "Windows CLI"
                 : OpHeader64.Subsystem == 3 ? "Windows CLI"
                 : OpHeader64.Subsystem == 9 ? "Windows CE GUI"
                                             : "Unknown");
    ShowMessages("\n%-36s%s", "Given file is a : ",
                 OpHeader64.Magic == 0x20b ? "PE32+(64)" : "PE32");
    ShowMessages("\n%-36s%d",
                 "Size of code segment(.text) : ", OpHeader64.SizeOfCode);
    ShowMessages("\n%-36s%#x",
                 "Base address of code segment(RVA) :", OpHeader64.BaseOfCode);
    ShowMessages("\n%-36s%d", "Size of Initialized data : ",
                 OpHeader64.SizeOfInitializedData);
    ShowMessages("\n%-36s%#x",
                 "Section Alignment :", OpHeader64.SectionAlignment);
    ShowMessages("\n%-36s%d",
                 "Major Linker Version : ", OpHeader64.MajorLinkerVersion);
    ShowMessages("\n%-36s%d",
                 "Minor Linker Version : ", OpHeader64.MinorLinkerVersion);
  }
  ShowMessages(
      "\n\nDumping Sections Header "
      "Info....\n--------------------------------");
  if (Is32Bit) {
    SecHeader = IMAGE_FIRST_SECTION(NtHeader32);
    NumberOfSections = NtHeader32->FileHeader.NumberOfSections;
  } else {
    SecHeader = IMAGE_FIRST_SECTION(NtHeader64);
    NumberOfSections = NtHeader64->FileHeader.NumberOfSections;
  }
  for (UINT32 i = 0; i < NumberOfSections; i++, SecHeader++) {
    if (Is32Bit) {
      ShowMessages("\n\nSection Info (%d of %d)", i + 1,
                   NtHeader32->FileHeader.NumberOfSections);
    } else {
      ShowMessages("\n\nSection Info (%d of %d)", i + 1,
                   NtHeader64->FileHeader.NumberOfSections);
    }
    ShowMessages("\n---------------------");
    ShowMessages("\n%-36s%s", "Section Header name : ", SecHeader->Name);
    ShowMessages("\n%-36s%#x",
                 "ActualSize of code or data : ", SecHeader->Misc.VirtualSize);
    ShowMessages("\n%-36s%#x",
                 "Virtual Address(RVA) :", SecHeader->VirtualAddress);
    ShowMessages("\n%-36s%#x", "Size of raw data (rounded to FA) : ",
                 SecHeader->SizeOfRawData);
    ShowMessages("\n%-36s%#x",
                 "Pointer to Raw Data : ", SecHeader->PointerToRawData);
    ShowMessages("\n%-36s%#x",
                 "Pointer to Relocations : ", SecHeader->PointerToRelocations);
    ShowMessages("\n%-36s%#x",
                 "Pointer to Line numbers : ", SecHeader->PointerToLinenumbers);
    ShowMessages("\n%-36s%#x",
                 "Number of relocations : ", SecHeader->NumberOfRelocations);
    ShowMessages("\n%-36s%#x",
                 "Number of line numbers : ", SecHeader->NumberOfLinenumbers);
    ShowMessages("\n%-36s%s", "Characteristics : ", "Contains ");
    if ((SecHeader->Characteristics & 0x20) == 0x20)
      ShowMessages("executable code, ");
    if ((SecHeader->Characteristics & 0x40) == 0x40)
      ShowMessages("initialized data, ");
    if ((SecHeader->Characteristics & 0x80) == 0x80)
      ShowMessages("uninitialized data, ");
    if ((SecHeader->Characteristics & 0x200) == 0x200)
      ShowMessages("comments and linker commands, ");
    if ((SecHeader->Characteristics & 0x10000000) == 0x10000000)
      ShowMessages("shareable data(via DLLs), ");
    if ((SecHeader->Characteristics & 0x40000000) == 0x40000000)
      ShowMessages("Readable, ");
    if ((SecHeader->Characteristics & 0x80000000) == 0x80000000)
      ShowMessages("Writable, ");
    if (SectionToShow != NULL) {
      if (!_strcmpi(SectionToShow, (const char *)SecHeader->Name)) {
        if (SecHeader->SizeOfRawData != 0) {
          if (Is32Bit) {
            PeHexDump((char *)((UINT64)DosHeader + SecHeader->PointerToRawData),
                      SecHeader->SizeOfRawData,
                      OpHeader32.ImageBase + SecHeader->VirtualAddress);
          } else {
            PeHexDump((char *)((UINT64)DosHeader + SecHeader->PointerToRawData),
                      SecHeader->SizeOfRawData,
                      (int)(OpHeader64.ImageBase + SecHeader->VirtualAddress));
          }
        }
      }
    }
  }
  ShowMessages(
      "\n==============================================================="
      "================\n");
  Result = TRUE;
Finished:
  UnmapViewOfFile(BaseAddr);
  CloseHandle(MapObjectHandle);
  return Result;
}

BOOLEAN PeIsPE32BitOr64Bit(const WCHAR *AddressOfFile, PBOOLEAN Is32Bit) {
  BOOLEAN Result = FALSE;
  HANDLE MapObjectHandle, FileHandle;
  LPVOID BaseAddr;
  PIMAGE_DOS_HEADER DosHeader;
  PIMAGE_NT_HEADERS32 NtHeader32 = NULL;
  IMAGE_OPTIONAL_HEADER32 OpHeader32;
  IMAGE_FILE_HEADER Header;
  FileHandle = CreateFileW(AddressOfFile, GENERIC_READ, FILE_SHARE_READ, NULL,
                           OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL);
  if (FileHandle == INVALID_HANDLE_VALUE) {
    ShowMessages("err, unable to read the file (%x)\n", GetLastError());
    return FALSE;
  };
  MapObjectHandle =
      CreateFileMapping(FileHandle, NULL, PAGE_READONLY, 0, 0, NULL);
  if (MapObjectHandle == NULL) {
    CloseHandle(FileHandle);
    ShowMessages("err, unable to create file mappings (%x)\n", GetLastError());
    return FALSE;
  }
  BaseAddr = MapViewOfFile(MapObjectHandle, FILE_MAP_READ, 0, 0, 0);
  if (BaseAddr == NULL) {
    CloseHandle(FileHandle);
    ShowMessages("err, unable to create map view of file (%x)\n",
                 GetLastError());
    return FALSE;
  }
  DosHeader = (PIMAGE_DOS_HEADER)BaseAddr;
  if (DosHeader->e_magic != IMAGE_DOS_SIGNATURE) {
    Result = FALSE;
    ShowMessages("err, the selected file is not in a valid PE format\n");
    goto Finished;
  }
  NtHeader32 =
      (PIMAGE_NT_HEADERS32)((UINT64)(DosHeader) + (DosHeader->e_lfanew));
  if (NtHeader32->Signature != IMAGE_NT_SIGNATURE) {
    Result = FALSE;
    ShowMessages("err, invalid image NT signature\n");
    goto Finished;
  }
  OpHeader32 = NtHeader32->OptionalHeader;
  Header = NtHeader32->FileHeader;
  switch (Header.Machine) {
    case IMAGE_FILE_MACHINE_I386:
      *Is32Bit = TRUE;
      Result = TRUE;
      goto Finished;
      break;
    case IMAGE_FILE_MACHINE_AMD64:
      *Is32Bit = FALSE;
      Result = TRUE;
      goto Finished;
      break;
    default:
      ShowMessages(
          "err, PE file is not i386 or AMD64; thus, it's not supported "
          "in HyperDbg\n");
      Result = FALSE;
      goto Finished;
      break;
  }
Finished:
  UnmapViewOfFile(BaseAddr);
  CloseHandle(MapObjectHandle);
  return Result;
}
