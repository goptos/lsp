package lsp

const (
	TextDocumentDidOpen           Method = "textDocument/didOpen"
	TextDocumentDidChange         Method = "textDocument/didChange"
	TextDocumentWillSave          Method = "textDocument/willSave"
	TextDocumentWillSaveWaitUntil Method = "textDocument/willSaveWaitUntil"
)

type TextDocumentSyncKind int

const (
	None        TextDocumentSyncKind = 0
	Full        TextDocumentSyncKind = 1
	Incremental TextDocumentSyncKind = 2
)

type TextDocumentSaveReason int

const (
	Manual     TextDocumentSaveReason = 1
	AfterDelay TextDocumentSaveReason = 2
	FocusOut   TextDocumentSaveReason = 3
)

type TextDocumentSyncOptions struct {
	OpenClose *bool                 `json:"openClose,omitempty"`
	Change    *TextDocumentSyncKind `json:"change,omitempty"`
}

type DidOpenTextDocumentNotification struct {
	Method string                    `json:"method"`
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type TextDocumentItem struct {
	Uri        DocumentUri `json:"uri"`
	LanguageId string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

type DocumentUri string
type Uri string

type DidChangeTextDocumentNotification struct {
	Method string                      `json:"method"`
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentIdentifier struct {
	Uri DocumentUri `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentContentChangeEvent struct {
	Range       *Range `json:"range,omitempty"`
	RangeLength *uint  `json:"rangeLength,omitempty"`
	Text        string `json:"text"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

type WillSaveTextDocumentNotification struct {
	Method string                     `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}

type WillSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Reason       TextDocumentSaveReason `json:"reason"`
}

type WillSaveWaitUntilTextDocumentRequest struct {
	Request
	Method string                     `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}
