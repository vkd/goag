package generator

type AssignData struct {
	From, To Templater
	IsNew    bool
}

func AssignTemplate(from, to Templater, isNew bool) Templater {
	return TemplateData("Assign", AssignData{From: from, To: to, IsNew: isNew})
}

func ToSliceStrings(t Templater) Templater {
	return TemplateData("ToSliceStrings", t)
}
