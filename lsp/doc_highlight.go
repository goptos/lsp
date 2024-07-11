package lsp

const (
	TextDocumentHighlight Method = "textDocument/documentHighlight"
)

type DocumentHighlightKind int

const (
	Text  DocumentHighlightKind = 1
	Read  DocumentHighlightKind = 2
	Write DocumentHighlightKind = 3
)

type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

type DocumentHighlightRequest struct {
	Request
	Params DocumentHighlightParams `json:"params"`
}

type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

type PartialResultParams struct {
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}

type DocumentHighlightResponse struct {
	Response
	Result []DocumentHighlight `json:"result"`
}

type DocumentHighlight struct {
	Range Range                  `json:"range"`
	Kind  *DocumentHighlightKind `json:"kind,omitempty"`
}

func NewDocumentHighlightOptions() *DocumentHighlightOptions {
	var progress = true
	return &DocumentHighlightOptions{
		WorkDoneProgressOptions: WorkDoneProgressOptions{
			WorkDoneProgress: &progress,
		},
	}
}

func NewDocumentHighlightResponse(req DocumentHighlightRequest) DocumentHighlightResponse {
	var kind = Text
	var highlights []DocumentHighlight
	highlights = append(highlights, DocumentHighlight{
		Range: Range{
			Start: Position{
				Line:      72,
				Character: 3,
			},
			End: Position{
				Line:      72,
				Character: 6,
			},
		},
		Kind: &kind,
	})
	return DocumentHighlightResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Result: highlights,
	}
}
