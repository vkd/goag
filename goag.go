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

	"github.com/vkd/goag/generator"
	"github.com/vkd/goag/specification"
)

type Generator struct {
	GenClient     bool
	GenAPIHandler bool
	DoNotEdit     bool
}

func (g Generator) GenerateDir(dir, out, packageName, specFilename, basePath, cfgFilename, specHandlerName string) error {
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

		cfgFile := filepath.Join(testpath, cfgFilename)

		err = g.generateFile(filepath.Join(testpath, out), packageName, specFile, basePath, cfgFile, specHandlerName)
		if err != nil {
			return fmt.Errorf("generate %q: %w", d.Name(), err)
		}
	}

	return nil
}

func (g Generator) GenerateFile(outDir, packageName, specFilename, basePath, cfgFilename, handlerSpecFilename string) error {
	return g.generateFile(outDir, packageName, specFilename, basePath, cfgFilename, handlerSpecFilename)
}

func (g Generator) generateFile(outDir, packageName, specFilename, basePath, cfgFilename, handlerSpecFilename string) error {
	specRaw, err := os.ReadFile(specFilename)
	if err != nil {
		return fmt.Errorf("read spec file: %w", err)
	}

	openapi3Spec, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(specFilename)
	if err != nil {
		return fmt.Errorf("load spec: %w", err)
	}

	if handlerSpecFilename == "" {
		handlerSpecFilename = filepath.Base(specFilename)
	}

	cfg, err := generator.LoadConfig(cfgFilename)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	err = g.Generate(openapi3Spec, outDir, packageName, specRaw, handlerSpecFilename, basePath, cfg)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	return nil
}

func (g Generator) Generate(openapi3Spec *openapi3.Swagger, outDir string, packageName string, specRaw []byte, baseFilename, basePath string, cfg generator.Config) error {
	s, err := specification.ParseSwagger(openapi3Spec, specification.SchemaOptions{
		IgnoreCustomType: cfg.CustomTypes.Ignore,
	})
	if err != nil {
		return fmt.Errorf("parse specification: %w", err)
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

	gen, err := generator.NewGenerator(s,
		cfg,
		generator.PackageName(packageName),
		generator.IfOption(generator.SkipDoNotEdit(), !g.DoNotEdit),
		generator.BasePath(basePath),
		generator.SpecFilename(baseFilename),
	)
	if err != nil {
		return fmt.Errorf("create new generator from spec file: %w", err)
	}

	if gen.Components.LenToRender() > 0 {
		err := RenderToFile(path.Join(outDir, "components.go"), gen.ComponentsFile(cfg))
		if err != nil {
			return fmt.Errorf("generate components: %w", err)
		}
	} else {
		componentsFile := path.Join(outDir, "components.go")
		err = os.Remove(componentsFile)
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("remove components.go (%s): %w", componentsFile, err)
			}
		}
	}

	if g.GenAPIHandler {
		err = RenderToFile(path.Join(outDir, "handler.go"), gen.HandlerFile(cfg))
		if err != nil {
			return fmt.Errorf("generate handler.go: %w", err)
		}

		err = RenderToFile(path.Join(outDir, "router.go"), gen.RouterFile())
		if err != nil {
			return fmt.Errorf("generate router.go: %w", err)
		}

		err = RenderToFile(path.Join(outDir, "spec_file.go"), gen.SpecFile(specRaw))
		if err != nil {
			return fmt.Errorf("generate spec_file.go: %w", err)
		}
	} else {
		err = os.Remove(path.Join(outDir, "handler.go"))
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("remove handler.go (%s): %w", path.Join(outDir, "handler.go"), err)
			}
		}
		err = os.Remove(path.Join(outDir, "router.go"))
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("remove router.go (%s): %w", path.Join(outDir, "router.go"), err)
			}
		}
		err = os.Remove(path.Join(outDir, "spec_file.go"))
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("remove spec_file.go (%s): %w", path.Join(outDir, "spec_file.go"), err)
			}
		}
	}

	clientFile := path.Join(outDir, "client.go")
	err = os.Remove(clientFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("remove client file (%s): %w", clientFile, err)
		}
	}
	if g.GenClient {
		err = RenderToFile(path.Join(outDir, "client.go"), gen.ClientFile(cfg))
		if err != nil {
			return fmt.Errorf("generate client.go: %w", err)
		}
	}

	return nil
}

func RenderToFile(filepath string, f generator.GoFile) error {
	s, err := f.Render()
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
