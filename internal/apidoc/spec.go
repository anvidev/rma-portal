package apidoc

type Method string

const (
	MethodGet    Method = "GET"
	MethodPost          = "POST"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	MethodPatch         = "PATCH"
)

type APIDocumentation struct {
	Info Info   `json:"info"`
	Tags []*Tag `json:"tags"`
}

type Tag struct {
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path    string       `json:"path"`
	Summary string       `json:"summary"`
	Method  Method       `json:"method"`
	Query   []QueryParam `json:"queries"`
	Form    []FormField  `json:"formdata"`
	Body    []BodyField  `json:"body"`
	// Responses  []Response `json:"responses"`
	Deprecated bool `json:"deprecated"`
}

// type Response struct {
// 	StatusCode  int         `json:"statusCode"`
// 	Description string      `json:"description"`
// 	Schema      interface{} `json:"schema"`
// 	Examples    []Example   `json:"examples"`
// }

// type Example struct {
// 	Name  string      `json:"name"`
// 	Value interface{} `json:"value"`
// }

type BodyField struct {
	Name        string
	Type        string
	JSONName    string
	Description string
	Required    bool
	Validation  map[string]string
	Fields      []BodyField
}

type FormField struct {
	Name        string `json:"name"`
	Schema      any    `json:"schema"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type QueryParam struct {
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Validation  ParamValidation `json:"validation"`
}

type ParamValidation struct {
	Required bool `json:"required"`
	Default  any  `json:"default"`
	Min      int  `json:"min"`
	Max      int  `json:"max"`
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
