// Automatically generated code - DO NOT EDIT.
package main

import (
	"fmt"
)

type ArrayElem struct {
	Type__     string      `json:"__type__"`
	Name       string      `json:"name"`
	Asname     interface{} `json:"asname"`
	Lineno     int         `json:"lineno"`
	Col_offset int         `json:"col_offset"`
}

type ArrayElem1 struct {
	Type__     string      `json:"__type__"`
	Names      []ArrayElem `json:"names"`
	Lineno     int         `json:"lineno"`
	Col_offset int         `json:"col_offset"`
}

type ConditionGroupValue struct {
	Id          string        `json:"id"`
	Conjunction string        `json:"conjunction"`
	Not         bool          `json:"not"`
	Children    []interface{} `json:"children"`
	IfVal       string        `json:"if"`
}

type QRCodeImageSettings struct {
	Src      string  `json:"src"`
	Height   float64 `json:"height"`
	Width    float64 `json:"width"`
	Excavate bool    `json:"excavate"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type RootConfig struct {
	Type__       string        `json:"__type__"`
	Body         []ArrayElem1  `json:"body"`
	Type_ignores []interface{} `json:"type_ignores"`
}

type RootConfig1 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties   `json:"patternProperties"`
	ID                   string              `json:"$$id"`
	ClassName            className           `json:"className"`
	Disabled             bool                `json:"disabled"`
	DisabledOn           className           `json:"disabledOn"`
	QRCodeImageSettings  QRCodeImageSettings `json:"QRCodeImageSettings"`
	ConditionGroupValue  ConditionGroupValue `json:"ConditionGroupValue"`
}

type className struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}

type patternProperties struct {
	PatternProperties interface{} `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|size|label|labelAlign|labelWidth|labelClassName|name|extraName|remark|labelRemark|hint|submitOnChange|readOnly|readOnlyOn|validateOnChange|description|desc|descriptionClassName|mode|horizontal|inline|inputClassName|placeholder|required|validationErrors|validations|value|clearValueOnHidden|validateApi|autoFill|initAutoFill|row|type|extractValue|joinValues|delimiter|allowCity|allowDistrict|allowStreet|searchable|itemClassName|columnClassName|columnRatio)$"`
}

type RootContainer struct {
	Data *RootConfig1 `json:"data"`
}

func main() {
	data := RootConfig1{
		Ref:                  "#/definitions/InputCityControlSchema",
		AdditionalProperties: false,
		PatternProperties:    patternProperties{},
		ID:                   "{\r\n    \"type\": \"string\",\r\n    \"description\": \"component unique id\"\r\n  }",
		ClassName: className{
			Ref:         "#/definitions/SchemaClassName",
			Description: "container className",
		},
		Disabled: false,
		DisabledOn: className{
			Ref:         "#/definitions/SchemaExpression",
			Description: "disabled expression",
		},
		QRCodeImageSettings: QRCodeImageSettings{},
		ConditionGroupValue: ConditionGroupValue{},
	}
	root := RootContainer{Data: &data}
	fmt.Printf("%#v\\n", root)
}
