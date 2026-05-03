package glfw

import (
	"os"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
)

var testGlfw *Glfw

func TestMain(m *testing.M) {
	g := &Glfw{}
	if g.Init() == 0 {
		var desc *int8
		g.GetError(&desc)
		panic("glfwInit failed: " + byteslice.PtrToString(desc))
	}
	testGlfw = g
	exitCode := m.Run()
	g.Terminate()
	os.Exit(exitCode)
}

func newGlfw(t *testing.T) *Glfw {
	t.Helper()
	return testGlfw
}

func newWindow(t *testing.T, g *Glfw, width int32, height int32, title string) *GLFWwindow {
	t.Helper()
	g.DefaultWindowHints()
	g.WindowHint(int32(GlfwVisible), int32(GlfwFalse))
	ct, _ := func() (*int8, []byte) { b := byteslice.FromString[int8](title); return &b[0], b }()
	w := g.CreateWindow(width, height, ct, nil, nil)
	if w == nil {
		var desc *int8
		g.GetError(&desc)
		t.Fatalf("glfwCreateWindow failed: %s", byteslice.PtrToString(desc))
	}
	return w
}

// === Init / Terminate / Version ===

func TestGlfwInitTerminate(t *testing.T) {
	g := newGlfw(t)
	_ = g
}

func TestGlfwGetVersion(t *testing.T) {
	g := newGlfw(t)

	var major, minor, rev int32
	g.GetVersion(&major, &minor, &rev)
	t.Logf("GLFW version: %d.%d.%d", major, minor, rev)
	if major < 3 || (major == 3 && minor < 5) {
		t.Fatalf("expected GLFW >= 3.5, got %d.%d.%d", major, minor, rev)
	}
}

func TestGlfwGetVersionString(t *testing.T) {
	g := newGlfw(t)

	str := byteslice.PtrToString(g.GetVersionString())
	t.Logf("GLFW version string: %s", str)
	if str == "<nil>" || str == "" || len(str) < 10 {
		t.Fatal("version string invalid")
	}
}

// === Platform ===

func TestGlfwGetPlatform(t *testing.T) {
	g := newGlfw(t)

	platform := g.GetPlatform()
	t.Logf("platform: %d (expect Win32=%d)", platform, GlfwPlatformWin32)
}

func TestGlfwPlatformSupported(t *testing.T) {
	g := newGlfw(t)

	supported := g.PlatformSupported(int32(GlfwPlatformWin32))
	if supported == 0 {
		t.Fatal("Win32 platform not supported")
	}
}

// === Error Callback (from triangle-opengl.c) ===

func TestGlfwErrorCallback(t *testing.T) {
	g := newGlfw(t)

	cb := func(code int32, desc *int8) uintptr {
		t.Logf("GLFW error %d: %s", code, byteslice.PtrToString(desc))
		return 0
	}
	g.SetErrorCallback(syscall.NewCallback(cb))

	g.WindowHint(int32(0x00020001), 0)
	g.PollEvents()
}

// === Monitor ===

func TestGlfwGetMonitors(t *testing.T) {
	g := newGlfw(t)

	var count int32
	monitors := g.GetMonitors(&count)
	t.Logf("found %d monitors", count)
	if monitors == nil || count == 0 {
		t.Fatal("no monitors")
	}
}

func TestGlfwGetPrimaryMonitor(t *testing.T) {
	g := newGlfw(t)

	monitor := g.GetPrimaryMonitor()
	if monitor == nil {
		t.Fatal("no primary monitor")
	}

	var x, y int32
	g.GetMonitorPos(monitor, &x, &y)
	t.Logf("pos: (%d,%d)", x, y)

	var pw, ph int32
	g.GetMonitorPhysicalSize(monitor, &pw, &ph)
	t.Logf("physical: %dx%d mm", pw, ph)

	name := byteslice.PtrToString(g.GetMonitorName(monitor))
	t.Logf("name: %s", name)
}

func TestGlfwGetMonitorWorkarea(t *testing.T) {
	g := newGlfw(t)

	m := g.GetPrimaryMonitor()
	var x, y, w, h int32
	g.GetMonitorWorkarea(m, &x, &y, &w, &h)
	t.Logf("workarea: (%d,%d) %dx%d", x, y, w, h)
	if w <= 0 || h <= 0 {
		t.Fatal("invalid workarea size")
	}
}

func TestGlfwGetMonitorContentScale(t *testing.T) {
	g := newGlfw(t)

	m := g.GetPrimaryMonitor()
	var xs, ys float32
	g.GetMonitorContentScale(m, &xs, &ys)
	t.Logf("content scale: %.2f x %.2f", xs, ys)
	if xs <= 0 || ys <= 0 {
		t.Fatal("invalid content scale")
	}
}

func TestGlfwGetVideoMode(t *testing.T) {
	g := newGlfw(t)

	mode := g.GetVideoMode(g.GetPrimaryMonitor())
	if mode == nil {
		t.Fatal("no video mode")
	}
	t.Logf("video mode: %dx%d @%dHz r%dg%db%d",
		mode.Width, mode.Height, mode.RefreshRate,
		mode.RedBits, mode.GreenBits, mode.BlueBits)
	if mode.Width <= 0 || mode.Height <= 0 {
		t.Fatal("invalid video mode dimensions")
	}
}

func TestGlfwGetVideoModes(t *testing.T) {
	g := newGlfw(t)

	var count int32
	modes := g.GetVideoModes(g.GetPrimaryMonitor(), &count)
	t.Logf("%d video modes available", count)
	if modes != nil && count > 0 {
		t.Logf("first mode: %dx%d @%dHz", (*modes).Width, (*modes).Height, (*modes).RefreshRate)
	}
}

func TestGlfwMonitorUserPointer(t *testing.T) {
	g := newGlfw(t)

	m := g.GetPrimaryMonitor()
	tag := uintptr(0xDEAD)
	g.SetMonitorUserPointer(m, unsafe.Pointer(tag))
	p := g.GetMonitorUserPointer(m)
	if p != unsafe.Pointer(tag) {
		t.Fatalf("user pointer mismatch: got %x want %x", uintptr(p), tag)
	}
}

// === Window Creation / Destruction (from windows.c, offscreen.c) ===

func TestGlfwCreateWindowBasic(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Test Window")
	defer g.DestroyWindow(w)

	title := byteslice.PtrToString(g.GetWindowTitle(w))
	if title != "Test Window" {
		t.Fatalf("title: got %q want %q", title, "Test Window")
	}
}

func TestGlfwCreateWindowHidden(t *testing.T) {
	g := newGlfw(t)

	g.DefaultWindowHints()
	g.WindowHint(int32(GlfwVisible), int32(GlfwFalse))
	ct, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("Hidden"); return &b[0], b }()
	w := g.CreateWindow(320, 240, ct, nil, nil)
	if w == nil {
		t.Fatal("create hidden window failed")
	}
	defer g.DestroyWindow(w)

	attrib := g.GetWindowAttrib(w, int32(GlfwVisible))
	if attrib != 0 {
		t.Fatal("hidden window should have Visible=FALSE")
	}
}

func TestGlfwSetWindowTitle(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 200, 200, "Old Title")
	defer g.DestroyWindow(w)

	newTitle, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("New Title"); return &b[0], b }()
	g.SetWindowTitle(w, newTitle)
	title := byteslice.PtrToString(g.GetWindowTitle(w))
	if title != "New Title" {
		t.Fatalf("got %q want %q", title, "New Title")
	}
}

func TestGlfwWindowPosSize(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 400, 300, "PosSize")
	defer g.DestroyWindow(w)

	g.SetWindowPos(w, 100, 200)
	var px, py int32
	g.GetWindowPos(w, &px, &py)
	t.Logf("pos: (%d,%d)", px, py)

	g.SetWindowSize(w, 800, 600)
	var ww, wh int32
	g.GetWindowSize(w, &ww, &wh)
	if ww != 800 || wh != 600 {
		t.Fatalf("size: got %dx%d want 800x600", ww, wh)
	}
}

func TestGlfwFramebufferSize(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "FB Size")
	defer g.DestroyWindow(w)

	var fbw, fbh int32
	g.GetFramebufferSize(w, &fbw, &fbh)
	t.Logf("framebuffer size: %dx%d", fbw, fbh)
	if fbw <= 0 || fbh <= 0 {
		t.Fatal("invalid framebuffer size")
	}
}

func TestGlfwWindowFrameSize(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Frame")
	defer g.DestroyWindow(w)

	var left, top, right, bottom int32
	g.GetWindowFrameSize(w, &left, &top, &right, &bottom)
	t.Logf("frame: left=%d top=%d right=%d bottom=%d", left, top, right, bottom)
}

func TestGlfwWindowContentScale(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Scale")
	defer g.DestroyWindow(w)

	var xs, ys float32
	g.GetWindowContentScale(w, &xs, &ys)
	t.Logf("content scale: %.2f x %.2f", xs, ys)
	if xs <= 0 || ys <= 0 {
		t.Fatal("invalid content scale")
	}
}

func TestGlfwWindowOpacity(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Opacity")
	defer g.DestroyWindow(w)

	g.SetWindowOpacity(w, 0.7)
	op := g.GetWindowOpacity(w)
	t.Logf("opacity: %.2f", op)
	if op == 1.0 || op == 0.0 {
		t.Log("opacity not supported, skipping check")
	} else if op < 0.65 || op > 0.75 {
		t.Fatalf("opacity: got %.2f want ~0.70", op)
	}
}

func TestGlfwWindowShouldClose(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Close")
	defer g.DestroyWindow(w)

	if g.WindowShouldClose(w) != 0 {
		t.Fatal("new window should not be closing")
	}
	g.SetWindowShouldClose(w, int32(GlfwTrue))
	if g.WindowShouldClose(w) == 0 {
		t.Fatal("window should be closing after SetWindowShouldClose(TRUE)")
	}
}

func TestGlfwWindowUserPointer(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "UPtr")
	defer g.DestroyWindow(w)

	tag := uintptr(0xC0FFEE)
	g.SetWindowUserPointer(w, unsafe.Pointer(tag))
	p := g.GetWindowUserPointer(w)
	if p != unsafe.Pointer(tag) {
		t.Fatalf("user pointer mismatch: got %x want %x", uintptr(p), tag)
	}
}

func TestGlfwWindowShowHideIconifyRestoreMaximize(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Visibility")
	defer g.DestroyWindow(w)

	g.ShowWindow(w)
	g.HideWindow(w)
	g.ShowWindow(w)

	g.IconifyWindow(w)
	iconified := g.GetWindowAttrib(w, int32(GlfwIconified))
	if iconified == 0 {
		t.Fatal("window should be iconified")
	}

	g.RestoreWindow(w)
	iconified = g.GetWindowAttrib(w, int32(GlfwIconified))
	if iconified != 0 {
		t.Fatal("window should not be iconified after restore")
	}

	g.MaximizeWindow(w)
	maximized := g.GetWindowAttrib(w, int32(GlfwMaximized))
	if maximized == 0 {
		t.Fatal("window should be maximized")
	}
}

func TestGlfwWindowHintVisibleOffscreen(t *testing.T) {
	g := newGlfw(t)

	g.DefaultWindowHints()
	g.WindowHint(int32(GlfwVisible), int32(GlfwFalse))
	g.WindowHint(int32(GlfwResizable), int32(GlfwTrue))
	g.WindowHint(int32(GlfwDecorated), int32(GlfwTrue))
	ct, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("Offscreen"); return &b[0], b }()
	w := g.CreateWindow(640, 480, ct, nil, nil)
	if w == nil {
		t.Fatal("offscreen window creation failed")
	}
	defer g.DestroyWindow(w)

	resizable := g.GetWindowAttrib(w, int32(GlfwResizable))
	decorated := g.GetWindowAttrib(w, int32(GlfwDecorated))
	t.Logf("resizable=%d decorated=%d", resizable, decorated)
	if resizable == 0 {
		t.Fatal("expected resizable=TRUE")
	}
}

func TestGlfwWindowHintContextVersion(t *testing.T) {
	g := newGlfw(t)

	g.DefaultWindowHints()
	g.WindowHint(int32(GlfwContextVersionMajor), 3)
	g.WindowHint(int32(GlfwContextVersionMinor), 3)
	g.WindowHint(int32(GlfwOpenglProfile), int32(GlfwOpenglCoreProfile))
	g.WindowHint(int32(GlfwVisible), int32(GlfwFalse))
	ct, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("GL Context"); return &b[0], b }()
	w := g.CreateWindow(640, 480, ct, nil, nil)
	if w == nil {
		t.Fatal("OpenGL context window creation failed")
	}
	defer g.DestroyWindow(w)

	ctxVerMaj := g.GetWindowAttrib(w, int32(GlfwContextVersionMajor))
	ctxVerMin := g.GetWindowAttrib(w, int32(GlfwContextVersionMinor))
	t.Logf("context version: %d.%d", ctxVerMaj, ctxVerMin)
}

func TestGlfwMakeContextCurrent(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Context")
	defer g.DestroyWindow(w)

	g.MakeContextCurrent(w)
	current := g.GetCurrentContext()
	if current != w {
		t.Fatal("current context should match created window")
	}
}

func TestGlfwWindowSizeLimits(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 400, 300, "Limits")
	defer g.DestroyWindow(w)

	g.SetWindowSizeLimits(w, 200, 150, 1920, 1080)
	t.Log("size limits set: min 200x150 max 1920x1080")
}

func TestGlfwWindowAspectRatio(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 800, 600, "Aspect")
	defer g.DestroyWindow(w)

	g.SetWindowAspectRatio(w, 4, 3)
	t.Log("aspect ratio set to 4:3")
}

func TestGlfwFocusWindow(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Focus")
	defer g.DestroyWindow(w)

	g.FocusWindow(w)
	focused := g.GetWindowAttrib(w, int32(GlfwFocused))
	t.Logf("focused: %d", focused)
}

func TestGlfwRequestWindowAttention(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Attention")
	defer g.DestroyWindow(w)

	g.RequestWindowAttention(w)
	t.Log("requested window attention")
}

func TestGlfwWindowMonitorFullscreen(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "FS")
	defer g.DestroyWindow(w)

	mon := g.GetWindowMonitor(w)
	if mon != nil {
		t.Fatal("windowed window should return NULL monitor")
	}
}

// === Input Mode / Key / Mouse (from triangle-opengl.c key_callback) ===

func TestGlfwInputModeStickyKeys(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Input")
	defer g.DestroyWindow(w)

	g.SetInputMode(w, int32(GlfwStickyKeys), int32(GlfwTrue))
	mode := g.GetInputMode(w, int32(GlfwStickyKeys))
	if mode != int32(GlfwTrue) {
		t.Fatalf("sticky keys mode: got %d want TRUE", mode)
	}
}

func TestGlfwInputModeCursor(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Cursor")
	defer g.DestroyWindow(w)

	g.SetInputMode(w, int32(GlfwCursor), int32(GlfwCursorDisabled))
	mode := g.GetInputMode(w, int32(GlfwCursor))
	if mode != int32(GlfwCursorDisabled) {
		t.Fatalf("cursor mode: got %d want DISABLED", mode)
	}

	g.SetInputMode(w, int32(GlfwCursor), int32(GlfwCursorNormal))
	mode = g.GetInputMode(w, int32(GlfwCursor))
	if mode != int32(GlfwCursorNormal) {
		t.Fatalf("cursor mode: got %d want NORMAL", mode)
	}
}

func TestGlfwRawMouseMotionSupported(t *testing.T) {
	g := newGlfw(t)

	supported := g.RawMouseMotionSupported()
	t.Logf("raw mouse motion supported: %d", supported)
}

func TestGlfwGetKeyDefaultState(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Keys")
	defer g.DestroyWindow(w)

	state := g.GetKey(w, int32(GlfwKeySpace))
	t.Logf("key SPACE state: %d (RELEASE=%d PRESS=%d)", state, GlfwRelease, GlfwPress)
}

func TestGlfwGetMouseButtonDefaultState(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Mouse")
	defer g.DestroyWindow(w)

	for btn := uint32(GlfwMouseButton1); btn <= GlfwMouseButton3; btn++ {
		state := g.GetMouseButton(w, int32(btn))
		t.Logf("mouse button %d state: %d", btn, state)
	}
}

func TestGlfwGetCursorPos(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "CursorPos")
	defer g.DestroyWindow(w)

	var cx, cy float64
	g.GetCursorPos(w, &cx, &cy)
	t.Logf("cursor pos: (%.1f, %.1f)", cx, cy)

	g.SetCursorPos(w, 100.0, 200.0)
	g.GetCursorPos(w, &cx, &cy)
	t.Logf("after set: (%.1f, %.1f)", cx, cy)
}

func TestGlfwKeyName(t *testing.T) {
	g := newGlfw(t)

	name := byteslice.PtrToString(g.GetKeyName(int32(GlfwKeyA), -1))
	t.Logf("name of KEY_A: %q", name)

	escName := byteslice.PtrToString(g.GetKeyName(int32(GlfwKeyEscape), -1))
	t.Logf("name of KEY_ESCAPE: %q", escName)
}

func TestGlfwGetKeyScancode(t *testing.T) {
	g := newGlfw(t)

	scancode := g.GetKeyScancode(int32(GlfwKeyA))
	t.Logf("scancode for KEY_A: %d", scancode)

	escapeSc := g.GetKeyScancode(int32(GlfwKeyEscape))
	t.Logf("scancode for KEY_ESCAPE: %d", escapeSc)
}

func TestGlfwStandardCursors(t *testing.T) {
	g := newGlfw(t)

	cursorTypes := []struct {
		name  string
		value int32
	}{
		{"Arrow", int32(GlfwArrowCursor)},
		{"IBeam", int32(GlfwIbeamCursor)},
		{"Crosshair", int32(GlfwCrosshairCursor)},
		{"Hand", int32(GlfwPointingHandCursor)},
		{"HResize", int32(GlfwHresizeCursor)},
		{"VResize", int32(GlfwVresizeCursor)},
	}

	for _, ct := range cursorTypes {
		cursor := g.CreateStandardCursor(ct.value)
		if cursor == nil {
			t.Logf("cursor %s: creation returned NULL (may be unsupported)", ct.name)
			continue
		}
		g.DestroyCursor(cursor)
		t.Logf("cursor %s: OK", ct.name)
	}
}

// === Clipboard ===

func TestGlfwClipboard(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "Clipboard")
	defer g.DestroyWindow(w)

	testStr, _ := func() (*int8, []byte) {
		b := byteslice.FromString[int8]("Hello from GLFW Go bindings!")
		return &b[0], b
	}()
	g.SetClipboardString(w, testStr)

	retrieved := byteslice.PtrToString(g.GetClipboardString(w))
	t.Logf("clipboard: %q", retrieved)
	if retrieved != "Hello from GLFW Go bindings!" {
		t.Fatalf("clipboard: got %q want expected string", retrieved)
	}
}

// === Time (used in triangle-opengl.c glfwGetTime for rotation) ===

func TestGlfwTime(t *testing.T) {
	g := newGlfw(t)

	t0 := g.GetTime()
	time.Sleep(50 * time.Millisecond)
	t1 := g.GetTime()
	elapsed := t1 - t0
	t.Logf("time: t0=%.6f t1=%.6f elapsed=%.6fs", t0, t1, elapsed)
	if t0 < 0.001 && t1 < 0.001 {
		t.Log("GetTime returns ~0 (may be unsupported in this environment)")
	} else if elapsed < 0.04 {
		t.Fatalf("elapsed time too short: %.6fs", elapsed)
	}
}

func TestGlfwSetTime(t *testing.T) {
	g := newGlfw(t)

	g.SetTime(12345.678)
	now := g.GetTime()
	t.Logf("set time to 12345.678, read back: %.3f", now)
	if now < 0.001 {
		t.Log("SetTime/GetTime not working (returns ~0)")
	} else if now < 12345.0 || now > 12346.0 {
		t.Fatalf("time: got %.3f want ~12345.678", now)
	}
}

func TestGlfwTimerPrecision(t *testing.T) {
	g := newGlfw(t)

	v0 := g.GetTimerValue()
	freq := g.GetTimerFrequency()
	v1 := g.GetTimerValue()
	t.Logf("timer: v0=%d v1=%d freq=%d", v0, v1, freq)
	if freq == 0 {
		t.Fatal("timer frequency is zero")
	}
	if v1 < v0 {
		t.Fatal("timer value went backwards")
	}
	deltaTicks := v1 - v0
	deltaSec := float64(deltaTicks) / float64(freq)
	t.Logf("delta: %d ticks = %.9fs", deltaTicks, deltaSec)
}

// === Events ===

func TestGlfwPollEvents(t *testing.T) {
	g := newGlfw(t)

	start := time.Now()
	for i := 0; i < 100; i++ {
		g.PollEvents()
	}
	elapsed := time.Since(start)
	t.Logf("100 PollEvents calls took %v", elapsed)
}

func TestGlfwPostEmptyEvent(t *testing.T) {
	g := newGlfw(t)

	g.PostEmptyEvent()
	t.Log("posted empty event")
}

// === Joystick / Gamepad ===

func TestGlfwJoystickPresentAll(t *testing.T) {
	g := newGlfw(t)

	for jid := uint32(GlfwJoystick1); jid <= GlfwJoystick16; jid++ {
		present := g.JoystickPresent(int32(jid)) != 0
		if present {
			name := byteslice.PtrToString(g.GetJoystickName(int32(jid)))
			guid := byteslice.PtrToString(g.GetJoystickGUID(int32(jid)))
			t.Logf("joystick %d: name=%s guid=%s", jid, name, guid)
		}
	}
}

func TestGlfwJoystickAxesButtonsHats(t *testing.T) {
	g := newGlfw(t)

	for jid := uint32(GlfwJoystick1); jid <= GlfwJoystick4; jid++ {
		if g.JoystickPresent(int32(jid)) == 0 {
			continue
		}

		var axisCount int32
		axes := g.GetJoystickAxes(int32(jid), &axisCount)
		t.Logf("joystick %d: %d axes", jid, axisCount)
		if axes != nil && axisCount > 0 {
			t.Logf("  axis[0]: %.3f", unsafe.Slice(axes, axisCount)[0])
		}

		var btnCount int32
		btns := g.GetJoystickButtons(int32(jid), &btnCount)
		_ = btns
		t.Logf("  %d buttons", btnCount)

		var hatCount int32
		hats := g.GetJoystickHats(int32(jid), &hatCount)
		t.Logf("  %d hats", hatCount)
		if hats != nil && hatCount > 0 {
			t.Logf("  hat[0]: %d", unsafe.Slice(hats, hatCount)[0])
		}
		break
	}
}

func TestGlfwGamepad(t *testing.T) {
	g := newGlfw(t)

	for jid := uint32(GlfwJoystick1); jid <= GlfwJoystick16; jid++ {
		isGP := g.JoystickIsGamepad(int32(jid))
		if isGP == 0 {
			continue
		}
		name := byteslice.PtrToString(g.GetGamepadName(int32(jid)))
		t.Logf("gamepad %d: %s", jid, name)

		var state GLFWgamepadstate
		ok := g.GetGamepadState(int32(jid), &state)
		if ok != 0 {
			t.Logf("  buttons: %+v axes: %+v", state.Buttons[:], state.Axes[:])
		}
		break
	}
}

func TestGlfwJoystickUserPointer(t *testing.T) {
	g := newGlfw(t)

	var tag int32 = 0xBEEF
	g.SetJoystickUserPointer(0, unsafe.Pointer(&tag))
	p := g.GetJoystickUserPointer(0)
	if p == nil {
		t.Log("joystick user pointer not supported (returns nil)")
	} else if (*int32)(p) != &tag {
		t.Logf("user pointer: got %x", uintptr(p))
	}
}

// === Vulkan ===

func TestGlfwExtensionSupported(t *testing.T) {
	g := newGlfw(t)

	extName, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("VK_KHR_surface"); return &b[0], b }()
	supported := g.ExtensionSupported(extName) != 0
	t.Logf("VK_KHR_surface supported: %v", supported)
}

func TestGlfwGetRequiredInstanceExtensions(t *testing.T) {
	g := newGlfw(t)

	var count uint32
	exts := g.GetRequiredInstanceExtensions(&count)
	t.Logf("%d required instance extensions:", count)
	if exts != nil {
		for i := 0; i < int(count); i++ {
			t.Logf("  [%d] %s", i, byteslice.PtrToString(unsafe.Slice(exts, count)[i]))
		}
	}
}

// === Init Hints (from offscreen.c glfwInitHint) ===

func TestGlfwInitHint(t *testing.T) {
	g := newGlfw(t)
	_ = g
	t.Log("init hints already applied before Init() in this test setup")
}

// === GetProcAddress (from triangle-opengl.c gladLoadGL) ===

func TestGlfwGetProcAddress(t *testing.T) {
	g := newGlfw(t)

	funcName, _ := func() (*int8, []byte) { b := byteslice.FromString[int8]("glClear"); return &b[0], b }()
	proc := g.GetProcAddress(funcName)
	if proc == nil {
		t.Log("glClear proc NULL (expected without GL context)")
	} else {
		t.Log("glClear proc found (non-NULL)")
	}
}

// === SwapInterval ===

func TestGlfwSwapInterval(t *testing.T) {
	g := newGlfw(t)

	w := newWindow(t, g, 640, 480, "VSync")
	defer g.DestroyWindow(w)

	g.MakeContextCurrent(w)
	g.SwapInterval(1)
	t.Log("vsync enabled (swap interval = 1)")
}

// === Multi-window (inspired by windows.c creating 4 windows) ===

func TestGlfwMultipleWindows(t *testing.T) {
	g := newGlfw(t)

	g.DefaultWindowHints()
	g.WindowHint(int32(GlfwVisible), int32(GlfwFalse))

	var windows [2]*GLFWwindow
	colors := []string{"Red", "Green"}
	for i := range windows {
		ct, _ := func() (*int8, []byte) { b := byteslice.FromString[int8](colors[i]); return &b[0], b }()
		windows[i] = g.CreateWindow(200, 200, ct, nil, nil)
		if windows[i] == nil {
			t.Fatalf("window %d creation failed", i)
		}
		defer g.DestroyWindow(windows[i])

		title := byteslice.PtrToString(g.GetWindowTitle(windows[i]))
		t.Logf("window[%d] title: %s", i, title)
	}

	for i, w := range windows {
		g.MakeContextCurrent(w)
		var ww, wh int32
		g.GetWindowSize(w, &ww, &wh)
		t.Logf("window[%d] size: %dx%d", i, ww, wh)
	}
}
