package goag

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/tools/imports"

	"github.com/vkd/goag/generator"
	"github.com/vkd/goag/generator/render"
)

func GenerateDir(dir, packageName, specFilename, basePath string) error {
	err := ParseTemplates()
	if err != nil {
		return fmt.Errorf("parse templates: %w", err)
	}

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

		err = generateFile(testpath, packageName, specFile, basePath)
		if err != nil {
			return fmt.Errorf("generate: %w", err)
		}
	}

	return nil
}

func GenerateFile(outDir, packageName, specFilename, basePath string) error {
	err := ParseTemplates()
	if err != nil {
		return fmt.Errorf("parse templates: %w", err)
	}
	return generateFile(outDir, packageName, specFilename, basePath)
}

func generateFile(outDir, packageName, specFilename, basePath string) error {
	specRaw, err := ioutil.ReadFile(specFilename)
	if err != nil {
		return fmt.Errorf("read spec file: %w", err)
	}

	spec, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(specFilename)
	if err != nil {
		return fmt.Errorf("load spec: %w", err)
	}

	specBaseFilename := filepath.Base(specFilename)

	err = Generate(spec, outDir, packageName, specRaw, specBaseFilename, basePath)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	return nil
}

func Generate(spec *openapi3.Swagger, outDir string, packageName string, specRaw []byte, baseFilename, basePath string) error {
	components := generator.NewComponents(spec.Components)
	if len(components.Schemas) > 0 {
		goFile := generator.GoFile{
			PackageName: packageName,
			Renders:     []generator.Render{components},
		}
		err := RenderToFile(path.Join(outDir, "components.go"), goFile)
		if err != nil {
			return fmt.Errorf("generate components: %w", err)
		}
	}

	if basePath == "" && len(spec.Servers) > 0 {
		s := spec.Servers[0]
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

	hs, err := generator.NewHandlers(packageName, spec, basePath)
	if err != nil {
		return fmt.Errorf("generate handlers: %w", err)
	}

	err = ExecToFile("handler.gotmpl", path.Join(outDir, "handler.go"), hs)
	if err != nil {
		return fmt.Errorf("generate handler: %w", err)
	}

	r, err := generator.NewRouter(packageName, hs.Handlers, spec, specRaw, baseFilename, basePath)
	if err != nil {
		return fmt.Errorf("generate router: %w", err)
	}

	err = ExecToFile("router.gotmpl", path.Join(outDir, "router.go"), r)
	if err != nil {
		return fmt.Errorf("generate router: %w", err)
	}

	specFile := generator.SpecFile(packageName, specRaw)

	err = RenderToFile(path.Join(outDir, "spec_file.go"), specFile)
	if err != nil {
		return fmt.Errorf("generate spec_file: %w", err)
	}

	return nil
}

func ExecToFile(templateName string, filepath string, data interface{}) error {
	var bb bytes.Buffer
	err := tmps.ExecuteTemplate(&bb, templateName, data)
	if err != nil {
		return fmt.Errorf("error on exec template: %w", err)
	}

	return WriteToFile(bb.Bytes(), filepath)
}

func RenderToFile(filepath string, f generator.Render) error {
	bs, err := render.Bytes(f)
	if err != nil {
		return fmt.Errorf("to bytes: %w", err)
	}
	err = WriteToFile(bs, filepath)
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
