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

type RootContainer struct {
	Data *RootConfig2 `json:"data"`
}

func main() {
	data := RootConfig2{
		Fsq_id: "string",
		Categories: []ArrayElem2{
			ArrayElem2{
				Id:          0,
				Name:        "string",
				Short_name:  "string",
				Plural_name: "string",
				Icon: icon{
					Id:         "string",
					Created_at: "2024-07-18T03:33:17.149Z",
					Prefix:     "string",
					Suffix:     "string",
					Width:      0,
					Height:     0,
					Classifications: []string{
						"string",
					},
					Tip: tip{
						Id:             "string",
						Created_at:     "2024-07-18T03:33:17.149Z",
						Text:           "string",
						Url:            "string",
						Lang:           "string",
						Agree_count:    0,
						Disagree_count: 0,
					},
				},
			},
		},
		Chains: []ArrayElem3{
			ArrayElem3{
				Id:   "string",
				Name: "string",
			},
		},
		Closed_bucket: "string",
		Date_closed:   "2024-07-18",
		Description:   "string",
		Distance:      0,
		Email:         "string",
		Fax:           "string",
		Features: features{
			Payment: payment{
				Credit_cards: credit_cards{
					Accepts_credit_cards: accepts_credit_cards{},
					Amex:                 accepts_credit_cards{},
					Discover:             accepts_credit_cards{},
					Visa:                 accepts_credit_cards{},
					Diners_club:          accepts_credit_cards{},
					Master_card:          accepts_credit_cards{},
					Union_pay:            accepts_credit_cards{},
				},
				Digital_wallet: digital_wallet{
					Accepts_nfc: accepts_credit_cards{},
				},
			},
			Food_and_drink: food_and_drink{
				Alcohol: alcohol{
					Bar_service: accepts_credit_cards{},
					Beer:        accepts_credit_cards{},
					Byo:         accepts_credit_cards{},
					Cocktails:   accepts_credit_cards{},
					Full_bar:    accepts_credit_cards{},
					Wine:        accepts_credit_cards{},
				},
				Meals: meals{
					Bar_snacks:   accepts_credit_cards{},
					Breakfast:    accepts_credit_cards{},
					Brunch:       accepts_credit_cards{},
					Lunch:        accepts_credit_cards{},
					Happy_hour:   accepts_credit_cards{},
					Dessert:      accepts_credit_cards{},
					Dinner:       accepts_credit_cards{},
					Tasting_menu: accepts_credit_cards{},
				},
			},
			Services: services{
				Delivery:      accepts_credit_cards{},
				Takeout:       accepts_credit_cards{},
				Drive_through: accepts_credit_cards{},
				Dine_in: dine_in{
					Reservations:             accepts_credit_cards{},
					Online_reservations:      accepts_credit_cards{},
					Groups_only_reservations: accepts_credit_cards{},
					Essential_reservations:   accepts_credit_cards{},
				},
			},
			Amenities: amenities{
				Restroom:              accepts_credit_cards{},
				Smoking:               accepts_credit_cards{},
				Jukebox:               accepts_credit_cards{},
				Music:                 accepts_credit_cards{},
				Live_music:            accepts_credit_cards{},
				Private_room:          accepts_credit_cards{},
				Outdoor_seating:       accepts_credit_cards{},
				Tvs:                   accepts_credit_cards{},
				Atm:                   accepts_credit_cards{},
				Coat_check:            accepts_credit_cards{},
				Wheelchair_accessible: accepts_credit_cards{},
				Parking: parking{
					Parking:        accepts_credit_cards{},
					Street_parking: accepts_credit_cards{},
					Valet_parking:  accepts_credit_cards{},
					Public_lot:     accepts_credit_cards{},
					Private_lot:    accepts_credit_cards{},
				},
				Sit_down_dining: accepts_credit_cards{},
				Wifi:            "string",
			},
			Attributes: attributes{
				Business_meeting: "string",
				Clean:            "string",
				Crowded:          "string",
				Dates_popular:    "string",
				Dressy:           "string",
				Families_popular: "string",
				Gluten_free_diet: "string",
				Good_for_dogs:    "string",
				Groups_popular:   "string",
				Healthy_diet:     "string",
				Late_night:       "string",
				Noisy:            "string",
				Quick_bite:       "string",
				Romantic:         "string",
				Service_quality:  "string",
				Singles_popular:  "string",
				Special_occasion: "string",
				Trendy:           "string",
				Value_for_money:  "string",
				Vegan_diet:       "string",
				Vegetarian_diet:  "string",
			},
		},
		Geocodes: geocodes{
			Drop_off: drop_off{
				Latitude:  0,
				Longitude: 0,
			},
			Front_door: drop_off{
				Latitude:  0,
				Longitude: 0,
			},
			Main: drop_off{
				Latitude:  0,
				Longitude: 0,
			},
			Road: drop_off{
				Latitude:  0,
				Longitude: 0,
			},
			Roof: drop_off{
				Latitude:  0,
				Longitude: 0,
			},
		},
		Hours: hours{
			Display:          "string",
			Is_local_holiday: true,
			Open_now:         true,
			Regular: []ArrayElem4{
				ArrayElem4{
					Close: "string",
					Day:   0,
					Open:  "string",
				},
			},
		},
		Hours_popular: []ArrayElem4{
			ArrayElem4{
				Close: "string",
				Day:   0,
				Open:  "string",
			},
		},
		Link: "string",
		Location: location{
			Address:           "string",
			Address_extended:  "string",
			Admin_region:      "string",
			Census_block:      "string",
			Country:           "string",
			Cross_street:      "string",
			Dma:               "string",
			Formatted_address: "string",
			Locality:          "string",
			Neighborhood: []string{
				"string",
			},
			Po_box:    "string",
			Post_town: "string",
			Postcode:  "string",
			Region:    "string",
		},
		Menu: "string",
		Name: "string",
		Photos: []icon{
			icon{
				Id:         "string",
				Created_at: "2024-07-18T03:33:17.149Z",
				Prefix:     "string",
				Suffix:     "string",
				Width:      0,
				Height:     0,
				Classifications: []string{
					"string",
				},
				Tip: tip{
					Id:             "string",
					Created_at:     "2024-07-18T03:33:17.149Z",
					Text:           "string",
					Url:            "string",
					Lang:           "string",
					Agree_count:    0,
					Disagree_count: 0,
				},
			},
		},
		Popularity:     0,
		Price:          0,
		Rating:         0,
		Related_places: accepts_credit_cards{},
		Social_media: social_media{
			Facebook_id: "string",
			Instagram:   "string",
			Twitter:     "string",
		},
		Stats: stats{
			Total_photos:  0,
			Total_ratings: 0,
			Total_tips:    0,
		},
		Store_id: "string",
		Tastes: []string{
			"string",
		},
		Tel:      "string",
		Timezone: "string",
		Tips: []tip{
			tip{
				Id:             "string",
				Created_at:     "2024-07-18T03:33:17.149Z",
				Text:           "string",
				Url:            "string",
				Lang:           "string",
				Agree_count:    0,
				Disagree_count: 0,
			},
		},
		Venue_reality_bucket: "string",
		Verified:             true,
		Website:              "string",
	}
	root := RootContainer{Data: &data}
	fmt.Printf("%#v\\n", root)
}
