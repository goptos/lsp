package documents

import "github.com/goptos/lsp/lsp"

type State struct {
	Documents map[lsp.DocumentUri]string
}

func NewState() State {
	return State{
		Documents: map[lsp.DocumentUri]string{},
	}
}

func (_self *State) OpenDocument(document lsp.DocumentUri, text string) {
	_self.Documents[document] = text
}

func (_self *State) UpdateDocument(document lsp.DocumentUri, changes []lsp.TextDocumentContentChangeEvent) {
	for _, change := range changes {
		_self.Documents[document] = change.Text
	}
}
