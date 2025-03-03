package apidoc

type ParamType string
type Method string

const (
	MethodGet    Method = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPatch         = "PATCH"

	ParamTypePath ParamType = "path"
	ParamTypeBody           = "body"
	ParamTypeForm           = "formData"
)

type APIDocumentation struct {
	Info      Info                 `json:"info"`
	Endpoints map[string]*Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Summary    string     `json:"summary"`
	Method     Method     `json:"method"`
	Params     []Param    `json:"params,omitempty"`
	Responses  []Response `json:"responses"`
	Tag        string     `json:"tag"`
	Deprecated bool       `json:"deprecated"`
}

type Response struct {
	StatusCode  int         `json:"statusCode"`
	Description string      `json:"description"`
	Schema      interface{} `json:"schema"`
	Examples    []Example   `json:"examples"`
}

type Example struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Param struct {
	Name        string      `json:"name"`
	Schema      interface{} `json:"schema"`
	ParamType   ParamType   `json:"paramType"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
}

type Info struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Contact     InfoContact `json:"contact"`
	Version     string      `json:"version"`
}

type InfoContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}
