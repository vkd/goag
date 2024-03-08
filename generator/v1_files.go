package generator

func (g *Generator) ClientFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Body:          g.Client,
	}
}

func (g *Generator) RouterFile(basePath, baseFilename string, oldRouter any) (Templater, error) {
	return g.goFile(nil, g.Router), nil
}
