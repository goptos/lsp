package handler

import (
	"encoding/json"

	"github.com/goptos/lsp/documents"
	"github.com/goptos/lsp/lsp"
	"github.com/goptos/lsp/rpc"
)

func receive[T any](connection *rpc.Connection, content []byte) T {
	var req T
	err := json.Unmarshal(content, &req)
	if err != nil {
		connection.Log("Couldn't parse content: %s", err)
	}
	return req
}

type Handler struct {
	state documents.State
}

func New(state *documents.State) *Handler {
	return &Handler{
		state: *state,
	}
}

func (_self *Handler) HandleMessage(connection *rpc.Connection, method lsp.Method, content []byte) {
	switch method {

	case lsp.Cancel:
		var req = receive[lsp.CancelRequest](connection, content)
		connection.Send(lsp.NewCancelResponse(req))

	case lsp.Shutdown:
		connection.Close()

	case lsp.Initialize:
		var req = receive[lsp.InitializeRequest](connection, content)
		connection.Log("Connected to: %s %s", req.Params.ClientInfo.Name, req.Params.ClientInfo.Version)
		connection.Log("Client SemanticTokensClientCapabilities capabilities: %v", *req.Params.Capabilities.TextDocument.SemanticTokens)
		connection.Send(lsp.NewInitializeResponse(req))

	case lsp.DocumentColor:
		var req = receive[lsp.DocumentColorRequest](connection, content)
		connection.Send(lsp.NewDocumentColorResponse(req))

	case lsp.TextDocumentHover:
		var req = receive[lsp.HoverRequest](connection, content)
		connection.Send(lsp.NewHoverResponse(req))

	case lsp.TextDocumentHighlight:
		var req = receive[lsp.DocumentHighlightRequest](connection, content)
		connection.Send(lsp.NewDocumentHighlightResponse(req))

	case lsp.TextDocumentSemanticTokensFull:
		var req = receive[lsp.SemanticTokensRequest](connection, content)
		connection.Send(lsp.NewSemanticTokensResponse(req))

	case lsp.TextDocumentDidOpen:
		var req = receive[lsp.DidOpenTextDocumentNotification](connection, content)
		_self.state.OpenDocument(req.Params.TextDocument.Uri, req.Params.TextDocument.Text)
		connection.Log("Opened: %s", req.Params.TextDocument.Uri)

	case lsp.TextDocumentDidChange:
		var req = receive[lsp.DidChangeTextDocumentNotification](connection, content)
		_self.state.UpdateDocument(req.Params.TextDocument.Uri, req.Params.ContentChanges)
		connection.Log("Updated: %s", req.Params.TextDocument.Uri)
	}
}
