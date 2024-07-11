package main

import (
	"os"

	"github.com/goptos/lsp/documents"
	"github.com/goptos/lsp/handler"
	"github.com/goptos/lsp/logger"
	"github.com/goptos/lsp/rpc"
)

func main() {
	logger := logger.New("C:/Users/algonz/repo/goptos/lsp/goptoslsp.log")
	logger.Println("Started")
	state := documents.NewState()
	handler := handler.New(&state)
	connection := rpc.NewConnection(logger, os.Stdin, os.Stdout, handler.HandleMessage)
	for connection.Open() {
		connection.Receive()
	}
}
