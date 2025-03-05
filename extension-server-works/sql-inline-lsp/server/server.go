package server

import (
	"context"
	"fmt"
	"os"
	"sql-inline-lsp/handler"

	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var (
	serverName    = "Go-Templ-LSP"
	serverVersion = "0.0.1"
)

type LspServer struct {
}

// Handle implements jsonrpc2.Handler.
func (l *LspServer) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	fusionHandler := handler.DefaultFusionHandler
	defer func() {
		if err := recover(); err != nil {
			conn.ReplyWithError(ctx, req.ID, &jsonrpc2.Error{
				Code:    1,
				Message: fmt.Sprintf("resolve %s cause a error: %v", req.Method, err),
			})
		}
	}()

	switch req.Method {
	case protocol.MethodInitialize:
		resp := protocol.InitializeResult{
			Capabilities: protocol.ServerCapabilities{
				CodeLensProvider: &protocol.CodeLensOptions{
					ResolveProvider: &protocol.True,
				},
				Workspace: &protocol.ServerCapabilitiesWorkspace{},
				CompletionProvider: &protocol.CompletionOptions{
					ResolveProvider:   &protocol.True,
					TriggerCharacters: []string{},
				},
			},
			ServerInfo: &protocol.InitializeResultServerInfo{
				Name:    serverName,
				Version: &serverVersion,
			},
		}
		conn.Reply(ctx, req.ID, &resp)
	case protocol.MethodInitialized:
		resp := protocol.InitializeResult{
			ServerInfo: &protocol.InitializeResultServerInfo{
				Name:    serverName,
				Version: &serverVersion,
			},
		}
		conn.Reply(ctx, req.ID, resp)
		conn.Notify(context.TODO(), protocol.MethodSetTrace, "init successfully")
	case protocol.MethodTextDocumentCompletion:
		fusionHandler.ParseCompletionRequest(ctx, conn, req)
	case protocol.MethodCodeLensResolve:
	case protocol.ServerWorkspaceCodeLensRefresh:
	case protocol.MethodExit:
		conn.Reply(ctx, req.ID, 0)
	case protocol.MethodShutdown:
		conn.Close()
		os.Exit(0)
	default:
	}

}

var _ jsonrpc2.Handler = (*LspServer)(nil)
