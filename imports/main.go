package main

import (
	"log"

	"golang_practice/imports/subfolder"
	"golang_practice/imports/subfolder/subsubfolder"
	subfolder2 "golang_practice/imports/subfolder_not_matching_package_name"
	"golang_practice/imports/submodule-module"
)

func main() {
	log.Print("in main")
	printFromAdjacent()
	subfolder.Print()
	subfolder2.Print()
	subsubfolder.Print()
	submodule.Print()
}
