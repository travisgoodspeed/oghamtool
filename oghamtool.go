// This is a quick little tool for transliterating text in and out of
// the Ogham script of Old Irish.

package main

import (
	"flag"
	"fmt"
	"strings"
)

var tolatin = flag.String("latin", "", "Transliterates to ASCII.")
var toogham = flag.String("ogham", "", "Transliterates to Ogham.")

// Ogham alphabet in unicode order.  Three X's are reserved.
//oblock := " ᚁᚂᚃᚄᚅᚆᚇᚈᚉᚊᚋᚌᚍᚎᚏᚐᚑᚒᚓᚔᚕᚖᚗᚘᚙᚚ᚛᚜XXX"

//Go out of Ogham, every glyph is a key.
var oblock = map[string]string{
	" ": " ", //Space
	"ᚁ": "B",
	"ᚂ": "L",
	"ᚃ": "F",
	"ᚄ": "S",
	"ᚅ": "N",
	"ᚆ": "H",
	"ᚇ": "D",
	"ᚈ": "T",
	"ᚉ": "C",
	"ᚊ": "Q",
	"ᚋ": "M",
	"ᚌ": "G",
	"ᚍ": "NG",
	"ᚎ": "Z",
	"ᚏ": "R",
	"ᚐ": "A",
	"ᚑ": "O",
	"ᚒ": "U",
	"ᚓ": "E",
	"ᚔ": "I",
	"ᚕ": "EA", //Also K or X.
	"ᚖ": "OI",
	"ᚗ": "UI",
	"ᚘ": "IA",
	"ᚙ": "AE",
	"ᚚ": "P",
	"᚛": "`", //Totally nonstandard.
	"᚜": "'",
	//Three reserved letters haven't been defined yet.
}

var lblock = map[string]string{
	" ":  " ", //Space
	"B":  "ᚁ",
	"L":  "ᚂ",
	"F":  "ᚃ",
	"S":  "ᚄ",
	"N":  "ᚅ",
	"H":  "ᚆ",
	"D":  "ᚇ",
	"T":  "ᚈ",
	"C":  "ᚉ",
	"Q":  "ᚊ",
	"M":  "ᚋ",
	"G":  "ᚌ",
	"NG": "ᚍ",
	"Z":  "ᚎ",
	"R":  "ᚏ",
	"A":  "ᚐ",
	"O":  "ᚑ",
	"U":  "ᚒ",
	"E":  "ᚓ",
	"I":  "ᚔ",
	"EA": "ᚕ",
	"OI": "ᚖ",
	"UI": "ᚗ",
	"IA": "ᚘ",
	"AE": "ᚙ",
	"P":  "ᚚ",
	//Quotes are repeated in different styles.
	"`": "᚛", //LaTeX style.
	"'": "᚜",
	">": "᚛", //Easier to type.
	"<": "᚜",
	//These are duplicates in the other direction.
	"K": "ᚕ",
	"X": "ᚕ",
}

//! Prints Ogham text as Latin.
func latin(foo string) {
	for _, c := range string(foo) {
		fmt.Printf("%s", oblock[string(c)])
	}
	fmt.Printf("\n")
}

//! Converts Latin into Ogham.
func ogham(foo string) {
	skip := 0
	for i, c := range string(foo) {
		pair := ""
		opair := ""
		oletter := lblock[strings.ToUpper(string(c))]

		//Maybe we have a pair of letters, not standalone.
		if i+1 < len(foo) {
			pair = strings.ToUpper(foo[i : i+2])
			opair = lblock[pair]
		}

		if skip > 0 {
			skip = skip - 1
		} else if opair != "" {
			fmt.Printf("%s", opair)
			skip = 1 //Skip the next letter.
		} else {
			fmt.Printf("%s", oletter)
		}
	}
	fmt.Printf("\n")
}

//! Main method.
func main() {
	//Parses the command-line runes.
	flag.Parse()

	//Handle the runes here.
	if strings.Compare(*tolatin, "") != 0 {
		latin(*tolatin)
	} else if strings.Compare(*toogham, "") != 0 {
		ogham(*toogham)
	} else {
		fmt.Println("Try --help.")
	}
}
