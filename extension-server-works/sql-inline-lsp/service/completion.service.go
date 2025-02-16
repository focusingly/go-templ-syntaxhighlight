package service

import (
	"sql-inline-lsp/model"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type provideService struct{}

var DefaultProvideService = &provideService{}

func (*provideService) ProvideCompletionItems() []protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	resp := []protocol.CompletionItem{}
	for _, word := range model.AllSQLKeywords {
		resp = append(resp, protocol.CompletionItem{
			Label:         word,
			InsertText:    &word,
			Kind:          &kind,
			Documentation: "SQL2016 word",
		})
	}

	return resp
}

type CompletionReqObj struct {
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
