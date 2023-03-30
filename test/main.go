package main

import (
	"github.com/Godyu97/vege9/vegePcre"
)

func main() {
	vegePcre.Replace("\\((.*?)\\)", "<\\1>", "Hello (world)!")
}
