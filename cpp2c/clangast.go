package cpp2c

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/golibrary/gjson"
)

func isSourceLoc(field string) bool {
	switch field {
	case "loc", "spellingLoc", "expansionLoc", "begin", "end", "attrLoc":
		return true
	}
	return false
}

type HexUint64 uint64

func (h HexUint64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"0x%x\"", uint64(h))), nil
}

func (h *HexUint64) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	val, err := strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
	if err != nil {
		return err
	}
	*h = HexUint64(val)
	return nil
}

func (h HexUint64) String() string {
	return fmt.Sprintf("0x%x", uint64(h))
}

type AstNode struct {
	Id     HexUint64              `json:"id"`
	Kind   string                 `json:"kind"`
	Inner  []*AstNode             `json:"inner,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type ShouldParseInner func(ctx *ParseCtx, kind string) bool
type ShouldKeepField func(ctx *ParseCtx, kind, field string, value interface{}) bool
type ShouldKeepInner func(ctx *ParseCtx, parentKind string, child *AstNode) bool

type ParseCtx struct {
	stack            []*AstNode
	file             string
	IsQtProject      bool
	ShouldParseInner ShouldParseInner
	ShouldKeepField  ShouldKeepField
	ShouldKeepInner  ShouldKeepInner
}

func gjsonToInterface(result gjson.Result, ctx *ParseCtx, sourceLoc bool) interface{} {
	switch result.Type {
	case gjson.Null:
		return nil
	case gjson.True:
		return true
	case gjson.False:
		return false
	case gjson.Number:
		return result.Float()
	case gjson.String:
		return result.String()
	case gjson.JSON:
		if result.IsArray() {
			var arr []interface{}
			result.ForEach(func(_, v gjson.Result) bool {
				arr = append(arr, gjsonToInterface(v, ctx, false))
				return true
			})
			return arr
		}
		if result.IsObject() {
			m := make(map[string]interface{})
			hasFileField := false
			subLoc := false
			result.ForEach(func(key, val gjson.Result) bool {
				keyStr := key.String()
				isSrcLoc := isSourceLoc(keyStr)
				valInterface := gjsonToInterface(val, ctx, isSrcLoc)
				m[keyStr] = valInterface
				if sourceLoc {
					if keyStr == "file" {
						if s, ok := valInterface.(string); ok {
							ctx.file = s
							hasFileField = true
						}
					} else if keyStr == "spellingLoc" {
						subLoc = true
					}
				}
				return true
			})
			if sourceLoc && !hasFileField && !subLoc {
				m["file"] = ctx.file
			}
			return m
		}
	}
	return nil
}

func gjsonToInterfaceSimple(result gjson.Result) interface{} {
	switch result.Type {
	case gjson.Null:
		return nil
	case gjson.True:
		return true
	case gjson.False:
		return false
	case gjson.Number:
		return result.Float()
	case gjson.String:
		return result.String()
	case gjson.JSON:
		if result.IsArray() {
			var arr []interface{}
			result.ForEach(func(_, v gjson.Result) bool {
				arr = append(arr, gjsonToInterfaceSimple(v))
				return true
			})
			return arr
		}
		if result.IsObject() {
			m := make(map[string]interface{})
			result.ForEach(func(key, val gjson.Result) bool {
				m[key.String()] = gjsonToInterfaceSimple(val)
				return true
			})
			return m
		}
	}
	return nil
}

func trackFileInResult(result gjson.Result, ctx *ParseCtx) {
	if result.IsArray() {
		result.ForEach(func(_, child gjson.Result) bool {
			trackFileInNode(child, ctx)
			return true
		})
	}
}

func trackFileInNode(result gjson.Result, ctx *ParseCtx) {
	if !result.IsObject() {
		return
	}
	result.ForEach(func(key, val gjson.Result) bool {
		keyStr := key.String()
		if isSourceLoc(keyStr) && val.IsObject() {
			fileResult := val.Get("file")
			if fileResult.Exists() {
				ctx.file = fileResult.String()
			}
			trackFileInNode(val, ctx)
		} else if keyStr == "inner" && val.IsArray() {
			trackFileInResult(val, ctx)
		} else if val.IsObject() {
			trackFileInNode(val, ctx)
		} else if val.IsArray() {
			trackFileInResult(val, ctx)
		}
		return true
	})
}

func parseClangAstFromResult(result gjson.Result, ctx *ParseCtx) (*AstNode, error) {
	node := &AstNode{}

	idResult := result.Get("id")
	if idResult.Exists() {
		idStr := idResult.String()
		id, err := strconv.ParseUint(strings.TrimPrefix(idStr, "0x"), 16, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse id %q as hex: %w", idStr, err)
		}
		node.Id = HexUint64(id)
	}

	kindResult := result.Get("kind")
	if kindResult.Exists() {
		node.Kind = kindResult.String()
	}

	innerResult := result.Get("inner")
	if innerResult.Exists() && innerResult.IsArray() {
		if ctx.ShouldParseInner != nil && !ctx.ShouldParseInner(ctx, node.Kind) {
			trackFileInResult(innerResult, ctx)
		} else {
			ctx.stack = append(ctx.stack, node)
			var innerNodes []*AstNode
			var parseErr error
			innerResult.ForEach(func(_, childResult gjson.Result) bool {
				child, err := parseClangAstFromResult(childResult, ctx)
				if err != nil {
					parseErr = err
					return false
				}
				if ctx.ShouldKeepInner == nil || ctx.ShouldKeepInner(ctx, node.Kind, child) {
					innerNodes = append(innerNodes, child)
				}
				return true
			})
			ctx.stack = ctx.stack[:len(ctx.stack)-1]
			if parseErr != nil {
				return nil, parseErr
			}
			node.Inner = innerNodes
		}
	}

	result.ForEach(func(key, val gjson.Result) bool {
		keyStr := key.String()
		switch keyStr {
		case "id", "kind", "inner":
			return true
		default:
			valInterface := gjsonToInterface(val, ctx, isSourceLoc(keyStr))
			if ctx.ShouldKeepField == nil || ctx.ShouldKeepField(ctx, node.Kind, keyStr, valInterface) {
				if node.Fields == nil {
					node.Fields = make(map[string]interface{})
				}
				node.Fields[keyStr] = valInterface
			}
		}
		return true
	})

	return node, nil
}

func parseClangAst(ctx *ParseCtx, jsonStr string) (*AstNode, error) {
	result := gjson.Parse(jsonStr)
	if !result.IsObject() {
		return nil, fmt.Errorf("expected JSON object, got %v", result.Type)
	}
	return parseClangAstFromResult(result, ctx)
}

func parseCachedAstFromResult(result gjson.Result) (*AstNode, error) {
	node := &AstNode{}

	idResult := result.Get("id")
	if idResult.Exists() {
		idStr := idResult.String()
		id, err := strconv.ParseUint(strings.TrimPrefix(idStr, "0x"), 16, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cached id %q as hex: %w", idStr, err)
		}
		node.Id = HexUint64(id)
	}

	kindResult := result.Get("kind")
	if kindResult.Exists() {
		node.Kind = kindResult.String()
	}

	innerResult := result.Get("inner")
	if innerResult.Exists() && innerResult.IsArray() {
		var innerNodes []*AstNode
		var parseErr error
		innerResult.ForEach(func(_, childResult gjson.Result) bool {
			child, err := parseCachedAstFromResult(childResult)
			if err != nil {
				parseErr = err
				return false
			}
			innerNodes = append(innerNodes, child)
			return true
		})
		if parseErr != nil {
			return nil, parseErr
		}
		node.Inner = innerNodes
	}

	fieldsResult := result.Get("fields")
	if fieldsResult.Exists() && fieldsResult.IsObject() {
		node.Fields = make(map[string]interface{})
		fieldsResult.ForEach(func(key, val gjson.Result) bool {
			node.Fields[key.String()] = gjsonToInterfaceSimple(val)
			return true
		})
	}

	return node, nil
}
