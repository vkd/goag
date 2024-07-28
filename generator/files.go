package generator

import "strings"

func (g *Generator) ClientFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       nil,
		Body:          g.Client,
	}
}

func (g *Generator) ComponentsFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       g.Imports,
		Body:          g.Components,
	}
}

func (g *Generator) HandlerFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       append(g.FileHandler.Imports, g.Imports...),
		Body:          g.FileHandler,
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
