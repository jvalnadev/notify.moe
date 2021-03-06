package main

import (
	"fmt"
	"html"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Unescape HTML entities in character names")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for character := range arn.StreamCharacters() {
		unescapedName := html.UnescapeString(character.Name.Canonical)

		if character.Name.Canonical != unescapedName {
			fmt.Println(character.Name.Canonical, "->", unescapedName)
			character.Name.Canonical = unescapedName
			character.Save()
		}
	}
}
