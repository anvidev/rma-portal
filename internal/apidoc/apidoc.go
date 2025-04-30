package apidoc

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS
var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		panic("failed to parse api documentation templates: " + err.Error())
	}
}

func NewDocumentation(info Info) *APIDocumentation {
	return &APIDocumentation{
		Info: info,
		Tags: []*Tag{},
	}
}

func (docs *APIDocumentation) AddTag(name string) *Tag {
	tag := &Tag{Name: name}
	docs.Tags = append(docs.Tags, tag)
	return tag
}

type EndpointOption func(e *Endpoint)

func (tag *Tag) AddEndpoint(method Method, path, summary string, opts ...EndpointOption) {
	endpoint := Endpoint{
		Path:    path,
		Summary: summary,
		Method:  method,
	}

	for _, opt := range opts {
		opt(&endpoint)
	}

	tag.Endpoints = append(tag.Endpoints, endpoint)
}

func WithBody(data any) EndpointOption {
	return func(e *Endpoint) {
		e.Body = structToSlice(data)
	}
}

func (end *Endpoint) WithForm(formdata FormField) *Endpoint {
	end.Form = append(end.Form, formdata)
	return end
}

func (end *Endpoint) WithDeprecate(deprecated bool) *Endpoint {
	end.Deprecated = deprecated
	return end
}

type ParamOption func(q *QueryParam)

func WithRequired(b bool) ParamOption {
	return func(q *QueryParam) {
		q.Validation.Required = b
	}
}

func WithDefault(i any) ParamOption {
	return func(q *QueryParam) {
		q.Validation.Default = i
	}
}

func WithMin(min int) ParamOption {
	return func(q *QueryParam) {
		q.Validation.Min = min
	}
}

func WithMax(max int) ParamOption {
	return func(q *QueryParam) {
		q.Validation.Max = max
	}
}

func WithQuery(name, typ string, opts ...ParamOption) EndpointOption {
	return func(e *Endpoint) {
		query := QueryParam{Name: name, Type: typ}
		for _, opt := range opts {
			opt(&query)
		}
		e.Query = append(e.Query, query)
	}
}

func (d *APIDocumentation) Serve(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "base.html", d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
