package lsp

const (
	DocumentColor Method = "textDocument/documentColor"
)

type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
}

type DocumentColorRequest struct {
	Request
	Params DocumentColorParams `json:"params"`
}

type DocumentColorParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type DocumentColorResponse struct {
	Response
	Result []ColorInformation `json:"result"`
}

type ColorInformation struct {
	Range Range `json:"range"`
	Color Color `json:"color"`
}

type Color struct {
	Red   float32 `json:"red"`
	Green float32 `json:"green"`
	Blue  float32 `json:"blue"`
	Alpha float32 `json:"alpha"`
}

func NewDocumentColorOptions() *DocumentColorOptions {
	var progress = true
	return &DocumentColorOptions{
		WorkDoneProgressOptions: WorkDoneProgressOptions{
			WorkDoneProgress: &progress,
		},
	}

}

func NewDocumentColorResponse(req DocumentColorRequest) DocumentColorResponse {
	var colorInformation = []ColorInformation{}
	colorInformation = append(colorInformation, ColorInformation{
		Range: Range{
			Start: Position{Line: 71, Character: 3},
			End:   Position{Line: 72, Character: 6},
		},
		Color: Color{
			Red:   0.5,
			Green: 0.5,
			Blue:  0.5,
		},
	})
	return DocumentColorResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Result: colorInformation,
	}
}
