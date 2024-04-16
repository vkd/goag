package generator

import "strings"

func (g *Generator) ClientFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Body:          g.Client,
	}
}

func (g *Generator) ComponentsFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Body:          g.Components,
	}
}

func (g *Generator) HandlerFile() (GoFile, error) {
	return g.goFile(g.Imports, g.FileHandler), nil
}

func (g *Generator) RouterFile(basePath, baseFilename string) (GoFile, error) {
	return g.goFile(nil, g.Router), nil
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

func (g *Generator) goFile(ims []Import, body Render) GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       ims,
		Body:          body,
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
