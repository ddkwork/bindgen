package c2go

import (
	"fmt"
	"strings"

	"github.com/ddkwork/bindgen/cc"
)
func GenerateUnionFields(u *cc.UnionType, unionSize int64) string {
	if u.NumFields() == 0 {
		return ""
	}
	var fields []string
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		fieldType := MapCTypeToGo(ft)
		fieldName := CFieldNameToGoName(f.Name())
		if fieldName == "_" || fieldType == "" {
			return ""
		}
		if sv, ok := ft.(*cc.StructType); ok {
			if HasBitfields(sv) {
				_, backingType, _ := ExtractUnionBitfields(u)
				if backingType != "" {
					fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, backingType))
				} else {
					fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
				}
			} else {
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			}
			continue
		}
		if _, ok := ft.(*cc.UnionType); ok {
			return ""
		}
		if _, ok := ft.(*cc.ArrayType); ok {
			return ""
		}
		fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
	}
	if len(fields) == 0 {
		return ""
	}
	return strings.Join(fields, "\n") + "\n"
}

func UnionAlignedFFIType(u *cc.UnionType, unionSize int64) string {
	maxAlign := 1
	maxSize := int64(0)
	var maxType string
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		align := ft.Align()
		goType := MapCTypeToGo(ft)
		if goType == "" {
			continue
		}
		if ft.Size() > maxSize || (ft.Size() == maxSize && align > maxAlign) {
			maxSize = ft.Size()
			maxAlign = align
			maxType = goType
		}
	}
	if maxType != "" && maxSize >= unionSize {
		return "Data " + maxType
	}
	return fmt.Sprintf("Data [%d]byte", unionSize)
}

func UnionAlignedFFITypeWithSize(types []cc.Type, unionSize int64) string {
	maxAlign := 1
	maxSize := int64(0)
	var maxType string
	for _, t := range types {
		align := t.Align()
		goType := MapCTypeToGo(t)
		if goType == "" {
			continue
		}
		if t.Size() > maxSize || (t.Size() == maxSize && align > maxAlign) {
			maxSize = t.Size()
			maxAlign = align
			maxType = goType
		}
	}
	if maxType != "" && maxSize >= unionSize {
		return "Data " + maxType
	}
	return fmt.Sprintf("Data [%d]byte", unionSize)
}

func GenerateStructFields(t *cc.StructType, structGoName string, forcePacked bool) (fieldStr string, methodStr string, innerTypes []StructInfo) {
	if IsPureBitfieldStruct(t) {
		f, m := GeneratePureBitfieldStruct(t, structGoName)
		return f, m, nil
	}

	isPacked := forcePacked || t.Align() == 1

	var fields []string
	var methods []string
	anonCount := 0
	currentOffset := int64(0)
	var padding int64

	type bitfieldGroup struct {
		offset      int64
		size        int64
		backingName string
		backingType string
		bitfields   []BitfieldInfo
	}

	type anonUnionField struct {
		name   string
		goType string
		ccType cc.Type
		offset int64
		size   int64
		align  int
	}

	var currentBfGroup *bitfieldGroup
	var prevFieldOffset int64 = -1
	var anonUnionCandidates []anonUnionField

	isAnonUnionCandidate := func(ft cc.Type) bool {
		switch v := ft.(type) {
		case *cc.StructType:
			if v.NumFields() == 0 || ft.Size() == 0 {
				return false
			}
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr == "" {
				return false
			}
			goName := CTagToGoName(tagStr)
			if goName == "" || goName == "struct" || strings.Contains(goName, "{") {
				return false
			}
			return true
		case *cc.UnionType:
			if v.NumFields() == 0 || ft.Size() == 0 {
				return false
			}
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr == "" {
				return false
			}
			goName := CTagToGoName(tagStr)
			if goName == "" || goName == "struct" || strings.Contains(goName, "{") {
				return false
			}
			return true
		}
		return false
	}

	flushAnonUnion := func() bool {
		if len(anonUnionCandidates) < 2 {
			return false
		}
		firstOffset := anonUnionCandidates[0].offset
		allSameOffset := true
		for _, af := range anonUnionCandidates {
			if af.offset != firstOffset {
				allSameOffset = false
				break
			}
		}
		if !allSameOffset {
			return false
		}
		maxSize := int64(0)
		for _, af := range anonUnionCandidates {
			if af.size > maxSize {
				maxSize = af.size
			}
		}
		anonCount++
		unionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
		var unionFieldStrs []string
		for _, af := range anonUnionCandidates {
			unionFieldStrs = append(unionFieldStrs, fmt.Sprintf("\t%s %s", af.name, af.goType))
			switch v := af.ccType.(type) {
			case *cc.StructType:
				tag := v.Tag()
				tagStr := string(tag.Src())
				nestedFields, nestedMethods, _ := GenerateStructFields(v, af.goType, isPacked)
				innerTypes = append(innerTypes, StructInfo{
					goName:  af.goType,
					cName:   tagStr,
					fields:  nestedFields,
					methods: nestedMethods,
				})
			case *cc.UnionType:
				tag := v.Tag()
				tagStr := string(tag.Src())
				if tagStr == "" {
					innerTypes = append(innerTypes, StructInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
					})
				} else if unionFields := GenerateUnionFields(v, v.Size()); unionFields != "" {
					for j := 0; j < v.NumFields(); j++ {
						if uf := v.FieldByIndex(j); uf != nil {
							if usv, ok := uf.Type().(*cc.StructType); ok {
								nestedTag := usv.Tag()
								nestedTagStr := string(nestedTag.Src())
								if nestedTagStr != "" && !HasBitfields(usv) {
									nestedF, nestedM, _ := GenerateStructFields(usv, MapCTypeToGo(uf.Type()), isPacked)
									innerTypes = append(innerTypes, StructInfo{
										goName:  MapCTypeToGo(uf.Type()),
										cName:   nestedTagStr,
										fields:  nestedF,
										methods: nestedM,
									})
								}
							}
						}
					}
					innerTypes = append(innerTypes, StructInfo{
						goName: af.goType + "_",
						cName:  tagStr,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, StructInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
					})
				} else {
					innerTypes = append(innerTypes, StructInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
					})
				}
			}
		}
		if !isPacked {
			padding := anonUnionCandidates[0].offset - currentOffset
			if padding > 0 {
				fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
			}
		}
		accessorName := unionName + "_"
		innerTypes = append(innerTypes, StructInfo{
			goName: accessorName,
			cName:  unionName,
			fields: strings.Join(unionFieldStrs, "\n") + "\n",
		})
		var candidateTypes []cc.Type
		for _, af := range anonUnionCandidates {
			candidateTypes = append(candidateTypes, af.ccType)
		}
		innerTypes = append(innerTypes, StructInfo{
			goName: unionName,
			cName:  unionName,
			fields: fmt.Sprintf("\t%s\n", UnionAlignedFFITypeWithSize(candidateTypes, maxSize)),
		})
		fields = append(fields, fmt.Sprintf("\t%s %s", "U", unionName))
		currentOffset = anonUnionCandidates[0].offset + maxSize
		anonUnionCandidates = nil
		prevFieldOffset = -1
		return true
	}

	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		ft := f.Type()
		fieldOffset := f.Offset()

		if !f.IsBitfield() && prevFieldOffset >= 0 && isAnonUnionCandidate(ft) {
			fieldType := MapCTypeToGo(ft)
			fieldName := CFieldNameToGoName(f.Name())
			if fieldName != "_" && fieldType != "" {
				if len(anonUnionCandidates) > 0 && fieldOffset != anonUnionCandidates[0].offset {
					if flushAnonUnion() {
						padding = fieldOffset - currentOffset
					} else {
						for _, af := range anonUnionCandidates {
							fieldPadding := af.offset - currentOffset
							currentOffset = af.offset + af.size
							if !isPacked && fieldPadding > 0 {
								fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
							}
							switch v := af.ccType.(type) {
							case *cc.StructType:
								tag := v.Tag()
								tagStr := string(tag.Src())
								nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, af.goType, isPacked)
								innerTypes = append(innerTypes, StructInfo{
									goName:  af.goType,
									cName:   tagStr,
									fields:  nestedFields,
									methods: nestedMethods,
								})
								innerTypes = append(innerTypes, nestedInner...)
								fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
							case *cc.UnionType:
								tag := v.Tag()
								tagStr := string(tag.Src())
								if tagStr == "" {
									innerTypes = append(innerTypes, StructInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
									})
								} else if unionFields := GenerateUnionFields(v, v.Size()); unionFields != "" {
									innerTypes = append(innerTypes, StructInfo{
										goName: af.goType + "_",
										cName:  tagStr,
										fields: unionFields,
									})
									innerTypes = append(innerTypes, StructInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
									})
								} else {
									innerTypes = append(innerTypes, StructInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
									})
								}
								fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
							}
						}
					}
					anonUnionCandidates = nil
				}
				anonUnionCandidates = append(anonUnionCandidates, anonUnionField{
					name:   fieldName,
					goType: fieldType,
					ccType: ft,
					offset: fieldOffset,
					size:   ft.Size(),
					align:  ft.Align(),
				})
				prevFieldOffset = fieldOffset
				continue
			}
		}

		if len(anonUnionCandidates) > 0 {
			if flushAnonUnion() {
				padding = fieldOffset - currentOffset
			} else {
				for _, af := range anonUnionCandidates {
					fieldPadding := af.offset - currentOffset
					currentOffset = af.offset + af.size
					if !isPacked && fieldPadding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
					}
					switch v := af.ccType.(type) {
					case *cc.StructType:
						tag := v.Tag()
						tagStr := string(tag.Src())
						nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, af.goType, isPacked)
						innerTypes = append(innerTypes, StructInfo{
							goName:  af.goType,
							cName:   tagStr,
							fields:  nestedFields,
							methods: nestedMethods,
						})
						innerTypes = append(innerTypes, nestedInner...)
					case *cc.UnionType:
						tag := v.Tag()
						tagStr := string(tag.Src())
						if tagStr != "" {
							if uf := GenerateUnionFields(v, v.Size()); uf != "" {
								innerTypes = append(innerTypes, StructInfo{
									goName: af.goType + "_",
									cName:  tagStr,
									fields: uf,
								})
								innerTypes = append(innerTypes, StructInfo{
									goName: af.goType,
									cName:  tagStr,
									fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
								})
							} else {
								innerTypes = append(innerTypes, StructInfo{
									goName: af.goType,
									cName:  tagStr,
									fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
								})
							}
						}
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
				}
			}
			anonUnionCandidates = nil
		}
		prevFieldOffset = fieldOffset

		if f.IsBitfield() {
			bfName := CFieldNameToGoName(f.Name())
			bfBits := f.ValueBits()
			bfOffset := f.OffsetBits()
			bfType := MapCTypeToGo(ft)

			if currentBfGroup == nil || fieldOffset != currentBfGroup.offset {
				if currentBfGroup != nil {
					padding := currentBfGroup.offset - currentOffset
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
					for _, bf := range currentBfGroup.bitfields {
						methods = append(methods, GenerateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
					}
					currentOffset = currentBfGroup.offset + currentBfGroup.size
				}
				totalBits := ft.Size() * 8
				backingType := "uint8"
				switch totalBits {
				case 8:
					backingType = "uint8"
				case 16:
					backingType = "uint16"
				case 32:
					backingType = "uint32"
				case 64:
					backingType = "uint64"
				}
				currentBfGroup = &bitfieldGroup{
					offset:      fieldOffset,
					size:        ft.Size(),
					backingName: bfName + "Bits",
					backingType: backingType,
				}
			}
			currentBfGroup.bitfields = append(currentBfGroup.bitfields, BitfieldInfo{
				goName:    bfName,
				goType:    bfType,
				bitOffset: bfOffset,
				bitWidth:  bfBits,
			})
			continue
		}

		if currentBfGroup != nil {
			padding := currentBfGroup.offset - currentOffset
			if !isPacked && padding > 0 {
				fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
			}
			fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
			for _, bf := range currentBfGroup.bitfields {
				methods = append(methods, GenerateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
			}
			currentOffset = currentBfGroup.offset + currentBfGroup.size
			currentBfGroup = nil
		}

		padding = fieldOffset - currentOffset

		fieldType := MapCTypeToGo(ft)
		fieldName := CFieldNameToGoName(f.Name())

		if fieldName == "_" && f.Name() == "" {
			switch v := ft.(type) {
			case *cc.UnionType:
				backingName, backingType, bfs := ExtractUnionBitfields(v)
				if backingName != "" && len(bfs) > 0 {
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", backingName, backingType))
					for _, bf := range bfs {
						methods = append(methods, GenerateBitfieldGetterSetter(bf, structGoName, backingName, backingType))
					}
					currentOffset = fieldOffset + ft.Size()
					continue
				}
				if unionFields := GenerateUnionFields(v, v.Size()); unionFields != "" {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					innerTypes = append(innerTypes, StructInfo{
						goName: anonUnionName + "_",
						cName:  anonUnionName,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, StructInfo{
						goName: anonUnionName,
						cName:  anonUnionName,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
					})
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, anonUnionName))
					currentOffset = fieldOffset + ft.Size()
					continue
				}
				if v.NumFields() > 0 {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					var unionFieldStrs []string
					for i := 0; i < v.NumFields(); i++ {
						f := v.FieldByIndex(i)
						ft := f.Type()
						fieldType := MapCTypeToGo(ft)
						fieldName := CFieldNameToGoName(f.Name())
						if fieldName == "_" || fieldType == "" {
							continue
						}
						unionFieldStrs = append(unionFieldStrs, fmt.Sprintf("\t%s %s", fieldName, fieldType))
						switch vt := ft.(type) {
						case *cc.StructType:
							tag := vt.Tag()
							tagStr := string(tag.Src())
							nestedFields, nestedMethods, nestedInner := GenerateStructFields(vt, fieldType, isPacked)
							innerTypes = append(innerTypes, StructInfo{
								goName:  fieldType,
								cName:   tagStr,
								fields:  nestedFields,
								methods: nestedMethods,
							})
							innerTypes = append(innerTypes, nestedInner...)
						case *cc.UnionType:
							tag := vt.Tag()
							tagStr := string(tag.Src())
							if tagStr != "" {
								if uf := GenerateUnionFields(vt, vt.Size()); uf != "" {
									innerTypes = append(innerTypes, StructInfo{
										goName: fieldType + "_",
										cName:  tagStr,
										fields: uf,
									})
									innerTypes = append(innerTypes, StructInfo{
										goName: fieldType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(vt, vt.Size())),
									})
								} else {
									innerTypes = append(innerTypes, StructInfo{
										goName: fieldType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(vt, vt.Size())),
									})
								}
							}
						}
					}
					if len(unionFieldStrs) > 0 {
						accessorName := anonUnionName + "_"
						innerTypes = append(innerTypes, StructInfo{
							goName: accessorName,
							cName:  anonUnionName,
							fields: strings.Join(unionFieldStrs, "\n") + "\n",
						})
						innerTypes = append(innerTypes, StructInfo{
							goName: anonUnionName,
							cName:  anonUnionName,
							fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
						})
						if !isPacked && padding > 0 {
							fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
						}
						fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, anonUnionName))
						currentOffset = fieldOffset + v.Size()
						continue
					}
				}
			case *cc.StructType:
				if HasBitfields(v) {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					bitfieldGoName := fieldType
					bitfieldFields, bitfieldMethods := GeneratePureBitfieldStruct(v, bitfieldGoName)
					innerTypes = append(innerTypes, StructInfo{
						goName:  bitfieldGoName,
						cName:   bitfieldGoName,
						fields:  bitfieldFields,
						methods: bitfieldMethods,
					})
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, bitfieldGoName))
					currentOffset = fieldOffset + ft.Size()
					continue
				}
			}
		}

		if !isPacked && padding > 0 {
			fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
		}

		if v, ok := ft.(*cc.StructType); ok {
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr != "" && HasBitfields(v) {
				bitfieldGoName := fieldType
				bitfieldFields, bitfieldMethods := GeneratePureBitfieldStruct(v, bitfieldGoName)
				innerTypes = append(innerTypes, StructInfo{
					goName:  bitfieldGoName,
					cName:   bitfieldGoName,
					fields:  bitfieldFields,
					methods: bitfieldMethods,
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, bitfieldGoName))
			} else if tagStr != "" {
				nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, fieldType, isPacked)
				innerTypes = append(innerTypes, StructInfo{
					goName:  fieldType,
					cName:   fieldType,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				innerTypes = append(innerTypes, nestedInner...)
				for _, it := range GenerateStructFieldsInnerTypes(v, fieldType, isPacked) {
					innerTypes = append(innerTypes, it)
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			} else if HasBitfields(v) {
				bitfieldGoName := fieldType
				bitfieldFields, bitfieldMethods := GeneratePureBitfieldStruct(v, bitfieldGoName)
				innerTypes = append(innerTypes, StructInfo{
					goName:  bitfieldGoName,
					cName:   bitfieldGoName,
					fields:  bitfieldFields,
					methods: bitfieldMethods,
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, bitfieldGoName))
			} else {
				anonCount++
				anonGoName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount)
				nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, anonGoName, isPacked)
				innerTypes = append(innerTypes, StructInfo{
					goName:  anonGoName,
					cName:   anonGoName,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				innerTypes = append(innerTypes, nestedInner...)
				for _, it := range GenerateStructFieldsInnerTypes(v, anonGoName, isPacked) {
					innerTypes = append(innerTypes, it)
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, anonGoName))
			}
		} else if uv, ok := ft.(*cc.UnionType); ok {
			tag := uv.Tag()
			tagStr := string(tag.Src())
			unionSize := uv.Size()
			if fieldType != "uint64" {
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			} else if tagStr == "" {
				anonCount++
				unionGoName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
				if unionFields := GenerateUnionFields(uv, unionSize); unionFields != "" {
					for j := 0; j < uv.NumFields(); j++ {
						if uf := uv.FieldByIndex(j); uf != nil {
							if usv, ok := uf.Type().(*cc.StructType); ok {
								nestedTag := usv.Tag()
								nestedTagStr := string(nestedTag.Src())
								if nestedTagStr != "" && !HasBitfields(usv) {
									nestedF, nestedM, _ := GenerateStructFields(usv, MapCTypeToGo(uf.Type()), isPacked)
									innerTypes = append(innerTypes, StructInfo{
										goName:  MapCTypeToGo(uf.Type()),
										cName:   nestedTagStr,
										fields:  nestedF,
										methods: nestedM,
									})
								}
							}
						}
					}
					innerTypes = append(innerTypes, StructInfo{
						goName: unionGoName + "_",
						cName:  tagStr,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, StructInfo{
						goName: unionGoName,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
					})
				} else {
					innerTypes = append(innerTypes, StructInfo{
						goName: unionGoName,
						cName:  unionGoName,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
					})
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			} else if unionFields := GenerateUnionFields(uv, unionSize); unionFields != "" {
				unionGoName := fieldType
				for j := 0; j < uv.NumFields(); j++ {
					if uf := uv.FieldByIndex(j); uf != nil {
						if usv, ok := uf.Type().(*cc.StructType); ok {
							nestedTag := usv.Tag()
							nestedTagStr := string(nestedTag.Src())
							if nestedTagStr != "" && !HasBitfields(usv) {
								nestedF, nestedM, _ := GenerateStructFields(usv, MapCTypeToGo(uf.Type()), isPacked)
								innerTypes = append(innerTypes, StructInfo{
									goName:  MapCTypeToGo(uf.Type()),
									cName:   nestedTagStr,
									fields:  nestedF,
									methods: nestedM,
								})
							}
						}
					}
				}
				innerTypes = append(innerTypes, StructInfo{
					goName: unionGoName + "_",
					cName:  tagStr,
					fields: unionFields,
				})
				innerTypes = append(innerTypes, StructInfo{
					goName: unionGoName,
					cName:  tagStr,
					fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			} else {
				unionGoName := fieldType
				innerTypes = append(innerTypes, StructInfo{
					goName: unionGoName,
					cName:  tagStr,
					fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			}
		} else if arr, ok := ft.(*cc.ArrayType); ok {
			elem := arr.Elem()
			for innerArr, ok := elem.(*cc.ArrayType); ok; innerArr, ok = elem.(*cc.ArrayType) {
				elem = innerArr.Elem()
			}
			if sv, ok := elem.(*cc.StructType); ok {
				tag := sv.Tag()
				tagStr := string(tag.Src())
				if tagStr != "" {
					elemGoName := CTagToGoName(tagStr)
					nestedFields, nestedMethods, nestedInner := GenerateStructFields(sv, elemGoName, isPacked)
					innerTypes = append(innerTypes, StructInfo{
						goName:  elemGoName,
						cName:   tagStr,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
					fieldType = elemGoName
				} else {
					anonCount++
					anonElemName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Elem"
					nestedFields, nestedMethods, nestedInner := GenerateStructFields(sv, anonElemName, isPacked)
					innerTypes = append(innerTypes, StructInfo{
						goName:  anonElemName,
						cName:   anonElemName,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
					fieldType = anonElemName
				}
			} else if uv, ok := elem.(*cc.UnionType); ok {
				tag := uv.Tag()
				tagStr := string(tag.Src())
				unionSize := uv.Size()
				if unionFields := GenerateUnionFields(uv, unionSize); unionFields != "" {
					elemGoName := CTagToGoName(tagStr)
					if elemGoName == "" {
						elemGoName = fieldType
					}
					innerTypes = append(innerTypes, StructInfo{
						goName: elemGoName,
						cName:  tagStr,
						fields: unionFields,
					})
					fieldType = elemGoName
				} else if tagStr != "" {
					elemGoName := CTagToGoName(tagStr)
					innerTypes = append(innerTypes, StructInfo{
						goName: elemGoName,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
					})
					fieldType = elemGoName
				} else {
					anonCount++
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					innerTypes = append(innerTypes, StructInfo{
						goName: anonUnionName,
						cName:  anonUnionName,
						fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(uv, unionSize)),
					})
					fieldType = anonUnionName
				}
			} else {
				fieldType = MapCTypeToGo(arr.Elem())
			}
			fields = append(fields, fmt.Sprintf("\t%s [%d]%s", fieldName, arr.Len(), fieldType))
		} else {
			fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
		}
		currentOffset = fieldOffset + ft.Size()
	}

	if currentBfGroup != nil {
		padding := currentBfGroup.offset - currentOffset
		if !isPacked && padding > 0 {
			fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
		}
		fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
		for _, bf := range currentBfGroup.bitfields {
			methods = append(methods, GenerateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
		}
		currentOffset = currentBfGroup.offset + currentBfGroup.size
	}

	if len(anonUnionCandidates) > 0 {
		if !flushAnonUnion() {
			for _, af := range anonUnionCandidates {
				fieldPadding := af.offset - currentOffset
				currentOffset = af.offset + af.size
				if !isPacked && fieldPadding > 0 {
					fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
				}
				switch v := af.ccType.(type) {
				case *cc.StructType:
					tag := v.Tag()
					tagStr := string(tag.Src())
					nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, af.goType, isPacked)
					innerTypes = append(innerTypes, StructInfo{
						goName:  af.goType,
						cName:   tagStr,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
				case *cc.UnionType:
					tag := v.Tag()
					tagStr := string(tag.Src())
					if tagStr != "" {
						if uf := GenerateUnionFields(v, v.Size()); uf != "" {
							innerTypes = append(innerTypes, StructInfo{
								goName: af.goType + "_",
								cName:  tagStr,
								fields: uf,
							})
							innerTypes = append(innerTypes, StructInfo{
								goName: af.goType,
								cName:  tagStr,
								fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
							})
						} else {
							innerTypes = append(innerTypes, StructInfo{
								goName: af.goType,
								cName:  tagStr,
								fields: fmt.Sprintf("\t%s\n", UnionAlignedFFIType(v, v.Size())),
							})
						}
					}
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
			}
		}
	}

	totalSize := t.Size()
	trailingPadding := totalSize - currentOffset
	if !isPacked && trailingPadding > 0 {
		fields = append(fields, fmt.Sprintf("\t_  [%d]byte", trailingPadding))
	}

	fieldNameCount := make(map[string]int)
	for i, f := range fields {
		parts := strings.SplitN(strings.TrimSpace(f), " ", 2)
		if len(parts) < 2 {
			continue
		}
		name := parts[0]
		if name == "_" {
			continue
		}
		fieldNameCount[name]++
		if fieldNameCount[name] > 1 {
			newName := fmt.Sprintf("%s%d", name, fieldNameCount[name])
			fields[i] = strings.Replace(f, name, newName, 1)
		}
	}

	return strings.Join(fields, "\n") + "\n", strings.Join(methods, "\n"), innerTypes
}

func GenerateStructFieldsInnerTypes(t *cc.StructType, structGoName string, forcePacked bool) []StructInfo {
	var result []StructInfo
	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		ft := f.Type()
		if v, ok := ft.(*cc.StructType); ok {
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr != "" {
				nestedFields, nestedMethods, nestedInner := GenerateStructFields(v, tagStr, forcePacked)
				result = append(result, StructInfo{
					goName:  tagStr,
					cName:   tagStr,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				result = append(result, nestedInner...)
				result = append(result, GenerateStructFieldsInnerTypes(v, tagStr, forcePacked)...)
			}
		}
	}
	return result
}

func HasBitfields(t *cc.StructType) bool {
	for i := 0; i < t.NumFields(); i++ {
		if t.FieldByIndex(i).IsBitfield() {
			return true
		}
	}
	return false
}

func IsPureBitfieldStruct(t *cc.StructType) bool {
	if t.NumFields() == 0 {
		return false
	}
	for i := 0; i < t.NumFields(); i++ {
		if !t.FieldByIndex(i).IsBitfield() {
			return false
		}
	}
	return true
}

func GeneratePureBitfieldStruct(t *cc.StructType, structGoName string) (string, string) {
	totalBytes := t.Size()

	var backingType string
	switch totalBytes {
	case 1:
		backingType = "uint8"
	case 2:
		backingType = "uint16"
	case 4:
		backingType = "uint32"
	case 8:
		backingType = "uint64"
	default:
		if totalBytes <= 16 {
			backingType = "uint64"
		} else {
			backingType = fmt.Sprintf("[%d]byte", totalBytes)
		}
	}

	bfs := ExtractStructBitfields(t)
	var methods []string
	for _, bf := range bfs {
		methods = append(methods, GenerateBitfieldGetterSetter(bf, structGoName, "Value", backingType))
	}

	fieldStr := fmt.Sprintf("\tValue %s\n", backingType)
	return fieldStr, strings.Join(methods, "\n")
}

func ExtractUnionBitfields(u *cc.UnionType) (backingName string, backingType string, bfs []BitfieldInfo) {
	backingIdx := -1
	bitfieldIdx := -1
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		switch v := ft.(type) {
		case *cc.StructType:
			if HasBitfields(v) {
				bitfieldIdx = i
			}
		default:
			if !f.IsBitfield() && backingIdx == -1 {
				backingIdx = i
			}
		}
	}
	if backingIdx == -1 || bitfieldIdx == -1 {
		return "", "", nil
	}
	bf := u.FieldByIndex(backingIdx)
	backingName = CFieldNameToGoName(bf.Name())
	backingType = MapCTypeToGo(bf.Type())
	bs := u.FieldByIndex(bitfieldIdx).Type().(*cc.StructType)
	bfs = ExtractStructBitfields(bs)
	return backingName, backingType, bfs
}

func ExtractStructBitfields(t *cc.StructType) []BitfieldInfo {
	var bfs []BitfieldInfo
	bitPos := 0
	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		if !f.IsBitfield() {
			continue
		}
		bfs = append(bfs, BitfieldInfo{
			goName:    CFieldNameToGoName(f.Name()),
			goType:    MapCTypeToGo(f.Type()),
			bitOffset: bitPos,
			bitWidth:  f.ValueBits(),
		})
		bitPos += int(f.ValueBits())
	}
	return bfs
}

func GenerateBitfieldGetterSetter(bf BitfieldInfo, structName, backingName, backingType string) string {
	mask := uint64(1)<<bf.bitWidth - 1
	shift := bf.bitOffset
	goName := bf.goName
	receiver := strings.ToLower(structName[:1])

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nfunc (%s *%s) Get%s() %s {\n", receiver, structName, goName, bf.goType))
	if shift == 0 {
		sb.WriteString(fmt.Sprintf("\treturn %s(%s.%s & %s(0x%X))\n", bf.goType, receiver, backingName, backingType, mask))
	} else {
		sb.WriteString(fmt.Sprintf("\treturn %s((%s.%s >> %d) & %s(0x%X))\n", bf.goType, receiver, backingName, shift, backingType, mask))
	}
	sb.WriteString("}\n")

	sb.WriteString(fmt.Sprintf("\nfunc (%s *%s) Set%s(val %s) {\n", receiver, structName, goName, bf.goType))
	if shift == 0 {
		sb.WriteString(fmt.Sprintf("\t%s.%s = (%s.%s & ^%s(0x%X)) | %s(%s(val) & 0x%X)\n", receiver, backingName, receiver, backingName, backingType, mask, backingType, bf.goType, mask))
	} else {
		sb.WriteString(fmt.Sprintf("\t%s.%s = (%s.%s & ^%s(0x%X<<%d)) | (%s(%s(val) & 0x%X) << %d)\n", receiver, backingName, receiver, backingName, backingType, mask, shift, backingType, bf.goType, mask, shift))
	}
	sb.WriteString("}\n")

	return sb.String()
}

