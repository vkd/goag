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
	cfgFilename  = flag.String("config", ".goag.yaml", "config filename")
	basePath     = flag.String("basepath", "", "Base path prefix")
	genClient    = flag.Bool("client", false, "Generate client code")
	doNotEdit    = flag.Bool("donotedit", true, "Add 'DO NOT EDIT' headers")
	specHandler  = flag.String("spec-handler-name", "openapi.yaml", "Handler spec filename")
	isApiHandler = flag.Bool("api-handler", true, "Generate api handler")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	var g goag.Generator
	g.GenClient = *genClient
	g.GenAPIHandler = *isApiHandler
	g.DoNotEdit = *doNotEdit

	var err error
	if dir != nil && *dir != "" {
		err = g.GenerateDir(*dir, *outDir, *packageName, *specFilename, *basePath, *cfgFilename, *specHandler)
	} else {
		err = g.GenerateFile(*outDir, *packageName, *specFile, *basePath, *cfgFilename, *specHandler)
	}
	if err != nil {
		log.Fatalf("Error on generate: %v", err)
	}
}
