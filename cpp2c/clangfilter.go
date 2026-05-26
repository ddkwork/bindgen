package cpp2c

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/gjson"
)

func filteredAstPath(inputHeader string) string {
	return CacheFileRoot(inputHeader) + ".filtered.json"
}

func dumpStack(stack []*AstNode) string {
	var ret string
	for i := len(stack) - 1; i >= 0; i-- {
		ret += "<" + stack[i].Kind
	}
	return ret
}

func allowExpressions(stack []*AstNode) bool {
	for _, node := range stack {
		if node.Kind == "ParmVarDecl" || node.Kind == "EnumConstantDecl" {
			return true
		}
	}
	return false
}

func parseInner(ctx *ParseCtx, kind string) bool {
	switch kind {
	case "FieldDecl", "TranslationUnitDecl", "NamespaceDecl":
		return true
	case "CXXRecordDecl", "EnumDecl", "EnumConstantDecl":
		return true
	case "CXXConstructorDecl", "CXXMethodDecl", "CXXConversionDecl":
		return ctx.stack[len(ctx.stack)-1].Kind == "CXXRecordDecl"
	case "CXXDestructorDecl":
		return false
	case "ParmVarDecl":
		return true
	case "TypeAliasDecl", "TypedefDecl":
		return false
	case "CompoundStmt", "CXXCtorInitializer":
		return false
	case "LinkageSpecDecl", "AccessSpecDecl", "VarDecl",
		"FileScopeAsmDecl", "FriendDecl", "UsingShadowDecl", "UsingDecl",
		"StaticAssertDecl", "ElaboratedType", "FullComment", "ParagraphComment",
		"EmptyDecl", "IndirectFieldDecl":
		return false
	case "FunctionDecl":
		return true
	case "ClassTemplateDecl", "TypeAliasTemplateDecl",
		"ClassTemplateSpecializationDecl",
		"ClassTemplatePartialSpecializationDecl", "FunctionTemplateDecl",
		"VarTemplatePartialSpecializationDecl",
		"VarTemplateSpecializationDecl",
		"VarTemplateDecl", "BuiltinTemplateDecl":
		return false
	case "ConstantExpr":
		return true
	case "ImplicitCastExpr", "IntegerLiteral", "ParenExpr":
		return true
	case "ConstructorUsingShadowDecl":
		return false
	case "PragmaDetectMismatchDecl", "PragmaCommentDecl":
		return false
	}

	if strings.HasSuffix(kind, "Attr") {
		return false
	}

	if allowExpressions(ctx.stack) {
		return true
	}

	return true
}

func keepField(ctx *ParseCtx, kind, field string, val interface{}) bool {
	switch field {
	case "loc":
		switch kind {
		case "TranslationUnitDecl", "NamespaceDecl", "CXXRecordDecl", "EnumDecl",
			"TypeAliasDecl", "TypedefDecl":
			return true
		case "AccessSpecDecl":
			return true
		}
		return false

	case "name", "type", "tagUsed", "definitionData", "bases", "access",
		"isImplicit", "explicitlyDeleted", "explicitlyDefaulted",
		"fixedUnderlyingType", "value", "storageClass", "virtual", "pure",
		"file":
		return true
	case "range":
		return false
	case "isReferenced", "language", "previousDecl", "originalNamespace", "isInline",
		"isUsed", "mangledName", "inline", "constexpr", "parentDeclContextId",
		"completeDefinition", "variadic", "target", "hasBraces", "init", "baseInit",
		"nominatedNamespace", "implicit", "anyInit", "valueCategory", "castKind",
		"scopedEnumTag", "argType", "opcode", "referencedDecl", "hasInClassInitializer",
		"isPostfix", "canOverflow", "ctorType", "elidable", "hadMultipleCandidates",
		"storageDuration", "conversionFunc", "constructionKind", "ownedTagDecl",
		"isBitfield", "delegatingInit", "mutable", "boundToLValueRef",
		"cleanupsHaveSideEffects", "temp", "dtor", "inherited", "isPartOfExplicitCast",
		"adl", "decl", "list", "tls", "nonOdrUseReason", "zeroing":
		return false
	}
	return true
}

func isFromPrivateHeader(node *AstNode) bool {
	if loc, ok := node.Fields["loc"].(map[string]interface{}); ok {
		if file, ok := loc["file"].(string); ok {
			return strings.Contains(file, `/private/`) ||
				strings.Contains(file, `\private\`) ||
				strings.Contains(strings.ToLower(filepath.Base(file)), "private")
		}
	}
	return false
}

func isNotFromQtSdk(node *AstNode) bool {
	loc, ok := node.Fields["loc"].(map[string]interface{})
	if !ok {
		return false // no loc info, keep it
	}
	// Try loc.file first, then expansionLoc.file, then spellingLoc.file
	candidates := []string{}
	if file, ok := loc["file"].(string); ok {
		candidates = append(candidates, file)
	}
	if el, ok := loc["expansionLoc"].(map[string]interface{}); ok {
		if file, ok := el["file"].(string); ok {
			candidates = append(candidates, file)
		}
	}
	if sl, ok := loc["spellingLoc"].(map[string]interface{}); ok {
		if file, ok := sl["file"].(string); ok {
			candidates = append(candidates, file)
		}
	}
	if len(candidates) == 0 {
		return false // no file paths, keep it
	}
	for _, f := range candidates {
		normalized := strings.ReplaceAll(f, `\`, `/`)
		if strings.Contains(normalized, `qt_static/include/`) {
			return false // at least one candidate is from Qt SDK
		}
	}
	return true // none of the candidates are from Qt SDK
}

func isQtProjectFromCflags(cflags []string, inputHeader string) bool {
	for _, flag := range cflags {
		normalized := strings.ReplaceAll(flag, `\`, `/`)
		if strings.Contains(normalized, `qt_static/include/`) ||
			strings.Contains(normalized, `qt_static\\include\\`) {
			return true
		}
		if strings.HasPrefix(flag, "-I") {
			path := flag[2:]
			normalized := strings.ReplaceAll(path, `\`, `/`)
			if strings.Contains(normalized, `qt_static/include/`) {
				return true
			}
		}
		if strings.Contains(flag, "QT_VERSION_MAJOR=6") {
			return true
		}
	}

	if strings.Contains(inputHeader, `qt_static`) {
		return true
	}

	return false
}

func keepInner(ctx *ParseCtx, parentKind string, child *AstNode) bool {
	if isFromPrivateHeader(child) {
		return false
	}
	switch child.Kind {
	case "NamespaceDecl", "CXXRecordDecl", "EnumDecl", "TypeAliasDecl", "TypedefDecl",
		"ParmVarDecl":
		if ctx.IsQtProject && isNotFromQtSdk(child) {
			return false
		}
		return true

	case "StaticAssertDecl", "ClassTemplateDecl",
		"ClassTemplateSpecializationDecl",
		"ClassTemplatePartialSpecializationDecl",
		"FunctionTemplateDecl",
		"BuiltinTemplateDecl",
		"VarTemplatePartialSpecializationDecl",
		"VarTemplateSpecializationDecl",
		"TypeAliasTemplateDecl",
		"VarTemplateDecl":
		return false
	case "FileScopeAsmDecl",
		"FriendDecl",
		"ElaboratedType":
		return false
	case "FunctionDecl":
		return parentKind == "TranslationUnitDecl" || parentKind == "LinkageSpecDecl" || parentKind == "NamespaceDecl"
	case "VarDecl", "FieldDecl", "IndirectFieldDecl":
		return parentKind == "CXXRecordDecl"

	case "CXXConstructorDecl", "CXXMethodDecl", "CXXConversionDecl", "CXXDestructorDecl":
		return parentKind == "CXXRecordDecl"

	case "LinkageSpecDecl":
		return false

	case "OverrideAttr", "DeprecatedAttr":
		return true
	case "AccessSpecDecl":
		return parentKind == "CXXRecordDecl"
	case "UsingDirectiveDecl",
		"UsingDecl",
		"UsingShadowDecl":
		return false
	case "FullComment", "ParagraphComment", "EmptyDecl":
		return false
	case "CompoundStmt", "CXXCtorInitializer":
		return false
	case "EnumConstantDecl":
		return true
	case "ConstructorUsingShadowDecl":
		return false
	}

	if strings.HasSuffix(child.Kind, "Attr") {
		return false
	}

	if allowExpressions(ctx.stack) {
		return true
	}

	return true
}

func filterAst(in io.Reader, isQtProject bool) (*AstNode, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	var pc = ParseCtx{
		IsQtProject:      isQtProject,
		ShouldParseInner: parseInner,
		ShouldKeepField:  keepField,
		ShouldKeepInner:  keepInner,
	}
	return parseClangAst(&pc, string(data))
}

func writeFilteredCache(ast *AstNode, inputHeader string) {
	path := filteredAstPath(inputHeader)
	os.MkdirAll(filepath.Dir(path), 0755)
	f, err := os.Create(path)
	if err != nil {
		panic("could not create filtered AST cache for " + inputHeader + ": " + err.Error())
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(ast)
	if err != nil {
		panic("could not encode filtered AST cache: " + err.Error())
	}
}

func readFilteredCache(inputHeader string) (*AstNode, error) {
	data, err := os.ReadFile(filteredAstPath(inputHeader))
	if err != nil {
		return nil, err
	}
	result := gjson.ParseBytes(data)
	if !result.IsObject() {
		return nil, fmt.Errorf("expected JSON object in cache")
	}
	ast, err := parseCachedAstFromResult(result)
	if err != nil {
		return nil, fmt.Errorf("could not decode filtered AST cache: %w", err)
	}
	return ast, nil
}

// Get or create the filtered AST.
// Two-layer cache:
//  1. *.filtered.json => filtered result (fastest reload)
//  2. *.json          => full clang output (re-filter without re-running clang)
//  3. clang exec      => slow fallback
//
// To re-filter after changing filter rules: delete only *.filtered.json
// To force full rebuild: delete cachedir/ entirely
func getFilteredAst(inputHeader string, cflags []string) *AstNode {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	isQtProject := isQtProjectFromCflags(cflags, inputHeader)

	ast, err := readFilteredCache(inputHeader)
	if err == nil {
		return ast
	}

	raw, err := readAstCache(inputHeader)
	if err == nil {
		ast, err = filterAst(bytes.NewReader(raw), isQtProject)
		if err != nil {
			return nil
		}
		writeFilteredCache(ast, inputHeader)
		return ast
	}

	ast, err = clangExec(inputHeader, cflags)
	if err != nil {
		return nil
	}

	writeFilteredCache(ast, inputHeader)
	return ast
}

func ClearFilteredCache(cacheDir string) {
	entries, err := os.ReadDir(cacheDir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".filtered.json") {
			path := filepath.Join(cacheDir, e.Name())
			os.Remove(path)
		}
	}
}
