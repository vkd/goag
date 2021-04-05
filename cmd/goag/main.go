package main

import (
	"flag"
	"log"

	"github.com/vkd/goag"
)

var (
	specFile     = flag.String("file", "openapi.yaml", "Spec file")
	dir          = flag.String("dir", "", "Specs dir")
	outDir       = flag.String("out", "./", "output dif")
	packageName  = flag.String("package", "simple", "package name")
	specFilename = flag.String("spec", "openapi.yaml", "spec filename")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	var err error
	if dir != nil && *dir != "" {
		err = goag.GenerateDir(*dir, *packageName, *specFilename)
	} else {
		err = goag.GenerateFile(*outDir, *packageName, *specFile)
	}
	if err != nil {
		log.Fatalf("Error on generate: %v", err)
	}
}
