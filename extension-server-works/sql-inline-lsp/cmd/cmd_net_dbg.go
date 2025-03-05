//go:build !stdProd
// +build !stdProd

package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"sql-inline-lsp/server"

	"github.com/sourcegraph/jsonrpc2"
)

func Run() {
	lspServerInstance := &server.LspServer{}
	fmt.Println("hello world")
	c, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		if conn, err := c.Accept(); err != nil {
			fmt.Println(err)
			continue
		} else {
			go func(conn net.Conn) {
				defer conn.Close()
				stream := jsonrpc2.NewBufferedStream(conn, &jsonrpc2.VSCodeObjectCodec{})
				rpcConn := jsonrpc2.NewConn(context.Background(), stream, lspServerInstance)
				<-rpcConn.DisconnectNotify()
				log.Printf("Go-Templ: %s was closed", conn.RemoteAddr().String())
			}(conn)
		}
	}
}
