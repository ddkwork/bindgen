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

type ArrayElem10 struct {
	Account_id          string       `json:"account_id"`
	Balances            balances     `json:"balances"`
	Days_available      int          `json:"days_available"`
	Historical_balances []ArrayElem5 `json:"historical_balances"`
	Mask                string       `json:"mask"`
	Name                string       `json:"name"`
	Official_name       string       `json:"official_name"`
	Owners              []ArrayElem8 `json:"owners"`
	Ownership_type      interface{}  `json:"ownership_type"`
	Subtype             string       `json:"subtype"`
	Transactions        []ArrayElem9 `json:"transactions"`
	TypeVal             string       `json:"type"`
}

type ArrayElem11 struct {
	Accounts          []ArrayElem10 `json:"accounts"`
	Date_last_updated string        `json:"date_last_updated"`
	Institution_id    string        `json:"institution_id"`
	Institution_name  string        `json:"institution_name"`
	Item_id           string        `json:"item_id"`
}

type ArrayElem2 struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Short_name  string `json:"short_name"`
	Plural_name string `json:"plural_name"`
	Icon        icon   `json:"icon"`
}

type ArrayElem3 struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ArrayElem4 struct {
	Close string `json:"close"`
	Day   int    `json:"day"`
	Open  string `json:"open"`
}

type ArrayElem5 struct {
	Current                  int         `json:"current"`
	Date                     string      `json:"date"`
	Iso_currency_code        string      `json:"iso_currency_code"`
	Unofficial_currency_code interface{} `json:"unofficial_currency_code"`
}

type ArrayElem6 struct {
	Data    data `json:"data"`
	Primary bool `json:"primary"`
}

type ArrayElem7 struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	TypeVal string `json:"type"`
}

type ArrayElem8 struct {
	Addresses     []ArrayElem6 `json:"addresses"`
	Emails        []ArrayElem7 `json:"emails"`
	Names         []string     `json:"names"`
	Phone_numbers []ArrayElem7 `json:"phone_numbers"`
}

type ArrayElem9 struct {
	Account_id               string      `json:"account_id"`
	Amount                   int         `json:"amount"`
	Date                     string      `json:"date"`
	Iso_currency_code        string      `json:"iso_currency_code"`
	Original_description     string      `json:"original_description"`
	Pending                  bool        `json:"pending"`
	Transaction_id           string      `json:"transaction_id"`
	Unofficial_currency_code interface{} `json:"unofficial_currency_code"`
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

type RootConfig2 struct {
	Fsq_id               string               `json:"fsq_id"`
	Categories           []ArrayElem2         `json:"categories"`
	Chains               []ArrayElem3         `json:"chains"`
	Closed_bucket        string               `json:"closed_bucket"`
	Date_closed          string               `json:"date_closed"`
	Description          string               `json:"description"`
	Distance             int                  `json:"distance"`
	Email                string               `json:"email"`
	Fax                  string               `json:"fax"`
	Features             features             `json:"features"`
	Geocodes             geocodes             `json:"geocodes"`
	Hours                hours                `json:"hours"`
	Hours_popular        []ArrayElem4         `json:"hours_popular"`
	Link                 string               `json:"link"`
	Location             location             `json:"location"`
	Menu                 string               `json:"menu"`
	Name                 string               `json:"name"`
	Photos               []icon               `json:"photos"`
	Popularity           int                  `json:"popularity"`
	Price                int                  `json:"price"`
	Rating               int                  `json:"rating"`
	Related_places       accepts_credit_cards `json:"related_places"`
	Social_media         social_media         `json:"social_media"`
	Stats                stats                `json:"stats"`
	Store_id             string               `json:"store_id"`
	Tastes               []string             `json:"tastes"`
	Tel                  string               `json:"tel"`
	Timezone             string               `json:"timezone"`
	Tips                 []tip                `json:"tips"`
	Venue_reality_bucket string               `json:"venue_reality_bucket"`
	Verified             bool                 `json:"verified"`
	Website              string               `json:"website"`
}

type RootConfig3 struct {
	Report     report        `json:"report"`
	Request_id string        `json:"request_id"`
	Warnings   []interface{} `json:"warnings"`
}

type accepts_credit_cards struct {
}

type alcohol struct {
	Bar_service accepts_credit_cards `json:"bar_service"`
	Beer        accepts_credit_cards `json:"beer"`
	Byo         accepts_credit_cards `json:"byo"`
	Cocktails   accepts_credit_cards `json:"cocktails"`
	Full_bar    accepts_credit_cards `json:"full_bar"`
	Wine        accepts_credit_cards `json:"wine"`
}

type amenities struct {
	Restroom              accepts_credit_cards `json:"restroom"`
	Smoking               accepts_credit_cards `json:"smoking"`
	Jukebox               accepts_credit_cards `json:"jukebox"`
	Music                 accepts_credit_cards `json:"music"`
	Live_music            accepts_credit_cards `json:"live_music"`
	Private_room          accepts_credit_cards `json:"private_room"`
	Outdoor_seating       accepts_credit_cards `json:"outdoor_seating"`
	Tvs                   accepts_credit_cards `json:"tvs"`
	Atm                   accepts_credit_cards `json:"atm"`
	Coat_check            accepts_credit_cards `json:"coat_check"`
	Wheelchair_accessible accepts_credit_cards `json:"wheelchair_accessible"`
	Parking               parking              `json:"parking"`
	Sit_down_dining       accepts_credit_cards `json:"sit_down_dining"`
	Wifi                  string               `json:"wifi"`
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

type balances struct {
	Available                int         `json:"available"`
	Current                  int         `json:"current"`
	Limit                    interface{} `json:"limit"`
	Margin_loan_amount       interface{} `json:"margin_loan_amount"`
	Iso_currency_code        string      `json:"iso_currency_code"`
	Unofficial_currency_code interface{} `json:"unofficial_currency_code"`
}

type className struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}

type credit_cards struct {
	Accepts_credit_cards accepts_credit_cards `json:"accepts_credit_cards"`
	Amex                 accepts_credit_cards `json:"amex"`
	Discover             accepts_credit_cards `json:"discover"`
	Visa                 accepts_credit_cards `json:"visa"`
	Diners_club          accepts_credit_cards `json:"diners_club"`
	Master_card          accepts_credit_cards `json:"master_card"`
	Union_pay            accepts_credit_cards `json:"union_pay"`
}

type data struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	Region      string `json:"region"`
	Street      string `json:"street"`
	Postal_code string `json:"postal_code"`
}

type digital_wallet struct {
	Accepts_nfc accepts_credit_cards `json:"accepts_nfc"`
}

type dine_in struct {
	Reservations             accepts_credit_cards `json:"reservations"`
	Online_reservations      accepts_credit_cards `json:"online_reservations"`
	Groups_only_reservations accepts_credit_cards `json:"groups_only_reservations"`
	Essential_reservations   accepts_credit_cards `json:"essential_reservations"`
}

type drop_off struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type features struct {
	Payment        payment        `json:"payment"`
	Food_and_drink food_and_drink `json:"food_and_drink"`
	Services       services       `json:"services"`
	Amenities      amenities      `json:"amenities"`
	Attributes     attributes     `json:"attributes"`
}

type food_and_drink struct {
	Alcohol alcohol `json:"alcohol"`
	Meals   meals   `json:"meals"`
}

type geocodes struct {
	Drop_off   drop_off `json:"drop_off"`
	Front_door drop_off `json:"front_door"`
	Main       drop_off `json:"main"`
	Road       drop_off `json:"road"`
	Roof       drop_off `json:"roof"`
}

type hours struct {
	Display          string       `json:"display"`
	Is_local_holiday bool         `json:"is_local_holiday"`
	Open_now         bool         `json:"open_now"`
	Regular          []ArrayElem4 `json:"regular"`
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
	Bar_snacks   accepts_credit_cards `json:"bar_snacks"`
	Breakfast    accepts_credit_cards `json:"breakfast"`
	Brunch       accepts_credit_cards `json:"brunch"`
	Lunch        accepts_credit_cards `json:"lunch"`
	Happy_hour   accepts_credit_cards `json:"happy_hour"`
	Dessert      accepts_credit_cards `json:"dessert"`
	Dinner       accepts_credit_cards `json:"dinner"`
	Tasting_menu accepts_credit_cards `json:"tasting_menu"`
}

type parking struct {
	Parking        accepts_credit_cards `json:"parking"`
	Street_parking accepts_credit_cards `json:"street_parking"`
	Valet_parking  accepts_credit_cards `json:"valet_parking"`
	Public_lot     accepts_credit_cards `json:"public_lot"`
	Private_lot    accepts_credit_cards `json:"private_lot"`
}

type patternProperties struct {
	PatternProperties interface{} `json:"^(loadingConfig|$$id|className|$ref|disabled|disabledOn|hidden|hiddenOn|visible|visibleOn|id|onEvent|static|staticOn|staticPlaceholder|staticClassName|staticLabelClassName|staticInputClassName|staticSchema|style|editorSetting|useMobileUI|testIdBuilder|size|label|labelAlign|labelWidth|labelClassName|name|extraName|remark|labelRemark|hint|submitOnChange|readOnly|readOnlyOn|validateOnChange|description|desc|descriptionClassName|mode|horizontal|inline|inputClassName|placeholder|required|validationErrors|validations|value|clearValueOnHidden|validateApi|autoFill|initAutoFill|row|type|extractValue|joinValues|delimiter|allowCity|allowDistrict|allowStreet|searchable|itemClassName|columnClassName|columnRatio)$"`
}

type payment struct {
	Credit_cards   credit_cards   `json:"credit_cards"`
	Digital_wallet digital_wallet `json:"digital_wallet"`
}

type report struct {
	Asset_report_id  string        `json:"asset_report_id"`
	Client_report_id string        `json:"client_report_id"`
	Date_generated   string        `json:"date_generated"`
	Days_requested   int           `json:"days_requested"`
	Items            []ArrayElem11 `json:"items"`
	User             user          `json:"user"`
}

type services struct {
	Delivery      accepts_credit_cards `json:"delivery"`
	Takeout       accepts_credit_cards `json:"takeout"`
	Drive_through accepts_credit_cards `json:"drive_through"`
	Dine_in       dine_in              `json:"dine_in"`
}

type social_media struct {
	Facebook_id string `json:"facebook_id"`
	Instagram   string `json:"instagram"`
	Twitter     string `json:"twitter"`
}

type stats struct {
	Total_photos  int `json:"total_photos"`
	Total_ratings int `json:"total_ratings"`
	Total_tips    int `json:"total_tips"`
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

type user struct {
	Client_user_id string `json:"client_user_id"`
	Email          string `json:"email"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Middle_name    string `json:"middle_name"`
	Phone_number   string `json:"phone_number"`
	Ssn            string `json:"ssn"`
}

type RootContainer struct {
	Data *RootConfig3 `json:"data"`
}

func main() {
	data := RootConfig3{
		Report: report{
			Asset_report_id:  "028e8404-a013-4a45-ac9e-002482f9cafc",
			Client_report_id: "client_report_id_1221",
			Date_generated:   "2023-03-30T18:27:37Z",
			Days_requested:   5,
			Items: []ArrayElem11{
				ArrayElem11{
					Accounts: []ArrayElem10{
						ArrayElem10{
							Account_id: "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
							Balances: balances{
								Available:                43200,
								Current:                  43200,
								Limit:                    nil,
								Margin_loan_amount:       nil,
								Iso_currency_code:        "USD",
								Unofficial_currency_code: nil,
							},
							Days_available: 5,
							Historical_balances: []ArrayElem5{
								ArrayElem5{
									Current:                  49050,
									Date:                     "2023-03-29",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  49050,
									Date:                     "2023-03-28",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  49050,
									Date:                     "2023-03-27",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  49050,
									Date:                     "2023-03-26",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  49050,
									Date:                     "2023-03-25",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
							},
							Mask:          "4444",
							Name:          "Plaid Money Market",
							Official_name: "Plaid Platinum Standard 1.85% Interest Money Market",
							Owners: []ArrayElem8{
								ArrayElem8{
									Addresses: []ArrayElem6{
										ArrayElem6{
											Data: data{
												City:        "Malakoff",
												Country:     "US",
												Region:      "NY",
												Street:      "2992 Cameron Road",
												Postal_code: "14236",
											},
											Primary: true,
										},
										ArrayElem6{
											Data: data{
												City:        "San Matias",
												Country:     "US",
												Region:      "CA",
												Street:      "2493 Leisure Lane",
												Postal_code: "93405-2255",
											},
											Primary: false,
										},
									},
									Emails: []ArrayElem7{
										ArrayElem7{
											Data:    "accountholder0@example.com",
											Primary: true,
											TypeVal: "primary",
										},
										ArrayElem7{
											Data:    "accountholder1@example.com",
											Primary: false,
											TypeVal: "secondary",
										},
										ArrayElem7{
											Data:    "extraordinarily.long.email.username.123456@reallylonghostname.com",
											Primary: false,
											TypeVal: "other",
										},
									},
									Names: []string{
										"Alberta Bobbeth Charleson",
									},
									Phone_numbers: []ArrayElem7{
										ArrayElem7{
											Data:    "1112223333",
											Primary: false,
											TypeVal: "home",
										},
										ArrayElem7{
											Data:    "1112224444",
											Primary: false,
											TypeVal: "work",
										},
										ArrayElem7{
											Data:    "1112225555",
											Primary: false,
											TypeVal: "mobile",
										},
									},
								},
							},
							Ownership_type: nil,
							Subtype:        "money market",
							Transactions: []ArrayElem9{
								ArrayElem9{
									Account_id:               "1qKRXQjk8xUWDJojNwPXTj8gEmR48piqRNye8",
									Amount:                   5850,
									Date:                     "2023-03-30",
									Iso_currency_code:        "USD",
									Original_description:     "ACH Electronic CreditGUSTO PAY 123456",
									Pending:                  false,
									Transaction_id:           "gGQgjoeyqBF89PND6K14Sow1wddZBmtLomJ78",
									Unofficial_currency_code: nil,
								},
							},
							TypeVal: "depository",
						},
						ArrayElem10{
							Account_id: "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
							Balances: balances{
								Available:                100,
								Current:                  110,
								Limit:                    nil,
								Margin_loan_amount:       nil,
								Iso_currency_code:        "USD",
								Unofficial_currency_code: nil,
							},
							Days_available: 5,
							Historical_balances: []ArrayElem5{
								ArrayElem5{
									Current:                  110,
									Date:                     "2023-03-29",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  -390,
									Date:                     "2023-03-28",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  -373,
									Date:                     "2023-03-27",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  -284,
									Date:                     "2023-03-26",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
								ArrayElem5{
									Current:                  -284,
									Date:                     "2023-03-25",
									Iso_currency_code:        "USD",
									Unofficial_currency_code: nil,
								},
							},
							Mask:          "0000",
							Name:          "Plaid Checking",
							Official_name: "Plaid Gold Standard 0% Interest Checking",
							Owners: []ArrayElem8{
								ArrayElem8{
									Addresses: []ArrayElem6{
										ArrayElem6{
											Data: data{
												City:        "Malakoff",
												Country:     "US",
												Region:      "NY",
												Street:      "2992 Cameron Road",
												Postal_code: "14236",
											},
											Primary: true,
										},
										ArrayElem6{
											Data: data{
												City:        "San Matias",
												Country:     "US",
												Region:      "CA",
												Street:      "2493 Leisure Lane",
												Postal_code: "93405-2255",
											},
											Primary: false,
										},
									},
									Emails: []ArrayElem7{
										ArrayElem7{
											Data:    "accountholder0@example.com",
											Primary: true,
											TypeVal: "primary",
										},
										ArrayElem7{
											Data:    "accountholder1@example.com",
											Primary: false,
											TypeVal: "secondary",
										},
										ArrayElem7{
											Data:    "extraordinarily.long.email.username.123456@reallylonghostname.com",
											Primary: false,
											TypeVal: "other",
										},
									},
									Names: []string{
										"Alberta Bobbeth Charleson",
									},
									Phone_numbers: []ArrayElem7{
										ArrayElem7{
											Data:    "1112223333",
											Primary: false,
											TypeVal: "home",
										},
										ArrayElem7{
											Data:    "1112224444",
											Primary: false,
											TypeVal: "work",
										},
										ArrayElem7{
											Data:    "1112225555",
											Primary: false,
											TypeVal: "mobile",
										},
									},
								},
							},
							Ownership_type: nil,
							Subtype:        "checking",
							Transactions: []ArrayElem9{
								ArrayElem9{
									Account_id:               "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
									Amount:                   89,
									Date:                     "2023-03-27",
									Iso_currency_code:        "USD",
									Original_description:     "SparkFun",
									Pending:                  false,
									Transaction_id:           "4zBRq1Qem4uAPnoyKjJNTRQpQddM4ztlo1PLD",
									Unofficial_currency_code: nil,
								},
								ArrayElem9{
									Account_id:               "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
									Amount:                   12,
									Date:                     "2023-03-28",
									Iso_currency_code:        "USD",
									Original_description:     "McDonalds #3322",
									Pending:                  false,
									Transaction_id:           "dkjL41PnbKsPral79jpxhMWdW55gkPfBkWpRL",
									Unofficial_currency_code: nil,
								},
								ArrayElem9{
									Account_id:               "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
									Amount:                   4,
									Date:                     "2023-03-28",
									Iso_currency_code:        "USD",
									Original_description:     "Starbucks",
									Pending:                  false,
									Transaction_id:           "a84ZxQaWDAtDL3dRgmazT57K7jjN3WFkNWMDy",
									Unofficial_currency_code: nil,
								},
								ArrayElem9{
									Account_id:               "eG7pNLjknrFpWvP7Dkbdf3Pq6GVBPKTaQJK5v",
									Amount:                   -500,
									Date:                     "2023-03-29",
									Iso_currency_code:        "USD",
									Original_description:     "United Airlines **** REFUND ****",
									Pending:                  false,
									Transaction_id:           "xG9jbv3eMoFWepzB7wQLT3LoLggX5Duy1Gbe5",
									Unofficial_currency_code: nil,
								},
							},
							TypeVal: "depository",
						},
					},
					Date_last_updated: "2023-03-30T18:25:26Z",
					Institution_id:    "ins_109508",
					Institution_name:  "First Platypus Bank",
					Item_id:           "AZMP7JrGXgtPd3AQMeg7hwMKgk5E8qU1V5ME7",
				},
			},
			User: user{
				Client_user_id: "uid_40332",
				Email:          "abcharleston@example.com",
				First_name:     "Anna",
				Last_name:      "Charleston",
				Middle_name:    "B",
				Phone_number:   "1-415-867-5309",
				Ssn:            "111-22-1234",
			},
		},
		Request_id: "GVzMdiDd8DDAQK4",
		Warnings:   []interface{}{},
	}
	root := RootContainer{Data: &data}
	fmt.Printf("%#v\\n", root)
}
