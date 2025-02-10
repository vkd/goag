package generator

import "strings"

func (g *Generator) ClientFile(cfg Config) GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       append(cfg.Imports(), g.Imports...),
		Body:          g.Client,
	}
}

func (g *Generator) ComponentsFile(cfg Config) GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       append(cfg.Imports(), g.Imports...),
		Body:          g.Components,
	}
}

func (g *Generator) HandlerFile(cfg Config) GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       append(cfg.Imports(), g.Imports...),
		Body:          g.HandlersFile,
	}
}

func (g *Generator) RouterFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       nil,
		Body:          g.Router,
	}
}

func (g *Generator) SpecFile(fileContent []byte) GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Body: RenderFunc(func() (string, error) {
			return "const SpecFile string = " + encodeRawFileAsString(string(fileContent)), nil
		}),
	}
}

func encodeRawFileAsString(s string) string {
	if strings.Contains(string(s), "\n") {
		s = "`" + strings.ReplaceAll(string(s), "`", "`+\"`\"+`") + "`"
	} else {
		s = `"` + strings.ReplaceAll(string(s), `"`, `\"`) + `"`
	}
	return s
}
