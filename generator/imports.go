package generator

var customImports []string

func AddImport(i string) {
	if i == "" {
		return
	}
	for _, imp := range customImports {
		if imp == i {
			return
		}
	}
	customImports = append(customImports, i)
}

func CustomImports() []string { return customImports }
