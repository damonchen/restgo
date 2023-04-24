package restgo

type Request struct {
	URL     string                 `yaml:"url"`
	Method  string                 `yaml:"method,omitempty"`
	JSON    map[string]interface{} `yaml:"json,omitempty"`
	Params  map[string]string      `yaml:"params,omitempty"`
	Data    map[string]string      `yaml:"data,omitempty"`
	Headers map[string]string      `yaml:"headers,omitempty"`
}

type ResponseSave struct {
	JSON map[string]string `yaml:"json"`
}

type Response struct {
	StatusCode          int                    `yaml:"status_code"`
	JSON                map[string]interface{} `yaml:"json"`
	Save                ResponseSave           `yaml:"save"`
	Headers             map[string]string      `yaml:"headers"`
	RedirectQueryParams map[string]string      `yaml:"redirect_query_params"`
}

type Stage struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description,omitempty"`
	Request     *Request  `yaml:"request"`
	Response    *Response `yaml:"response"`
}

// TestStage A file test
type TestStage struct {
	Description string  `yaml:"description"`
	Stages      []Stage `yaml:"stages"`
}
