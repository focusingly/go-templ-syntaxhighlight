package model

type (
	ClientCompletionReq struct {
		Context struct {
			TriggerKind int `json:"triggerKind" yaml:"triggerKind" xml:"triggerKind" toml:"triggerKind"`
		} `json:"context" yaml:"context" xml:"context" toml:"context"`
		Position struct {
			Character int `json:"character" yaml:"character" xml:"character" toml:"character"`
			Line      int `json:"line" yaml:"line" xml:"line" toml:"line"`
		} `json:"position" yaml:"position" xml:"position" toml:"position"`
		TextDocument struct {
			URI string `json:"uri" yaml:"uri" xml:"uri" toml:"uri"`
		} `json:"textDocument" yaml:"textDocument" xml:"textDocument" toml:"textDocument"`
	}
)
