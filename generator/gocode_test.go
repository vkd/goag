package generator

import (
	"testing"
)

func TestGoStruct_String(t *testing.T) {
	type fields struct {
		Fields []GoStructField
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, `struct{}`},
		{"one field", fields{Fields: []GoStructField{
			{Name: "ID", Type: Int},
		}}, `struct{
	ID int
}`},
		{"two fields", fields{Fields: []GoStructField{
			{Name: "ID", Type: Int},
			{Name: "Age", Type: Int32},
		}}, `struct{
	ID int
	Age int32
}`},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := GoStruct{
				Fields: tt.fields.Fields,
			}
			got, err := g.String()
			if err != nil {
				t.Errorf("GoStruct.String() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("GoStruct.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGoStructField_String(t *testing.T) {
	type fields struct {
		Name    string
		Comment string
		Type    Render
		Tags    []GoFieldTag
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"no tags", fields{Name: "ID", Type: Int}, `ID int`},
		{"one tags", fields{Name: "ID", Type: Int, Tags: []GoFieldTag{
			{"json", "id"},
		}}, `ID int ` + "`" + `json:"id"` + "`"},
		{"two tags", fields{Name: "ID", Type: Int, Tags: []GoFieldTag{
			{"json", "id"},
			{"yaml", "id"},
		}}, `ID int ` + "`" + `json:"id" yaml:"id"` + "`"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			g := GoStructField{
				Name:    tt.fields.Name,
				Comment: tt.fields.Comment,
				Type:    tt.fields.Type,
				Tags:    tt.fields.Tags,
			}
			got, err := g.String()
			if err != nil {
				t.Errorf("GoStructField.String() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("GoStructField.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FileImports(is ...string) (out []GoFileImport) {
	for _, i := range is {
		out = append(out, GoFileImport{Package: i})
	}
	return out
}

func TestGoFile_String(t *testing.T) {
	tests := []struct {
		name string
		file GoFile
		want string
	}{
		{"no imports", GoFile{PackageName: "test"}, "package test\n"},
		{"one import", GoFile{PackageName: "test", Imports: FileImports("fmt")}, "package test\n\nimport (\n\t\"fmt\"\n)\n"},
		{"two imports", GoFile{PackageName: "test", Imports: FileImports("fmt", "io")}, "package test\n\nimport (\n\t\"fmt\"\n\t\"io\"\n)\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.file.String()
			if err != nil {
				t.Errorf("GoFile.String() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("GoFile.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
