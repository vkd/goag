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
	basePath     = flag.String("basepath", "", "Base path prefix")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	var err error
	if dir != nil && *dir != "" {
		err = goag.GenerateDir(*dir, *packageName, *specFilename, *basePath)
	} else {
		err = goag.GenerateFile(*outDir, *packageName, *specFile, *basePath)
	}
	if err != nil {
		log.Fatalf("Error on generate: %v", err)
	}
}
