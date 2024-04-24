package generator

import "fmt"

type FileHandler struct {
	Imports []Import

	Handlers []*Handler

	IsWriteJSONFunc bool
	IsCustomMaybe   bool
}

func NewFileHandler(os []*Operation, basePathPrefix string, cfg Config) (zero FileHandler, _ error) {
	out := FileHandler{
		IsCustomMaybe: cfg.Maybe.Type != "",
	}
	for _, o := range os {
		h, ims, err := NewHandler(o, basePathPrefix, cfg)
		if err != nil {
			return zero, fmt.Errorf("handler %q: %w", o.Name, err)
		}
		out.Imports = append(out.Imports, ims...)
		out.Handlers = append(out.Handlers, h)

		for _, r := range h.Responses {
			if r.ContentJSON.IsSet {
				out.IsWriteJSONFunc = true
			}
		}
		if h.DefaultResponse != nil && h.DefaultResponse.ContentJSON.IsSet {
			out.IsWriteJSONFunc = true
		}
	}
	return out, nil
}

func (f FileHandler) Render() (string, error) {
	return ExecuteTemplate("FileHandler", f)
}
