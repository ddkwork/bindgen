package glfw

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Glfw struct{}

//go:embed glfw3.dll
var dllBytes []byte

var (
	dll                                    *windows.LazyDLL
	proc_glfwInit                          *windows.LazyProc
	proc_glfwTerminate                     *windows.LazyProc
	proc_glfwInitHint                      *windows.LazyProc
	proc_glfwInitAllocator                 *windows.LazyProc
	proc_glfwGetVersion                    *windows.LazyProc
	proc_glfwGetVersionString              *windows.LazyProc
	proc_glfwGetError                      *windows.LazyProc
	proc_glfwSetErrorCallback              *windows.LazyProc
	proc_glfwGetPlatform                   *windows.LazyProc
	proc_glfwPlatformSupported             *windows.LazyProc
	proc_glfwGetMonitors                   *windows.LazyProc
	proc_glfwGetPrimaryMonitor             *windows.LazyProc
	proc_glfwGetMonitorPos                 *windows.LazyProc
	proc_glfwGetMonitorWorkarea            *windows.LazyProc
	proc_glfwGetMonitorPhysicalSize        *windows.LazyProc
	proc_glfwGetMonitorContentScale        *windows.LazyProc
	proc_glfwGetMonitorName                *windows.LazyProc
	proc_glfwSetMonitorUserPointer         *windows.LazyProc
	proc_glfwGetMonitorUserPointer         *windows.LazyProc
	proc_glfwSetMonitorCallback            *windows.LazyProc
	proc_glfwGetVideoModes                 *windows.LazyProc
	proc_glfwGetVideoMode                  *windows.LazyProc
	proc_glfwSetGamma                      *windows.LazyProc
	proc_glfwGetGammaRamp                  *windows.LazyProc
	proc_glfwSetGammaRamp                  *windows.LazyProc
	proc_glfwDefaultWindowHints            *windows.LazyProc
	proc_glfwWindowHint                    *windows.LazyProc
	proc_glfwWindowHintString              *windows.LazyProc
	proc_glfwCreateWindow                  *windows.LazyProc
	proc_glfwDestroyWindow                 *windows.LazyProc
	proc_glfwWindowShouldClose             *windows.LazyProc
	proc_glfwSetWindowShouldClose          *windows.LazyProc
	proc_glfwGetWindowTitle                *windows.LazyProc
	proc_glfwSetWindowTitle                *windows.LazyProc
	proc_glfwSetWindowIcon                 *windows.LazyProc
	proc_glfwGetWindowPos                  *windows.LazyProc
	proc_glfwSetWindowPos                  *windows.LazyProc
	proc_glfwGetWindowSize                 *windows.LazyProc
	proc_glfwSetWindowSizeLimits           *windows.LazyProc
	proc_glfwSetWindowAspectRatio          *windows.LazyProc
	proc_glfwSetWindowSize                 *windows.LazyProc
	proc_glfwGetFramebufferSize            *windows.LazyProc
	proc_glfwGetWindowFrameSize            *windows.LazyProc
	proc_glfwGetWindowContentScale         *windows.LazyProc
	proc_glfwGetWindowOpacity              *windows.LazyProc
	proc_glfwSetWindowOpacity              *windows.LazyProc
	proc_glfwIconifyWindow                 *windows.LazyProc
	proc_glfwRestoreWindow                 *windows.LazyProc
	proc_glfwMaximizeWindow                *windows.LazyProc
	proc_glfwShowWindow                    *windows.LazyProc
	proc_glfwHideWindow                    *windows.LazyProc
	proc_glfwFocusWindow                   *windows.LazyProc
	proc_glfwRequestWindowAttention        *windows.LazyProc
	proc_glfwGetWindowMonitor              *windows.LazyProc
	proc_glfwSetWindowMonitor              *windows.LazyProc
	proc_glfwGetWindowAttrib               *windows.LazyProc
	proc_glfwSetWindowAttrib               *windows.LazyProc
	proc_glfwSetWindowUserPointer          *windows.LazyProc
	proc_glfwGetWindowUserPointer          *windows.LazyProc
	proc_glfwSetWindowPosCallback          *windows.LazyProc
	proc_glfwSetWindowSizeCallback         *windows.LazyProc
	proc_glfwSetWindowCloseCallback        *windows.LazyProc
	proc_glfwSetWindowRefreshCallback      *windows.LazyProc
	proc_glfwSetWindowFocusCallback        *windows.LazyProc
	proc_glfwSetWindowIconifyCallback      *windows.LazyProc
	proc_glfwSetWindowMaximizeCallback     *windows.LazyProc
	proc_glfwSetFramebufferSizeCallback    *windows.LazyProc
	proc_glfwSetWindowContentScaleCallback *windows.LazyProc
	proc_glfwPollEvents                    *windows.LazyProc
	proc_glfwWaitEvents                    *windows.LazyProc
	proc_glfwWaitEventsTimeout             *windows.LazyProc
	proc_glfwPostEmptyEvent                *windows.LazyProc
	proc_glfwGetInputMode                  *windows.LazyProc
	proc_glfwSetInputMode                  *windows.LazyProc
	proc_glfwRawMouseMotionSupported       *windows.LazyProc
	proc_glfwGetKeyName                    *windows.LazyProc
	proc_glfwGetKeyScancode                *windows.LazyProc
	proc_glfwGetKey                        *windows.LazyProc
	proc_glfwGetMouseButton                *windows.LazyProc
	proc_glfwGetCursorPos                  *windows.LazyProc
	proc_glfwSetCursorPos                  *windows.LazyProc
	proc_glfwCreateCursor                  *windows.LazyProc
	proc_glfwCreateStandardCursor          *windows.LazyProc
	proc_glfwDestroyCursor                 *windows.LazyProc
	proc_glfwSetCursor                     *windows.LazyProc
	proc_glfwSetKeyCallback                *windows.LazyProc
	proc_glfwSetCharCallback               *windows.LazyProc
	proc_glfwSetCharModsCallback           *windows.LazyProc
	proc_glfwSetMouseButtonCallback        *windows.LazyProc
	proc_glfwSetCursorPosCallback          *windows.LazyProc
	proc_glfwSetCursorEnterCallback        *windows.LazyProc
	proc_glfwSetScrollCallback             *windows.LazyProc
	proc_glfwSetDropCallback               *windows.LazyProc
	proc_glfwJoystickPresent               *windows.LazyProc
	proc_glfwGetJoystickAxes               *windows.LazyProc
	proc_glfwGetJoystickButtons            *windows.LazyProc
	proc_glfwGetJoystickHats               *windows.LazyProc
	proc_glfwGetJoystickName               *windows.LazyProc
	proc_glfwGetJoystickGUID               *windows.LazyProc
	proc_glfwSetJoystickUserPointer        *windows.LazyProc
	proc_glfwGetJoystickUserPointer        *windows.LazyProc
	proc_glfwJoystickIsGamepad             *windows.LazyProc
	proc_glfwSetJoystickCallback           *windows.LazyProc
	proc_glfwUpdateGamepadMappings         *windows.LazyProc
	proc_glfwGetGamepadName                *windows.LazyProc
	proc_glfwGetGamepadState               *windows.LazyProc
	proc_glfwSetClipboardString            *windows.LazyProc
	proc_glfwGetClipboardString            *windows.LazyProc
	proc_glfwGetTime                       *windows.LazyProc
	proc_glfwSetTime                       *windows.LazyProc
	proc_glfwGetTimerValue                 *windows.LazyProc
	proc_glfwGetTimerFrequency             *windows.LazyProc
	proc_glfwMakeContextCurrent            *windows.LazyProc
	proc_glfwGetCurrentContext             *windows.LazyProc
	proc_glfwSwapBuffers                   *windows.LazyProc
	proc_glfwSwapInterval                  *windows.LazyProc
	proc_glfwExtensionSupported            *windows.LazyProc
	proc_glfwGetProcAddress                *windows.LazyProc
	proc_glfwVulkanSupported               *windows.LazyProc
	proc_glfwGetRequiredInstanceExtensions *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "glfw3.dll"))
	proc_glfwInit = dll.NewProc("glfwInit")
	proc_glfwTerminate = dll.NewProc("glfwTerminate")
	proc_glfwInitHint = dll.NewProc("glfwInitHint")
	proc_glfwInitAllocator = dll.NewProc("glfwInitAllocator")
	proc_glfwGetVersion = dll.NewProc("glfwGetVersion")
	proc_glfwGetVersionString = dll.NewProc("glfwGetVersionString")
	proc_glfwGetError = dll.NewProc("glfwGetError")
	proc_glfwSetErrorCallback = dll.NewProc("glfwSetErrorCallback")
	proc_glfwGetPlatform = dll.NewProc("glfwGetPlatform")
	proc_glfwPlatformSupported = dll.NewProc("glfwPlatformSupported")
	proc_glfwGetMonitors = dll.NewProc("glfwGetMonitors")
	proc_glfwGetPrimaryMonitor = dll.NewProc("glfwGetPrimaryMonitor")
	proc_glfwGetMonitorPos = dll.NewProc("glfwGetMonitorPos")
	proc_glfwGetMonitorWorkarea = dll.NewProc("glfwGetMonitorWorkarea")
	proc_glfwGetMonitorPhysicalSize = dll.NewProc("glfwGetMonitorPhysicalSize")
	proc_glfwGetMonitorContentScale = dll.NewProc("glfwGetMonitorContentScale")
	proc_glfwGetMonitorName = dll.NewProc("glfwGetMonitorName")
	proc_glfwSetMonitorUserPointer = dll.NewProc("glfwSetMonitorUserPointer")
	proc_glfwGetMonitorUserPointer = dll.NewProc("glfwGetMonitorUserPointer")
	proc_glfwSetMonitorCallback = dll.NewProc("glfwSetMonitorCallback")
	proc_glfwGetVideoModes = dll.NewProc("glfwGetVideoModes")
	proc_glfwGetVideoMode = dll.NewProc("glfwGetVideoMode")
	proc_glfwSetGamma = dll.NewProc("glfwSetGamma")
	proc_glfwGetGammaRamp = dll.NewProc("glfwGetGammaRamp")
	proc_glfwSetGammaRamp = dll.NewProc("glfwSetGammaRamp")
	proc_glfwDefaultWindowHints = dll.NewProc("glfwDefaultWindowHints")
	proc_glfwWindowHint = dll.NewProc("glfwWindowHint")
	proc_glfwWindowHintString = dll.NewProc("glfwWindowHintString")
	proc_glfwCreateWindow = dll.NewProc("glfwCreateWindow")
	proc_glfwDestroyWindow = dll.NewProc("glfwDestroyWindow")
	proc_glfwWindowShouldClose = dll.NewProc("glfwWindowShouldClose")
	proc_glfwSetWindowShouldClose = dll.NewProc("glfwSetWindowShouldClose")
	proc_glfwGetWindowTitle = dll.NewProc("glfwGetWindowTitle")
	proc_glfwSetWindowTitle = dll.NewProc("glfwSetWindowTitle")
	proc_glfwSetWindowIcon = dll.NewProc("glfwSetWindowIcon")
	proc_glfwGetWindowPos = dll.NewProc("glfwGetWindowPos")
	proc_glfwSetWindowPos = dll.NewProc("glfwSetWindowPos")
	proc_glfwGetWindowSize = dll.NewProc("glfwGetWindowSize")
	proc_glfwSetWindowSizeLimits = dll.NewProc("glfwSetWindowSizeLimits")
	proc_glfwSetWindowAspectRatio = dll.NewProc("glfwSetWindowAspectRatio")
	proc_glfwSetWindowSize = dll.NewProc("glfwSetWindowSize")
	proc_glfwGetFramebufferSize = dll.NewProc("glfwGetFramebufferSize")
	proc_glfwGetWindowFrameSize = dll.NewProc("glfwGetWindowFrameSize")
	proc_glfwGetWindowContentScale = dll.NewProc("glfwGetWindowContentScale")
	proc_glfwGetWindowOpacity = dll.NewProc("glfwGetWindowOpacity")
	proc_glfwSetWindowOpacity = dll.NewProc("glfwSetWindowOpacity")
	proc_glfwIconifyWindow = dll.NewProc("glfwIconifyWindow")
	proc_glfwRestoreWindow = dll.NewProc("glfwRestoreWindow")
	proc_glfwMaximizeWindow = dll.NewProc("glfwMaximizeWindow")
	proc_glfwShowWindow = dll.NewProc("glfwShowWindow")
	proc_glfwHideWindow = dll.NewProc("glfwHideWindow")
	proc_glfwFocusWindow = dll.NewProc("glfwFocusWindow")
	proc_glfwRequestWindowAttention = dll.NewProc("glfwRequestWindowAttention")
	proc_glfwGetWindowMonitor = dll.NewProc("glfwGetWindowMonitor")
	proc_glfwSetWindowMonitor = dll.NewProc("glfwSetWindowMonitor")
	proc_glfwGetWindowAttrib = dll.NewProc("glfwGetWindowAttrib")
	proc_glfwSetWindowAttrib = dll.NewProc("glfwSetWindowAttrib")
	proc_glfwSetWindowUserPointer = dll.NewProc("glfwSetWindowUserPointer")
	proc_glfwGetWindowUserPointer = dll.NewProc("glfwGetWindowUserPointer")
	proc_glfwSetWindowPosCallback = dll.NewProc("glfwSetWindowPosCallback")
	proc_glfwSetWindowSizeCallback = dll.NewProc("glfwSetWindowSizeCallback")
	proc_glfwSetWindowCloseCallback = dll.NewProc("glfwSetWindowCloseCallback")
	proc_glfwSetWindowRefreshCallback = dll.NewProc("glfwSetWindowRefreshCallback")
	proc_glfwSetWindowFocusCallback = dll.NewProc("glfwSetWindowFocusCallback")
	proc_glfwSetWindowIconifyCallback = dll.NewProc("glfwSetWindowIconifyCallback")
	proc_glfwSetWindowMaximizeCallback = dll.NewProc("glfwSetWindowMaximizeCallback")
	proc_glfwSetFramebufferSizeCallback = dll.NewProc("glfwSetFramebufferSizeCallback")
	proc_glfwSetWindowContentScaleCallback = dll.NewProc("glfwSetWindowContentScaleCallback")
	proc_glfwPollEvents = dll.NewProc("glfwPollEvents")
	proc_glfwWaitEvents = dll.NewProc("glfwWaitEvents")
	proc_glfwWaitEventsTimeout = dll.NewProc("glfwWaitEventsTimeout")
	proc_glfwPostEmptyEvent = dll.NewProc("glfwPostEmptyEvent")
	proc_glfwGetInputMode = dll.NewProc("glfwGetInputMode")
	proc_glfwSetInputMode = dll.NewProc("glfwSetInputMode")
	proc_glfwRawMouseMotionSupported = dll.NewProc("glfwRawMouseMotionSupported")
	proc_glfwGetKeyName = dll.NewProc("glfwGetKeyName")
	proc_glfwGetKeyScancode = dll.NewProc("glfwGetKeyScancode")
	proc_glfwGetKey = dll.NewProc("glfwGetKey")
	proc_glfwGetMouseButton = dll.NewProc("glfwGetMouseButton")
	proc_glfwGetCursorPos = dll.NewProc("glfwGetCursorPos")
	proc_glfwSetCursorPos = dll.NewProc("glfwSetCursorPos")
	proc_glfwCreateCursor = dll.NewProc("glfwCreateCursor")
	proc_glfwCreateStandardCursor = dll.NewProc("glfwCreateStandardCursor")
	proc_glfwDestroyCursor = dll.NewProc("glfwDestroyCursor")
	proc_glfwSetCursor = dll.NewProc("glfwSetCursor")
	proc_glfwSetKeyCallback = dll.NewProc("glfwSetKeyCallback")
	proc_glfwSetCharCallback = dll.NewProc("glfwSetCharCallback")
	proc_glfwSetCharModsCallback = dll.NewProc("glfwSetCharModsCallback")
	proc_glfwSetMouseButtonCallback = dll.NewProc("glfwSetMouseButtonCallback")
	proc_glfwSetCursorPosCallback = dll.NewProc("glfwSetCursorPosCallback")
	proc_glfwSetCursorEnterCallback = dll.NewProc("glfwSetCursorEnterCallback")
	proc_glfwSetScrollCallback = dll.NewProc("glfwSetScrollCallback")
	proc_glfwSetDropCallback = dll.NewProc("glfwSetDropCallback")
	proc_glfwJoystickPresent = dll.NewProc("glfwJoystickPresent")
	proc_glfwGetJoystickAxes = dll.NewProc("glfwGetJoystickAxes")
	proc_glfwGetJoystickButtons = dll.NewProc("glfwGetJoystickButtons")
	proc_glfwGetJoystickHats = dll.NewProc("glfwGetJoystickHats")
	proc_glfwGetJoystickName = dll.NewProc("glfwGetJoystickName")
	proc_glfwGetJoystickGUID = dll.NewProc("glfwGetJoystickGUID")
	proc_glfwSetJoystickUserPointer = dll.NewProc("glfwSetJoystickUserPointer")
	proc_glfwGetJoystickUserPointer = dll.NewProc("glfwGetJoystickUserPointer")
	proc_glfwJoystickIsGamepad = dll.NewProc("glfwJoystickIsGamepad")
	proc_glfwSetJoystickCallback = dll.NewProc("glfwSetJoystickCallback")
	proc_glfwUpdateGamepadMappings = dll.NewProc("glfwUpdateGamepadMappings")
	proc_glfwGetGamepadName = dll.NewProc("glfwGetGamepadName")
	proc_glfwGetGamepadState = dll.NewProc("glfwGetGamepadState")
	proc_glfwSetClipboardString = dll.NewProc("glfwSetClipboardString")
	proc_glfwGetClipboardString = dll.NewProc("glfwGetClipboardString")
	proc_glfwGetTime = dll.NewProc("glfwGetTime")
	proc_glfwSetTime = dll.NewProc("glfwSetTime")
	proc_glfwGetTimerValue = dll.NewProc("glfwGetTimerValue")
	proc_glfwGetTimerFrequency = dll.NewProc("glfwGetTimerFrequency")
	proc_glfwMakeContextCurrent = dll.NewProc("glfwMakeContextCurrent")
	proc_glfwGetCurrentContext = dll.NewProc("glfwGetCurrentContext")
	proc_glfwSwapBuffers = dll.NewProc("glfwSwapBuffers")
	proc_glfwSwapInterval = dll.NewProc("glfwSwapInterval")
	proc_glfwExtensionSupported = dll.NewProc("glfwExtensionSupported")
	proc_glfwGetProcAddress = dll.NewProc("glfwGetProcAddress")
	proc_glfwVulkanSupported = dll.NewProc("glfwVulkanSupported")
	proc_glfwGetRequiredInstanceExtensions = dll.NewProc("glfwGetRequiredInstanceExtensions")
}

func saveEmbeddedDLL(data []byte, name string) string {
	tmpDir := os.TempDir()
	p := filepath.Join(tmpDir, name)
	os.WriteFile(p, data, 0o644)
	return p
}

func (g *Glfw) Init() int32 {
	r1, _, _ := proc_glfwInit.Call()
	return int32(r1)
}

func (g *Glfw) Terminate() {
	proc_glfwTerminate.Call()
}

func (g *Glfw) InitHint(Hint int32, Value int32) {
	proc_glfwInitHint.Call(uintptr(Hint), uintptr(Value))
}

func (g *Glfw) InitAllocator(Allocator *GLFWallocator) {
	proc_glfwInitAllocator.Call(uintptr(unsafe.Pointer(Allocator)))
}

func (g *Glfw) GetVersion(Major *int32, Minor *int32, Rev *int32) {
	proc_glfwGetVersion.Call(uintptr(unsafe.Pointer(Major)), uintptr(unsafe.Pointer(Minor)), uintptr(unsafe.Pointer(Rev)))
}

func (g *Glfw) GetVersionString() *int8 {
	r1, _, _ := proc_glfwGetVersionString.Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetError(Description **int8) int32 {
	r1, _, _ := proc_glfwGetError.Call(uintptr(unsafe.Pointer(Description)))
	return int32(r1)
}

func (g *Glfw) SetErrorCallback(Callback GLFWerrorfun) GLFWerrorfun {
	r1, _, _ := proc_glfwSetErrorCallback.Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWerrorfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetPlatform() int32 {
	r1, _, _ := proc_glfwGetPlatform.Call()
	return int32(r1)
}

func (g *Glfw) PlatformSupported(Platform int32) int32 {
	r1, _, _ := proc_glfwPlatformSupported.Call(uintptr(Platform))
	return int32(r1)
}

func (g *Glfw) GetMonitors(Count *int32) **GLFWmonitor {
	r1, _, _ := proc_glfwGetMonitors.Call(uintptr(unsafe.Pointer(Count)))
	return (**GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) GetPrimaryMonitor() *GLFWmonitor {
	r1, _, _ := proc_glfwGetPrimaryMonitor.Call()
	return (*GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) GetMonitorPos(Monitor *GLFWmonitor, Xpos *int32, Ypos *int32) {
	proc_glfwGetMonitorPos.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) GetMonitorWorkarea(Monitor *GLFWmonitor, Xpos *int32, Ypos *int32, Width *int32, Height *int32) {
	proc_glfwGetMonitorWorkarea.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) GetMonitorPhysicalSize(Monitor *GLFWmonitor, WidthMM *int32, HeightMM *int32) {
	proc_glfwGetMonitorPhysicalSize.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(WidthMM)), uintptr(unsafe.Pointer(HeightMM)))
}

func (g *Glfw) GetMonitorContentScale(Monitor *GLFWmonitor, Xscale *float32, Yscale *float32) {
	proc_glfwGetMonitorContentScale.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xscale)), uintptr(unsafe.Pointer(Yscale)))
}

func (g *Glfw) GetMonitorName(Monitor *GLFWmonitor) *int8 {
	r1, _, _ := proc_glfwGetMonitorName.Call(uintptr(unsafe.Pointer(Monitor)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetMonitorUserPointer(Monitor *GLFWmonitor, Pointer unsafe.Pointer) {
	proc_glfwSetMonitorUserPointer.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(Pointer))
}

func (g *Glfw) GetMonitorUserPointer(Monitor *GLFWmonitor) unsafe.Pointer {
	r1, _, _ := proc_glfwGetMonitorUserPointer.Call(uintptr(unsafe.Pointer(Monitor)))
	return unsafe.Pointer(r1)
}

func (g *Glfw) SetMonitorCallback(Callback GLFWmonitorfun) GLFWmonitorfun {
	r1, _, _ := proc_glfwSetMonitorCallback.Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWmonitorfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetVideoModes(Monitor *GLFWmonitor, Count *int32) *GLFWvidmode {
	r1, _, _ := proc_glfwGetVideoModes.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Count)))
	return (*GLFWvidmode)(unsafe.Pointer(r1))
}

func (g *Glfw) GetVideoMode(Monitor *GLFWmonitor) *GLFWvidmode {
	r1, _, _ := proc_glfwGetVideoMode.Call(uintptr(unsafe.Pointer(Monitor)))
	return (*GLFWvidmode)(unsafe.Pointer(r1))
}

func (g *Glfw) SetGamma(Monitor *GLFWmonitor, Gamma float32) {
	proc_glfwSetGamma.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(Gamma))
}

func (g *Glfw) GetGammaRamp(Monitor *GLFWmonitor) *GLFWgammaramp {
	r1, _, _ := proc_glfwGetGammaRamp.Call(uintptr(unsafe.Pointer(Monitor)))
	return (*GLFWgammaramp)(unsafe.Pointer(r1))
}

func (g *Glfw) SetGammaRamp(Monitor *GLFWmonitor, Ramp *GLFWgammaramp) {
	proc_glfwSetGammaRamp.Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Ramp)))
}

func (g *Glfw) DefaultWindowHints() {
	proc_glfwDefaultWindowHints.Call()
}

func (g *Glfw) WindowHint(Hint int32, Value int32) {
	proc_glfwWindowHint.Call(uintptr(Hint), uintptr(Value))
}

func (g *Glfw) WindowHintString(Hint int32, Value *int8) {
	proc_glfwWindowHintString.Call(uintptr(Hint), uintptr(unsafe.Pointer(Value)))
}

func (g *Glfw) CreateWindow(Width int32, Height int32, Title *int8, Monitor *GLFWmonitor, Share *GLFWwindow) *GLFWwindow {
	r1, _, _ := proc_glfwCreateWindow.Call(uintptr(Width), uintptr(Height), uintptr(unsafe.Pointer(Title)), uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Share)))
	return (*GLFWwindow)(unsafe.Pointer(r1))
}

func (g *Glfw) DestroyWindow(Window *GLFWwindow) {
	proc_glfwDestroyWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) WindowShouldClose(Window *GLFWwindow) int32 {
	r1, _, _ := proc_glfwWindowShouldClose.Call(uintptr(unsafe.Pointer(Window)))
	return int32(r1)
}

func (g *Glfw) SetWindowShouldClose(Window *GLFWwindow, Value int32) {
	proc_glfwSetWindowShouldClose.Call(uintptr(unsafe.Pointer(Window)), uintptr(Value))
}

func (g *Glfw) GetWindowTitle(Window *GLFWwindow) *int8 {
	r1, _, _ := proc_glfwGetWindowTitle.Call(uintptr(unsafe.Pointer(Window)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetWindowTitle(Window *GLFWwindow, Title *int8) {
	proc_glfwSetWindowTitle.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Title)))
}

func (g *Glfw) SetWindowIcon(Window *GLFWwindow, Count int32, Images *GLFWimage) {
	proc_glfwSetWindowIcon.Call(uintptr(unsafe.Pointer(Window)), uintptr(Count), uintptr(unsafe.Pointer(Images)))
}

func (g *Glfw) GetWindowPos(Window *GLFWwindow, Xpos *int32, Ypos *int32) {
	proc_glfwGetWindowPos.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) SetWindowPos(Window *GLFWwindow, Xpos int32, Ypos int32) {
	proc_glfwSetWindowPos.Call(uintptr(unsafe.Pointer(Window)), uintptr(Xpos), uintptr(Ypos))
}

func (g *Glfw) GetWindowSize(Window *GLFWwindow, Width *int32, Height *int32) {
	proc_glfwGetWindowSize.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) SetWindowSizeLimits(Window *GLFWwindow, Minwidth int32, Minheight int32, Maxwidth int32, Maxheight int32) {
	proc_glfwSetWindowSizeLimits.Call(uintptr(unsafe.Pointer(Window)), uintptr(Minwidth), uintptr(Minheight), uintptr(Maxwidth), uintptr(Maxheight))
}

func (g *Glfw) SetWindowAspectRatio(Window *GLFWwindow, Numer int32, Denom int32) {
	proc_glfwSetWindowAspectRatio.Call(uintptr(unsafe.Pointer(Window)), uintptr(Numer), uintptr(Denom))
}

func (g *Glfw) SetWindowSize(Window *GLFWwindow, Width int32, Height int32) {
	proc_glfwSetWindowSize.Call(uintptr(unsafe.Pointer(Window)), uintptr(Width), uintptr(Height))
}

func (g *Glfw) GetFramebufferSize(Window *GLFWwindow, Width *int32, Height *int32) {
	proc_glfwGetFramebufferSize.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) GetWindowFrameSize(Window *GLFWwindow, Left *int32, Top *int32, Right *int32, Bottom *int32) {
	proc_glfwGetWindowFrameSize.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Left)), uintptr(unsafe.Pointer(Top)), uintptr(unsafe.Pointer(Right)), uintptr(unsafe.Pointer(Bottom)))
}

func (g *Glfw) GetWindowContentScale(Window *GLFWwindow, Xscale *float32, Yscale *float32) {
	proc_glfwGetWindowContentScale.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xscale)), uintptr(unsafe.Pointer(Yscale)))
}

func (g *Glfw) GetWindowOpacity(Window *GLFWwindow) float32 {
	r1, _, _ := proc_glfwGetWindowOpacity.Call(uintptr(unsafe.Pointer(Window)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowOpacity(Window *GLFWwindow, Opacity float32) {
	proc_glfwSetWindowOpacity.Call(uintptr(unsafe.Pointer(Window)), uintptr(Opacity))
}

func (g *Glfw) IconifyWindow(Window *GLFWwindow) {
	proc_glfwIconifyWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) RestoreWindow(Window *GLFWwindow) {
	proc_glfwRestoreWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) MaximizeWindow(Window *GLFWwindow) {
	proc_glfwMaximizeWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) ShowWindow(Window *GLFWwindow) {
	proc_glfwShowWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) HideWindow(Window *GLFWwindow) {
	proc_glfwHideWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) FocusWindow(Window *GLFWwindow) {
	proc_glfwFocusWindow.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) RequestWindowAttention(Window *GLFWwindow) {
	proc_glfwRequestWindowAttention.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) GetWindowMonitor(Window *GLFWwindow) *GLFWmonitor {
	r1, _, _ := proc_glfwGetWindowMonitor.Call(uintptr(unsafe.Pointer(Window)))
	return (*GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) SetWindowMonitor(Window *GLFWwindow, Monitor *GLFWmonitor, Xpos int32, Ypos int32, Width int32, Height int32, RefreshRate int32) {
	proc_glfwSetWindowMonitor.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Monitor)), uintptr(Xpos), uintptr(Ypos), uintptr(Width), uintptr(Height), uintptr(RefreshRate))
}

func (g *Glfw) GetWindowAttrib(Window *GLFWwindow, Attrib int32) int32 {
	r1, _, _ := proc_glfwGetWindowAttrib.Call(uintptr(unsafe.Pointer(Window)), uintptr(Attrib))
	return int32(r1)
}

func (g *Glfw) SetWindowAttrib(Window *GLFWwindow, Attrib int32, Value int32) {
	proc_glfwSetWindowAttrib.Call(uintptr(unsafe.Pointer(Window)), uintptr(Attrib), uintptr(Value))
}

func (g *Glfw) SetWindowUserPointer(Window *GLFWwindow, Pointer unsafe.Pointer) {
	proc_glfwSetWindowUserPointer.Call(uintptr(unsafe.Pointer(Window)), uintptr(Pointer))
}

func (g *Glfw) GetWindowUserPointer(Window *GLFWwindow) unsafe.Pointer {
	r1, _, _ := proc_glfwGetWindowUserPointer.Call(uintptr(unsafe.Pointer(Window)))
	return unsafe.Pointer(r1)
}

func (g *Glfw) SetWindowPosCallback(Window *GLFWwindow, Callback GLFWwindowposfun) GLFWwindowposfun {
	r1, _, _ := proc_glfwSetWindowPosCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowposfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowSizeCallback(Window *GLFWwindow, Callback GLFWwindowsizefun) GLFWwindowsizefun {
	r1, _, _ := proc_glfwSetWindowSizeCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowsizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowCloseCallback(Window *GLFWwindow, Callback GLFWwindowclosefun) GLFWwindowclosefun {
	r1, _, _ := proc_glfwSetWindowCloseCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowclosefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowRefreshCallback(Window *GLFWwindow, Callback GLFWwindowrefreshfun) GLFWwindowrefreshfun {
	r1, _, _ := proc_glfwSetWindowRefreshCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowrefreshfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowFocusCallback(Window *GLFWwindow, Callback GLFWwindowfocusfun) GLFWwindowfocusfun {
	r1, _, _ := proc_glfwSetWindowFocusCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowfocusfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowIconifyCallback(Window *GLFWwindow, Callback GLFWwindowiconifyfun) GLFWwindowiconifyfun {
	r1, _, _ := proc_glfwSetWindowIconifyCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowiconifyfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowMaximizeCallback(Window *GLFWwindow, Callback GLFWwindowmaximizefun) GLFWwindowmaximizefun {
	r1, _, _ := proc_glfwSetWindowMaximizeCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowmaximizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetFramebufferSizeCallback(Window *GLFWwindow, Callback GLFWframebuffersizefun) GLFWframebuffersizefun {
	r1, _, _ := proc_glfwSetFramebufferSizeCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWframebuffersizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowContentScaleCallback(Window *GLFWwindow, Callback GLFWwindowcontentscalefun) GLFWwindowcontentscalefun {
	r1, _, _ := proc_glfwSetWindowContentScaleCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowcontentscalefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) PollEvents() {
	proc_glfwPollEvents.Call()
}

func (g *Glfw) WaitEvents() {
	proc_glfwWaitEvents.Call()
}

func (g *Glfw) WaitEventsTimeout(Timeout float64) {
	proc_glfwWaitEventsTimeout.Call(uintptr(Timeout))
}

func (g *Glfw) PostEmptyEvent() {
	proc_glfwPostEmptyEvent.Call()
}

func (g *Glfw) GetInputMode(Window *GLFWwindow, Mode int32) int32 {
	r1, _, _ := proc_glfwGetInputMode.Call(uintptr(unsafe.Pointer(Window)), uintptr(Mode))
	return int32(r1)
}

func (g *Glfw) SetInputMode(Window *GLFWwindow, Mode int32, Value int32) {
	proc_glfwSetInputMode.Call(uintptr(unsafe.Pointer(Window)), uintptr(Mode), uintptr(Value))
}

func (g *Glfw) RawMouseMotionSupported() int32 {
	r1, _, _ := proc_glfwRawMouseMotionSupported.Call()
	return int32(r1)
}

func (g *Glfw) GetKeyName(Key int32, Scancode int32) *int8 {
	r1, _, _ := proc_glfwGetKeyName.Call(uintptr(Key), uintptr(Scancode))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetKeyScancode(Key int32) int32 {
	r1, _, _ := proc_glfwGetKeyScancode.Call(uintptr(Key))
	return int32(r1)
}

func (g *Glfw) GetKey(Window *GLFWwindow, Key int32) int32 {
	r1, _, _ := proc_glfwGetKey.Call(uintptr(unsafe.Pointer(Window)), uintptr(Key))
	return int32(r1)
}

func (g *Glfw) GetMouseButton(Window *GLFWwindow, Button int32) int32 {
	r1, _, _ := proc_glfwGetMouseButton.Call(uintptr(unsafe.Pointer(Window)), uintptr(Button))
	return int32(r1)
}

func (g *Glfw) GetCursorPos(Window *GLFWwindow, Xpos *float64, Ypos *float64) {
	proc_glfwGetCursorPos.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) SetCursorPos(Window *GLFWwindow, Xpos float64, Ypos float64) {
	proc_glfwSetCursorPos.Call(uintptr(unsafe.Pointer(Window)), uintptr(Xpos), uintptr(Ypos))
}

func (g *Glfw) CreateCursor(Image *GLFWimage, Xhot int32, Yhot int32) *GLFWcursor {
	r1, _, _ := proc_glfwCreateCursor.Call(uintptr(unsafe.Pointer(Image)), uintptr(Xhot), uintptr(Yhot))
	return (*GLFWcursor)(unsafe.Pointer(r1))
}

func (g *Glfw) CreateStandardCursor(Shape int32) *GLFWcursor {
	r1, _, _ := proc_glfwCreateStandardCursor.Call(uintptr(Shape))
	return (*GLFWcursor)(unsafe.Pointer(r1))
}

func (g *Glfw) DestroyCursor(Cursor *GLFWcursor) {
	proc_glfwDestroyCursor.Call(uintptr(unsafe.Pointer(Cursor)))
}

func (g *Glfw) SetCursor(Window *GLFWwindow, Cursor *GLFWcursor) {
	proc_glfwSetCursor.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Cursor)))
}

func (g *Glfw) SetKeyCallback(Window *GLFWwindow, Callback GLFWkeyfun) GLFWkeyfun {
	r1, _, _ := proc_glfwSetKeyCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWkeyfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCharCallback(Window *GLFWwindow, Callback GLFWcharfun) GLFWcharfun {
	r1, _, _ := proc_glfwSetCharCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcharfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCharModsCallback(Window *GLFWwindow, Callback GLFWcharmodsfun) GLFWcharmodsfun {
	r1, _, _ := proc_glfwSetCharModsCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcharmodsfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetMouseButtonCallback(Window *GLFWwindow, Callback GLFWmousebuttonfun) GLFWmousebuttonfun {
	r1, _, _ := proc_glfwSetMouseButtonCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWmousebuttonfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCursorPosCallback(Window *GLFWwindow, Callback GLFWcursorposfun) GLFWcursorposfun {
	r1, _, _ := proc_glfwSetCursorPosCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcursorposfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCursorEnterCallback(Window *GLFWwindow, Callback GLFWcursorenterfun) GLFWcursorenterfun {
	r1, _, _ := proc_glfwSetCursorEnterCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcursorenterfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetScrollCallback(Window *GLFWwindow, Callback GLFWscrollfun) GLFWscrollfun {
	r1, _, _ := proc_glfwSetScrollCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWscrollfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetDropCallback(Window *GLFWwindow, Callback GLFWdropfun) GLFWdropfun {
	r1, _, _ := proc_glfwSetDropCallback.Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWdropfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) JoystickPresent(Jid int32) int32 {
	r1, _, _ := proc_glfwJoystickPresent.Call(uintptr(Jid))
	return int32(r1)
}

func (g *Glfw) GetJoystickAxes(Jid int32, Count *int32) *float32 {
	r1, _, _ := proc_glfwGetJoystickAxes.Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*float32)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickButtons(Jid int32, Count *int32) *uint8 {
	r1, _, _ := proc_glfwGetJoystickButtons.Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*uint8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickHats(Jid int32, Count *int32) *uint8 {
	r1, _, _ := proc_glfwGetJoystickHats.Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*uint8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickName(Jid int32) *int8 {
	r1, _, _ := proc_glfwGetJoystickName.Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickGUID(Jid int32) *int8 {
	r1, _, _ := proc_glfwGetJoystickGUID.Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetJoystickUserPointer(Jid int32, Pointer unsafe.Pointer) {
	proc_glfwSetJoystickUserPointer.Call(uintptr(Jid), uintptr(Pointer))
}

func (g *Glfw) GetJoystickUserPointer(Jid int32) unsafe.Pointer {
	r1, _, _ := proc_glfwGetJoystickUserPointer.Call(uintptr(Jid))
	return unsafe.Pointer(r1)
}

func (g *Glfw) JoystickIsGamepad(Jid int32) int32 {
	r1, _, _ := proc_glfwJoystickIsGamepad.Call(uintptr(Jid))
	return int32(r1)
}

func (g *Glfw) SetJoystickCallback(Callback GLFWjoystickfun) GLFWjoystickfun {
	r1, _, _ := proc_glfwSetJoystickCallback.Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWjoystickfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) UpdateGamepadMappings(String *int8) int32 {
	r1, _, _ := proc_glfwUpdateGamepadMappings.Call(uintptr(unsafe.Pointer(String)))
	return int32(r1)
}

func (g *Glfw) GetGamepadName(Jid int32) *int8 {
	r1, _, _ := proc_glfwGetGamepadName.Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetGamepadState(Jid int32, State *GLFWgamepadstate) int32 {
	r1, _, _ := proc_glfwGetGamepadState.Call(uintptr(Jid), uintptr(unsafe.Pointer(State)))
	return int32(r1)
}

func (g *Glfw) SetClipboardString(Window *GLFWwindow, String *int8) {
	proc_glfwSetClipboardString.Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(String)))
}

func (g *Glfw) GetClipboardString(Window *GLFWwindow) *int8 {
	r1, _, _ := proc_glfwGetClipboardString.Call(uintptr(unsafe.Pointer(Window)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetTime() float64 {
	r1, _, _ := proc_glfwGetTime.Call()
	return *(*float64)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetTime(Time float64) {
	proc_glfwSetTime.Call(uintptr(Time))
}

func (g *Glfw) GetTimerValue() uint64 {
	r1, _, _ := proc_glfwGetTimerValue.Call()
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetTimerFrequency() uint64 {
	r1, _, _ := proc_glfwGetTimerFrequency.Call()
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (g *Glfw) MakeContextCurrent(Window *GLFWwindow) {
	proc_glfwMakeContextCurrent.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) GetCurrentContext() *GLFWwindow {
	r1, _, _ := proc_glfwGetCurrentContext.Call()
	return (*GLFWwindow)(unsafe.Pointer(r1))
}

func (g *Glfw) SwapBuffers(Window *GLFWwindow) {
	proc_glfwSwapBuffers.Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) SwapInterval(Interval int32) {
	proc_glfwSwapInterval.Call(uintptr(Interval))
}

func (g *Glfw) ExtensionSupported(Extension *int8) int32 {
	r1, _, _ := proc_glfwExtensionSupported.Call(uintptr(unsafe.Pointer(Extension)))
	return int32(r1)
}

func (g *Glfw) GetProcAddress(Procname *int8) GLFWglproc {
	r1, _, _ := proc_glfwGetProcAddress.Call(uintptr(unsafe.Pointer(Procname)))
	return *(*GLFWglproc)(unsafe.Pointer(&r1))
}

func (g *Glfw) VulkanSupported() int32 {
	r1, _, _ := proc_glfwVulkanSupported.Call()
	return int32(r1)
}

func (g *Glfw) GetRequiredInstanceExtensions(Count *uint32) **int8 {
	r1, _, _ := proc_glfwGetRequiredInstanceExtensions.Call(uintptr(unsafe.Pointer(Count)))
	return (**int8)(unsafe.Pointer(r1))
}
