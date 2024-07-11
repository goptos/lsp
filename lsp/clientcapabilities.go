package lsp

type TokenFormat string

const (
	Relative TokenFormat = "relative"
)

type ClientCapabilities struct {
	Workspace        *CcWorkspace                        `json:"workspace,omitempty"`
	TextDocument     *TextDocumentClientCapabilities     `json:"textDocument,omitempty"`
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`
	Window           *CcWindow                           `json:"window,omitempty"`
	General          *CcGeneral                          `json:"general,omitempty"`
	Experimental     *LSPAny                             `json:"experimental,omitempty"`
}

type CcWorkspace struct {
	ApplyEdit *bool `json:"applyEdit,omitempty"`
	// WorkspaceEdit          *WorkspaceEditClientCapabilities           `json:"workspaceEdit,omitempty"`
	// DidChangeConfiguration *DidChangeConfigurationClientCapabilities  `json:"didChangeConfiguration,omitempty"`
	// DidChangeWatchedFiles  *DidChangeWatchedFilesClientCapabilities   `json:"didChangeWatchedFiles,omitempty"`
	// Symbol                 *WorkspaceSymbolClientCapabilities         `json:"symbol,omitempty"`
	// ExecuteCommand         *ExecuteCommandClientCapabilities          `json:"executeCommand,omitempty"`
	WorkspaceFolders *bool                                      `json:"workspaceFolders,omitempty"`
	Configuration    *bool                                      `json:"configuration,omitempty"`
	SemanticTokens   *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`
	// CodeLens               *CodeLensWorkspaceClientCapabilities       `json:"codeLens,omitempty"`
	FileOperations **CcFileOperations `json:"fileOperations,omitempty"`
	// InlineValue            *InlineValueWorkspaceClientCapabilities    `json:"inlineValue,omitempty"`
	// InlayHint              *InlayHintWorkspaceClientCapabilities      `json:"inlayHint,omitempty"`
	// Diagnostics            *DiagnosticWorkspaceClientCapabilities     `json:"diagnostics,omitempty"`
}

type CcFileOperations struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	DidCreate           *bool `json:"didCreate,omitempty"`
	WillCreate          *bool `json:"willCreate,omitempty"`
	DidRename           *bool `json:"didRename,omitempty"`
	WillRename          *bool `json:"willRename,omitempty"`
	DidDelete           *bool `json:"didDelete,omitempty"`
	WillDelete          *bool `json:"willDelete,omitempty"`
}

type CcWindow struct {
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`
	// ShowMessage      *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`
	// ShowDocument     *ShowDocumentClientCapabilities       `json:"showDocument,omitempty"`
}

type CcGeneral struct {
	StaleRequestSupport *CcStaleRequestSupport `json:"staleRequestSupport,omitempty"`
	// RegularExpressions  *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`
	// Markdown            *MarkdownClientCapabilities           `json:"markdown,omitempty"`
	// PositionEncodings   *[]PositionEncodingKind               `json:"positionEncodings,omitempty"`
}

type CcStaleRequestSupport struct {
	Cancel                 bool     `json:"cancel"`
	RetryOnContentModified []string `json:"retryOnContentModified"`
}

type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`
	// Completion         *CompletionClientCapabilities               `json:"completion,omitempty"`
	// Hover              *HoverClientCapabilities                    `json:"hover,omitempty"`
	// SignatureHelp      *SignatureHelpClientCapabilities            `json:"signatureHelp,omitempty"`
	// Declaration        *DeclarationClientCapabilities              `json:"declaration,omitempty"`
	// Definition         *DefinitionClientCapabilities               `json:"definition,omitempty"`
	// TypeDefinition     *TypeDefinitionClientCapabilities           `json:"typeDefinition,omitempty"`
	// Implementation     *ImplementationClientCapabilities           `json:"implementation,omitempty"`
	// References         *ReferenceClientCapabilities                `json:"references,omitempty"`
	// DocumentHighlight  *DocumentHighlightClientCapabilities        `json:"documentHighlight,omitempty"`
	// DocumentSymbol     *DocumentSymbolClientCapabilities           `json:"documentSymbol,omitempty"`
	// CodeAction         *CodeActionClientCapabilities               `json:"codeAction,omitempty"`
	// CodeLens           *CodeLensClientCapabilities                 `json:"codeLens,omitempty"`
	// DocumentLink       *DocumentLinkClientCapabilities             `json:"documentLink,omitempty"`
	// ColorProvider      *DocumentColorClientCapabilities            `json:"colorProvider,omitempty"`
	// Formatting         *DocumentFormattingClientCapabilities       `json:"formatting,omitempty"`
	// RangeFormatting    *DocumentRangeFormattingClientCapabilities  `json:"rangeFormatting,omitempty"`
	// OnTypeFormatting   *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`
	// Rename             *RenameClientCapabilities                   `json:"rename,omitempty"`
	// PublishDiagnostics *PublishDiagnosticsClientCapabilities       `json:"publishDiagnostics,omitempty"`
	// FoldingRange       *FoldingRangeClientCapabilities             `json:"foldingRange,omitempty"`
	// SelectionRange     *SelectionRangeClientCapabilities           `json:"selectionRange,omitempty"`
	// LinkedEditingRange *LinkedEditingRangeClientCapabilities       `json:"linkedEditingRange,omitempty"`
	// CallHierarchy      *CallHierarchyClientCapabilities            `json:"callHierarchy,omitempty"`
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`
	// Moniker            *MonikerClientCapabilities                  `json:"moniker,omitempty"`
	// TypeHierarchy      *TypeHierarchyClientCapabilities            `json:"typeHierarchy,omitempty"`
	// InlineValue        *InlineValueClientCapabilities              `json:"inlineValue,omitempty"`
	// InlayHint          *InlayHintClientCapabilities                `json:"inlayHint,omitempty"`
	// Diagnostic         *DiagnosticClientCapabilities               `json:"diagnostic,omitempty"`
}

type NotebookDocumentClientCapabilities struct {
	Synchronization NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}

type NotebookDocumentSyncClientCapabilities struct {
	DynamicRegistration     *bool `json:"dynamicRegistration,omitempty"`
	ExecutionSummarySupport *bool `json:"executionSummarySupport,omitempty"`
}

type SemanticTokensWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	WillSave            *bool `json:"willSave,omitempty"`
	WillSaveWaitUntil   *bool `json:"willSaveWaitUntil,omitempty"`
	DidSave             *bool `json:"didSave,omitempty"`
}

type SemanticTokensClientCapabilities struct {
	DynamicRegistration     *bool                                    `json:"dynamicRegistration,omitempty"`
	Requests                SemanticTokensClientCapabilitiesRequests `json:"requests"`
	TokenTypes              []string                                 `json:"tokenTypes"`
	TokenModifiers          []string                                 `json:"tokenModifiers"`
	Formats                 []TokenFormat                            `json:"formats"`
	OverlappingTokenSupport *bool                                    `json:"overlappingTokenSupport,omitempty"`
	MultilineTokenSupport   *bool                                    `json:"multilineTokenSupport,omitempty"`
	ServerCancelSupport     *bool                                    `json:"serverCancelSupport,omitempty"`
	AugmentsSyntaxTokens    *bool                                    `json:"augmentsSyntaxTokens,omitempty"`
}

type SemanticTokensClientCapabilitiesRequests struct {
	Range *bool  `json:"range,omitempty"`
	Full  *Delta `json:"full,omitempty"`
}

type Delta struct {
	Delta bool `json:"delta"`
}
