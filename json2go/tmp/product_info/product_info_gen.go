// Automatically generated code - DO NOT EDIT.
package main

import (
	"fmt"
)

type ActionSchema struct {
	AllOf []ArrayElem5 `json:"allOf"`
}

type AjaxActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Api                  className       `json:"api"`
	Feedback             RootSchema      `json:"feedback"`
	Reload               RootSchema      `json:"reload"`
	Redirect             string          `json:"redirect"`
	IgnoreConfirm        bool            `json:"ignoreConfirm"`
	IsolateScope         bool            `json:"isolateScope"`
}

type AlertSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Title                string        `json:"title"`
	Body                 className     `json:"body"`
	Level                string        `json:"level"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	CloseButtonClassName string        `json:"closeButtonClassName"`
	ShowIcon             bool          `json:"showIcon"`
	Icon                 className     `json:"icon"`
	IconClassName        string        `json:"iconClassName"`
	Actions              className     `json:"actions"`
}

type AnchorNavSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Links                []any         `json:"links"`
	Active               any           `json:"active"`
	LinkClassName        className     `json:"linkClassName"`
	SectionClassName     className     `json:"sectionClassName"`
	Direction            string        `json:"direction"`
}

type AnchorNavSectionSchema struct {
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	Title                string        `json:"title"`
	Href                 string        `json:"href"`
	Body                 className     `json:"body"`
	Children             []any         `json:"children"`
}

type ArrayControlSchema struct {
	Scaffold                           mock             `json:"scaffold"`
	NoBorder                           bool             `json:"noBorder"`
	DeleteConfirmText                  string           `json:"deleteConfirmText"`
	DeleteApi                          className        `json:"deleteApi"`
	TypeSwitchable                     bool             `json:"typeSwitchable"`
	FormClassName                      className        `json:"formClassName"`
	AddButtonClassName                 className        `json:"addButtonClassName"`
	AddButtonText                      string           `json:"addButtonText"`
	Addable                            bool             `json:"addable"`
	Addattop                           bool             `json:"addattop"`
	Draggable                          bool             `json:"draggable"`
	DraggableTip                       string           `json:"draggableTip"`
	Flat                               bool             `json:"flat"`
	Delimiter                          string           `json:"delimiter"`
	JoinValues                         bool             `json:"joinValues"`
	MaxLength                          SchemaClassName  `json:"maxLength"`
	MinLength                          SchemaClassName  `json:"minLength"`
	MultiLine                          bool             `json:"multiLine"`
	Multiple                           bool             `json:"multiple"`
	Removable                          bool             `json:"removable"`
	SubFormMode                        string           `json:"subFormMode"`
	SubFormHorizontal                  className        `json:"subFormHorizontal"`
	Placeholder                        string           `json:"placeholder"`
	CanAccessSuperData                 bool             `json:"canAccessSuperData"`
	TabsMode                           bool             `json:"tabsMode"`
	TabsStyle                          string           `json:"tabsStyle"`
	TabsLabelTpl                       className        `json:"tabsLabelTpl"`
	LazyLoad                           bool             `json:"lazyLoad"`
	StrictMode                         bool             `json:"strictMode"`
	SyncFields                         []string         `json:"syncFields"`
	Nullable                           bool             `json:"nullable"`
	Messages                           messages2        `json:"messages"`
	UpdatePristineAfterStoreDataReInit bool             `json:"updatePristineAfterStoreDataReInit"`
	TestIdBuilder                      RootSchema       `json:"testIdBuilder"`
	Remark                             className        `json:"remark"`
	LabelRemark                        className        `json:"labelRemark"`
	Size                               string           `json:"size"`
	Label                              SchemaClassName  `json:"label"`
	LabelAlign                         className        `json:"labelAlign"`
	LabelWidth                         any              `json:"labelWidth"`
	LabelClassName                     string           `json:"labelClassName"`
	Name                               string           `json:"name"`
	ExtraName                          string           `json:"extraName"`
	Hint                               string           `json:"hint"`
	SubmitOnChange                     bool             `json:"submitOnChange"`
	ReadOnly                           bool             `json:"readOnly"`
	ReadOnlyOn                         string           `json:"readOnlyOn"`
	ValidateOnChange                   bool             `json:"validateOnChange"`
	Description                        string           `json:"description"`
	Desc                               string           `json:"desc"`
	DescriptionClassName               className        `json:"descriptionClassName"`
	Mode                               string           `json:"mode"`
	Horizontal                         className        `json:"horizontal"`
	Inline                             bool             `json:"inline"`
	InputClassName                     className        `json:"inputClassName"`
	Required                           bool             `json:"required"`
	ValidationErrors                   validationErrors `json:"validationErrors"`
	Validations                        data             `json:"validations"`
	Value                              mock             `json:"value"`
	ClearValueOnHidden                 bool             `json:"clearValueOnHidden"`
	ValidateApi                        SchemaClassName  `json:"validateApi"`
	AutoFill                           autoFill         `json:"autoFill"`
	InitAutoFill                       data             `json:"initAutoFill"`
	Row                                float64          `json:"row"`
	ID                                 string           `json:"$$id"`
	ClassName                          className        `json:"className"`
	Ref                                string           `json:"$ref"`
	Disabled                           bool             `json:"disabled"`
	DisabledOn                         className        `json:"disabledOn"`
	Hidden                             bool             `json:"hidden"`
	HiddenOn                           className        `json:"hiddenOn"`
	Visible                            bool             `json:"visible"`
	VisibleOn                          className        `json:"visibleOn"`
	Id                                 string           `json:"id"`
	OnEvent                            any              `json:"onEvent"`
	Static                             bool             `json:"static"`
	StaticOn                           className        `json:"staticOn"`
	StaticPlaceholder                  string           `json:"staticPlaceholder"`
	StaticClassName                    className        `json:"staticClassName"`
	StaticLabelClassName               className        `json:"staticLabelClassName"`
	StaticInputClassName               className        `json:"staticInputClassName"`
	StaticSchema                       staticSchema     `json:"staticSchema"`
	Style                              any              `json:"style"`
	EditorSetting                      editorSetting    `json:"editorSetting"`
	UseMobileUI                        bool             `json:"useMobileUI"`
	TypeVal                            string           `json:"type"`
	Items                              className        `json:"items"`
}

type ArrayElem struct {
	Type__     string `json:"__type__"`
	Name       string `json:"name"`
	Asname     any    `json:"asname"`
	Lineno     int    `json:"lineno"`
	Col_offset int    `json:"col_offset"`
}

type ArrayElem1 struct {
	Type__     string      `json:"__type__"`
	Names      []ArrayElem `json:"names"`
	Lineno     int         `json:"lineno"`
	Col_offset int         `json:"col_offset"`
}

type ArrayElem10 struct {
	TypeVal              string        `json:"type"`
	Header               header        `json:"header"`
	Body                 []any         `json:"body"`
	Media                media         `json:"media"`
	Actions              []any         `json:"actions"`
	Toolbar              []any         `json:"toolbar"`
	Secondary            className     `json:"secondary"`
	UseCardLabel         bool          `json:"useCardLabel"`
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
}

type ArrayElem11 struct {
	TypeVal string `json:"type"`
	Items   any    `json:"items"`
}

type ArrayElem12 struct {
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Top    float64 `json:"top"`
	Bottom float64 `json:"bottom"`
}

type ArrayElem13 struct {
	Enable bool     `json:"enable"`
	Types  []string `json:"types"`
}

type ArrayElem14 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties3 `json:"patternProperties"`
}

type ArrayElem15 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties4 `json:"patternProperties"`
}

type ArrayElem16 struct {
	Children             []any           `json:"children"`
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Api                  className       `json:"api"`
	Feedback             RootSchema      `json:"feedback"`
	Reload               RootSchema      `json:"reload"`
	Redirect             string          `json:"redirect"`
	IgnoreConfirm        bool            `json:"ignoreConfirm"`
	IsolateScope         bool            `json:"isolateScope"`
}

type ArrayElem17 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties5 `json:"patternProperties"`
}

type ArrayElem18 struct {
	Value RootSchema `json:"value"`
	Color string     `json:"color"`
}

type ArrayElem19 struct {
	TypeVal string `json:"type"`
	Items   string `json:"items"`
}

type ArrayElem2 struct {
	TypeVal string `json:"type"`
}

type ArrayElem20 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties6 `json:"patternProperties"`
}

type ArrayElem21 struct {
	TypeVal              string `json:"type"`
	AdditionalProperties string `json:"additionalProperties"`
}

type ArrayElem22 struct {
	TypeVal string     `json:"type"`
	Items   RootSchema `json:"items"`
}

type ArrayElem23 struct {
	Color string `json:"color"`
	Title string `json:"title"`
}

type ArrayElem24 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties7 `json:"patternProperties"`
}

type ArrayElem25 struct {
	TypeVal      string       `json:"type"`
	Label        string       `json:"label"`
	ValueTypes   []string     `json:"valueTypes"`
	Operators    []any        `json:"operators"`
	Funcs        []string     `json:"funcs"`
	DefaultValue staticSchema `json:"defaultValue"`
	Placeholder  string       `json:"placeholder"`
	Name         string       `json:"name"`
	MinLength    float64      `json:"minLength"`
	MaxLength    float64      `json:"maxLength"`
}

type ArrayElem26 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ArrayElem27 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties8 `json:"patternProperties"`
}

type ArrayElem28 struct {
	ID                   string          `json:"$$id"`
	ClassName            string          `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 string          `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Api                  className       `json:"api"`
	Feedback             RootSchema      `json:"feedback"`
	Reload               RootSchema      `json:"reload"`
	Redirect             string          `json:"redirect"`
	IgnoreConfirm        bool            `json:"ignoreConfirm"`
	IsolateScope         bool            `json:"isolateScope"`
	Position             string          `json:"position"`
}

type ArrayElem29 struct {
	Ref                  string             `json:"$ref"`
	AdditionalProperties bool               `json:"additionalProperties"`
	PatternProperties    patternProperties9 `json:"patternProperties"`
}

type ArrayElem3 struct {
	TypeVal  string `json:"type"`
	ConstVal string `json:"const"`
}

type ArrayElem30 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties10 `json:"patternProperties"`
}

type ArrayElem31 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties11 `json:"patternProperties"`
}

type ArrayElem32 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties12 `json:"patternProperties"`
}

type ArrayElem33 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties13 `json:"patternProperties"`
}

type ArrayElem34 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties14 `json:"patternProperties"`
}

type ArrayElem35 struct {
	Ref                  string              `json:"$ref"`
	AdditionalProperties bool                `json:"additionalProperties"`
	PatternProperties    patternProperties15 `json:"patternProperties"`
}

type ArrayElem36 struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Short_name  string `json:"short_name"`
	Plural_name string `json:"plural_name"`
	Icon        icon   `json:"icon"`
}

type ArrayElem37 struct {
	Close string `json:"close"`
	Day   int    `json:"day"`
	Open  string `json:"open"`
}

type ArrayElem38 struct {
	Current                  int    `json:"current"`
	Date                     string `json:"date"`
	Iso_currency_code        string `json:"iso_currency_code"`
	Unofficial_currency_code any    `json:"unofficial_currency_code"`
}

type ArrayElem39 struct {
	Data    data1 `json:"data"`
	Primary bool  `json:"primary"`
}

type ArrayElem4 struct {
	IfVal ArrayElem2 `json:"if"`
	Then  RootSchema `json:"then"`
}

type ArrayElem40 struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	TypeVal string `json:"type"`
}

type ArrayElem41 struct {
	Addresses     []ArrayElem39 `json:"addresses"`
	Emails        []ArrayElem40 `json:"emails"`
	Names         []string      `json:"names"`
	Phone_numbers []ArrayElem40 `json:"phone_numbers"`
}

type ArrayElem42 struct {
	Account_id               string `json:"account_id"`
	Amount                   int    `json:"amount"`
	Date                     string `json:"date"`
	Iso_currency_code        string `json:"iso_currency_code"`
	Original_description     string `json:"original_description"`
	Pending                  bool   `json:"pending"`
	Transaction_id           string `json:"transaction_id"`
	Unofficial_currency_code any    `json:"unofficial_currency_code"`
}

type ArrayElem43 struct {
	Account_id          string        `json:"account_id"`
	Balances            balances      `json:"balances"`
	Days_available      int           `json:"days_available"`
	Historical_balances []ArrayElem38 `json:"historical_balances"`
	Mask                string        `json:"mask"`
	Name                string        `json:"name"`
	Official_name       string        `json:"official_name"`
	Owners              []ArrayElem41 `json:"owners"`
	Ownership_type      any           `json:"ownership_type"`
	Subtype             string        `json:"subtype"`
	Transactions        []ArrayElem42 `json:"transactions"`
	TypeVal             string        `json:"type"`
}

type ArrayElem44 struct {
	Accounts          []ArrayElem43 `json:"accounts"`
	Date_last_updated string        `json:"date_last_updated"`
	Institution_id    string        `json:"institution_id"`
	Institution_name  string        `json:"institution_name"`
	Item_id           string        `json:"item_id"`
}

type ArrayElem45 struct {
	Commands               []string `json:"commands"`
	VmOptionsFilePath      string   `json:"vmOptionsFilePath"`
	BootClassPathJarNames  []string `json:"bootClassPathJarNames"`
	AdditionalJvmArguments []string `json:"additionalJvmArguments"`
	MainClass              string   `json:"mainClass"`
	EnvVarBaseName         string   `json:"envVarBaseName"`
	DataDirectoryName      string   `json:"dataDirectoryName"`
}

type ArrayElem46 struct {
	Os                     string        `json:"os"`
	Arch                   string        `json:"arch"`
	LauncherPath           string        `json:"launcherPath"`
	JavaExecutablePath     string        `json:"javaExecutablePath"`
	VmOptionsFilePath      string        `json:"vmOptionsFilePath"`
	BootClassPathJarNames  []string      `json:"bootClassPathJarNames"`
	AdditionalJvmArguments []string      `json:"additionalJvmArguments"`
	MainClass              string        `json:"mainClass"`
	CustomCommands         []ArrayElem45 `json:"customCommands"`
}

type ArrayElem47 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ArrayElem48 struct {
	Name      string   `json:"name"`
	Kind      string   `json:"kind"`
	ClassPath []string `json:"classPath"`
}

type ArrayElem5 struct {
	IfVal ifVal      `json:"if"`
	Then  RootSchema `json:"then"`
}

type ArrayElem6 struct {
	TypeVal string `json:"type"`
	Items   items  `json:"items"`
}

type ArrayElem7 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties1 `json:"patternProperties"`
}

type ArrayElem8 struct {
	Ref               string             `json:"$ref"`
	PatternProperties patternProperties2 `json:"patternProperties"`
}

type ArrayElem9 struct {
	AllOf []ArrayElem8 `json:"allOf"`
}

type ArrayItem struct {
	Title       RootSchema `json:"title"`
	Body        RootSchema `json:"body"`
	Level       string     `json:"level"`
	Id          string     `json:"id"`
	Position    string     `json:"position"`
	CloseButton bool       `json:"closeButton"`
	ShowIcon    bool       `json:"showIcon"`
	Timeout     float64    `json:"timeout"`
}

type ArrayItem1 struct {
	Length float64 `json:"length"`
}

type ArrayItem2 struct {
	Key    string  `json:"key"`
	Label  string  `json:"label"`
	Remark string  `json:"remark"`
	Status float64 `json:"status"`
}

type ArrayItem3 struct {
	Rule    string `json:"rule"`
	Message string `json:"message"`
	Name    data   `json:"name"`
}

type ArrayItem4 struct {
	Label string `json:"label"`
}

type AudioSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Inline               bool          `json:"inline"`
	Src                  className     `json:"src"`
	Loop                 bool          `json:"loop"`
	AutoPlay             bool          `json:"autoPlay"`
	Rates                []float64     `json:"rates"`
	Controls             []string      `json:"controls"`
}

type AutoFillHeightObject struct {
	Height    float64 `json:"height"`
	MaxHeight float64 `json:"maxHeight"`
}

type AutoGenerateFilterObject struct {
	ColumnsNum       float64 `json:"columnsNum"`
	ShowBtnToolbar   bool    `json:"showBtnToolbar"`
	DefaultCollapsed bool    `json:"defaultCollapsed"`
}

type AvatarSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Badge                className       `json:"badge"`
	Src                  string          `json:"src"`
	DefaultAvatar        string          `json:"defaultAvatar"`
	Icon                 string          `json:"icon"`
	Fit                  string          `json:"fit"`
	Shape                string          `json:"shape"`
	Size                 SchemaClassName `json:"size"`
	Text                 string          `json:"text"`
	Gap                  float64         `json:"gap"`
	Alt                  string          `json:"alt"`
	Draggable            bool            `json:"draggable"`
	CrossOrigin          string          `json:"crossOrigin"`
	OnError              string          `json:"onError"`
}

type BadgeObject struct {
	ClassName     string  `json:"className"`
	Text          any     `json:"text"`
	Size          float64 `json:"size"`
	Mode          string  `json:"mode"`
	Offset        []any   `json:"offset"`
	Position      string  `json:"position"`
	OverflowCount float64 `json:"overflowCount"`
	VisibleOn     string  `json:"visibleOn"`
	Animation     bool    `json:"animation"`
	Style         any     `json:"style"`
	Level         string  `json:"level"`
}

type BaseApiObject struct {
	Method                 string          `json:"method"`
	Url                    string          `json:"url"`
	Data                   any             `json:"data"`
	ConvertKeyToPath       bool            `json:"convertKeyToPath"`
	ResponseData           any             `json:"responseData"`
	AttachDataToQuery      bool            `json:"attachDataToQuery"`
	DataType               string          `json:"dataType"`
	ResponseType           string          `json:"responseType"`
	Headers                any             `json:"headers"`
	SendOn                 string          `json:"sendOn"`
	ReplaceData            bool            `json:"replaceData"`
	ConcatDataFields       SchemaClassName `json:"concatDataFields"`
	AutoRefresh            bool            `json:"autoRefresh"`
	TrackExpression        string          `json:"trackExpression"`
	Cache                  float64         `json:"cache"`
	ForceAppendDataToQuery bool            `json:"forceAppendDataToQuery"`
	QsOptions              qsOptions       `json:"qsOptions"`
	Silent                 bool            `json:"silent"`
	Messages               messages1       `json:"messages"`
}

type ButtonGroupControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               mock             `json:"remark"`
	LabelRemark          mock             `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	BtnClassName         RootSchema       `json:"btnClassName"`
	BtnActiveClassName   string           `json:"btnActiveClassName"`
	Buttons              []any            `json:"buttons"`
	BtnLevel             string           `json:"btnLevel"`
	BtnActiveLevel       string           `json:"btnActiveLevel"`
	Vertical             bool             `json:"vertical"`
	Tiled                bool             `json:"tiled"`
	Size                 string           `json:"size"`
	Testid               string           `json:"testid"`
	TypeVal              string           `json:"type"`
}

type ButtonGroupSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	BtnClassName         RootSchema    `json:"btnClassName"`
	BtnActiveClassName   string        `json:"btnActiveClassName"`
	Buttons              []any         `json:"buttons"`
	BtnLevel             string        `json:"btnLevel"`
	BtnActiveLevel       string        `json:"btnActiveLevel"`
	Vertical             bool          `json:"vertical"`
	Tiled                bool          `json:"tiled"`
	Size                 string        `json:"size"`
}

type ButtonToolbarSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Buttons              []any         `json:"buttons"`
}

type CRUD2CardsSchema struct {
	Card                          card                `json:"card"`
	HeaderClassName               className           `json:"headerClassName"`
	FooterClassName               className           `json:"footerClassName"`
	ItemClassName                 className           `json:"itemClassName"`
	Placeholder                   className           `json:"placeholder"`
	ShowFooter                    bool                `json:"showFooter"`
	ShowHeader                    bool                `json:"showHeader"`
	Source                        className           `json:"source"`
	Title                         className           `json:"title"`
	HideCheckToggler              bool                `json:"hideCheckToggler"`
	AffixHeader                   bool                `json:"affixHeader"`
	AffixFooter                   bool                `json:"affixFooter"`
	Header                        className           `json:"header"`
	Footer                        className           `json:"footer"`
	ItemCheckableOn               className           `json:"itemCheckableOn"`
	ItemDraggableOn               className           `json:"itemDraggableOn"`
	CheckOnItemClick              bool                `json:"checkOnItemClick"`
	MasonryLayout                 bool                `json:"masonryLayout"`
	ValueField                    string              `json:"valueField"`
	Testid                        string              `json:"testid"`
	ID                            string              `json:"$$id"`
	ClassName                     className           `json:"className"`
	Ref                           string              `json:"$ref"`
	Disabled                      bool                `json:"disabled"`
	DisabledOn                    className           `json:"disabledOn"`
	Hidden                        bool                `json:"hidden"`
	HiddenOn                      className           `json:"hiddenOn"`
	Visible                       bool                `json:"visible"`
	VisibleOn                     className           `json:"visibleOn"`
	Id                            string              `json:"id"`
	OnEvent                       any                 `json:"onEvent"`
	Static                        bool                `json:"static"`
	StaticOn                      className           `json:"staticOn"`
	StaticPlaceholder             string              `json:"staticPlaceholder"`
	StaticClassName               className           `json:"staticClassName"`
	StaticLabelClassName          className           `json:"staticLabelClassName"`
	StaticInputClassName          className           `json:"staticInputClassName"`
	StaticSchema                  staticSchema        `json:"staticSchema"`
	Style                         any                 `json:"style"`
	EditorSetting                 editorSetting       `json:"editorSetting"`
	UseMobileUI                   bool                `json:"useMobileUI"`
	TestIdBuilder                 RootSchema          `json:"testIdBuilder"`
	LoadingConfig                 loadingConfig       `json:"loadingConfig"`
	Mode                          string              `json:"mode"`
	TypeVal                       string              `json:"type"`
	Api                           className           `json:"api"`
	SilentPolling                 bool                `json:"silentPolling"`
	Interval                      float64             `json:"interval"`
	StopAutoRefreshWhen           RootSchema          `json:"stopAutoRefreshWhen"`
	LoadType                      string              `json:"loadType"`
	PerPage                       float64             `json:"perPage"`
	LoadDataOnce                  bool                `json:"loadDataOnce"`
	Selectable                    bool                `json:"selectable"`
	Multiple                      bool                `json:"multiple"`
	ShowSelection                 bool                `json:"showSelection"`
	QuickSaveApi                  className           `json:"quickSaveApi"`
	QuickSaveItemApi              className           `json:"quickSaveItemApi"`
	SaveOrderApi                  className           `json:"saveOrderApi"`
	SyncLocation                  bool                `json:"syncLocation"`
	PageField                     string              `json:"pageField"`
	PerPageField                  string              `json:"perPageField"`
	Name                          RootSchema          `json:"name"`
	HideQuickSaveBtn              bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange    bool                `json:"autoJumpToTopOnPagerChange"`
	HeaderToolbar                 className           `json:"headerToolbar"`
	HeaderToolbarClassName        string              `json:"headerToolbarClassName"`
	FooterToolbar                 className           `json:"footerToolbar"`
	FooterToolbarClassName        string              `json:"footerToolbarClassName"`
	SyncResponse2Query            bool                `json:"syncResponse2Query"`
	KeepItemSelectionOnPageChange bool                `json:"keepItemSelectionOnPageChange"`
	AutoFillHeight                bool                `json:"autoFillHeight"`
	PrimaryField                  string              `json:"primaryField"`
	ParsePrimitiveQuery           parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CRUD2ListSchema struct {
	Title                         className           `json:"title"`
	Footer                        className           `json:"footer"`
	FooterClassName               className           `json:"footerClassName"`
	Header                        className           `json:"header"`
	HeaderClassName               className           `json:"headerClassName"`
	ListItem                      className           `json:"listItem"`
	Source                        className           `json:"source"`
	ShowFooter                    bool                `json:"showFooter"`
	ShowHeader                    bool                `json:"showHeader"`
	Placeholder                   className           `json:"placeholder"`
	HideCheckToggler              bool                `json:"hideCheckToggler"`
	AffixHeader                   bool                `json:"affixHeader"`
	AffixFooter                   bool                `json:"affixFooter"`
	ItemCheckableOn               className           `json:"itemCheckableOn"`
	ItemDraggableOn               className           `json:"itemDraggableOn"`
	CheckOnItemClick              bool                `json:"checkOnItemClick"`
	ValueField                    string              `json:"valueField"`
	Size                          string              `json:"size"`
	ItemAction                    className           `json:"itemAction"`
	Testid                        string              `json:"testid"`
	ID                            string              `json:"$$id"`
	ClassName                     className           `json:"className"`
	Ref                           string              `json:"$ref"`
	Disabled                      bool                `json:"disabled"`
	DisabledOn                    className           `json:"disabledOn"`
	Hidden                        bool                `json:"hidden"`
	HiddenOn                      className           `json:"hiddenOn"`
	Visible                       bool                `json:"visible"`
	VisibleOn                     className           `json:"visibleOn"`
	Id                            string              `json:"id"`
	OnEvent                       any                 `json:"onEvent"`
	Static                        bool                `json:"static"`
	StaticOn                      className           `json:"staticOn"`
	StaticPlaceholder             string              `json:"staticPlaceholder"`
	StaticClassName               className           `json:"staticClassName"`
	StaticLabelClassName          className           `json:"staticLabelClassName"`
	StaticInputClassName          className           `json:"staticInputClassName"`
	StaticSchema                  staticSchema        `json:"staticSchema"`
	Style                         any                 `json:"style"`
	EditorSetting                 editorSetting       `json:"editorSetting"`
	UseMobileUI                   bool                `json:"useMobileUI"`
	TestIdBuilder                 RootSchema          `json:"testIdBuilder"`
	Mode                          string              `json:"mode"`
	LoadingConfig                 loadingConfig       `json:"loadingConfig"`
	TypeVal                       string              `json:"type"`
	Api                           className           `json:"api"`
	SilentPolling                 bool                `json:"silentPolling"`
	Interval                      float64             `json:"interval"`
	StopAutoRefreshWhen           RootSchema          `json:"stopAutoRefreshWhen"`
	LoadType                      string              `json:"loadType"`
	PerPage                       float64             `json:"perPage"`
	LoadDataOnce                  bool                `json:"loadDataOnce"`
	Selectable                    bool                `json:"selectable"`
	Multiple                      bool                `json:"multiple"`
	ShowSelection                 bool                `json:"showSelection"`
	QuickSaveApi                  className           `json:"quickSaveApi"`
	QuickSaveItemApi              className           `json:"quickSaveItemApi"`
	SaveOrderApi                  className           `json:"saveOrderApi"`
	SyncLocation                  bool                `json:"syncLocation"`
	PageField                     string              `json:"pageField"`
	PerPageField                  string              `json:"perPageField"`
	Name                          RootSchema          `json:"name"`
	HideQuickSaveBtn              bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange    bool                `json:"autoJumpToTopOnPagerChange"`
	HeaderToolbar                 className           `json:"headerToolbar"`
	HeaderToolbarClassName        string              `json:"headerToolbarClassName"`
	FooterToolbar                 className           `json:"footerToolbar"`
	FooterToolbarClassName        string              `json:"footerToolbarClassName"`
	SyncResponse2Query            bool                `json:"syncResponse2Query"`
	KeepItemSelectionOnPageChange bool                `json:"keepItemSelectionOnPageChange"`
	AutoFillHeight                bool                `json:"autoFillHeight"`
	PrimaryField                  string              `json:"primaryField"`
	ParsePrimitiveQuery           parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CRUD2TableSchema struct {
	Title                         SchemaClassName     `json:"title"`
	Source                        className           `json:"source"`
	ColumnsTogglable              columnsTogglable    `json:"columnsTogglable"`
	Columns                       []any               `json:"columns"`
	RowSelection                  draggableConfig     `json:"rowSelection"`
	Expandable                    className           `json:"expandable"`
	Sticky                        bool                `json:"sticky"`
	Loading                       SchemaClassName     `json:"loading"`
	ItemBadge                     className           `json:"itemBadge"`
	ShowBadge                     bool                `json:"showBadge"`
	PopOverContainer              mock                `json:"popOverContainer"`
	KeyField                      string              `json:"keyField"`
	ChildrenColumnName            string              `json:"childrenColumnName"`
	RowClassNameExpr              string              `json:"rowClassNameExpr"`
	LineHeight                    string              `json:"lineHeight"`
	Bordered                      bool                `json:"bordered"`
	ShowHeader                    bool                `json:"showHeader"`
	Footer                        SchemaClassName     `json:"footer"`
	QuickSaveApi                  className           `json:"quickSaveApi"`
	QuickSaveItemApi              className           `json:"quickSaveItemApi"`
	Messages                      className           `json:"messages"`
	Reload                        string              `json:"reload"`
	Actions                       []any               `json:"actions"`
	MaxKeepItemSelectionLength    float64             `json:"maxKeepItemSelectionLength"`
	KeepItemSelectionOnPageChange bool                `json:"keepItemSelectionOnPageChange"`
	Selectable                    bool                `json:"selectable"`
	Multiple                      bool                `json:"multiple"`
	PrimaryField                  string              `json:"primaryField"`
	TableLayout                   string              `json:"tableLayout"`
	AutoFillHeight                bool                `json:"autoFillHeight"`
	CanAccessSuperData            bool                `json:"canAccessSuperData"`
	LazyRenderAfter               float64             `json:"lazyRenderAfter"`
	Testid                        string              `json:"testid"`
	ID                            string              `json:"$$id"`
	ClassName                     className           `json:"className"`
	Ref                           string              `json:"$ref"`
	Disabled                      bool                `json:"disabled"`
	DisabledOn                    className           `json:"disabledOn"`
	Hidden                        bool                `json:"hidden"`
	HiddenOn                      className           `json:"hiddenOn"`
	Visible                       bool                `json:"visible"`
	VisibleOn                     className           `json:"visibleOn"`
	Id                            string              `json:"id"`
	OnEvent                       any                 `json:"onEvent"`
	Static                        bool                `json:"static"`
	StaticOn                      className           `json:"staticOn"`
	StaticPlaceholder             string              `json:"staticPlaceholder"`
	StaticClassName               className           `json:"staticClassName"`
	StaticLabelClassName          className           `json:"staticLabelClassName"`
	StaticInputClassName          className           `json:"staticInputClassName"`
	StaticSchema                  staticSchema        `json:"staticSchema"`
	Style                         any                 `json:"style"`
	EditorSetting                 editorSetting       `json:"editorSetting"`
	UseMobileUI                   bool                `json:"useMobileUI"`
	TestIdBuilder                 RootSchema          `json:"testIdBuilder"`
	Mode                          string              `json:"mode"`
	LoadingConfig                 loadingConfig       `json:"loadingConfig"`
	TypeVal                       string              `json:"type"`
	Api                           className           `json:"api"`
	SilentPolling                 bool                `json:"silentPolling"`
	Interval                      float64             `json:"interval"`
	StopAutoRefreshWhen           RootSchema          `json:"stopAutoRefreshWhen"`
	LoadType                      string              `json:"loadType"`
	PerPage                       float64             `json:"perPage"`
	LoadDataOnce                  bool                `json:"loadDataOnce"`
	ShowSelection                 bool                `json:"showSelection"`
	SaveOrderApi                  className           `json:"saveOrderApi"`
	SyncLocation                  bool                `json:"syncLocation"`
	PageField                     string              `json:"pageField"`
	PerPageField                  string              `json:"perPageField"`
	Name                          RootSchema          `json:"name"`
	HideQuickSaveBtn              bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange    bool                `json:"autoJumpToTopOnPagerChange"`
	HeaderToolbar                 className           `json:"headerToolbar"`
	HeaderToolbarClassName        string              `json:"headerToolbarClassName"`
	FooterToolbar                 className           `json:"footerToolbar"`
	FooterToolbarClassName        string              `json:"footerToolbarClassName"`
	SyncResponse2Query            bool                `json:"syncResponse2Query"`
	ParsePrimitiveQuery           parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CRUDCardsSchema struct {
	Card                           card                `json:"card"`
	HeaderClassName                className           `json:"headerClassName"`
	FooterClassName                className           `json:"footerClassName"`
	ItemClassName                  className           `json:"itemClassName"`
	Placeholder                    className           `json:"placeholder"`
	ShowFooter                     bool                `json:"showFooter"`
	ShowHeader                     bool                `json:"showHeader"`
	Source                         className           `json:"source"`
	Title                          className           `json:"title"`
	HideCheckToggler               bool                `json:"hideCheckToggler"`
	AffixHeader                    bool                `json:"affixHeader"`
	AffixFooter                    bool                `json:"affixFooter"`
	Header                         className           `json:"header"`
	Footer                         className           `json:"footer"`
	ItemCheckableOn                className           `json:"itemCheckableOn"`
	ItemDraggableOn                className           `json:"itemDraggableOn"`
	CheckOnItemClick               bool                `json:"checkOnItemClick"`
	MasonryLayout                  bool                `json:"masonryLayout"`
	ValueField                     string              `json:"valueField"`
	Testid                         string              `json:"testid"`
	ID                             string              `json:"$$id"`
	ClassName                      className           `json:"className"`
	Ref                            string              `json:"$ref"`
	Disabled                       bool                `json:"disabled"`
	DisabledOn                     className           `json:"disabledOn"`
	Hidden                         bool                `json:"hidden"`
	HiddenOn                       className           `json:"hiddenOn"`
	Visible                        bool                `json:"visible"`
	VisibleOn                      className           `json:"visibleOn"`
	Id                             string              `json:"id"`
	OnEvent                        any                 `json:"onEvent"`
	Static                         bool                `json:"static"`
	StaticOn                       className           `json:"staticOn"`
	StaticPlaceholder              string              `json:"staticPlaceholder"`
	StaticClassName                className           `json:"staticClassName"`
	StaticLabelClassName           className           `json:"staticLabelClassName"`
	StaticInputClassName           className           `json:"staticInputClassName"`
	StaticSchema                   staticSchema        `json:"staticSchema"`
	Style                          any                 `json:"style"`
	EditorSetting                  editorSetting       `json:"editorSetting"`
	UseMobileUI                    bool                `json:"useMobileUI"`
	TestIdBuilder                  RootSchema          `json:"testIdBuilder"`
	LoadingConfig                  loadingConfig       `json:"loadingConfig"`
	Mode                           string              `json:"mode"`
	TypeVal                        string              `json:"type"`
	Api                            className           `json:"api"`
	DeferApi                       className           `json:"deferApi"`
	BulkActions                    []any               `json:"bulkActions"`
	ItemActions                    []any               `json:"itemActions"`
	PerPage                        float64             `json:"perPage"`
	OrderBy                        string              `json:"orderBy"`
	OrderDir                       string              `json:"orderDir"`
	DefaultParams                  className           `json:"defaultParams"`
	Draggable                      bool                `json:"draggable"`
	DraggableOn                    className           `json:"draggableOn"`
	Name                           RootSchema          `json:"name"`
	Filter                         mock                `json:"filter"`
	InitFetch                      bool                `json:"initFetch"`
	InitFetchOn                    className           `json:"initFetchOn"`
	InnerClassName                 className           `json:"innerClassName"`
	Interval                       float64             `json:"interval"`
	OrderField                     string              `json:"orderField"`
	PageField                      string              `json:"pageField"`
	PerPageField                   string              `json:"perPageField"`
	PageDirectionField             string              `json:"pageDirectionField"`
	QuickSaveApi                   className           `json:"quickSaveApi"`
	QuickSaveItemApi               className           `json:"quickSaveItemApi"`
	SaveOrderApi                   className           `json:"saveOrderApi"`
	SyncLocation                   bool                `json:"syncLocation"`
	HeaderToolbar                  []any               `json:"headerToolbar"`
	FooterToolbar                  []any               `json:"footerToolbar"`
	PerPageAvailable               []float64           `json:"perPageAvailable"`
	Messages                       RootSchema          `json:"messages"`
	HideQuickSaveBtn               bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange     bool                `json:"autoJumpToTopOnPagerChange"`
	SilentPolling                  bool                `json:"silentPolling"`
	StopAutoRefreshWhen            RootSchema          `json:"stopAutoRefreshWhen"`
	StopAutoRefreshWhenModalIsOpen bool                `json:"stopAutoRefreshWhenModalIsOpen"`
	FilterTogglable                data                `json:"filterTogglable"`
	FilterDefaultVisible           bool                `json:"filterDefaultVisible"`
	SyncResponse2Query             bool                `json:"syncResponse2Query"`
	KeepItemSelectionOnPageChange  bool                `json:"keepItemSelectionOnPageChange"`
	LabelTpl                       className           `json:"labelTpl"`
	LoadDataOnce                   bool                `json:"loadDataOnce"`
	LoadDataOnceFetchOnFilter      bool                `json:"loadDataOnceFetchOnFilter"`
	MatchFunc                      SchemaClassName     `json:"matchFunc"`
	ExpandConfig                   expandConfig        `json:"expandConfig"`
	AlwaysShowPagination           bool                `json:"alwaysShowPagination"`
	AutoGenerateFilter             draggableConfig     `json:"autoGenerateFilter"`
	AutoFillHeight                 SchemaClassName     `json:"autoFillHeight"`
	ParsePrimitiveQuery            parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CRUDListSchema struct {
	Title                          className           `json:"title"`
	Footer                         className           `json:"footer"`
	FooterClassName                className           `json:"footerClassName"`
	Header                         className           `json:"header"`
	HeaderClassName                className           `json:"headerClassName"`
	ListItem                       className           `json:"listItem"`
	Source                         className           `json:"source"`
	ShowFooter                     bool                `json:"showFooter"`
	ShowHeader                     bool                `json:"showHeader"`
	Placeholder                    className           `json:"placeholder"`
	HideCheckToggler               bool                `json:"hideCheckToggler"`
	AffixHeader                    bool                `json:"affixHeader"`
	AffixFooter                    bool                `json:"affixFooter"`
	ItemCheckableOn                className           `json:"itemCheckableOn"`
	ItemDraggableOn                className           `json:"itemDraggableOn"`
	CheckOnItemClick               bool                `json:"checkOnItemClick"`
	ValueField                     string              `json:"valueField"`
	Size                           string              `json:"size"`
	ItemAction                     className           `json:"itemAction"`
	Testid                         string              `json:"testid"`
	ID                             string              `json:"$$id"`
	ClassName                      className           `json:"className"`
	Ref                            string              `json:"$ref"`
	Disabled                       bool                `json:"disabled"`
	DisabledOn                     className           `json:"disabledOn"`
	Hidden                         bool                `json:"hidden"`
	HiddenOn                       className           `json:"hiddenOn"`
	Visible                        bool                `json:"visible"`
	VisibleOn                      className           `json:"visibleOn"`
	Id                             string              `json:"id"`
	OnEvent                        any                 `json:"onEvent"`
	Static                         bool                `json:"static"`
	StaticOn                       className           `json:"staticOn"`
	StaticPlaceholder              string              `json:"staticPlaceholder"`
	StaticClassName                className           `json:"staticClassName"`
	StaticLabelClassName           className           `json:"staticLabelClassName"`
	StaticInputClassName           className           `json:"staticInputClassName"`
	StaticSchema                   staticSchema        `json:"staticSchema"`
	Style                          any                 `json:"style"`
	EditorSetting                  editorSetting       `json:"editorSetting"`
	UseMobileUI                    bool                `json:"useMobileUI"`
	TestIdBuilder                  RootSchema          `json:"testIdBuilder"`
	Mode                           string              `json:"mode"`
	LoadingConfig                  loadingConfig       `json:"loadingConfig"`
	TypeVal                        string              `json:"type"`
	Api                            className           `json:"api"`
	DeferApi                       className           `json:"deferApi"`
	BulkActions                    []any               `json:"bulkActions"`
	ItemActions                    []any               `json:"itemActions"`
	PerPage                        float64             `json:"perPage"`
	OrderBy                        string              `json:"orderBy"`
	OrderDir                       string              `json:"orderDir"`
	DefaultParams                  className           `json:"defaultParams"`
	Draggable                      bool                `json:"draggable"`
	DraggableOn                    className           `json:"draggableOn"`
	Name                           RootSchema          `json:"name"`
	Filter                         mock                `json:"filter"`
	InitFetch                      bool                `json:"initFetch"`
	InitFetchOn                    className           `json:"initFetchOn"`
	InnerClassName                 className           `json:"innerClassName"`
	Interval                       float64             `json:"interval"`
	OrderField                     string              `json:"orderField"`
	PageField                      string              `json:"pageField"`
	PerPageField                   string              `json:"perPageField"`
	PageDirectionField             string              `json:"pageDirectionField"`
	QuickSaveApi                   className           `json:"quickSaveApi"`
	QuickSaveItemApi               className           `json:"quickSaveItemApi"`
	SaveOrderApi                   className           `json:"saveOrderApi"`
	SyncLocation                   bool                `json:"syncLocation"`
	HeaderToolbar                  []any               `json:"headerToolbar"`
	FooterToolbar                  []any               `json:"footerToolbar"`
	PerPageAvailable               []float64           `json:"perPageAvailable"`
	Messages                       RootSchema          `json:"messages"`
	HideQuickSaveBtn               bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange     bool                `json:"autoJumpToTopOnPagerChange"`
	SilentPolling                  bool                `json:"silentPolling"`
	StopAutoRefreshWhen            RootSchema          `json:"stopAutoRefreshWhen"`
	StopAutoRefreshWhenModalIsOpen bool                `json:"stopAutoRefreshWhenModalIsOpen"`
	FilterTogglable                data                `json:"filterTogglable"`
	FilterDefaultVisible           bool                `json:"filterDefaultVisible"`
	SyncResponse2Query             bool                `json:"syncResponse2Query"`
	KeepItemSelectionOnPageChange  bool                `json:"keepItemSelectionOnPageChange"`
	LabelTpl                       className           `json:"labelTpl"`
	LoadDataOnce                   bool                `json:"loadDataOnce"`
	LoadDataOnceFetchOnFilter      bool                `json:"loadDataOnceFetchOnFilter"`
	MatchFunc                      SchemaClassName     `json:"matchFunc"`
	ExpandConfig                   expandConfig        `json:"expandConfig"`
	AlwaysShowPagination           bool                `json:"alwaysShowPagination"`
	AutoGenerateFilter             draggableConfig     `json:"autoGenerateFilter"`
	AutoFillHeight                 SchemaClassName     `json:"autoFillHeight"`
	ParsePrimitiveQuery            parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CRUDTableSchema struct {
	AffixHeader                    bool                `json:"affixHeader"`
	AffixFooter                    bool                `json:"affixFooter"`
	Columns                        []any               `json:"columns"`
	ColumnsTogglable               SchemaClassName     `json:"columnsTogglable"`
	Footable                       SchemaClassName     `json:"footable"`
	FooterClassName                className           `json:"footerClassName"`
	HeaderClassName                className           `json:"headerClassName"`
	Placeholder                    SchemaClassName     `json:"placeholder"`
	ShowFooter                     bool                `json:"showFooter"`
	ShowHeader                     bool                `json:"showHeader"`
	Source                         className           `json:"source"`
	TableClassName                 className           `json:"tableClassName"`
	Title                          string              `json:"title"`
	ToolbarClassName               className           `json:"toolbarClassName"`
	CombineNum                     SchemaClassName     `json:"combineNum"`
	CombineFromIndex               float64             `json:"combineFromIndex"`
	PrefixRow                      []any               `json:"prefixRow"`
	AffixRow                       []any               `json:"affixRow"`
	Resizable                      bool                `json:"resizable"`
	RowClassNameExpr               string              `json:"rowClassNameExpr"`
	ItemBadge                      className           `json:"itemBadge"`
	AutoGenerateFilter             draggableConfig     `json:"autoGenerateFilter"`
	CanAccessSuperData             bool                `json:"canAccessSuperData"`
	AutoFillHeight                 SchemaClassName     `json:"autoFillHeight"`
	TableLayout                    string              `json:"tableLayout"`
	DeferApi                       className           `json:"deferApi"`
	Testid                         string              `json:"testid"`
	ID                             string              `json:"$$id"`
	ClassName                      className           `json:"className"`
	Ref                            string              `json:"$ref"`
	Disabled                       bool                `json:"disabled"`
	DisabledOn                     className           `json:"disabledOn"`
	Hidden                         bool                `json:"hidden"`
	HiddenOn                       className           `json:"hiddenOn"`
	Visible                        bool                `json:"visible"`
	VisibleOn                      className           `json:"visibleOn"`
	Id                             string              `json:"id"`
	OnEvent                        any                 `json:"onEvent"`
	Static                         bool                `json:"static"`
	StaticOn                       className           `json:"staticOn"`
	StaticPlaceholder              string              `json:"staticPlaceholder"`
	StaticClassName                className           `json:"staticClassName"`
	StaticLabelClassName           className           `json:"staticLabelClassName"`
	StaticInputClassName           className           `json:"staticInputClassName"`
	StaticSchema                   staticSchema        `json:"staticSchema"`
	Style                          any                 `json:"style"`
	EditorSetting                  editorSetting       `json:"editorSetting"`
	UseMobileUI                    bool                `json:"useMobileUI"`
	TestIdBuilder                  RootSchema          `json:"testIdBuilder"`
	Mode                           string              `json:"mode"`
	LoadingConfig                  loadingConfig       `json:"loadingConfig"`
	TypeVal                        string              `json:"type"`
	Api                            className           `json:"api"`
	BulkActions                    []any               `json:"bulkActions"`
	ItemActions                    []any               `json:"itemActions"`
	PerPage                        float64             `json:"perPage"`
	OrderBy                        string              `json:"orderBy"`
	OrderDir                       string              `json:"orderDir"`
	DefaultParams                  className           `json:"defaultParams"`
	Draggable                      bool                `json:"draggable"`
	DraggableOn                    className           `json:"draggableOn"`
	Name                           RootSchema          `json:"name"`
	Filter                         mock                `json:"filter"`
	InitFetch                      bool                `json:"initFetch"`
	InitFetchOn                    className           `json:"initFetchOn"`
	InnerClassName                 className           `json:"innerClassName"`
	Interval                       float64             `json:"interval"`
	OrderField                     string              `json:"orderField"`
	PageField                      string              `json:"pageField"`
	PerPageField                   string              `json:"perPageField"`
	PageDirectionField             string              `json:"pageDirectionField"`
	QuickSaveApi                   className           `json:"quickSaveApi"`
	QuickSaveItemApi               className           `json:"quickSaveItemApi"`
	SaveOrderApi                   className           `json:"saveOrderApi"`
	SyncLocation                   bool                `json:"syncLocation"`
	HeaderToolbar                  []any               `json:"headerToolbar"`
	FooterToolbar                  []any               `json:"footerToolbar"`
	PerPageAvailable               []float64           `json:"perPageAvailable"`
	Messages                       RootSchema          `json:"messages"`
	HideQuickSaveBtn               bool                `json:"hideQuickSaveBtn"`
	AutoJumpToTopOnPagerChange     bool                `json:"autoJumpToTopOnPagerChange"`
	SilentPolling                  bool                `json:"silentPolling"`
	StopAutoRefreshWhen            RootSchema          `json:"stopAutoRefreshWhen"`
	StopAutoRefreshWhenModalIsOpen bool                `json:"stopAutoRefreshWhenModalIsOpen"`
	FilterTogglable                data                `json:"filterTogglable"`
	FilterDefaultVisible           bool                `json:"filterDefaultVisible"`
	SyncResponse2Query             bool                `json:"syncResponse2Query"`
	KeepItemSelectionOnPageChange  bool                `json:"keepItemSelectionOnPageChange"`
	LabelTpl                       className           `json:"labelTpl"`
	LoadDataOnce                   bool                `json:"loadDataOnce"`
	LoadDataOnceFetchOnFilter      bool                `json:"loadDataOnceFetchOnFilter"`
	MatchFunc                      SchemaClassName     `json:"matchFunc"`
	ExpandConfig                   expandConfig        `json:"expandConfig"`
	AlwaysShowPagination           bool                `json:"alwaysShowPagination"`
	ParsePrimitiveQuery            parsePrimitiveQuery `json:"parsePrimitiveQuery"`
}

type CalendarSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Schedules            schedules     `json:"schedules"`
	ScheduleClassNames   []string      `json:"scheduleClassNames"`
	ScheduleAction       className     `json:"scheduleAction"`
	LargeMode            bool          `json:"largeMode"`
	TodayActiveStyle     any           `json:"todayActiveStyle"`
}

type Card2Schema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	HideCheckToggler     bool          `json:"hideCheckToggler"`
	CheckOnItemClick     bool          `json:"checkOnItemClick"`
	WrapperComponent     string        `json:"wrapperComponent"`
}

type CardBodyField struct {
	AllOf []ArrayElem7 `json:"allOf"`
}

type CardSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Header               header        `json:"header"`
	Body                 []any         `json:"body"`
	Media                media         `json:"media"`
	Actions              []any         `json:"actions"`
	Toolbar              []any         `json:"toolbar"`
	Secondary            className     `json:"secondary"`
	UseCardLabel         bool          `json:"useCardLabel"`
}

type CardsSchema struct {
	LoadingConfig        loadingConfig `json:"loadingConfig"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Card                 card          `json:"card"`
	HeaderClassName      className     `json:"headerClassName"`
	FooterClassName      className     `json:"footerClassName"`
	ItemClassName        className     `json:"itemClassName"`
	Placeholder          className     `json:"placeholder"`
	ShowFooter           bool          `json:"showFooter"`
	ShowHeader           bool          `json:"showHeader"`
	Source               className     `json:"source"`
	Title                className     `json:"title"`
	HideCheckToggler     bool          `json:"hideCheckToggler"`
	AffixHeader          bool          `json:"affixHeader"`
	AffixFooter          bool          `json:"affixFooter"`
	Header               className     `json:"header"`
	Footer               className     `json:"footer"`
	ItemCheckableOn      className     `json:"itemCheckableOn"`
	ItemDraggableOn      className     `json:"itemDraggableOn"`
	CheckOnItemClick     bool          `json:"checkOnItemClick"`
	MasonryLayout        bool          `json:"masonryLayout"`
	ValueField           string        `json:"valueField"`
}

type CarouselSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Auto                 bool          `json:"auto"`
	Interval             any           `json:"interval"`
	Duration             float64       `json:"duration"`
	Width                float64       `json:"width"`
	Height               float64       `json:"height"`
	ControlsTheme        string        `json:"controlsTheme"`
	Placeholder          string        `json:"placeholder"`
	Controls             []string      `json:"controls"`
	Animation            string        `json:"animation"`
	ItemSchema           className     `json:"itemSchema"`
	Name                 RootSchema    `json:"name"`
	ThumbMode            string        `json:"thumbMode"`
	Options              []any         `json:"options"`
	AlwaysShowArrow      bool          `json:"alwaysShowArrow"`
	Multiple             multiple      `json:"multiple"`
	Icons                icons         `json:"icons"`
}

type ChainedSelectControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
}

type ChartSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	ChartTheme           mock          `json:"chartTheme"`
	Api                  className     `json:"api"`
	InitFetch            bool          `json:"initFetch"`
	InitFetchOn          className     `json:"initFetchOn"`
	Config               mock          `json:"config"`
	TrackExpression      string        `json:"trackExpression"`
	Width                any           `json:"width"`
	Height               any           `json:"height"`
	Interval             float64       `json:"interval"`
	Name                 RootSchema    `json:"name"`
	DataFilter           RootSchema    `json:"dataFilter"`
	Source               RootSchema    `json:"source"`
	DisableDataMapping   bool          `json:"disableDataMapping"`
	ClickAction          className     `json:"clickAction"`
	ReplaceChartOption   bool          `json:"replaceChartOption"`
	UnMountOnHidden      bool          `json:"unMountOnHidden"`
	MapURL               className     `json:"mapURL"`
	MapName              string        `json:"mapName"`
	LoadBaiduMap         bool          `json:"loadBaiduMap"`
}

type CheckboxControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	TrueValue            any              `json:"trueValue"`
	FalseValue           any              `json:"falseValue"`
	Option               string           `json:"option"`
	Badge                className        `json:"badge"`
	Partial              bool             `json:"partial"`
	OptionType           string           `json:"optionType"`
	Checked              bool             `json:"checked"`
}

type CheckboxesControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	CheckAll             bool             `json:"checkAll"`
	DefaultCheckAll      bool             `json:"defaultCheckAll"`
	CheckAllText         string           `json:"checkAllText"`
	ColumnsCount         SchemaClassName  `json:"columnsCount"`
	MenuTpl              string           `json:"menuTpl"`
}

type CollapseGroupSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	ActiveKey            activeKey     `json:"activeKey"`
	Accordion            bool          `json:"accordion"`
	ExpandIcon           className     `json:"expandIcon"`
	ExpandIconPosition   string        `json:"expandIconPosition"`
	Body                 className     `json:"body"`
	EnableFieldSetStyle  bool          `json:"enableFieldSetStyle"`
}

type CollapseSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Key                  string          `json:"key"`
	HeaderPosition       string          `json:"headerPosition"`
	Header               SchemaClassName `json:"header"`
	Body                 className       `json:"body"`
	BodyClassName        string          `json:"bodyClassName"`
	Collapsable          bool            `json:"collapsable"`
	Collapsed            bool            `json:"collapsed"`
	ShowArrow            bool            `json:"showArrow"`
	ExpandIcon           className       `json:"expandIcon"`
	HeadingClassName     string          `json:"headingClassName"`
	CollapseHeader       className       `json:"collapseHeader"`
	Size                 string          `json:"size"`
	MountOnEnter         bool            `json:"mountOnEnter"`
	UnmountOnExit        bool            `json:"unmountOnExit"`
	DivideLine           bool            `json:"divideLine"`
}

type ColorMapType struct {
	AnyOf []ArrayElem19 `json:"anyOf"`
}

type ColorSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	DefaultColor         string        `json:"defaultColor"`
	ShowValue            bool          `json:"showValue"`
}

type ColumnSchema struct {
	Name               string          `json:"name"`
	Title              SchemaClassName `json:"title"`
	TypeVal            string          `json:"type"`
	RowSpanExpr        string          `json:"rowSpanExpr"`
	ColSpanExpr        string          `json:"colSpanExpr"`
	Children           []any           `json:"children"`
	Copyable           bool            `json:"copyable"`
	Remark             string          `json:"remark"`
	Searchable         SchemaClassName `json:"searchable"`
	Sorter             bool            `json:"sorter"`
	Sortable           bool            `json:"sortable"`
	Filterable         filterable      `json:"filterable"`
	Align              string          `json:"align"`
	Fixed              any             `json:"fixed"`
	Toggled            bool            `json:"toggled"`
	ClassName          string          `json:"className"`
	TitleClassName     string          `json:"titleClassName"`
	ClassNameExpr      string          `json:"classNameExpr"`
	QuickEdit          className       `json:"quickEdit"`
	Width              any             `json:"width"`
	CanAccessSuperData bool            `json:"canAccessSuperData"`
}

type ComboCondition struct {
	Test     string       `json:"test"`
	Items    []any        `json:"items"`
	Label    string       `json:"label"`
	Scaffold staticSchema `json:"scaffold"`
	Mode     string       `json:"mode"`
}

type ComboControlSchema struct {
	ID                                 string           `json:"$$id"`
	ClassName                          className        `json:"className"`
	Ref                                string           `json:"$ref"`
	Disabled                           bool             `json:"disabled"`
	DisabledOn                         className        `json:"disabledOn"`
	Hidden                             bool             `json:"hidden"`
	HiddenOn                           className        `json:"hiddenOn"`
	Visible                            bool             `json:"visible"`
	VisibleOn                          className        `json:"visibleOn"`
	Id                                 string           `json:"id"`
	OnEvent                            any              `json:"onEvent"`
	Static                             bool             `json:"static"`
	StaticOn                           className        `json:"staticOn"`
	StaticPlaceholder                  string           `json:"staticPlaceholder"`
	StaticClassName                    className        `json:"staticClassName"`
	StaticLabelClassName               className        `json:"staticLabelClassName"`
	StaticInputClassName               className        `json:"staticInputClassName"`
	StaticSchema                       staticSchema     `json:"staticSchema"`
	Style                              any              `json:"style"`
	EditorSetting                      editorSetting    `json:"editorSetting"`
	UseMobileUI                        bool             `json:"useMobileUI"`
	TestIdBuilder                      RootSchema       `json:"testIdBuilder"`
	Size                               string           `json:"size"`
	Label                              SchemaClassName  `json:"label"`
	LabelAlign                         className        `json:"labelAlign"`
	LabelWidth                         any              `json:"labelWidth"`
	LabelClassName                     string           `json:"labelClassName"`
	Name                               string           `json:"name"`
	ExtraName                          string           `json:"extraName"`
	Remark                             className        `json:"remark"`
	LabelRemark                        className        `json:"labelRemark"`
	Hint                               string           `json:"hint"`
	SubmitOnChange                     bool             `json:"submitOnChange"`
	ReadOnly                           bool             `json:"readOnly"`
	ReadOnlyOn                         string           `json:"readOnlyOn"`
	ValidateOnChange                   bool             `json:"validateOnChange"`
	Description                        string           `json:"description"`
	Desc                               string           `json:"desc"`
	DescriptionClassName               className        `json:"descriptionClassName"`
	Mode                               string           `json:"mode"`
	Horizontal                         className        `json:"horizontal"`
	Inline                             bool             `json:"inline"`
	InputClassName                     className        `json:"inputClassName"`
	Placeholder                        string           `json:"placeholder"`
	Required                           bool             `json:"required"`
	ValidationErrors                   validationErrors `json:"validationErrors"`
	Validations                        data             `json:"validations"`
	Value                              mock             `json:"value"`
	ClearValueOnHidden                 bool             `json:"clearValueOnHidden"`
	ValidateApi                        SchemaClassName  `json:"validateApi"`
	AutoFill                           autoFill         `json:"autoFill"`
	InitAutoFill                       data             `json:"initAutoFill"`
	Row                                float64          `json:"row"`
	TypeVal                            string           `json:"type"`
	Scaffold                           scaffold         `json:"scaffold"`
	NoBorder                           bool             `json:"noBorder"`
	DeleteConfirmText                  string           `json:"deleteConfirmText"`
	DeleteApi                          className        `json:"deleteApi"`
	TypeSwitchable                     bool             `json:"typeSwitchable"`
	Conditions                         []any            `json:"conditions"`
	FormClassName                      className        `json:"formClassName"`
	AddButtonClassName                 className        `json:"addButtonClassName"`
	AddButtonText                      string           `json:"addButtonText"`
	Addable                            bool             `json:"addable"`
	Addattop                           bool             `json:"addattop"`
	Items                              []any            `json:"items"`
	Draggable                          bool             `json:"draggable"`
	DraggableTip                       string           `json:"draggableTip"`
	Flat                               bool             `json:"flat"`
	Delimiter                          string           `json:"delimiter"`
	JoinValues                         bool             `json:"joinValues"`
	MaxLength                          SchemaClassName  `json:"maxLength"`
	MinLength                          SchemaClassName  `json:"minLength"`
	MultiLine                          bool             `json:"multiLine"`
	Multiple                           bool             `json:"multiple"`
	Removable                          bool             `json:"removable"`
	SubFormMode                        string           `json:"subFormMode"`
	SubFormHorizontal                  className        `json:"subFormHorizontal"`
	CanAccessSuperData                 bool             `json:"canAccessSuperData"`
	TabsMode                           bool             `json:"tabsMode"`
	TabsStyle                          string           `json:"tabsStyle"`
	TabsLabelTpl                       className        `json:"tabsLabelTpl"`
	LazyLoad                           bool             `json:"lazyLoad"`
	StrictMode                         bool             `json:"strictMode"`
	SyncFields                         []string         `json:"syncFields"`
	Nullable                           bool             `json:"nullable"`
	Messages                           messages2        `json:"messages"`
	UpdatePristineAfterStoreDataReInit bool             `json:"updatePristineAfterStoreDataReInit"`
}

type ComboSubControl struct {
	AllOf []ArrayElem24 `json:"allOf"`
}

type ConditionBuilderConfig struct {
	ValueTypes []string   `json:"valueTypes"`
	Fields     RootSchema `json:"fields"`
	Funcs      RootSchema `json:"funcs"`
	MaxLevel   float64    `json:"maxLevel"`
	Types      any        `json:"types"`
}

type ConditionBuilderControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Embed                bool             `json:"embed"`
	PickerIcon           className        `json:"pickerIcon"`
	Funcs                className        `json:"funcs"`
	Fields               className        `json:"fields"`
	Config               className        `json:"config"`
	Source               draggableConfig  `json:"source"`
	BuilderMode          string           `json:"builderMode"`
	ShowANDOR            bool             `json:"showANDOR"`
	Draggable            bool             `json:"draggable"`
	AddBtnVisibleOn      string           `json:"addBtnVisibleOn"`
	AddGroupBtnVisibleOn string           `json:"addGroupBtnVisibleOn"`
	Formula              formula          `json:"formula"`
	FormulaForIf         mock             `json:"formulaForIf"`
}

type ConditionBuilderFuncArg struct {
	TypeVal      RootSchema   `json:"type"`
	Label        string       `json:"label"`
	ValueTypes   []string     `json:"valueTypes"`
	Operators    []any        `json:"operators"`
	Funcs        []string     `json:"funcs"`
	DefaultValue staticSchema `json:"defaultValue"`
	Placeholder  string       `json:"placeholder"`
	IsOptional   bool         `json:"isOptional"`
}

type ConditionBuilderType struct {
	DefaultOp   RootSchema `json:"defaultOp"`
	Operators   []any      `json:"operators"`
	Placeholder string     `json:"placeholder"`
	ValueTypes  []string   `json:"valueTypes"`
}

type ConditionFieldFunc struct {
	TypeVal    string     `json:"type"`
	ReturnType RootSchema `json:"returnType"`
	Args       []any      `json:"args"`
	Label      string     `json:"label"`
}

type ConditionGroupValue struct {
	Id          string `json:"id"`
	Conjunction string `json:"conjunction"`
	Not         bool   `json:"not"`
	Children    []any  `json:"children"`
	IfVal       string `json:"if"`
}

type ConditionRule struct {
	Id    staticSchema `json:"id"`
	Left  RootSchema   `json:"left"`
	Op    RootSchema   `json:"op"`
	Right right        `json:"right"`
	IfVal string       `json:"if"`
}

type ContainerDraggableConfig struct {
	Axis             string          `json:"axis"`
	DefaultPosition  defaultPosition `json:"defaultPosition"`
	Bounds           bounds          `json:"bounds"`
	Grid             []float64       `json:"grid"`
	Handle           string          `json:"handle"`
	Cancel           string          `json:"cancel"`
	Scale            float64         `json:"scale"`
	EnableUserSelect bool            `json:"enableUserSelect"`
}

type ContainerSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Body                 className       `json:"body"`
	BodyClassName        className       `json:"bodyClassName"`
	WrapperComponent     string          `json:"wrapperComponent"`
	WrapperBody          bool            `json:"wrapperBody"`
	Draggable            any             `json:"draggable"`
	DraggableConfig      draggableConfig `json:"draggableConfig"`
}

type CopyActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	CopyVal              className       `json:"copy"`
}

type CustomField struct {
	TypeVal      string       `json:"type"`
	Label        string       `json:"label"`
	ValueTypes   []string     `json:"valueTypes"`
	Operators    []any        `json:"operators"`
	Funcs        []string     `json:"funcs"`
	DefaultValue staticSchema `json:"defaultValue"`
	Placeholder  string       `json:"placeholder"`
	Name         string       `json:"name"`
	Value        staticSchema `json:"value"`
}

type DataProviderCollection struct {
	Inited             RootSchema `json:"inited"`
	OnApiFetched       RootSchema `json:"onApiFetched"`
	OnSchemaApiFetched RootSchema `json:"onSchemaApiFetched"`
	OnWsFetched        RootSchema `json:"onWsFetched"`
}

type DateControlSchema struct {
	ID                       string           `json:"$$id"`
	ClassName                className        `json:"className"`
	Ref                      string           `json:"$ref"`
	Disabled                 bool             `json:"disabled"`
	DisabledOn               className        `json:"disabledOn"`
	Hidden                   bool             `json:"hidden"`
	HiddenOn                 className        `json:"hiddenOn"`
	Visible                  bool             `json:"visible"`
	VisibleOn                className        `json:"visibleOn"`
	Id                       string           `json:"id"`
	OnEvent                  any              `json:"onEvent"`
	Static                   bool             `json:"static"`
	StaticOn                 className        `json:"staticOn"`
	StaticPlaceholder        string           `json:"staticPlaceholder"`
	StaticClassName          className        `json:"staticClassName"`
	StaticLabelClassName     className        `json:"staticLabelClassName"`
	StaticInputClassName     className        `json:"staticInputClassName"`
	StaticSchema             staticSchema     `json:"staticSchema"`
	Style                    any              `json:"style"`
	EditorSetting            editorSetting    `json:"editorSetting"`
	UseMobileUI              bool             `json:"useMobileUI"`
	TestIdBuilder            RootSchema       `json:"testIdBuilder"`
	Size                     string           `json:"size"`
	Label                    SchemaClassName  `json:"label"`
	LabelAlign               className        `json:"labelAlign"`
	LabelWidth               any              `json:"labelWidth"`
	LabelClassName           string           `json:"labelClassName"`
	Name                     string           `json:"name"`
	ExtraName                string           `json:"extraName"`
	Remark                   className        `json:"remark"`
	LabelRemark              className        `json:"labelRemark"`
	Hint                     string           `json:"hint"`
	SubmitOnChange           bool             `json:"submitOnChange"`
	ReadOnly                 bool             `json:"readOnly"`
	ReadOnlyOn               string           `json:"readOnlyOn"`
	ValidateOnChange         bool             `json:"validateOnChange"`
	Description              string           `json:"description"`
	Desc                     string           `json:"desc"`
	DescriptionClassName     className        `json:"descriptionClassName"`
	Mode                     string           `json:"mode"`
	Horizontal               className        `json:"horizontal"`
	Inline                   bool             `json:"inline"`
	InputClassName           className        `json:"inputClassName"`
	Placeholder              string           `json:"placeholder"`
	Required                 bool             `json:"required"`
	ValidationErrors         validationErrors `json:"validationErrors"`
	Validations              data             `json:"validations"`
	Value                    mock             `json:"value"`
	ClearValueOnHidden       bool             `json:"clearValueOnHidden"`
	ValidateApi              SchemaClassName  `json:"validateApi"`
	AutoFill                 autoFill         `json:"autoFill"`
	InitAutoFill             data             `json:"initAutoFill"`
	Row                      float64          `json:"row"`
	TypeVal                  string           `json:"type"`
	Clearable                bool             `json:"clearable"`
	Format                   string           `json:"format"`
	ValueFormat              string           `json:"valueFormat"`
	InputFormat              string           `json:"inputFormat"`
	DisplayFormat            string           `json:"displayFormat"`
	Utc                      bool             `json:"utc"`
	Emebed                   bool             `json:"emebed"`
	BorderMode               string           `json:"borderMode"`
	Shortcuts                SchemaClassName  `json:"shortcuts"`
	DisabledDate             string           `json:"disabledDate"`
	CloseOnSelect            bool             `json:"closeOnSelect"`
	MinDate                  string           `json:"minDate"`
	MaxDate                  string           `json:"maxDate"`
	PopOverContainerSelector string           `json:"popOverContainerSelector"`
}

type DateRangeControlSchema struct {
	ID                       string           `json:"$$id"`
	ClassName                className        `json:"className"`
	Ref                      string           `json:"$ref"`
	Disabled                 bool             `json:"disabled"`
	DisabledOn               className        `json:"disabledOn"`
	Hidden                   bool             `json:"hidden"`
	HiddenOn                 className        `json:"hiddenOn"`
	Visible                  bool             `json:"visible"`
	VisibleOn                className        `json:"visibleOn"`
	Id                       string           `json:"id"`
	OnEvent                  any              `json:"onEvent"`
	Static                   bool             `json:"static"`
	StaticOn                 className        `json:"staticOn"`
	StaticPlaceholder        string           `json:"staticPlaceholder"`
	StaticClassName          className        `json:"staticClassName"`
	StaticLabelClassName     className        `json:"staticLabelClassName"`
	StaticInputClassName     className        `json:"staticInputClassName"`
	StaticSchema             staticSchema     `json:"staticSchema"`
	Style                    any              `json:"style"`
	EditorSetting            editorSetting    `json:"editorSetting"`
	UseMobileUI              bool             `json:"useMobileUI"`
	TestIdBuilder            RootSchema       `json:"testIdBuilder"`
	Size                     string           `json:"size"`
	Label                    SchemaClassName  `json:"label"`
	LabelAlign               className        `json:"labelAlign"`
	LabelWidth               any              `json:"labelWidth"`
	LabelClassName           string           `json:"labelClassName"`
	Name                     string           `json:"name"`
	ExtraName                string           `json:"extraName"`
	Remark                   className        `json:"remark"`
	LabelRemark              className        `json:"labelRemark"`
	Hint                     string           `json:"hint"`
	SubmitOnChange           bool             `json:"submitOnChange"`
	ReadOnly                 bool             `json:"readOnly"`
	ReadOnlyOn               string           `json:"readOnlyOn"`
	ValidateOnChange         bool             `json:"validateOnChange"`
	Description              string           `json:"description"`
	Desc                     string           `json:"desc"`
	DescriptionClassName     className        `json:"descriptionClassName"`
	Mode                     string           `json:"mode"`
	Horizontal               className        `json:"horizontal"`
	Inline                   bool             `json:"inline"`
	InputClassName           className        `json:"inputClassName"`
	Placeholder              string           `json:"placeholder"`
	Required                 bool             `json:"required"`
	ValidationErrors         validationErrors `json:"validationErrors"`
	Validations              data             `json:"validations"`
	Value                    mock             `json:"value"`
	ClearValueOnHidden       bool             `json:"clearValueOnHidden"`
	ValidateApi              SchemaClassName  `json:"validateApi"`
	AutoFill                 autoFill         `json:"autoFill"`
	InitAutoFill             data             `json:"initAutoFill"`
	Row                      float64          `json:"row"`
	TypeVal                  string           `json:"type"`
	Delimiter                string           `json:"delimiter"`
	Format                   string           `json:"format"`
	ValueFormat              string           `json:"valueFormat"`
	InputFormat              string           `json:"inputFormat"`
	DisplayFormat            string           `json:"displayFormat"`
	JoinValues               bool             `json:"joinValues"`
	MaxDate                  string           `json:"maxDate"`
	MinDate                  string           `json:"minDate"`
	MaxDuration              string           `json:"maxDuration"`
	MinDuration              string           `json:"minDuration"`
	BorderMode               string           `json:"borderMode"`
	Embed                    bool             `json:"embed"`
	Ranges                   SchemaClassName  `json:"ranges"`
	Shortcuts                SchemaClassName  `json:"shortcuts"`
	StartPlaceholder         string           `json:"startPlaceholder"`
	EndPlaceholder           string           `json:"endPlaceholder"`
	Animation                bool             `json:"animation"`
	Transform                string           `json:"transform"`
	PopOverContainerSelector string           `json:"popOverContainerSelector"`
}

type DateRangeSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	ValueFormat          string        `json:"valueFormat"`
	Format               string        `json:"format"`
	DisplayFormat        string        `json:"displayFormat"`
	Delimiter            string        `json:"delimiter"`
	Connector            string        `json:"connector"`
}

type DateSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Format               string        `json:"format"`
	DisplayFormat        string        `json:"displayFormat"`
	Placeholder          string        `json:"placeholder"`
	ValueFormat          string        `json:"valueFormat"`
	FromNow              bool          `json:"fromNow"`
	UpdateFrequency      float64       `json:"updateFrequency"`
	DisplayTimeZone      string        `json:"displayTimeZone"`
}

type DateTimeControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Clearable            bool             `json:"clearable"`
	Format               string           `json:"format"`
	ValueFormat          string           `json:"valueFormat"`
	InputFormat          string           `json:"inputFormat"`
	DisplayFormat        string           `json:"displayFormat"`
	Utc                  bool             `json:"utc"`
	Emebed               bool             `json:"emebed"`
	BorderMode           string           `json:"borderMode"`
	Shortcuts            SchemaClassName  `json:"shortcuts"`
	DisabledDate         string           `json:"disabledDate"`
	TimeFormat           string           `json:"timeFormat"`
	MinDate              string           `json:"minDate"`
	MaxDate              string           `json:"maxDate"`
	TimeConstraints      mock             `json:"timeConstraints"`
	IsEndDate            bool             `json:"isEndDate"`
}

type DialogActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Dialog               className       `json:"dialog"`
	NextCondition        className       `json:"nextCondition"`
	Reload               RootSchema      `json:"reload"`
	Redirect             string          `json:"redirect"`
	Data                 mock            `json:"data"`
}

type DialogSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	InputParams          mock          `json:"inputParams"`
	Actions              []any         `json:"actions"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	CloseOnEsc           bool          `json:"closeOnEsc"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	Name                 RootSchema    `json:"name"`
	Size                 string        `json:"size"`
	Height               string        `json:"height"`
	Width                string        `json:"width"`
	Title                className     `json:"title"`
	Header               RootSchema    `json:"header"`
	HeaderClassName      RootSchema    `json:"headerClassName"`
	Footer               RootSchema    `json:"footer"`
	Confirm              bool          `json:"confirm"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	ShowLoading          bool          `json:"showLoading"`
	Overlay              bool          `json:"overlay"`
	DialogType           string        `json:"dialogType"`
	Draggable            bool          `json:"draggable"`
	Data                 any           `json:"data"`
}

type DialogSchemaBase struct {
	InputParams          mock          `json:"inputParams"`
	Actions              []any         `json:"actions"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	CloseOnEsc           bool          `json:"closeOnEsc"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	Name                 RootSchema    `json:"name"`
	Size                 string        `json:"size"`
	Height               string        `json:"height"`
	Width                string        `json:"width"`
	Title                className     `json:"title"`
	Header               RootSchema    `json:"header"`
	HeaderClassName      RootSchema    `json:"headerClassName"`
	Footer               RootSchema    `json:"footer"`
	Confirm              bool          `json:"confirm"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	ShowLoading          bool          `json:"showLoading"`
	Overlay              bool          `json:"overlay"`
	DialogType           string        `json:"dialogType"`
	Draggable            bool          `json:"draggable"`
	Data                 any           `json:"data"`
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
}

type DiffControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	DiffValue            className        `json:"diffValue"`
	Language             string           `json:"language"`
	Options              mock             `json:"options"`
}

type DividerSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	LineStyle            string        `json:"lineStyle"`
	Direction            string        `json:"direction"`
	Color                string        `json:"color"`
	Rotate               float64       `json:"rotate"`
	Title                RootSchema    `json:"title"`
	TitleClassName       string        `json:"titleClassName"`
	TitlePosition        string        `json:"titlePosition"`
}

type DrawerActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Drawer               className       `json:"drawer"`
	NextCondition        className       `json:"nextCondition"`
	Reload               RootSchema      `json:"reload"`
	Redirect             string          `json:"redirect"`
	Data                 mock            `json:"data"`
}

type DrawerSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	InputParams          mock          `json:"inputParams"`
	Actions              []any         `json:"actions"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	HeaderClassName      className     `json:"headerClassName"`
	FooterClassName      className     `json:"footerClassName"`
	CloseOnEsc           bool          `json:"closeOnEsc"`
	Name                 RootSchema    `json:"name"`
	Size                 string        `json:"size"`
	Title                className     `json:"title"`
	Position             string        `json:"position"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	Width                any           `json:"width"`
	Height               any           `json:"height"`
	Header               className     `json:"header"`
	Footer               className     `json:"footer"`
	Confirm              bool          `json:"confirm"`
	Resizable            bool          `json:"resizable"`
	Overlay              bool          `json:"overlay"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	Data                 any           `json:"data"`
}

type DrawerSchemaBase struct {
	InputParams          mock          `json:"inputParams"`
	Actions              []any         `json:"actions"`
	Body                 className     `json:"body"`
	ClassName            className     `json:"className"`
	BodyClassName        className     `json:"bodyClassName"`
	HeaderClassName      className     `json:"headerClassName"`
	FooterClassName      className     `json:"footerClassName"`
	CloseOnEsc           bool          `json:"closeOnEsc"`
	Name                 RootSchema    `json:"name"`
	Size                 string        `json:"size"`
	Title                className     `json:"title"`
	Position             string        `json:"position"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	Width                any           `json:"width"`
	Height               any           `json:"height"`
	Header               className     `json:"header"`
	Footer               className     `json:"footer"`
	Confirm              bool          `json:"confirm"`
	Resizable            bool          `json:"resizable"`
	Overlay              bool          `json:"overlay"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	Data                 any           `json:"data"`
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
}

type DropdownButton struct {
	AnyOf []ArrayElem16 `json:"anyOf"`
}

type DropdownButtonSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Block                bool          `json:"block"`
	BtnClassName         className     `json:"btnClassName"`
	Buttons              []any         `json:"buttons"`
	Body                 className     `json:"body"`
	Label                string        `json:"label"`
	Level                string        `json:"level"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	CloseOnClick         bool          `json:"closeOnClick"`
	Size                 string        `json:"size"`
	Align                string        `json:"align"`
	IconOnly             bool          `json:"iconOnly"`
	RightIcon            className     `json:"rightIcon"`
	Trigger              string        `json:"trigger"`
	HideCaret            bool          `json:"hideCaret"`
	MenuClassName        string        `json:"menuClassName"`
	OverlayPlacement     string        `json:"overlayPlacement"`
}

type EachSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Name                 string        `json:"name"`
	Source               string        `json:"source"`
	ItemKeyName          string        `json:"itemKeyName"`
	IndexKeyName         string        `json:"indexKeyName"`
	Items                RootSchema    `json:"items"`
	Placeholder          string        `json:"placeholder"`
}

type EditorControlSchema struct {
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               mock             `json:"remark"`
	LabelRemark          mock             `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	TypeVal              string           `json:"type"`
	Language             string           `json:"language"`
	Size                 string           `json:"size"`
	AllowFullscreen      bool             `json:"allowFullscreen"`
	EditorDidMount       string           `json:"editorDidMount"`
}

type EmailActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 string          `json:"body"`
	ActionType           string          `json:"actionType"`
	To                   string          `json:"to"`
	Cc                   string          `json:"cc"`
	Bcc                  string          `json:"bcc"`
	Subject              string          `json:"subject"`
}

type ExpandableSchema struct {
	TypeVal                  string `json:"type"`
	KeyField                 string `json:"keyField"`
	ExpandableOn             string `json:"expandableOn"`
	ExpandedRowClassNameExpr string `json:"expandedRowClassNameExpr"`
	ExpandedRowKeys          []any  `json:"expandedRowKeys"`
	ExpandedRowKeysExpr      string `json:"expandedRowKeysExpr"`
}

type ExpressionField struct {
	TypeVal string `json:"type"`
	Field   string `json:"field"`
}

type ExpressionFormula struct {
	TypeVal string `json:"type"`
	Value   string `json:"value"`
}

type ExpressionFunc struct {
	TypeVal string `json:"type"`
	FuncVal string `json:"func"`
	Args    []any  `json:"args"`
}

type FeedbackDialog struct {
	InputParams          mock          `json:"inputParams"`
	Actions              []any         `json:"actions"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	CloseOnEsc           bool          `json:"closeOnEsc"`
	CloseOnOutside       bool          `json:"closeOnOutside"`
	Name                 RootSchema    `json:"name"`
	Size                 string        `json:"size"`
	Height               string        `json:"height"`
	Width                string        `json:"width"`
	Title                className     `json:"title"`
	Header               RootSchema    `json:"header"`
	HeaderClassName      RootSchema    `json:"headerClassName"`
	Footer               RootSchema    `json:"footer"`
	Confirm              bool          `json:"confirm"`
	ShowCloseButton      bool          `json:"showCloseButton"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	ShowLoading          bool          `json:"showLoading"`
	Overlay              bool          `json:"overlay"`
	DialogType           string        `json:"dialogType"`
	Draggable            bool          `json:"draggable"`
	Data                 any           `json:"data"`
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            string        `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	SkipRestOnCancel     bool          `json:"skipRestOnCancel"`
	SkipRestOnConfirm    bool          `json:"skipRestOnConfirm"`
}

type FieldGroup struct {
	Label    string `json:"label"`
	Children []any  `json:"children"`
}

type FieldSetControlSchema struct {
	Key                  string           `json:"key"`
	HeaderPosition       string           `json:"headerPosition"`
	Header               SchemaClassName  `json:"header"`
	BodyClassName        string           `json:"bodyClassName"`
	Disabled             bool             `json:"disabled"`
	Collapsable          bool             `json:"collapsable"`
	Collapsed            bool             `json:"collapsed"`
	ShowArrow            bool             `json:"showArrow"`
	ExpandIcon           className        `json:"expandIcon"`
	HeadingClassName     string           `json:"headingClassName"`
	CollapseHeader       className        `json:"collapseHeader"`
	Size                 string           `json:"size"`
	MountOnEnter         bool             `json:"mountOnEnter"`
	UnmountOnExit        bool             `json:"unmountOnExit"`
	DivideLine           bool             `json:"divideLine"`
	Testid               string           `json:"testid"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               mock             `json:"remark"`
	LabelRemark          mock             `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	TitlePosition        string           `json:"titlePosition"`
	Body                 className        `json:"body"`
	Title                className        `json:"title"`
	CollapseTitle        className        `json:"collapseTitle"`
	SubFormMode          string           `json:"subFormMode"`
	SubFormHorizontal    className        `json:"subFormHorizontal"`
}

type FieldSimple struct {
	AnyOf []ArrayElem25 `json:"anyOf"`
}

type Field_161becb8 struct {
	Prototype staticSchema `json:"prototype"`
	Length    float64      `json:"length"`
	Arguments staticSchema `json:"arguments"`
	Caller    RootSchema   `json:"caller"`
}

type FileControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             any              `json:"autoFill"`
	InitAutoFill         bool             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	BtnLabel             string           `json:"btnLabel"`
	Accept               string           `json:"accept"`
	Capture              string           `json:"capture"`
	AsBase64             bool             `json:"asBase64"`
	AsBlob               bool             `json:"asBlob"`
	AutoUpload           bool             `json:"autoUpload"`
	ChunkApi             className        `json:"chunkApi"`
	ChunkSize            float64          `json:"chunkSize"`
	Concurrency          float64          `json:"concurrency"`
	Delimiter            string           `json:"delimiter"`
	DownloadUrl          className        `json:"downloadUrl"`
	TemplateUrl          className        `json:"templateUrl"`
	FileField            string           `json:"fileField"`
	FinishChunkApi       className        `json:"finishChunkApi"`
	HideUploadButton     bool             `json:"hideUploadButton"`
	MaxLength            float64          `json:"maxLength"`
	MaxSize              float64          `json:"maxSize"`
	Receiver             className        `json:"receiver"`
	StartChunkApi        string           `json:"startChunkApi"`
	UseChunk             columnsTogglable `json:"useChunk"`
	BtnClassName         className        `json:"btnClassName"`
	BtnUploadClassName   className        `json:"btnUploadClassName"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	ExtractValue         bool             `json:"extractValue"`
	ResetValue           mock             `json:"resetValue"`
	ValueField           string           `json:"valueField"`
	NameField            string           `json:"nameField"`
	UrlField             string           `json:"urlField"`
	StateTextMap         stateTextMap     `json:"stateTextMap"`
	Documentation        string           `json:"documentation"`
	DocumentLink         string           `json:"documentLink"`
	Drag                 bool             `json:"drag"`
}

type FlexSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Justify              string        `json:"justify"`
	AlignItems           string        `json:"alignItems"`
	AlignContent         string        `json:"alignContent"`
	Direction            string        `json:"direction"`
	Items                className     `json:"items"`
}

type FormControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Body                 className        `json:"body"`
}

type FormHorizontal struct {
	Left       float64 `json:"left"`
	Right      float64 `json:"right"`
	LeftFixed  data    `json:"leftFixed"`
	Justify    bool    `json:"justify"`
	LabelAlign string  `json:"labelAlign"`
	LabelWidth any     `json:"labelWidth"`
}

type FormSchema struct {
	ID                          string          `json:"$$id"`
	ClassName                   className       `json:"className"`
	Ref                         string          `json:"$ref"`
	Disabled                    bool            `json:"disabled"`
	DisabledOn                  className       `json:"disabledOn"`
	Hidden                      bool            `json:"hidden"`
	HiddenOn                    className       `json:"hiddenOn"`
	Visible                     bool            `json:"visible"`
	VisibleOn                   className       `json:"visibleOn"`
	Id                          string          `json:"id"`
	OnEvent                     any             `json:"onEvent"`
	Static                      bool            `json:"static"`
	StaticOn                    RootSchema      `json:"staticOn"`
	StaticPlaceholder           string          `json:"staticPlaceholder"`
	StaticClassName             RootSchema      `json:"staticClassName"`
	StaticLabelClassName        className       `json:"staticLabelClassName"`
	StaticInputClassName        className       `json:"staticInputClassName"`
	StaticSchema                staticSchema    `json:"staticSchema"`
	Style                       any             `json:"style"`
	EditorSetting               editorSetting   `json:"editorSetting"`
	UseMobileUI                 bool            `json:"useMobileUI"`
	TestIdBuilder               RootSchema      `json:"testIdBuilder"`
	TypeVal                     string          `json:"type"`
	Testid                      string          `json:"testid"`
	Title                       string          `json:"title"`
	Actions                     []any           `json:"actions"`
	Body                        className       `json:"body"`
	Tabs                        staticSchema    `json:"tabs"`
	FieldSet                    staticSchema    `json:"fieldSet"`
	Data                        staticSchema    `json:"data"`
	Debug                       bool            `json:"debug"`
	DebugConfig                 debugConfig     `json:"debugConfig"`
	InitApi                     SchemaClassName `json:"initApi"`
	InitAsyncApi                SchemaClassName `json:"initAsyncApi"`
	InitFinishedField           string          `json:"initFinishedField"`
	InitCheckInterval           float64         `json:"initCheckInterval"`
	InitFetch                   bool            `json:"initFetch"`
	InitFetchOn                 string          `json:"initFetchOn"`
	Interval                    float64         `json:"interval"`
	SilentPolling               bool            `json:"silentPolling"`
	StopAutoRefreshWhen         string          `json:"stopAutoRefreshWhen"`
	PersistData                 string          `json:"persistData"`
	PersistDataKeys             []string        `json:"persistDataKeys"`
	ClearPersistDataAfterSubmit bool            `json:"clearPersistDataAfterSubmit"`
	Api                         SchemaClassName `json:"api"`
	Feedback                    mock            `json:"feedback"`
	AsyncApi                    SchemaClassName `json:"asyncApi"`
	CheckInterval               float64         `json:"checkInterval"`
	FinishedField               string          `json:"finishedField"`
	ResetAfterSubmit            bool            `json:"resetAfterSubmit"`
	ClearAfterSubmit            bool            `json:"clearAfterSubmit"`
	Mode                        string          `json:"mode"`
	ColumnCount                 float64         `json:"columnCount"`
	Horizontal                  className       `json:"horizontal"`
	AutoFocus                   bool            `json:"autoFocus"`
	Messages                    messages        `json:"messages"`
	Name                        string          `json:"name"`
	PanelClassName              className       `json:"panelClassName"`
	PrimaryField                string          `json:"primaryField"`
	Redirect                    string          `json:"redirect"`
	Reload                      string          `json:"reload"`
	SubmitOnChange              bool            `json:"submitOnChange"`
	SubmitOnInit                bool            `json:"submitOnInit"`
	SubmitText                  string          `json:"submitText"`
	Target                      string          `json:"target"`
	WrapWithPanel               bool            `json:"wrapWithPanel"`
	AffixFooter                 bool            `json:"affixFooter"`
	PromptPageLeave             bool            `json:"promptPageLeave"`
	PromptPageLeaveMessage      string          `json:"promptPageLeaveMessage"`
	Rules                       []ArrayItem3    `json:"rules"`
	PreventEnterSubmit          bool            `json:"preventEnterSubmit"`
	LabelAlign                  className       `json:"labelAlign"`
	LabelWidth                  any             `json:"labelWidth"`
}

type FormulaControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Condition            string           `json:"condition"`
	AutoSet              bool             `json:"autoSet"`
	Formula              string           `json:"formula"`
	InitSet              bool             `json:"initSet"`
}

type FormulaPickerInputSettings struct {
	TypeVal RootSchema `json:"type"`
}

type FuncGroup struct {
	GroupName string `json:"groupName"`
	Items     []any  `json:"items"`
}

type FuncItem struct {
	Name        string `json:"name"`
	Example     string `json:"example"`
	Description string `json:"description"`
}

type Grid struct {
	AllOf []ArrayElem17 `json:"allOf"`
}

type Grid2DSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Cols                 float64         `json:"cols"`
	Width                SchemaClassName `json:"width"`
	Gap                  any             `json:"gap"`
	GapRow               any             `json:"gapRow"`
	RowHeight            any             `json:"rowHeight"`
	Grids                []any           `json:"grids"`
}

type GridColumnObject struct {
	Id                 string          `json:"id"`
	Xs                 SchemaClassName `json:"xs"`
	Sm                 SchemaClassName `json:"sm"`
	Md                 SchemaClassName `json:"md"`
	Lg                 SchemaClassName `json:"lg"`
	Valign             string          `json:"valign"`
	Mode               string          `json:"mode"`
	Horizontal         className       `json:"horizontal"`
	Body               RootSchema      `json:"body"`
	ColumnClassName    className       `json:"columnClassName"`
	Style              mock            `json:"style"`
	WrapperCustomStyle staticSchema    `json:"wrapperCustomStyle"`
	ThemeCss           staticSchema    `json:"themeCss"`
}

type GridSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Columns              []any         `json:"columns"`
	Gap                  string        `json:"gap"`
	Valign               string        `json:"valign"`
	Align                string        `json:"align"`
}

type GroupControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Body                 []any            `json:"body"`
	Gap                  string           `json:"gap"`
	Direction            string           `json:"direction"`
	SubFormMode          string           `json:"subFormMode"`
	SubFormHorizontal    className        `json:"subFormHorizontal"`
}

type GroupSubControl struct {
	AllOf []ArrayElem27 `json:"allOf"`
}

type HBoxColumnObject struct {
	ColumnClassName string    `json:"columnClassName"`
	Valign          string    `json:"valign"`
	Width           any       `json:"width"`
	Height          any       `json:"height"`
	Style           any       `json:"style"`
	Mode            string    `json:"mode"`
	Horizontal      className `json:"horizontal"`
	Body            className `json:"body"`
	Visible         bool      `json:"visible"`
	VisibleOn       className `json:"visibleOn"`
}

type HBoxSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Columns              []any         `json:"columns"`
	SubFormMode          string        `json:"subFormMode"`
	SubFormHorizontal    className     `json:"subFormHorizontal"`
	Gap                  string        `json:"gap"`
	Valign               string        `json:"valign"`
	Align                string        `json:"align"`
}

type HboxRow struct {
	AllOf []ArrayElem20 `json:"allOf"`
}

type HiddenControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
}

type IFrameSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Src                  className     `json:"src"`
	Events               any           `json:"events"`
	Width                any           `json:"width"`
	Height               any           `json:"height"`
	Allow                string        `json:"allow"`
	Name                 string        `json:"name"`
	Referrerpolicy       string        `json:"referrerpolicy"`
	Sandbox              string        `json:"sandbox"`
}

type IconCheckedSchema struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Svg  string `json:"svg"`
}

type IconItemSchema struct {
	Icon     RootSchema `json:"icon"`
	Position string     `json:"position"`
}

type IconSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Icon                 draggableConfig `json:"icon"`
	Vendor               string          `json:"vendor"`
	Badge                className       `json:"badge"`
}

type ImageControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             any              `json:"autoFill"`
	InitAutoFill         bool             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Src                  className        `json:"src"`
	ImageClassName       string           `json:"imageClassName"`
	Accept               string           `json:"accept"`
	AllowInput           bool             `json:"allowInput"`
	AutoUpload           bool             `json:"autoUpload"`
	UploadBtnText        SchemaClassName  `json:"uploadBtnText"`
	BtnClassName         className        `json:"btnClassName"`
	BtnUploadClassName   className        `json:"btnUploadClassName"`
	Compress             bool             `json:"compress"`
	CompressOptions      compressOptions  `json:"compressOptions"`
	Crop                 data             `json:"crop"`
	CropFormat           string           `json:"cropFormat"`
	CropQuality          float64          `json:"cropQuality"`
	ReCropable           bool             `json:"reCropable"`
	HideUploadButton     bool             `json:"hideUploadButton"`
	Limit                limit            `json:"limit"`
	MaxLength            float64          `json:"maxLength"`
	MaxSize              float64          `json:"maxSize"`
	Receiver             className        `json:"receiver"`
	ShowCompressOptions  bool             `json:"showCompressOptions"`
	Multiple             bool             `json:"multiple"`
	Capture              string           `json:"capture"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ExtractValue         bool             `json:"extractValue"`
	ResetValue           mock             `json:"resetValue"`
	ThumbMode            string           `json:"thumbMode"`
	ThumbRatio           string           `json:"thumbRatio"`
	InitCrop             bool             `json:"initCrop"`
	DropCrop             bool             `json:"dropCrop"`
	FrameImage           className        `json:"frameImage"`
	FixedSize            bool             `json:"fixedSize"`
	FixedSizeClassName   className        `json:"fixedSizeClassName"`
	Draggable            bool             `json:"draggable"`
	DraggableTip         string           `json:"draggableTip"`
}

type ImageSchema struct {
	ID                    string        `json:"$$id"`
	ClassName             className     `json:"className"`
	Ref                   string        `json:"$ref"`
	Disabled              bool          `json:"disabled"`
	DisabledOn            className     `json:"disabledOn"`
	Hidden                bool          `json:"hidden"`
	HiddenOn              className     `json:"hiddenOn"`
	Visible               bool          `json:"visible"`
	VisibleOn             className     `json:"visibleOn"`
	Id                    string        `json:"id"`
	OnEvent               any           `json:"onEvent"`
	Static                bool          `json:"static"`
	StaticOn              className     `json:"staticOn"`
	StaticPlaceholder     string        `json:"staticPlaceholder"`
	StaticClassName       className     `json:"staticClassName"`
	StaticLabelClassName  className     `json:"staticLabelClassName"`
	StaticInputClassName  className     `json:"staticInputClassName"`
	StaticSchema          staticSchema  `json:"staticSchema"`
	Style                 any           `json:"style"`
	EditorSetting         editorSetting `json:"editorSetting"`
	UseMobileUI           bool          `json:"useMobileUI"`
	TestIdBuilder         RootSchema    `json:"testIdBuilder"`
	TypeVal               string        `json:"type"`
	Testid                string        `json:"testid"`
	DefaultImage          className     `json:"defaultImage"`
	Title                 className     `json:"title"`
	Name                  string        `json:"name"`
	ImageCaption          className     `json:"imageCaption"`
	Src                   className     `json:"src"`
	OriginalSrc           className     `json:"originalSrc"`
	EnlargeAble           bool          `json:"enlargeAble"`
	EnlargeWithGallary    bool          `json:"enlargeWithGallary"`
	Alt                   string        `json:"alt"`
	Height                float64       `json:"height"`
	Width                 float64       `json:"width"`
	InnerClassName        className     `json:"innerClassName"`
	ImageClassName        className     `json:"imageClassName"`
	ThumbClassName        className     `json:"thumbClassName"`
	ImageGallaryClassName className     `json:"imageGallaryClassName"`
	Caption               className     `json:"caption"`
	ImageMode             string        `json:"imageMode"`
	ThumbMode             string        `json:"thumbMode"`
	ThumbRatio            string        `json:"thumbRatio"`
	Href                  className     `json:"href"`
	Blank                 bool          `json:"blank"`
	HtmlTarget            string        `json:"htmlTarget"`
	ShowToolbar           bool          `json:"showToolbar"`
	ToolbarActions        []any         `json:"toolbarActions"`
}

type ImageToolbarAction struct {
	Key           string `json:"key"`
	Label         string `json:"label"`
	Icon          string `json:"icon"`
	IconClassName string `json:"iconClassName"`
	Disabled      bool   `json:"disabled"`
}

type ImagesSchema struct {
	ID                    string        `json:"$$id"`
	ClassName             className     `json:"className"`
	Ref                   string        `json:"$ref"`
	Disabled              bool          `json:"disabled"`
	DisabledOn            className     `json:"disabledOn"`
	Hidden                bool          `json:"hidden"`
	HiddenOn              className     `json:"hiddenOn"`
	Visible               bool          `json:"visible"`
	VisibleOn             className     `json:"visibleOn"`
	Id                    string        `json:"id"`
	OnEvent               any           `json:"onEvent"`
	Static                bool          `json:"static"`
	StaticOn              className     `json:"staticOn"`
	StaticPlaceholder     string        `json:"staticPlaceholder"`
	StaticClassName       className     `json:"staticClassName"`
	StaticLabelClassName  className     `json:"staticLabelClassName"`
	StaticInputClassName  className     `json:"staticInputClassName"`
	StaticSchema          staticSchema  `json:"staticSchema"`
	Style                 any           `json:"style"`
	EditorSetting         editorSetting `json:"editorSetting"`
	UseMobileUI           bool          `json:"useMobileUI"`
	TestIdBuilder         RootSchema    `json:"testIdBuilder"`
	TypeVal               string        `json:"type"`
	Testid                string        `json:"testid"`
	DefaultImage          className     `json:"defaultImage"`
	Placeholder           string        `json:"placeholder"`
	Delimiter             string        `json:"delimiter"`
	ThumbMode             string        `json:"thumbMode"`
	ThumbRatio            string        `json:"thumbRatio"`
	Name                  string        `json:"name"`
	Value                 staticSchema  `json:"value"`
	Source                string        `json:"source"`
	Options               []any         `json:"options"`
	Src                   string        `json:"src"`
	OriginalSrc           string        `json:"originalSrc"`
	EnlargeAble           bool          `json:"enlargeAble"`
	EnlargetWithImages    bool          `json:"enlargetWithImages"`
	ShowDimensions        bool          `json:"showDimensions"`
	ListClassName         className     `json:"listClassName"`
	ImageGallaryClassName className     `json:"imageGallaryClassName"`
	ShowToolbar           bool          `json:"showToolbar"`
	ToolbarActions        []any         `json:"toolbarActions"`
}

type InputCityControlSchema struct {
	LoadingConfig        loadingConfig    `json:"loadingConfig"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	ExtractValue         bool             `json:"extractValue"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	AllowCity            bool             `json:"allowCity"`
	AllowDistrict        bool             `json:"allowDistrict"`
	AllowStreet          bool             `json:"allowStreet"`
	Searchable           bool             `json:"searchable"`
	ItemClassName        string           `json:"itemClassName"`
}

type InputColorControlSchema struct {
	ID                       string           `json:"$$id"`
	ClassName                className        `json:"className"`
	Ref                      string           `json:"$ref"`
	Disabled                 bool             `json:"disabled"`
	DisabledOn               className        `json:"disabledOn"`
	Hidden                   bool             `json:"hidden"`
	HiddenOn                 className        `json:"hiddenOn"`
	Visible                  bool             `json:"visible"`
	VisibleOn                className        `json:"visibleOn"`
	Id                       string           `json:"id"`
	OnEvent                  any              `json:"onEvent"`
	Static                   bool             `json:"static"`
	StaticOn                 className        `json:"staticOn"`
	StaticPlaceholder        string           `json:"staticPlaceholder"`
	StaticClassName          className        `json:"staticClassName"`
	StaticLabelClassName     className        `json:"staticLabelClassName"`
	StaticInputClassName     className        `json:"staticInputClassName"`
	StaticSchema             staticSchema     `json:"staticSchema"`
	Style                    any              `json:"style"`
	EditorSetting            editorSetting    `json:"editorSetting"`
	UseMobileUI              bool             `json:"useMobileUI"`
	TestIdBuilder            RootSchema       `json:"testIdBuilder"`
	Size                     string           `json:"size"`
	Label                    SchemaClassName  `json:"label"`
	LabelAlign               className        `json:"labelAlign"`
	LabelWidth               any              `json:"labelWidth"`
	LabelClassName           string           `json:"labelClassName"`
	Name                     string           `json:"name"`
	ExtraName                string           `json:"extraName"`
	Remark                   className        `json:"remark"`
	LabelRemark              className        `json:"labelRemark"`
	Hint                     string           `json:"hint"`
	SubmitOnChange           bool             `json:"submitOnChange"`
	ReadOnly                 bool             `json:"readOnly"`
	ReadOnlyOn               string           `json:"readOnlyOn"`
	ValidateOnChange         bool             `json:"validateOnChange"`
	Description              string           `json:"description"`
	Desc                     string           `json:"desc"`
	DescriptionClassName     className        `json:"descriptionClassName"`
	Mode                     string           `json:"mode"`
	Horizontal               className        `json:"horizontal"`
	Inline                   bool             `json:"inline"`
	InputClassName           className        `json:"inputClassName"`
	Placeholder              string           `json:"placeholder"`
	Required                 bool             `json:"required"`
	ValidationErrors         validationErrors `json:"validationErrors"`
	Validations              data             `json:"validations"`
	Value                    mock             `json:"value"`
	ClearValueOnHidden       bool             `json:"clearValueOnHidden"`
	ValidateApi              SchemaClassName  `json:"validateApi"`
	AutoFill                 autoFill         `json:"autoFill"`
	InitAutoFill             data             `json:"initAutoFill"`
	Row                      float64          `json:"row"`
	TypeVal                  string           `json:"type"`
	Clearable                bool             `json:"clearable"`
	Format                   string           `json:"format"`
	CloseOnSelect            bool             `json:"closeOnSelect"`
	PresetColors             []any            `json:"presetColors"`
	AllowCustomColor         bool             `json:"allowCustomColor"`
	PopOverContainerSelector string           `json:"popOverContainerSelector"`
}

type InputGroupControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Body                 className        `json:"body"`
	ValidationConfig     validationConfig `json:"validationConfig"`
}

type InputSignatureSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Width                float64          `json:"width"`
	Height               float64          `json:"height"`
	Color                string           `json:"color"`
	BgColor              string           `json:"bgColor"`
	ClearBtnLabel        string           `json:"clearBtnLabel"`
	ClearBtnIcon         string           `json:"clearBtnIcon"`
	UndoBtnLabel         string           `json:"undoBtnLabel"`
	UndoBtnIcon          string           `json:"undoBtnIcon"`
	ConfirmBtnLabel      string           `json:"confirmBtnLabel"`
	ConfirmBtnIcon       string           `json:"confirmBtnIcon"`
	Embed                bool             `json:"embed"`
	EmbedConfirmLabel    string           `json:"embedConfirmLabel"`
	EmbedConfirmIcon     string           `json:"embedConfirmIcon"`
	EbmedCancelLabel     string           `json:"ebmedCancelLabel"`
	EbmedCancelIcon      string           `json:"ebmedCancelIcon"`
	EmbedBtnIcon         string           `json:"embedBtnIcon"`
	EmbedBtnLabel        string           `json:"embedBtnLabel"`
}

type JSONSchemaEditorControlSchema struct {
	TypeVal               string           `json:"type"`
	Remark                className        `json:"remark"`
	LabelRemark           className        `json:"labelRemark"`
	Size                  string           `json:"size"`
	Label                 SchemaClassName  `json:"label"`
	LabelAlign            className        `json:"labelAlign"`
	LabelWidth            any              `json:"labelWidth"`
	LabelClassName        string           `json:"labelClassName"`
	Name                  string           `json:"name"`
	ExtraName             string           `json:"extraName"`
	Hint                  string           `json:"hint"`
	SubmitOnChange        bool             `json:"submitOnChange"`
	ReadOnly              bool             `json:"readOnly"`
	ReadOnlyOn            string           `json:"readOnlyOn"`
	ValidateOnChange      bool             `json:"validateOnChange"`
	Description           string           `json:"description"`
	Desc                  string           `json:"desc"`
	DescriptionClassName  className        `json:"descriptionClassName"`
	Mode                  string           `json:"mode"`
	Horizontal            className        `json:"horizontal"`
	Inline                bool             `json:"inline"`
	InputClassName        className        `json:"inputClassName"`
	Required              bool             `json:"required"`
	ValidationErrors      validationErrors `json:"validationErrors"`
	Validations           data             `json:"validations"`
	Value                 mock             `json:"value"`
	ClearValueOnHidden    bool             `json:"clearValueOnHidden"`
	ValidateApi           SchemaClassName  `json:"validateApi"`
	AutoFill              autoFill         `json:"autoFill"`
	InitAutoFill          data             `json:"initAutoFill"`
	Row                   float64          `json:"row"`
	ID                    string           `json:"$$id"`
	ClassName             className        `json:"className"`
	Ref                   string           `json:"$ref"`
	Disabled              bool             `json:"disabled"`
	DisabledOn            className        `json:"disabledOn"`
	Hidden                bool             `json:"hidden"`
	HiddenOn              className        `json:"hiddenOn"`
	Visible               bool             `json:"visible"`
	VisibleOn             className        `json:"visibleOn"`
	Id                    string           `json:"id"`
	OnEvent               any              `json:"onEvent"`
	Static                bool             `json:"static"`
	StaticOn              className        `json:"staticOn"`
	StaticPlaceholder     string           `json:"staticPlaceholder"`
	StaticClassName       className        `json:"staticClassName"`
	StaticLabelClassName  className        `json:"staticLabelClassName"`
	StaticInputClassName  className        `json:"staticInputClassName"`
	StaticSchema          staticSchema     `json:"staticSchema"`
	Style                 any              `json:"style"`
	EditorSetting         editorSetting    `json:"editorSetting"`
	UseMobileUI           bool             `json:"useMobileUI"`
	TestIdBuilder         RootSchema       `json:"testIdBuilder"`
	Definitions           any              `json:"definitions"`
	RootTypeMutable       bool             `json:"rootTypeMutable"`
	ShowRootInfo          bool             `json:"showRootInfo"`
	DisabledTypes         []string         `json:"disabledTypes"`
	EnableAdvancedSetting bool             `json:"enableAdvancedSetting"`
	AdvancedSettings      any              `json:"advancedSettings"`
	Placeholder           className        `json:"placeholder"`
	Mini                  bool             `json:"mini"`
}

type JsonSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Value                SchemaClassName `json:"value"`
	LevelExpand          float64         `json:"levelExpand"`
	Source               string          `json:"source"`
	Mutable              bool            `json:"mutable"`
	DisplayDataTypes     bool            `json:"displayDataTypes"`
	EnableClipboard      bool            `json:"enableClipboard"`
	IconStyle            string          `json:"iconStyle"`
	QuotesOnKeys         bool            `json:"quotesOnKeys"`
	SortKeys             bool            `json:"sortKeys"`
	EllipsisThreshold    SchemaClassName `json:"ellipsisThreshold"`
}

type LinkActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Link                 string          `json:"link"`
}

type LinkSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Blank                bool          `json:"blank"`
	Href                 string        `json:"href"`
	Body                 className     `json:"body"`
	Badge                className     `json:"badge"`
	HtmlTarget           string        `json:"htmlTarget"`
	Icon                 string        `json:"icon"`
	RightIcon            string        `json:"rightIcon"`
}

type ListBodyField struct {
	AllOf []ArrayElem14 `json:"allOf"`
}

type ListBodyFieldObject struct {
	Label          string    `json:"label"`
	LabelClassName className `json:"labelClassName"`
	InnerClassName className `json:"innerClassName"`
	Name           string    `json:"name"`
	PopOver        className `json:"popOver"`
	QuickEdit      className `json:"quickEdit"`
	Copyable       className `json:"copyable"`
}

type ListControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	SubmitOnDBClick      bool             `json:"submitOnDBClick"`
	ImageClassName       string           `json:"imageClassName"`
	ItemSchema           className        `json:"itemSchema"`
	ActiveItemSchema     className        `json:"activeItemSchema"`
	ListClassName        className        `json:"listClassName"`
}

type ListItemSchema struct {
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	Actions              []any         `json:"actions"`
	ActionsPosition      string        `json:"actionsPosition"`
	Avatar               className     `json:"avatar"`
	Body                 []any         `json:"body"`
	Desc                 className     `json:"desc"`
	Remark               className     `json:"remark"`
	Title                className     `json:"title"`
	SubTitle             className     `json:"subTitle"`
}

type ListSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Title                className     `json:"title"`
	Footer               className     `json:"footer"`
	FooterClassName      className     `json:"footerClassName"`
	Header               className     `json:"header"`
	HeaderClassName      className     `json:"headerClassName"`
	ListItem             className     `json:"listItem"`
	Source               className     `json:"source"`
	ShowFooter           bool          `json:"showFooter"`
	ShowHeader           bool          `json:"showHeader"`
	Placeholder          className     `json:"placeholder"`
	HideCheckToggler     bool          `json:"hideCheckToggler"`
	AffixHeader          bool          `json:"affixHeader"`
	AffixFooter          bool          `json:"affixFooter"`
	ItemCheckableOn      className     `json:"itemCheckableOn"`
	ItemDraggableOn      className     `json:"itemDraggableOn"`
	CheckOnItemClick     bool          `json:"checkOnItemClick"`
	ValueField           string        `json:"valueField"`
	Size                 string        `json:"size"`
	ItemAction           className     `json:"itemAction"`
}

type ListenerAction struct {
	ActionType      string `json:"actionType"`
	Description     string `json:"description"`
	ComponentId     string `json:"componentId"`
	ComponentName   string `json:"componentName"`
	IgnoreError     bool   `json:"ignoreError"`
	Args            any    `json:"args"`
	Data            data   `json:"data"`
	DataMergeMode   string `json:"dataMergeMode"`
	OutputVar       string `json:"outputVar"`
	PreventDefault  bool   `json:"preventDefault"`
	StopPropagation bool   `json:"stopPropagation"`
	Expression      data   `json:"expression"`
	ExecOn          string `json:"execOn"`
}

type LocationControlSchema struct {
	ID                     string           `json:"$$id"`
	ClassName              className        `json:"className"`
	Ref                    string           `json:"$ref"`
	Disabled               bool             `json:"disabled"`
	DisabledOn             className        `json:"disabledOn"`
	Hidden                 bool             `json:"hidden"`
	HiddenOn               className        `json:"hiddenOn"`
	Visible                bool             `json:"visible"`
	VisibleOn              className        `json:"visibleOn"`
	Id                     string           `json:"id"`
	OnEvent                any              `json:"onEvent"`
	Static                 bool             `json:"static"`
	StaticOn               className        `json:"staticOn"`
	StaticPlaceholder      string           `json:"staticPlaceholder"`
	StaticClassName        className        `json:"staticClassName"`
	StaticLabelClassName   className        `json:"staticLabelClassName"`
	StaticInputClassName   className        `json:"staticInputClassName"`
	StaticSchema           staticSchema     `json:"staticSchema"`
	Style                  any              `json:"style"`
	EditorSetting          editorSetting    `json:"editorSetting"`
	UseMobileUI            bool             `json:"useMobileUI"`
	TestIdBuilder          RootSchema       `json:"testIdBuilder"`
	Size                   string           `json:"size"`
	Label                  SchemaClassName  `json:"label"`
	LabelAlign             className        `json:"labelAlign"`
	LabelWidth             any              `json:"labelWidth"`
	LabelClassName         string           `json:"labelClassName"`
	Name                   string           `json:"name"`
	ExtraName              string           `json:"extraName"`
	Remark                 className        `json:"remark"`
	LabelRemark            className        `json:"labelRemark"`
	Hint                   string           `json:"hint"`
	SubmitOnChange         bool             `json:"submitOnChange"`
	ReadOnly               bool             `json:"readOnly"`
	ReadOnlyOn             string           `json:"readOnlyOn"`
	ValidateOnChange       bool             `json:"validateOnChange"`
	Description            string           `json:"description"`
	Desc                   string           `json:"desc"`
	DescriptionClassName   className        `json:"descriptionClassName"`
	Mode                   string           `json:"mode"`
	Horizontal             className        `json:"horizontal"`
	Inline                 bool             `json:"inline"`
	InputClassName         className        `json:"inputClassName"`
	Placeholder            string           `json:"placeholder"`
	Required               bool             `json:"required"`
	ValidationErrors       validationErrors `json:"validationErrors"`
	Validations            data             `json:"validations"`
	Value                  mock             `json:"value"`
	ClearValueOnHidden     bool             `json:"clearValueOnHidden"`
	ValidateApi            SchemaClassName  `json:"validateApi"`
	AutoFill               autoFill         `json:"autoFill"`
	InitAutoFill           data             `json:"initAutoFill"`
	Row                    float64          `json:"row"`
	TypeVal                string           `json:"type"`
	Vendor                 string           `json:"vendor"`
	Ak                     string           `json:"ak"`
	AutoSelectCurrentLoc   bool             `json:"autoSelectCurrentLoc"`
	OnlySelectCurrentLoc   bool             `json:"onlySelectCurrentLoc"`
	GetLocationPlaceholder string           `json:"getLocationPlaceholder"`
}

type MappingSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Name                 string          `json:"name"`
	MapVal               any             `json:"map"`
	ValueField           string          `json:"valueField"`
	LabelField           string          `json:"labelField"`
	ItemSchema           className       `json:"itemSchema"`
	Source               draggableConfig `json:"source"`
	Placeholder          string          `json:"placeholder"`
}

type MatrixControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Multiple             bool             `json:"multiple"`
	SingleSelectMode     bool             `json:"singleSelectMode"`
	Source               className        `json:"source"`
	Columns              []ArrayItem4     `json:"columns"`
	Rows                 []ArrayItem4     `json:"rows"`
	RowLabel             string           `json:"rowLabel"`
}

type MonthControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Clearable            bool             `json:"clearable"`
	Format               string           `json:"format"`
	ValueFormat          string           `json:"valueFormat"`
	InputFormat          string           `json:"inputFormat"`
	DisplayFormat        string           `json:"displayFormat"`
	Utc                  bool             `json:"utc"`
	Emebed               bool             `json:"emebed"`
	BorderMode           string           `json:"borderMode"`
	Shortcuts            SchemaClassName  `json:"shortcuts"`
	DisabledDate         string           `json:"disabledDate"`
}

type MonthRangeControlSchema struct {
	Delimiter                string           `json:"delimiter"`
	Format                   string           `json:"format"`
	ValueFormat              string           `json:"valueFormat"`
	InputFormat              string           `json:"inputFormat"`
	DisplayFormat            string           `json:"displayFormat"`
	JoinValues               bool             `json:"joinValues"`
	MaxDate                  string           `json:"maxDate"`
	MinDate                  string           `json:"minDate"`
	MaxDuration              string           `json:"maxDuration"`
	MinDuration              string           `json:"minDuration"`
	Value                    mock             `json:"value"`
	BorderMode               string           `json:"borderMode"`
	Embed                    bool             `json:"embed"`
	Ranges                   SchemaClassName  `json:"ranges"`
	Shortcuts                SchemaClassName  `json:"shortcuts"`
	StartPlaceholder         string           `json:"startPlaceholder"`
	EndPlaceholder           string           `json:"endPlaceholder"`
	Animation                bool             `json:"animation"`
	Transform                string           `json:"transform"`
	PopOverContainerSelector string           `json:"popOverContainerSelector"`
	Remark                   className        `json:"remark"`
	LabelRemark              className        `json:"labelRemark"`
	Size                     string           `json:"size"`
	Label                    SchemaClassName  `json:"label"`
	LabelAlign               className        `json:"labelAlign"`
	LabelWidth               any              `json:"labelWidth"`
	LabelClassName           string           `json:"labelClassName"`
	Name                     string           `json:"name"`
	ExtraName                string           `json:"extraName"`
	Hint                     string           `json:"hint"`
	SubmitOnChange           bool             `json:"submitOnChange"`
	ReadOnly                 bool             `json:"readOnly"`
	ReadOnlyOn               string           `json:"readOnlyOn"`
	ValidateOnChange         bool             `json:"validateOnChange"`
	Description              string           `json:"description"`
	Desc                     string           `json:"desc"`
	DescriptionClassName     className        `json:"descriptionClassName"`
	Mode                     string           `json:"mode"`
	Horizontal               className        `json:"horizontal"`
	Inline                   bool             `json:"inline"`
	InputClassName           className        `json:"inputClassName"`
	Placeholder              string           `json:"placeholder"`
	Required                 bool             `json:"required"`
	ValidationErrors         validationErrors `json:"validationErrors"`
	Validations              data             `json:"validations"`
	ClearValueOnHidden       bool             `json:"clearValueOnHidden"`
	ValidateApi              SchemaClassName  `json:"validateApi"`
	AutoFill                 autoFill         `json:"autoFill"`
	InitAutoFill             data             `json:"initAutoFill"`
	Row                      float64          `json:"row"`
	ID                       string           `json:"$$id"`
	ClassName                className        `json:"className"`
	Ref                      string           `json:"$ref"`
	Disabled                 bool             `json:"disabled"`
	DisabledOn               className        `json:"disabledOn"`
	Hidden                   bool             `json:"hidden"`
	HiddenOn                 className        `json:"hiddenOn"`
	Visible                  bool             `json:"visible"`
	VisibleOn                className        `json:"visibleOn"`
	Id                       string           `json:"id"`
	OnEvent                  any              `json:"onEvent"`
	Static                   bool             `json:"static"`
	StaticOn                 className        `json:"staticOn"`
	StaticPlaceholder        string           `json:"staticPlaceholder"`
	StaticClassName          className        `json:"staticClassName"`
	StaticLabelClassName     className        `json:"staticLabelClassName"`
	StaticInputClassName     className        `json:"staticInputClassName"`
	StaticSchema             staticSchema     `json:"staticSchema"`
	Style                    any              `json:"style"`
	EditorSetting            editorSetting    `json:"editorSetting"`
	UseMobileUI              bool             `json:"useMobileUI"`
	TestIdBuilder            RootSchema       `json:"testIdBuilder"`
	TypeVal                  string           `json:"type"`
}

type MultilineTextSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Text                 string        `json:"text"`
	MaxRows              float64       `json:"maxRows"`
	ExpendButtonText     string        `json:"expendButtonText"`
	CollapseButtonText   string        `json:"collapseButtonText"`
}

type MultipleValue struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type NavItemSchema struct {
	Testid               string          `json:"testid"`
	ID                   string          `json:"$$id"`
	ClassName            string          `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	Label                SchemaClassName `json:"label"`
	Icon                 draggableConfig `json:"icon"`
	To                   RootSchema      `json:"to"`
	Target               string          `json:"target"`
	Unfolded             bool            `json:"unfolded"`
	Active               bool            `json:"active"`
	DeferVal             bool            `json:"defer"`
	DeferApi             RootSchema      `json:"deferApi"`
	Children             []any           `json:"children"`
	Key                  string          `json:"key"`
	DisabledTip          string          `json:"disabledTip"`
	Mode                 string          `json:"mode"`
}

type NavOverflow struct {
	Enable                   bool              `json:"enable"`
	OverflowLabel            SchemaClassName   `json:"overflowLabel"`
	OverflowIndicator        overflowIndicator `json:"overflowIndicator"`
	OverflowClassName        className         `json:"overflowClassName"`
	OverflowPopoverClassName className         `json:"overflowPopoverClassName"`
	OverflowListClassName    className         `json:"overflowListClassName"`
	MaxVisibleCount          float64           `json:"maxVisibleCount"`
	WrapperComponent         string            `json:"wrapperComponent"`
	ItemWidth                float64           `json:"itemWidth"`
	OverflowSuffix           className         `json:"overflowSuffix"`
	Style                    className         `json:"style"`
}

type NavSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Links                []any           `json:"links"`
	IndentSize           float64         `json:"indentSize"`
	Source               className       `json:"source"`
	DeferApi             className       `json:"deferApi"`
	Stacked              bool            `json:"stacked"`
	ItemActions          className       `json:"itemActions"`
	Draggable            bool            `json:"draggable"`
	SaveOrderApi         className       `json:"saveOrderApi"`
	ItemBadge            className       `json:"itemBadge"`
	Badge                className       `json:"badge"`
	DragOnSameLevel      bool            `json:"dragOnSameLevel"`
	Overflow             className       `json:"overflow"`
	Level                float64         `json:"level"`
	DefaultOpenLevel     float64         `json:"defaultOpenLevel"`
	ShowKey              string          `json:"showKey"`
	Collapsed            bool            `json:"collapsed"`
	Mode                 string          `json:"mode"`
	ExpandIcon           SchemaClassName `json:"expandIcon"`
	ExpandPosition       string          `json:"expandPosition"`
	ThemeColor           string          `json:"themeColor"`
	Accordion            bool            `json:"accordion"`
	PopupClassName       string          `json:"popupClassName"`
	Searchable           bool            `json:"searchable"`
	SearchConfig         searchConfig    `json:"searchConfig"`
}

type NestedSelectControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	BorderMode           string           `json:"borderMode"`
	MenuClassName        string           `json:"menuClassName"`
	Cascade              bool             `json:"cascade"`
	WithChildren         bool             `json:"withChildren"`
	OnlyChildren         bool             `json:"onlyChildren"`
	OnlyLeaf             bool             `json:"onlyLeaf"`
	HideNodePathLabel    bool             `json:"hideNodePathLabel"`
	MaxTagCount          float64          `json:"maxTagCount"`
	OverflowTagPopover   any              `json:"overflowTagPopover"`
}

type NumberControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Max                  float64          `json:"max"`
	Min                  float64          `json:"min"`
	Step                 float64          `json:"step"`
	Precision            float64          `json:"precision"`
	ShowSteps            bool             `json:"showSteps"`
	BorderMode           string           `json:"borderMode"`
	Prefix               string           `json:"prefix"`
	Suffix               string           `json:"suffix"`
	UnitOptions          SchemaClassName  `json:"unitOptions"`
	Big                  bool             `json:"big"`
	KilobitSeparator     bool             `json:"kilobitSeparator"`
	Keyboard             bool             `json:"keyboard"`
	DisplayMode          string           `json:"displayMode"`
	ShowAsPercent        bool             `json:"showAsPercent"`
}

type OperationSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Placeholder          string        `json:"placeholder"`
	Buttons              []any         `json:"buttons"`
}

type OperatorType struct {
	AnyOf []ArrayElem3 `json:"anyOf"`
}

type Option struct {
	Label       string          `json:"label"`
	ScopeLabel  string          `json:"scopeLabel"`
	Value       mock            `json:"value"`
	Disabled    bool            `json:"disabled"`
	Children    className       `json:"children"`
	Visible     bool            `json:"visible"`
	Hidden      bool            `json:"hidden"`
	Description string          `json:"description"`
	DeferVal    bool            `json:"defer"`
	DeferApi    draggableConfig `json:"deferApi"`
	Loading     bool            `json:"loading"`
	Loaded      bool            `json:"loaded"`
}

type OtherActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
}

type PageSchema struct {
	LoadingConfig        loadingConfig `json:"loadingConfig"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Title                string        `json:"title"`
	SubTitle             string        `json:"subTitle"`
	Remark               className     `json:"remark"`
	Body                 className     `json:"body"`
	BodyClassName        className     `json:"bodyClassName"`
	Aside                className     `json:"aside"`
	AsideResizor         bool          `json:"asideResizor"`
	AsideSticky          bool          `json:"asideSticky"`
	AsideMinWidth        float64       `json:"asideMinWidth"`
	AsideMaxWidth        float64       `json:"asideMaxWidth"`
	AsideClassName       className     `json:"asideClassName"`
	Css                  any           `json:"css"`
	MobileCSS            any           `json:"mobileCSS"`
	Data                 className     `json:"data"`
	HeaderClassName      className     `json:"headerClassName"`
	InitApi              className     `json:"initApi"`
	InitFetch            bool          `json:"initFetch"`
	InitFetchOn          className     `json:"initFetchOn"`
	Messages             RootSchema    `json:"messages"`
	Name                 RootSchema    `json:"name"`
	Toolbar              className     `json:"toolbar"`
	ToolbarClassName     className     `json:"toolbarClassName"`
	Definitions          staticSchema  `json:"definitions"`
	Interval             float64       `json:"interval"`
	SilentPolling        bool          `json:"silentPolling"`
	StopAutoRefreshWhen  className     `json:"stopAutoRefreshWhen"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
	CssVars              mock          `json:"cssVars"`
	Regions              []string      `json:"regions"`
	PullRefresh          pullRefresh   `json:"pullRefresh"`
}

type PaginationSchema struct {
	ID                       string            `json:"$$id"`
	ClassName                className         `json:"className"`
	Ref                      string            `json:"$ref"`
	Disabled                 bool              `json:"disabled"`
	DisabledOn               className         `json:"disabledOn"`
	Hidden                   bool              `json:"hidden"`
	HiddenOn                 className         `json:"hiddenOn"`
	Visible                  bool              `json:"visible"`
	VisibleOn                className         `json:"visibleOn"`
	Id                       string            `json:"id"`
	OnEvent                  any               `json:"onEvent"`
	Static                   bool              `json:"static"`
	StaticOn                 className         `json:"staticOn"`
	StaticPlaceholder        string            `json:"staticPlaceholder"`
	StaticClassName          className         `json:"staticClassName"`
	StaticLabelClassName     className         `json:"staticLabelClassName"`
	StaticInputClassName     className         `json:"staticInputClassName"`
	StaticSchema             staticSchema      `json:"staticSchema"`
	Style                    any               `json:"style"`
	EditorSetting            editorSetting     `json:"editorSetting"`
	UseMobileUI              bool              `json:"useMobileUI"`
	TestIdBuilder            RootSchema        `json:"testIdBuilder"`
	TypeVal                  string            `json:"type"`
	Testid                   string            `json:"testid"`
	Layout                   layout            `json:"layout"`
	MaxButtons               float64           `json:"maxButtons"`
	Mode                     overflowIndicator `json:"mode"`
	ActivePage               float64           `json:"activePage"`
	Total                    float64           `json:"total"`
	PerPage                  float64           `json:"perPage"`
	ShowPerPage              bool              `json:"showPerPage"`
	PerPageAvailable         []float64         `json:"perPageAvailable"`
	ShowPageInput            bool              `json:"showPageInput"`
	HasNext                  bool              `json:"hasNext"`
	PopOverContainerSelector string            `json:"popOverContainerSelector"`
}

type PaginationWrapperSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	ShowPageInput        bool          `json:"showPageInput"`
	MaxButtons           float64       `json:"maxButtons"`
	InputName            string        `json:"inputName"`
	OutputName           string        `json:"outputName"`
	PerPage              float64       `json:"perPage"`
	Position             string        `json:"position"`
	Body                 className     `json:"body"`
}

type PanelSchema struct {
	ID                      string          `json:"$$id"`
	ClassName               className       `json:"className"`
	Ref                     string          `json:"$ref"`
	Disabled                bool            `json:"disabled"`
	DisabledOn              className       `json:"disabledOn"`
	Hidden                  bool            `json:"hidden"`
	HiddenOn                className       `json:"hiddenOn"`
	Visible                 bool            `json:"visible"`
	VisibleOn               className       `json:"visibleOn"`
	Id                      string          `json:"id"`
	OnEvent                 any             `json:"onEvent"`
	Static                  bool            `json:"static"`
	StaticOn                className       `json:"staticOn"`
	StaticPlaceholder       string          `json:"staticPlaceholder"`
	StaticClassName         className       `json:"staticClassName"`
	StaticLabelClassName    className       `json:"staticLabelClassName"`
	StaticInputClassName    className       `json:"staticInputClassName"`
	StaticSchema            staticSchema    `json:"staticSchema"`
	Style                   any             `json:"style"`
	EditorSetting           editorSetting   `json:"editorSetting"`
	UseMobileUI             bool            `json:"useMobileUI"`
	TestIdBuilder           RootSchema      `json:"testIdBuilder"`
	TypeVal                 string          `json:"type"`
	Testid                  string          `json:"testid"`
	Actions                 []any           `json:"actions"`
	ActionsClassName        className       `json:"actionsClassName"`
	Body                    className       `json:"body"`
	BodyClassName           className       `json:"bodyClassName"`
	Footer                  className       `json:"footer"`
	FooterClassName         className       `json:"footerClassName"`
	FooterWrapClassName     className       `json:"footerWrapClassName"`
	Header                  className       `json:"header"`
	HeaderClassName         className       `json:"headerClassName"`
	Title                   className       `json:"title"`
	AffixFooter             SchemaClassName `json:"affixFooter"`
	SubFormMode             string          `json:"subFormMode"`
	SubFormHorizontal       className       `json:"subFormHorizontal"`
	HeaderControlClassName  string          `json:"headerControlClassName"`
	BodyControlClassName    string          `json:"bodyControlClassName"`
	ActionsControlClassName string          `json:"actionsControlClassName"`
}

type PasswordSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	MosaicText           string        `json:"mosaicText"`
}

type PickerControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	LabelTpl             className        `json:"labelTpl"`
	LabelField           string           `json:"labelField"`
	ValueField           string           `json:"valueField"`
	PickerSchema         mock             `json:"pickerSchema"`
	ModalMode            string           `json:"modalMode"`
	ModalTitle           string           `json:"modalTitle"`
	Embed                bool             `json:"embed"`
	OverflowConfig       overflowConfig   `json:"overflowConfig"`
}

type PlainSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Tpl                  RootSchema    `json:"tpl"`
	Text                 RootSchema    `json:"text"`
	Inline               bool          `json:"inline"`
	Placeholder          string        `json:"placeholder"`
}

type PortletSchema struct {
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Tabs                 []any         `json:"tabs"`
	Source               string        `json:"source"`
	TabsClassName        className     `json:"tabsClassName"`
	TabsMode             string        `json:"tabsMode"`
	ContentClassName     className     `json:"contentClassName"`
	LinksClassName       className     `json:"linksClassName"`
	MountOnEnter         bool          `json:"mountOnEnter"`
	UnmountOnExit        bool          `json:"unmountOnExit"`
	Toolbar              []any         `json:"toolbar"`
	Scrollable           bool          `json:"scrollable"`
	Divider              bool          `json:"divider"`
	Description          className     `json:"description"`
	HideHeader           bool          `json:"hideHeader"`
}

type PortletTabSchema struct {
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	Title                string        `json:"title"`
	Tab                  className     `json:"tab"`
	Toolbar              []any         `json:"toolbar"`
	Body                 className     `json:"body"`
	Icon                 className     `json:"icon"`
	IconPosition         string        `json:"iconPosition"`
	Reload               bool          `json:"reload"`
	MountOnEnter         bool          `json:"mountOnEnter"`
	UnmountOnExit        bool          `json:"unmountOnExit"`
}

type PresetColor struct {
	AnyOf []ArrayElem23 `json:"anyOf"`
}

type ProgressSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Name                 string        `json:"name"`
	Value                float64       `json:"value"`
	Mode                 string        `json:"mode"`
	ProgressClassName    className     `json:"progressClassName"`
	MapVal               className     `json:"map"`
	ShowLabel            bool          `json:"showLabel"`
	Placeholder          string        `json:"placeholder"`
	Stripe               bool          `json:"stripe"`
	Animate              bool          `json:"animate"`
	StrokeWidth          float64       `json:"strokeWidth"`
	GapDegree            float64       `json:"gapDegree"`
	GapPosition          string        `json:"gapPosition"`
	ValueTpl             string        `json:"valueTpl"`
	Threshold            threshold     `json:"threshold"`
	ShowThresholdText    bool          `json:"showThresholdText"`
}

type QRCodeImageSettings struct {
	Src      string  `json:"src"`
	Height   float64 `json:"height"`
	Width    float64 `json:"width"`
	Excavate bool    `json:"excavate"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type QRCodeSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Name                 string        `json:"name"`
	QrcodeClassName      className     `json:"qrcodeClassName"`
	CodeSize             float64       `json:"codeSize"`
	BackgroundColor      string        `json:"backgroundColor"`
	ForegroundColor      string        `json:"foregroundColor"`
	Level                string        `json:"level"`
	Placeholder          string        `json:"placeholder"`
	ImageSettings        className     `json:"imageSettings"`
	Mode                 string        `json:"mode"`
}

type RadioControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	TrueValue            any              `json:"trueValue"`
	FalseValue           any              `json:"falseValue"`
	Option               string           `json:"option"`
	Badge                className        `json:"badge"`
	Partial              bool             `json:"partial"`
	OptionType           string           `json:"optionType"`
}

type RadiosControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	ColumnsCount         float64          `json:"columnsCount"`
}

type RangeControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                className        `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Max                  SchemaClassName  `json:"max"`
	Min                  SchemaClassName  `json:"min"`
	Step                 SchemaClassName  `json:"step"`
	Unit                 string           `json:"unit"`
	ShowSteps            bool             `json:"showSteps"`
	Parts                SchemaClassName  `json:"parts"`
	Marks                any              `json:"marks"`
	TooltipVisible       bool             `json:"tooltipVisible"`
	TooltipPlacement     className        `json:"tooltipPlacement"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ShowInput            bool             `json:"showInput"`
	Clearable            bool             `json:"clearable"`
}

type RatingControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Count                float64          `json:"count"`
	Half                 bool             `json:"half"`
	AllowClear           bool             `json:"allowClear"`
	Readonly             bool             `json:"readonly"`
	Colors               SchemaClassName  `json:"colors"`
	InactiveColor        string           `json:"inactiveColor"`
	Texts                any              `json:"texts"`
	TextPosition         className        `json:"textPosition"`
	Char                 string           `json:"char"`
	CharClassName        string           `json:"charClassName"`
	TextClassName        string           `json:"textClassName"`
}

type ReactCSSProperties struct {
	AlignmentBaseline              RootSchema `json:"alignmentBaseline"`
	BaselineShift                  RootSchema `json:"baselineShift"`
	Clip                           className  `json:"clip"`
	ClipPath                       className  `json:"clipPath"`
	ClipRule                       RootSchema `json:"clipRule"`
	Color                          className  `json:"color"`
	ColorInterpolation             RootSchema `json:"colorInterpolation"`
	ColorRendering                 RootSchema `json:"colorRendering"`
	Cursor                         className  `json:"cursor"`
	Direction                      className  `json:"direction"`
	Display                        className  `json:"display"`
	DominantBaseline               RootSchema `json:"dominantBaseline"`
	Fill                           RootSchema `json:"fill"`
	FillOpacity                    RootSchema `json:"fillOpacity"`
	FillRule                       RootSchema `json:"fillRule"`
	Filter                         className  `json:"filter"`
	FloodColor                     RootSchema `json:"floodColor"`
	FloodOpacity                   RootSchema `json:"floodOpacity"`
	Font                           className  `json:"font"`
	FontFamily                     className  `json:"fontFamily"`
	FontSize                       className  `json:"fontSize"`
	FontSizeAdjust                 className  `json:"fontSizeAdjust"`
	FontStretch                    className  `json:"fontStretch"`
	FontStyle                      className  `json:"fontStyle"`
	FontVariant                    className  `json:"fontVariant"`
	FontWeight                     className  `json:"fontWeight"`
	GlyphOrientationVertical       RootSchema `json:"glyphOrientationVertical"`
	ImageRendering                 className  `json:"imageRendering"`
	LetterSpacing                  className  `json:"letterSpacing"`
	LightingColor                  RootSchema `json:"lightingColor"`
	LineHeight                     className  `json:"lineHeight"`
	Marker                         RootSchema `json:"marker"`
	MarkerEnd                      RootSchema `json:"markerEnd"`
	MarkerMid                      RootSchema `json:"markerMid"`
	MarkerStart                    RootSchema `json:"markerStart"`
	Mask                           className  `json:"mask"`
	Opacity                        className  `json:"opacity"`
	Overflow                       className  `json:"overflow"`
	PaintOrder                     className  `json:"paintOrder"`
	PointerEvents                  className  `json:"pointerEvents"`
	ShapeRendering                 RootSchema `json:"shapeRendering"`
	StopColor                      RootSchema `json:"stopColor"`
	StopOpacity                    RootSchema `json:"stopOpacity"`
	Stroke                         RootSchema `json:"stroke"`
	StrokeDasharray                RootSchema `json:"strokeDasharray"`
	StrokeDashoffset               RootSchema `json:"strokeDashoffset"`
	StrokeLinecap                  RootSchema `json:"strokeLinecap"`
	StrokeLinejoin                 RootSchema `json:"strokeLinejoin"`
	StrokeMiterlimit               RootSchema `json:"strokeMiterlimit"`
	StrokeOpacity                  RootSchema `json:"strokeOpacity"`
	StrokeWidth                    RootSchema `json:"strokeWidth"`
	TextAnchor                     RootSchema `json:"textAnchor"`
	TextDecoration                 className  `json:"textDecoration"`
	TextRendering                  className  `json:"textRendering"`
	UnicodeBidi                    className  `json:"unicodeBidi"`
	VectorEffect                   RootSchema `json:"vectorEffect"`
	Visibility                     className  `json:"visibility"`
	WhiteSpace                     className  `json:"whiteSpace"`
	WordSpacing                    className  `json:"wordSpacing"`
	WritingMode                    className  `json:"writingMode"`
	Azimuth                        className  `json:"azimuth"`
	BoxAlign                       className  `json:"boxAlign"`
	BoxDirection                   className  `json:"boxDirection"`
	BoxFlex                        className  `json:"boxFlex"`
	BoxFlexGroup                   className  `json:"boxFlexGroup"`
	BoxLines                       className  `json:"boxLines"`
	BoxOrdinalGroup                className  `json:"boxOrdinalGroup"`
	BoxOrient                      className  `json:"boxOrient"`
	BoxPack                        className  `json:"boxPack"`
	GridColumnGap                  className  `json:"gridColumnGap"`
	GridGap                        className  `json:"gridGap"`
	GridRowGap                     className  `json:"gridRowGap"`
	ImeMode                        className  `json:"imeMode"`
	OffsetBlock                    className  `json:"offsetBlock"`
	OffsetBlockEnd                 className  `json:"offsetBlockEnd"`
	OffsetBlockStart               className  `json:"offsetBlockStart"`
	OffsetInline                   className  `json:"offsetInline"`
	OffsetInlineEnd                className  `json:"offsetInlineEnd"`
	OffsetInlineStart              className  `json:"offsetInlineStart"`
	ScrollSnapCoordinate           className  `json:"scrollSnapCoordinate"`
	ScrollSnapDestination          className  `json:"scrollSnapDestination"`
	ScrollSnapPointsX              className  `json:"scrollSnapPointsX"`
	ScrollSnapPointsY              className  `json:"scrollSnapPointsY"`
	ScrollSnapTypeX                className  `json:"scrollSnapTypeX"`
	ScrollSnapTypeY                className  `json:"scrollSnapTypeY"`
	KhtmlBoxAlign                  className  `json:"KhtmlBoxAlign"`
	KhtmlBoxDirection              className  `json:"KhtmlBoxDirection"`
	KhtmlBoxFlex                   className  `json:"KhtmlBoxFlex"`
	KhtmlBoxFlexGroup              className  `json:"KhtmlBoxFlexGroup"`
	KhtmlBoxLines                  className  `json:"KhtmlBoxLines"`
	KhtmlBoxOrdinalGroup           className  `json:"KhtmlBoxOrdinalGroup"`
	KhtmlBoxOrient                 className  `json:"KhtmlBoxOrient"`
	KhtmlBoxPack                   className  `json:"KhtmlBoxPack"`
	KhtmlLineBreak                 className  `json:"KhtmlLineBreak"`
	KhtmlOpacity                   className  `json:"KhtmlOpacity"`
	KhtmlUserSelect                className  `json:"KhtmlUserSelect"`
	MozBackfaceVisibility          className  `json:"MozBackfaceVisibility"`
	MozBackgroundClip              className  `json:"MozBackgroundClip"`
	MozBackgroundInlinePolicy      className  `json:"MozBackgroundInlinePolicy"`
	MozBackgroundOrigin            className  `json:"MozBackgroundOrigin"`
	MozBackgroundSize              className  `json:"MozBackgroundSize"`
	MozBorderRadius                className  `json:"MozBorderRadius"`
	MozBorderRadiusBottomleft      className  `json:"MozBorderRadiusBottomleft"`
	MozBorderRadiusBottomright     className  `json:"MozBorderRadiusBottomright"`
	MozBorderRadiusTopleft         className  `json:"MozBorderRadiusTopleft"`
	MozBorderRadiusTopright        className  `json:"MozBorderRadiusTopright"`
	MozBoxAlign                    className  `json:"MozBoxAlign"`
	MozBoxDirection                className  `json:"MozBoxDirection"`
	MozBoxFlex                     className  `json:"MozBoxFlex"`
	MozBoxOrdinalGroup             className  `json:"MozBoxOrdinalGroup"`
	MozBoxOrient                   className  `json:"MozBoxOrient"`
	MozBoxPack                     className  `json:"MozBoxPack"`
	MozBoxShadow                   className  `json:"MozBoxShadow"`
	MozFloatEdge                   className  `json:"MozFloatEdge"`
	MozForceBrokenImageIcon        className  `json:"MozForceBrokenImageIcon"`
	MozOpacity                     className  `json:"MozOpacity"`
	MozOutline                     className  `json:"MozOutline"`
	MozOutlineColor                className  `json:"MozOutlineColor"`
	MozOutlineStyle                className  `json:"MozOutlineStyle"`
	MozOutlineWidth                className  `json:"MozOutlineWidth"`
	MozPerspective                 className  `json:"MozPerspective"`
	MozPerspectiveOrigin           className  `json:"MozPerspectiveOrigin"`
	MozTextAlignLast               className  `json:"MozTextAlignLast"`
	MozTextDecorationColor         className  `json:"MozTextDecorationColor"`
	MozTextDecorationLine          className  `json:"MozTextDecorationLine"`
	MozTextDecorationStyle         className  `json:"MozTextDecorationStyle"`
	MozTransform                   className  `json:"MozTransform"`
	MozTransformOrigin             className  `json:"MozTransformOrigin"`
	MozTransformStyle              className  `json:"MozTransformStyle"`
	MozTransition                  className  `json:"MozTransition"`
	MozTransitionDelay             className  `json:"MozTransitionDelay"`
	MozTransitionDuration          className  `json:"MozTransitionDuration"`
	MozTransitionProperty          className  `json:"MozTransitionProperty"`
	MozTransitionTimingFunction    className  `json:"MozTransitionTimingFunction"`
	MozUserInput                   className  `json:"MozUserInput"`
	MsImeMode                      className  `json:"msImeMode"`
	OAnimation                     className  `json:"OAnimation"`
	OAnimationDelay                className  `json:"OAnimationDelay"`
	OAnimationDirection            className  `json:"OAnimationDirection"`
	OAnimationDuration             className  `json:"OAnimationDuration"`
	OAnimationFillMode             className  `json:"OAnimationFillMode"`
	OAnimationIterationCount       className  `json:"OAnimationIterationCount"`
	OAnimationName                 className  `json:"OAnimationName"`
	OAnimationPlayState            className  `json:"OAnimationPlayState"`
	OAnimationTimingFunction       className  `json:"OAnimationTimingFunction"`
	OBackgroundSize                className  `json:"OBackgroundSize"`
	OBorderImage                   className  `json:"OBorderImage"`
	OObjectFit                     className  `json:"OObjectFit"`
	OObjectPosition                className  `json:"OObjectPosition"`
	OTabSize                       className  `json:"OTabSize"`
	OTextOverflow                  className  `json:"OTextOverflow"`
	OTransform                     className  `json:"OTransform"`
	OTransformOrigin               className  `json:"OTransformOrigin"`
	OTransition                    className  `json:"OTransition"`
	OTransitionDelay               className  `json:"OTransitionDelay"`
	OTransitionDuration            className  `json:"OTransitionDuration"`
	OTransitionProperty            className  `json:"OTransitionProperty"`
	OTransitionTimingFunction      className  `json:"OTransitionTimingFunction"`
	WebkitBoxAlign                 className  `json:"WebkitBoxAlign"`
	WebkitBoxDirection             className  `json:"WebkitBoxDirection"`
	WebkitBoxFlex                  className  `json:"WebkitBoxFlex"`
	WebkitBoxFlexGroup             className  `json:"WebkitBoxFlexGroup"`
	WebkitBoxLines                 className  `json:"WebkitBoxLines"`
	WebkitBoxOrdinalGroup          className  `json:"WebkitBoxOrdinalGroup"`
	WebkitBoxOrient                className  `json:"WebkitBoxOrient"`
	WebkitBoxPack                  className  `json:"WebkitBoxPack"`
	MozAnimation                   className  `json:"MozAnimation"`
	MozBorderImage                 className  `json:"MozBorderImage"`
	MozColumnRule                  className  `json:"MozColumnRule"`
	MozColumns                     className  `json:"MozColumns"`
	MozOutlineRadius               className  `json:"MozOutlineRadius"`
	MsContentZoomLimit             className  `json:"msContentZoomLimit"`
	MsContentZoomSnap              className  `json:"msContentZoomSnap"`
	MsFlex                         className  `json:"msFlex"`
	MsScrollLimit                  className  `json:"msScrollLimit"`
	MsScrollSnapX                  className  `json:"msScrollSnapX"`
	MsScrollSnapY                  className  `json:"msScrollSnapY"`
	MsTransition                   className  `json:"msTransition"`
	WebkitAnimation                className  `json:"WebkitAnimation"`
	WebkitBorderBefore             className  `json:"WebkitBorderBefore"`
	WebkitBorderImage              className  `json:"WebkitBorderImage"`
	WebkitBorderRadius             className  `json:"WebkitBorderRadius"`
	WebkitColumnRule               className  `json:"WebkitColumnRule"`
	WebkitColumns                  className  `json:"WebkitColumns"`
	WebkitFlex                     className  `json:"WebkitFlex"`
	WebkitFlexFlow                 className  `json:"WebkitFlexFlow"`
	WebkitMask                     className  `json:"WebkitMask"`
	WebkitMaskBoxImage             className  `json:"WebkitMaskBoxImage"`
	WebkitTextEmphasis             className  `json:"WebkitTextEmphasis"`
	WebkitTextStroke               className  `json:"WebkitTextStroke"`
	WebkitTransition               className  `json:"WebkitTransition"`
	MozAnimationDelay              className  `json:"MozAnimationDelay"`
	MozAnimationDirection          className  `json:"MozAnimationDirection"`
	MozAnimationDuration           className  `json:"MozAnimationDuration"`
	MozAnimationFillMode           className  `json:"MozAnimationFillMode"`
	MozAnimationIterationCount     className  `json:"MozAnimationIterationCount"`
	MozAnimationName               className  `json:"MozAnimationName"`
	MozAnimationPlayState          className  `json:"MozAnimationPlayState"`
	MozAnimationTimingFunction     className  `json:"MozAnimationTimingFunction"`
	MozAppearance                  className  `json:"MozAppearance"`
	MozBinding                     className  `json:"MozBinding"`
	MozBorderBottomColors          className  `json:"MozBorderBottomColors"`
	MozBorderEndColor              className  `json:"MozBorderEndColor"`
	MozBorderEndStyle              className  `json:"MozBorderEndStyle"`
	MozBorderEndWidth              className  `json:"MozBorderEndWidth"`
	MozBorderLeftColors            className  `json:"MozBorderLeftColors"`
	MozBorderRightColors           className  `json:"MozBorderRightColors"`
	MozBorderStartColor            className  `json:"MozBorderStartColor"`
	MozBorderStartStyle            className  `json:"MozBorderStartStyle"`
	MozBorderTopColors             className  `json:"MozBorderTopColors"`
	MozBoxSizing                   className  `json:"MozBoxSizing"`
	MozColumnCount                 className  `json:"MozColumnCount"`
	MozColumnFill                  className  `json:"MozColumnFill"`
	MozColumnRuleColor             className  `json:"MozColumnRuleColor"`
	MozColumnRuleStyle             className  `json:"MozColumnRuleStyle"`
	MozColumnRuleWidth             className  `json:"MozColumnRuleWidth"`
	MozColumnWidth                 className  `json:"MozColumnWidth"`
	MozContextProperties           className  `json:"MozContextProperties"`
	MozFontFeatureSettings         className  `json:"MozFontFeatureSettings"`
	MozFontLanguageOverride        className  `json:"MozFontLanguageOverride"`
	MozHyphens                     className  `json:"MozHyphens"`
	MozImageRegion                 className  `json:"MozImageRegion"`
	MozMarginEnd                   className  `json:"MozMarginEnd"`
	MozMarginStart                 className  `json:"MozMarginStart"`
	MozOrient                      className  `json:"MozOrient"`
	MozOsxFontSmoothing            className  `json:"MozOsxFontSmoothing"`
	MozOutlineRadiusBottomleft     className  `json:"MozOutlineRadiusBottomleft"`
	MozOutlineRadiusBottomright    className  `json:"MozOutlineRadiusBottomright"`
	MozOutlineRadiusTopleft        className  `json:"MozOutlineRadiusTopleft"`
	MozOutlineRadiusTopright       className  `json:"MozOutlineRadiusTopright"`
	MozPaddingEnd                  className  `json:"MozPaddingEnd"`
	MozPaddingStart                className  `json:"MozPaddingStart"`
	MozStackSizing                 className  `json:"MozStackSizing"`
	MozTabSize                     className  `json:"MozTabSize"`
	MozTextBlink                   className  `json:"MozTextBlink"`
	MozTextSizeAdjust              className  `json:"MozTextSizeAdjust"`
	MozUserFocus                   className  `json:"MozUserFocus"`
	MozUserModify                  className  `json:"MozUserModify"`
	MozUserSelect                  className  `json:"MozUserSelect"`
	MozWindowDragging              className  `json:"MozWindowDragging"`
	MozWindowShadow                className  `json:"MozWindowShadow"`
	MsAccelerator                  className  `json:"msAccelerator"`
	MsBlockProgression             className  `json:"msBlockProgression"`
	MsContentZoomChaining          className  `json:"msContentZoomChaining"`
	MsContentZoomLimitMax          className  `json:"msContentZoomLimitMax"`
	MsContentZoomLimitMin          className  `json:"msContentZoomLimitMin"`
	MsContentZoomSnapPoints        className  `json:"msContentZoomSnapPoints"`
	MsContentZoomSnapType          className  `json:"msContentZoomSnapType"`
	MsContentZooming               className  `json:"msContentZooming"`
	MsFilter                       className  `json:"msFilter"`
	MsFlexDirection                className  `json:"msFlexDirection"`
	MsFlexPositive                 className  `json:"msFlexPositive"`
	MsFlowFrom                     className  `json:"msFlowFrom"`
	MsFlowInto                     className  `json:"msFlowInto"`
	MsGridColumns                  className  `json:"msGridColumns"`
	MsGridRows                     className  `json:"msGridRows"`
	MsHighContrastAdjust           className  `json:"msHighContrastAdjust"`
	MsHyphenateLimitChars          className  `json:"msHyphenateLimitChars"`
	MsHyphenateLimitLines          className  `json:"msHyphenateLimitLines"`
	MsHyphenateLimitZone           className  `json:"msHyphenateLimitZone"`
	MsHyphens                      className  `json:"msHyphens"`
	MsImeAlign                     className  `json:"msImeAlign"`
	MsLineBreak                    className  `json:"msLineBreak"`
	MsOrder                        className  `json:"msOrder"`
	MsOverflowStyle                className  `json:"msOverflowStyle"`
	MsOverflowX                    className  `json:"msOverflowX"`
	MsOverflowY                    className  `json:"msOverflowY"`
	MsScrollChaining               className  `json:"msScrollChaining"`
	MsScrollLimitXMax              className  `json:"msScrollLimitXMax"`
	MsScrollLimitXMin              className  `json:"msScrollLimitXMin"`
	MsScrollLimitYMax              className  `json:"msScrollLimitYMax"`
	MsScrollLimitYMin              className  `json:"msScrollLimitYMin"`
	MsScrollRails                  className  `json:"msScrollRails"`
	MsScrollSnapPointsX            className  `json:"msScrollSnapPointsX"`
	MsScrollSnapPointsY            className  `json:"msScrollSnapPointsY"`
	MsScrollSnapType               className  `json:"msScrollSnapType"`
	MsScrollTranslation            className  `json:"msScrollTranslation"`
	MsScrollbar3dlightColor        className  `json:"msScrollbar3dlightColor"`
	MsScrollbarArrowColor          className  `json:"msScrollbarArrowColor"`
	MsScrollbarBaseColor           className  `json:"msScrollbarBaseColor"`
	MsScrollbarDarkshadowColor     className  `json:"msScrollbarDarkshadowColor"`
	MsScrollbarFaceColor           className  `json:"msScrollbarFaceColor"`
	MsScrollbarHighlightColor      className  `json:"msScrollbarHighlightColor"`
	MsScrollbarShadowColor         className  `json:"msScrollbarShadowColor"`
	MsScrollbarTrackColor          className  `json:"msScrollbarTrackColor"`
	MsTextAutospace                className  `json:"msTextAutospace"`
	MsTextCombineHorizontal        className  `json:"msTextCombineHorizontal"`
	MsTextOverflow                 className  `json:"msTextOverflow"`
	MsTouchAction                  className  `json:"msTouchAction"`
	MsTouchSelect                  className  `json:"msTouchSelect"`
	MsTransform                    className  `json:"msTransform"`
	MsTransformOrigin              className  `json:"msTransformOrigin"`
	MsTransitionDelay              className  `json:"msTransitionDelay"`
	MsTransitionDuration           className  `json:"msTransitionDuration"`
	MsTransitionProperty           className  `json:"msTransitionProperty"`
	MsTransitionTimingFunction     className  `json:"msTransitionTimingFunction"`
	MsUserSelect                   className  `json:"msUserSelect"`
	MsWordBreak                    className  `json:"msWordBreak"`
	MsWrapFlow                     className  `json:"msWrapFlow"`
	MsWrapMargin                   className  `json:"msWrapMargin"`
	MsWrapThrough                  className  `json:"msWrapThrough"`
	MsWritingMode                  className  `json:"msWritingMode"`
	WebkitAlignContent             className  `json:"WebkitAlignContent"`
	WebkitAlignItems               className  `json:"WebkitAlignItems"`
	WebkitAlignSelf                className  `json:"WebkitAlignSelf"`
	WebkitAnimationDelay           className  `json:"WebkitAnimationDelay"`
	WebkitAnimationDirection       className  `json:"WebkitAnimationDirection"`
	WebkitAnimationDuration        className  `json:"WebkitAnimationDuration"`
	WebkitAnimationFillMode        className  `json:"WebkitAnimationFillMode"`
	WebkitAnimationIterationCount  className  `json:"WebkitAnimationIterationCount"`
	WebkitAnimationName            className  `json:"WebkitAnimationName"`
	WebkitAnimationPlayState       className  `json:"WebkitAnimationPlayState"`
	WebkitAnimationTimingFunction  className  `json:"WebkitAnimationTimingFunction"`
	WebkitAppearance               className  `json:"WebkitAppearance"`
	WebkitBackdropFilter           className  `json:"WebkitBackdropFilter"`
	WebkitBackfaceVisibility       className  `json:"WebkitBackfaceVisibility"`
	WebkitBackgroundClip           className  `json:"WebkitBackgroundClip"`
	WebkitBackgroundOrigin         className  `json:"WebkitBackgroundOrigin"`
	WebkitBackgroundSize           className  `json:"WebkitBackgroundSize"`
	WebkitBorderBeforeColor        className  `json:"WebkitBorderBeforeColor"`
	WebkitBorderBeforeStyle        className  `json:"WebkitBorderBeforeStyle"`
	WebkitBorderBeforeWidth        className  `json:"WebkitBorderBeforeWidth"`
	WebkitBorderBottomLeftRadius   className  `json:"WebkitBorderBottomLeftRadius"`
	WebkitBorderBottomRightRadius  className  `json:"WebkitBorderBottomRightRadius"`
	WebkitBorderImageSlice         className  `json:"WebkitBorderImageSlice"`
	WebkitBorderTopLeftRadius      className  `json:"WebkitBorderTopLeftRadius"`
	WebkitBorderTopRightRadius     className  `json:"WebkitBorderTopRightRadius"`
	WebkitBoxDecorationBreak       className  `json:"WebkitBoxDecorationBreak"`
	WebkitBoxReflect               className  `json:"WebkitBoxReflect"`
	WebkitBoxShadow                className  `json:"WebkitBoxShadow"`
	WebkitBoxSizing                className  `json:"WebkitBoxSizing"`
	WebkitClipPath                 className  `json:"WebkitClipPath"`
	WebkitColumnCount              className  `json:"WebkitColumnCount"`
	WebkitColumnFill               className  `json:"WebkitColumnFill"`
	WebkitColumnRuleColor          className  `json:"WebkitColumnRuleColor"`
	WebkitColumnRuleStyle          className  `json:"WebkitColumnRuleStyle"`
	WebkitColumnRuleWidth          className  `json:"WebkitColumnRuleWidth"`
	WebkitColumnSpan               className  `json:"WebkitColumnSpan"`
	WebkitColumnWidth              className  `json:"WebkitColumnWidth"`
	WebkitFilter                   className  `json:"WebkitFilter"`
	WebkitFlexBasis                className  `json:"WebkitFlexBasis"`
	WebkitFlexDirection            className  `json:"WebkitFlexDirection"`
	WebkitFlexGrow                 className  `json:"WebkitFlexGrow"`
	WebkitFlexShrink               className  `json:"WebkitFlexShrink"`
	WebkitFlexWrap                 className  `json:"WebkitFlexWrap"`
	WebkitFontFeatureSettings      className  `json:"WebkitFontFeatureSettings"`
	WebkitFontKerning              className  `json:"WebkitFontKerning"`
	WebkitFontSmoothing            className  `json:"WebkitFontSmoothing"`
	WebkitFontVariantLigatures     className  `json:"WebkitFontVariantLigatures"`
	WebkitHyphenateCharacter       className  `json:"WebkitHyphenateCharacter"`
	WebkitHyphens                  className  `json:"WebkitHyphens"`
	WebkitInitialLetter            className  `json:"WebkitInitialLetter"`
	WebkitJustifyContent           className  `json:"WebkitJustifyContent"`
	WebkitLineBreak                className  `json:"WebkitLineBreak"`
	WebkitLineClamp                className  `json:"WebkitLineClamp"`
	WebkitMarginEnd                className  `json:"WebkitMarginEnd"`
	WebkitMarginStart              className  `json:"WebkitMarginStart"`
	WebkitMaskAttachment           className  `json:"WebkitMaskAttachment"`
	WebkitMaskBoxImageOutset       className  `json:"WebkitMaskBoxImageOutset"`
	WebkitMaskBoxImageRepeat       className  `json:"WebkitMaskBoxImageRepeat"`
	WebkitMaskBoxImageSlice        className  `json:"WebkitMaskBoxImageSlice"`
	WebkitMaskBoxImageSource       className  `json:"WebkitMaskBoxImageSource"`
	WebkitMaskBoxImageWidth        className  `json:"WebkitMaskBoxImageWidth"`
	WebkitMaskClip                 className  `json:"WebkitMaskClip"`
	WebkitMaskComposite            className  `json:"WebkitMaskComposite"`
	WebkitMaskImage                className  `json:"WebkitMaskImage"`
	WebkitMaskOrigin               className  `json:"WebkitMaskOrigin"`
	WebkitMaskPosition             className  `json:"WebkitMaskPosition"`
	WebkitMaskPositionX            className  `json:"WebkitMaskPositionX"`
	WebkitMaskPositionY            className  `json:"WebkitMaskPositionY"`
	WebkitMaskRepeat               className  `json:"WebkitMaskRepeat"`
	WebkitMaskRepeatX              className  `json:"WebkitMaskRepeatX"`
	WebkitMaskRepeatY              className  `json:"WebkitMaskRepeatY"`
	WebkitMaskSize                 className  `json:"WebkitMaskSize"`
	WebkitMaxInlineSize            className  `json:"WebkitMaxInlineSize"`
	WebkitOrder                    className  `json:"WebkitOrder"`
	WebkitOverflowScrolling        className  `json:"WebkitOverflowScrolling"`
	WebkitPaddingEnd               className  `json:"WebkitPaddingEnd"`
	WebkitPaddingStart             className  `json:"WebkitPaddingStart"`
	WebkitPerspective              className  `json:"WebkitPerspective"`
	WebkitPerspectiveOrigin        className  `json:"WebkitPerspectiveOrigin"`
	WebkitPrintColorAdjust         className  `json:"WebkitPrintColorAdjust"`
	WebkitRubyPosition             className  `json:"WebkitRubyPosition"`
	WebkitScrollSnapType           className  `json:"WebkitScrollSnapType"`
	WebkitShapeMargin              className  `json:"WebkitShapeMargin"`
	WebkitTapHighlightColor        className  `json:"WebkitTapHighlightColor"`
	WebkitTextCombine              className  `json:"WebkitTextCombine"`
	WebkitTextDecorationColor      className  `json:"WebkitTextDecorationColor"`
	WebkitTextDecorationLine       className  `json:"WebkitTextDecorationLine"`
	WebkitTextDecorationSkip       className  `json:"WebkitTextDecorationSkip"`
	WebkitTextDecorationStyle      className  `json:"WebkitTextDecorationStyle"`
	WebkitTextEmphasisColor        className  `json:"WebkitTextEmphasisColor"`
	WebkitTextEmphasisPosition     className  `json:"WebkitTextEmphasisPosition"`
	WebkitTextEmphasisStyle        className  `json:"WebkitTextEmphasisStyle"`
	WebkitTextFillColor            className  `json:"WebkitTextFillColor"`
	WebkitTextOrientation          className  `json:"WebkitTextOrientation"`
	WebkitTextSizeAdjust           className  `json:"WebkitTextSizeAdjust"`
	WebkitTextStrokeColor          className  `json:"WebkitTextStrokeColor"`
	WebkitTextStrokeWidth          className  `json:"WebkitTextStrokeWidth"`
	WebkitTextUnderlinePosition    className  `json:"WebkitTextUnderlinePosition"`
	WebkitTouchCallout             className  `json:"WebkitTouchCallout"`
	WebkitTransform                className  `json:"WebkitTransform"`
	WebkitTransformOrigin          className  `json:"WebkitTransformOrigin"`
	WebkitTransformStyle           className  `json:"WebkitTransformStyle"`
	WebkitTransitionDelay          className  `json:"WebkitTransitionDelay"`
	WebkitTransitionDuration       className  `json:"WebkitTransitionDuration"`
	WebkitTransitionProperty       className  `json:"WebkitTransitionProperty"`
	WebkitTransitionTimingFunction className  `json:"WebkitTransitionTimingFunction"`
	WebkitUserModify               className  `json:"WebkitUserModify"`
	WebkitUserSelect               className  `json:"WebkitUserSelect"`
	WebkitWritingMode              className  `json:"WebkitWritingMode"`
	All                            className  `json:"all"`
	Animation                      className  `json:"animation"`
	AnimationRange                 className  `json:"animationRange"`
	Background                     className  `json:"background"`
	BackgroundPosition             className  `json:"backgroundPosition"`
	Border                         className  `json:"border"`
	BorderBlock                    className  `json:"borderBlock"`
	BorderBlockEnd                 className  `json:"borderBlockEnd"`
	BorderBlockStart               className  `json:"borderBlockStart"`
	BorderBottom                   className  `json:"borderBottom"`
	BorderColor                    className  `json:"borderColor"`
	BorderImage                    className  `json:"borderImage"`
	BorderInline                   className  `json:"borderInline"`
	BorderInlineEnd                className  `json:"borderInlineEnd"`
	BorderInlineStart              className  `json:"borderInlineStart"`
	BorderLeft                     className  `json:"borderLeft"`
	BorderRadius                   className  `json:"borderRadius"`
	BorderRight                    className  `json:"borderRight"`
	BorderStyle                    className  `json:"borderStyle"`
	BorderTop                      className  `json:"borderTop"`
	BorderWidth                    className  `json:"borderWidth"`
	Caret                          className  `json:"caret"`
	ColumnRule                     className  `json:"columnRule"`
	Columns                        className  `json:"columns"`
	ContainIntrinsicSize           className  `json:"containIntrinsicSize"`
	Container                      className  `json:"container"`
	Flex                           className  `json:"flex"`
	FlexFlow                       className  `json:"flexFlow"`
	Gap                            className  `json:"gap"`
	Grid                           className  `json:"grid"`
	GridArea                       className  `json:"gridArea"`
	GridColumn                     className  `json:"gridColumn"`
	GridRow                        className  `json:"gridRow"`
	GridTemplate                   className  `json:"gridTemplate"`
	Inset                          className  `json:"inset"`
	InsetBlock                     className  `json:"insetBlock"`
	InsetInline                    className  `json:"insetInline"`
	LineClamp                      className  `json:"lineClamp"`
	ListStyle                      className  `json:"listStyle"`
	Margin                         className  `json:"margin"`
	MarginBlock                    className  `json:"marginBlock"`
	MarginInline                   className  `json:"marginInline"`
	MaskBorder                     className  `json:"maskBorder"`
	Motion                         className  `json:"motion"`
	Offset                         className  `json:"offset"`
	Outline                        className  `json:"outline"`
	OverscrollBehavior             className  `json:"overscrollBehavior"`
	Padding                        className  `json:"padding"`
	PaddingBlock                   className  `json:"paddingBlock"`
	PaddingInline                  className  `json:"paddingInline"`
	PlaceContent                   className  `json:"placeContent"`
	PlaceItems                     className  `json:"placeItems"`
	PlaceSelf                      className  `json:"placeSelf"`
	ScrollMargin                   className  `json:"scrollMargin"`
	ScrollMarginBlock              className  `json:"scrollMarginBlock"`
	ScrollMarginInline             className  `json:"scrollMarginInline"`
	ScrollPadding                  className  `json:"scrollPadding"`
	ScrollPaddingBlock             className  `json:"scrollPaddingBlock"`
	ScrollPaddingInline            className  `json:"scrollPaddingInline"`
	ScrollSnapMargin               className  `json:"scrollSnapMargin"`
	ScrollTimeline                 className  `json:"scrollTimeline"`
	TextEmphasis                   className  `json:"textEmphasis"`
	Transition                     className  `json:"transition"`
	ViewTimeline                   className  `json:"viewTimeline"`
	AccentColor                    className  `json:"accentColor"`
	AlignContent                   className  `json:"alignContent"`
	AlignItems                     className  `json:"alignItems"`
	AlignSelf                      className  `json:"alignSelf"`
	AlignTracks                    className  `json:"alignTracks"`
	AnimationComposition           className  `json:"animationComposition"`
	AnimationDelay                 className  `json:"animationDelay"`
	AnimationDirection             className  `json:"animationDirection"`
	AnimationDuration              className  `json:"animationDuration"`
	AnimationFillMode              className  `json:"animationFillMode"`
	AnimationIterationCount        className  `json:"animationIterationCount"`
	AnimationName                  className  `json:"animationName"`
	AnimationPlayState             className  `json:"animationPlayState"`
	AnimationRangeEnd              className  `json:"animationRangeEnd"`
	AnimationRangeStart            className  `json:"animationRangeStart"`
	AnimationTimeline              className  `json:"animationTimeline"`
	AnimationTimingFunction        className  `json:"animationTimingFunction"`
	Appearance                     className  `json:"appearance"`
	AspectRatio                    className  `json:"aspectRatio"`
	BackdropFilter                 className  `json:"backdropFilter"`
	BackfaceVisibility             className  `json:"backfaceVisibility"`
	BackgroundAttachment           className  `json:"backgroundAttachment"`
	BackgroundBlendMode            className  `json:"backgroundBlendMode"`
	BackgroundClip                 className  `json:"backgroundClip"`
	BackgroundColor                className  `json:"backgroundColor"`
	BackgroundImage                className  `json:"backgroundImage"`
	BackgroundOrigin               className  `json:"backgroundOrigin"`
	BackgroundPositionX            className  `json:"backgroundPositionX"`
	BackgroundPositionY            className  `json:"backgroundPositionY"`
	BackgroundRepeat               className  `json:"backgroundRepeat"`
	BackgroundSize                 className  `json:"backgroundSize"`
	BlockOverflow                  className  `json:"blockOverflow"`
	BlockSize                      className  `json:"blockSize"`
	BorderBlockColor               className  `json:"borderBlockColor"`
	BorderBlockEndColor            className  `json:"borderBlockEndColor"`
	BorderBlockEndStyle            className  `json:"borderBlockEndStyle"`
	BorderBlockEndWidth            className  `json:"borderBlockEndWidth"`
	BorderBlockStartColor          className  `json:"borderBlockStartColor"`
	BorderBlockStartStyle          className  `json:"borderBlockStartStyle"`
	BorderBlockStartWidth          className  `json:"borderBlockStartWidth"`
	BorderBlockStyle               className  `json:"borderBlockStyle"`
	BorderBlockWidth               className  `json:"borderBlockWidth"`
	BorderBottomColor              className  `json:"borderBottomColor"`
	BorderBottomLeftRadius         className  `json:"borderBottomLeftRadius"`
	BorderBottomRightRadius        className  `json:"borderBottomRightRadius"`
	BorderBottomStyle              className  `json:"borderBottomStyle"`
	BorderBottomWidth              className  `json:"borderBottomWidth"`
	BorderCollapse                 className  `json:"borderCollapse"`
	BorderEndEndRadius             className  `json:"borderEndEndRadius"`
	BorderEndStartRadius           className  `json:"borderEndStartRadius"`
	BorderImageOutset              className  `json:"borderImageOutset"`
	BorderImageRepeat              className  `json:"borderImageRepeat"`
	BorderImageSlice               className  `json:"borderImageSlice"`
	BorderImageSource              className  `json:"borderImageSource"`
	BorderImageWidth               className  `json:"borderImageWidth"`
	BorderInlineColor              className  `json:"borderInlineColor"`
	BorderInlineEndColor           className  `json:"borderInlineEndColor"`
	BorderInlineEndStyle           className  `json:"borderInlineEndStyle"`
	BorderInlineEndWidth           className  `json:"borderInlineEndWidth"`
	BorderInlineStartColor         className  `json:"borderInlineStartColor"`
	BorderInlineStartStyle         className  `json:"borderInlineStartStyle"`
	BorderInlineStartWidth         className  `json:"borderInlineStartWidth"`
	BorderInlineStyle              className  `json:"borderInlineStyle"`
	BorderInlineWidth              className  `json:"borderInlineWidth"`
	BorderLeftColor                className  `json:"borderLeftColor"`
	BorderLeftStyle                className  `json:"borderLeftStyle"`
	BorderLeftWidth                className  `json:"borderLeftWidth"`
	BorderRightColor               className  `json:"borderRightColor"`
	BorderRightStyle               className  `json:"borderRightStyle"`
	BorderRightWidth               className  `json:"borderRightWidth"`
	BorderSpacing                  className  `json:"borderSpacing"`
	BorderStartEndRadius           className  `json:"borderStartEndRadius"`
	BorderStartStartRadius         className  `json:"borderStartStartRadius"`
	BorderTopColor                 className  `json:"borderTopColor"`
	BorderTopLeftRadius            className  `json:"borderTopLeftRadius"`
	BorderTopRightRadius           className  `json:"borderTopRightRadius"`
	BorderTopStyle                 className  `json:"borderTopStyle"`
	BorderTopWidth                 className  `json:"borderTopWidth"`
	Bottom                         className  `json:"bottom"`
	BoxDecorationBreak             className  `json:"boxDecorationBreak"`
	BoxShadow                      className  `json:"boxShadow"`
	BoxSizing                      className  `json:"boxSizing"`
	BreakAfter                     className  `json:"breakAfter"`
	BreakBefore                    className  `json:"breakBefore"`
	BreakInside                    className  `json:"breakInside"`
	CaptionSide                    className  `json:"captionSide"`
	CaretColor                     className  `json:"caretColor"`
	CaretShape                     className  `json:"caretShape"`
	Clear                          className  `json:"clear"`
	ColorAdjust                    className  `json:"colorAdjust"`
	ColorScheme                    className  `json:"colorScheme"`
	ColumnCount                    className  `json:"columnCount"`
	ColumnFill                     className  `json:"columnFill"`
	ColumnGap                      className  `json:"columnGap"`
	ColumnRuleColor                className  `json:"columnRuleColor"`
	ColumnRuleStyle                className  `json:"columnRuleStyle"`
	ColumnRuleWidth                className  `json:"columnRuleWidth"`
	ColumnSpan                     className  `json:"columnSpan"`
	ColumnWidth                    className  `json:"columnWidth"`
	Contain                        className  `json:"contain"`
	ContainIntrinsicBlockSize      className  `json:"containIntrinsicBlockSize"`
	ContainIntrinsicHeight         className  `json:"containIntrinsicHeight"`
	ContainIntrinsicInlineSize     className  `json:"containIntrinsicInlineSize"`
	ContainIntrinsicWidth          className  `json:"containIntrinsicWidth"`
	ContainerName                  className  `json:"containerName"`
	ContainerType                  className  `json:"containerType"`
	Content                        className  `json:"content"`
	ContentVisibility              className  `json:"contentVisibility"`
	CounterIncrement               className  `json:"counterIncrement"`
	CounterReset                   className  `json:"counterReset"`
	CounterSet                     className  `json:"counterSet"`
	EmptyCells                     className  `json:"emptyCells"`
	FlexBasis                      className  `json:"flexBasis"`
	FlexDirection                  className  `json:"flexDirection"`
	FlexGrow                       className  `json:"flexGrow"`
	FlexShrink                     className  `json:"flexShrink"`
	FlexWrap                       className  `json:"flexWrap"`
	Float                          className  `json:"float"`
	FontFeatureSettings            className  `json:"fontFeatureSettings"`
	FontKerning                    className  `json:"fontKerning"`
	FontLanguageOverride           className  `json:"fontLanguageOverride"`
	FontOpticalSizing              className  `json:"fontOpticalSizing"`
	FontPalette                    className  `json:"fontPalette"`
	FontSmooth                     className  `json:"fontSmooth"`
	FontSynthesis                  className  `json:"fontSynthesis"`
	FontSynthesisPosition          className  `json:"fontSynthesisPosition"`
	FontSynthesisSmallCaps         className  `json:"fontSynthesisSmallCaps"`
	FontSynthesisStyle             className  `json:"fontSynthesisStyle"`
	FontSynthesisWeight            className  `json:"fontSynthesisWeight"`
	FontVariantAlternates          className  `json:"fontVariantAlternates"`
	FontVariantCaps                className  `json:"fontVariantCaps"`
	FontVariantEastAsian           className  `json:"fontVariantEastAsian"`
	FontVariantEmoji               className  `json:"fontVariantEmoji"`
	FontVariantLigatures           className  `json:"fontVariantLigatures"`
	FontVariantNumeric             className  `json:"fontVariantNumeric"`
	FontVariantPosition            className  `json:"fontVariantPosition"`
	FontVariationSettings          className  `json:"fontVariationSettings"`
	ForcedColorAdjust              className  `json:"forcedColorAdjust"`
	GridAutoColumns                className  `json:"gridAutoColumns"`
	GridAutoFlow                   className  `json:"gridAutoFlow"`
	GridAutoRows                   className  `json:"gridAutoRows"`
	GridColumnEnd                  className  `json:"gridColumnEnd"`
	GridColumnStart                className  `json:"gridColumnStart"`
	GridRowEnd                     className  `json:"gridRowEnd"`
	GridRowStart                   className  `json:"gridRowStart"`
	GridTemplateAreas              className  `json:"gridTemplateAreas"`
	GridTemplateColumns            className  `json:"gridTemplateColumns"`
	GridTemplateRows               className  `json:"gridTemplateRows"`
	HangingPunctuation             className  `json:"hangingPunctuation"`
	Height                         className  `json:"height"`
	HyphenateCharacter             className  `json:"hyphenateCharacter"`
	HyphenateLimitChars            className  `json:"hyphenateLimitChars"`
	Hyphens                        className  `json:"hyphens"`
	ImageOrientation               className  `json:"imageOrientation"`
	ImageResolution                className  `json:"imageResolution"`
	InitialLetter                  className  `json:"initialLetter"`
	InlineSize                     className  `json:"inlineSize"`
	InputSecurity                  className  `json:"inputSecurity"`
	InsetBlockEnd                  className  `json:"insetBlockEnd"`
	InsetBlockStart                className  `json:"insetBlockStart"`
	InsetInlineEnd                 className  `json:"insetInlineEnd"`
	InsetInlineStart               className  `json:"insetInlineStart"`
	Isolation                      className  `json:"isolation"`
	JustifyContent                 className  `json:"justifyContent"`
	JustifyItems                   className  `json:"justifyItems"`
	JustifySelf                    className  `json:"justifySelf"`
	JustifyTracks                  className  `json:"justifyTracks"`
	Left                           className  `json:"left"`
	LineBreak                      className  `json:"lineBreak"`
	LineHeightStep                 className  `json:"lineHeightStep"`
	ListStyleImage                 className  `json:"listStyleImage"`
	ListStylePosition              className  `json:"listStylePosition"`
	ListStyleType                  className  `json:"listStyleType"`
	MarginBlockEnd                 className  `json:"marginBlockEnd"`
	MarginBlockStart               className  `json:"marginBlockStart"`
	MarginBottom                   className  `json:"marginBottom"`
	MarginInlineEnd                className  `json:"marginInlineEnd"`
	MarginInlineStart              className  `json:"marginInlineStart"`
	MarginLeft                     className  `json:"marginLeft"`
	MarginRight                    className  `json:"marginRight"`
	MarginTop                      className  `json:"marginTop"`
	MarginTrim                     className  `json:"marginTrim"`
	MaskBorderMode                 className  `json:"maskBorderMode"`
	MaskBorderOutset               className  `json:"maskBorderOutset"`
	MaskBorderRepeat               className  `json:"maskBorderRepeat"`
	MaskBorderSlice                className  `json:"maskBorderSlice"`
	MaskBorderSource               className  `json:"maskBorderSource"`
	MaskBorderWidth                className  `json:"maskBorderWidth"`
	MaskClip                       className  `json:"maskClip"`
	MaskComposite                  className  `json:"maskComposite"`
	MaskImage                      className  `json:"maskImage"`
	MaskMode                       className  `json:"maskMode"`
	MaskOrigin                     className  `json:"maskOrigin"`
	MaskPosition                   className  `json:"maskPosition"`
	MaskRepeat                     className  `json:"maskRepeat"`
	MaskSize                       className  `json:"maskSize"`
	MaskType                       className  `json:"maskType"`
	MasonryAutoFlow                className  `json:"masonryAutoFlow"`
	MathDepth                      className  `json:"mathDepth"`
	MathShift                      className  `json:"mathShift"`
	MathStyle                      className  `json:"mathStyle"`
	MaxBlockSize                   className  `json:"maxBlockSize"`
	MaxHeight                      className  `json:"maxHeight"`
	MaxInlineSize                  className  `json:"maxInlineSize"`
	MaxLines                       className  `json:"maxLines"`
	MaxWidth                       className  `json:"maxWidth"`
	MinBlockSize                   className  `json:"minBlockSize"`
	MinHeight                      className  `json:"minHeight"`
	MinInlineSize                  className  `json:"minInlineSize"`
	MinWidth                       className  `json:"minWidth"`
	MixBlendMode                   className  `json:"mixBlendMode"`
	MotionDistance                 className  `json:"motionDistance"`
	MotionPath                     className  `json:"motionPath"`
	MotionRotation                 className  `json:"motionRotation"`
	ObjectFit                      className  `json:"objectFit"`
	ObjectPosition                 className  `json:"objectPosition"`
	OffsetAnchor                   className  `json:"offsetAnchor"`
	OffsetDistance                 className  `json:"offsetDistance"`
	OffsetPath                     className  `json:"offsetPath"`
	OffsetPosition                 className  `json:"offsetPosition"`
	OffsetRotate                   className  `json:"offsetRotate"`
	OffsetRotation                 className  `json:"offsetRotation"`
	Order                          className  `json:"order"`
	Orphans                        className  `json:"orphans"`
	OutlineColor                   className  `json:"outlineColor"`
	OutlineOffset                  className  `json:"outlineOffset"`
	OutlineStyle                   className  `json:"outlineStyle"`
	OutlineWidth                   className  `json:"outlineWidth"`
	OverflowAnchor                 className  `json:"overflowAnchor"`
	OverflowBlock                  className  `json:"overflowBlock"`
	OverflowClipBox                className  `json:"overflowClipBox"`
	OverflowClipMargin             className  `json:"overflowClipMargin"`
	OverflowInline                 className  `json:"overflowInline"`
	OverflowWrap                   className  `json:"overflowWrap"`
	OverflowX                      className  `json:"overflowX"`
	OverflowY                      className  `json:"overflowY"`
	Overlay                        className  `json:"overlay"`
	OverscrollBehaviorBlock        className  `json:"overscrollBehaviorBlock"`
	OverscrollBehaviorInline       className  `json:"overscrollBehaviorInline"`
	OverscrollBehaviorX            className  `json:"overscrollBehaviorX"`
	OverscrollBehaviorY            className  `json:"overscrollBehaviorY"`
	PaddingBlockEnd                className  `json:"paddingBlockEnd"`
	PaddingBlockStart              className  `json:"paddingBlockStart"`
	PaddingBottom                  className  `json:"paddingBottom"`
	PaddingInlineEnd               className  `json:"paddingInlineEnd"`
	PaddingInlineStart             className  `json:"paddingInlineStart"`
	PaddingLeft                    className  `json:"paddingLeft"`
	PaddingRight                   className  `json:"paddingRight"`
	PaddingTop                     className  `json:"paddingTop"`
	Page                           className  `json:"page"`
	PageBreakAfter                 className  `json:"pageBreakAfter"`
	PageBreakBefore                className  `json:"pageBreakBefore"`
	PageBreakInside                className  `json:"pageBreakInside"`
	Perspective                    className  `json:"perspective"`
	PerspectiveOrigin              className  `json:"perspectiveOrigin"`
	Position                       className  `json:"position"`
	PrintColorAdjust               className  `json:"printColorAdjust"`
	Quotes                         className  `json:"quotes"`
	Resize                         className  `json:"resize"`
	Right                          className  `json:"right"`
	Rotate                         className  `json:"rotate"`
	RowGap                         className  `json:"rowGap"`
	RubyAlign                      className  `json:"rubyAlign"`
	RubyMerge                      className  `json:"rubyMerge"`
	RubyPosition                   className  `json:"rubyPosition"`
	Scale                          className  `json:"scale"`
	ScrollBehavior                 className  `json:"scrollBehavior"`
	ScrollMarginBlockEnd           className  `json:"scrollMarginBlockEnd"`
	ScrollMarginBlockStart         className  `json:"scrollMarginBlockStart"`
	ScrollMarginBottom             className  `json:"scrollMarginBottom"`
	ScrollMarginInlineEnd          className  `json:"scrollMarginInlineEnd"`
	ScrollMarginInlineStart        className  `json:"scrollMarginInlineStart"`
	ScrollMarginLeft               className  `json:"scrollMarginLeft"`
	ScrollMarginRight              className  `json:"scrollMarginRight"`
	ScrollMarginTop                className  `json:"scrollMarginTop"`
	ScrollPaddingBlockEnd          className  `json:"scrollPaddingBlockEnd"`
	ScrollPaddingBlockStart        className  `json:"scrollPaddingBlockStart"`
	ScrollPaddingBottom            className  `json:"scrollPaddingBottom"`
	ScrollPaddingInlineEnd         className  `json:"scrollPaddingInlineEnd"`
	ScrollPaddingInlineStart       className  `json:"scrollPaddingInlineStart"`
	ScrollPaddingLeft              className  `json:"scrollPaddingLeft"`
	ScrollPaddingRight             className  `json:"scrollPaddingRight"`
	ScrollPaddingTop               className  `json:"scrollPaddingTop"`
	ScrollSnapAlign                className  `json:"scrollSnapAlign"`
	ScrollSnapMarginBottom         className  `json:"scrollSnapMarginBottom"`
	ScrollSnapMarginLeft           className  `json:"scrollSnapMarginLeft"`
	ScrollSnapMarginRight          className  `json:"scrollSnapMarginRight"`
	ScrollSnapMarginTop            className  `json:"scrollSnapMarginTop"`
	ScrollSnapStop                 className  `json:"scrollSnapStop"`
	ScrollSnapType                 className  `json:"scrollSnapType"`
	ScrollTimelineAxis             className  `json:"scrollTimelineAxis"`
	ScrollTimelineName             className  `json:"scrollTimelineName"`
	ScrollbarColor                 className  `json:"scrollbarColor"`
	ScrollbarGutter                className  `json:"scrollbarGutter"`
	ScrollbarWidth                 className  `json:"scrollbarWidth"`
	ShapeImageThreshold            className  `json:"shapeImageThreshold"`
	ShapeMargin                    className  `json:"shapeMargin"`
	ShapeOutside                   className  `json:"shapeOutside"`
	TabSize                        className  `json:"tabSize"`
	TableLayout                    className  `json:"tableLayout"`
	TextAlign                      className  `json:"textAlign"`
	TextAlignLast                  className  `json:"textAlignLast"`
	TextCombineUpright             className  `json:"textCombineUpright"`
	TextDecorationColor            className  `json:"textDecorationColor"`
	TextDecorationLine             className  `json:"textDecorationLine"`
	TextDecorationSkip             className  `json:"textDecorationSkip"`
	TextDecorationSkipInk          className  `json:"textDecorationSkipInk"`
	TextDecorationStyle            className  `json:"textDecorationStyle"`
	TextDecorationThickness        className  `json:"textDecorationThickness"`
	TextEmphasisColor              className  `json:"textEmphasisColor"`
	TextEmphasisPosition           className  `json:"textEmphasisPosition"`
	TextEmphasisStyle              className  `json:"textEmphasisStyle"`
	TextIndent                     className  `json:"textIndent"`
	TextJustify                    className  `json:"textJustify"`
	TextOrientation                className  `json:"textOrientation"`
	TextOverflow                   className  `json:"textOverflow"`
	TextShadow                     className  `json:"textShadow"`
	TextSizeAdjust                 className  `json:"textSizeAdjust"`
	TextTransform                  className  `json:"textTransform"`
	TextUnderlineOffset            className  `json:"textUnderlineOffset"`
	TextUnderlinePosition          className  `json:"textUnderlinePosition"`
	TextWrap                       className  `json:"textWrap"`
	TimelineScope                  className  `json:"timelineScope"`
	Top                            className  `json:"top"`
	TouchAction                    className  `json:"touchAction"`
	Transform                      className  `json:"transform"`
	TransformBox                   className  `json:"transformBox"`
	TransformOrigin                className  `json:"transformOrigin"`
	TransformStyle                 className  `json:"transformStyle"`
	TransitionBehavior             className  `json:"transitionBehavior"`
	TransitionDelay                className  `json:"transitionDelay"`
	TransitionDuration             className  `json:"transitionDuration"`
	TransitionProperty             className  `json:"transitionProperty"`
	TransitionTimingFunction       className  `json:"transitionTimingFunction"`
	Translate                      className  `json:"translate"`
	UserSelect                     className  `json:"userSelect"`
	VerticalAlign                  className  `json:"verticalAlign"`
	ViewTimelineAxis               className  `json:"viewTimelineAxis"`
	ViewTimelineInset              className  `json:"viewTimelineInset"`
	ViewTimelineName               className  `json:"viewTimelineName"`
	ViewTransitionName             className  `json:"viewTransitionName"`
	WhiteSpaceCollapse             className  `json:"whiteSpaceCollapse"`
	WhiteSpaceTrim                 className  `json:"whiteSpaceTrim"`
	Widows                         className  `json:"widows"`
	Width                          className  `json:"width"`
	WillChange                     className  `json:"willChange"`
	WordBreak                      className  `json:"wordBreak"`
	WordWrap                       className  `json:"wordWrap"`
	ZIndex                         className  `json:"zIndex"`
	Zoom                           className  `json:"zoom"`
}

type ReloadActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               className       `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
}

type RemarkSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Label                string        `json:"label"`
	Icon                 RootSchema    `json:"icon"`
	TooltipClassName     RootSchema    `json:"tooltipClassName"`
	Trigger              []string      `json:"trigger"`
	Title                string        `json:"title"`
	Content              className     `json:"content"`
	Placement            string        `json:"placement"`
	RootClose            bool          `json:"rootClose"`
	Shape                string        `json:"shape"`
}

type RepeatControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Options              string           `json:"options"`
}

type RichTextControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Vendor               string           `json:"vendor"`
	Receiver             className        `json:"receiver"`
	VideoReceiver        className        `json:"videoReceiver"`
	FileField            string           `json:"fileField"`
	BorderMode           string           `json:"borderMode"`
	Options              mock             `json:"options"`
}

type RootConfig struct {
	Type__       string       `json:"__type__"`
	Body         []ArrayElem1 `json:"body"`
	Type_ignores []any        `json:"type_ignores"`
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

type RootConfig2 struct {
	Schema      string      `json:"$schema"`
	Ref         string      `json:"$ref"`
	Definitions definitions `json:"definitions"`
}

type RootConfig3 struct {
	Fsq_id               string        `json:"fsq_id"`
	Categories           []ArrayElem36 `json:"categories"`
	Chains               []trackConfig `json:"chains"`
	Closed_bucket        string        `json:"closed_bucket"`
	Date_closed          string        `json:"date_closed"`
	Description          string        `json:"description"`
	Distance             int           `json:"distance"`
	Email                string        `json:"email"`
	Fax                  string        `json:"fax"`
	Features             features      `json:"features"`
	Geocodes             geocodes      `json:"geocodes"`
	Hours                hours         `json:"hours"`
	Hours_popular        []ArrayElem37 `json:"hours_popular"`
	Link                 string        `json:"link"`
	Location             location      `json:"location"`
	Menu                 string        `json:"menu"`
	Name                 string        `json:"name"`
	Photos               []icon        `json:"photos"`
	Popularity           int           `json:"popularity"`
	Price                int           `json:"price"`
	Rating               int           `json:"rating"`
	Related_places       staticSchema  `json:"related_places"`
	Social_media         social_media  `json:"social_media"`
	Stats                stats         `json:"stats"`
	Store_id             string        `json:"store_id"`
	Tastes               []string      `json:"tastes"`
	Tel                  string        `json:"tel"`
	Timezone             string        `json:"timezone"`
	Tips                 []tip         `json:"tips"`
	Venue_reality_bucket string        `json:"venue_reality_bucket"`
	Verified             bool          `json:"verified"`
	Website              string        `json:"website"`
}

type RootConfig4 struct {
	Report     report `json:"report"`
	Request_id string `json:"request_id"`
	Warnings   []any  `json:"warnings"`
}

type RootConfig5 struct {
	Id                  string           `json:"@id"`
	TypeVal             string           `json:"@type"`
	Cwa                 string           `json:"cwa"`
	ForecastOffice      string           `json:"forecastOffice"`
	GridId              string           `json:"gridId"`
	GridX               int              `json:"gridX"`
	GridY               int              `json:"gridY"`
	Forecast            string           `json:"forecast"`
	ForecastHourly      string           `json:"forecastHourly"`
	ForecastGridData    string           `json:"forecastGridData"`
	ObservationStations string           `json:"observationStations"`
	RelativeLocation    relativeLocation `json:"relativeLocation"`
	ForecastZone        string           `json:"forecastZone"`
	County              string           `json:"county"`
	FireWeatherZone     string           `json:"fireWeatherZone"`
	TimeZone            string           `json:"timeZone"`
	RadarStation        string           `json:"radarStation"`
}

type RootConfig6 struct {
	Name              string        `json:"name"`
	Version           string        `json:"version"`
	BuildNumber       string        `json:"buildNumber"`
	ProductCode       string        `json:"productCode"`
	EnvVarBaseName    string        `json:"envVarBaseName"`
	DataDirectoryName string        `json:"dataDirectoryName"`
	SvgIconPath       string        `json:"svgIconPath"`
	ProductVendor     string        `json:"productVendor"`
	Launch            []ArrayElem46 `json:"launch"`
	CustomProperties  []ArrayElem47 `json:"customProperties"`
	BundledPlugins    []string      `json:"bundledPlugins"`
	Modules           []string      `json:"modules"`
	FileExtensions    []string      `json:"fileExtensions"`
	Layout            []ArrayElem48 `json:"layout"`
}

type RootSchema struct {
	Ref string `json:"$ref"`
}

type RowSelectionOptionsSchema struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

type RowSelectionSchema struct {
	TypeVal             string  `json:"type"`
	KeyField            string  `json:"keyField"`
	DisableOn           string  `json:"disableOn"`
	Selections          []any   `json:"selections"`
	SelectedRowKeys     []any   `json:"selectedRowKeys"`
	SelectedRowKeysExpr string  `json:"selectedRowKeysExpr"`
	ColumnWidth         float64 `json:"columnWidth"`
	RowClick            bool    `json:"rowClick"`
}

type SchemaApiObject struct {
	Method                 string    `json:"method"`
	Url                    className `json:"url"`
	Data                   any       `json:"data"`
	ConvertKeyToPath       bool      `json:"convertKeyToPath"`
	ResponseData           any       `json:"responseData"`
	AttachDataToQuery      bool      `json:"attachDataToQuery"`
	DataType               string    `json:"dataType"`
	ResponseType           string    `json:"responseType"`
	Headers                any       `json:"headers"`
	SendOn                 className `json:"sendOn"`
	ReplaceData            bool      `json:"replaceData"`
	AutoRefresh            bool      `json:"autoRefresh"`
	TrackExpression        string    `json:"trackExpression"`
	Cache                  float64   `json:"cache"`
	ForceAppendDataToQuery bool      `json:"forceAppendDataToQuery"`
	QsOptions              qsOptions `json:"qsOptions"`
	Silent                 bool      `json:"silent"`
	DownloadFileName       string    `json:"downloadFileName"`
}

type SchemaClassName struct {
	AnyOf       []ArrayElem2 `json:"anyOf"`
	Description string       `json:"description"`
}

type SchemaCopyableObject struct {
	Icon    className `json:"icon"`
	Content className `json:"content"`
	Tooltip string    `json:"tooltip"`
}

type SchemaEditorItemPlaceholder struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DefaultVal  string `json:"default"`
	Empty       string `json:"empty"`
}

type SchemaMessage struct {
	FetchFailed  string `json:"fetchFailed"`
	FetchSuccess string `json:"fetchSuccess"`
	SaveFailed   string `json:"saveFailed"`
	SaveSuccess  string `json:"saveSuccess"`
}

type SchemaObject struct {
	AllOf []ArrayElem4 `json:"allOf"`
}

type SchemaObjectCardBodyField struct {
	AnyOf []ArrayElem29 `json:"anyOf"`
}

type SchemaObjectComboSubControl struct {
	AnyOf []ArrayElem34 `json:"anyOf"`
}

type SchemaObjectGrid struct {
	AnyOf []ArrayElem32 `json:"anyOf"`
}

type SchemaObjectGroupSubControl struct {
	AnyOf []ArrayElem35 `json:"anyOf"`
}

type SchemaObjectHboxRow struct {
	AnyOf []ArrayElem33 `json:"anyOf"`
}

type SchemaObjectListBodyField struct {
	AnyOf []ArrayElem30 `json:"anyOf"`
}

type SchemaObjectTableColumnWithType struct {
	AnyOf []ArrayElem31 `json:"anyOf"`
}

type SchemaPopOverObject struct {
	ClassName        string     `json:"className"`
	PopOverClassName string     `json:"popOverClassName"`
	PopOverEnableOn  className  `json:"popOverEnableOn"`
	Mode             string     `json:"mode"`
	Size             string     `json:"size"`
	Position         string     `json:"position"`
	Trigger          string     `json:"trigger"`
	ShowIcon         bool       `json:"showIcon"`
	Offset           offset     `json:"offset"`
	Title            string     `json:"title"`
	Body             RootSchema `json:"body"`
}

type SchemaQuickEditObject struct {
	AnyOf []ArrayElem9 `json:"anyOf"`
}

type SearchBoxSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Name                 string        `json:"name"`
	Placeholder          string        `json:"placeholder"`
	Mini                 bool          `json:"mini"`
	Enhance              bool          `json:"enhance"`
	Clearable            bool          `json:"clearable"`
	SearchImediately     bool          `json:"searchImediately"`
	ClearAndSubmit       bool          `json:"clearAndSubmit"`
	Loading              bool          `json:"loading"`
}

type SelectControlSchema struct {
	LoadingConfig        loadingConfig    `json:"loadingConfig"`
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	AutoComplete         className        `json:"autoComplete"`
	MenuTpl              string           `json:"menuTpl"`
	ShowInvalidMatch     bool             `json:"showInvalidMatch"`
	BorderMode           string           `json:"borderMode"`
	SelectMode           string           `json:"selectMode"`
	LeftOptions          []any            `json:"leftOptions"`
	LeftMode             string           `json:"leftMode"`
	RightMode            string           `json:"rightMode"`
	SearchResultMode     string           `json:"searchResultMode"`
	Columns              []any            `json:"columns"`
	SearchResultColumns  []any            `json:"searchResultColumns"`
	Searchable           bool             `json:"searchable"`
	SearchApi            className        `json:"searchApi"`
	ItemHeight           float64          `json:"itemHeight"`
	VirtualThreshold     float64          `json:"virtualThreshold"`
	CheckAll             bool             `json:"checkAll"`
	DefaultCheckAll      bool             `json:"defaultCheckAll"`
	CheckAllLabel        string           `json:"checkAllLabel"`
	MaxTagCount          float64          `json:"maxTagCount"`
	OverflowTagPopover   any              `json:"overflowTagPopover"`
	OptionClassName      className        `json:"optionClassName"`
	Overlay              overlay          `json:"overlay"`
}

type ServiceSchema struct {
	LoadingConfig        loadingConfig `json:"loadingConfig"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Api                  className     `json:"api"`
	Ws                   string        `json:"ws"`
	DataProvider         className     `json:"dataProvider"`
	Body                 className     `json:"body"`
	FetchOn              RootSchema    `json:"fetchOn"`
	InitFetch            bool          `json:"initFetch"`
	InitFetchOn          className     `json:"initFetchOn"`
	SchemaApi            className     `json:"schemaApi"`
	InitFetchSchema      bool          `json:"initFetchSchema"`
	InitFetchSchemaOn    className     `json:"initFetchSchemaOn"`
	Interval             float64       `json:"interval"`
	SilentPolling        bool          `json:"silentPolling"`
	StopAutoRefreshWhen  className     `json:"stopAutoRefreshWhen"`
	Messages             RootSchema    `json:"messages"`
	Name                 RootSchema    `json:"name"`
	ShowErrorMsg         bool          `json:"showErrorMsg"`
}

type ShortCutDate struct {
	Label string          `json:"label"`
	Date  draggableConfig `json:"date"`
}

type ShortCutDateRange struct {
	Label     string `json:"label"`
	StartDate right  `json:"startDate"`
	EndDate   right  `json:"endDate"`
}

type ShortCuts struct {
	AnyOf []ArrayElem26 `json:"anyOf"`
}

type SparkLineSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Name                 string        `json:"name"`
	Width                float64       `json:"width"`
	Height               float64       `json:"height"`
	ClickAction          className     `json:"clickAction"`
	Placeholder          string        `json:"placeholder"`
	Value                []any         `json:"value"`
}

type SpinnerSchema struct {
	LoadingConfig        loadingConfig `json:"loadingConfig"`
	ID                   string        `json:"$$id"`
	ClassName            string        `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Show                 bool          `json:"show"`
	SpinnerClassName     string        `json:"spinnerClassName"`
	SpinnerWrapClassName string        `json:"spinnerWrapClassName"`
	Mode                 string        `json:"mode"`
	Size                 string        `json:"size"`
	Icon                 string        `json:"icon"`
	Tip                  string        `json:"tip"`
	TipPlacement         string        `json:"tipPlacement"`
	Delay                float64       `json:"delay"`
	Overlay              bool          `json:"overlay"`
	Body                 className     `json:"body"`
}

type StateSchema struct {
	Testid               string        `json:"testid"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            string        `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	Title                string        `json:"title"`
	Body                 className     `json:"body"`
}

type StaticExactControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Tpl                  className        `json:"tpl"`
	Text                 className        `json:"text"`
	PopOver              className        `json:"popOver"`
	QuickEdit            className        `json:"quickEdit"`
	Copyable             className        `json:"copyable"`
	BorderMode           string           `json:"borderMode"`
}

type StatusSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Placeholder          string        `json:"placeholder"`
	MapVal               any           `json:"map"`
	LabelMap             any           `json:"labelMap"`
	Source               className     `json:"source"`
}

type StepSchema struct {
	Testid               string          `json:"testid"`
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	Title                SchemaClassName `json:"title"`
	SubTitle             SchemaClassName `json:"subTitle"`
	Icon                 string          `json:"icon"`
	Value                any             `json:"value"`
	Description          SchemaClassName `json:"description"`
}

type StepsSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Steps                []any         `json:"steps"`
	Source               string        `json:"source"`
	Value                any           `json:"value"`
	Name                 string        `json:"name"`
	Status               right         `json:"status"`
	Mode                 string        `json:"mode"`
	LabelPlacement       string        `json:"labelPlacement"`
	ProgressDot          bool          `json:"progressDot"`
}

type SubFormControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Multiple             bool             `json:"multiple"`
	Draggable            bool             `json:"draggable"`
	DraggableTip         string           `json:"draggableTip"`
	Addable              bool             `json:"addable"`
	Removable            bool             `json:"removable"`
	MinLength            float64          `json:"minLength"`
	MaxLength            float64          `json:"maxLength"`
	LabelField           string           `json:"labelField"`
	BtnLabel             string           `json:"btnLabel"`
	AddButtonText        string           `json:"addButtonText"`
	AddButtonClassName   className        `json:"addButtonClassName"`
	ItemClassName        className        `json:"itemClassName"`
	ItemsClassName       className        `json:"itemsClassName"`
	ShowErrorMsg         bool             `json:"showErrorMsg"`
	Form                 form             `json:"form"`
	Scaffold             staticSchema     `json:"scaffold"`
}

type SwitchContainerSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Items                []any         `json:"items"`
}

type SwitchControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	TrueValue            any              `json:"trueValue"`
	FalseValue           any              `json:"falseValue"`
	Option               string           `json:"option"`
	OnText               SchemaClassName  `json:"onText"`
	OffText              SchemaClassName  `json:"offText"`
	Loading              bool             `json:"loading"`
}

type TabSchema struct {
	Testid               string          `json:"testid"`
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	Title                SchemaClassName `json:"title"`
	Tab                  className       `json:"tab"`
	Body                 className       `json:"body"`
	Badge                float64         `json:"badge"`
	Hash                 string          `json:"hash"`
	Icon                 className       `json:"icon"`
	IconPosition         string          `json:"iconPosition"`
	Reload               bool            `json:"reload"`
	MountOnEnter         bool            `json:"mountOnEnter"`
	UnmountOnExit        bool            `json:"unmountOnExit"`
	Mode                 string          `json:"mode"`
	Horizontal           className       `json:"horizontal"`
	Closable             bool            `json:"closable"`
}

type TableColumnObject struct {
	Label              string          `json:"label"`
	Fixed              string          `json:"fixed"`
	Name               string          `json:"name"`
	PopOver            className       `json:"popOver"`
	QuickEdit          className       `json:"quickEdit"`
	QuickEditOnUpdate  className       `json:"quickEditOnUpdate"`
	Copyable           className       `json:"copyable"`
	Sortable           bool            `json:"sortable"`
	Searchable         SchemaClassName `json:"searchable"`
	Toggled            bool            `json:"toggled"`
	Width              any             `json:"width"`
	Align              string          `json:"align"`
	VAlign             string          `json:"vAlign"`
	HeaderAlign        string          `json:"headerAlign"`
	ClassName          string          `json:"className"`
	ClassNameExpr      string          `json:"classNameExpr"`
	LabelClassName     string          `json:"labelClassName"`
	Filterable         SchemaClassName `json:"filterable"`
	Breakpoint         string          `json:"breakpoint"`
	Remark             className       `json:"remark"`
	Value              mock            `json:"value"`
	Unique             bool            `json:"unique"`
	CanAccessSuperData bool            `json:"canAccessSuperData"`
	LazyRenderAfter    float64         `json:"lazyRenderAfter"`
	InnerStyle         any             `json:"innerStyle"`
}

type TableColumnWithType struct {
	AllOf []ArrayElem15 `json:"allOf"`
}

type TableControlSchema struct {
	AffixHeader           bool             `json:"affixHeader"`
	AffixFooter           bool             `json:"affixFooter"`
	Columns               []any            `json:"columns"`
	ColumnsTogglable      SchemaClassName  `json:"columnsTogglable"`
	Footable              SchemaClassName  `json:"footable"`
	FooterClassName       className        `json:"footerClassName"`
	HeaderClassName       className        `json:"headerClassName"`
	Placeholder           string           `json:"placeholder"`
	ShowFooter            bool             `json:"showFooter"`
	ShowHeader            bool             `json:"showHeader"`
	Source                className        `json:"source"`
	TableClassName        className        `json:"tableClassName"`
	Title                 string           `json:"title"`
	ToolbarClassName      className        `json:"toolbarClassName"`
	CombineNum            SchemaClassName  `json:"combineNum"`
	CombineFromIndex      float64          `json:"combineFromIndex"`
	PrefixRow             []any            `json:"prefixRow"`
	AffixRow              []any            `json:"affixRow"`
	Resizable             bool             `json:"resizable"`
	RowClassNameExpr      string           `json:"rowClassNameExpr"`
	ItemBadge             className        `json:"itemBadge"`
	AutoGenerateFilter    draggableConfig  `json:"autoGenerateFilter"`
	CanAccessSuperData    bool             `json:"canAccessSuperData"`
	AutoFillHeight        SchemaClassName  `json:"autoFillHeight"`
	TableLayout           string           `json:"tableLayout"`
	DeferApi              className        `json:"deferApi"`
	Testid                string           `json:"testid"`
	ID                    string           `json:"$$id"`
	ClassName             className        `json:"className"`
	Ref                   string           `json:"$ref"`
	Disabled              bool             `json:"disabled"`
	DisabledOn            className        `json:"disabledOn"`
	Hidden                bool             `json:"hidden"`
	HiddenOn              className        `json:"hiddenOn"`
	Visible               bool             `json:"visible"`
	VisibleOn             className        `json:"visibleOn"`
	Id                    string           `json:"id"`
	OnEvent               any              `json:"onEvent"`
	Static                bool             `json:"static"`
	StaticOn              className        `json:"staticOn"`
	StaticPlaceholder     string           `json:"staticPlaceholder"`
	StaticClassName       className        `json:"staticClassName"`
	StaticLabelClassName  className        `json:"staticLabelClassName"`
	StaticInputClassName  className        `json:"staticInputClassName"`
	StaticSchema          staticSchema     `json:"staticSchema"`
	Style                 any              `json:"style"`
	EditorSetting         editorSetting    `json:"editorSetting"`
	UseMobileUI           bool             `json:"useMobileUI"`
	TestIdBuilder         RootSchema       `json:"testIdBuilder"`
	Size                  string           `json:"size"`
	Label                 SchemaClassName  `json:"label"`
	LabelAlign            className        `json:"labelAlign"`
	LabelWidth            any              `json:"labelWidth"`
	LabelClassName        string           `json:"labelClassName"`
	Name                  string           `json:"name"`
	ExtraName             string           `json:"extraName"`
	Remark                mock             `json:"remark"`
	LabelRemark           mock             `json:"labelRemark"`
	Hint                  string           `json:"hint"`
	SubmitOnChange        bool             `json:"submitOnChange"`
	ReadOnly              bool             `json:"readOnly"`
	ReadOnlyOn            string           `json:"readOnlyOn"`
	ValidateOnChange      bool             `json:"validateOnChange"`
	Description           string           `json:"description"`
	Desc                  string           `json:"desc"`
	DescriptionClassName  className        `json:"descriptionClassName"`
	Mode                  string           `json:"mode"`
	Horizontal            className        `json:"horizontal"`
	Inline                bool             `json:"inline"`
	InputClassName        className        `json:"inputClassName"`
	Required              bool             `json:"required"`
	ValidationErrors      validationErrors `json:"validationErrors"`
	Validations           data             `json:"validations"`
	Value                 mock             `json:"value"`
	ClearValueOnHidden    bool             `json:"clearValueOnHidden"`
	ValidateApi           SchemaClassName  `json:"validateApi"`
	AutoFill              autoFill         `json:"autoFill"`
	InitAutoFill          data             `json:"initAutoFill"`
	Row                   float64          `json:"row"`
	TypeVal               string           `json:"type"`
	Addable               bool             `json:"addable"`
	ChildrenAddable       bool             `json:"childrenAddable"`
	Copyable              bool             `json:"copyable"`
	CopyBtnLabel          string           `json:"copyBtnLabel"`
	CopyBtnIcon           string           `json:"copyBtnIcon"`
	CopyAddBtn            bool             `json:"copyAddBtn"`
	CopyData              any              `json:"copyData"`
	Draggable             bool             `json:"draggable"`
	AddApi                className        `json:"addApi"`
	AddBtnLabel           string           `json:"addBtnLabel"`
	AddBtnIcon            string           `json:"addBtnIcon"`
	SubAddBtnLabel        string           `json:"subAddBtnLabel"`
	SubAddBtnIcon         string           `json:"subAddBtnIcon"`
	Removable             bool             `json:"removable"`
	DeleteApi             className        `json:"deleteApi"`
	Editable              bool             `json:"editable"`
	EditBtnLabel          string           `json:"editBtnLabel"`
	EditBtnIcon           string           `json:"editBtnIcon"`
	ConfirmBtnLabel       string           `json:"confirmBtnLabel"`
	ConfirmBtnIcon        string           `json:"confirmBtnIcon"`
	CancelBtnLabel        string           `json:"cancelBtnLabel"`
	CancelBtnIcon         string           `json:"cancelBtnIcon"`
	DeleteBtnLabel        string           `json:"deleteBtnLabel"`
	DeleteBtnIcon         string           `json:"deleteBtnIcon"`
	UpdateApi             className        `json:"updateApi"`
	Scaffold              mock             `json:"scaffold"`
	DeleteConfirmText     string           `json:"deleteConfirmText"`
	ValueField            string           `json:"valueField"`
	NeedConfirm           bool             `json:"needConfirm"`
	ShowIndex             bool             `json:"showIndex"`
	PerPage               float64          `json:"perPage"`
	MaxLength             SchemaClassName  `json:"maxLength"`
	MinLength             SchemaClassName  `json:"minLength"`
	ShowFooterAddBtn      bool             `json:"showFooterAddBtn"`
	ShowTableAddBtn       bool             `json:"showTableAddBtn"`
	FooterAddBtn          className        `json:"footerAddBtn"`
	EnableStaticTransform bool             `json:"enableStaticTransform"`
}

type TableSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	AffixHeader          bool            `json:"affixHeader"`
	AffixFooter          bool            `json:"affixFooter"`
	Columns              []any           `json:"columns"`
	ColumnsTogglable     SchemaClassName `json:"columnsTogglable"`
	Footable             SchemaClassName `json:"footable"`
	FooterClassName      className       `json:"footerClassName"`
	HeaderClassName      className       `json:"headerClassName"`
	Placeholder          SchemaClassName `json:"placeholder"`
	ShowFooter           bool            `json:"showFooter"`
	ShowHeader           bool            `json:"showHeader"`
	Source               className       `json:"source"`
	TableClassName       className       `json:"tableClassName"`
	Title                string          `json:"title"`
	ToolbarClassName     className       `json:"toolbarClassName"`
	CombineNum           SchemaClassName `json:"combineNum"`
	CombineFromIndex     float64         `json:"combineFromIndex"`
	PrefixRow            []any           `json:"prefixRow"`
	AffixRow             []any           `json:"affixRow"`
	Resizable            bool            `json:"resizable"`
	RowClassNameExpr     string          `json:"rowClassNameExpr"`
	ItemBadge            className       `json:"itemBadge"`
	AutoGenerateFilter   draggableConfig `json:"autoGenerateFilter"`
	CanAccessSuperData   bool            `json:"canAccessSuperData"`
	AutoFillHeight       SchemaClassName `json:"autoFillHeight"`
	TableLayout          string          `json:"tableLayout"`
	DeferApi             className       `json:"deferApi"`
}

type TableSchema2 struct {
	ID                            string           `json:"$$id"`
	ClassName                     className        `json:"className"`
	Ref                           string           `json:"$ref"`
	Disabled                      bool             `json:"disabled"`
	DisabledOn                    className        `json:"disabledOn"`
	Hidden                        bool             `json:"hidden"`
	HiddenOn                      className        `json:"hiddenOn"`
	Visible                       bool             `json:"visible"`
	VisibleOn                     className        `json:"visibleOn"`
	Id                            string           `json:"id"`
	OnEvent                       any              `json:"onEvent"`
	Static                        bool             `json:"static"`
	StaticOn                      className        `json:"staticOn"`
	StaticPlaceholder             string           `json:"staticPlaceholder"`
	StaticClassName               className        `json:"staticClassName"`
	StaticLabelClassName          className        `json:"staticLabelClassName"`
	StaticInputClassName          className        `json:"staticInputClassName"`
	StaticSchema                  staticSchema     `json:"staticSchema"`
	Style                         any              `json:"style"`
	EditorSetting                 editorSetting    `json:"editorSetting"`
	UseMobileUI                   bool             `json:"useMobileUI"`
	TestIdBuilder                 RootSchema       `json:"testIdBuilder"`
	TypeVal                       string           `json:"type"`
	Testid                        string           `json:"testid"`
	Title                         SchemaClassName  `json:"title"`
	Source                        className        `json:"source"`
	ColumnsTogglable              columnsTogglable `json:"columnsTogglable"`
	Columns                       []any            `json:"columns"`
	RowSelection                  draggableConfig  `json:"rowSelection"`
	Expandable                    className        `json:"expandable"`
	Sticky                        bool             `json:"sticky"`
	Loading                       SchemaClassName  `json:"loading"`
	ItemBadge                     className        `json:"itemBadge"`
	ShowBadge                     bool             `json:"showBadge"`
	PopOverContainer              mock             `json:"popOverContainer"`
	KeyField                      string           `json:"keyField"`
	ChildrenColumnName            string           `json:"childrenColumnName"`
	RowClassNameExpr              string           `json:"rowClassNameExpr"`
	LineHeight                    string           `json:"lineHeight"`
	Bordered                      bool             `json:"bordered"`
	ShowHeader                    bool             `json:"showHeader"`
	Footer                        SchemaClassName  `json:"footer"`
	QuickSaveApi                  className        `json:"quickSaveApi"`
	QuickSaveItemApi              className        `json:"quickSaveItemApi"`
	Messages                      className        `json:"messages"`
	Reload                        string           `json:"reload"`
	Actions                       []any            `json:"actions"`
	MaxKeepItemSelectionLength    float64          `json:"maxKeepItemSelectionLength"`
	KeepItemSelectionOnPageChange bool             `json:"keepItemSelectionOnPageChange"`
	Selectable                    bool             `json:"selectable"`
	Multiple                      bool             `json:"multiple"`
	PrimaryField                  string           `json:"primaryField"`
	TableLayout                   string           `json:"tableLayout"`
	AutoFillHeight                SchemaClassName  `json:"autoFillHeight"`
	CanAccessSuperData            bool             `json:"canAccessSuperData"`
	LazyRenderAfter               float64          `json:"lazyRenderAfter"`
}

type TabsSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Tabs                 []any           `json:"tabs"`
	Source               string          `json:"source"`
	TabsMode             className       `json:"tabsMode"`
	ContentClassName     className       `json:"contentClassName"`
	LinksClassName       className       `json:"linksClassName"`
	MountOnEnter         bool            `json:"mountOnEnter"`
	UnmountOnExit        bool            `json:"unmountOnExit"`
	Toolbar              className       `json:"toolbar"`
	SubFormMode          string          `json:"subFormMode"`
	SubFormHorizontal    className       `json:"subFormHorizontal"`
	Addable              bool            `json:"addable"`
	Closable             bool            `json:"closable"`
	Draggable            bool            `json:"draggable"`
	ShowTip              bool            `json:"showTip"`
	ShowTipClassName     string          `json:"showTipClassName"`
	Editable             bool            `json:"editable"`
	Scrollable           bool            `json:"scrollable"`
	SidePosition         string          `json:"sidePosition"`
	AddBtnText           string          `json:"addBtnText"`
	DefaultKey           draggableConfig `json:"defaultKey"`
	ActiveKey            draggableConfig `json:"activeKey"`
	CollapseOnExceed     float64         `json:"collapseOnExceed"`
	CollapseBtnLabel     string          `json:"collapseBtnLabel"`
	Swipeable            bool            `json:"swipeable"`
}

type TabsTransferControlSchema struct {
	LoadingConfig              loadingConfig    `json:"loadingConfig"`
	ShowArrow                  bool             `json:"showArrow"`
	Sortable                   bool             `json:"sortable"`
	SelectMode                 string           `json:"selectMode"`
	ResultListModeFollowSelect bool             `json:"resultListModeFollowSelect"`
	LeftOptions                []any            `json:"leftOptions"`
	LeftMode                   string           `json:"leftMode"`
	RightMode                  string           `json:"rightMode"`
	SearchResultMode           string           `json:"searchResultMode"`
	Columns                    []any            `json:"columns"`
	SearchResultColumns        []any            `json:"searchResultColumns"`
	Searchable                 bool             `json:"searchable"`
	ResultSearchable           bool             `json:"resultSearchable"`
	SearchApi                  className        `json:"searchApi"`
	SelectTitle                string           `json:"selectTitle"`
	ResultTitle                string           `json:"resultTitle"`
	MenuTpl                    className        `json:"menuTpl"`
	ValueTpl                   className        `json:"valueTpl"`
	SearchPlaceholder          string           `json:"searchPlaceholder"`
	ResultSearchPlaceholder    string           `json:"resultSearchPlaceholder"`
	Statistics                 bool             `json:"statistics"`
	ItemHeight                 float64          `json:"itemHeight"`
	VirtualThreshold           float64          `json:"virtualThreshold"`
	ShowInvalidMatch           bool             `json:"showInvalidMatch"`
	OnlyChildren               bool             `json:"onlyChildren"`
	InitiallyOpen              bool             `json:"initiallyOpen"`
	AutoCheckChildren          bool             `json:"autoCheckChildren"`
	Pagination                 pagination       `json:"pagination"`
	Remark                     className        `json:"remark"`
	LabelRemark                className        `json:"labelRemark"`
	Size                       string           `json:"size"`
	Label                      SchemaClassName  `json:"label"`
	LabelAlign                 className        `json:"labelAlign"`
	LabelWidth                 any              `json:"labelWidth"`
	LabelClassName             string           `json:"labelClassName"`
	Name                       string           `json:"name"`
	ExtraName                  string           `json:"extraName"`
	Hint                       string           `json:"hint"`
	SubmitOnChange             bool             `json:"submitOnChange"`
	ReadOnly                   bool             `json:"readOnly"`
	ReadOnlyOn                 string           `json:"readOnlyOn"`
	ValidateOnChange           bool             `json:"validateOnChange"`
	Description                string           `json:"description"`
	Desc                       string           `json:"desc"`
	DescriptionClassName       className        `json:"descriptionClassName"`
	Mode                       string           `json:"mode"`
	Horizontal                 className        `json:"horizontal"`
	Inline                     bool             `json:"inline"`
	InputClassName             className        `json:"inputClassName"`
	Placeholder                string           `json:"placeholder"`
	Required                   bool             `json:"required"`
	ValidationErrors           validationErrors `json:"validationErrors"`
	Validations                data             `json:"validations"`
	Value                      mock             `json:"value"`
	ClearValueOnHidden         bool             `json:"clearValueOnHidden"`
	ValidateApi                SchemaClassName  `json:"validateApi"`
	AutoFill                   autoFill         `json:"autoFill"`
	InitAutoFill               data             `json:"initAutoFill"`
	Row                        float64          `json:"row"`
	ID                         string           `json:"$$id"`
	ClassName                  className        `json:"className"`
	Ref                        string           `json:"$ref"`
	Disabled                   bool             `json:"disabled"`
	DisabledOn                 className        `json:"disabledOn"`
	Hidden                     bool             `json:"hidden"`
	HiddenOn                   className        `json:"hiddenOn"`
	Visible                    bool             `json:"visible"`
	VisibleOn                  className        `json:"visibleOn"`
	Id                         string           `json:"id"`
	OnEvent                    any              `json:"onEvent"`
	Static                     bool             `json:"static"`
	StaticOn                   className        `json:"staticOn"`
	StaticPlaceholder          string           `json:"staticPlaceholder"`
	StaticClassName            className        `json:"staticClassName"`
	StaticLabelClassName       className        `json:"staticLabelClassName"`
	StaticInputClassName       className        `json:"staticInputClassName"`
	StaticSchema               staticSchema     `json:"staticSchema"`
	Style                      any              `json:"style"`
	EditorSetting              editorSetting    `json:"editorSetting"`
	UseMobileUI                bool             `json:"useMobileUI"`
	TestIdBuilder              RootSchema       `json:"testIdBuilder"`
	Options                    options          `json:"options"`
	Source                     draggableConfig  `json:"source"`
	SelectFirst                bool             `json:"selectFirst"`
	InitFetchOn                string           `json:"initFetchOn"`
	InitFetch                  bool             `json:"initFetch"`
	Multiple                   bool             `json:"multiple"`
	JoinValues                 bool             `json:"joinValues"`
	Delimiter                  string           `json:"delimiter"`
	ValuesNoWrap               bool             `json:"valuesNoWrap"`
	ExtractValue               bool             `json:"extractValue"`
	Clearable                  bool             `json:"clearable"`
	ResetValue                 string           `json:"resetValue"`
	DeferField                 string           `json:"deferField"`
	DeferApi                   draggableConfig  `json:"deferApi"`
	AddApi                     draggableConfig  `json:"addApi"`
	AddControls                []any            `json:"addControls"`
	AddDialog                  className        `json:"addDialog"`
	Creatable                  bool             `json:"creatable"`
	CreateBtnLabel             string           `json:"createBtnLabel"`
	Editable                   bool             `json:"editable"`
	EditApi                    draggableConfig  `json:"editApi"`
	EditControls               []any            `json:"editControls"`
	EditDialog                 className        `json:"editDialog"`
	Removable                  bool             `json:"removable"`
	DeleteApi                  draggableConfig  `json:"deleteApi"`
	DeleteConfirmText          string           `json:"deleteConfirmText"`
	TypeVal                    string           `json:"type"`
}

type TagControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             draggableConfig  `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	OptionsTip           string           `json:"optionsTip"`
	Dropdown             bool             `json:"dropdown"`
	Max                  float64          `json:"max"`
	MaxTagLength         float64          `json:"maxTagLength"`
	MaxTagCount          float64          `json:"maxTagCount"`
	OverflowTagPopover   className        `json:"overflowTagPopover"`
	EnableBatchAdd       bool             `json:"enableBatchAdd"`
	Separator            string           `json:"separator"`
}

type TagSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Color                string        `json:"color"`
	Label                string        `json:"label"`
	DisplayMode          string        `json:"displayMode"`
	Icon                 className     `json:"icon"`
	Closable             bool          `json:"closable"`
	CloseIcon            className     `json:"closeIcon"`
	Checkable            bool          `json:"checkable"`
	Checked              bool          `json:"checked"`
}

type TasksSchema struct {
	LoadingConfig        loadingConfig `json:"loadingConfig"`
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	BtnClassName         RootSchema    `json:"btnClassName"`
	BtnText              string        `json:"btnText"`
	CheckApi             className     `json:"checkApi"`
	Interval             float64       `json:"interval"`
	Items                []ArrayItem2  `json:"items"`
	Name                 RootSchema    `json:"name"`
	OperationLabel       string        `json:"operationLabel"`
	ReSubmitApi          className     `json:"reSubmitApi"`
	RemarkLabel          string        `json:"remarkLabel"`
	RetryBtnClassName    className     `json:"retryBtnClassName"`
	RetryBtnText         string        `json:"retryBtnText"`
	StatusLabel          string        `json:"statusLabel"`
	StatusLabelMap       []string      `json:"statusLabelMap"`
	StatusTextMap        []string      `json:"statusTextMap"`
	SubmitApi            className     `json:"submitApi"`
	TableClassName       className     `json:"tableClassName"`
	TaskNameLabel        string        `json:"taskNameLabel"`
	InitialStatusCode    float64       `json:"initialStatusCode"`
	ReadyStatusCode      float64       `json:"readyStatusCode"`
	LoadingStatusCode    float64       `json:"loadingStatusCode"`
	CanRetryStatusCode   float64       `json:"canRetryStatusCode"`
	FinishStatusCode     float64       `json:"finishStatusCode"`
	ErrorStatusCode      float64       `json:"errorStatusCode"`
}

type TestIdBuilder struct {
	TestId string `json:"testId"`
}

type TextControlSchema struct {
	Options               options          `json:"options"`
	Source                draggableConfig  `json:"source"`
	SelectFirst           bool             `json:"selectFirst"`
	InitFetchOn           string           `json:"initFetchOn"`
	InitFetch             bool             `json:"initFetch"`
	Multiple              bool             `json:"multiple"`
	JoinValues            bool             `json:"joinValues"`
	Delimiter             string           `json:"delimiter"`
	ValuesNoWrap          bool             `json:"valuesNoWrap"`
	ExtractValue          bool             `json:"extractValue"`
	Clearable             bool             `json:"clearable"`
	ResetValue            string           `json:"resetValue"`
	DeferField            string           `json:"deferField"`
	DeferApi              draggableConfig  `json:"deferApi"`
	AddApi                draggableConfig  `json:"addApi"`
	AddControls           []any            `json:"addControls"`
	AddDialog             className        `json:"addDialog"`
	Creatable             bool             `json:"creatable"`
	CreateBtnLabel        string           `json:"createBtnLabel"`
	Editable              bool             `json:"editable"`
	EditApi               draggableConfig  `json:"editApi"`
	EditControls          []any            `json:"editControls"`
	EditDialog            className        `json:"editDialog"`
	Removable             bool             `json:"removable"`
	DeleteApi             draggableConfig  `json:"deleteApi"`
	DeleteConfirmText     string           `json:"deleteConfirmText"`
	Size                  string           `json:"size"`
	Label                 SchemaClassName  `json:"label"`
	LabelAlign            className        `json:"labelAlign"`
	LabelWidth            any              `json:"labelWidth"`
	LabelClassName        string           `json:"labelClassName"`
	Name                  string           `json:"name"`
	ExtraName             string           `json:"extraName"`
	Hint                  string           `json:"hint"`
	SubmitOnChange        bool             `json:"submitOnChange"`
	ReadOnly              bool             `json:"readOnly"`
	ReadOnlyOn            string           `json:"readOnlyOn"`
	ValidateOnChange      bool             `json:"validateOnChange"`
	Description           string           `json:"description"`
	Desc                  string           `json:"desc"`
	DescriptionClassName  className        `json:"descriptionClassName"`
	Mode                  string           `json:"mode"`
	Horizontal            className        `json:"horizontal"`
	Inline                bool             `json:"inline"`
	InputClassName        className        `json:"inputClassName"`
	Placeholder           string           `json:"placeholder"`
	Required              bool             `json:"required"`
	ValidationErrors      validationErrors `json:"validationErrors"`
	Validations           data             `json:"validations"`
	Value                 mock             `json:"value"`
	ClearValueOnHidden    bool             `json:"clearValueOnHidden"`
	ValidateApi           SchemaClassName  `json:"validateApi"`
	AutoFill              autoFill         `json:"autoFill"`
	InitAutoFill          data             `json:"initAutoFill"`
	Row                   float64          `json:"row"`
	ID                    string           `json:"$$id"`
	ClassName             className        `json:"className"`
	Ref                   string           `json:"$ref"`
	Disabled              bool             `json:"disabled"`
	DisabledOn            className        `json:"disabledOn"`
	Hidden                bool             `json:"hidden"`
	HiddenOn              className        `json:"hiddenOn"`
	Visible               bool             `json:"visible"`
	VisibleOn             className        `json:"visibleOn"`
	Id                    string           `json:"id"`
	OnEvent               any              `json:"onEvent"`
	Static                bool             `json:"static"`
	StaticOn              className        `json:"staticOn"`
	StaticPlaceholder     string           `json:"staticPlaceholder"`
	StaticClassName       className        `json:"staticClassName"`
	StaticLabelClassName  className        `json:"staticLabelClassName"`
	StaticInputClassName  className        `json:"staticInputClassName"`
	StaticSchema          staticSchema     `json:"staticSchema"`
	Style                 any              `json:"style"`
	EditorSetting         editorSetting    `json:"editorSetting"`
	UseMobileUI           bool             `json:"useMobileUI"`
	TestIdBuilder         RootSchema       `json:"testIdBuilder"`
	Remark                className        `json:"remark"`
	LabelRemark           className        `json:"labelRemark"`
	TypeVal               string           `json:"type"`
	AddOn                 addOn            `json:"addOn"`
	TrimContents          bool             `json:"trimContents"`
	AutoComplete          className        `json:"autoComplete"`
	NativeAutoComplete    string           `json:"nativeAutoComplete"`
	BorderMode            string           `json:"borderMode"`
	MinLength             float64          `json:"minLength"`
	MaxLength             float64          `json:"maxLength"`
	ShowCounter           bool             `json:"showCounter"`
	Prefix                string           `json:"prefix"`
	Suffix                string           `json:"suffix"`
	Transform             transform        `json:"transform"`
	InputControlClassName string           `json:"inputControlClassName"`
	NativeInputClassName  string           `json:"nativeInputClassName"`
	ClearValueOnEmpty     bool             `json:"clearValueOnEmpty"`
}

type TextareaControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	MaxRows              float64          `json:"maxRows"`
	MinRows              float64          `json:"minRows"`
	BorderMode           string           `json:"borderMode"`
	MaxLength            float64          `json:"maxLength"`
	ShowCounter          bool             `json:"showCounter"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
}

type TimeControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Clearable            bool             `json:"clearable"`
	Format               string           `json:"format"`
	ValueFormat          string           `json:"valueFormat"`
	InputFormat          string           `json:"inputFormat"`
	DisplayFormat        string           `json:"displayFormat"`
	Utc                  bool             `json:"utc"`
	Emebed               bool             `json:"emebed"`
	BorderMode           string           `json:"borderMode"`
	Shortcuts            SchemaClassName  `json:"shortcuts"`
	DisabledDate         string           `json:"disabledDate"`
	TimeFormat           string           `json:"timeFormat"`
	TimeConstraints      mock             `json:"timeConstraints"`
}

type TimelineItemSchema struct {
	Testid               string          `json:"testid"`
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	Time                 string          `json:"time"`
	Title                className       `json:"title"`
	Detail               string          `json:"detail"`
	DetailCollapsedText  string          `json:"detailCollapsedText"`
	DetailExpandedText   string          `json:"detailExpandedText"`
	Color                string          `json:"color"`
	Icon                 SchemaClassName `json:"icon"`
	IconClassName        string          `json:"iconClassName"`
	TimeClassName        string          `json:"timeClassName"`
	TitleClassName       string          `json:"titleClassName"`
	DetailClassName      string          `json:"detailClassName"`
}

type TimelineSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Items                []any           `json:"items"`
	Source               draggableConfig `json:"source"`
	Mode                 string          `json:"mode"`
	Direction            string          `json:"direction"`
	Reverse              bool            `json:"reverse"`
	ItemTitleSchema      className       `json:"itemTitleSchema"`
	IconClassName        string          `json:"iconClassName"`
	TimeClassName        string          `json:"timeClassName"`
	TitleClassName       string          `json:"titleClassName"`
	DetailClassName      string          `json:"detailClassName"`
}

type ToastActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Toast                className       `json:"toast"`
}

type ToastSchemaBase struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              RootSchema    `json:"type"`
	Testid               string        `json:"testid"`
	Items                []ArrayItem   `json:"items"`
	Position             string        `json:"position"`
	CloseButton          bool          `json:"closeButton"`
	ShowIcon             bool          `json:"showIcon"`
	Timeout              float64       `json:"timeout"`
}

type TooltipWrapperSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            string          `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Title                string          `json:"title"`
	Content              string          `json:"content"`
	Tooltip              string          `json:"tooltip"`
	Placement            string          `json:"placement"`
	Offset               []float64       `json:"offset"`
	ShowArrow            bool            `json:"showArrow"`
	Trigger              draggableConfig `json:"trigger"`
	MouseEnterDelay      float64         `json:"mouseEnterDelay"`
	MouseLeaveDelay      float64         `json:"mouseLeaveDelay"`
	RootClose            bool            `json:"rootClose"`
	Body                 className       `json:"body"`
	WrapperComponent     string          `json:"wrapperComponent"`
	Inline               bool            `json:"inline"`
	TooltipTheme         string          `json:"tooltipTheme"`
	Enterable            bool            `json:"enterable"`
	TooltipStyle         any             `json:"tooltipStyle"`
	TooltipClassName     string          `json:"tooltipClassName"`
}

type TplSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Tpl                  RootSchema    `json:"tpl"`
	Html                 RootSchema    `json:"html"`
	Text                 RootSchema    `json:"text"`
	Raw                  string        `json:"raw"`
	Inline               bool          `json:"inline"`
	WrapperComponent     mock          `json:"wrapperComponent"`
	Badge                className     `json:"badge"`
	TestidBuilder        RootSchema    `json:"testidBuilder"`
}

type TransferControlSchema struct {
	LoadingConfig              loadingConfig    `json:"loadingConfig"`
	Options                    options          `json:"options"`
	Source                     draggableConfig  `json:"source"`
	SelectFirst                bool             `json:"selectFirst"`
	InitFetchOn                string           `json:"initFetchOn"`
	InitFetch                  bool             `json:"initFetch"`
	Multiple                   bool             `json:"multiple"`
	JoinValues                 bool             `json:"joinValues"`
	Delimiter                  string           `json:"delimiter"`
	ValuesNoWrap               bool             `json:"valuesNoWrap"`
	ExtractValue               bool             `json:"extractValue"`
	Clearable                  bool             `json:"clearable"`
	ResetValue                 string           `json:"resetValue"`
	DeferField                 string           `json:"deferField"`
	DeferApi                   draggableConfig  `json:"deferApi"`
	AddApi                     draggableConfig  `json:"addApi"`
	AddControls                []any            `json:"addControls"`
	AddDialog                  className        `json:"addDialog"`
	Creatable                  bool             `json:"creatable"`
	CreateBtnLabel             string           `json:"createBtnLabel"`
	Editable                   bool             `json:"editable"`
	EditApi                    draggableConfig  `json:"editApi"`
	EditControls               []any            `json:"editControls"`
	EditDialog                 className        `json:"editDialog"`
	Removable                  bool             `json:"removable"`
	DeleteApi                  draggableConfig  `json:"deleteApi"`
	DeleteConfirmText          string           `json:"deleteConfirmText"`
	Size                       string           `json:"size"`
	Label                      SchemaClassName  `json:"label"`
	LabelAlign                 className        `json:"labelAlign"`
	LabelWidth                 any              `json:"labelWidth"`
	LabelClassName             string           `json:"labelClassName"`
	Name                       string           `json:"name"`
	ExtraName                  string           `json:"extraName"`
	Hint                       string           `json:"hint"`
	SubmitOnChange             bool             `json:"submitOnChange"`
	ReadOnly                   bool             `json:"readOnly"`
	ReadOnlyOn                 string           `json:"readOnlyOn"`
	ValidateOnChange           bool             `json:"validateOnChange"`
	Description                string           `json:"description"`
	Desc                       string           `json:"desc"`
	DescriptionClassName       className        `json:"descriptionClassName"`
	Mode                       string           `json:"mode"`
	Horizontal                 className        `json:"horizontal"`
	Inline                     bool             `json:"inline"`
	InputClassName             className        `json:"inputClassName"`
	Placeholder                string           `json:"placeholder"`
	Required                   bool             `json:"required"`
	ValidationErrors           validationErrors `json:"validationErrors"`
	Validations                data             `json:"validations"`
	Value                      mock             `json:"value"`
	ClearValueOnHidden         bool             `json:"clearValueOnHidden"`
	ValidateApi                SchemaClassName  `json:"validateApi"`
	AutoFill                   autoFill         `json:"autoFill"`
	InitAutoFill               data             `json:"initAutoFill"`
	Row                        float64          `json:"row"`
	ID                         string           `json:"$$id"`
	ClassName                  className        `json:"className"`
	Ref                        string           `json:"$ref"`
	Disabled                   bool             `json:"disabled"`
	DisabledOn                 className        `json:"disabledOn"`
	Hidden                     bool             `json:"hidden"`
	HiddenOn                   className        `json:"hiddenOn"`
	Visible                    bool             `json:"visible"`
	VisibleOn                  className        `json:"visibleOn"`
	Id                         string           `json:"id"`
	OnEvent                    any              `json:"onEvent"`
	Static                     bool             `json:"static"`
	StaticOn                   className        `json:"staticOn"`
	StaticPlaceholder          string           `json:"staticPlaceholder"`
	StaticClassName            className        `json:"staticClassName"`
	StaticLabelClassName       className        `json:"staticLabelClassName"`
	StaticInputClassName       className        `json:"staticInputClassName"`
	StaticSchema               staticSchema     `json:"staticSchema"`
	Style                      any              `json:"style"`
	EditorSetting              editorSetting    `json:"editorSetting"`
	UseMobileUI                bool             `json:"useMobileUI"`
	TestIdBuilder              RootSchema       `json:"testIdBuilder"`
	Remark                     className        `json:"remark"`
	LabelRemark                className        `json:"labelRemark"`
	TypeVal                    string           `json:"type"`
	ShowArrow                  bool             `json:"showArrow"`
	Sortable                   bool             `json:"sortable"`
	SelectMode                 string           `json:"selectMode"`
	ResultListModeFollowSelect bool             `json:"resultListModeFollowSelect"`
	LeftOptions                []any            `json:"leftOptions"`
	LeftMode                   string           `json:"leftMode"`
	RightMode                  string           `json:"rightMode"`
	SearchResultMode           string           `json:"searchResultMode"`
	Columns                    []any            `json:"columns"`
	SearchResultColumns        []any            `json:"searchResultColumns"`
	Searchable                 bool             `json:"searchable"`
	ResultSearchable           bool             `json:"resultSearchable"`
	SearchApi                  className        `json:"searchApi"`
	SelectTitle                string           `json:"selectTitle"`
	ResultTitle                string           `json:"resultTitle"`
	MenuTpl                    className        `json:"menuTpl"`
	ValueTpl                   className        `json:"valueTpl"`
	SearchPlaceholder          string           `json:"searchPlaceholder"`
	ResultSearchPlaceholder    string           `json:"resultSearchPlaceholder"`
	Statistics                 bool             `json:"statistics"`
	ItemHeight                 float64          `json:"itemHeight"`
	VirtualThreshold           float64          `json:"virtualThreshold"`
	ShowInvalidMatch           bool             `json:"showInvalidMatch"`
	OnlyChildren               bool             `json:"onlyChildren"`
	InitiallyOpen              bool             `json:"initiallyOpen"`
	AutoCheckChildren          bool             `json:"autoCheckChildren"`
	Pagination                 pagination       `json:"pagination"`
}

type TransferPickerControlSchema struct {
	LoadingConfig              loadingConfig    `json:"loadingConfig"`
	ShowArrow                  bool             `json:"showArrow"`
	Sortable                   bool             `json:"sortable"`
	SelectMode                 string           `json:"selectMode"`
	ResultListModeFollowSelect bool             `json:"resultListModeFollowSelect"`
	LeftOptions                []any            `json:"leftOptions"`
	LeftMode                   string           `json:"leftMode"`
	RightMode                  string           `json:"rightMode"`
	SearchResultMode           string           `json:"searchResultMode"`
	Columns                    []any            `json:"columns"`
	SearchResultColumns        []any            `json:"searchResultColumns"`
	Searchable                 bool             `json:"searchable"`
	ResultSearchable           bool             `json:"resultSearchable"`
	SearchApi                  className        `json:"searchApi"`
	SelectTitle                string           `json:"selectTitle"`
	ResultTitle                string           `json:"resultTitle"`
	MenuTpl                    className        `json:"menuTpl"`
	ValueTpl                   className        `json:"valueTpl"`
	SearchPlaceholder          string           `json:"searchPlaceholder"`
	ResultSearchPlaceholder    string           `json:"resultSearchPlaceholder"`
	Statistics                 bool             `json:"statistics"`
	ItemHeight                 float64          `json:"itemHeight"`
	VirtualThreshold           float64          `json:"virtualThreshold"`
	ShowInvalidMatch           bool             `json:"showInvalidMatch"`
	OnlyChildren               bool             `json:"onlyChildren"`
	InitiallyOpen              bool             `json:"initiallyOpen"`
	AutoCheckChildren          bool             `json:"autoCheckChildren"`
	Pagination                 pagination       `json:"pagination"`
	Remark                     className        `json:"remark"`
	LabelRemark                className        `json:"labelRemark"`
	Size                       string           `json:"size"`
	Label                      SchemaClassName  `json:"label"`
	LabelAlign                 className        `json:"labelAlign"`
	LabelWidth                 any              `json:"labelWidth"`
	LabelClassName             string           `json:"labelClassName"`
	Name                       string           `json:"name"`
	ExtraName                  string           `json:"extraName"`
	Hint                       string           `json:"hint"`
	SubmitOnChange             bool             `json:"submitOnChange"`
	ReadOnly                   bool             `json:"readOnly"`
	ReadOnlyOn                 string           `json:"readOnlyOn"`
	ValidateOnChange           bool             `json:"validateOnChange"`
	Description                string           `json:"description"`
	Desc                       string           `json:"desc"`
	DescriptionClassName       className        `json:"descriptionClassName"`
	Mode                       string           `json:"mode"`
	Horizontal                 className        `json:"horizontal"`
	Inline                     bool             `json:"inline"`
	InputClassName             className        `json:"inputClassName"`
	Placeholder                string           `json:"placeholder"`
	Required                   bool             `json:"required"`
	ValidationErrors           validationErrors `json:"validationErrors"`
	Validations                data             `json:"validations"`
	Value                      mock             `json:"value"`
	ClearValueOnHidden         bool             `json:"clearValueOnHidden"`
	ValidateApi                SchemaClassName  `json:"validateApi"`
	AutoFill                   autoFill         `json:"autoFill"`
	InitAutoFill               data             `json:"initAutoFill"`
	Row                        float64          `json:"row"`
	ID                         string           `json:"$$id"`
	ClassName                  className        `json:"className"`
	Ref                        string           `json:"$ref"`
	Disabled                   bool             `json:"disabled"`
	DisabledOn                 className        `json:"disabledOn"`
	Hidden                     bool             `json:"hidden"`
	HiddenOn                   className        `json:"hiddenOn"`
	Visible                    bool             `json:"visible"`
	VisibleOn                  className        `json:"visibleOn"`
	Id                         string           `json:"id"`
	OnEvent                    any              `json:"onEvent"`
	Static                     bool             `json:"static"`
	StaticOn                   className        `json:"staticOn"`
	StaticPlaceholder          string           `json:"staticPlaceholder"`
	StaticClassName            className        `json:"staticClassName"`
	StaticLabelClassName       className        `json:"staticLabelClassName"`
	StaticInputClassName       className        `json:"staticInputClassName"`
	StaticSchema               staticSchema     `json:"staticSchema"`
	Style                      any              `json:"style"`
	EditorSetting              editorSetting    `json:"editorSetting"`
	UseMobileUI                bool             `json:"useMobileUI"`
	TestIdBuilder              RootSchema       `json:"testIdBuilder"`
	Options                    options          `json:"options"`
	Source                     draggableConfig  `json:"source"`
	SelectFirst                bool             `json:"selectFirst"`
	InitFetchOn                string           `json:"initFetchOn"`
	InitFetch                  bool             `json:"initFetch"`
	Multiple                   bool             `json:"multiple"`
	JoinValues                 bool             `json:"joinValues"`
	Delimiter                  string           `json:"delimiter"`
	ValuesNoWrap               bool             `json:"valuesNoWrap"`
	ExtractValue               bool             `json:"extractValue"`
	Clearable                  bool             `json:"clearable"`
	ResetValue                 string           `json:"resetValue"`
	DeferField                 string           `json:"deferField"`
	DeferApi                   draggableConfig  `json:"deferApi"`
	AddApi                     draggableConfig  `json:"addApi"`
	AddControls                []any            `json:"addControls"`
	AddDialog                  className        `json:"addDialog"`
	Creatable                  bool             `json:"creatable"`
	CreateBtnLabel             string           `json:"createBtnLabel"`
	Editable                   bool             `json:"editable"`
	EditApi                    draggableConfig  `json:"editApi"`
	EditControls               []any            `json:"editControls"`
	EditDialog                 className        `json:"editDialog"`
	Removable                  bool             `json:"removable"`
	DeleteApi                  draggableConfig  `json:"deleteApi"`
	DeleteConfirmText          string           `json:"deleteConfirmText"`
	TypeVal                    string           `json:"type"`
	BorderMode                 string           `json:"borderMode"`
	PickerSize                 string           `json:"pickerSize"`
}

type TreeControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             className        `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	HideRoot             bool             `json:"hideRoot"`
	RootLabel            string           `json:"rootLabel"`
	RootValue            mock             `json:"rootValue"`
	ShowIcon             bool             `json:"showIcon"`
	AutoCheckChildren    bool             `json:"autoCheckChildren"`
	Cascade              bool             `json:"cascade"`
	WithChildren         bool             `json:"withChildren"`
	OnlyChildren         bool             `json:"onlyChildren"`
	OnlyLeaf             bool             `json:"onlyLeaf"`
	RootCreatable        bool             `json:"rootCreatable"`
	EnableNodePath       bool             `json:"enableNodePath"`
	PathSeparator        string           `json:"pathSeparator"`
	ShowOutline          bool             `json:"showOutline"`
	HighlightTxt         string           `json:"highlightTxt"`
	EnableDefaultIcon    bool             `json:"enableDefaultIcon"`
	Searchable           bool             `json:"searchable"`
	SearchApi            className        `json:"searchApi"`
	SearchConfig         searchConfig1    `json:"searchConfig"`
	HeightAuto           bool             `json:"heightAuto"`
}

type TreeSelectControlSchema struct {
	Options              options          `json:"options"`
	Source               draggableConfig  `json:"source"`
	SelectFirst          bool             `json:"selectFirst"`
	InitFetchOn          string           `json:"initFetchOn"`
	InitFetch            bool             `json:"initFetch"`
	Multiple             bool             `json:"multiple"`
	JoinValues           bool             `json:"joinValues"`
	Delimiter            string           `json:"delimiter"`
	ValuesNoWrap         bool             `json:"valuesNoWrap"`
	ExtractValue         bool             `json:"extractValue"`
	Clearable            bool             `json:"clearable"`
	ResetValue           string           `json:"resetValue"`
	DeferField           string           `json:"deferField"`
	DeferApi             className        `json:"deferApi"`
	AddApi               draggableConfig  `json:"addApi"`
	AddControls          []any            `json:"addControls"`
	AddDialog            className        `json:"addDialog"`
	Creatable            bool             `json:"creatable"`
	CreateBtnLabel       string           `json:"createBtnLabel"`
	Editable             bool             `json:"editable"`
	EditApi              draggableConfig  `json:"editApi"`
	EditControls         []any            `json:"editControls"`
	EditDialog           className        `json:"editDialog"`
	Removable            bool             `json:"removable"`
	DeleteApi            draggableConfig  `json:"deleteApi"`
	DeleteConfirmText    string           `json:"deleteConfirmText"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	TypeVal              string           `json:"type"`
	HideRoot             bool             `json:"hideRoot"`
	RootLabel            string           `json:"rootLabel"`
	RootValue            mock             `json:"rootValue"`
	ShowIcon             bool             `json:"showIcon"`
	Cascade              bool             `json:"cascade"`
	WithChildren         bool             `json:"withChildren"`
	OnlyChildren         bool             `json:"onlyChildren"`
	OnlyLeaf             bool             `json:"onlyLeaf"`
	RootCreatable        bool             `json:"rootCreatable"`
	HideNodePathLabel    bool             `json:"hideNodePathLabel"`
	EnableNodePath       bool             `json:"enableNodePath"`
	PathSeparator        string           `json:"pathSeparator"`
	ShowOutline          bool             `json:"showOutline"`
	MaxTagCount          float64          `json:"maxTagCount"`
	OverflowTagPopover   className        `json:"overflowTagPopover"`
	MenuTpl              string           `json:"menuTpl"`
	EnableDefaultIcon    bool             `json:"enableDefaultIcon"`
}

type UUIDControlSchema struct {
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Placeholder          string           `json:"placeholder"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	TypeVal              string           `json:"type"`
	Length               float64          `json:"length"`
}

type UrlActionSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	Blank                bool            `json:"blank"`
	Url                  string          `json:"url"`
}

type VBoxSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Rows                 []any         `json:"rows"`
}

type VanillaAction struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Block                bool            `json:"block"`
	DisabledTip          string          `json:"disabledTip"`
	Icon                 className       `json:"icon"`
	IconClassName        className       `json:"iconClassName"`
	RightIcon            className       `json:"rightIcon"`
	RightIconClassName   className       `json:"rightIconClassName"`
	LoadingClassName     className       `json:"loadingClassName"`
	Label                string          `json:"label"`
	Level                string          `json:"level"`
	Primary              bool            `json:"primary"`
	Size                 string          `json:"size"`
	Tooltip              RootSchema      `json:"tooltip"`
	TooltipPlacement     string          `json:"tooltipPlacement"`
	ConfirmText          string          `json:"confirmText"`
	Required             []string        `json:"required"`
	ActiveLevel          string          `json:"activeLevel"`
	ActiveClassName      string          `json:"activeClassName"`
	Close                any             `json:"close"`
	RequireSelected      bool            `json:"requireSelected"`
	MergeData            bool            `json:"mergeData"`
	Target               string          `json:"target"`
	CountDown            float64         `json:"countDown"`
	CountDownTpl         string          `json:"countDownTpl"`
	Badge                className       `json:"badge"`
	HotKey               string          `json:"hotKey"`
	LoadingOn            string          `json:"loadingOn"`
	OnClick              SchemaClassName `json:"onClick"`
	Body                 className       `json:"body"`
	ActionType           string          `json:"actionType"`
	DownloadFileName     string          `json:"downloadFileName"`
}

type VariableItem struct {
	Label      string `json:"label"`
	Value      string `json:"value"`
	Path       string `json:"path"`
	Children   []any  `json:"children"`
	TypeVal    string `json:"type"`
	Tag        string `json:"tag"`
	SelectMode string `json:"selectMode"`
	IsMember   bool   `json:"isMember"`
}

type VideoSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	AutoPlay             bool          `json:"autoPlay"`
	ColumnsCount         float64       `json:"columnsCount"`
	Frames               any           `json:"frames"`
	FramesClassName      className     `json:"framesClassName"`
	IsLive               bool          `json:"isLive"`
	JumpFrame            bool          `json:"jumpFrame"`
	Muted                bool          `json:"muted"`
	Loop                 bool          `json:"loop"`
	PlayerClassName      className     `json:"playerClassName"`
	Poster               className     `json:"poster"`
	SplitPoster          bool          `json:"splitPoster"`
	Src                  className     `json:"src"`
	VideoType            string        `json:"videoType"`
	AspectRatio          string        `json:"aspectRatio"`
	Rates                []float64     `json:"rates"`
	JumpBufferDuration   float64       `json:"jumpBufferDuration"`
	StopOnNextFrame      bool          `json:"stopOnNextFrame"`
}

type WizardSchema struct {
	LoadingConfig        loadingConfig   `json:"loadingConfig"`
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	ActionClassName      className       `json:"actionClassName"`
	ActionFinishLabel    string          `json:"actionFinishLabel"`
	ActionNextLabel      string          `json:"actionNextLabel"`
	ActionNextSaveLabel  string          `json:"actionNextSaveLabel"`
	ActionPrevLabel      string          `json:"actionPrevLabel"`
	Api                  className       `json:"api"`
	BulkSubmit           bool            `json:"bulkSubmit"`
	InitApi              className       `json:"initApi"`
	Mode                 string          `json:"mode"`
	Name                 RootSchema      `json:"name"`
	ReadOnly             bool            `json:"readOnly"`
	Redirect             string          `json:"redirect"`
	Reload               RootSchema      `json:"reload"`
	Target               string          `json:"target"`
	AffixFooter          SchemaClassName `json:"affixFooter"`
	Steps                []any           `json:"steps"`
	StartStep            string          `json:"startStep"`
	StepsClassName       string          `json:"stepsClassName"`
	BodyClassName        string          `json:"bodyClassName"`
	StepClassName        string          `json:"stepClassName"`
	FooterClassName      string          `json:"footerClassName"`
	WrapWithPanel        bool            `json:"wrapWithPanel"`
}

type WizardStepSchema struct {
	Api                         SchemaClassName `json:"api"`
	AsyncApi                    SchemaClassName `json:"asyncApi"`
	InitApi                     SchemaClassName `json:"initApi"`
	Jumpable                    bool            `json:"jumpable"`
	JumpableOn                  className       `json:"jumpableOn"`
	Title                       string          `json:"title"`
	Label                       string          `json:"label"`
	Actions                     []any           `json:"actions"`
	Redirect                    string          `json:"redirect"`
	Reload                      string          `json:"reload"`
	Target                      string          `json:"target"`
	Testid                      string          `json:"testid"`
	ID                          string          `json:"$$id"`
	ClassName                   className       `json:"className"`
	Ref                         string          `json:"$ref"`
	Disabled                    bool            `json:"disabled"`
	DisabledOn                  className       `json:"disabledOn"`
	Hidden                      bool            `json:"hidden"`
	HiddenOn                    className       `json:"hiddenOn"`
	Visible                     bool            `json:"visible"`
	VisibleOn                   className       `json:"visibleOn"`
	Id                          string          `json:"id"`
	OnEvent                     any             `json:"onEvent"`
	Static                      bool            `json:"static"`
	StaticOn                    RootSchema      `json:"staticOn"`
	StaticPlaceholder           string          `json:"staticPlaceholder"`
	StaticClassName             RootSchema      `json:"staticClassName"`
	StaticLabelClassName        className       `json:"staticLabelClassName"`
	StaticInputClassName        className       `json:"staticInputClassName"`
	StaticSchema                staticSchema    `json:"staticSchema"`
	Style                       any             `json:"style"`
	EditorSetting               editorSetting   `json:"editorSetting"`
	UseMobileUI                 bool            `json:"useMobileUI"`
	TestIdBuilder               RootSchema      `json:"testIdBuilder"`
	SubTitle                    SchemaClassName `json:"subTitle"`
	Icon                        string          `json:"icon"`
	Value                       any             `json:"value"`
	Description                 SchemaClassName `json:"description"`
	Body                        className       `json:"body"`
	Tabs                        staticSchema    `json:"tabs"`
	FieldSet                    staticSchema    `json:"fieldSet"`
	Data                        staticSchema    `json:"data"`
	Debug                       bool            `json:"debug"`
	DebugConfig                 debugConfig     `json:"debugConfig"`
	InitAsyncApi                SchemaClassName `json:"initAsyncApi"`
	InitFinishedField           string          `json:"initFinishedField"`
	InitCheckInterval           float64         `json:"initCheckInterval"`
	InitFetch                   bool            `json:"initFetch"`
	InitFetchOn                 string          `json:"initFetchOn"`
	Interval                    float64         `json:"interval"`
	SilentPolling               bool            `json:"silentPolling"`
	StopAutoRefreshWhen         string          `json:"stopAutoRefreshWhen"`
	PersistData                 string          `json:"persistData"`
	PersistDataKeys             []string        `json:"persistDataKeys"`
	ClearPersistDataAfterSubmit bool            `json:"clearPersistDataAfterSubmit"`
	Feedback                    mock            `json:"feedback"`
	CheckInterval               float64         `json:"checkInterval"`
	FinishedField               string          `json:"finishedField"`
	ResetAfterSubmit            bool            `json:"resetAfterSubmit"`
	ClearAfterSubmit            bool            `json:"clearAfterSubmit"`
	Mode                        string          `json:"mode"`
	ColumnCount                 float64         `json:"columnCount"`
	Horizontal                  className       `json:"horizontal"`
	AutoFocus                   bool            `json:"autoFocus"`
	Messages                    messages        `json:"messages"`
	Name                        string          `json:"name"`
	PanelClassName              className       `json:"panelClassName"`
	PrimaryField                string          `json:"primaryField"`
	SubmitOnChange              bool            `json:"submitOnChange"`
	SubmitOnInit                bool            `json:"submitOnInit"`
	SubmitText                  string          `json:"submitText"`
	WrapWithPanel               bool            `json:"wrapWithPanel"`
	AffixFooter                 bool            `json:"affixFooter"`
	PromptPageLeave             bool            `json:"promptPageLeave"`
	PromptPageLeaveMessage      string          `json:"promptPageLeaveMessage"`
	Rules                       []ArrayItem3    `json:"rules"`
	PreventEnterSubmit          bool            `json:"preventEnterSubmit"`
	LabelAlign                  className       `json:"labelAlign"`
	LabelWidth                  any             `json:"labelWidth"`
}

type WordsSchema struct {
	ID                   string          `json:"$$id"`
	ClassName            className       `json:"className"`
	Ref                  string          `json:"$ref"`
	Disabled             bool            `json:"disabled"`
	DisabledOn           className       `json:"disabledOn"`
	Hidden               bool            `json:"hidden"`
	HiddenOn             className       `json:"hiddenOn"`
	Visible              bool            `json:"visible"`
	VisibleOn            className       `json:"visibleOn"`
	Id                   string          `json:"id"`
	OnEvent              any             `json:"onEvent"`
	Static               bool            `json:"static"`
	StaticOn             className       `json:"staticOn"`
	StaticPlaceholder    string          `json:"staticPlaceholder"`
	StaticClassName      className       `json:"staticClassName"`
	StaticLabelClassName className       `json:"staticLabelClassName"`
	StaticInputClassName className       `json:"staticInputClassName"`
	StaticSchema         staticSchema    `json:"staticSchema"`
	Style                any             `json:"style"`
	EditorSetting        editorSetting   `json:"editorSetting"`
	UseMobileUI          bool            `json:"useMobileUI"`
	TestIdBuilder        RootSchema      `json:"testIdBuilder"`
	TypeVal              string          `json:"type"`
	Testid               string          `json:"testid"`
	Limit                float64         `json:"limit"`
	ExpendButtonText     string          `json:"expendButtonText"`
	ExpendButton         className       `json:"expendButton"`
	CollapseButtonText   string          `json:"collapseButtonText"`
	CollapseButton       className       `json:"collapseButton"`
	Words                SchemaClassName `json:"words"`
	InTag                SchemaClassName `json:"inTag"`
	Delimiter            string          `json:"delimiter"`
}

type WrapperSchema struct {
	ID                   string        `json:"$$id"`
	ClassName            className     `json:"className"`
	Ref                  string        `json:"$ref"`
	Disabled             bool          `json:"disabled"`
	DisabledOn           className     `json:"disabledOn"`
	Hidden               bool          `json:"hidden"`
	HiddenOn             className     `json:"hiddenOn"`
	Visible              bool          `json:"visible"`
	VisibleOn            className     `json:"visibleOn"`
	Id                   string        `json:"id"`
	OnEvent              any           `json:"onEvent"`
	Static               bool          `json:"static"`
	StaticOn             className     `json:"staticOn"`
	StaticPlaceholder    string        `json:"staticPlaceholder"`
	StaticClassName      className     `json:"staticClassName"`
	StaticLabelClassName className     `json:"staticLabelClassName"`
	StaticInputClassName className     `json:"staticInputClassName"`
	StaticSchema         staticSchema  `json:"staticSchema"`
	Style                any           `json:"style"`
	EditorSetting        editorSetting `json:"editorSetting"`
	UseMobileUI          bool          `json:"useMobileUI"`
	TestIdBuilder        RootSchema    `json:"testIdBuilder"`
	TypeVal              string        `json:"type"`
	Testid               string        `json:"testid"`
	Body                 className     `json:"body"`
	Size                 string        `json:"size"`
	Wrap                 bool          `json:"wrap"`
}

type activeKey struct {
	AnyOf       []ArrayElem11 `json:"anyOf"`
	Description string        `json:"description"`
}

type addOn struct {
	AnyOf []ArrayElem28 `json:"anyOf"`
}

type alcohol struct {
	Bar_service staticSchema `json:"bar_service"`
	Beer        staticSchema `json:"beer"`
	Byo         staticSchema `json:"byo"`
	Cocktails   staticSchema `json:"cocktails"`
	Full_bar    staticSchema `json:"full_bar"`
	Wine        staticSchema `json:"wine"`
}

type amenities struct {
	Restroom              staticSchema `json:"restroom"`
	Smoking               staticSchema `json:"smoking"`
	Jukebox               staticSchema `json:"jukebox"`
	Music                 staticSchema `json:"music"`
	Live_music            staticSchema `json:"live_music"`
	Private_room          staticSchema `json:"private_room"`
	Outdoor_seating       staticSchema `json:"outdoor_seating"`
	Tvs                   staticSchema `json:"tvs"`
	Atm                   staticSchema `json:"atm"`
	Coat_check            staticSchema `json:"coat_check"`
	Wheelchair_accessible staticSchema `json:"wheelchair_accessible"`
	Parking               parking      `json:"parking"`
	Sit_down_dining       staticSchema `json:"sit_down_dining"`
	Wifi                  string       `json:"wifi"`
}

type attributes struct {
	Business_meeting string `json:"business_meeting"`
	Clean            string `json:"clean"`
	Crowded          string `json:"crowded"`
	Dates_popular    string `json:"dates_popular"`
	Dressy           string `json:"dressy"`
	Families_popular string `json:"families_popular"`
	Gluten_free_diet string `json:"gluten_free_diet"`
	Good_for_dogs    string `json:"good_for_dogs"`
	Groups_popular   string `json:"groups_popular"`
	Healthy_diet     string `json:"healthy_diet"`
	Late_night       string `json:"late_night"`
	Noisy            string `json:"noisy"`
	Quick_bite       string `json:"quick_bite"`
	Romantic         string `json:"romantic"`
	Service_quality  string `json:"service_quality"`
	Singles_popular  string `json:"singles_popular"`
	Special_occasion string `json:"special_occasion"`
	Trendy           string `json:"trendy"`
	Value_for_money  string `json:"value_for_money"`
	Vegan_diet       string `json:"vegan_diet"`
	Vegetarian_diet  string `json:"vegetarian_diet"`
}

type autoFill struct {
	AnyOf       []ArrayElem21 `json:"anyOf"`
	Description string        `json:"description"`
}

type balances struct {
	Available                int    `json:"available"`
	Current                  int    `json:"current"`
	Limit                    any    `json:"limit"`
	Margin_loan_amount       any    `json:"margin_loan_amount"`
	Iso_currency_code        string `json:"iso_currency_code"`
	Unofficial_currency_code any    `json:"unofficial_currency_code"`
}

type bearing struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}

type bounds struct {
	AnyOf       []ArrayElem12 `json:"anyOf"`
	Description string        `json:"description"`
}

type card struct {
	AnyOf []ArrayElem10 `json:"anyOf"`
}

type className struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}

type columnsTogglable struct {
	AnyOf       []ArrayElem3 `json:"anyOf"`
	Description string       `json:"description"`
}

type compressOptions struct {
	MaxHeight float64 `json:"maxHeight"`
	MaxWidth  float64 `json:"maxWidth"`
}

type credit_cards struct {
	Accepts_credit_cards staticSchema `json:"accepts_credit_cards"`
	Amex                 staticSchema `json:"amex"`
	Discover             staticSchema `json:"discover"`
	Visa                 staticSchema `json:"visa"`
	Diners_club          staticSchema `json:"diners_club"`
	Master_card          staticSchema `json:"master_card"`
	Union_pay            staticSchema `json:"union_pay"`
}

type data struct {
	AnyOf []ArrayElem2 `json:"anyOf"`
}

type data1 struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	Region      string `json:"region"`
	Street      string `json:"street"`
	Postal_code string `json:"postal_code"`
}

type debounceConfig struct {
	MaxWait  float64 `json:"maxWait"`
	Wait     float64 `json:"wait"`
	Leading  bool    `json:"leading"`
	Trailing bool    `json:"trailing"`
}

type debugConfig struct {
	LevelExpand       float64         `json:"levelExpand"`
	EnableClipboard   bool            `json:"enableClipboard"`
	IconStyle         string          `json:"iconStyle"`
	QuotesOnKeys      bool            `json:"quotesOnKeys"`
	SortKeys          bool            `json:"sortKeys"`
	EllipsisThreshold SchemaClassName `json:"ellipsisThreshold"`
}

type defaultPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type definitions struct {
	RootSchema                         RootSchema                      `json:"RootSchema"`
	PageSchema                         PageSchema                      `json:"PageSchema"`
	SchemaClassName                    SchemaClassName                 `json:"SchemaClassName"`
	SchemaExpression                   string                          `json:"SchemaExpression"`
	ListenerAction                     ListenerAction                  `json:"ListenerAction"`
	ConditionGroupValue                ConditionGroupValue             `json:"ConditionGroupValue"`
	ConditionRule                      ConditionRule                   `json:"ConditionRule"`
	ExpressionComplex                  right                           `json:"ExpressionComplex"`
	ExpressionValue                    right                           `json:"ExpressionValue"`
	ExpressionSimple                   data                            `json:"ExpressionSimple"`
	ExpressionFunc                     ExpressionFunc                  `json:"ExpressionFunc"`
	ExpressionField                    ExpressionField                 `json:"ExpressionField"`
	ExpressionFormula                  ExpressionFormula               `json:"ExpressionFormula"`
	OperatorType                       OperatorType                    `json:"OperatorType"`
	DebounceConfig                     debounceConfig                  `json:"debounceConfig"`
	TrackConfig                        trackConfig                     `json:"trackConfig"`
	TestIdBuilder                      TestIdBuilder                   `json:"TestIdBuilder"`
	SchemaRemark                       data                            `json:"SchemaRemark"`
	SchemaIcon                         string                          `json:"SchemaIcon"`
	SchemaTpl                          string                          `json:"SchemaTpl"`
	SchemaCollection                   right                           `json:"SchemaCollection"`
	SchemaObject                       SchemaObject                    `json:"SchemaObject"`
	FlexSchema                         FlexSchema                      `json:"FlexSchema"`
	TplSchema                          TplSchema                       `json:"TplSchema"`
	BadgeObject                        BadgeObject                     `json:"BadgeObject"`
	RemarkSchema                       RemarkSchema                    `json:"RemarkSchema"`
	ActionSchema                       ActionSchema                    `json:"ActionSchema"`
	AjaxActionSchema                   AjaxActionSchema                `json:"AjaxActionSchema"`
	SchemaTooltip                      data                            `json:"SchemaTooltip"`
	SchemaApi                          data                            `json:"SchemaApi"`
	SchemaApiObject                    SchemaApiObject                 `json:"SchemaApiObject"`
	SchemaUrlPath                      RootSchema                      `json:"SchemaUrlPath"`
	SchemaTokenizeableString           string                          `json:"SchemaTokenizeableString"`
	FeedbackDialog                     FeedbackDialog                  `json:"FeedbackDialog"`
	SchemaName                         string                          `json:"SchemaName"`
	SchemaReload                       string                          `json:"SchemaReload"`
	UrlActionSchema                    UrlActionSchema                 `json:"UrlActionSchema"`
	LinkActionSchema                   LinkActionSchema                `json:"LinkActionSchema"`
	DialogActionSchema                 DialogActionSchema              `json:"DialogActionSchema"`
	DialogSchemaBase                   DialogSchemaBase                `json:"DialogSchemaBase"`
	DrawerActionSchema                 DrawerActionSchema              `json:"DrawerActionSchema"`
	DrawerSchemaBase                   DrawerSchemaBase                `json:"DrawerSchemaBase"`
	ToastActionSchema                  ToastActionSchema               `json:"ToastActionSchema"`
	ToastSchemaBase                    ToastSchemaBase                 `json:"ToastSchemaBase"`
	SchemaType                         string                          `json:"SchemaType"`
	CopyActionSchema                   CopyActionSchema                `json:"CopyActionSchema"`
	ReloadActionSchema                 ReloadActionSchema              `json:"ReloadActionSchema"`
	EmailActionSchema                  EmailActionSchema               `json:"EmailActionSchema"`
	OtherActionSchema                  OtherActionSchema               `json:"OtherActionSchema"`
	VanillaAction                      VanillaAction                   `json:"VanillaAction"`
	AlertSchema                        AlertSchema                     `json:"AlertSchema"`
	AudioSchema                        AudioSchema                     `json:"AudioSchema"`
	AvatarSchema                       AvatarSchema                    `json:"AvatarSchema"`
	ButtonGroupSchema                  ButtonGroupSchema               `json:"ButtonGroupSchema"`
	ButtonToolbarSchema                ButtonToolbarSchema             `json:"ButtonToolbarSchema"`
	CalendarSchema                     CalendarSchema                  `json:"CalendarSchema"`
	CardSchema                         CardSchema                      `json:"CardSchema"`
	CardBodyField                      CardBodyField                   `json:"CardBodyField"`
	SchemaPopOver                      data                            `json:"SchemaPopOver"`
	SchemaPopOverObject                SchemaPopOverObject             `json:"SchemaPopOverObject"`
	SchemaQuickEdit                    data                            `json:"SchemaQuickEdit"`
	SchemaQuickEditObject              SchemaQuickEditObject           `json:"SchemaQuickEditObject"`
	SchemaCopyable                     data                            `json:"SchemaCopyable"`
	SchemaCopyableObject               SchemaCopyableObject            `json:"SchemaCopyableObject"`
	CardsSchema                        CardsSchema                     `json:"CardsSchema"`
	Card2Schema                        Card2Schema                     `json:"Card2Schema"`
	CarouselSchema                     CarouselSchema                  `json:"CarouselSchema"`
	ChartSchema                        ChartSchema                     `json:"ChartSchema"`
	SchemaFunction                     data                            `json:"SchemaFunction"`
	Field_161becb8                     Field_161becb8                  `json:"interface-1922134811-9821-11285-1922134811-0-212510"`
	CollapseSchema                     CollapseSchema                  `json:"CollapseSchema"`
	CollapseGroupSchema                CollapseGroupSchema             `json:"CollapseGroupSchema"`
	ColorSchema                        ColorSchema                     `json:"ColorSchema"`
	ContainerSchema                    ContainerSchema                 `json:"ContainerSchema"`
	ContainerDraggableConfig           ContainerDraggableConfig        `json:"ContainerDraggableConfig"`
	SwitchContainerSchema              SwitchContainerSchema           `json:"SwitchContainerSchema"`
	StateSchema                        StateSchema                     `json:"StateSchema"`
	CRUDSchema                         draggableConfig                 `json:"CRUDSchema"`
	CRUDCardsSchema                    CRUDCardsSchema                 `json:"CRUDCardsSchema"`
	PlainObject                        any                             `json:"PlainObject"`
	CRUDBultinToolbarType              string                          `json:"CRUDBultinToolbarType"`
	SchemaMessage                      SchemaMessage                   `json:"SchemaMessage"`
	AutoGenerateFilterObject           AutoGenerateFilterObject        `json:"AutoGenerateFilterObject"`
	CRUDListSchema                     CRUDListSchema                  `json:"CRUDListSchema"`
	ListItemSchema                     ListItemSchema                  `json:"ListItemSchema"`
	ListBodyField                      ListBodyField                   `json:"ListBodyField"`
	ListBodyFieldObject                ListBodyFieldObject             `json:"ListBodyFieldObject"`
	CRUDTableSchema                    CRUDTableSchema                 `json:"CRUDTableSchema"`
	TableColumn                        right                           `json:"TableColumn"`
	TableColumnWithType                TableColumnWithType             `json:"TableColumnWithType"`
	TableColumnObject                  TableColumnObject               `json:"TableColumnObject"`
	CRUD2Schema                        right                           `json:"CRUD2Schema"`
	CRUD2CardsSchema                   CRUD2CardsSchema                `json:"CRUD2CardsSchema"`
	CRUD2ListSchema                    CRUD2ListSchema                 `json:"CRUD2ListSchema"`
	CRUD2TableSchema                   CRUD2TableSchema                `json:"CRUD2TableSchema"`
	ColumnSchema                       ColumnSchema                    `json:"ColumnSchema"`
	RowSelectionSchema                 RowSelectionSchema              `json:"RowSelectionSchema"`
	RowSelectionOptionsSchema          RowSelectionOptionsSchema       `json:"RowSelectionOptionsSchema"`
	ExpandableSchema                   ExpandableSchema                `json:"ExpandableSchema"`
	DateSchema                         DateSchema                      `json:"DateSchema"`
	DialogSchema                       DialogSchema                    `json:"DialogSchema"`
	DividerSchema                      DividerSchema                   `json:"DividerSchema"`
	DrawerSchema                       DrawerSchema                    `json:"DrawerSchema"`
	DropdownButtonSchema               DropdownButtonSchema            `json:"DropdownButtonSchema"`
	DropdownButton                     DropdownButton                  `json:"DropdownButton"`
	EachSchema                         EachSchema                      `json:"EachSchema"`
	GridSchema                         GridSchema                      `json:"GridSchema"`
	GridColumn                         RootSchema                      `json:"GridColumn"`
	GridColumnObject                   GridColumnObject                `json:"GridColumnObject"`
	FormHorizontal                     FormHorizontal                  `json:"FormHorizontal"`
	Grid2DSchema                       Grid2DSchema                    `json:"Grid2DSchema"`
	Grid                               Grid                            `json:"Grid"`
	HBoxSchema                         HBoxSchema                      `json:"HBoxSchema"`
	HBoxColumn                         RootSchema                      `json:"HBoxColumn"`
	HBoxColumnObject                   HBoxColumnObject                `json:"HBoxColumnObject"`
	IconSchema                         IconSchema                      `json:"IconSchema"`
	IconCheckedSchema                  IconCheckedSchema               `json:"IconCheckedSchema"`
	IFrameSchema                       IFrameSchema                    `json:"IFrameSchema"`
	ImageSchema                        ImageSchema                     `json:"ImageSchema"`
	ImageToolbarAction                 ImageToolbarAction              `json:"ImageToolbarAction"`
	ImagesSchema                       ImagesSchema                    `json:"ImagesSchema"`
	JsonSchema                         JsonSchema                      `json:"JsonSchema"`
	LinkSchema                         LinkSchema                      `json:"LinkSchema"`
	ListSchema                         ListSchema                      `json:"ListSchema"`
	MappingSchema                      MappingSchema                   `json:"MappingSchema"`
	NavSchema                          NavSchema                       `json:"NavSchema"`
	NavItemSchema                      NavItemSchema                   `json:"NavItemSchema"`
	IconItemSchema                     IconItemSchema                  `json:"IconItemSchema"`
	NavOverflow                        NavOverflow                     `json:"NavOverflow"`
	ReactCSSProperties                 ReactCSSProperties              `json:"React.CSSProperties"`
	PropertyAlignmentBaseline          right                           `json:"Property.AlignmentBaseline"`
	Globals                            string                          `json:"Globals"`
	PropertyBaselineShiftstringnumber  right                           `json:"Property.BaselineShift<(string|number)>"`
	PropertyClip                       right                           `json:"Property.Clip"`
	PropertyClipPath                   right                           `json:"Property.ClipPath"`
	DataTypeGeometryBox                right                           `json:"DataType.GeometryBox"`
	DataTypeBox                        string                          `json:"DataType.Box"`
	PropertyClipRule                   right                           `json:"Property.ClipRule"`
	PropertyColor                      right                           `json:"Property.Color"`
	DataTypeColor                      right                           `json:"DataType.Color"`
	DataTypeNamedColor                 string                          `json:"DataType.NamedColor"`
	DataTypeDeprecatedSystemColor      string                          `json:"DataType.DeprecatedSystemColor"`
	PropertyColorInterpolation         right                           `json:"Property.ColorInterpolation"`
	PropertyColorRendering             right                           `json:"Property.ColorRendering"`
	PropertyCursor                     right                           `json:"Property.Cursor"`
	PropertyDirection                  right                           `json:"Property.Direction"`
	PropertyDisplay                    right                           `json:"Property.Display"`
	DataTypeDisplayOutside             string                          `json:"DataType.DisplayOutside"`
	DataTypeDisplayInside              string                          `json:"DataType.DisplayInside"`
	DataTypeDisplayInternal            string                          `json:"DataType.DisplayInternal"`
	DataTypeDisplayLegacy              string                          `json:"DataType.DisplayLegacy"`
	PropertyDominantBaseline           right                           `json:"Property.DominantBaseline"`
	PropertyFill                       right                           `json:"Property.Fill"`
	DataTypePaint                      right                           `json:"DataType.Paint"`
	PropertyFillOpacity                right                           `json:"Property.FillOpacity"`
	PropertyFillRule                   right                           `json:"Property.FillRule"`
	PropertyFilter                     right                           `json:"Property.Filter"`
	PropertyFloodColor                 right                           `json:"Property.FloodColor"`
	PropertyFloodOpacity               right                           `json:"Property.FloodOpacity"`
	PropertyFont                       right                           `json:"Property.Font"`
	PropertyFontFamily                 right                           `json:"Property.FontFamily"`
	DataTypeGenericFamily              string                          `json:"DataType.GenericFamily"`
	PropertyFontSizestringnumber       right                           `json:"Property.FontSize<(string|number)>"`
	DataTypeAbsoluteSize               string                          `json:"DataType.AbsoluteSize"`
	PropertyFontSizeAdjust             right                           `json:"Property.FontSizeAdjust"`
	PropertyFontStretch                right                           `json:"Property.FontStretch"`
	DataTypeFontStretchAbsolute        string                          `json:"DataType.FontStretchAbsolute"`
	PropertyFontStyle                  right                           `json:"Property.FontStyle"`
	PropertyFontVariant                right                           `json:"Property.FontVariant"`
	DataTypeEastAsianVariantValues     string                          `json:"DataType.EastAsianVariantValues"`
	PropertyFontWeight                 right                           `json:"Property.FontWeight"`
	DataTypeFontWeightAbsolute         OperatorType                    `json:"DataType.FontWeightAbsolute"`
	PropertyGlyphOrientationVertical   right                           `json:"Property.GlyphOrientationVertical"`
	PropertyImageRendering             right                           `json:"Property.ImageRendering"`
	PropertyLetterSpacingstringnumber  right                           `json:"Property.LetterSpacing<(string|number)>"`
	PropertyLightingColor              right                           `json:"Property.LightingColor"`
	PropertyLineHeightstringnumber     right                           `json:"Property.LineHeight<(string|number)>"`
	PropertyMarker                     right                           `json:"Property.Marker"`
	PropertyMarkerEnd                  right                           `json:"Property.MarkerEnd"`
	PropertyMarkerMid                  right                           `json:"Property.MarkerMid"`
	PropertyMarkerStart                right                           `json:"Property.MarkerStart"`
	PropertyMaskstringnumber           right                           `json:"Property.Mask<(string|number)>"`
	DataTypeMaskLayerstringnumber      right                           `json:"DataType.MaskLayer<(string|number)>"`
	DataTypePositionstringnumber       data                            `json:"DataType.Position<(string|number)>"`
	DataTypeRepeatStyle                string                          `json:"DataType.RepeatStyle"`
	DataTypeCompositingOperator        string                          `json:"DataType.CompositingOperator"`
	DataTypeMaskingMode                string                          `json:"DataType.MaskingMode"`
	PropertyOpacity                    right                           `json:"Property.Opacity"`
	PropertyOverflow                   right                           `json:"Property.Overflow"`
	PropertyPaintOrder                 right                           `json:"Property.PaintOrder"`
	PropertyPointerEvents              right                           `json:"Property.PointerEvents"`
	PropertyShapeRendering             right                           `json:"Property.ShapeRendering"`
	PropertyStopColor                  right                           `json:"Property.StopColor"`
	PropertyStopOpacity                right                           `json:"Property.StopOpacity"`
	PropertyStroke                     right                           `json:"Property.Stroke"`
	Field_0a048aec                     right                           `json:"Property.StrokeDasharray<(string|number)>"`
	DataTypeDasharraystringnumber      any                             `json:"DataType.Dasharray<(string|number)>"`
	Field_39430d8b                     right                           `json:"Property.StrokeDashoffset<(string|number)>"`
	PropertyStrokeLinecap              right                           `json:"Property.StrokeLinecap"`
	PropertyStrokeLinejoin             right                           `json:"Property.StrokeLinejoin"`
	PropertyStrokeMiterlimit           right                           `json:"Property.StrokeMiterlimit"`
	PropertyStrokeOpacity              right                           `json:"Property.StrokeOpacity"`
	PropertyStrokeWidthstringnumber    right                           `json:"Property.StrokeWidth<(string|number)>"`
	PropertyTextAnchor                 right                           `json:"Property.TextAnchor"`
	PropertyTextDecorationstringnumber right                           `json:"Property.TextDecoration<(string|number)>"`
	PropertyTextRendering              right                           `json:"Property.TextRendering"`
	PropertyUnicodeBidi                right                           `json:"Property.UnicodeBidi"`
	PropertyVectorEffect               right                           `json:"Property.VectorEffect"`
	PropertyVisibility                 right                           `json:"Property.Visibility"`
	PropertyWhiteSpace                 right                           `json:"Property.WhiteSpace"`
	PropertyWordSpacingstringnumber    right                           `json:"Property.WordSpacing<(string|number)>"`
	PropertyWritingMode                right                           `json:"Property.WritingMode"`
	PropertyAzimuth                    right                           `json:"Property.Azimuth"`
	PropertyBoxAlign                   right                           `json:"Property.BoxAlign"`
	PropertyBoxDirection               right                           `json:"Property.BoxDirection"`
	PropertyBoxFlex                    right                           `json:"Property.BoxFlex"`
	PropertyBoxFlexGroup               right                           `json:"Property.BoxFlexGroup"`
	PropertyBoxLines                   right                           `json:"Property.BoxLines"`
	PropertyBoxOrdinalGroup            right                           `json:"Property.BoxOrdinalGroup"`
	PropertyBoxOrient                  right                           `json:"Property.BoxOrient"`
	PropertyBoxPack                    right                           `json:"Property.BoxPack"`
	PropertyGridColumnGapstringnumber  right                           `json:"Property.GridColumnGap<(string|number)>"`
	PropertyGridGapstringnumber        right                           `json:"Property.GridGap<(string|number)>"`
	PropertyGridRowGapstringnumber     right                           `json:"Property.GridRowGap<(string|number)>"`
	PropertyImeMode                    right                           `json:"Property.ImeMode"`
	PropertyInsetBlockstringnumber     right                           `json:"Property.InsetBlock<(string|number)>"`
	PropertyInsetBlockEndstringnumber  right                           `json:"Property.InsetBlockEnd<(string|number)>"`
	Field_0c0d7952                     right                           `json:"Property.InsetBlockStart<(string|number)>"`
	PropertyInsetInlinestringnumber    right                           `json:"Property.InsetInline<(string|number)>"`
	PropertyInsetInlineEndstringnumber right                           `json:"Property.InsetInlineEnd<(string|number)>"`
	Field_a6ec5815                     right                           `json:"Property.InsetInlineStart<(string|number)>"`
	Field_800f1918                     right                           `json:"Property.ScrollSnapCoordinate<(string|number)>"`
	Field_8cea0839                     right                           `json:"Property.ScrollSnapDestination<(string|number)>"`
	PropertyScrollSnapPointsX          right                           `json:"Property.ScrollSnapPointsX"`
	PropertyScrollSnapPointsY          right                           `json:"Property.ScrollSnapPointsY"`
	PropertyScrollSnapTypeX            right                           `json:"Property.ScrollSnapTypeX"`
	PropertyScrollSnapTypeY            right                           `json:"Property.ScrollSnapTypeY"`
	PropertyLineBreak                  right                           `json:"Property.LineBreak"`
	PropertyUserSelect                 right                           `json:"Property.UserSelect"`
	PropertyBackfaceVisibility         right                           `json:"Property.BackfaceVisibility"`
	PropertyBackgroundClip             right                           `json:"Property.BackgroundClip"`
	PropertyBoxDecorationBreak         right                           `json:"Property.BoxDecorationBreak"`
	PropertyBackgroundOrigin           right                           `json:"Property.BackgroundOrigin"`
	PropertyBackgroundSizestringnumber right                           `json:"Property.BackgroundSize<(string|number)>"`
	DataTypeBgSizestringnumber         data                            `json:"DataType.BgSize<(string|number)>"`
	PropertyBorderRadiusstringnumber   right                           `json:"Property.BorderRadius<(string|number)>"`
	Field_2d82f04b                     right                           `json:"Property.BorderBottomLeftRadius<(string|number)>"`
	Field_bba18b0c                     right                           `json:"Property.BorderBottomRightRadius<(string|number)>"`
	Field_51341ad2                     right                           `json:"Property.BorderTopLeftRadius<(string|number)>"`
	Field_bea291ef                     right                           `json:"Property.BorderTopRightRadius<(string|number)>"`
	PropertyBoxShadow                  right                           `json:"Property.BoxShadow"`
	PropertyMozFloatEdge               right                           `json:"Property.MozFloatEdge"`
	PropertyMozForceBrokenImageIcon    right                           `json:"Property.MozForceBrokenImageIcon"`
	PropertyOutlinestringnumber        right                           `json:"Property.Outline<(string|number)>"`
	DataTypeLineStyle                  string                          `json:"DataType.LineStyle"`
	DataTypeLineWidthstringnumber      data                            `json:"DataType.LineWidth<(string|number)>"`
	PropertyOutlineColor               right                           `json:"Property.OutlineColor"`
	PropertyOutlineStyle               right                           `json:"Property.OutlineStyle"`
	PropertyOutlineWidthstringnumber   right                           `json:"Property.OutlineWidth<(string|number)>"`
	PropertyPerspectivestringnumber    right                           `json:"Property.Perspective<(string|number)>"`
	Field_38b841cf                     right                           `json:"Property.PerspectiveOrigin<(string|number)>"`
	PropertyTextAlignLast              right                           `json:"Property.TextAlignLast"`
	PropertyTextDecorationColor        right                           `json:"Property.TextDecorationColor"`
	PropertyTextDecorationLine         right                           `json:"Property.TextDecorationLine"`
	PropertyTextDecorationStyle        right                           `json:"Property.TextDecorationStyle"`
	PropertyTransform                  right                           `json:"Property.Transform"`
	Field_a5a96495                     right                           `json:"Property.TransformOrigin<(string|number)>"`
	PropertyTransformStyle             right                           `json:"Property.TransformStyle"`
	PropertyTransitionstring           right                           `json:"Property.Transition<string>"`
	DataTypeSingleTransitionstring     right                           `json:"DataType.SingleTransition<string>"`
	DataTypeEasingFunction             right                           `json:"DataType.EasingFunction"`
	DataTypeCubicBezierTimingFunction  string                          `json:"DataType.CubicBezierTimingFunction"`
	DataTypeStepTimingFunction         string                          `json:"DataType.StepTimingFunction"`
	PropertyTransitionDelaystring      right                           `json:"Property.TransitionDelay<string>"`
	PropertyTransitionDurationstring   right                           `json:"Property.TransitionDuration<string>"`
	PropertyTransitionProperty         right                           `json:"Property.TransitionProperty"`
	PropertyTransitionTimingFunction   right                           `json:"Property.TransitionTimingFunction"`
	PropertyMozUserInput               right                           `json:"Property.MozUserInput"`
	PropertyAnimationstring            right                           `json:"Property.Animation<string>"`
	DataTypeSingleAnimationstring      right                           `json:"DataType.SingleAnimation<string>"`
	DataTypeSingleAnimationDirection   string                          `json:"DataType.SingleAnimationDirection"`
	DataTypeSingleAnimationFillMode    string                          `json:"DataType.SingleAnimationFillMode"`
	DataTypeSingleAnimationTimeline    string                          `json:"DataType.SingleAnimationTimeline"`
	PropertyAnimationDelaystring       right                           `json:"Property.AnimationDelay<string>"`
	PropertyAnimationDirection         right                           `json:"Property.AnimationDirection"`
	PropertyAnimationDurationstring    right                           `json:"Property.AnimationDuration<string>"`
	PropertyAnimationFillMode          right                           `json:"Property.AnimationFillMode"`
	PropertyAnimationIterationCount    right                           `json:"Property.AnimationIterationCount"`
	PropertyAnimationName              right                           `json:"Property.AnimationName"`
	PropertyAnimationPlayState         right                           `json:"Property.AnimationPlayState"`
	PropertyAnimationTimingFunction    right                           `json:"Property.AnimationTimingFunction"`
	PropertyBorderImage                right                           `json:"Property.BorderImage"`
	PropertyObjectFit                  right                           `json:"Property.ObjectFit"`
	PropertyObjectPositionstringnumber right                           `json:"Property.ObjectPosition<(string|number)>"`
	PropertyTabSizestringnumber        right                           `json:"Property.TabSize<(string|number)>"`
	PropertyTextOverflow               right                           `json:"Property.TextOverflow"`
	PropertyColumnRulestringnumber     right                           `json:"Property.ColumnRule<(string|number)>"`
	PropertyColumnsstringnumber        right                           `json:"Property.Columns<(string|number)>"`
	Field_396a0a9b                     right                           `json:"Property.MozOutlineRadius<(string|number)>"`
	PropertyMsContentZoomLimit         right                           `json:"Property.MsContentZoomLimit"`
	PropertyMsContentZoomSnap          right                           `json:"Property.MsContentZoomSnap"`
	PropertyFlexstringnumber           right                           `json:"Property.Flex<(string|number)>"`
	PropertyMsScrollLimit              right                           `json:"Property.MsScrollLimit"`
	PropertyMsScrollSnapX              right                           `json:"Property.MsScrollSnapX"`
	PropertyMsScrollSnapY              right                           `json:"Property.MsScrollSnapY"`
	Field_198a7253                     right                           `json:"Property.WebkitBorderBefore<(string|number)>"`
	PropertyFlexFlow                   right                           `json:"Property.FlexFlow"`
	PropertyWebkitMaskstringnumber     right                           `json:"Property.WebkitMask<(string|number)>"`
	PropertyMaskBorder                 right                           `json:"Property.MaskBorder"`
	PropertyTextEmphasis               right                           `json:"Property.TextEmphasis"`
	Field_644866cf                     right                           `json:"Property.WebkitTextStroke<(string|number)>"`
	PropertyMozAppearance              right                           `json:"Property.MozAppearance"`
	PropertyMozBinding                 right                           `json:"Property.MozBinding"`
	PropertyMozBorderBottomColors      right                           `json:"Property.MozBorderBottomColors"`
	PropertyBorderInlineEndColor       right                           `json:"Property.BorderInlineEndColor"`
	PropertyBorderInlineEndStyle       right                           `json:"Property.BorderInlineEndStyle"`
	Field_ff1eba0b                     right                           `json:"Property.BorderInlineEndWidth<(string|number)>"`
	PropertyMozBorderLeftColors        right                           `json:"Property.MozBorderLeftColors"`
	PropertyMozBorderRightColors       right                           `json:"Property.MozBorderRightColors"`
	PropertyBorderInlineStartColor     right                           `json:"Property.BorderInlineStartColor"`
	PropertyBorderInlineStartStyle     right                           `json:"Property.BorderInlineStartStyle"`
	PropertyMozBorderTopColors         right                           `json:"Property.MozBorderTopColors"`
	PropertyBoxSizing                  right                           `json:"Property.BoxSizing"`
	PropertyColumnCount                right                           `json:"Property.ColumnCount"`
	PropertyColumnFill                 right                           `json:"Property.ColumnFill"`
	PropertyColumnRuleColor            right                           `json:"Property.ColumnRuleColor"`
	PropertyColumnRuleStyle            right                           `json:"Property.ColumnRuleStyle"`
	Field_8eba95b9                     right                           `json:"Property.ColumnRuleWidth<(string|number)>"`
	PropertyColumnWidthstringnumber    right                           `json:"Property.ColumnWidth<(string|number)>"`
	PropertyMozContextProperties       right                           `json:"Property.MozContextProperties"`
	PropertyFontFeatureSettings        right                           `json:"Property.FontFeatureSettings"`
	PropertyFontLanguageOverride       right                           `json:"Property.FontLanguageOverride"`
	PropertyHyphens                    right                           `json:"Property.Hyphens"`
	PropertyMozImageRegion             right                           `json:"Property.MozImageRegion"`
	Field_9a10601f                     right                           `json:"Property.MarginInlineEnd<(string|number)>"`
	Field_eeb35265                     right                           `json:"Property.MarginInlineStart<(string|number)>"`
	PropertyMozOrient                  right                           `json:"Property.MozOrient"`
	PropertyFontSmoothstringnumber     right                           `json:"Property.FontSmooth<(string|number)>"`
	Field_7d93d28d                     right                           `json:"Property.MozOutlineRadiusBottomleft<(string|number)>"`
	Field_aac4dd8c                     right                           `json:"Property.MozOutlineRadiusBottomright<(string|number)>"`
	Field_b0ce29e1                     right                           `json:"Property.MozOutlineRadiusTopleft<(string|number)>"`
	Field_598eaca3                     right                           `json:"Property.MozOutlineRadiusTopright<(string|number)>"`
	Field_14cf1d9f                     right                           `json:"Property.PaddingInlineEnd<(string|number)>"`
	Field_3f35d291                     right                           `json:"Property.PaddingInlineStart<(string|number)>"`
	PropertyMozStackSizing             right                           `json:"Property.MozStackSizing"`
	PropertyMozTextBlink               right                           `json:"Property.MozTextBlink"`
	PropertyTextSizeAdjust             right                           `json:"Property.TextSizeAdjust"`
	PropertyMozUserFocus               right                           `json:"Property.MozUserFocus"`
	PropertyMozUserModify              right                           `json:"Property.MozUserModify"`
	PropertyMozWindowDragging          right                           `json:"Property.MozWindowDragging"`
	PropertyMozWindowShadow            right                           `json:"Property.MozWindowShadow"`
	PropertyMsAccelerator              right                           `json:"Property.MsAccelerator"`
	PropertyMsBlockProgression         right                           `json:"Property.MsBlockProgression"`
	PropertyMsContentZoomChaining      right                           `json:"Property.MsContentZoomChaining"`
	PropertyMsContentZoomLimitMax      right                           `json:"Property.MsContentZoomLimitMax"`
	PropertyMsContentZoomLimitMin      right                           `json:"Property.MsContentZoomLimitMin"`
	PropertyMsContentZoomSnapPoints    right                           `json:"Property.MsContentZoomSnapPoints"`
	PropertyMsContentZoomSnapType      right                           `json:"Property.MsContentZoomSnapType"`
	PropertyMsContentZooming           right                           `json:"Property.MsContentZooming"`
	PropertyMsFilter                   right                           `json:"Property.MsFilter"`
	PropertyFlexDirection              right                           `json:"Property.FlexDirection"`
	PropertyFlexGrow                   right                           `json:"Property.FlexGrow"`
	PropertyMsFlowFrom                 right                           `json:"Property.MsFlowFrom"`
	PropertyMsFlowInto                 right                           `json:"Property.MsFlowInto"`
	PropertyMsGridColumnsstringnumber  right                           `json:"Property.MsGridColumns<(string|number)>"`
	DataTypeTrackBreadthstringnumber   data                            `json:"DataType.TrackBreadth<(string|number)>"`
	PropertyMsGridRowsstringnumber     right                           `json:"Property.MsGridRows<(string|number)>"`
	PropertyMsHighContrastAdjust       right                           `json:"Property.MsHighContrastAdjust"`
	PropertyMsHyphenateLimitChars      right                           `json:"Property.MsHyphenateLimitChars"`
	PropertyMsHyphenateLimitLines      right                           `json:"Property.MsHyphenateLimitLines"`
	Field_5606d2b5                     right                           `json:"Property.MsHyphenateLimitZone<(string|number)>"`
	PropertyMsImeAlign                 right                           `json:"Property.MsImeAlign"`
	PropertyOrder                      right                           `json:"Property.Order"`
	PropertyMsOverflowStyle            right                           `json:"Property.MsOverflowStyle"`
	PropertyOverflowX                  right                           `json:"Property.OverflowX"`
	PropertyOverflowY                  right                           `json:"Property.OverflowY"`
	PropertyMsScrollChaining           right                           `json:"Property.MsScrollChaining"`
	Field_51de3038                     right                           `json:"Property.MsScrollLimitXMax<(string|number)>"`
	Field_005c0ae3                     right                           `json:"Property.MsScrollLimitXMin<(string|number)>"`
	Field_fbd9d68d                     right                           `json:"Property.MsScrollLimitYMax<(string|number)>"`
	Field_12ef1f7d                     right                           `json:"Property.MsScrollLimitYMin<(string|number)>"`
	PropertyMsScrollRails              right                           `json:"Property.MsScrollRails"`
	PropertyMsScrollSnapPointsX        right                           `json:"Property.MsScrollSnapPointsX"`
	PropertyMsScrollSnapPointsY        right                           `json:"Property.MsScrollSnapPointsY"`
	PropertyMsScrollSnapType           right                           `json:"Property.MsScrollSnapType"`
	PropertyMsScrollTranslation        right                           `json:"Property.MsScrollTranslation"`
	PropertyMsScrollbar3dlightColor    right                           `json:"Property.MsScrollbar3dlightColor"`
	PropertyMsScrollbarArrowColor      right                           `json:"Property.MsScrollbarArrowColor"`
	PropertyMsScrollbarBaseColor       right                           `json:"Property.MsScrollbarBaseColor"`
	PropertyMsScrollbarDarkshadowColor right                           `json:"Property.MsScrollbarDarkshadowColor"`
	PropertyMsScrollbarFaceColor       right                           `json:"Property.MsScrollbarFaceColor"`
	PropertyMsScrollbarHighlightColor  right                           `json:"Property.MsScrollbarHighlightColor"`
	PropertyMsScrollbarShadowColor     right                           `json:"Property.MsScrollbarShadowColor"`
	PropertyMsScrollbarTrackColor      right                           `json:"Property.MsScrollbarTrackColor"`
	PropertyMsTextAutospace            right                           `json:"Property.MsTextAutospace"`
	PropertyTextCombineUpright         right                           `json:"Property.TextCombineUpright"`
	PropertyTouchAction                right                           `json:"Property.TouchAction"`
	PropertyMsTouchSelect              right                           `json:"Property.MsTouchSelect"`
	PropertyMsUserSelect               right                           `json:"Property.MsUserSelect"`
	PropertyWordBreak                  right                           `json:"Property.WordBreak"`
	PropertyMsWrapFlow                 right                           `json:"Property.MsWrapFlow"`
	PropertyMsWrapMarginstringnumber   right                           `json:"Property.MsWrapMargin<(string|number)>"`
	PropertyMsWrapThrough              right                           `json:"Property.MsWrapThrough"`
	PropertyAlignContent               right                           `json:"Property.AlignContent"`
	DataTypeContentDistribution        string                          `json:"DataType.ContentDistribution"`
	DataTypeContentPosition            string                          `json:"DataType.ContentPosition"`
	PropertyAlignItems                 right                           `json:"Property.AlignItems"`
	DataTypeSelfPosition               string                          `json:"DataType.SelfPosition"`
	PropertyAlignSelf                  right                           `json:"Property.AlignSelf"`
	PropertyWebkitAppearance           right                           `json:"Property.WebkitAppearance"`
	PropertyBackdropFilter             right                           `json:"Property.BackdropFilter"`
	PropertyWebkitBorderBeforeColor    right                           `json:"Property.WebkitBorderBeforeColor"`
	PropertyWebkitBorderBeforeStyle    right                           `json:"Property.WebkitBorderBeforeStyle"`
	Field_7af05760                     right                           `json:"Property.WebkitBorderBeforeWidth<(string|number)>"`
	PropertyBorderImageSlice           right                           `json:"Property.BorderImageSlice"`
	Field_113d0d97                     right                           `json:"Property.WebkitBoxReflect<(string|number)>"`
	PropertyColumnSpan                 right                           `json:"Property.ColumnSpan"`
	PropertyFlexBasisstringnumber      right                           `json:"Property.FlexBasis<(string|number)>"`
	PropertyFlexShrink                 right                           `json:"Property.FlexShrink"`
	PropertyFlexWrap                   right                           `json:"Property.FlexWrap"`
	PropertyFontKerning                right                           `json:"Property.FontKerning"`
	PropertyFontVariantLigatures       right                           `json:"Property.FontVariantLigatures"`
	PropertyHyphenateCharacter         right                           `json:"Property.HyphenateCharacter"`
	PropertyInitialLetter              right                           `json:"Property.InitialLetter"`
	PropertyJustifyContent             right                           `json:"Property.JustifyContent"`
	PropertyWebkitLineClamp            right                           `json:"Property.WebkitLineClamp"`
	PropertyWebkitMaskAttachment       right                           `json:"Property.WebkitMaskAttachment"`
	DataTypeAttachment                 string                          `json:"DataType.Attachment"`
	Field_e2538c6e                     right                           `json:"Property.MaskBorderOutset<(string|number)>"`
	PropertyMaskBorderRepeat           right                           `json:"Property.MaskBorderRepeat"`
	PropertyMaskBorderSlice            right                           `json:"Property.MaskBorderSlice"`
	PropertyMaskBorderSource           right                           `json:"Property.MaskBorderSource"`
	Field_7883ae54                     right                           `json:"Property.MaskBorderWidth<(string|number)>"`
	PropertyWebkitMaskClip             right                           `json:"Property.WebkitMaskClip"`
	PropertyWebkitMaskComposite        right                           `json:"Property.WebkitMaskComposite"`
	DataTypeCompositeStyle             string                          `json:"DataType.CompositeStyle"`
	PropertyWebkitMaskImage            right                           `json:"Property.WebkitMaskImage"`
	PropertyWebkitMaskOrigin           right                           `json:"Property.WebkitMaskOrigin"`
	Field_54a1328e                     right                           `json:"Property.WebkitMaskPosition<(string|number)>"`
	Field_be8e52c1                     right                           `json:"Property.WebkitMaskPositionX<(string|number)>"`
	Field_4f41f3af                     right                           `json:"Property.WebkitMaskPositionY<(string|number)>"`
	PropertyWebkitMaskRepeat           right                           `json:"Property.WebkitMaskRepeat"`
	PropertyWebkitMaskRepeatX          right                           `json:"Property.WebkitMaskRepeatX"`
	PropertyWebkitMaskRepeatY          right                           `json:"Property.WebkitMaskRepeatY"`
	PropertyWebkitMaskSizestringnumber right                           `json:"Property.WebkitMaskSize<(string|number)>"`
	PropertyMaxInlineSizestringnumber  right                           `json:"Property.MaxInlineSize<(string|number)>"`
	PropertyWebkitOverflowScrolling    right                           `json:"Property.WebkitOverflowScrolling"`
	PropertyPrintColorAdjust           right                           `json:"Property.PrintColorAdjust"`
	PropertyRubyPosition               right                           `json:"Property.RubyPosition"`
	PropertyScrollSnapType             right                           `json:"Property.ScrollSnapType"`
	PropertyShapeMarginstringnumber    right                           `json:"Property.ShapeMargin<(string|number)>"`
	PropertyWebkitTapHighlightColor    right                           `json:"Property.WebkitTapHighlightColor"`
	PropertyTextDecorationSkip         right                           `json:"Property.TextDecorationSkip"`
	PropertyTextEmphasisColor          right                           `json:"Property.TextEmphasisColor"`
	PropertyTextEmphasisPosition       right                           `json:"Property.TextEmphasisPosition"`
	PropertyTextEmphasisStyle          right                           `json:"Property.TextEmphasisStyle"`
	PropertyWebkitTextFillColor        right                           `json:"Property.WebkitTextFillColor"`
	PropertyTextOrientation            right                           `json:"Property.TextOrientation"`
	PropertyWebkitTextStrokeColor      right                           `json:"Property.WebkitTextStrokeColor"`
	Field_c8e57b6c                     right                           `json:"Property.WebkitTextStrokeWidth<(string|number)>"`
	PropertyTextUnderlinePosition      right                           `json:"Property.TextUnderlinePosition"`
	PropertyWebkitTouchCallout         right                           `json:"Property.WebkitTouchCallout"`
	PropertyWebkitUserModify           right                           `json:"Property.WebkitUserModify"`
	PropertyAll                        RootSchema                      `json:"Property.All"`
	PropertyAnimationRangestringnumber right                           `json:"Property.AnimationRange<(string|number)>"`
	DataTypeTimelineRangeName          string                          `json:"DataType.TimelineRangeName"`
	PropertyBackgroundstringnumber     right                           `json:"Property.Background<(string|number)>"`
	DataTypeFinalBgLayerstringnumber   right                           `json:"DataType.FinalBgLayer<(string|number)>"`
	DataTypeBgPositionstringnumber     data                            `json:"DataType.BgPosition<(string|number)>"`
	Field_3030d98a                     right                           `json:"Property.BackgroundPosition<(string|number)>"`
	PropertyBorderstringnumber         right                           `json:"Property.Border<(string|number)>"`
	PropertyBorderBlockstringnumber    right                           `json:"Property.BorderBlock<(string|number)>"`
	PropertyBorderBlockEndstringnumber right                           `json:"Property.BorderBlockEnd<(string|number)>"`
	Field_9b54f67b                     right                           `json:"Property.BorderBlockStart<(string|number)>"`
	PropertyBorderBottomstringnumber   right                           `json:"Property.BorderBottom<(string|number)>"`
	PropertyBorderColor                right                           `json:"Property.BorderColor"`
	PropertyBorderInlinestringnumber   right                           `json:"Property.BorderInline<(string|number)>"`
	Field_bbf3394b                     right                           `json:"Property.BorderInlineEnd<(string|number)>"`
	Field_626f0a7d                     right                           `json:"Property.BorderInlineStart<(string|number)>"`
	PropertyBorderLeftstringnumber     right                           `json:"Property.BorderLeft<(string|number)>"`
	PropertyBorderRightstringnumber    right                           `json:"Property.BorderRight<(string|number)>"`
	PropertyBorderStyle                right                           `json:"Property.BorderStyle"`
	PropertyBorderTopstringnumber      right                           `json:"Property.BorderTop<(string|number)>"`
	PropertyBorderWidthstringnumber    right                           `json:"Property.BorderWidth<(string|number)>"`
	PropertyCaret                      right                           `json:"Property.Caret"`
	Field_de93c262                     right                           `json:"Property.ContainIntrinsicSize<(string|number)>"`
	PropertyContainer                  right                           `json:"Property.Container"`
	PropertyGapstringnumber            right                           `json:"Property.Gap<(string|number)>"`
	PropertyGrid                       right                           `json:"Property.Grid"`
	PropertyGridArea                   right                           `json:"Property.GridArea"`
	DataTypeGridLine                   OperatorType                    `json:"DataType.GridLine"`
	PropertyGridColumn                 right                           `json:"Property.GridColumn"`
	PropertyGridRow                    right                           `json:"Property.GridRow"`
	PropertyGridTemplate               right                           `json:"Property.GridTemplate"`
	PropertyInsetstringnumber          right                           `json:"Property.Inset<(string|number)>"`
	PropertyLineClamp                  right                           `json:"Property.LineClamp"`
	PropertyListStyle                  right                           `json:"Property.ListStyle"`
	PropertyMarginstringnumber         right                           `json:"Property.Margin<(string|number)>"`
	PropertyMarginBlockstringnumber    right                           `json:"Property.MarginBlock<(string|number)>"`
	PropertyMarginInlinestringnumber   right                           `json:"Property.MarginInline<(string|number)>"`
	PropertyOffsetstringnumber         right                           `json:"Property.Offset<(string|number)>"`
	PropertyOverscrollBehavior         right                           `json:"Property.OverscrollBehavior"`
	PropertyPaddingstringnumber        right                           `json:"Property.Padding<(string|number)>"`
	PropertyPaddingBlockstringnumber   right                           `json:"Property.PaddingBlock<(string|number)>"`
	PropertyPaddingInlinestringnumber  right                           `json:"Property.PaddingInline<(string|number)>"`
	PropertyPlaceContent               right                           `json:"Property.PlaceContent"`
	PropertyPlaceItems                 right                           `json:"Property.PlaceItems"`
	PropertyPlaceSelf                  right                           `json:"Property.PlaceSelf"`
	PropertyScrollMarginstringnumber   right                           `json:"Property.ScrollMargin<(string|number)>"`
	Field_69de88de                     right                           `json:"Property.ScrollMarginBlock<(string|number)>"`
	Field_82cd4911                     right                           `json:"Property.ScrollMarginInline<(string|number)>"`
	PropertyScrollPaddingstringnumber  right                           `json:"Property.ScrollPadding<(string|number)>"`
	Field_094090c6                     right                           `json:"Property.ScrollPaddingBlock<(string|number)>"`
	Field_3465a829                     right                           `json:"Property.ScrollPaddingInline<(string|number)>"`
	PropertyScrollTimeline             right                           `json:"Property.ScrollTimeline"`
	PropertyViewTimeline               right                           `json:"Property.ViewTimeline"`
	PropertyAccentColor                right                           `json:"Property.AccentColor"`
	PropertyAlignTracks                right                           `json:"Property.AlignTracks"`
	PropertyAnimationComposition       right                           `json:"Property.AnimationComposition"`
	DataTypeSingleAnimationComposition string                          `json:"DataType.SingleAnimationComposition"`
	Field_ab42eb79                     right                           `json:"Property.AnimationRangeEnd<(string|number)>"`
	Field_0ba8715f                     right                           `json:"Property.AnimationRangeStart<(string|number)>"`
	PropertyAnimationTimeline          right                           `json:"Property.AnimationTimeline"`
	PropertyAppearance                 right                           `json:"Property.Appearance"`
	DataTypeCompatAuto                 string                          `json:"DataType.CompatAuto"`
	PropertyAspectRatio                right                           `json:"Property.AspectRatio"`
	PropertyBackgroundAttachment       right                           `json:"Property.BackgroundAttachment"`
	PropertyBackgroundBlendMode        right                           `json:"Property.BackgroundBlendMode"`
	DataTypeBlendMode                  string                          `json:"DataType.BlendMode"`
	PropertyBackgroundColor            right                           `json:"Property.BackgroundColor"`
	PropertyBackgroundImage            right                           `json:"Property.BackgroundImage"`
	Field_908da3e9                     right                           `json:"Property.BackgroundPositionX<(string|number)>"`
	Field_e2ba4068                     right                           `json:"Property.BackgroundPositionY<(string|number)>"`
	PropertyBackgroundRepeat           right                           `json:"Property.BackgroundRepeat"`
	PropertyBlockOverflow              right                           `json:"Property.BlockOverflow"`
	PropertyBlockSizestringnumber      right                           `json:"Property.BlockSize<(string|number)>"`
	PropertyBorderBlockColor           right                           `json:"Property.BorderBlockColor"`
	PropertyBorderBlockEndColor        right                           `json:"Property.BorderBlockEndColor"`
	PropertyBorderBlockEndStyle        right                           `json:"Property.BorderBlockEndStyle"`
	Field_e4103678                     right                           `json:"Property.BorderBlockEndWidth<(string|number)>"`
	PropertyBorderBlockStartColor      right                           `json:"Property.BorderBlockStartColor"`
	PropertyBorderBlockStartStyle      right                           `json:"Property.BorderBlockStartStyle"`
	Field_27264080                     right                           `json:"Property.BorderBlockStartWidth<(string|number)>"`
	PropertyBorderBlockStyle           right                           `json:"Property.BorderBlockStyle"`
	Field_dd5bb43d                     right                           `json:"Property.BorderBlockWidth<(string|number)>"`
	PropertyBorderBottomColor          right                           `json:"Property.BorderBottomColor"`
	PropertyBorderBottomStyle          right                           `json:"Property.BorderBottomStyle"`
	Field_e12142ad                     right                           `json:"Property.BorderBottomWidth<(string|number)>"`
	PropertyBorderCollapse             right                           `json:"Property.BorderCollapse"`
	Field_18eaeae3                     right                           `json:"Property.BorderEndEndRadius<(string|number)>"`
	Field_db7478fd                     right                           `json:"Property.BorderEndStartRadius<(string|number)>"`
	Field_a4eac0cb                     right                           `json:"Property.BorderImageOutset<(string|number)>"`
	PropertyBorderImageRepeat          right                           `json:"Property.BorderImageRepeat"`
	PropertyBorderImageSource          right                           `json:"Property.BorderImageSource"`
	Field_77038b3a                     right                           `json:"Property.BorderImageWidth<(string|number)>"`
	PropertyBorderInlineColor          right                           `json:"Property.BorderInlineColor"`
	Field_6a94bf8a                     right                           `json:"Property.BorderInlineStartWidth<(string|number)>"`
	PropertyBorderInlineStyle          right                           `json:"Property.BorderInlineStyle"`
	Field_60b42e25                     right                           `json:"Property.BorderInlineWidth<(string|number)>"`
	PropertyBorderLeftColor            right                           `json:"Property.BorderLeftColor"`
	PropertyBorderLeftStyle            right                           `json:"Property.BorderLeftStyle"`
	Field_c64e646c                     right                           `json:"Property.BorderLeftWidth<(string|number)>"`
	PropertyBorderRightColor           right                           `json:"Property.BorderRightColor"`
	PropertyBorderRightStyle           right                           `json:"Property.BorderRightStyle"`
	Field_c3d0c280                     right                           `json:"Property.BorderRightWidth<(string|number)>"`
	PropertyBorderSpacingstringnumber  right                           `json:"Property.BorderSpacing<(string|number)>"`
	Field_44c501b0                     right                           `json:"Property.BorderStartEndRadius<(string|number)>"`
	Field_9db2cba3                     right                           `json:"Property.BorderStartStartRadius<(string|number)>"`
	PropertyBorderTopColor             right                           `json:"Property.BorderTopColor"`
	PropertyBorderTopStyle             right                           `json:"Property.BorderTopStyle"`
	PropertyBorderTopWidthstringnumber right                           `json:"Property.BorderTopWidth<(string|number)>"`
	PropertyBottomstringnumber         right                           `json:"Property.Bottom<(string|number)>"`
	PropertyBreakAfter                 right                           `json:"Property.BreakAfter"`
	PropertyBreakBefore                right                           `json:"Property.BreakBefore"`
	PropertyBreakInside                right                           `json:"Property.BreakInside"`
	PropertyCaptionSide                right                           `json:"Property.CaptionSide"`
	PropertyCaretColor                 right                           `json:"Property.CaretColor"`
	PropertyCaretShape                 right                           `json:"Property.CaretShape"`
	PropertyClear                      right                           `json:"Property.Clear"`
	PropertyColorScheme                right                           `json:"Property.ColorScheme"`
	PropertyColumnGapstringnumber      right                           `json:"Property.ColumnGap<(string|number)>"`
	PropertyContain                    right                           `json:"Property.Contain"`
	Field_02975c25                     right                           `json:"Property.ContainIntrinsicBlockSize<(string|number)>"`
	Field_95e98a8b                     right                           `json:"Property.ContainIntrinsicHeight<(string|number)>"`
	Field_700e662a                     right                           `json:"Property.ContainIntrinsicInlineSize<(string|number)>"`
	Field_4f21f771                     right                           `json:"Property.ContainIntrinsicWidth<(string|number)>"`
	PropertyContainerName              right                           `json:"Property.ContainerName"`
	PropertyContainerType              right                           `json:"Property.ContainerType"`
	PropertyContent                    right                           `json:"Property.Content"`
	DataTypeContentList                right                           `json:"DataType.ContentList"`
	DataTypeQuote                      string                          `json:"DataType.Quote"`
	PropertyContentVisibility          right                           `json:"Property.ContentVisibility"`
	PropertyCounterIncrement           right                           `json:"Property.CounterIncrement"`
	PropertyCounterReset               right                           `json:"Property.CounterReset"`
	PropertyCounterSet                 right                           `json:"Property.CounterSet"`
	PropertyEmptyCells                 right                           `json:"Property.EmptyCells"`
	PropertyFloat                      right                           `json:"Property.Float"`
	PropertyFontOpticalSizing          right                           `json:"Property.FontOpticalSizing"`
	PropertyFontPalette                right                           `json:"Property.FontPalette"`
	PropertyFontSynthesis              right                           `json:"Property.FontSynthesis"`
	PropertyFontSynthesisPosition      right                           `json:"Property.FontSynthesisPosition"`
	PropertyFontSynthesisSmallCaps     right                           `json:"Property.FontSynthesisSmallCaps"`
	PropertyFontSynthesisStyle         right                           `json:"Property.FontSynthesisStyle"`
	PropertyFontSynthesisWeight        right                           `json:"Property.FontSynthesisWeight"`
	PropertyFontVariantAlternates      right                           `json:"Property.FontVariantAlternates"`
	PropertyFontVariantCaps            right                           `json:"Property.FontVariantCaps"`
	PropertyFontVariantEastAsian       right                           `json:"Property.FontVariantEastAsian"`
	PropertyFontVariantEmoji           right                           `json:"Property.FontVariantEmoji"`
	PropertyFontVariantNumeric         right                           `json:"Property.FontVariantNumeric"`
	PropertyFontVariantPosition        right                           `json:"Property.FontVariantPosition"`
	PropertyFontVariationSettings      right                           `json:"Property.FontVariationSettings"`
	PropertyForcedColorAdjust          right                           `json:"Property.ForcedColorAdjust"`
	Field_b71a1771                     right                           `json:"Property.GridAutoColumns<(string|number)>"`
	PropertyGridAutoFlow               right                           `json:"Property.GridAutoFlow"`
	PropertyGridAutoRowsstringnumber   right                           `json:"Property.GridAutoRows<(string|number)>"`
	PropertyGridColumnEnd              right                           `json:"Property.GridColumnEnd"`
	PropertyGridColumnStart            right                           `json:"Property.GridColumnStart"`
	PropertyGridRowEnd                 right                           `json:"Property.GridRowEnd"`
	PropertyGridRowStart               right                           `json:"Property.GridRowStart"`
	PropertyGridTemplateAreas          right                           `json:"Property.GridTemplateAreas"`
	Field_3a8540af                     right                           `json:"Property.GridTemplateColumns<(string|number)>"`
	Field_c571f87b                     right                           `json:"Property.GridTemplateRows<(string|number)>"`
	PropertyHangingPunctuation         right                           `json:"Property.HangingPunctuation"`
	PropertyHeightstringnumber         right                           `json:"Property.Height<(string|number)>"`
	PropertyHyphenateLimitChars        right                           `json:"Property.HyphenateLimitChars"`
	PropertyImageOrientation           right                           `json:"Property.ImageOrientation"`
	PropertyImageResolution            right                           `json:"Property.ImageResolution"`
	PropertyInlineSizestringnumber     right                           `json:"Property.InlineSize<(string|number)>"`
	PropertyInputSecurity              right                           `json:"Property.InputSecurity"`
	PropertyIsolation                  right                           `json:"Property.Isolation"`
	PropertyJustifyItems               right                           `json:"Property.JustifyItems"`
	PropertyJustifySelf                right                           `json:"Property.JustifySelf"`
	PropertyJustifyTracks              right                           `json:"Property.JustifyTracks"`
	PropertyLeftstringnumber           right                           `json:"Property.Left<(string|number)>"`
	PropertyLineHeightStepstringnumber right                           `json:"Property.LineHeightStep<(string|number)>"`
	PropertyListStyleImage             right                           `json:"Property.ListStyleImage"`
	PropertyListStylePosition          right                           `json:"Property.ListStylePosition"`
	PropertyListStyleType              right                           `json:"Property.ListStyleType"`
	PropertyMarginBlockEndstringnumber right                           `json:"Property.MarginBlockEnd<(string|number)>"`
	Field_c50d997d                     right                           `json:"Property.MarginBlockStart<(string|number)>"`
	PropertyMarginBottomstringnumber   right                           `json:"Property.MarginBottom<(string|number)>"`
	PropertyMarginLeftstringnumber     right                           `json:"Property.MarginLeft<(string|number)>"`
	PropertyMarginRightstringnumber    right                           `json:"Property.MarginRight<(string|number)>"`
	PropertyMarginTopstringnumber      right                           `json:"Property.MarginTop<(string|number)>"`
	PropertyMarginTrim                 right                           `json:"Property.MarginTrim"`
	PropertyMaskBorderMode             right                           `json:"Property.MaskBorderMode"`
	PropertyMaskClip                   right                           `json:"Property.MaskClip"`
	PropertyMaskComposite              right                           `json:"Property.MaskComposite"`
	PropertyMaskImage                  right                           `json:"Property.MaskImage"`
	PropertyMaskMode                   right                           `json:"Property.MaskMode"`
	PropertyMaskOrigin                 right                           `json:"Property.MaskOrigin"`
	PropertyMaskPositionstringnumber   right                           `json:"Property.MaskPosition<(string|number)>"`
	PropertyMaskRepeat                 right                           `json:"Property.MaskRepeat"`
	PropertyMaskSizestringnumber       right                           `json:"Property.MaskSize<(string|number)>"`
	PropertyMaskType                   right                           `json:"Property.MaskType"`
	PropertyMasonryAutoFlow            right                           `json:"Property.MasonryAutoFlow"`
	PropertyMathDepth                  right                           `json:"Property.MathDepth"`
	PropertyMathShift                  right                           `json:"Property.MathShift"`
	PropertyMathStyle                  right                           `json:"Property.MathStyle"`
	PropertyMaxBlockSizestringnumber   right                           `json:"Property.MaxBlockSize<(string|number)>"`
	PropertyMaxHeightstringnumber      right                           `json:"Property.MaxHeight<(string|number)>"`
	PropertyMaxLines                   right                           `json:"Property.MaxLines"`
	PropertyMaxWidthstringnumber       right                           `json:"Property.MaxWidth<(string|number)>"`
	PropertyMinBlockSizestringnumber   right                           `json:"Property.MinBlockSize<(string|number)>"`
	PropertyMinHeightstringnumber      right                           `json:"Property.MinHeight<(string|number)>"`
	PropertyMinInlineSizestringnumber  right                           `json:"Property.MinInlineSize<(string|number)>"`
	PropertyMinWidthstringnumber       right                           `json:"Property.MinWidth<(string|number)>"`
	PropertyMixBlendMode               right                           `json:"Property.MixBlendMode"`
	PropertyOffsetDistancestringnumber right                           `json:"Property.OffsetDistance<(string|number)>"`
	PropertyOffsetPath                 right                           `json:"Property.OffsetPath"`
	PropertyOffsetRotate               right                           `json:"Property.OffsetRotate"`
	PropertyOffsetAnchorstringnumber   right                           `json:"Property.OffsetAnchor<(string|number)>"`
	PropertyOffsetPositionstringnumber right                           `json:"Property.OffsetPosition<(string|number)>"`
	PropertyOrphans                    right                           `json:"Property.Orphans"`
	PropertyOutlineOffsetstringnumber  right                           `json:"Property.OutlineOffset<(string|number)>"`
	PropertyOverflowAnchor             right                           `json:"Property.OverflowAnchor"`
	PropertyOverflowBlock              right                           `json:"Property.OverflowBlock"`
	PropertyOverflowClipBox            right                           `json:"Property.OverflowClipBox"`
	Field_70c2cf43                     right                           `json:"Property.OverflowClipMargin<(string|number)>"`
	DataTypeVisualBox                  string                          `json:"DataType.VisualBox"`
	PropertyOverflowInline             right                           `json:"Property.OverflowInline"`
	PropertyOverflowWrap               right                           `json:"Property.OverflowWrap"`
	PropertyOverlay                    right                           `json:"Property.Overlay"`
	PropertyOverscrollBehaviorBlock    right                           `json:"Property.OverscrollBehaviorBlock"`
	PropertyOverscrollBehaviorInline   right                           `json:"Property.OverscrollBehaviorInline"`
	PropertyOverscrollBehaviorX        right                           `json:"Property.OverscrollBehaviorX"`
	PropertyOverscrollBehaviorY        right                           `json:"Property.OverscrollBehaviorY"`
	Field_f2d771e5                     right                           `json:"Property.PaddingBlockEnd<(string|number)>"`
	Field_80db4a85                     right                           `json:"Property.PaddingBlockStart<(string|number)>"`
	PropertyPaddingBottomstringnumber  right                           `json:"Property.PaddingBottom<(string|number)>"`
	PropertyPaddingLeftstringnumber    right                           `json:"Property.PaddingLeft<(string|number)>"`
	PropertyPaddingRightstringnumber   right                           `json:"Property.PaddingRight<(string|number)>"`
	PropertyPaddingTopstringnumber     right                           `json:"Property.PaddingTop<(string|number)>"`
	PropertyPage                       right                           `json:"Property.Page"`
	PropertyPageBreakAfter             right                           `json:"Property.PageBreakAfter"`
	PropertyPageBreakBefore            right                           `json:"Property.PageBreakBefore"`
	PropertyPageBreakInside            right                           `json:"Property.PageBreakInside"`
	PropertyPosition                   right                           `json:"Property.Position"`
	PropertyQuotes                     right                           `json:"Property.Quotes"`
	PropertyResize                     right                           `json:"Property.Resize"`
	PropertyRightstringnumber          right                           `json:"Property.Right<(string|number)>"`
	PropertyRotate                     right                           `json:"Property.Rotate"`
	PropertyRowGapstringnumber         right                           `json:"Property.RowGap<(string|number)>"`
	PropertyRubyAlign                  right                           `json:"Property.RubyAlign"`
	PropertyRubyMerge                  right                           `json:"Property.RubyMerge"`
	PropertyScale                      right                           `json:"Property.Scale"`
	PropertyScrollBehavior             right                           `json:"Property.ScrollBehavior"`
	Field_96c2bf02                     right                           `json:"Property.ScrollMarginBlockEnd<(string|number)>"`
	Field_ff99c3e3                     right                           `json:"Property.ScrollMarginBlockStart<(string|number)>"`
	Field_d767f7fb                     right                           `json:"Property.ScrollMarginBottom<(string|number)>"`
	Field_0a2f943e                     right                           `json:"Property.ScrollMarginInlineEnd<(string|number)>"`
	Field_c405bd95                     right                           `json:"Property.ScrollMarginInlineStart<(string|number)>"`
	Field_ac0dd0d0                     right                           `json:"Property.ScrollMarginLeft<(string|number)>"`
	Field_0064529c                     right                           `json:"Property.ScrollMarginRight<(string|number)>"`
	Field_c5da99a9                     right                           `json:"Property.ScrollMarginTop<(string|number)>"`
	Field_44caef0c                     right                           `json:"Property.ScrollPaddingBlockEnd<(string|number)>"`
	Field_48a3726a                     right                           `json:"Property.ScrollPaddingBlockStart<(string|number)>"`
	Field_54832dfe                     right                           `json:"Property.ScrollPaddingBottom<(string|number)>"`
	Field_fb0447ef                     right                           `json:"Property.ScrollPaddingInlineEnd<(string|number)>"`
	Field_5e9eb7ba                     right                           `json:"Property.ScrollPaddingInlineStart<(string|number)>"`
	Field_5693ff34                     right                           `json:"Property.ScrollPaddingLeft<(string|number)>"`
	Field_45704e15                     right                           `json:"Property.ScrollPaddingRight<(string|number)>"`
	Field_6e5644c1                     right                           `json:"Property.ScrollPaddingTop<(string|number)>"`
	PropertyScrollSnapAlign            right                           `json:"Property.ScrollSnapAlign"`
	PropertyScrollSnapStop             right                           `json:"Property.ScrollSnapStop"`
	PropertyScrollTimelineAxis         right                           `json:"Property.ScrollTimelineAxis"`
	PropertyScrollTimelineName         right                           `json:"Property.ScrollTimelineName"`
	PropertyScrollbarColor             right                           `json:"Property.ScrollbarColor"`
	PropertyScrollbarGutter            right                           `json:"Property.ScrollbarGutter"`
	PropertyScrollbarWidth             right                           `json:"Property.ScrollbarWidth"`
	PropertyShapeImageThreshold        right                           `json:"Property.ShapeImageThreshold"`
	PropertyShapeOutside               right                           `json:"Property.ShapeOutside"`
	PropertyTableLayout                right                           `json:"Property.TableLayout"`
	PropertyTextAlign                  right                           `json:"Property.TextAlign"`
	PropertyTextDecorationSkipInk      right                           `json:"Property.TextDecorationSkipInk"`
	Field_88a84f60                     right                           `json:"Property.TextDecorationThickness<(string|number)>"`
	PropertyTextIndentstringnumber     right                           `json:"Property.TextIndent<(string|number)>"`
	PropertyTextJustify                right                           `json:"Property.TextJustify"`
	PropertyTextShadow                 right                           `json:"Property.TextShadow"`
	PropertyTextTransform              right                           `json:"Property.TextTransform"`
	Field_94f20510                     right                           `json:"Property.TextUnderlineOffset<(string|number)>"`
	PropertyTextWrap                   right                           `json:"Property.TextWrap"`
	PropertyTimelineScope              right                           `json:"Property.TimelineScope"`
	PropertyTopstringnumber            right                           `json:"Property.Top<(string|number)>"`
	PropertyTransformBox               right                           `json:"Property.TransformBox"`
	PropertyTransitionBehavior         right                           `json:"Property.TransitionBehavior"`
	PropertyTranslatestringnumber      right                           `json:"Property.Translate<(string|number)>"`
	PropertyVerticalAlignstringnumber  right                           `json:"Property.VerticalAlign<(string|number)>"`
	PropertyViewTimelineAxis           right                           `json:"Property.ViewTimelineAxis"`
	Field_b40103be                     right                           `json:"Property.ViewTimelineInset<(string|number)>"`
	PropertyViewTimelineName           right                           `json:"Property.ViewTimelineName"`
	PropertyViewTransitionName         right                           `json:"Property.ViewTransitionName"`
	PropertyWhiteSpaceCollapse         right                           `json:"Property.WhiteSpaceCollapse"`
	PropertyWhiteSpaceTrim             right                           `json:"Property.WhiteSpaceTrim"`
	PropertyWidows                     right                           `json:"Property.Widows"`
	PropertyWidthstringnumber          right                           `json:"Property.Width<(string|number)>"`
	PropertyWillChange                 right                           `json:"Property.WillChange"`
	DataTypeAnimateableFeature         string                          `json:"DataType.AnimateableFeature"`
	PropertyWordWrap                   right                           `json:"Property.WordWrap"`
	PropertyZIndex                     right                           `json:"Property.ZIndex"`
	PropertyZoom                       right                           `json:"Property.Zoom"`
	OperationSchema                    OperationSchema                 `json:"OperationSchema"`
	PaginationSchema                   PaginationSchema                `json:"PaginationSchema"`
	MODE_TYPE                          string                          `json:"MODE_TYPE"`
	PaginationWrapperSchema            PaginationWrapperSchema         `json:"PaginationWrapperSchema"`
	PanelSchema                        PanelSchema                     `json:"PanelSchema"`
	PlainSchema                        PlainSchema                     `json:"PlainSchema"`
	ProgressSchema                     ProgressSchema                  `json:"ProgressSchema"`
	ColorMapType                       ColorMapType                    `json:"ColorMapType"`
	QRCodeSchema                       QRCodeSchema                    `json:"QRCodeSchema"`
	QRCodeImageSettings                QRCodeImageSettings             `json:"QRCodeImageSettings"`
	SearchBoxSchema                    SearchBoxSchema                 `json:"SearchBoxSchema"`
	ServiceSchema                      ServiceSchema                   `json:"ServiceSchema"`
	ComposedDataProvider               right                           `json:"ComposedDataProvider"`
	DataProvider                       data                            `json:"DataProvider"`
	DataProviderCollection             DataProviderCollection          `json:"DataProviderCollection"`
	SparkLineSchema                    SparkLineSchema                 `json:"SparkLineSchema"`
	StatusSchema                       StatusSchema                    `json:"StatusSchema"`
	StatusSource                       any                             `json:"StatusSource"`
	SpinnerSchema                      SpinnerSchema                   `json:"SpinnerSchema"`
	TableSchema                        TableSchema                     `json:"TableSchema"`
	TableSchema2                       TableSchema2                    `json:"TableSchema2"`
	AutoFillHeightObject               AutoFillHeightObject            `json:"AutoFillHeightObject"`
	TabsSchema                         TabsSchema                      `json:"TabsSchema"`
	TabSchema                          TabSchema                       `json:"TabSchema"`
	TabsMode                           string                          `json:"TabsMode"`
	TasksSchema                        TasksSchema                     `json:"TasksSchema"`
	VBoxSchema                         VBoxSchema                      `json:"VBoxSchema"`
	HboxRow                            HboxRow                         `json:"HboxRow"`
	VideoSchema                        VideoSchema                     `json:"VideoSchema"`
	WizardSchema                       WizardSchema                    `json:"WizardSchema"`
	WizardStepSchema                   WizardStepSchema                `json:"WizardStepSchema"`
	BaseApiObject                      BaseApiObject                   `json:"BaseApiObject"`
	ClassName                          data                            `json:"ClassName"`
	LabelAlign                         string                          `json:"LabelAlign"`
	WrapperSchema                      WrapperSchema                   `json:"WrapperSchema"`
	TooltipWrapperSchema               TooltipWrapperSchema            `json:"TooltipWrapperSchema"`
	Trigger                            string                          `json:"Trigger"`
	FormSchema                         FormSchema                      `json:"FormSchema"`
	AnchorNavSchema                    AnchorNavSchema                 `json:"AnchorNavSchema"`
	AnchorNavSectionSchema             AnchorNavSectionSchema          `json:"AnchorNavSectionSchema"`
	StepsSchema                        StepsSchema                     `json:"StepsSchema"`
	StepSchema                         StepSchema                      `json:"StepSchema"`
	StepStatus                         string                          `json:"StepStatus"`
	PortletSchema                      PortletSchema                   `json:"PortletSchema"`
	PortletTabSchema                   PortletTabSchema                `json:"PortletTabSchema"`
	TimelineSchema                     TimelineSchema                  `json:"TimelineSchema"`
	TimelineItemSchema                 TimelineItemSchema              `json:"TimelineItemSchema"`
	FormControlSchema                  FormControlSchema               `json:"FormControlSchema"`
	ArrayControlSchema                 ArrayControlSchema              `json:"ArrayControlSchema"`
	ButtonGroupControlSchema           ButtonGroupControlSchema        `json:"ButtonGroupControlSchema"`
	Option                             Option                          `json:"Option"`
	Options                            []any                           `json:"Options"`
	ChainedSelectControlSchema         ChainedSelectControlSchema      `json:"ChainedSelectControlSchema"`
	CheckboxControlSchema              CheckboxControlSchema           `json:"CheckboxControlSchema"`
	CheckboxesControlSchema            CheckboxesControlSchema         `json:"CheckboxesControlSchema"`
	InputCityControlSchema             InputCityControlSchema          `json:"InputCityControlSchema"`
	InputColorControlSchema            InputColorControlSchema         `json:"InputColorControlSchema"`
	PresetColor                        PresetColor                     `json:"PresetColor"`
	ComboControlSchema                 ComboControlSchema              `json:"ComboControlSchema"`
	ComboCondition                     ComboCondition                  `json:"ComboCondition"`
	ComboSubControl                    ComboSubControl                 `json:"ComboSubControl"`
	ConditionBuilderControlSchema      ConditionBuilderControlSchema   `json:"ConditionBuilderControlSchema"`
	ConditionBuilderFuncs              []any                           `json:"ConditionBuilderFuncs"`
	ConditionFieldFunc                 ConditionFieldFunc              `json:"ConditionFieldFunc"`
	FieldTypes                         string                          `json:"FieldTypes"`
	ConditionBuilderFuncArg            ConditionBuilderFuncArg         `json:"ConditionBuilderFuncArg"`
	ConditionBuilderFields             []any                           `json:"ConditionBuilderFields"`
	ConditionBuilderField              right                           `json:"ConditionBuilderField"`
	FieldSimple                        FieldSimple                     `json:"FieldSimple"`
	CustomField                        CustomField                     `json:"CustomField"`
	FieldGroup                         FieldGroup                      `json:"FieldGroup"`
	ConditionBuilderConfig             ConditionBuilderConfig          `json:"ConditionBuilderConfig"`
	ConditionBuilderType               ConditionBuilderType            `json:"ConditionBuilderType"`
	VariableItem                       VariableItem                    `json:"VariableItem"`
	FuncGroup                          FuncGroup                       `json:"FuncGroup"`
	FuncItem                           FuncItem                        `json:"FuncItem"`
	FormulaPickerInputSettings         FormulaPickerInputSettings      `json:"FormulaPickerInputSettings"`
	FormulaPickerInputSettingType      string                          `json:"FormulaPickerInputSettingType"`
	DateControlSchema                  DateControlSchema               `json:"DateControlSchema"`
	ShortCuts                          ShortCuts                       `json:"ShortCuts"`
	ShortCutDate                       ShortCutDate                    `json:"ShortCutDate"`
	MomentMoment                       momentMoment                    `json:"moment.Moment"`
	ShortCutDateRange                  ShortCutDateRange               `json:"ShortCutDateRange"`
	DateTimeControlSchema              DateTimeControlSchema           `json:"DateTimeControlSchema"`
	TimeControlSchema                  TimeControlSchema               `json:"TimeControlSchema"`
	MonthControlSchema                 MonthControlSchema              `json:"MonthControlSchema"`
	QuarterControlSchema               MonthControlSchema              `json:"QuarterControlSchema"`
	YearControlSchema                  MonthControlSchema              `json:"YearControlSchema"`
	DateRangeControlSchema             DateRangeControlSchema          `json:"DateRangeControlSchema"`
	DiffControlSchema                  DiffControlSchema               `json:"DiffControlSchema"`
	EditorControlSchema                EditorControlSchema             `json:"EditorControlSchema"`
	FieldSetControlSchema              FieldSetControlSchema           `json:"FieldSetControlSchema"`
	FileControlSchema                  FileControlSchema               `json:"FileControlSchema"`
	FormulaControlSchema               FormulaControlSchema            `json:"FormulaControlSchema"`
	GroupControlSchema                 GroupControlSchema              `json:"GroupControlSchema"`
	GroupSubControl                    GroupSubControl                 `json:"GroupSubControl"`
	HiddenControlSchema                HiddenControlSchema             `json:"HiddenControlSchema"`
	IconPickerControlSchema            HiddenControlSchema             `json:"IconPickerControlSchema"`
	ImageControlSchema                 ImageControlSchema              `json:"ImageControlSchema"`
	InputGroupControlSchema            InputGroupControlSchema         `json:"InputGroupControlSchema"`
	ListControlSchema                  ListControlSchema               `json:"ListControlSchema"`
	JSONSchemaEditorControlSchema      JSONSchemaEditorControlSchema   `json:"JSONSchemaEditorControlSchema"`
	SchemaEditorItemPlaceholder        SchemaEditorItemPlaceholder     `json:"SchemaEditorItemPlaceholder"`
	InputSignatureSchema               InputSignatureSchema            `json:"InputSignatureSchema"`
	LocationControlSchema              LocationControlSchema           `json:"LocationControlSchema"`
	UUIDControlSchema                  UUIDControlSchema               `json:"UUIDControlSchema"`
	MatrixControlSchema                MatrixControlSchema             `json:"MatrixControlSchema"`
	MonthRangeControlSchema            MonthRangeControlSchema         `json:"MonthRangeControlSchema"`
	QuarterRangeControlSchema          MonthRangeControlSchema         `json:"QuarterRangeControlSchema"`
	NestedSelectControlSchema          NestedSelectControlSchema       `json:"NestedSelectControlSchema"`
	NumberControlSchema                NumberControlSchema             `json:"NumberControlSchema"`
	PickerControlSchema                PickerControlSchema             `json:"PickerControlSchema"`
	RadiosControlSchema                RadiosControlSchema             `json:"RadiosControlSchema"`
	RadioControlSchema                 RadioControlSchema              `json:"RadioControlSchema"`
	RangeControlSchema                 RangeControlSchema              `json:"RangeControlSchema"`
	Value                              SchemaClassName                 `json:"Value"`
	MultipleValue                      MultipleValue                   `json:"MultipleValue"`
	TooltipPosType                     string                          `json:"TooltipPosType"`
	RatingControlSchema                RatingControlSchema             `json:"RatingControlSchema"`
	TextPositionType                   string                          `json:"textPositionType"`
	RichTextControlSchema              RichTextControlSchema           `json:"RichTextControlSchema"`
	RepeatControlSchema                RepeatControlSchema             `json:"RepeatControlSchema"`
	SelectControlSchema                SelectControlSchema             `json:"SelectControlSchema"`
	SubFormControlSchema               SubFormControlSchema            `json:"SubFormControlSchema"`
	SwitchControlSchema                SwitchControlSchema             `json:"SwitchControlSchema"`
	StaticExactControlSchema           StaticExactControlSchema        `json:"StaticExactControlSchema"`
	TableControlSchema                 TableControlSchema              `json:"TableControlSchema"`
	TabsTransferControlSchema          TabsTransferControlSchema       `json:"TabsTransferControlSchema"`
	TagControlSchema                   TagControlSchema                `json:"TagControlSchema"`
	TextControlSchema                  TextControlSchema               `json:"TextControlSchema"`
	TextareaControlSchema              TextareaControlSchema           `json:"TextareaControlSchema"`
	TransferControlSchema              TransferControlSchema           `json:"TransferControlSchema"`
	TransferPickerControlSchema        TransferPickerControlSchema     `json:"TransferPickerControlSchema"`
	TabsTransferPickerControlSchema    TabsTransferControlSchema       `json:"TabsTransferPickerControlSchema"`
	TreeControlSchema                  TreeControlSchema               `json:"TreeControlSchema"`
	TreeSelectControlSchema            TreeSelectControlSchema         `json:"TreeSelectControlSchema"`
	UserSelectControlSchema            ChainedSelectControlSchema      `json:"UserSelectControlSchema"`
	DateRangeSchema                    DateRangeSchema                 `json:"DateRangeSchema"`
	MultilineTextSchema                MultilineTextSchema             `json:"MultilineTextSchema"`
	PasswordSchema                     PasswordSchema                  `json:"PasswordSchema"`
	WordsSchema                        WordsSchema                     `json:"WordsSchema"`
	TagSchema                          TagSchema                       `json:"TagSchema"`
	SchemaDefaultData                  any                             `json:"SchemaDefaultData"`
	SchemaObjectCardBodyField          SchemaObjectCardBodyField       `json:"SchemaObjectCardBodyField"`
	SchemaObjectListBodyField          SchemaObjectListBodyField       `json:"SchemaObjectListBodyField"`
	SchemaObjectTableColumnWithType    SchemaObjectTableColumnWithType `json:"SchemaObjectTableColumnWithType"`
	SchemaObjectGrid                   SchemaObjectGrid                `json:"SchemaObjectGrid"`
	SchemaObjectHboxRow                SchemaObjectHboxRow             `json:"SchemaObjectHboxRow"`
	SchemaObjectComboSubControl        SchemaObjectComboSubControl     `json:"SchemaObjectComboSubControl"`
	SchemaObjectGroupSubControl        SchemaObjectGroupSubControl     `json:"SchemaObjectGroupSubControl"`
	UnkownSchema                       any                             `json:"UnkownSchema"`
}

type digital_wallet struct {
	Accepts_nfc staticSchema `json:"accepts_nfc"`
}

type dine_in struct {
	Reservations             staticSchema `json:"reservations"`
	Online_reservations      staticSchema `json:"online_reservations"`
	Groups_only_reservations staticSchema `json:"groups_only_reservations"`
	Essential_reservations   staticSchema `json:"essential_reservations"`
}

type distance struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type draggableConfig struct {
	AnyOf       []RootSchema `json:"anyOf"`
	Description string       `json:"description"`
}

type drop_off struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type editorSetting struct {
	Behavior    string `json:"behavior"`
	DisplayName string `json:"displayName"`
	Mock        mock   `json:"mock"`
}

type expandConfig struct {
	Expand    string `json:"expand"`
	ExpandAll bool   `json:"expandAll"`
	Accordion bool   `json:"accordion"`
}

type features struct {
	Payment        payment        `json:"payment"`
	Food_and_drink food_and_drink `json:"food_and_drink"`
	Services       services       `json:"services"`
	Amenities      amenities      `json:"amenities"`
	Attributes     attributes     `json:"attributes"`
}

type filterable struct {
	Source  string `json:"source"`
	Options []any  `json:"options"`
}

type food_and_drink struct {
	Alcohol alcohol `json:"alcohol"`
	Meals   meals   `json:"meals"`
}

type form struct {
	Actions                     []any           `json:"actions"`
	Body                        className       `json:"body"`
	Title                       string          `json:"title"`
	Tabs                        staticSchema    `json:"tabs"`
	FieldSet                    staticSchema    `json:"fieldSet"`
	Data                        staticSchema    `json:"data"`
	Debug                       bool            `json:"debug"`
	DebugConfig                 debugConfig     `json:"debugConfig"`
	InitApi                     SchemaClassName `json:"initApi"`
	InitAsyncApi                SchemaClassName `json:"initAsyncApi"`
	InitFinishedField           string          `json:"initFinishedField"`
	InitCheckInterval           float64         `json:"initCheckInterval"`
	InitFetch                   bool            `json:"initFetch"`
	InitFetchOn                 string          `json:"initFetchOn"`
	Interval                    float64         `json:"interval"`
	SilentPolling               bool            `json:"silentPolling"`
	StopAutoRefreshWhen         string          `json:"stopAutoRefreshWhen"`
	PersistData                 string          `json:"persistData"`
	PersistDataKeys             []string        `json:"persistDataKeys"`
	ClearPersistDataAfterSubmit bool            `json:"clearPersistDataAfterSubmit"`
	Api                         SchemaClassName `json:"api"`
	Feedback                    mock            `json:"feedback"`
	AsyncApi                    SchemaClassName `json:"asyncApi"`
	CheckInterval               float64         `json:"checkInterval"`
	FinishedField               string          `json:"finishedField"`
	ResetAfterSubmit            bool            `json:"resetAfterSubmit"`
	ClearAfterSubmit            bool            `json:"clearAfterSubmit"`
	Mode                        string          `json:"mode"`
	ColumnCount                 float64         `json:"columnCount"`
	Horizontal                  className       `json:"horizontal"`
	AutoFocus                   bool            `json:"autoFocus"`
	Messages                    messages        `json:"messages"`
	Name                        string          `json:"name"`
	PanelClassName              className       `json:"panelClassName"`
	PrimaryField                string          `json:"primaryField"`
	Redirect                    string          `json:"redirect"`
	Reload                      string          `json:"reload"`
	SubmitOnChange              bool            `json:"submitOnChange"`
	SubmitOnInit                bool            `json:"submitOnInit"`
	SubmitText                  string          `json:"submitText"`
	Target                      string          `json:"target"`
	WrapWithPanel               bool            `json:"wrapWithPanel"`
	AffixFooter                 bool            `json:"affixFooter"`
	PromptPageLeave             bool            `json:"promptPageLeave"`
	PromptPageLeaveMessage      string          `json:"promptPageLeaveMessage"`
	Rules                       []ArrayItem3    `json:"rules"`
	PreventEnterSubmit          bool            `json:"preventEnterSubmit"`
	LabelAlign                  className       `json:"labelAlign"`
	LabelWidth                  any             `json:"labelWidth"`
	Static                      bool            `json:"static"`
	StaticOn                    RootSchema      `json:"staticOn"`
	StaticClassName             RootSchema      `json:"staticClassName"`
	Testid                      string          `json:"testid"`
	ID                          string          `json:"$$id"`
	ClassName                   className       `json:"className"`
	Ref                         string          `json:"$ref"`
	Disabled                    bool            `json:"disabled"`
	DisabledOn                  className       `json:"disabledOn"`
	Hidden                      bool            `json:"hidden"`
	HiddenOn                    className       `json:"hiddenOn"`
	Visible                     bool            `json:"visible"`
	VisibleOn                   className       `json:"visibleOn"`
	Id                          string          `json:"id"`
	OnEvent                     any             `json:"onEvent"`
	StaticPlaceholder           string          `json:"staticPlaceholder"`
	StaticLabelClassName        className       `json:"staticLabelClassName"`
	StaticInputClassName        className       `json:"staticInputClassName"`
	StaticSchema                staticSchema    `json:"staticSchema"`
	Style                       any             `json:"style"`
	EditorSetting               editorSetting   `json:"editorSetting"`
	UseMobileUI                 bool            `json:"useMobileUI"`
	TestIdBuilder               RootSchema      `json:"testIdBuilder"`
}

type formula struct {
	EvalMode             bool             `json:"evalMode"`
	MixedMode            bool             `json:"mixedMode"`
	Variables            []any            `json:"variables"`
	VariableMode         string           `json:"variableMode"`
	Functions            []any            `json:"functions"`
	Title                string           `json:"title"`
	Header               string           `json:"header"`
	InputMode            string           `json:"inputMode"`
	AllowInput           bool             `json:"allowInput"`
	Icon                 className        `json:"icon"`
	BtnLabel             string           `json:"btnLabel"`
	Level                string           `json:"level"`
	BtnSize              string           `json:"btnSize"`
	BorderMode           string           `json:"borderMode"`
	Placeholder          string           `json:"placeholder"`
	VariableClassName    string           `json:"variableClassName"`
	FunctionClassName    string           `json:"functionClassName"`
	SelfVariableName     string           `json:"selfVariableName"`
	InputSettings        className        `json:"inputSettings"`
	Remark               className        `json:"remark"`
	LabelRemark          className        `json:"labelRemark"`
	Size                 string           `json:"size"`
	Label                SchemaClassName  `json:"label"`
	LabelAlign           className        `json:"labelAlign"`
	LabelWidth           any              `json:"labelWidth"`
	LabelClassName       string           `json:"labelClassName"`
	Name                 string           `json:"name"`
	ExtraName            string           `json:"extraName"`
	Hint                 string           `json:"hint"`
	SubmitOnChange       bool             `json:"submitOnChange"`
	ReadOnly             bool             `json:"readOnly"`
	ReadOnlyOn           string           `json:"readOnlyOn"`
	ValidateOnChange     bool             `json:"validateOnChange"`
	Description          string           `json:"description"`
	Desc                 string           `json:"desc"`
	DescriptionClassName className        `json:"descriptionClassName"`
	Mode                 string           `json:"mode"`
	Horizontal           className        `json:"horizontal"`
	Inline               bool             `json:"inline"`
	InputClassName       className        `json:"inputClassName"`
	Required             bool             `json:"required"`
	ValidationErrors     validationErrors `json:"validationErrors"`
	Validations          data             `json:"validations"`
	Value                mock             `json:"value"`
	ClearValueOnHidden   bool             `json:"clearValueOnHidden"`
	ValidateApi          SchemaClassName  `json:"validateApi"`
	AutoFill             autoFill         `json:"autoFill"`
	InitAutoFill         data             `json:"initAutoFill"`
	Row                  float64          `json:"row"`
	ID                   string           `json:"$$id"`
	ClassName            className        `json:"className"`
	Ref                  string           `json:"$ref"`
	Disabled             bool             `json:"disabled"`
	DisabledOn           className        `json:"disabledOn"`
	Hidden               bool             `json:"hidden"`
	HiddenOn             className        `json:"hiddenOn"`
	Visible              bool             `json:"visible"`
	VisibleOn            className        `json:"visibleOn"`
	Id                   string           `json:"id"`
	OnEvent              any              `json:"onEvent"`
	Static               bool             `json:"static"`
	StaticOn             className        `json:"staticOn"`
	StaticPlaceholder    string           `json:"staticPlaceholder"`
	StaticClassName      className        `json:"staticClassName"`
	StaticLabelClassName className        `json:"staticLabelClassName"`
	StaticInputClassName className        `json:"staticInputClassName"`
	StaticSchema         staticSchema     `json:"staticSchema"`
	Style                any              `json:"style"`
	EditorSetting        editorSetting    `json:"editorSetting"`
	UseMobileUI          bool             `json:"useMobileUI"`
	TestIdBuilder        RootSchema       `json:"testIdBuilder"`
}

type geocodes struct {
	Drop_off   drop_off `json:"drop_off"`
	Front_door drop_off `json:"front_door"`
	Main       drop_off `json:"main"`
	Road       drop_off `json:"road"`
	Roof       drop_off `json:"roof"`
}

type header struct {
	ClassName              RootSchema   `json:"className"`
	Title                  className    `json:"title"`
	TitleClassName         RootSchema   `json:"titleClassName"`
	SubTitle               className    `json:"subTitle"`
	SubTitleClassName      RootSchema   `json:"subTitleClassName"`
	SubTitlePlaceholder    string       `json:"subTitlePlaceholder"`
	Description            className    `json:"description"`
	DescriptionPlaceholder string       `json:"descriptionPlaceholder"`
	DescriptionClassName   className    `json:"descriptionClassName"`
	Desc                   RootSchema   `json:"desc"`
	DescPlaceholder        RootSchema   `json:"descPlaceholder"`
	DescClassName          RootSchema   `json:"descClassName"`
	Avatar                 className    `json:"avatar"`
	AvatarText             RootSchema   `json:"avatarText"`
	AvatarTextBackground   []ArrayItem1 `json:"avatarTextBackground"`
	AvatarTextClassName    RootSchema   `json:"avatarTextClassName"`
	AvatarClassName        className    `json:"avatarClassName"`
	ImageClassName         className    `json:"imageClassName"`
	Highlight              className    `json:"highlight"`
	HighlightClassName     RootSchema   `json:"highlightClassName"`
	Href                   className    `json:"href"`
	Blank                  bool         `json:"blank"`
}

type hours struct {
	Display          string        `json:"display"`
	Is_local_holiday bool          `json:"is_local_holiday"`
	Open_now         bool          `json:"open_now"`
	Regular          []ArrayElem37 `json:"regular"`
}

type icon struct {
	Id              string   `json:"id"`
	Created_at      string   `json:"created_at"`
	Prefix          string   `json:"prefix"`
	Suffix          string   `json:"suffix"`
	Width           int      `json:"width"`
	Height          int      `json:"height"`
	Classifications []string `json:"classifications"`
	Tip             tip      `json:"tip"`
}

type icons struct {
	Prev RootSchema `json:"prev"`
	Next RootSchema `json:"next"`
}

type ifVal struct {
	TypeVal    string `json:"type"`
	ActionType string `json:"actionType"`
}

type items struct {
	StartTime string       `json:"startTime"`
	EndTime   string       `json:"endTime"`
	Content   staticSchema `json:"content"`
	ClassName string       `json:"className"`
}

type layout struct {
	AnyOf       []ArrayElem2 `json:"anyOf"`
	Description string       `json:"description"`
	DefaultVal  string       `json:"default"`
}

type limit struct {
	AspectRatioLabel string  `json:"aspectRatioLabel"`
	AspectRatio      float64 `json:"aspectRatio"`
	Height           float64 `json:"height"`
	Width            float64 `json:"width"`
	MaxHeight        float64 `json:"maxHeight"`
	MaxWidth         float64 `json:"maxWidth"`
	MinHeight        float64 `json:"minHeight"`
	MinWidth         float64 `json:"minWidth"`
}

type loadingConfig struct {
	Root string `json:"root"`
	Show bool   `json:"show"`
}

type location struct {
	Address           string   `json:"address"`
	Address_extended  string   `json:"address_extended"`
	Admin_region      string   `json:"admin_region"`
	Census_block      string   `json:"census_block"`
	Country           string   `json:"country"`
	Cross_street      string   `json:"cross_street"`
	Dma               string   `json:"dma"`
	Formatted_address string   `json:"formatted_address"`
	Locality          string   `json:"locality"`
	Neighborhood      []string `json:"neighborhood"`
	Po_box            string   `json:"po_box"`
	Post_town         string   `json:"post_town"`
	Postcode          string   `json:"postcode"`
	Region            string   `json:"region"`
}

type meals struct {
	Bar_snacks   staticSchema `json:"bar_snacks"`
	Breakfast    staticSchema `json:"breakfast"`
	Brunch       staticSchema `json:"brunch"`
	Lunch        staticSchema `json:"lunch"`
	Happy_hour   staticSchema `json:"happy_hour"`
	Dessert      staticSchema `json:"dessert"`
	Dinner       staticSchema `json:"dinner"`
	Tasting_menu staticSchema `json:"tasting_menu"`
}

type media struct {
	ClassName RootSchema `json:"className"`
	TypeVal   string     `json:"type"`
	Url       className  `json:"url"`
	Position  string     `json:"position"`
	AutoPlay  bool       `json:"autoPlay"`
	IsLive    bool       `json:"isLive"`
	Poster    className  `json:"poster"`
}

type messages struct {
	ValidateFailed string `json:"validateFailed"`
}

type messages1 struct {
	Success string `json:"success"`
	Failed  string `json:"failed"`
}

type messages2 struct {
	ValidateFailed          string `json:"validateFailed"`
	MinLengthValidateFailed string `json:"minLengthValidateFailed"`
	MaxLengthValidateFailed string `json:"maxLengthValidateFailed"`
}

type mock struct {
	Description string `json:"description"`
}

type momentMoment struct {
	Constructor Field_161becb8 `json:"constructor"`
}

type multiple struct {
	Count float64 `json:"count"`
}

type offset struct {
	Top  float64 `json:"top"`
	Left float64 `json:"left"`
}

type options struct {
	AnyOf       []ArrayElem22 `json:"anyOf"`
	Description string        `json:"description"`
}

type overflowConfig struct {
	MaxTagCount              float64   `json:"maxTagCount"`
	DisplayPosition          []string  `json:"displayPosition"`
	OverflowTagPopover       className `json:"overflowTagPopover"`
	OverflowTagPopoverInCRUD className `json:"overflowTagPopoverInCRUD"`
}

type overflowIndicator struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	DefaultVal  string `json:"default"`
}

type overlay struct {
	Width        any    `json:"width"`
	Align        string `json:"align"`
	FilterOption string `json:"filterOption"`
}

type pagination struct {
	Layout                   layout    `json:"layout"`
	MaxButtons               float64   `json:"maxButtons"`
	PerPageAvailable         []float64 `json:"perPageAvailable"`
	PopOverContainerSelector string    `json:"popOverContainerSelector"`
	Enable                   className `json:"enable"`
	ClassName                className `json:"className"`
	LoadDataOnce             bool      `json:"loadDataOnce"`
}

type parking struct {
	Parking        staticSchema `json:"parking"`
	Street_parking staticSchema `json:"street_parking"`
	Valet_parking  staticSchema `json:"valet_parking"`
	Public_lot     staticSchema `json:"public_lot"`
	Private_lot    staticSchema `json:"private_lot"`
}

type parsePrimitiveQuery struct {
	AnyOf       []ArrayElem13 `json:"anyOf"`
	Description string        `json:"description"`
}

type patternProperties struct {
	PatternProperties any `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|size|label|labelAlign|labelWidth|labelClassName|name|extraName|remark|labelRemark|hint|submitOnChange|readOnly|readOnlyOn|validateOnChange|description|desc|descriptionClassName|mode|horizontal|inline|inputClassName|placeholder|required|validationErrors|validations|value|clearValueOnHidden|validateApi|autoFill|initAutoFill|row|type|extractValue|joinValues|delimiter|allowCity|allowDistrict|allowStreet|searchable|itemClassName|columnClassName|columnRatio)$"`
}

type patternProperties1 struct {
	PatternProperties staticSchema `json:"^(label|labelClassName|name|popOver|quickEdit|copyable)$"`
}

type patternProperties10 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|label|labelClassName|innerClassName|popOver|quickEdit|copyable)$"`
}

type patternProperties11 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|label|fixed|popOver|quickEdit|quickEditOnUpdate|copyable|sortable|searchable|toggled|width|align|vAlign|headerAlign|classNameExpr|labelClassName|filterable|breakpoint|value|unique|canAccessSuperData|lazyRenderAfter|innerStyle)$"`
}

type patternProperties12 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|x|y|w|h|width|height|align|valign|gridClassName)$"`
}

type patternProperties13 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|rowClassName|cellClassName)$"`
}

type patternProperties14 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|unique|columnClassName)$"`
}

type patternProperties15 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|columnClassName|columnRatio)$"`
}

type patternProperties2 struct {
	PatternProperties staticSchema `json:"^(saveImmediately|resetOnFailed|reload|mode|icon)$"`
}

type patternProperties3 struct {
	PatternProperties staticSchema `json:"^(label|labelClassName|innerClassName|name|popOver|quickEdit|copyable)$"`
}

type patternProperties4 struct {
	PatternProperties staticSchema `json:"^(label|fixed|name|popOver|quickEdit|quickEditOnUpdate|copyable|sortable|searchable|toggled|width|align|vAlign|headerAlign|className|classNameExpr|labelClassName|filterable|breakpoint|remark|value|unique|canAccessSuperData|lazyRenderAfter|innerStyle)$"`
}

type patternProperties5 struct {
	PatternProperties staticSchema `json:"^(x|y|w|h|width|height|align|valign|gridClassName)$"`
}

type patternProperties6 struct {
	PatternProperties staticSchema `json:"^(rowClassName|cellClassName)$"`
}

type patternProperties7 struct {
	PatternProperties staticSchema `json:"^(unique|columnClassName|testid)$"`
}

type patternProperties8 struct {
	PatternProperties staticSchema `json:"^(columnClassName|columnRatio|name)$"`
}

type patternProperties9 struct {
	PatternProperties staticSchema `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|type|testid|title|subTitle|remark|body|bodyClassName|aside|asideResizor|asideSticky|asideMinWidth|asideMaxWidth|asideClassName|css|mobileCSS|data|headerClassName|initApi|initFetch|initFetchOn|messages|name|toolbar|toolbarClassName|definitions|interval|silentPolling|stopAutoRefreshWhen|showErrorMsg|cssVars|regions|pullRefresh|label|labelClassName|popOver|quickEdit|copyable)$"`
}

type payment struct {
	Credit_cards   credit_cards   `json:"credit_cards"`
	Digital_wallet digital_wallet `json:"digital_wallet"`
}

type pullRefresh struct {
	Disabled    bool   `json:"disabled"`
	PullingText string `json:"pullingText"`
	LoosingText string `json:"loosingText"`
}

type qsOptions struct {
	ArrayFormat string `json:"arrayFormat"`
	Indices     bool   `json:"indices"`
	AllowDots   bool   `json:"allowDots"`
}

type relativeLocation struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Distance distance `json:"distance"`
	Bearing  bearing  `json:"bearing"`
}

type report struct {
	Asset_report_id  string        `json:"asset_report_id"`
	Client_report_id string        `json:"client_report_id"`
	Date_generated   string        `json:"date_generated"`
	Days_requested   int           `json:"days_requested"`
	Items            []ArrayElem44 `json:"items"`
	User             user          `json:"user"`
}

type right struct {
	AnyOf []RootSchema `json:"anyOf"`
}

type scaffold struct {
	Description string       `json:"description"`
	DefaultVal  staticSchema `json:"default"`
}

type schedules struct {
	AnyOf       []ArrayElem6 `json:"anyOf"`
	Description string       `json:"description"`
}

type searchConfig struct {
	ClassName        string          `json:"className"`
	MatchFunc        SchemaClassName `json:"matchFunc"`
	Placeholder      string          `json:"placeholder"`
	Mini             bool            `json:"mini"`
	Enhance          bool            `json:"enhance"`
	Clearable        bool            `json:"clearable"`
	SearchImediately bool            `json:"searchImediately"`
	ValueField       string          `json:"valueField"`
}

type searchConfig1 struct {
	ClassName        string `json:"className"`
	Placeholder      string `json:"placeholder"`
	Mini             bool   `json:"mini"`
	Enhance          bool   `json:"enhance"`
	Clearable        bool   `json:"clearable"`
	SearchImediately bool   `json:"searchImediately"`
	Sticky           bool   `json:"sticky"`
}

type services struct {
	Delivery      staticSchema `json:"delivery"`
	Takeout       staticSchema `json:"takeout"`
	Drive_through staticSchema `json:"drive_through"`
	Dine_in       dine_in      `json:"dine_in"`
}

type social_media struct {
	Facebook_id string `json:"facebook_id"`
	Instagram   string `json:"instagram"`
	Twitter     string `json:"twitter"`
}

type stateTextMap struct {
	Init      string `json:"init"`
	Pending   string `json:"pending"`
	Uploading string `json:"uploading"`
	Error     string `json:"error"`
	Uploaded  string `json:"uploaded"`
	Ready     string `json:"ready"`
}

type staticSchema struct {
}

type stats struct {
	Total_photos  int `json:"total_photos"`
	Total_ratings int `json:"total_ratings"`
	Total_tips    int `json:"total_tips"`
}

type threshold struct {
	AnyOf       []ArrayElem18 `json:"anyOf"`
	Description string        `json:"description"`
}

type tip struct {
	Id             string `json:"id"`
	Created_at     string `json:"created_at"`
	Text           string `json:"text"`
	Url            string `json:"url"`
	Lang           string `json:"lang"`
	Agree_count    int    `json:"agree_count"`
	Disagree_count int    `json:"disagree_count"`
}

type trackConfig struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type transform struct {
	LowerCase bool `json:"lowerCase"`
	UpperCase bool `json:"upperCase"`
}

type user struct {
	Client_user_id string `json:"client_user_id"`
	Email          string `json:"email"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Middle_name    string `json:"middle_name"`
	Phone_number   string `json:"phone_number"`
	Ssn            string `json:"ssn"`
}

type validationConfig struct {
	ErrorMode string `json:"errorMode"`
	Delimiter string `json:"delimiter"`
}

type validationErrors struct {
	IsAlpha                string `json:"isAlpha"`
	IsAlphanumeric         string `json:"isAlphanumeric"`
	IsEmail                string `json:"isEmail"`
	IsFloat                string `json:"isFloat"`
	IsInt                  string `json:"isInt"`
	IsJson                 string `json:"isJson"`
	IsLength               string `json:"isLength"`
	IsNumeric              string `json:"isNumeric"`
	IsRequired             string `json:"isRequired"`
	IsUrl                  string `json:"isUrl"`
	MatchRegexp            string `json:"matchRegexp"`
	MatchRegexp2           string `json:"matchRegexp2"`
	MatchRegexp3           string `json:"matchRegexp3"`
	MatchRegexp4           string `json:"matchRegexp4"`
	MatchRegexp5           string `json:"matchRegexp5"`
	MaxLength              string `json:"maxLength"`
	Maximum                string `json:"maximum"`
	MinLength              string `json:"minLength"`
	Minimum                string `json:"minimum"`
	IsDateTimeSame         string `json:"isDateTimeSame"`
	IsDateTimeBefore       string `json:"isDateTimeBefore"`
	IsDateTimeAfter        string `json:"isDateTimeAfter"`
	IsDateTimeSameOrBefore string `json:"isDateTimeSameOrBefore"`
	IsDateTimeSameOrAfter  string `json:"isDateTimeSameOrAfter"`
	IsDateTimeBetween      string `json:"isDateTimeBetween"`
	IsTimeSame             string `json:"isTimeSame"`
	IsTimeBefore           string `json:"isTimeBefore"`
	IsTimeAfter            string `json:"isTimeAfter"`
	IsTimeSameOrBefore     string `json:"isTimeSameOrBefore"`
	IsTimeSameOrAfter      string `json:"isTimeSameOrAfter"`
	IsTimeBetween          string `json:"isTimeBetween"`
}

type RootContainer struct {
	Data *RootConfig6 `json:"data"`
}

func main() {
	data := RootConfig6{
		Name:              "GoLand",
		Version:           "2025.1.1",
		BuildNumber:       "251.25410.140",
		ProductCode:       "GO",
		EnvVarBaseName:    "GOLAND",
		DataDirectoryName: "GoLand2025.1",
		SvgIconPath:       "bin/goland.svg",
		ProductVendor:     "JetBrains",
		Launch: []ArrayElem46{
			ArrayElem46{
				Os:                 "Windows",
				Arch:               "amd64",
				LauncherPath:       "bin/goland64.exe",
				JavaExecutablePath: "jbr/bin/java.exe",
				VmOptionsFilePath:  "bin/goland64.exe.vmoptions",
				BootClassPathJarNames: []string{
					"platform-loader.jar",
					"util-8.jar",
					"util.jar",
					"app-client.jar",
					"util_rt.jar",
					"product.jar",
					"lib-client.jar",
					"trove.jar",
					"app.jar",
					"opentelemetry.jar",
					"jps-model.jar",
					"stats.jar",
					"rd.jar",
					"external-system-rt.jar",
					"protobuf.jar",
					"bouncy-castle.jar",
					"forms_rt.jar",
					"lib.jar",
					"externalProcess-rt.jar",
					"groovy.jar",
					"annotations.jar",
					"jsch-agent.jar",
					"kotlinx-coroutines-slf4j-1.8.0-intellij.jar",
					"product-client.jar",
				},
				AdditionalJvmArguments: []string{
					"-Xbootclasspath/a:%IDE_HOME%/lib/nio-fs.jar",
					"-Djava.system.class.loader=com.intellij.util.lang.PathClassLoader",
					"-Didea.vendor.name=JetBrains",
					"-Didea.paths.selector=GoLand2025.1",
					"-Djna.boot.library.path=%IDE_HOME%/lib/jna/amd64",
					"-Djna.nosys=true",
					"-Djna.noclasspath=true",
					"-Dpty4j.preferred.native.folder=%IDE_HOME%/lib/pty4j",
					"-Dio.netty.allocator.type=pooled",
					"-Dintellij.platform.runtime.repository.path=%IDE_HOME%/modules/module-descriptors.jar",
					"-Didea.platform.prefix=GoLand",
					"-Dwsl.use.remote.agent.for.nio.filesystem=true",
					"-Djava.nio.file.spi.DefaultFileSystemProvider=com.intellij.platform.core.nio.fs.MultiRoutingFileSystemProvider",
					"-Dsplash=true",
					"-Daether.connector.resumeDownloads=false",
					"-Dcompose.swing.render.on.graphics=true",
					"--add-opens=java.base/java.io=ALL-UNNAMED",
					"--add-opens=java.base/java.lang=ALL-UNNAMED",
					"--add-opens=java.base/java.lang.ref=ALL-UNNAMED",
					"--add-opens=java.base/java.lang.reflect=ALL-UNNAMED",
					"--add-opens=java.base/java.net=ALL-UNNAMED",
					"--add-opens=java.base/java.nio=ALL-UNNAMED",
					"--add-opens=java.base/java.nio.charset=ALL-UNNAMED",
					"--add-opens=java.base/java.text=ALL-UNNAMED",
					"--add-opens=java.base/java.time=ALL-UNNAMED",
					"--add-opens=java.base/java.util=ALL-UNNAMED",
					"--add-opens=java.base/java.util.concurrent=ALL-UNNAMED",
					"--add-opens=java.base/java.util.concurrent.atomic=ALL-UNNAMED",
					"--add-opens=java.base/java.util.concurrent.locks=ALL-UNNAMED",
					"--add-opens=java.base/jdk.internal.vm=ALL-UNNAMED",
					"--add-opens=java.base/sun.net.dns=ALL-UNNAMED",
					"--add-opens=java.base/sun.nio.ch=ALL-UNNAMED",
					"--add-opens=java.base/sun.nio.fs=ALL-UNNAMED",
					"--add-opens=java.base/sun.security.ssl=ALL-UNNAMED",
					"--add-opens=java.base/sun.security.util=ALL-UNNAMED",
					"--add-opens=java.desktop/com.sun.java.swing=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt.dnd.peer=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt.event=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt.font=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt.image=ALL-UNNAMED",
					"--add-opens=java.desktop/java.awt.peer=ALL-UNNAMED",
					"--add-opens=java.desktop/javax.swing=ALL-UNNAMED",
					"--add-opens=java.desktop/javax.swing.plaf.basic=ALL-UNNAMED",
					"--add-opens=java.desktop/javax.swing.text=ALL-UNNAMED",
					"--add-opens=java.desktop/javax.swing.text.html=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.awt=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.awt.datatransfer=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.awt.image=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.awt.windows=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.font=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.java2d=ALL-UNNAMED",
					"--add-opens=java.desktop/sun.swing=ALL-UNNAMED",
					"--add-opens=java.management/sun.management=ALL-UNNAMED",
					"--add-opens=jdk.attach/sun.tools.attach=ALL-UNNAMED",
					"--add-opens=jdk.compiler/com.sun.tools.javac.api=ALL-UNNAMED",
					"--add-opens=jdk.internal.jvmstat/sun.jvmstat.monitor=ALL-UNNAMED",
					"--add-opens=jdk.jdi/com.sun.tools.jdi=ALL-UNNAMED",
				},
				MainClass: "com.intellij.idea.Main",
				CustomCommands: []ArrayElem45{
					ArrayElem45{
						Commands: []string{
							"thinClient",
							"thinClient-headless",
							"installFrontendPlugins",
						},
						VmOptionsFilePath: "bin/jetbrains_client64.exe.vmoptions",
						BootClassPathJarNames: []string{
							"platform-loader.jar",
						},
						AdditionalJvmArguments: []string{
							"-Djava.system.class.loader=com.intellij.util.lang.PathClassLoader",
							"-Didea.vendor.name=JetBrains",
							"-Didea.paths.selector=GoLand2025.1",
							"-Djna.boot.library.path=%IDE_HOME%/lib/jna/amd64",
							"-Djna.nosys=true",
							"-Djna.noclasspath=true",
							"-Dpty4j.preferred.native.folder=%IDE_HOME%/lib/pty4j",
							"-Dio.netty.allocator.type=pooled",
							"-Dintellij.platform.runtime.repository.path=%IDE_HOME%/modules/module-descriptors.jar",
							"-Dintellij.platform.root.module=intellij.goland.frontend.split",
							"-Dintellij.platform.product.mode=frontend",
							"-Didea.platform.prefix=JetBrainsClient",
							"-Dwsl.use.remote.agent.for.nio.filesystem=false",
							"-Dide.no.platform.update=true",
							"-Didea.initially.ask.config=never",
							"-Dnosplash=true",
							"-Didea.paths.customizer=com.intellij.platform.ide.impl.startup.multiProcess.FrontendProcessPathCustomizer",
							"-Dintellij.platform.full.ide.product.code=GO",
							"-Dsplash=true",
							"-Daether.connector.resumeDownloads=false",
							"-Dcompose.swing.render.on.graphics=true",
							"--add-opens=java.base/java.io=ALL-UNNAMED",
							"--add-opens=java.base/java.lang=ALL-UNNAMED",
							"--add-opens=java.base/java.lang.ref=ALL-UNNAMED",
							"--add-opens=java.base/java.lang.reflect=ALL-UNNAMED",
							"--add-opens=java.base/java.net=ALL-UNNAMED",
							"--add-opens=java.base/java.nio=ALL-UNNAMED",
							"--add-opens=java.base/java.nio.charset=ALL-UNNAMED",
							"--add-opens=java.base/java.text=ALL-UNNAMED",
							"--add-opens=java.base/java.time=ALL-UNNAMED",
							"--add-opens=java.base/java.util=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent.atomic=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent.locks=ALL-UNNAMED",
							"--add-opens=java.base/jdk.internal.vm=ALL-UNNAMED",
							"--add-opens=java.base/sun.net.dns=ALL-UNNAMED",
							"--add-opens=java.base/sun.nio.ch=ALL-UNNAMED",
							"--add-opens=java.base/sun.nio.fs=ALL-UNNAMED",
							"--add-opens=java.base/sun.security.ssl=ALL-UNNAMED",
							"--add-opens=java.base/sun.security.util=ALL-UNNAMED",
							"--add-opens=java.desktop/com.sun.java.swing=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.dnd.peer=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.event=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.font=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.image=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.peer=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.plaf.basic=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.text=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.text.html=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.datatransfer=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.image=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.windows=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.font=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.java2d=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.swing=ALL-UNNAMED",
							"--add-opens=java.management/sun.management=ALL-UNNAMED",
							"--add-opens=jdk.attach/sun.tools.attach=ALL-UNNAMED",
							"--add-opens=jdk.compiler/com.sun.tools.javac.api=ALL-UNNAMED",
							"--add-opens=jdk.internal.jvmstat/sun.jvmstat.monitor=ALL-UNNAMED",
							"--add-opens=jdk.jdi/com.sun.tools.jdi=ALL-UNNAMED",
							"-Dintellij.platform.load.app.info.from.resources=true",
						},
						MainClass:         "com.intellij.platform.runtime.loader.IntellijLoader",
						EnvVarBaseName:    "JETBRAINS_CLIENT",
						DataDirectoryName: "GoLand2025.1",
					},
					ArrayElem45{
						Commands: []string{
							"qodana",
						},
						BootClassPathJarNames: []string{
							"platform-loader.jar",
							"util-8.jar",
							"util.jar",
							"app-client.jar",
							"util_rt.jar",
							"product.jar",
							"lib-client.jar",
							"trove.jar",
							"app.jar",
							"opentelemetry.jar",
							"jps-model.jar",
							"stats.jar",
							"rd.jar",
							"external-system-rt.jar",
							"protobuf.jar",
							"bouncy-castle.jar",
							"forms_rt.jar",
							"lib.jar",
							"externalProcess-rt.jar",
							"groovy.jar",
							"annotations.jar",
							"jsch-agent.jar",
							"kotlinx-coroutines-slf4j-1.8.0-intellij.jar",
							"product-client.jar",
							"nio-fs.jar",
						},
						AdditionalJvmArguments: []string{
							"-Djava.system.class.loader=com.intellij.util.lang.PathClassLoader",
							"-Didea.vendor.name=JetBrains",
							"-Didea.paths.selector=GoLand2025.1",
							"-Djna.boot.library.path=%IDE_HOME%/lib/jna/amd64",
							"-Djna.nosys=true",
							"-Djna.noclasspath=true",
							"-Dpty4j.preferred.native.folder=%IDE_HOME%/lib/pty4j",
							"-Dio.netty.allocator.type=pooled",
							"-Dintellij.platform.runtime.repository.path=%IDE_HOME%/modules/module-descriptors.jar",
							"-Didea.platform.prefix=GoLand",
							"-Dwsl.use.remote.agent.for.nio.filesystem=false",
							"-Dsplash=true",
							"-Daether.connector.resumeDownloads=false",
							"-Dcompose.swing.render.on.graphics=true",
							"--add-opens=java.base/java.io=ALL-UNNAMED",
							"--add-opens=java.base/java.lang=ALL-UNNAMED",
							"--add-opens=java.base/java.lang.ref=ALL-UNNAMED",
							"--add-opens=java.base/java.lang.reflect=ALL-UNNAMED",
							"--add-opens=java.base/java.net=ALL-UNNAMED",
							"--add-opens=java.base/java.nio=ALL-UNNAMED",
							"--add-opens=java.base/java.nio.charset=ALL-UNNAMED",
							"--add-opens=java.base/java.text=ALL-UNNAMED",
							"--add-opens=java.base/java.time=ALL-UNNAMED",
							"--add-opens=java.base/java.util=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent.atomic=ALL-UNNAMED",
							"--add-opens=java.base/java.util.concurrent.locks=ALL-UNNAMED",
							"--add-opens=java.base/jdk.internal.vm=ALL-UNNAMED",
							"--add-opens=java.base/sun.net.dns=ALL-UNNAMED",
							"--add-opens=java.base/sun.nio.ch=ALL-UNNAMED",
							"--add-opens=java.base/sun.nio.fs=ALL-UNNAMED",
							"--add-opens=java.base/sun.security.ssl=ALL-UNNAMED",
							"--add-opens=java.base/sun.security.util=ALL-UNNAMED",
							"--add-opens=java.desktop/com.sun.java.swing=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.dnd.peer=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.event=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.font=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.image=ALL-UNNAMED",
							"--add-opens=java.desktop/java.awt.peer=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.plaf.basic=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.text=ALL-UNNAMED",
							"--add-opens=java.desktop/javax.swing.text.html=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.datatransfer=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.image=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.awt.windows=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.font=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.java2d=ALL-UNNAMED",
							"--add-opens=java.desktop/sun.swing=ALL-UNNAMED",
							"--add-opens=java.management/sun.management=ALL-UNNAMED",
							"--add-opens=jdk.attach/sun.tools.attach=ALL-UNNAMED",
							"--add-opens=jdk.compiler/com.sun.tools.javac.api=ALL-UNNAMED",
							"--add-opens=jdk.internal.jvmstat/sun.jvmstat.monitor=ALL-UNNAMED",
							"--add-opens=jdk.jdi/com.sun.tools.jdi=ALL-UNNAMED",
							"-Dqodana.application=true",
							"-Dintellij.platform.load.app.info.from.resources=true",
							"-Dfus.internal.reduce.initial.delay=true",
							"-Didea.headless.statistics.max.files.to.send=5000",
							"-Djava.awt.headless=true",
							"-Didea.qodana.thirdpartyplugins.accept=true",
							"-Dide.warmup.use.predicates=false",
							"-Dvcs.log.index.enable=false",
							"-Didea.job.launcher.without.timeout=true",
							"-Dscanning.in.smart.mode=false",
							"-Deap.login.enabled=false",
							"-Dsdk.download.consent=true",
							"-Dide.activity.tracking.enable.debug=true",
							"-Deslint.service.expiration.timeout.ms=5000",
							"-Dphp.additional.library.manager.new.updater=true",
							"-Dgradle.force.project.sdk.import=true",
							"-Dide.region.url.mapping.expiration.timeout=7200",
							"-Dqodana.product.name=Qodana for Go",
							"-Dqodana.build.number=QDGO-251.25410.140",
							"-Dqodana.eap=false",
						},
					},
				},
			},
		},
		CustomProperties: []ArrayElem47{
			ArrayElem47{
				Key:   "source.git.revision",
				Value: "1f3176f2429c2",
			},
		},
		BundledPlugins: []string{
			"Docker",
			"Git4Idea",
			"HtmlTools",
			"JSIntentionPowerPack",
			"JavaScript",
			"JavaScriptDebugger",
			"Refactor-X",
			"Remote Development Server",
			"com.intellij",
			"com.intellij.completion.ml.ranking",
			"com.intellij.configurationScript",
			"com.intellij.copyright",
			"com.intellij.css",
			"com.intellij.database",
			"com.intellij.dev",
			"com.intellij.diagram",
			"com.intellij.grpc",
			"com.intellij.ja",
			"com.intellij.jsonpath",
			"com.intellij.ko",
			"com.intellij.kubernetes",
			"com.intellij.microservices.ui",
			"com.intellij.modules.json",
			"com.intellij.platform.images",
			"com.intellij.plugins.colorful.darcula.colorscheme",
			"com.intellij.plugins.goland.solarized.colorscheme",
			"com.intellij.plugins.monokai.colorscheme",
			"com.intellij.plugins.vscodekeymap",
			"com.intellij.plugins.watcher",
			"com.intellij.react",
			"com.intellij.searcheverywhere.ml",
			"com.intellij.settingsSync",
			"com.intellij.stylelint",
			"com.intellij.swagger",
			"com.intellij.zh",
			"com.jetbrains.codeWithMe",
			"com.jetbrains.gateway",
			"com.jetbrains.gateway.terminal",
			"com.jetbrains.performancePlugin",
			"com.jetbrains.performancePlugin.async",
			"com.jetbrains.plugins.webDeployment",
			"com.jetbrains.restClient",
			"com.jetbrains.sh",
			"com.jetbrains.station",
			"com.ypwang.plugin.go-linter",
			"hg4idea",
			"idea.plugin.protoeditor",
			"intellij.charts",
			"intellij.git.commit.modal",
			"intellij.grid.plugin",
			"intellij.indexing.shared",
			"intellij.indexing.shared.core",
			"intellij.platform.ijent.impl",
			"intellij.webpack",
			"name.kropp.intellij.makefile",
			"org.editorconfig.editorconfigjetbrains",
			"org.intellij.intelliLang",
			"org.intellij.plugins.markdown",
			"org.intellij.qodana",
			"org.jetbrains.completion.full.line",
			"org.jetbrains.plugins.docker.gateway",
			"org.jetbrains.plugins.emojipicker",
			"org.jetbrains.plugins.github",
			"org.jetbrains.plugins.gitlab",
			"org.jetbrains.plugins.go",
			"org.jetbrains.plugins.go-template",
			"org.jetbrains.plugins.go.ide",
			"org.jetbrains.plugins.go.sharedIndexes.bundled",
			"org.jetbrains.plugins.remote-run",
			"org.jetbrains.plugins.terminal",
			"org.jetbrains.plugins.textmate",
			"org.jetbrains.plugins.yaml",
			"org.jetbrains.security.package-checker",
			"tanvd.grazi",
			"training",
		},
		Modules: []string{
			"com.intellij.marketplace",
			"com.intellij.ml.inline.completion",
			"com.intellij.modules.coverage",
			"com.intellij.modules.database",
			"com.intellij.modules.database.connectivity",
			"com.intellij.modules.database.core",
			"com.intellij.modules.duplicatesDetector",
			"com.intellij.modules.externalSystem",
			"com.intellij.modules.go",
			"com.intellij.modules.go-capable",
			"com.intellij.modules.goland",
			"com.intellij.modules.javascript",
			"com.intellij.modules.lang",
			"com.intellij.modules.microservices",
			"com.intellij.modules.nativeDebug-plugin-capable",
			"com.intellij.modules.os.linux",
			"com.intellij.modules.os.unix",
			"com.intellij.modules.os.xwindow",
			"com.intellij.modules.platform",
			"com.intellij.modules.profiler",
			"com.intellij.modules.python-core-capable",
			"com.intellij.modules.python-in-mini-ide-capable",
			"com.intellij.modules.python-in-non-pycharm-ide-capable",
			"com.intellij.modules.remoteServers",
			"com.intellij.modules.run.targets",
			"com.intellij.modules.sql",
			"com.intellij.modules.ssh",
			"com.intellij.modules.ssh.ui",
			"com.intellij.modules.structuralsearch",
			"com.intellij.modules.ultimate",
			"com.intellij.modules.vcs",
			"com.intellij.modules.xdebugger",
			"com.intellij.modules.xml",
			"com.intellij.platform.experimental.backend",
			"com.intellij.platform.experimental.frontend",
			"com.intellij.platform.experimental.monolith",
			"com.intellij.platform.ide.provisioner",
			"com.intellij.station.core",
		},
		FileExtensions: []string{
			"*.ane",
			"*.ant",
			"*.apk",
			"*.ats",
			"*.bash",
			"*.bigtiff",
			"*.bmp",
			"*.cgo",
			"*.cjs",
			"*.cql",
			"*.css",
			"*.csv",
			"*.cts",
			"*.dcx",
			"*.ddl",
			"*.dic",
			"*.diff",
			"*.doc",
			"*.dockerfile",
			"*.dockerignore",
			"*.docx",
			"*.dtd",
			"*.ear",
			"*.egg",
			"*.elt",
			"*.ent",
			"*.es6",
			"*.fnc",
			"*.fxml",
			"*.ghe",
			"*.gif",
			"*.gitignore",
			"*.go",
			"*.go2",
			"*.gohtml",
			"*.har",
			"*.hgignore",
			"*.hlp",
			"*.hql",
			"*.htm",
			"*.html",
			"*.http",
			"*.icns",
			"*.ico",
			"*.ignore",
			"*.iml",
			"*.ipr",
			"*.iws",
			"*.jar",
			"*.jbig2",
			"*.jhm",
			"*.jnlp",
			"*.jpeg",
			"*.jpg",
			"*.jrxml",
			"*.js",
			"*.js.flow",
			"*.jsb2",
			"*.jsb3",
			"*.jshintrc",
			"*.json",
			"*.json5",
			"*.jsonc",
			"*.jsonl",
			"*.jsonlines",
			"*.jsx",
			"*.kubernetes-helm-json",
			"*.kubernetes-helm-txt",
			"*.kubernetes-helm-yaml",
			"*.kubernetes-werf-yaml",
			"*.ldjson",
			"*.map",
			"*.markdown",
			"*.md",
			"*.mdb",
			"*.mjs",
			"*.mk",
			"*.mod",
			"*.mts",
			"*.ndjson",
			"*.odt",
			"*.pam",
			"*.patch",
			"*.pb",
			"*.pbm",
			"*.pck",
			"*.pcx",
			"*.pdf",
			"*.pgm",
			"*.pkb",
			"*.pks",
			"*.pls",
			"*.png",
			"*.pnm",
			"*.postman_collection",
			"*.postman_collection.json",
			"*.postman_environment",
			"*.postman_environment.json",
			"*.ppm",
			"*.ppt",
			"*.pptx",
			"*.prc",
			"*.proto",
			"*.psd",
			"*.redis",
			"*.regexp",
			"*.rest",
			"*.rgbe",
			"*.rnc",
			"*.rng",
			"*.s",
			"*.sarif",
			"*.sarif*.json",
			"*.sarif.json",
			"*.sh",
			"*.sht",
			"*.shtm",
			"*.shtml",
			"*.snap",
			"*.sql",
			"*.svg",
			"*.swc",
			"*.textproto",
			"*.tga",
			"*.tif",
			"*.tiff",
			"*.tld",
			"*.tpb",
			"*.tps",
			"*.trg",
			"*.ts",
			"*.tsv",
			"*.tsx",
			"*.txtpb",
			"*.uml",
			"*.vsd",
			"*.vw",
			"*.war",
			"*.wbmp",
			"*.wsdl",
			"*.xbm",
			"*.xhtml",
			"*.xml",
			"*.xpm",
			"*.xsd",
			"*.xsl",
			"*.xslt",
			"*.xul",
			"*.yaml",
			"*.yml",
			"*.zip",
			"*.zsh",
			".babelrc",
			".devcontainer.json",
			".editorconfig",
			".eslintrc",
			".stylelintrc",
			".ws-context",
			"Containerfile",
			"Dockerfile",
			"Dockerfile.fast-jar",
			"Dockerfile.jlink",
			"Dockerfile.jvm",
			"Dockerfile.legacy-jar",
			"Dockerfile.native",
			"Dockerfile.native-distroless",
			"Dockerfile.native-micro",
			"GNUmakefile",
			"GNUmakefile.inc",
			"devcontainer.json",
			"exclude",
			"go.mod",
			"go.work",
			"http-client.cookies",
			"jest.config",
			"jshint.json",
			"makefile",
			"makefile.inc",
			"tsconfig.app.json",
			"tsconfig.base.json",
			"tsconfig.e2e.json",
			"tsconfig.editor.json",
			"tsconfig.json",
			"tsconfig.lib.json",
			"tsconfig.spec.json",
			"tsconfig.test.json",
			"yarn.lock",
		},
		Layout: []ArrayElem48{
			ArrayElem48{
				Name: "Docker",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/clouds-docker-impl/lib/clouds-docker-impl.jar",
				},
			},
			ArrayElem48{
				Name: "Git4Idea",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/vcs-git/lib/vcs-git.jar",
					"plugins/vcs-git/lib/git4idea-rt.jar",
				},
			},
			ArrayElem48{
				Name: "HtmlTools",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/html-tools/lib/html-tools.jar",
				},
			},
			ArrayElem48{
				Name: "JSIntentionPowerPack",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/javascript-intentions/lib/javascript-intentions.jar",
				},
			},
			ArrayElem48{
				Name: "JavaScript",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/javascript-plugin/lib/javascript-plugin.jar",
					"plugins/javascript-plugin/lib/javascript-frontback.jar",
				},
			},
			ArrayElem48{
				Name: "JavaScriptDebugger",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/javascript-debugger/lib/javascript-debugger.jar",
					"plugins/javascript-debugger/lib/ChromeConnector.jar",
				},
			},
			ArrayElem48{
				Name: "Refactor-X",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/xml-refactoring/lib/xml-refactoring.jar",
				},
			},
			ArrayElem48{
				Name: "Remote Development Server",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/remote-dev-server/lib/remote-dev-server.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij",
				Kind: "plugin",
				ClassPath: []string{
					"lib/platform-loader.jar",
					"lib/util-8.jar",
					"lib/util.jar",
					"lib/app-client.jar",
					"lib/util_rt.jar",
					"lib/product.jar",
					"lib/lib-client.jar",
					"lib/trove.jar",
					"lib/app.jar",
					"lib/opentelemetry.jar",
					"lib/jps-model.jar",
					"lib/stats.jar",
					"lib/rd.jar",
					"lib/external-system-rt.jar",
					"lib/protobuf.jar",
					"lib/bouncy-castle.jar",
					"lib/nio-fs.jar",
					"lib/forms_rt.jar",
					"lib/lib.jar",
					"lib/externalProcess-rt.jar",
					"lib/groovy.jar",
					"lib/annotations.jar",
					"lib/jsch-agent.jar",
					"lib/kotlinx-coroutines-slf4j-1.8.0-intellij.jar",
					"lib/product-client.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.completion.ml.ranking",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/completionMlRanking/lib/completionMlRanking.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.configurationScript",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/configurationScript/lib/configurationScript.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.copyright",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/copyright/lib/copyright.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.css",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/css-impl/lib/css-impl.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.database",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/DatabaseTools/lib/database-plugin.jar",
					"plugins/DatabaseTools/lib/jdbc-console.jar",
					"plugins/DatabaseTools/lib/database-openapi.jar",
					"plugins/DatabaseTools/lib/aether-dependency-resolver.jar",
					"plugins/DatabaseTools/lib/maven-resolver-transport-file.jar",
					"plugins/DatabaseTools/lib/maven-resolver-transport-http.jar",
					"plugins/DatabaseTools/lib/maven-resolver-connector-basic.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.dev",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/dev/lib/dev.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.diagram",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/uml/lib/uml-support.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.grpc",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/grpc/lib/grpc.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.ja",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/localization-ja/lib/localization-ja.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.jsonpath",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/jsonpath/lib/jsonpath.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.ko",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/localization-ko/lib/localization-ko.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.kubernetes",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/clouds-kubernetes/lib/clouds-kubernetes.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.marketplace",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.microservices.ui",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/microservices-ui/lib/microservices-ui.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.ml.inline.completion",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.coverage",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.database",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.database.connectivity",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.database.core",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.duplicatesDetector",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.externalSystem",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.go",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.go-capable",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.goland",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.javascript",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.json",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/json/lib/json.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.modules.lang",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.microservices",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.nativeDebug-plugin-capable",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.os.linux",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.os.unix",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.os.xwindow",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.platform",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.profiler",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.python-core-capable",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.python-in-mini-ide-capable",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.python-in-non-pycharm-ide-capable",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.remoteServers",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.run.targets",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.sql",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.ssh",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.ssh.ui",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.structuralsearch",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.ultimate",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.vcs",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.xdebugger",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.modules.xml",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.platform.experimental.backend",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.platform.experimental.frontend",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.platform.experimental.monolith",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.platform.ide.provisioner",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.platform.images",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/platform-images/lib/platform-images.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.plugins.colorful.darcula.colorscheme",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/color-scheme-colorful-darcula/lib/color-scheme-colorful-darcula.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.plugins.goland.solarized.colorscheme",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/color-scheme-solarized-goland/lib/color-scheme-solarized-goland.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.plugins.monokai.colorscheme",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/color-scheme-monokai/lib/color-scheme-monokai.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.plugins.vscodekeymap",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/keymap-vscode/lib/keymap-vscode.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.plugins.watcher",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/fileWatcher/lib/fileWatcher.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.react",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/react/lib/react.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.searcheverywhere.ml",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/searchEverywhereMl/lib/searchEverywhereMl.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.settingsSync",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/settingsSync/lib/settingsSync.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.station.core",
				Kind: "pluginAlias",
			},
			ArrayElem48{
				Name: "com.intellij.stylelint",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/stylelint/lib/stylelint.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.swagger",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/swagger/lib/swagger.jar",
				},
			},
			ArrayElem48{
				Name: "com.intellij.zh",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/localization-zh/lib/localization-zh.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.codeWithMe",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/cwm-plugin.jar",
					"plugins/cwm-plugin/lib/cwm-plugin-common.jar",
					"plugins/cwm-plugin/lib/cwm-plugin-backend.jar",
					"plugins/cwm-plugin/lib/cwm-plugin-android.jar",
					"plugins/cwm-plugin/lib/rd.jar",
					"plugins/cwm-plugin/lib/split.jar",
					"plugins/cwm-plugin/lib/ice4j.jar",
					"plugins/cwm-plugin/lib/jstun.jar",
					"plugins/cwm-plugin/lib/jitsi-utils.jar",
					"plugins/cwm-plugin/lib/tls-channel.jar",
					"plugins/cwm-plugin/lib/vcs-protocol.jar",
					"plugins/cwm-plugin/lib/backend-split.jar",
					"plugins/cwm-plugin/lib/split-protocol.jar",
					"plugins/cwm-plugin/lib/java-websocket.jar",
					"plugins/cwm-plugin/lib/quiche-jna-stubs.jar",
					"plugins/cwm-plugin/lib/remote-controller.jar",
					"plugins/cwm-plugin/lib/bouncy-castle-pgp.jar",
					"plugins/cwm-plugin/lib/grpc-netty-shaded.jar",
					"plugins/cwm-plugin/lib/jnaerator-runtime.jar",
					"plugins/cwm-plugin/lib/remote-controller-backend.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.gateway",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/gateway-plugin/lib/gateway-plugin.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.gateway.terminal",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/gateway-terminal/lib/gateway-terminal.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.performancePlugin",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/performanceTesting/lib/performanceTesting.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.performancePlugin.async",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/performanceTesting-async/lib/performanceTesting-async.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.plugins.webDeployment",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/webDeployment/lib/webDeployment.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.restClient",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/restClient/lib/restClient.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.sh",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/sh/lib/sh.jar",
				},
			},
			ArrayElem48{
				Name: "com.jetbrains.station",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/station-plugin/lib/station-plugin.jar",
				},
			},
			ArrayElem48{
				Name: "com.ypwang.plugin.go-linter",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/go-linter/lib/go-linter.jar",
				},
			},
			ArrayElem48{
				Name: "hg4idea",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/vcs-hg/lib/vcs-hg.jar",
				},
			},
			ArrayElem48{
				Name: "idea.plugin.protoeditor",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/protoeditor/lib/protoeditor.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.aqua.runners.cypress",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.aqua.runners.js.core",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.aqua.runners.playwright.js",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.charts",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/charts/lib/charts.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.compose",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.file",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.gateway.git4idea",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.gateway/forwardPorts",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.gateway/json",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.kubernetes",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.remoteRun",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.docker.ssh",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.kubernetes.charts.gotpl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.kubernetes.cron",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.clouds.kubernetes.db",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.completionMlRanking.experiments",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.copyright.vcs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.cwm.vcs",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.cwm.vcs.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.database.diagrams",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.database.docker",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.database.flameGraphs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.database.grazie",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.database.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.dev.codeInsight",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.dev.psiViewer",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.diagram.impl/vcs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.diagram.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.diagram.maven",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.diagram.properties",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.driver.backend.split",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.execution.process.elevation",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.execution.process.elevation.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.execution.process.elevation.backend.split",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.execution.process.mediator.client",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.execution.process.mediator.client.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.execution.process.mediator.common",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.execution.process.mediator.common.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.execution.process.mediator.daemon",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.execution.process.mediator.daemon.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.featuresTrainer.onboarding",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fileWatcher/webDeployment",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.api",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.core",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.core.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.core.impl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.cpp",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.cpp.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.less",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.less.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.css.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.postcss",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.postcss.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.sass",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.css.sass.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.experiments",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.federatedCompute",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.go",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.go.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.go.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.html",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.html.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.html.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.java.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.java.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.js",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.astro",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.astro.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.js.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.svelte",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.svelte.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.vue",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.js.vue.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.kotlin",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.kotlin.jupyter",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.kotlin.jupyter.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.kotlin.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.kotlin.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.php",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.php.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.php.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.python",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.python.jupyter",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.python.jupyter.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.python.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.python.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider.cpp",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider.cpp.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.rider.cpp.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider.csharp",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider.csharp.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.rider.csharp.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.rider.local",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.ruby",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.ruby.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.ruby.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.rust",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.rust.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.rust.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.terraform",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.terraform.local",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.fullLine.terraform.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.fullLine.vcs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.web",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.fullLine.xdebugger",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.git.commit.modal",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/vcs-git-commit-modal/lib/vcs-git-commit-modal.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.go.completionMlRanking",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/go-plugin/lib/modules/intellij.go.completionMlRanking.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.go.copyright",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.dfa",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.diagram",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.docker",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.featuresTrainer",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.grazie",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.langInjection",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.microservices",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.profiler",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.sharedIndexes",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.sharedIndexes.bundled/sharedCore",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.shell",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.template.integration",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.terminal",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.go.watchers",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.goland.customization",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.goland.customization.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.goland.featuresTrainer",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.goland.featuresTrainer.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grazie.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grazie.json",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grazie.markdown",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grazie.properties",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grazie.xml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grazie.yaml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grid",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.grid.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.charts.impl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grid.core.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.grid.core.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.csv.core.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.grid.csv.core.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.images.impl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grid.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.grid.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.json.impl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grid.plugin",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/grid-plugin/lib/grid-plugin.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.scripting.impl",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grid.scripting.rt",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/grid-plugin/lib/modules/intellij.grid.scripting.rt.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.grid.types",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.grid.types.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.groovy.dev",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.grpc/intelliLang",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.httpClient.injection.graphql",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.httpClient.textMate",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ide.startup.importSettings",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.ide.startup.importSettings.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/indexing-shared-ultimate-plugin-bundled/lib/indexing-shared-ultimate-plugin-bundled.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.core",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/indexing-shared/lib/indexing-shared.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.generator",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/indexing-shared-ultimate-plugin-bundled/lib/modules/intellij.indexing.shared.generator.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.ultimate",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/indexing-shared-ultimate-plugin-bundled/lib/modules/intellij.indexing.shared.ultimate.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.ultimate.git",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/indexing-shared-ultimate-plugin-bundled/lib/modules/intellij.indexing.shared.ultimate.git.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.ultimate.java",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/indexing-shared-ultimate-plugin-bundled/lib/modules/intellij.indexing.shared.ultimate.java.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.indexing.shared.ultimate.plugin.bundled/inner",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.java.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.java.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.java.dev",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.backend",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/javascript-debugger/lib/modules/intellij.javascript.debugger.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.dap.launcher",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.frontend",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/javascript-debugger/lib/modules/intellij.javascript.debugger.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.javaee",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.javaee/jsp",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.debugger.shared",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/javascript-debugger/lib/modules/intellij.javascript.debugger.shared.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.javascript.devkit",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.impl/diagrams",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.sh",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.sharedIndexes",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.javascript.sharedIndexes.generator",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.json",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/json/lib/modules/intellij.json.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.json.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/json/lib/modules/intellij.json.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.kotlin.dev",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.libraries.compose.foundation.desktop",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.compose.foundation.desktop.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.grpc",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.grpc.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.grpc.netty.shaded",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.grpc.netty.shaded.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.ktor.client",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.ktor.client.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.ktor.client.cio",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.ktor.client.cio.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.microba",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.microba.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.libraries.skiko",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.libraries.skiko.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.llmInstaller",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.llmInstaller.enterprise",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.llmInstaller.pycharm.community",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.backend.split",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.compose.preview",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.fenceInjection",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.frontmatter",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.frontmatter.toml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.frontmatter.yaml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.images",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.model",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.spellchecker",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.markdown.xml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.microservices.ui/diagram",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.cpp.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.cpp.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.css.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.css.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.css.less.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.css.postcss.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.css.sass.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.go.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.go.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.html.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.html.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.java.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.java.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.javascript.astro.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.javascript.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.javascript.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.javascript.svelte.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.javascript.vue.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.jupyter.kotlin.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.jupyter.python.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.kotlin.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.kotlin.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.markdown.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.php.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.php.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.python.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.python.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.rbs.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.rider.cpp.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.rider.cpp.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.rider.csharp.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.rider.csharp.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.ruby.completion",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/fullLine/lib/modules/intellij.ml.llm.ruby.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.ml.llm.rust.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.sql.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.terraform.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.text.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.vcs.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.xdebugger.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.ml.llm.yaml.completion",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.modules.database/dialects.oracle.utplsql.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.notebooks.backend.split",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.amper",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.go",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.gradle",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.javascript",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.maven",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.php",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.packageChecker.python",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.performanceTesting.async.core",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.performanceTesting.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.performanceTesting.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.performanceTesting.frontend",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/performanceTesting/lib/modules/intellij.performanceTesting.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.performanceTesting.remoteDriver",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.performanceTesting.vcs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.platform.bookmarks.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.bookmarks.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.bookmarks.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.bookmarks.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.clouds",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.clouds.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.collaborationTools",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.collaborationTools.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.compose",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.compose.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.coverage",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.coverage.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.coverage.agent",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.coverage.agent.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.dap",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.dap.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.debugger.impl.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.debugger.impl.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.debugger.impl.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.debugger.impl.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.diagnostic.freezeAnalyzer",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.diagnostic.freezeAnalyzer.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.diagnostic.freezes",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.diagnostic.freezes.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.editor",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.editor.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.editor.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.editor.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.editor.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.editor.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.execution.dashboard",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.execution.dashboard.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.execution.serviceView",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.execution.serviceView.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.ide.newUiOnboarding",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.ide.newUiOnboarding.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.ijent.impl",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/platform-ijent-impl/lib/platform-ijent-impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.images.backend.svg",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.images.backend.svg.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.images.copyright",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.images.copyright.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.inline.completion",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.inline.completion.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.inline.completion.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.inline.completion.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.inline.completion.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.inline.completion.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.foundation",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.foundation.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.ideLafBridge",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.ideLafBridge.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.markdown.core",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.markdown.core.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.markdown.extensions.gfmAlerts",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.markdown.extensions.gfmAlerts.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.markdown.extensions.gfmStrikethrough",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.markdown.extensions.gfmStrikethrough.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.markdown.extensions.gfmTables",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.markdown.extensions.gfmTables.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.markdown.ideLafBridgeStyling",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.markdown.ideLafBridgeStyling.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.jewel.ui",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.jewel.ui.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.kernel.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.kernel.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.kernel.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.kernel.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.lvcs.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.lvcs.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.navbar",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.navbar.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.navbar.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.navbar.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.navbar.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.navbar.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.navbar.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.navbar.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.navbar.monolith",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.navbar.monolith.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.progress.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.progress.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.project.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.project.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.recentFiles",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.recentFiles.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.recentFiles.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.recentFiles.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.recentFiles.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.recentFiles.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.registry.cloud",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.registry.cloud.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.rpc.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.rpc.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.rpc.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.rpc.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.searchEverywhere",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.searchEverywhere.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.searchEverywhere.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.searchEverywhere.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.searchEverywhere.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.searchEverywhere.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.settings.local",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.settings.local.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.smRunner.vcs",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.smRunner.vcs.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.statistics.devkit",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.platform.tips",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.tips.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.vcs.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.common.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.platform.vcs.common.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.dvcs.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.dvcs.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.dvcs.impl.shared",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.dvcs.impl.shared.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.backend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.exec",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.exec.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.frontend",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.frontend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.lang",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.lang.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.lang.actions",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.lang.actions.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.impl.shared",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.impl.shared.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.platform.vcs.log.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.platform.vcs.log.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.profiler.asyncOne",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.profiler.asyncOne.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.profiler.common",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.profiler.common.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.profiler.ultimate.ideaAsyncProfiler",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.profiler.ultimate.ideaAsyncProfiler.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.protoeditor.go",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.protoeditor.jvm",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.protoeditor.python",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.cpp",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.go",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.inspectionKts.js",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.inspectionKts.kotlin",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.js",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.dev",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.gradle",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.groovy",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.kotlin",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.maven",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.jvm.swagger",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.php",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.python.community",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.python.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.qodana.sarif",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/qodana/lib/modules/intellij.qodana.sarif.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.qodana/yaml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.remoteRun.terminal.backend.split",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.restClient/intelliLang",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.restClient/intelliLangInJs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.restClient/javaScript",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.restClient/jsonPath",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.restClient/microservicesUI",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.rml.dfa.impl",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.rml.dfa.impl.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.ranking.core",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/searchEverywhereMl/lib/modules/intellij.searchEverywhereMl.ranking.core.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.ranking.ext",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.ranking.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.ranking.vcs",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.ranking.yaml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.searchEverywhereMl.typos",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.settingsSync.core",
				Kind: "productModuleV2",
				ClassPath: []string{
					"lib/modules/intellij.settingsSync.core.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.sh.copyright",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.sh.markdown",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.sh.python",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.sh.terminal",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.station.core",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.station.frontend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/station-plugin/lib/modules/intellij.station.frontend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.swagger.endpoints",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.swagger.gradle",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.swagger.java",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.swagger.maven",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.terminal.backend",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.terminal.backend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/cwm-plugin/lib/modules/intellij.terminal.backend.split.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.terminal.frontend",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.terminal.sh",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.terminal/cloud",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.git.coverage",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.git.featuresTrainer",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.git/localHistory",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/vcs-git/lib/modules/intellij.vcs.git/localHistory.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.vcs.git/newUiOnboarding",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.git/terminal",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.github",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/vcs-github-IU/lib/modules/intellij.vcs.github.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.vcs.github.json",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.github.tracker",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.github.ultimate.cron",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.github.ultimate.docker",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.github.ultimate.features",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.gitlab",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.gitlab.ultimate.features",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.vcs.gitlab.yaml",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.webDeployment/phpStorm",
				Kind: "moduleV2",
			},
			ArrayElem48{
				Name: "intellij.webpack",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/webpack/lib/webpack.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.yaml",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/yaml/lib/modules/intellij.yaml.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.yaml.backend",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/yaml/lib/modules/intellij.yaml.backend.jar",
				},
			},
			ArrayElem48{
				Name: "intellij.yaml.frontend.split",
				Kind: "moduleV2",
				ClassPath: []string{
					"plugins/yaml/lib/modules/intellij.yaml.frontend.split.jar",
				},
			},
			ArrayElem48{
				Name: "name.kropp.intellij.makefile",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/makefile/lib/makefile.jar",
				},
			},
			ArrayElem48{
				Name: "org.editorconfig.editorconfigjetbrains",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/editorconfig/lib/editorconfig.jar",
				},
			},
			ArrayElem48{
				Name: "org.intellij.intelliLang",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/platform-langInjection/lib/platform-langInjection.jar",
					"plugins/platform-langInjection/lib/java-langInjection-jps.jar",
				},
			},
			ArrayElem48{
				Name: "org.intellij.plugins.markdown",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/markdown/lib/markdown.jar",
				},
			},
			ArrayElem48{
				Name: "org.intellij.qodana",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/qodana/lib/qodana.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.completion.full.line",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/fullLine/lib/fullLine.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.docker.gateway",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/clouds-docker-gateway/lib/clouds-docker-gateway.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.emojipicker",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/emojipicker/lib/emojipicker.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.github",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/vcs-github-IU/lib/vcs-github.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.gitlab",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/vcs-gitlab-IU/lib/vcs-gitlab.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.go",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/go-plugin/lib/go-plugin.jar",
					"plugins/go-plugin/lib/go-frontback.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.go-template",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/go-template/lib/go-template.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.go.ide",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/go-ide/lib/go-ide.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.go.sharedIndexes.bundled",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/go-sharedIndexes-bundled/lib/go-sharedIndexes-bundled.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.remote-run",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/remoteRun/lib/remoteRun.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.terminal",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/terminal/lib/terminal.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.textmate",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/textmate/lib/textmate.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.plugins.yaml",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/yaml/lib/yaml.jar",
				},
			},
			ArrayElem48{
				Name: "org.jetbrains.security.package-checker",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/packageChecker/lib/packageChecker.jar",
				},
			},
			ArrayElem48{
				Name: "tanvd.grazi",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/grazie/lib/grazie.jar",
				},
			},
			ArrayElem48{
				Name: "training",
				Kind: "plugin",
				ClassPath: []string{
					"plugins/featuresTrainer/lib/featuresTrainer.jar",
				},
			},
		},
	}
	root := RootContainer{Data: &data}
	fmt.Printf("%#v\\n", root)
}
