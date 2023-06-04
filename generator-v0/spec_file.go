package generator

func SpecFile(packageName string, fileContent []byte) GoFile {
	return GoFile{
		PackageName: packageName,
		Renders: []Render{
			GoConstDef{
				Name:  "SpecFile",
				Type:  StringType,
				Value: StringValue(fileContent),
			},
		},
	}
}
