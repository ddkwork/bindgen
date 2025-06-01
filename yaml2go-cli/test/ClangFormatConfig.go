package main

// RawStringFormats
type RawStringFormats struct {
	Language           string   `yaml:"Language"`
	Delimiters         []string `yaml:"Delimiters"`
	CanonicalDelimiter string   `yaml:"CanonicalDelimiter"`
	BasedOnStyle       string   `yaml:"BasedOnStyle"`
	EnclosingFunctions []string `yaml:"EnclosingFunctions"`
}

// AlignConsecutiveBitFields
type AlignConsecutiveBitFields struct {
	AcrossComments            bool `yaml:"AcrossComments"`
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
}

// BraceWrapping
type BraceWrapping struct {
	AfterClass            bool   `yaml:"AfterClass"`
	AfterExternBlock      bool   `yaml:"AfterExternBlock"`
	SplitEmptyFunction    bool   `yaml:"SplitEmptyFunction"`
	SplitEmptyNamespace   bool   `yaml:"SplitEmptyNamespace"`
	AfterCaseLabel        bool   `yaml:"AfterCaseLabel"`
	AfterControlStatement string `yaml:"AfterControlStatement"`
	AfterNamespace        bool   `yaml:"AfterNamespace"`
	AfterUnion            bool   `yaml:"AfterUnion"`
	SplitEmptyRecord      bool   `yaml:"SplitEmptyRecord"`
	AfterEnum             bool   `yaml:"AfterEnum"`
	AfterObjCDeclaration  bool   `yaml:"AfterObjCDeclaration"`
	AfterStruct           bool   `yaml:"AfterStruct"`
	BeforeElse            bool   `yaml:"BeforeElse"`
	BeforeLambdaBody      bool   `yaml:"BeforeLambdaBody"`
	IndentBraces          bool   `yaml:"IndentBraces"`
	AfterFunction         bool   `yaml:"AfterFunction"`
	BeforeCatch           bool   `yaml:"BeforeCatch"`
	BeforeWhile           bool   `yaml:"BeforeWhile"`
}

// AlignTrailingComments
type AlignTrailingComments struct {
	Kind           string `yaml:"Kind"`
	OverEmptyLines int    `yaml:"OverEmptyLines"`
}

// SpacesInLineCommentPrefix
type SpacesInLineCommentPrefix struct {
	Minimum int `yaml:"Minimum"`
	Maximum int `yaml:"Maximum"`
}

// SpaceBeforeParensOptions
type SpaceBeforeParensOptions struct {
	AfterControlStatements       bool `yaml:"AfterControlStatements"`
	AfterForeachMacros           bool `yaml:"AfterForeachMacros"`
	AfterIfMacros                bool `yaml:"AfterIfMacros"`
	AfterOverloadedOperator      bool `yaml:"AfterOverloadedOperator"`
	AfterRequiresInClause        bool `yaml:"AfterRequiresInClause"`
	AfterFunctionDefinitionName  bool `yaml:"AfterFunctionDefinitionName"`
	AfterFunctionDeclarationName bool `yaml:"AfterFunctionDeclarationName"`
	AfterPlacementOperator       bool `yaml:"AfterPlacementOperator"`
	AfterRequiresInExpression    bool `yaml:"AfterRequiresInExpression"`
	BeforeNonEmptyParentheses    bool `yaml:"BeforeNonEmptyParentheses"`
}

// AlignConsecutiveAssignments
type AlignConsecutiveAssignments struct {
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
	AcrossComments            bool `yaml:"AcrossComments"`
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
}

// KeepEmptyLines
type KeepEmptyLines struct {
	AtEndOfFile    bool `yaml:"AtEndOfFile"`
	AtStartOfBlock bool `yaml:"AtStartOfBlock"`
	AtStartOfFile  bool `yaml:"AtStartOfFile"`
}

// AlignConsecutiveTableGenBreakingDAGArgColons
type AlignConsecutiveTableGenBreakingDAGArgColons struct {
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
	AcrossComments            bool `yaml:"AcrossComments"`
}

// AlignConsecutiveDeclarations
type AlignConsecutiveDeclarations struct {
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
	AcrossComments            bool `yaml:"AcrossComments"`
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
}

// SpacesInParensOptions
type SpacesInParensOptions struct {
	ExceptDoubleParentheses bool `yaml:"ExceptDoubleParentheses"`
	InCStyleCasts           bool `yaml:"InCStyleCasts"`
	InConditionalStatements bool `yaml:"InConditionalStatements"`
	InEmptyParentheses      bool `yaml:"InEmptyParentheses"`
	Other                   bool `yaml:"Other"`
}

// AlignConsecutiveShortCaseStatements
type AlignConsecutiveShortCaseStatements struct {
	AlignCaseColons  bool `yaml:"AlignCaseColons"`
	Enabled          bool `yaml:"Enabled"`
	AcrossEmptyLines bool `yaml:"AcrossEmptyLines"`
	AcrossComments   bool `yaml:"AcrossComments"`
	AlignCaseArrows  bool `yaml:"AlignCaseArrows"`
}

// IntegerLiteralSeparator
type IntegerLiteralSeparator struct {
	Hex              int `yaml:"Hex"`
	HexMinDigits     int `yaml:"HexMinDigits"`
	Binary           int `yaml:"Binary"`
	BinaryMinDigits  int `yaml:"BinaryMinDigits"`
	Decimal          int `yaml:"Decimal"`
	DecimalMinDigits int `yaml:"DecimalMinDigits"`
}

// ClangFormatConfig
type ClangFormatConfig struct {
	QualifierAlignment                           string                                       `yaml:"QualifierAlignment"`
	AlwaysBreakAfterDefinitionReturnType         string                                       `yaml:"AlwaysBreakAfterDefinitionReturnType"`
	AlwaysBreakBeforeMultilineStrings            bool                                         `yaml:"AlwaysBreakBeforeMultilineStrings"`
	BreakBeforeTernaryOperators                  bool                                         `yaml:"BreakBeforeTernaryOperators"`
	BreakFunctionDefinitionParameters            bool                                         `yaml:"BreakFunctionDefinitionParameters"`
	BreakInheritanceList                         string                                       `yaml:"BreakInheritanceList"`
	JavaScriptWrapImports                        bool                                         `yaml:"JavaScriptWrapImports"`
	PenaltyBreakOpenParenthesis                  int                                          `yaml:"PenaltyBreakOpenParenthesis"`
	RemoveBracesLLVM                             bool                                         `yaml:"RemoveBracesLLVM"`
	AllowShortLoopsOnASingleLine                 bool                                         `yaml:"AllowShortLoopsOnASingleLine"`
	AttributeMacros                              []string                                     `yaml:"AttributeMacros"`
	StatementMacros                              []string                                     `yaml:"StatementMacros"`
	AllowAllParametersOfDeclarationOnNextLine    bool                                         `yaml:"AllowAllParametersOfDeclarationOnNextLine"`
	DerivePointerAlignment                       bool                                         `yaml:"DerivePointerAlignment"`
	InsertBraces                                 bool                                         `yaml:"InsertBraces"`
	PenaltyBreakComment                          int                                          `yaml:"PenaltyBreakComment"`
	PenaltyBreakTemplateDeclaration              int                                          `yaml:"PenaltyBreakTemplateDeclaration"`
	SpacesBeforeTrailingComments                 int                                          `yaml:"SpacesBeforeTrailingComments"`
	AlignConsecutiveMacros                       AlignConsecutiveMacros                       `yaml:"AlignConsecutiveMacros"`
	AllowBreakBeforeNoexceptSpecifier            string                                       `yaml:"AllowBreakBeforeNoexceptSpecifier"`
	BreakStringLiterals                          bool                                         `yaml:"BreakStringLiterals"`
	ExperimentalAutoDetectBinPacking             bool                                         `yaml:"ExperimentalAutoDetectBinPacking"`
	AlignTrailingComments                        AlignTrailingComments                        `yaml:"AlignTrailingComments"`
	IndentWrappedFunctionNames                   bool                                         `yaml:"IndentWrappedFunctionNames"`
	MacroBlockEnd                                string                                       `yaml:"MacroBlockEnd"`
	PenaltyIndentedWhitespace                    int                                          `yaml:"PenaltyIndentedWhitespace"`
	DisableFormat                                bool                                         `yaml:"DisableFormat"`
	JavaScriptQuotes                             string                                       `yaml:"JavaScriptQuotes"`
	KeepFormFeed                                 bool                                         `yaml:"KeepFormFeed"`
	RemoveEmptyLinesInUnwrappedLines             bool                                         `yaml:"RemoveEmptyLinesInUnwrappedLines"`
	RemoveParentheses                            string                                       `yaml:"RemoveParentheses"`
	SkipMacroDefinitionBody                      bool                                         `yaml:"SkipMacroDefinitionBody"`
	SpaceAfterTemplateKeyword                    bool                                         `yaml:"SpaceAfterTemplateKeyword"`
	SpaceBeforeJsonColon                         bool                                         `yaml:"SpaceBeforeJsonColon"`
	IncludeBlocks                                string                                       `yaml:"IncludeBlocks"`
	RawStringFormats                             []RawStringFormats                           `yaml:"RawStringFormats"`
	SpacesInLineCommentPrefix                    SpacesInLineCommentPrefix                    `yaml:"SpacesInLineCommentPrefix"`
	BreakArrays                                  bool                                         `yaml:"BreakArrays"`
	CompactNamespaces                            bool                                         `yaml:"CompactNamespaces"`
	IndentGotoLabels                             bool                                         `yaml:"IndentGotoLabels"`
	IndentPPDirectives                           string                                       `yaml:"IndentPPDirectives"`
	SortUsingDeclarations                        string                                       `yaml:"SortUsingDeclarations"`
	SpaceBeforeInheritanceColon                  bool                                         `yaml:"SpaceBeforeInheritanceColon"`
	AlignConsecutiveTableGenBreakingDAGArgColons AlignConsecutiveTableGenBreakingDAGArgColons `yaml:"AlignConsecutiveTableGenBreakingDAGArgColons"`
	AlignEscapedNewlines                         string                                       `yaml:"AlignEscapedNewlines"`
	AllowShortFunctionsOnASingleLine             string                                       `yaml:"AllowShortFunctionsOnASingleLine"`
	AlignConsecutiveBitFields                    AlignConsecutiveBitFields                    `yaml:"AlignConsecutiveBitFields"`
	TabWidth                                     int                                          `yaml:"TabWidth"`
	AlignConsecutiveTableGenCondOperatorColons   AlignConsecutiveTableGenCondOperatorColons   `yaml:"AlignConsecutiveTableGenCondOperatorColons"`
	IncludeIsMainSourceRegex                     string                                       `yaml:"IncludeIsMainSourceRegex"`
	IndentCaseBlocks                             bool                                         `yaml:"IndentCaseBlocks"`
	PackConstructorInitializers                  string                                       `yaml:"PackConstructorInitializers"`
	PenaltyBreakBeforeMemberAccess               int                                          `yaml:"PenaltyBreakBeforeMemberAccess"`
	WrapNamespaceBodyWithEmptyLines              string                                       `yaml:"WrapNamespaceBodyWithEmptyLines"`
	PenaltyExcessCharacter                       int                                          `yaml:"PenaltyExcessCharacter"`
	AllowShortNamespacesOnASingleLine            bool                                         `yaml:"AllowShortNamespacesOnASingleLine"`
	BinPackArguments                             bool                                         `yaml:"BinPackArguments"`
	IndentAccessModifiers                        bool                                         `yaml:"IndentAccessModifiers"`
	LambdaBodyIndentation                        string                                       `yaml:"LambdaBodyIndentation"`
	PenaltyBreakScopeResolution                  int                                          `yaml:"PenaltyBreakScopeResolution"`
	PPIndentWidth                                int                                          `yaml:"PPIndentWidth"`
	SeparateDefinitionBlocks                     string                                       `yaml:"SeparateDefinitionBlocks"`
	AlignOperands                                string                                       `yaml:"AlignOperands"`
	AllowShortIfStatementsOnASingleLine          string                                       `yaml:"AllowShortIfStatementsOnASingleLine"`
	BitFieldColonSpacing                         string                                       `yaml:"BitFieldColonSpacing"`
	BraceWrapping                                BraceWrapping                                `yaml:"BraceWrapping"`
	BreakAfterAttributes                         string                                       `yaml:"BreakAfterAttributes"`
	ShortNamespaceLines                          int                                          `yaml:"ShortNamespaceLines"`
	SpaceBeforeCpp11BracedList                   bool                                         `yaml:"SpaceBeforeCpp11BracedList"`
	SpaceBeforeCtorInitializerColon              bool                                         `yaml:"SpaceBeforeCtorInitializerColon"`
	IndentRequiresClause                         bool                                         `yaml:"IndentRequiresClause"`
	MacroBlockBegin                              string                                       `yaml:"MacroBlockBegin"`
	PenaltyReturnTypeOnItsOwnLine                int                                          `yaml:"PenaltyReturnTypeOnItsOwnLine"`
	AllowShortCaseLabelsOnASingleLine            bool                                         `yaml:"AllowShortCaseLabelsOnASingleLine"`
	AlignConsecutiveDeclarations                 AlignConsecutiveDeclarations                 `yaml:"AlignConsecutiveDeclarations"`
	BreakBeforeBinaryOperators                   string                                       `yaml:"BreakBeforeBinaryOperators"`
	SpaceBeforeParens                            string                                       `yaml:"SpaceBeforeParens"`
	SpaceBeforeParensOptions                     SpaceBeforeParensOptions                     `yaml:"SpaceBeforeParensOptions"`
	AllowShortBlocksOnASingleLine                string                                       `yaml:"AllowShortBlocksOnASingleLine"`
	IncludeCategories                            []IncludeCategories                          `yaml:"IncludeCategories"`
	ReflowComments                               string                                       `yaml:"ReflowComments"`
	RemoveSemicolon                              bool                                         `yaml:"RemoveSemicolon"`
	SpacesInAngles                               string                                       `yaml:"SpacesInAngles"`
	SpacesInContainerLiterals                    bool                                         `yaml:"SpacesInContainerLiterals"`
	SpacesInParensOptions                        SpacesInParensOptions                        `yaml:"SpacesInParensOptions"`
	WhitespaceSensitiveMacros                    []string                                     `yaml:"WhitespaceSensitiveMacros"`
	AllowShortCaseExpressionOnASingleLine        bool                                         `yaml:"AllowShortCaseExpressionOnASingleLine"`
	ColumnLimit                                  int                                          `yaml:"ColumnLimit"`
	IndentExportBlock                            bool                                         `yaml:"IndentExportBlock"`
	IndentExternBlock                            string                                       `yaml:"IndentExternBlock"`
	SpaceBeforeRangeBasedForLoopColon            bool                                         `yaml:"SpaceBeforeRangeBasedForLoopColon"`
	AccessModifierOffset                         int                                          `yaml:"AccessModifierOffset"`
	InsertNewlineAtEOF                           bool                                         `yaml:"InsertNewlineAtEOF"`
	SpaceBeforeCaseColon                         bool                                         `yaml:"SpaceBeforeCaseColon"`
	SpaceBeforeSquareBrackets                    bool                                         `yaml:"SpaceBeforeSquareBrackets"`
	UseTab                                       string                                       `yaml:"UseTab"`
	Language                                     string                                       `yaml:"Language"`
	AlignConsecutiveAssignments                  AlignConsecutiveAssignments                  `yaml:"AlignConsecutiveAssignments"`
	AlignConsecutiveShortCaseStatements          AlignConsecutiveShortCaseStatements          `yaml:"AlignConsecutiveShortCaseStatements"`
	AllowAllArgumentsOnNextLine                  bool                                         `yaml:"AllowAllArgumentsOnNextLine"`
	LineEnding                                   string                                       `yaml:"LineEnding"`
	NamespaceIndentation                         string                                       `yaml:"NamespaceIndentation"`
	ObjCSpaceAfterProperty                       bool                                         `yaml:"ObjCSpaceAfterProperty"`
	SpaceInEmptyBlock                            bool                                         `yaml:"SpaceInEmptyBlock"`
	AlignConsecutiveTableGenDefinitionColons     AlignConsecutiveTableGenDefinitionColons     `yaml:"AlignConsecutiveTableGenDefinitionColons"`
	IndentCaseLabels                             bool                                         `yaml:"IndentCaseLabels"`
	PenaltyBreakFirstLessLess                    int                                          `yaml:"PenaltyBreakFirstLessLess"`
	ReferenceAlignment                           string                                       `yaml:"ReferenceAlignment"`
	SpacesInSquareBrackets                       bool                                         `yaml:"SpacesInSquareBrackets"`
	Standard                                     string                                       `yaml:"Standard"`
	BreakAfterJavaFieldAnnotations               bool                                         `yaml:"BreakAfterJavaFieldAnnotations"`
	EmptyLineBeforeAccessModifier                string                                       `yaml:"EmptyLineBeforeAccessModifier"`
	MainIncludeChar                              string                                       `yaml:"MainIncludeChar"`
	StatementAttributeLikeMacros                 []string                                     `yaml:"StatementAttributeLikeMacros"`
	TableGenBreakInsideDAGArg                    string                                       `yaml:"TableGenBreakInsideDAGArg"`
	Cpp11BracedListStyle                         bool                                         `yaml:"Cpp11BracedListStyle"`
	ObjCBinPackProtocolList                      string                                       `yaml:"ObjCBinPackProtocolList"`
	AlignAfterOpenBracket                        string                                       `yaml:"AlignAfterOpenBracket"`
	AlignArrayOfStructures                       string                                       `yaml:"AlignArrayOfStructures"`
	IntegerLiteralSeparator                      IntegerLiteralSeparator                      `yaml:"IntegerLiteralSeparator"`
	KeepEmptyLines                               KeepEmptyLines                               `yaml:"KeepEmptyLines"`
	MaxEmptyLinesToKeep                          int                                          `yaml:"MaxEmptyLinesToKeep"`
	ObjCBlockIndentWidth                         int                                          `yaml:"ObjCBlockIndentWidth"`
	PenaltyBreakString                           int                                          `yaml:"PenaltyBreakString"`
	SpaceAroundPointerQualifiers                 string                                       `yaml:"SpaceAroundPointerQualifiers"`
	BreakAdjacentStringLiterals                  bool                                         `yaml:"BreakAdjacentStringLiterals"`
	BreakAfterReturnType                         string                                       `yaml:"BreakAfterReturnType"`
	BreakBeforeBraces                            string                                       `yaml:"BreakBeforeBraces"`
	BreakTemplateDeclarations                    bool                                         `yaml:"BreakTemplateDeclarations"`
	RequiresClausePosition                       string                                       `yaml:"RequiresClausePosition"`
	SortIncludes                                 string                                       `yaml:"SortIncludes"`
	SpacesInParens                               string                                       `yaml:"SpacesInParens"`
	AllowShortEnumsOnASingleLine                 bool                                         `yaml:"AllowShortEnumsOnASingleLine"`
	BreakBeforeConceptDeclarations               string                                       `yaml:"BreakBeforeConceptDeclarations"`
	BreakBeforeInlineASMColon                    string                                       `yaml:"BreakBeforeInlineASMColon"`
	IndentWidth                                  int                                          `yaml:"IndentWidth"`
	InsertTrailingCommas                         string                                       `yaml:"InsertTrailingCommas"`
	SortJavaStaticImport                         string                                       `yaml:"SortJavaStaticImport"`
	VerilogBreakBetweenInstancePorts             bool                                         `yaml:"VerilogBreakBetweenInstancePorts"`
	AllowShortLambdasOnASingleLine               string                                       `yaml:"AllowShortLambdasOnASingleLine"`
	BinPackParameters                            string                                       `yaml:"BinPackParameters"`
	PenaltyBreakBeforeFirstCallParameter         int                                          `yaml:"PenaltyBreakBeforeFirstCallParameter"`
	SpaceAfterCStyleCast                         bool                                         `yaml:"SpaceAfterCStyleCast"`
	SpaceBeforeAssignmentOperators               bool                                         `yaml:"SpaceBeforeAssignmentOperators"`
	ForEachMacros                                []string                                     `yaml:"ForEachMacros"`
	ObjCBreakBeforeNestedBlockParam              bool                                         `yaml:"ObjCBreakBeforeNestedBlockParam"`
	PointerAlignment                             string                                       `yaml:"PointerAlignment"`
	AllowShortCompoundRequirementOnASingleLine   bool                                         `yaml:"AllowShortCompoundRequirementOnASingleLine"`
	BreakConstructorInitializers                 string                                       `yaml:"BreakConstructorInitializers"`
	PenaltyBreakAssignment                       int                                          `yaml:"PenaltyBreakAssignment"`
	RequiresExpressionIndentation                string                                       `yaml:"RequiresExpressionIndentation"`
	SpaceAfterLogicalNot                         bool                                         `yaml:"SpaceAfterLogicalNot"`
	BreakBinaryOperations                        string                                       `yaml:"BreakBinaryOperations"`
	FixNamespaceComments                         bool                                         `yaml:"FixNamespaceComments"`
	IfMacros                                     []string                                     `yaml:"IfMacros"`
	IncludeIsMainRegex                           string                                       `yaml:"IncludeIsMainRegex"`
	ObjCSpaceBeforeProtocolList                  bool                                         `yaml:"ObjCSpaceBeforeProtocolList"`
	CommentPragmas                               string                                       `yaml:"CommentPragmas"`
	ConstructorInitializerIndentWidth            int                                          `yaml:"ConstructorInitializerIndentWidth"`
	ContinuationIndentWidth                      int                                          `yaml:"ContinuationIndentWidth"`
	EmptyLineAfterAccessModifier                 string                                       `yaml:"EmptyLineAfterAccessModifier"`
}

// AlignConsecutiveTableGenCondOperatorColons
type AlignConsecutiveTableGenCondOperatorColons struct {
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
	AcrossComments            bool `yaml:"AcrossComments"`
}

// IncludeCategories
type IncludeCategories struct {
	SortPriority  int    `yaml:"SortPriority"`
	CaseSensitive bool   `yaml:"CaseSensitive"`
	Regex         string `yaml:"Regex"`
	Priority      int    `yaml:"Priority"`
}

// AlignConsecutiveTableGenDefinitionColons
type AlignConsecutiveTableGenDefinitionColons struct {
	AcrossComments            bool `yaml:"AcrossComments"`
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
}

// AlignConsecutiveMacros
type AlignConsecutiveMacros struct {
	Enabled                   bool `yaml:"Enabled"`
	AcrossEmptyLines          bool `yaml:"AcrossEmptyLines"`
	AcrossComments            bool `yaml:"AcrossComments"`
	AlignCompound             bool `yaml:"AlignCompound"`
	AlignFunctionDeclarations bool `yaml:"AlignFunctionDeclarations"`
	AlignFunctionPointers     bool `yaml:"AlignFunctionPointers"`
	PadOperators              bool `yaml:"PadOperators"`
}

