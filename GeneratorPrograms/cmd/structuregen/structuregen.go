package main

import (
	"fmt"

	"github.com/hhhzzzsss/procedura-generator/structuregen"
)

func main() {
	settings := structuregen.GetDefaultSettings()

	fmt.Println("Generating structure...")
	region := structuregen.GenerateStructure(&settings)

	region.CreateDump()
}
