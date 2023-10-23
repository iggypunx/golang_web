package main

import (
	"log"
	"github.com/iggypunx/golang_web/GoWebPackage/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
