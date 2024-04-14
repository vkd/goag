package main

import (
	"flag"
	"log"

	"github.com/vkd/goag"
)

var (
	specFile     = flag.String("file", "openapi.yaml", "Spec file")
	dir          = flag.String("dir", "", "Specs dir. Used for unittests.")
	outDir       = flag.String("out", "./", "output dif")
	packageName  = flag.String("package", "simple", "package name")
	specFilename = flag.String("spec", "openapi.yaml", "spec filename")
	cfgFilename  = flag.String("config", "goag.yaml", "config filename")
	basePath     = flag.String("basepath", "", "Base path prefix")
	genClient    = flag.Bool("client", false, "Generate client code")
	deleteOld    = flag.Bool("delete", false, "Delete old files")
	doNotEdit    = flag.Bool("donotedit", true, "Add 'DO NOT EDIT' headers")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	var g goag.Generator
	g.GenClient = *genClient
	g.DeleteOld = *deleteOld
	g.DoNotEdit = *doNotEdit

	var err error
	if dir != nil && *dir != "" {
		err = g.GenerateDir(*dir, *outDir, *packageName, *specFilename, *basePath, *cfgFilename)
	} else {
		err = g.GenerateFile(*outDir, *packageName, *specFile, *basePath, *cfgFilename)
	}
	if err != nil {
		log.Fatalf("Error on generate: %v", err)
	}
}
