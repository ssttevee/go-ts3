// Package ts3util contains helper functions for the ts3 package
package ts3util

import "strings"

var (
	escapeTable = map[rune][]rune{
		92:  {92, 92},
		47:  {92, 47},
		32:  {92, 115},
		124: {92, 112},
		7:   {92, 97},
		8:   {92, 98},
		12:  {92, 102},
		10:  {92, 110},
		13:  {92, 114},
		9:   {92, 116},
		11:  {92, 118},
	}

	escaper   *strings.Replacer
	unescaper *strings.Replacer
)

func init() {
	escaperPairs := make([]string, 0, len(escapeTable)*2)
	unescaperPairs := make([]string, 0, len(escapeTable)*2)

	for char, replace := range escapeTable {
		needle := string([]rune{char})
		replacement := string(replace)

		escaperPairs = append(escaperPairs, needle, replacement)
		unescaperPairs = append(unescaperPairs, replacement, needle)
	}

	escaper = strings.NewReplacer(escaperPairs...)
	unescaper = strings.NewReplacer(unescaperPairs...)
}

// Escape will escaped string according to TeamSpeak's Server Query specs
func Escape(s string) string {
	return escaper.Replace(s)
}

// Unescape will unescape the string according to TeamSpeak's Server Query specs
func Unescape(s string) string {
	return unescaper.Replace(s)
}
