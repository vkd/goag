package goag

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/tools/imports"

	generatorv2 "github.com/vkd/goag/generator"
	"github.com/vkd/goag/generator-v0"
	"github.com/vkd/goag/specification"
)

type Generator struct {
	GenClient bool
	DeleteOld bool
	DoNotEdit bool
}

func (g Generator) GenerateDir(dir, out, packageName, specFilename, basePath string) error {
	ts, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read dir %q: %w", dir, err)
	}

	for _, d := range ts {
		if !d.IsDir() {
			continue
		}
		testpath := filepath.Join(dir, d.Name())
		log.Printf("Generate file: %q", testpath)

		specFile := filepath.Join(testpath, specFilename)

		err = g.generateFile(filepath.Join(testpath, out), packageName, specFile, basePath)
		if err != nil {
			return fmt.Errorf("generate: %w", err)
		}
	}

	return nil
}

func (g Generator) GenerateFile(outDir, packageName, specFilename, basePath string) error {
	return g.generateFile(outDir, packageName, specFilename, basePath)
}

func (g Generator) generateFile(outDir, packageName, specFilename, basePath string) error {
	specRaw, err := os.ReadFile(specFilename)
	if err != nil {
		return fmt.Errorf("read spec file: %w", err)
	}

	openapi3Spec, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(specFilename)
	if err != nil {
		return fmt.Errorf("load spec: %w", err)
	}

	specBaseFilename := filepath.Base(specFilename)

	err = g.Generate(openapi3Spec, outDir, packageName, specRaw, specBaseFilename, basePath)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	return nil
}

func (g Generator) Generate(openapi3Spec *openapi3.Swagger, outDir string, packageName string, specRaw []byte, baseFilename, basePath string) error {
	s, err := specification.ParseSwagger(openapi3Spec)
	if err != nil {
		return fmt.Errorf("parse specification: %w", err)
	}

	componentsFile := path.Join(outDir, "components.go")
	err = os.Remove(componentsFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("remove components file (%s): %w", componentsFile, err)
		}
	}
	components, err := generatorv2.NewComponents(s.Components)
	if err != nil {
		return fmt.Errorf("generate 'components' file: %w", err)
	}
	if len(components.Schemas) > 0 {
		goFile := generatorv2.GoFile{
			SkipDoNotEdit: !g.DoNotEdit,
			PackageName:   packageName,
			Body:          components,
		}
		err := RenderToFile(componentsFile, generator.RenderFunc(goFile.Render))
		if err != nil {
			return fmt.Errorf("generate components: %w", err)
		}
	}

	if basePath == "" && len(openapi3Spec.Servers) > 0 {
		s := openapi3Spec.Servers[0]
		rawURL := s.URL
		for k, v := range s.Variables {
			if def, ok := v.Default.(string); ok {
				rawURL = strings.ReplaceAll(rawURL, "{"+k+"}", def)
			}
		}
		u, err := url.Parse(rawURL)
		if err != nil {
			return fmt.Errorf("parse servers.0.url: %w", err)
		}
		basePath = u.Path
	}

	hs, err := generator.NewHandlers(s, basePath)
	if err != nil {
		return fmt.Errorf("generate handlers: %w", err)
	}

	var handlers []generatorv2.Render
	for _, h := range hs.Handlers {
		handlers = append(handlers, h.HandlerOld)
	}

	gen, err := generatorv2.NewGenerator(s, generatorv2.PackageName(packageName), generatorv2.SkipDoNotEdit(), generatorv2.BasePath(basePath), generatorv2.SpecFilename(baseFilename))
	if err != nil {
		return fmt.Errorf("create new generator from spec file: %w", err)
	}

	hFile, err := gen.FileHandlerTemplater(handlers, hs.IsWriteJSONFunc)
	if err != nil {
		return fmt.Errorf("generate handlers file: %w", err)
	}

	err = RenderToFile(path.Join(outDir, "handler.go"), hFile)
	if err != nil {
		return fmt.Errorf("generate handler: %w", err)
	}

	r, err := generator.NewRouter(packageName, hs.Handlers, openapi3Spec, specRaw, baseFilename, basePath)
	if err != nil {
		return fmt.Errorf("generate router: %w", err)
	}

	rFile, err := gen.RouterFile(basePath, baseFilename, r)
	if err != nil {
		return fmt.Errorf("generate router file: %w", err)
	}

	err = RenderToFile(path.Join(outDir, "router.go"), rFile)
	if err != nil {
		return fmt.Errorf("generate router.go: %w", err)
	}

	specFile := generator.SpecFile(packageName, specRaw)

	err = RenderToFile(path.Join(outDir, "spec_file.go"), specFile)
	if err != nil {
		return fmt.Errorf("generate spec_file: %w", err)
	}

	clientFile := path.Join(outDir, "client.go")
	err = os.Remove(clientFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("remove client file (%s): %w", clientFile, err)
		}
	}
	if g.GenClient {
		clientFile := gen.ClientFile()

		err = RenderToFile(path.Join(outDir, "client.go"), generatorv2.TemplaterFunc(clientFile.Render))
		if err != nil {
			return fmt.Errorf("generate client source: %w", err)
		}
	}

	return nil
}

func RenderToFile(filepath string, f generator.Render) error {
	s, err := f.String()
	if err != nil {
		return fmt.Errorf("to bytes: %w", err)
	}
	err = WriteToFile([]byte(s), filepath)
	if err != nil {
		return fmt.Errorf("write to file: %w", err)
	}
	return nil
}

func WriteToFile(bs []byte, filepath string) error {
	dirpath := path.Dir(filepath)
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("mkdir all: %w", err)
	}

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error on open file: %w", err)
	}

	importedBs, err := imports.Process("", bs, nil)
	// bs, err := format.Source(bb.Bytes())
	if err != nil {
		// return fmt.Errorf("error on format go source: %w", err)
		log.Printf("Error on format go source (%s): %v", filepath, err)
	} else {
		bs = importedBs
	}

	_, err = f.Write(bs)
	if err != nil {
		return fmt.Errorf("error on copy file content: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("error on close file: %w", err)
	}
	return nil
}
