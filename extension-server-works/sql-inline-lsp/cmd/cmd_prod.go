//go:build stdProd
// +build stdProd

package cmd

import (
	"context"
	"sql-inline-lsp/server"
	"sql-inline-lsp/utility"

	"github.com/sourcegraph/jsonrpc2"
)

func Run() {
	lspServerInstance := &server.LspServer{}
	stdioBuf := utility.CreateStdReadWriteCloser()
	stream := jsonrpc2.NewBufferedStream(stdioBuf, jsonrpc2.VSCodeObjectCodec{})
	rpcConn := jsonrpc2.NewConn(context.Background(), stream, lspServerInstance)
	<-rpcConn.DisconnectNotify()
}
