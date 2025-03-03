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
		Info:      info,
		Endpoints: make(map[string]*Endpoint),
	}
}

func (d *APIDocumentation) AddEndpoint(path string, method Method) *Endpoint {
	endpoint := &Endpoint{Method: method}
	d.Endpoints[path] = endpoint
	return endpoint
}

func (e *Endpoint) WithSummary(summary string) *Endpoint {
	e.Summary = summary
	return e
}

func (e *Endpoint) WithParam(param Param) *Endpoint {
	e.Params = append(e.Params, param)
	return e
}

func (e *Endpoint) WithResponse(response Response) *Endpoint {
	e.Responses = append(e.Responses, response)
	return e
}

func (e *Endpoint) WithTag(tag string) *Endpoint {
	e.Tag = tag
	return e
}

func (e *Endpoint) WithDeprecated(deprecated bool) *Endpoint {
	e.Deprecated = deprecated
	return e
}

func (d *APIDocumentation) Serve(w http.ResponseWriter, r *http.Request) {
	endpointsByTags := make(map[string][]*Endpoint)

	for _, e := range d.Endpoints {
		endpointsByTags[e.Tag] = append(endpointsByTags[e.Tag], e)
	}

	data := struct {
		Info      Info
		Endpoints map[string][]*Endpoint
	}{
		Info:      d.Info,
		Endpoints: endpointsByTags,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
