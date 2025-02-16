package handler

import (
	"context"
	"encoding/json"
	"os"
	"sql-inline-lsp/model"
	"sql-inline-lsp/service"

	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type fuseHandler struct {
}

var DefaultFusionHandler = &fuseHandler{}

var (
	fileService       = service.DefaultFileService
	extractService    = service.DefaultExtractService
	completionService = service.DefaultProvideService
)

func (f *fuseHandler) ParseCompletionRequest(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	bf, err := req.Params.MarshalJSON()
	if err != nil {
		return
	}
	parsedReq := &model.ClientCompletionReq{}
	if err := json.Unmarshal(bf, &parsedReq); err != nil {
		return
	}
	url, ok := fileService.ParseFileProtocolURL(parsedReq.TextDocument.URI)
	if !ok {
		return
	}

	bf2, err := os.ReadFile(url)
	if err != nil {
		return
	}
	positions := extractService.GetAllRawSqlCommentedStringPos(string(bf2))
	if !extractService.PositionInRange(parsedReq, positions) {
		return
	}
	completions := completionService.ProvideCompletionItems()
	conn.Reply(ctx, req.ID, &protocol.CompletionList{
		IsIncomplete: false,
		Items:        completions,
	})
}
