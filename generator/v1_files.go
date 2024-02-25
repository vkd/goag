package generator

func (g *Generator) ClientFile() GoFile {
	return GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Body:          g.Client,
	}
}
