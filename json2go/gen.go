package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/waitgroup"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/tidwall/gjson"
)

type FieldInfo struct {
	JsonName string
	GoName   string
	Type     string
}

type StructInfo struct {
	Name   string
	Fields []FieldInfo
}

var (
	structMap          = make(map[string]StructInfo)
	fieldNameGenerator = make(map[string]map[string]string)
	structCache        = make(map[string]string)
)

func cleanName(s string) string {
	if s == "" {
		return "Field"
	}

	if strings.HasPrefix(s, "^(") && strings.HasSuffix(s, ")$") {
		return "PatternProperties"
	}

	if len(s) > 40 {
		hash := sha1.Sum([]byte(s))
		return "Field_" + hex.EncodeToString(hash[:])[:8]
	}

	var b strings.Builder
	first := true
	for _, r := range s {
		if first && !unicode.IsLetter(r) {
			continue
		}
		first = false

		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else if r == '$' {
			b.WriteString("Ref")
		}
	}
	s = b.String()

	if s == "" {
		s = "Field"
	}

	if isGoReservedWord(s) {
		s += "Val"
	}

	return s
}

func isGoReservedWord(s string) bool {
	reserved := []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
		"int", "string", "bool", "true", "false", "nil",
		"append", "make", "len", "cap", "new", "complex128",
		"float32", "float64", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"complex64", "complex", "real", "imag", "panic", "recover",
		"iota", "rune", "copy", "delete", "interface", "struct",
	}

	for _, word := range reserved {
		if strings.EqualFold(s, word) {
			return true
		}
	}

	return false
}

func ensureUniqueFieldName(structName, jsonName string) string {
	if _, ok := fieldNameGenerator[structName]; !ok {
		fieldNameGenerator[structName] = make(map[string]string)
	}

	if goName, ok := fieldNameGenerator[structName][jsonName]; ok {
		return goName
	}

	if jsonName == "$ref" {
		fieldNameGenerator[structName][jsonName] = "Ref"
		return "Ref"
	}

	if jsonName == "$$id" {
		fieldNameGenerator[structName][jsonName] = "ID"
		return "ID"
	}

	baseName := cleanName(jsonName)
	if baseName == "" {
		baseName = "Field"
	}

	r := []rune(baseName)
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}
	result := string(r)

	if isGoReservedWord(result) {
		result += "Val"
	}

	counter := 1
	originalResult := result
	for {
		exists := false
		for _, existingName := range fieldNameGenerator[structName] {
			if existingName == result {
				exists = true
				break
			}
		}

		if !exists {
			break
		}
		result = fmt.Sprintf("%s%d", originalResult, counter)
		counter++
	}

	fieldNameGenerator[structName][jsonName] = result
	return result
}

func structSignature(fields []FieldInfo) string {
	h := sha1.New()
	for _, f := range fields {
		fmt.Fprintf(h, "%s:%s", f.JsonName, f.Type)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func getOrCreateStruct(fields []FieldInfo, contextName string) string {
	sig := structSignature(fields)
	if name, ok := structCache[sig]; ok {
		return name
	}

	name := ""
	if contextName != "" {
		name = cleanName(contextName)
	}

	if name == "" {
		name = "Object"
	}

	counter := 1
	baseName := name
	for {
		if _, exists := structMap[name]; !exists {
			break
		}
		name = fmt.Sprintf("%s%d", baseName, counter)
		counter++
	}

	for i := range fields {
		fields[i].GoName = ensureUniqueFieldName(name, fields[i].JsonName)
	}

	structMap[name] = StructInfo{
		Name:   name,
		Fields: fields,
	}
	structCache[sig] = name
	return name
}

func processObject(data gjson.Result) []FieldInfo {
	if !data.IsObject() {
		return nil
	}

	fields := make([]FieldInfo, 0)

	if properties := data.Get("properties"); properties.Exists() && properties.IsObject() {
		return processObject(properties)
	}

	data.ForEach(func(key, val gjson.Result) bool {
		jsonName := key.String()
		var fieldType string

		if val.IsObject() {
			if properties := val.Get("properties"); properties.Exists() && properties.IsObject() {
				nestedFields := processObject(val)
				fieldType = getOrCreateStruct(nestedFields, cleanName(jsonName))
			} else if val.Get("type").Exists() {
				typeVal := val.Get("type").String()
				switch typeVal {
				case "string":
					fieldType = "string"
				case "number":
					fieldType = "float64"
				case "integer":
					fieldType = "int"
				case "boolean":
					fieldType = "bool"
				case "array":
					if items := val.Get("items"); items.Exists() {
						elemType := determineItemType(items)
						fieldType = "[]" + elemType
					} else {
						fieldType = "[]any"
					}
				case "object":
					if properties := val.Get("properties"); properties.Exists() {
						nestedFields := processObject(properties)
						fieldType = getOrCreateStruct(nestedFields, cleanName(jsonName))
					} else {
						fieldType = "any"
					}
				default:
					fieldType = "any"
				}
			} else {
				nestedFields := processObject(val)
				fieldType = getOrCreateStruct(nestedFields, cleanName(jsonName))
			}
		} else if val.IsArray() {
			fieldType = "[]" + determineArrayItemType(val)
		} else {
			switch val.Type {
			case gjson.String:
				fieldType = "string"
			case gjson.Number:
				if strings.Contains(val.Raw, ".") {
					fieldType = "float64"
				} else {
					fieldType = "int"
				}
			case gjson.True, gjson.False:
				fieldType = "bool"
			default:
				fieldType = "any"
			}
		}

		fields = append(fields, FieldInfo{
			JsonName: jsonName,
			Type:     fieldType,
		})
		return true
	})

	return fields
}

func determineArrayItemType(val gjson.Result) string {
	if !val.IsArray() {
		return "any"
	}

	if firstElem := val.Get("0"); firstElem.Exists() {
		switch firstElem.Type {
		case gjson.String:
			return "string"
		case gjson.Number:
			if strings.Contains(firstElem.Raw, ".") {
				return "float64"
			}
			return "int"
		case gjson.True, gjson.False:
			return "bool"
		case gjson.JSON:
			if firstElem.IsObject() {
				nestedFields := processObject(firstElem)
				return getOrCreateStruct(nestedFields, "ArrayElem")
			}
		}
	}
	return "any"
}

func determineItemType(items gjson.Result) string {
	if items.Get("type").Exists() {
		switch items.Get("type").String() {
		case "string":
			return "string"
		case "number":
			return "float64"
		case "integer":
			return "int"
		case "boolean":
			return "bool"
		case "object":
			if properties := items.Get("properties"); properties.Exists() {
				nestedFields := processObject(properties)
				return getOrCreateStruct(nestedFields, "ArrayItem")
			}
			return "any"
		}
	}
	return "any"
}

func generateStructCode() string {
	sb := &strings.Builder{}

	names := make([]string, 0, len(structMap))
	for name := range structMap {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		info := structMap[name]
		sb.WriteString(fmt.Sprintf("type %s struct {\n", name))

		for _, field := range info.Fields {
			sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", field.GoName, field.Type, field.JsonName))
		}
		sb.WriteString("}\n\n")
	}

	return sb.String()
}

func generateInstanceCode(data gjson.Result, structType string) string {
	structInfo, ok := structMap[structType]
	if !ok {
		return "nil"
	}

	sb := &strings.Builder{}
	sb.WriteString(structType + "{\n")

	for _, field := range structInfo.Fields {
		val := data.Get(field.JsonName)
		if !val.Exists() {
			continue
		}

		sb.WriteString(fmt.Sprintf("\t%s: ", field.GoName))

		switch field.Type {
		case "string":
			sb.WriteString(fmt.Sprintf("%q,\n", val.String()))
		case "int":
			sb.WriteString(fmt.Sprintf("%d,\n", val.Int()))
		case "float64":
			sb.WriteString(fmt.Sprintf("%v,\n", val.Float()))
		case "bool":
			sb.WriteString(fmt.Sprintf("%v,\n", val.Bool()))
		case "any":
			sb.WriteString("nil,\n")
		default:
			if strings.HasPrefix(field.Type, "[]") {
				// 修复点：为数组元素生成正确的字面量值
				sb.WriteString("[]" + field.Type[2:] + "{\n")
				val.ForEach(func(_, elem gjson.Result) bool {
					sb.WriteString("\t\t")

					// 处理字符串数组元素
					if field.Type[2:] == "string" {
						sb.WriteString(fmt.Sprintf("%q,\n", elem.String()))
						return true
					}

					// 处理其他类型的数组元素
					if strings.HasPrefix(field.Type, "[]*") {
						sb.WriteString("&")
						sb.WriteString(generateInstanceCode(elem, field.Type[3:]))
					} else {
						sb.WriteString(generateInstanceCode(elem, field.Type[2:]))
					}
					sb.WriteString(",\n")
					return true
				})
				sb.WriteString("\t},\n")
			} else {
				sb.WriteString(generateInstanceCode(val, field.Type))
				sb.WriteString(",\n")
			}
		}
	}

	sb.WriteString("}")
	return sb.String()
}

func Walk(root string) {
	g := waitgroup.New()
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			g.Go(func() {
				Run(path)
			})
		}
		return err
	})
	g.Wait()
}

func Run(path string) {
	data := gjson.ParseBytes(mylog.Check2(os.ReadFile(path)))
	rootFields := processObject(data)
	rootType := getOrCreateStruct(rootFields, "RootConfig")

	code := fmt.Sprintf(`
// Automatically generated code - DO NOT EDIT.
package main

import (
	"fmt"
)

%s

type RootContainer struct {
	Data *%s `+"`json:\"data\"`"+`
}

func main() {
	data := %s
	root := RootContainer{Data: &data}
	fmt.Printf("%%#v\\n", root)
}
`, generateStructCode(), rootType, generateInstanceCode(data, rootType))

	base := filepath.Base(path)
	base = stream.TrimExtension(base)
	filePath := filepath.Join("tmp", base, base+"_gen.go")
	stream.WriteGoFile(filePath, code)
	fmt.Println("try running the generated file...")
	stream.RunCommandArgs("go", "run", filePath)
}
