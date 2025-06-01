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

type RootConfig struct {
	Type__       string        `json:"__type__"`
	Body         []ArrayElem1  `json:"body"`
	Type_ignores []interface{} `json:"type_ignores"`
}

type RootContainer struct {
	Data *RootConfig `json:"data"`
}

func main() {
	data := RootConfig{
		Type__: "Module",
		Body: []ArrayElem1{
			ArrayElem1{
				Type__: "Import",
				Names: []ArrayElem{
					ArrayElem{
						Type__:     "alias",
						Name:       "socket",
						Asname:     nil,
						Lineno:     2,
						Col_offset: 7,
					},
					ArrayElem{
						Type__:     "alias",
						Name:       "struct",
						Asname:     nil,
						Lineno:     2,
						Col_offset: 15,
					},
				},
				Lineno:     2,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__: "ImportFrom",
				Names: []ArrayElem{
					ArrayElem{
						Type__:     "alias",
						Name:       "*",
						Asname:     nil,
						Lineno:     3,
						Col_offset: 19,
					},
				},
				Lineno:     3,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     9,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     11,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     12,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     13,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     14,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     16,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     17,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     18,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     19,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     21,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     22,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     23,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     24,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     26,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     27,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     28,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     29,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     31,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     32,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     33,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     34,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     35,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     36,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     37,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     38,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     39,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     41,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     42,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     43,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     44,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     45,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     46,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     48,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     49,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     50,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     51,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     52,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     53,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     54,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     55,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     56,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     57,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     59,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     60,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     61,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     62,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     63,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     64,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     65,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     66,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     67,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     68,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     70,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     71,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     72,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     73,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     74,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     76,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     77,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     78,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     79,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     81,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     82,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     83,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     84,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     86,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     87,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     88,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     89,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     91,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     92,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     93,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     94,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     95,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     96,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     97,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     98,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     99,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     101,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     102,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     103,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     104,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     105,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     106,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     108,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     109,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     110,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     111,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     112,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     113,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     114,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     115,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     116,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     117,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     119,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     120,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     121,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     122,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     123,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     124,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     125,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     126,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     127,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     128,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     131,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     132,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     133,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     134,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     135,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     136,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     137,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     138,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     139,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     140,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     141,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     142,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     143,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     146,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     147,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     148,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     149,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     150,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     151,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     152,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     153,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     156,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     157,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     158,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     159,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     160,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     161,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     162,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     165,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     166,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     167,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     168,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     171,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     172,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     173,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     174,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     175,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     176,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     177,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     178,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     179,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     180,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     181,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     182,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     183,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     184,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     187,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     188,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     189,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     190,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     191,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     194,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     195,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     196,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     197,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     200,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     201,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     202,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     203,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     204,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     207,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     208,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     209,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     210,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     211,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     212,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     215,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     216,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     217,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     218,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     219,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     220,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     223,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     224,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     225,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     226,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     227,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     228,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     229,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     230,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     233,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     234,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     235,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     236,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     237,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     238,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     239,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     240,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     241,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     242,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     243,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     244,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     245,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     246,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     247,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     248,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     249,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     250,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     251,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     252,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     253,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     254,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     255,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     256,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     257,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     258,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     259,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     260,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     261,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     262,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     263,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     264,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     265,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     268,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     269,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     270,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     271,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     272,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     273,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     274,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     275,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     276,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     277,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     278,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     279,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     280,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     281,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     282,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     283,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     284,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     285,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     288,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "Assign",
				Lineno:     289,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "ClassDef",
				Lineno:     294,
				Col_offset: 0,
			},
			ArrayElem1{
				Type__:     "ClassDef",
				Lineno:     352,
				Col_offset: 0,
			},
		},
		Type_ignores: []interface{}{},
	}
	root := RootContainer{Data: &data}
	fmt.Printf("%#v\\n", root)
}
