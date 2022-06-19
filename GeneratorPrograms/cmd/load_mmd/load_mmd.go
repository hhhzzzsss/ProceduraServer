package main

import "github.com/hhhzzzsss/procedura-generator/mmd_loader"

func main() {
	r := mmd_loader.LoadDumpAsRegion()
	r.CreateDump()
}
