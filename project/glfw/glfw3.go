package glfw

import (
	"time"
	"unsafe"
)

// Source: glfw3.h
type (
	_builtin_va_list      = *uintptr
	_predefined_size_t    = uint64
	_predefined_wchar_t   = uint16
	_predefined_ptrdiff_t = int64
	Wchar_t               = uint16
	Size_t                = uint64
	SIZE_T                = uintptr
	Ptrdiff_t             = int64
	Uintptr_t             = uint64
	Intptr_t              = int64
	Uint8_t               = uint8
	Uint16_t              = uint16
	Uint32_t              = uint32
	Uint64_t              = uint64
	Int8_t                = int8
	Int16_t               = int16
	Int32_t               = int32
	Int64_t               = int64
	Time_t                = time.Duration
	Errno_t               = int32
	_time32_t             = int32
	_time64_t             = time.Duration
	Va_list               = *int8
	PLIST_ENTRY           = *LIST_ENTRY
	BOOL                  = int32
	BOOLEAN               = bool
	DWORD                 = uint32
	WORD                  = uint16
	PVOID                 = uintptr
	LPVOID                = uintptr
	HANDLE                = uintptr
	QWORD                 = uint64
	UINT64                = uint64
	PUINT64               = *uint64
	UINT32                = uint32
	PUINT32               = *uint32
	UINT16                = uint16
	PUINT16               = *uint16
	UINT8                 = uint8
	PUINT8                = *uint8
	INT64                 = int64
	PINT64                = *int64
	INT32                 = int32
	PINT32                = *int32
	INT16                 = int16
	PINT16                = *int16
	INT8                  = int8
	PINT8                 = *int8
	ULONG64               = uint64
	PULONG64              = *uint64
	DWORD64               = uint64
	PDWORD64              = *uint64
	CHAR                  = int8
	WCHAR                 = uint16
	UCHAR                 = uint8
	SHORT                 = int16
	USHORT                = uint16
	INT                   = int32
	UINT                  = uint32
	PUINT                 = *uint32
	LONG                  = int32
	ULONG                 = uint32
	Int_least8_t          = int8
	Int_least16_t         = int16
	Int_least32_t         = int32
	Int_least64_t         = int64
	Uint_least8_t         = uint8
	Uint_least16_t        = uint16
	Uint_least32_t        = uint32
	Uint_least64_t        = uint64
	Int_fast8_t           = int8
	Int_fast16_t          = int32
	Int_fast32_t          = int32
	Int_fast64_t          = int64
	Uint_fast8_t          = uint8
	Uint_fast16_t         = uint32
	Uint_fast32_t         = uint32
	Uint_fast64_t         = uint64
	Intmax_t              = int64
	Uintmax_t             = uint64
)

type GLFWglproc func()

type GLFWvkproc func()

type GLFWallocatefun func(uintptr, unsafe.Pointer) unsafe.Pointer

type GLFWreallocatefun func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer

type GLFWdeallocatefun func(unsafe.Pointer, unsafe.Pointer)

type GLFWerrorfun func(int32, *int8)

type GLFWwindowposfun func(*GLFWwindow, int32, int32)

type GLFWwindowsizefun func(*GLFWwindow, int32, int32)

type GLFWwindowclosefun func(*GLFWwindow)

type GLFWwindowrefreshfun func(*GLFWwindow)

type GLFWwindowfocusfun func(*GLFWwindow, int32)

type GLFWwindowiconifyfun func(*GLFWwindow, int32)

type GLFWwindowmaximizefun func(*GLFWwindow, int32)

type GLFWframebuffersizefun func(*GLFWwindow, int32, int32)

type GLFWwindowcontentscalefun func(*GLFWwindow, float32, float32)

type GLFWmousebuttonfun func(*GLFWwindow, int32, int32, int32)

type GLFWcursorposfun func(*GLFWwindow, float64, float64)

type GLFWcursorenterfun func(*GLFWwindow, int32)

type GLFWscrollfun func(*GLFWwindow, float64, float64)

type GLFWkeyfun func(*GLFWwindow, int32, int32, int32, int32)

type GLFWcharfun func(*GLFWwindow, uint32)

type GLFWcharmodsfun func(*GLFWwindow, uint32, int32)

type GLFWdropfun func(*GLFWwindow, int32, **int8)

type GLFWmonitorfun func(*GLFWmonitor, int32)

type GLFWjoystickfun func(int32, int32)

// Source: unknown.h:21 -> _LIST_ENTRY
type LIST_ENTRY struct {
	Flink *LIST_ENTRY
	Blink *LIST_ENTRY
}

// Source: GLFW/glfw3.h:1396 -> GLFWmonitor
type GLFWmonitor struct{}

// Source: GLFW/glfw3.h:1408 -> GLFWwindow
type GLFWwindow struct{}

// Source: GLFW/glfw3.h:1420 -> GLFWcursor
type GLFWcursor struct{}

// Source: GLFW/glfw3.h:2031 -> GLFWvidmode
type GLFWvidmode struct {
	Width       int32
	Height      int32
	RedBits     int32
	GreenBits   int32
	BlueBits    int32
	RefreshRate int32
}

// Source: GLFW/glfw3.h:2065 -> GLFWgammaramp
type GLFWgammaramp struct {
	Red   *uint16
	Green *uint16
	Blue  *uint16
	Size  uint32
	_     [4]byte
}

// Source: GLFW/glfw3.h:2094 -> GLFWimage
type GLFWimage struct {
	Width  int32
	Height int32
	Pixels *uint8
}

// Source: GLFW/glfw3.h:2118 -> GLFWgamepadstate
type GLFWgamepadstate struct {
	Buttons [15]uint8
	_       [1]byte
	Axes    [6]float32
}

// Source: GLFW/glfw3.h:2142 -> GLFWallocator
type GLFWallocator struct {
	Allocate   GLFWallocatefun
	Reallocate GLFWreallocatefun
	Deallocate GLFWdeallocatefun
	User       unsafe.Pointer
}

// Source: glfw3.h -> Macro constants
const (
	GlfwVersionMajor              uint32 = 3
	GlfwVersionMinor              uint32 = 5
	GlfwVersionRevision           uint32 = 0
	GlfwTrue                      uint32 = 1
	GlfwFalse                     uint32 = 0
	GlfwRelease                   uint32 = 0
	GlfwPress                     uint32 = 1
	GlfwRepeat                    uint32 = 2
	GlfwHatCentered               uint32 = 0
	GlfwHatUp                     uint32 = 1
	GlfwHatRight                  uint32 = 2
	GlfwHatDown                   uint32 = 4
	GlfwHatLeft                   uint32 = 8
	GlfwHatRightUp                uint32 = (GlfwHatRight | GlfwHatUp)
	GlfwHatRightDown              uint32 = (GlfwHatRight | GlfwHatDown)
	GlfwHatLeftUp                 uint32 = (GlfwHatLeft | GlfwHatUp)
	GlfwHatLeftDown               uint32 = (GlfwHatLeft | GlfwHatDown)
	GlfwKeyUnknown                int32  = -1
	GlfwKeySpace                  uint32 = 32
	GlfwKeyApostrophe             uint32 = 39
	GlfwKeyComma                  uint32 = 44
	GlfwKeyMinus                  uint32 = 45
	GlfwKeyPeriod                 uint32 = 46
	GlfwKeySlash                  uint32 = 47
	GlfwKey0                      uint32 = 48
	GlfwKey1                      uint32 = 49
	GlfwKey2                      uint32 = 50
	GlfwKey3                      uint32 = 51
	GlfwKey4                      uint32 = 52
	GlfwKey5                      uint32 = 53
	GlfwKey6                      uint32 = 54
	GlfwKey7                      uint32 = 55
	GlfwKey8                      uint32 = 56
	GlfwKey9                      uint32 = 57
	GlfwKeySemicolon              uint32 = 59
	GlfwKeyEqual                  uint32 = 61
	GlfwKeyA                      uint32 = 65
	GlfwKeyB                      uint32 = 66
	GlfwKeyC                      uint32 = 67
	GlfwKeyD                      uint32 = 68
	GlfwKeyE                      uint32 = 69
	GlfwKeyF                      uint32 = 70
	GlfwKeyG                      uint32 = 71
	GlfwKeyH                      uint32 = 72
	GlfwKeyI                      uint32 = 73
	GlfwKeyJ                      uint32 = 74
	GlfwKeyK                      uint32 = 75
	GlfwKeyL                      uint32 = 76
	GlfwKeyM                      uint32 = 77
	GlfwKeyN                      uint32 = 78
	GlfwKeyO                      uint32 = 79
	GlfwKeyP                      uint32 = 80
	GlfwKeyQ                      uint32 = 81
	GlfwKeyR                      uint32 = 82
	GlfwKeyS                      uint32 = 83
	GlfwKeyT                      uint32 = 84
	GlfwKeyU                      uint32 = 85
	GlfwKeyV                      uint32 = 86
	GlfwKeyW                      uint32 = 87
	GlfwKeyX                      uint32 = 88
	GlfwKeyY                      uint32 = 89
	GlfwKeyZ                      uint32 = 90
	GlfwKeyLeftBracket            uint32 = 91
	GlfwKeyBackslash              uint32 = 92
	GlfwKeyRightBracket           uint32 = 93
	GlfwKeyGraveAccent            uint32 = 96
	GlfwKeyWorld1                 uint32 = 161
	GlfwKeyWorld2                 uint32 = 162
	GlfwKeyEscape                 uint32 = 256
	GlfwKeyEnter                  uint32 = 257
	GlfwKeyTab                    uint32 = 258
	GlfwKeyBackspace              uint32 = 259
	GlfwKeyInsert                 uint32 = 260
	GlfwKeyDelete                 uint32 = 261
	GlfwKeyRight                  uint32 = 262
	GlfwKeyLeft                   uint32 = 263
	GlfwKeyDown                   uint32 = 264
	GlfwKeyUp                     uint32 = 265
	GlfwKeyPageUp                 uint32 = 266
	GlfwKeyPageDown               uint32 = 267
	GlfwKeyHome                   uint32 = 268
	GlfwKeyEnd                    uint32 = 269
	GlfwKeyCapsLock               uint32 = 280
	GlfwKeyScrollLock             uint32 = 281
	GlfwKeyNumLock                uint32 = 282
	GlfwKeyPrintScreen            uint32 = 283
	GlfwKeyPause                  uint32 = 284
	GlfwKeyF1                     uint32 = 290
	GlfwKeyF2                     uint32 = 291
	GlfwKeyF3                     uint32 = 292
	GlfwKeyF4                     uint32 = 293
	GlfwKeyF5                     uint32 = 294
	GlfwKeyF6                     uint32 = 295
	GlfwKeyF7                     uint32 = 296
	GlfwKeyF8                     uint32 = 297
	GlfwKeyF9                     uint32 = 298
	GlfwKeyF10                    uint32 = 299
	GlfwKeyF11                    uint32 = 300
	GlfwKeyF12                    uint32 = 301
	GlfwKeyF13                    uint32 = 302
	GlfwKeyF14                    uint32 = 303
	GlfwKeyF15                    uint32 = 304
	GlfwKeyF16                    uint32 = 305
	GlfwKeyF17                    uint32 = 306
	GlfwKeyF18                    uint32 = 307
	GlfwKeyF19                    uint32 = 308
	GlfwKeyF20                    uint32 = 309
	GlfwKeyF21                    uint32 = 310
	GlfwKeyF22                    uint32 = 311
	GlfwKeyF23                    uint32 = 312
	GlfwKeyF24                    uint32 = 313
	GlfwKeyF25                    uint32 = 314
	GlfwKeyKp0                    uint32 = 320
	GlfwKeyKp1                    uint32 = 321
	GlfwKeyKp2                    uint32 = 322
	GlfwKeyKp3                    uint32 = 323
	GlfwKeyKp4                    uint32 = 324
	GlfwKeyKp5                    uint32 = 325
	GlfwKeyKp6                    uint32 = 326
	GlfwKeyKp7                    uint32 = 327
	GlfwKeyKp8                    uint32 = 328
	GlfwKeyKp9                    uint32 = 329
	GlfwKeyKpDecimal              uint32 = 330
	GlfwKeyKpDivide               uint32 = 331
	GlfwKeyKpMultiply             uint32 = 332
	GlfwKeyKpSubtract             uint32 = 333
	GlfwKeyKpAdd                  uint32 = 334
	GlfwKeyKpEnter                uint32 = 335
	GlfwKeyKpEqual                uint32 = 336
	GlfwKeyLeftShift              uint32 = 340
	GlfwKeyLeftControl            uint32 = 341
	GlfwKeyLeftAlt                uint32 = 342
	GlfwKeyLeftSuper              uint32 = 343
	GlfwKeyRightShift             uint32 = 344
	GlfwKeyRightControl           uint32 = 345
	GlfwKeyRightAlt               uint32 = 346
	GlfwKeyRightSuper             uint32 = 347
	GlfwKeyMenu                   uint32 = 348
	GlfwKeyLast                   uint32 = GlfwKeyMenu
	GlfwModShift                  uint32 = 0x0001
	GlfwModControl                uint32 = 0x0002
	GlfwModAlt                    uint32 = 0x0004
	GlfwModSuper                  uint32 = 0x0008
	GlfwModCapsLock               uint32 = 0x0010
	GlfwModNumLock                uint32 = 0x0020
	GlfwMouseButton1              uint32 = 0
	GlfwMouseButton2              uint32 = 1
	GlfwMouseButton3              uint32 = 2
	GlfwMouseButton4              uint32 = 3
	GlfwMouseButton5              uint32 = 4
	GlfwMouseButton6              uint32 = 5
	GlfwMouseButton7              uint32 = 6
	GlfwMouseButton8              uint32 = 7
	GlfwMouseButtonLast           uint32 = GlfwMouseButton8
	GlfwMouseButtonLeft           uint32 = GlfwMouseButton1
	GlfwMouseButtonRight          uint32 = GlfwMouseButton2
	GlfwMouseButtonMiddle         uint32 = GlfwMouseButton3
	GlfwJoystick1                 uint32 = 0
	GlfwJoystick2                 uint32 = 1
	GlfwJoystick3                 uint32 = 2
	GlfwJoystick4                 uint32 = 3
	GlfwJoystick5                 uint32 = 4
	GlfwJoystick6                 uint32 = 5
	GlfwJoystick7                 uint32 = 6
	GlfwJoystick8                 uint32 = 7
	GlfwJoystick9                 uint32 = 8
	GlfwJoystick10                uint32 = 9
	GlfwJoystick11                uint32 = 10
	GlfwJoystick12                uint32 = 11
	GlfwJoystick13                uint32 = 12
	GlfwJoystick14                uint32 = 13
	GlfwJoystick15                uint32 = 14
	GlfwJoystick16                uint32 = 15
	GlfwJoystickLast              uint32 = GlfwJoystick16
	GlfwGamepadButtonA            uint32 = 0
	GlfwGamepadButtonB            uint32 = 1
	GlfwGamepadButtonX            uint32 = 2
	GlfwGamepadButtonY            uint32 = 3
	GlfwGamepadButtonLeftBumper   uint32 = 4
	GlfwGamepadButtonRightBumper  uint32 = 5
	GlfwGamepadButtonBack         uint32 = 6
	GlfwGamepadButtonStart        uint32 = 7
	GlfwGamepadButtonGuide        uint32 = 8
	GlfwGamepadButtonLeftThumb    uint32 = 9
	GlfwGamepadButtonRightThumb   uint32 = 10
	GlfwGamepadButtonDpadUp       uint32 = 11
	GlfwGamepadButtonDpadRight    uint32 = 12
	GlfwGamepadButtonDpadDown     uint32 = 13
	GlfwGamepadButtonDpadLeft     uint32 = 14
	GlfwGamepadButtonLast         uint32 = GlfwGamepadButtonDpadLeft
	GlfwGamepadButtonCross        uint32 = GlfwGamepadButtonA
	GlfwGamepadButtonCircle       uint32 = GlfwGamepadButtonB
	GlfwGamepadButtonSquare       uint32 = GlfwGamepadButtonX
	GlfwGamepadButtonTriangle     uint32 = GlfwGamepadButtonY
	GlfwGamepadAxisLeftX          uint32 = 0
	GlfwGamepadAxisLeftY          uint32 = 1
	GlfwGamepadAxisRightX         uint32 = 2
	GlfwGamepadAxisRightY         uint32 = 3
	GlfwGamepadAxisLeftTrigger    uint32 = 4
	GlfwGamepadAxisRightTrigger   uint32 = 5
	GlfwGamepadAxisLast           uint32 = GlfwGamepadAxisRightTrigger
	GlfwNoError                   uint32 = 0
	GlfwNotInitialized            uint32 = 0x00010001
	GlfwNoCurrentContext          uint32 = 0x00010002
	GlfwInvalidEnum               uint32 = 0x00010003
	GlfwInvalidValue              uint32 = 0x00010004
	GlfwOutOfMemory               uint32 = 0x00010005
	GlfwApiUnavailable            uint32 = 0x00010006
	GlfwVersionUnavailable        uint32 = 0x00010007
	GlfwPlatformError             uint32 = 0x00010008
	GlfwFormatUnavailable         uint32 = 0x00010009
	GlfwNoWindowContext           uint32 = 0x0001000A
	GlfwCursorUnavailable         uint32 = 0x0001000B
	GlfwFeatureUnavailable        uint32 = 0x0001000C
	GlfwFeatureUnimplemented      uint32 = 0x0001000D
	GlfwPlatformUnavailable       uint32 = 0x0001000E
	GlfwFocused                   uint32 = 0x00020001
	GlfwIconified                 uint32 = 0x00020002
	GlfwResizable                 uint32 = 0x00020003
	GlfwVisible                   uint32 = 0x00020004
	GlfwDecorated                 uint32 = 0x00020005
	GlfwAutoIconify               uint32 = 0x00020006
	GlfwFloating                  uint32 = 0x00020007
	GlfwMaximized                 uint32 = 0x00020008
	GlfwCenterCursor              uint32 = 0x00020009
	GlfwTransparentFramebuffer    uint32 = 0x0002000A
	GlfwHovered                   uint32 = 0x0002000B
	GlfwFocusOnShow               uint32 = 0x0002000C
	GlfwMousePassthrough          uint32 = 0x0002000D
	GlfwPositionX                 uint32 = 0x0002000E
	GlfwPositionY                 uint32 = 0x0002000F
	GlfwRedBits                   uint32 = 0x00021001
	GlfwGreenBits                 uint32 = 0x00021002
	GlfwBlueBits                  uint32 = 0x00021003
	GlfwAlphaBits                 uint32 = 0x00021004
	GlfwDepthBits                 uint32 = 0x00021005
	GlfwStencilBits               uint32 = 0x00021006
	GlfwAccumRedBits              uint32 = 0x00021007
	GlfwAccumGreenBits            uint32 = 0x00021008
	GlfwAccumBlueBits             uint32 = 0x00021009
	GlfwAccumAlphaBits            uint32 = 0x0002100A
	GlfwAuxBuffers                uint32 = 0x0002100B
	GlfwStereo                    uint32 = 0x0002100C
	GlfwSamples                   uint32 = 0x0002100D
	GlfwSrgbCapable               uint32 = 0x0002100E
	GlfwRefreshRate               uint32 = 0x0002100F
	GlfwDoublebuffer              uint32 = 0x00021010
	GlfwClientApi                 uint32 = 0x00022001
	GlfwContextVersionMajor       uint32 = 0x00022002
	GlfwContextVersionMinor       uint32 = 0x00022003
	GlfwContextRevision           uint32 = 0x00022004
	GlfwContextRobustness         uint32 = 0x00022005
	GlfwOpenglForwardCompat       uint32 = 0x00022006
	GlfwContextDebug              uint32 = 0x00022007
	GlfwOpenglDebugContext        uint32 = GlfwContextDebug
	GlfwOpenglProfile             uint32 = 0x00022008
	GlfwContextReleaseBehavior    uint32 = 0x00022009
	GlfwContextNoError            uint32 = 0x0002200A
	GlfwContextCreationApi        uint32 = 0x0002200B
	GlfwScaleToMonitor            uint32 = 0x0002200C
	GlfwScaleFramebuffer          uint32 = 0x0002200D
	GlfwCocoaRetinaFramebuffer    uint32 = 0x00023001
	GlfwCocoaFrameName            uint32 = 0x00023002
	GlfwCocoaGraphicsSwitching    uint32 = 0x00023003
	GlfwX11ClassName              uint32 = 0x00024001
	GlfwX11InstanceName           uint32 = 0x00024002
	GlfwWin32KeyboardMenu         uint32 = 0x00025001
	GlfwWin32Showdefault          uint32 = 0x00025002
	GlfwWaylandAppId              uint32 = 0x00026001
	GlfwNoApi                     uint32 = 0
	GlfwOpenglApi                 uint32 = 0x00030001
	GlfwOpenglEsApi               uint32 = 0x00030002
	GlfwNoRobustness              uint32 = 0
	GlfwNoResetNotification       uint32 = 0x00031001
	GlfwLoseContextOnReset        uint32 = 0x00031002
	GlfwOpenglAnyProfile          uint32 = 0
	GlfwOpenglCoreProfile         uint32 = 0x00032001
	GlfwOpenglCompatProfile       uint32 = 0x00032002
	GlfwCursor                    uint32 = 0x00033001
	GlfwStickyKeys                uint32 = 0x00033002
	GlfwStickyMouseButtons        uint32 = 0x00033003
	GlfwLockKeyMods               uint32 = 0x00033004
	GlfwRawMouseMotion            uint32 = 0x00033005
	GlfwUnlimitedMouseButtons     uint32 = 0x00033006
	GlfwCursorNormal              uint32 = 0x00034001
	GlfwCursorHidden              uint32 = 0x00034002
	GlfwCursorDisabled            uint32 = 0x00034003
	GlfwCursorCaptured            uint32 = 0x00034004
	GlfwAnyReleaseBehavior        uint32 = 0
	GlfwReleaseBehaviorFlush      uint32 = 0x00035001
	GlfwReleaseBehaviorNone       uint32 = 0x00035002
	GlfwNativeContextApi          uint32 = 0x00036001
	GlfwEglContextApi             uint32 = 0x00036002
	GlfwOsmesaContextApi          uint32 = 0x00036003
	GlfwAnglePlatformTypeNone     uint32 = 0x00037001
	GlfwAnglePlatformTypeOpengl   uint32 = 0x00037002
	GlfwAnglePlatformTypeOpengles uint32 = 0x00037003
	GlfwAnglePlatformTypeD3d9     uint32 = 0x00037004
	GlfwAnglePlatformTypeD3d11    uint32 = 0x00037005
	GlfwAnglePlatformTypeVulkan   uint32 = 0x00037007
	GlfwAnglePlatformTypeMetal    uint32 = 0x00037008
	GlfwWaylandPreferLibdecor     uint32 = 0x00038001
	GlfwWaylandDisableLibdecor    uint32 = 0x00038002
	GlfwAnyPosition               uint32 = 0x80000000
	GlfwArrowCursor               uint32 = 0x00036001
	GlfwIbeamCursor               uint32 = 0x00036002
	GlfwCrosshairCursor           uint32 = 0x00036003
	GlfwPointingHandCursor        uint32 = 0x00036004
	GlfwResizeEwCursor            uint32 = 0x00036005
	GlfwResizeNsCursor            uint32 = 0x00036006
	GlfwResizeNwseCursor          uint32 = 0x00036007
	GlfwResizeNeswCursor          uint32 = 0x00036008
	GlfwResizeAllCursor           uint32 = 0x00036009
	GlfwNotAllowedCursor          uint32 = 0x0003600A
	GlfwHresizeCursor             uint32 = GlfwResizeEwCursor
	GlfwVresizeCursor             uint32 = GlfwResizeNsCursor
	GlfwConnected                 uint32 = 0x00040001
	GlfwDisconnected              uint32 = 0x00040002
	GlfwJoystickHatButtons        uint32 = 0x00050001
	GlfwAnglePlatformType         uint32 = 0x00050002
	GlfwPlatform                  uint32 = 0x00050003
	GlfwCocoaChdirResources       uint32 = 0x00051001
	GlfwCocoaMenubar              uint32 = 0x00051002
	GlfwX11XcbVulkanSurface       uint32 = 0x00052001
	GlfwWaylandLibdecor           uint32 = 0x00053001
	GlfwAnyPlatform               uint32 = 0x00060000
	GlfwPlatformWin32             uint32 = 0x00060001
	GlfwPlatformCocoa             uint32 = 0x00060002
	GlfwPlatformWayland           uint32 = 0x00060003
	GlfwPlatformX11               uint32 = 0x00060004
	GlfwPlatformNull              uint32 = 0x00060005
	GlfwDontCare                  int32  = -1
	GLAPIENTRY                    uint32 = 0
)
