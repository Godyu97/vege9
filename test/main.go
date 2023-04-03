package main

import (
	"github.com/Godyu97/vege9/vegePcre"
	"log"
)

func main() {
	str := "Hello (world)!"
	replace_str := "<$1>"
	patten := "\\((.*?)\\)"
	log.Println(vegePcre.PcreppReplaceImpl(patten, replace_str, str))
}
