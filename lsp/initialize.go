package lsp

const (
	Initialize  Method = "initialize"
	Initialized Method = "initialized"
)

type TraceValue string

const (
	Off      TraceValue = "off"
	Messages TraceValue = "messages"
	Verbose  TraceValue = "verbose"
)

type InitializeRequest struct {
	Request
	Params InitializeParams `json:"params"`
}

type InitializeParams struct {
	WorkDoneProgressParams
	ProcessId             int                `json:"processId"`
	ClientInfo            *ClientInfo        `json:"clientInfo"`
	Locale                *string            `json:"locale,omitempty"`
	RootPath              *string            `json:"rootPath,omitempty"`
	RootUri               DocumentUri        `json:"rootUri"`
	InitializationOptions *LSPAny            `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 *TraceValue        `json:"trace,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type WorkspaceFolder struct {
	Uri  Uri    `json:"Uri"`
	Name string `json:"name"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ServerInfo        `json:"serverInfo,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	TextDocumentSync          *TextDocumentSyncKind     `json:"textDocumentSync,omitempty"`
	SemanticTokensProvider    *SemanticTokensOptions    `json:"semanticTokensProvider,omitempty"`
	DocumentHighlightProvider *DocumentHighlightOptions `json:"documentHighlightProvider,omitempty"`
	ColorProvider             *DocumentColorOptions     `json:"colorProvider,omitempty"`
}

func NewInitializeResponse(req InitializeRequest) InitializeResponse {
	var full = Full
	return InitializeResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:       &full,
				SemanticTokensProvider: NewSemanticTokensOptions(),
				// DocumentHighlightProvider: NewDocumentHighlightOptions(),
				// ColorProvider:             NewDocumentColorOptions(),
			},
			ServerInfo: &ServerInfo{
				Name:    "goptos",
				Version: "0.0.1",
			},
		},
	}
}
