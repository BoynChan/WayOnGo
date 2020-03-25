package main

import (
	"fmt"
	"testing"
)

//Author: Boyn
//Date: 2020/3/24
// use test file as a test function of s expresion

func TestMarshal(t *testing.T) {
	movie := Movie{
		Title:    "Titanic",
		Subtitle: "Jack And Rose",
		Year:     1999,
		Color:    true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	marshal, err := Marshal(movie)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
}
