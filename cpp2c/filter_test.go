package cpp2c

import (
	"testing"
)

type mockFilter struct {
	allowClassFunc                func(className string) bool
	allowMethodFunc               func(className string, mm CppMethod) error
	allowFreeFunctionFunc         func(fn CppFreeFunction) error
	allowSignalFunc               func(mm CppMethod) bool
	allowInheritedParentFunc      func(className string) bool
	allowVirtualFunc              func(mm CppMethod) bool
	allowVirtualForClassFunc      func(className string) bool
	allowCtorFunc                 func(className string, mm CppMethod) bool
	allowTypeFunc                 func(p CppParameter, isReturnType bool) error
	insertTypedefsFunc            func()
	headerPlatformRestrictionFunc func(fullpath string) AllowedPlatformInfo
	getPureVirtualDefaultsFunc    func(className string) []PureVirtualDefault
	applyQuirksFunc               func(packageName, className string, mm *CppMethod)
}

func (m *mockFilter) AllowClass(className string) bool {
	if m.allowClassFunc != nil {
		return m.allowClassFunc(className)
	}
	return true
}

func (m *mockFilter) AllowMethod(className string, mm CppMethod) error {
	if m.allowMethodFunc != nil {
		return m.allowMethodFunc(className, mm)
	}
	return nil
}

func (m *mockFilter) AllowFreeFunction(fn CppFreeFunction) error {
	if m.allowFreeFunctionFunc != nil {
		return m.allowFreeFunctionFunc(fn)
	}
	return nil
}

func (m *mockFilter) AllowSignal(mm CppMethod) bool {
	if m.allowSignalFunc != nil {
		return m.allowSignalFunc(mm)
	}
	return true
}

func (m *mockFilter) AllowInheritedParent(className string) bool {
	if m.allowInheritedParentFunc != nil {
		return m.allowInheritedParentFunc(className)
	}
	return true
}

func (m *mockFilter) AllowVirtual(mm CppMethod) bool {
	if m.allowVirtualFunc != nil {
		return m.allowVirtualFunc(mm)
	}
	return true
}

func (m *mockFilter) AllowVirtualForClass(className string) bool {
	if m.allowVirtualForClassFunc != nil {
		return m.allowVirtualForClassFunc(className)
	}
	return true
}

func (m *mockFilter) AllowCtor(className string, mm CppMethod) bool {
	if m.allowCtorFunc != nil {
		return m.allowCtorFunc(className, mm)
	}
	return true
}

func (m *mockFilter) AllowType(p CppParameter, isReturnType bool) error {
	if m.allowTypeFunc != nil {
		return m.allowTypeFunc(p, isReturnType)
	}
	return nil
}

func (m *mockFilter) InsertTypedefs() {
	if m.insertTypedefsFunc != nil {
		m.insertTypedefsFunc()
	}
}

func (m *mockFilter) HeaderPlatformRestriction(fullpath string) AllowedPlatformInfo {
	if m.headerPlatformRestrictionFunc != nil {
		return m.headerPlatformRestrictionFunc(fullpath)
	}
	return nil
}

func (m *mockFilter) GetPureVirtualDefaults(className string) []PureVirtualDefault {
	if m.getPureVirtualDefaultsFunc != nil {
		return m.getPureVirtualDefaultsFunc(className)
	}
	return nil
}

func (m *mockFilter) ApplyQuirks(packageName, className string, mm *CppMethod) {
	if m.applyQuirksFunc != nil {
		m.applyQuirksFunc(packageName, className, mm)
	}
}

func (m *mockFilter) PreventStructDeclaration(className string) bool {
	return false
}

func (m *mockFilter) ShouldEmitForwardDeclaration(className string) bool {
	return false
}

func (m *mockFilter) ShouldEmitStructDefinition(className string) bool {
	return false
}

func (m *mockFilter) ShouldEmitEnumDefinition(enumName string) bool {
	return false
}

func TestSetActiveFilter(t *testing.T) {
	originalFilter := ActiveFilter
	defer func() { ActiveFilter = originalFilter }()

	mock := &mockFilter{
		allowClassFunc: func(className string) bool {
			return className == "TestClass"
		},
	}

	SetActiveFilter(mock)

	if !ActiveFilter.AllowClass("TestClass") {
		t.Error("Expected AllowClass to return true for TestClass")
	}
	if ActiveFilter.AllowClass("OtherClass") {
		t.Error("Expected AllowClass to return false for OtherClass")
	}
}

func TestDefaultFilter(t *testing.T) {
	filter := &DefaultFilter{}

	tests := []struct {
		name      string
		className string
		expected  bool
	}{
		{"primitive int", "int", true},
		{"primitive void", "void", true},
		{"primitive float", "float", true},
		{"primitive bool", "bool", true},
		{"primitive char", "char", true},
		{"size_t", "size_t", true},
		{"int8_t", "int8_t", true},
		{"uint64_t", "uint64_t", true},
		{"std::string", "std::string", false},
		{"std::vector", "std::vector", false},
		{"std::map", "std::map", false},
		{"QString", "QString", true},
		{"QObject", "QObject", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowClass(tt.className)
			if result != tt.expected {
				t.Errorf("AllowClass(%q) = %v, want %v", tt.className, result, tt.expected)
			}
		})
	}
}

func TestDefaultFilterAllowFreeFunction(t *testing.T) {
	filter := &DefaultFilter{}

	tests := []struct {
		name    string
		fn      CppFreeFunction
		wantErr bool
	}{
		{
			name: "normal function",
			fn: CppFreeFunction{
				FunctionName: "myFunction",
				Parameters:   []CppParameter{},
			},
			wantErr: false,
		},
		{
			name: "variadic function",
			fn: CppFreeFunction{
				FunctionName: "printf",
				Parameters:   []CppParameter{{ParameterType: "..."}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := filter.AllowFreeFunction(tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowFreeFunction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultFilterAllowMethod(t *testing.T) {
	filter := &DefaultFilter{}

	tests := []struct {
		name      string
		className string
		method    CppMethod
		wantErr   bool
	}{
		{
			name:      "normal method",
			className: "TestClass",
			method: CppMethod{
				MethodName: "doSomething",
				ReturnType: CppParameter{ParameterType: "void"},
			},
			wantErr: false,
		},
		{
			name:      "operator method",
			className: "TestClass",
			method: CppMethod{
				MethodName: "operator=",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := filter.AllowMethod(tt.className, tt.method)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowMethod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultFilterAllowSignal(t *testing.T) {
	filter := &DefaultFilter{}

	tests := []struct {
		name     string
		method   CppMethod
		expected bool
	}{
		{
			name: "normal signal",
			method: CppMethod{
				MethodName: "clicked",
				IsSignal:   true,
			},
			expected: true,
		},
		{
			name: "private signal",
			method: CppMethod{
				MethodName: "privateSignal",
				IsSignal:   true,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.AllowSignal(tt.method)
			if result != tt.expected {
				t.Errorf("AllowSignal() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDefaultFilterAllowType(t *testing.T) {
	filter := &DefaultFilter{}

	tests := []struct {
		name         string
		param        CppParameter
		isReturnType bool
		wantErr      bool
	}{
		{
			name:         "int type",
			param:        CppParameter{ParameterType: "int"},
			isReturnType: false,
			wantErr:      false,
		},
		{
			name:         "void pointer",
			param:        CppParameter{ParameterType: "void", Pointer: true},
			isReturnType: false,
			wantErr:      true,
		},
		{
			name:         "function pointer",
			param:        CppParameter{ParameterType: "void(*)()"},
			isReturnType: false,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := filter.AllowType(tt.param, tt.isReturnType)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllowType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
