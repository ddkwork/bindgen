package cpp2c

type Filter interface {
	AllowClass(className string) bool                              // 过滤类名前缀
	AllowMethod(className string, mm CppMethod) error              // 过滤类方法
	AllowFreeFunction(fn CppFreeFunction) error                    // 过滤自由函数
	AllowSignal(mm CppMethod) bool                                 // 过滤信号
	AllowInheritedParent(className string) bool                    // 过滤继承父类
	AllowVirtual(mm CppMethod) bool                                // 过滤虚函数
	AllowVirtualForClass(className string) bool                    // 过滤类的虚函数支持
	AllowCtor(className string, mm CppMethod) bool                 // 过滤构造函数
	AllowType(p CppParameter, isReturnType bool) error             // 过滤参数/返回类型
	InsertTypedefs()                                               // 注入已知typedef
	HeaderPlatformRestriction(fullpath string) AllowedPlatformInfo // 平台限制
	GetPureVirtualDefaults(className string) []PureVirtualDefault  // 纯虚函数默认实现
	ApplyQuirks(packageName, className string, mm *CppMethod)      // 应用特殊处理
	PreventStructDeclaration(className string) bool                // 防止结构体声明
	ShouldEmitForwardDeclaration(className string) bool            // 是否生成前向声明（Qt内置类型需要）
	ShouldEmitStructDefinition(className string) bool              // 是否从AST自动生成结构体字段定义
	ShouldEmitEnumDefinition(enumName string) bool                 // 是否从AST自动生成枚举定义
}

type DefaultFilter struct{}

func (d *DefaultFilter) AllowClass(className string) bool {
	return DefaultAllowClass(className)
}

func (d *DefaultFilter) AllowMethod(className string, mm CppMethod) error {
	return DefaultAllowMethod(className, mm)
}

func (d *DefaultFilter) AllowFreeFunction(fn CppFreeFunction) error {
	return DefaultAllowFreeFunction(fn)
}

func (d *DefaultFilter) AllowSignal(mm CppMethod) bool {
	return DefaultAllowSignal(mm)
}

func (d *DefaultFilter) AllowInheritedParent(className string) bool {
	return true
}

func (d *DefaultFilter) AllowVirtual(mm CppMethod) bool {
	return true
}

func (d *DefaultFilter) AllowVirtualForClass(className string) bool {
	return true
}

func (d *DefaultFilter) AllowCtor(className string, mm CppMethod) bool {
	return true
}

func (d *DefaultFilter) AllowType(p CppParameter, isReturnType bool) error {
	return nil
}

func (d *DefaultFilter) InsertTypedefs() {}

func (d *DefaultFilter) HeaderPlatformRestriction(fullpath string) AllowedPlatformInfo {
	return nil
}

func (d *DefaultFilter) GetPureVirtualDefaults(className string) []PureVirtualDefault {
	return nil
}

func (d *DefaultFilter) ApplyQuirks(packageName, className string, mm *CppMethod) {}

func (d *DefaultFilter) PreventStructDeclaration(className string) bool {
	return false
}

func (d *DefaultFilter) ShouldEmitForwardDeclaration(className string) bool {
	return false
}

func (d *DefaultFilter) ShouldEmitStructDefinition(className string) bool {
	return false
}

func (d *DefaultFilter) ShouldEmitEnumDefinition(enumName string) bool {
	return false
}

var ActiveFilter Filter = &DefaultFilter{}

func SetActiveFilter(f Filter) {
	ActiveFilter = f
}

func InsertTypedefs() {
	ActiveFilter.InsertTypedefs()
}
