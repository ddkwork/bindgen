package cpp2c

type Converter interface {
	ClearAstFilteredCache(cacheDir string) error  //清理旧的ast过滤文件
	FilterAst(astFile, filteredFile string) error //执行ast过滤
	EmitCabi(sourceFile, outputDir string) error  //输出cabi
}

type ConverterImpl struct {
	Filter Filter
}

func NewConverter(filter Filter) Converter {
	return &ConverterImpl{Filter: filter}
}

func (c *ConverterImpl) ClearAstFilteredCache(cacheDir string) error {
	ClearFilteredCache(cacheDir)
	return nil
}

func (c *ConverterImpl) FilterAst(astFile, filteredFile string) error {
	return nil
}

func (c *ConverterImpl) EmitCabi(sourceFile, outputDir string) error {
	return nil
}
