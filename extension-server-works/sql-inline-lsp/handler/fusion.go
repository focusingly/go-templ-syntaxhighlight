package handler

import (
	"context"
	"encoding/json"
	"sql-inline-lsp/model"
	"sql-inline-lsp/service"
	"sql-inline-lsp/utility"

	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type fuseHandler struct {
}

var DefaultFusionHandler = &fuseHandler{}

var (
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
	normalPath, err := utility.ConvertFileURLToPath(parsedReq.TextDocument.URI)
	if err != nil {
		return
	}
	result, err := utility.FindSQLPositions(normalPath)
	if err != nil {
		return
	}
	if !utility.PosInScopes(result, parsedReq) {
		return
	}
	completions := completionService.ProvideCompletionItems()
	conn.Reply(ctx, req.ID, &protocol.CompletionList{
		IsIncomplete: false,
		Items:        completions,
	})
}
