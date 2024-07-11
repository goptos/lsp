package lsp

import "fmt"

const (
	TextDocumentHover Method = "textDocument/hover"
)

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type WorkDoneProgressParams struct {
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

type ProgressToken int

type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

type HoverResponse struct {
	Response
	Result Hover `json:"result"`
}

type Hover struct {
	Contents []MarkedString `json:"contents"`
	Range    *Range         `json:"range,omitempty"`
}

type MarkedString string

func NewHoverResponse(req HoverRequest) HoverResponse {
	var document = req.Params.TextDocument.Uri
	var position = req.Params.Position
	var contents []MarkedString
	contents = append(contents, "Hello here is your hover response: ")
	contents = append(contents, MarkedString(fmt.Sprintf("%d %d %s",
		position.Line, position.Character, document)))
	return HoverResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Result: Hover{
			Contents: contents,
		},
	}
}
