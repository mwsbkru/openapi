package openapi

// RequestBody -.
type RequestBody struct {
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// RequestBodies -.
type RequestBodies map[string]*RequestBody
