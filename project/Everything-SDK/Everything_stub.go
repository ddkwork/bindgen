package everything

// Source: Everything_stub.h:21 -> _FILETIME
type FILETIME struct {
	DwLowDateTime uint32
	DwHighDateTime uint32
}

// Source: Everything_stub.h:0 -> _LARGE_INTEGER
type LARGE_INTEGER struct {
	Data int64
}

// Source: Everything_stub.h -> Macro constants
const (
	IncWindows uint32 = 1
	DWORD uint32 = unsignedlong
	UINT uint32 = unsignedint
	UintPtr uint32 = longlong
	IntPtr uint32 = longlong
	WPARAM uint32 = longlong
	LPARAM uint32 = longlong
	LPCSTR uint32 = constchar*
	LPCWSTR uint32 = constwchar_t*
)

