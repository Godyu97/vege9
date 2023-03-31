package main

import (
	"github.com/Godyu97/vege9/vegePcre"
	"log"
)

func main() {
	str := "Hello (world)!"
	replace_str := "<$1>"
	patten := "\\((.*?)\\)"
	log.Println(vegePcre.Replace(patten, replace_str, str))
}
