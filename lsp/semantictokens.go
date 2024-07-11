package lsp

func generateOrderedArrayFromIndexedMap(m map[string]int) []string {
	var tmp = make(map[int]string)
	for k, v := range m {
		tmp[v] = k
	}
	var a = []string{}
	for i := 0; i < len(tmp); i++ {
		a = append(a, tmp[i])
	}
	return a
}

const (
	TextDocumentSemanticTokensFull      Method = "textDocument/semanticTokens/full"
	TextDocumentSemanticTokensFullDelta Method = "textDocument/semanticTokens/full/delta"
)

var TokenSyntax = map[string]string{
	"View":                 "keyword",
	"StartTag":             "label",
	"EndTag":               "decorator",
	"Component":            "function",
	"Comment":              "comment",
	"AttributeName":        "namespace",
	"AttributeValue":       "string",
	"KeywordAttributeName": "",
	"DynamicAttributeName": "",
	"EventAttributeName":   "",
	"AttributeEffect":      "",
	"ArgumentAttribute":    "variable",
	"Text":                 "comment",
	"Code":                 "variable",
	"EndOfFile":            "comment",
}

var TokenTypes = map[string]int{
	"keyword":  0,
	"type":     1,
	"comment":  2,
	"string":   3,
	"variable": 4,
	"function": 5,
}

var TokenModifiers = map[string]int{
	"Component": 1,
}

type SemanticTokensOptions struct {
	WorkDoneProgressOptions
	Legend SemanticTokensLegend `json:"legend"`
	Range  *bool                `json:"range,omitempty"`
	Full   *bool                `json:"full,omitempty"`
}

type SemanticTokensLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`
}

type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

type TextDocumentRegistrationOptions struct {
	DocumentSelector DocumentSelector `json:"documentSelector"`
}

type DocumentSelector = []DocumentFilter

type DocumentFilter struct {
	Language *string `json:"language,omitempty"`
	Scheme   *string `json:"scheme,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
}

type StaticRegistrationOptions struct {
	Id *string `json:"id,omitempty"`
}

type SemanticTokensRequest struct {
	Request
	Params SemanticTokensParams `json:"params"`
}

type SemanticTokensParams struct {
	WorkDoneProgressParams
	PartialResultParams
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type SemanticTokensResponse struct {
	Response
	Result SemanticTokens `json:"result"`
}

type SemanticTokens struct {
	ResultId *string `json:"resultId,omitempty"`
	Data     []uint  `json:"data"`
}

type SemanticTokensPartialResult struct {
	Data []uint `json:"data"`
}

type SemanticTokensDeltaRequest struct {
	Request
	Params SemanticTokensDeltaParams `json:"params"`
}

type SemanticTokensDeltaParams struct {
	WorkDoneProgressParams
	PartialResultParams
	TextDocument     TextDocumentIdentifier `json:"textDocument"`
	PreviousResultId string                 `json:"previousResultId"`
}

type SemanticTokensDeltaResponse struct {
	Response
	Result SemanticTokensDelta `json:"result"`
}

type SemanticTokensDelta struct {
	ResultId *string              `json:"resultId,omitempty"`
	Edits    []SemanticTokensEdit `json:"edits"`
}

type SemanticTokensEdit struct {
	Start       uint    `json:"start"`
	DeleteCount uint    `json:"deleteCount"`
	Data        *[]uint `json:"data,omitempty"`
}

type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit `json:"edits"`
}

func NewSemanticTokensOptions() *SemanticTokensOptions {
	var progress = true
	var full = true
	return &SemanticTokensOptions{
		WorkDoneProgressOptions: WorkDoneProgressOptions{
			WorkDoneProgress: &progress,
		},
		Legend: SemanticTokensLegend{
			TokenTypes:     generateOrderedArrayFromIndexedMap(TokenTypes),
			TokenModifiers: generateOrderedArrayFromIndexedMap(TokenModifiers),
		},
		Full: &full,
	}
}

func NewSemanticTokensResponse(req SemanticTokensRequest) SemanticTokensResponse {
	var tokens = []token{
		{
			sLn:       72,
			sCol:      1,
			eLn:       72,
			eCol:      2,
			tokenType: "Comment",
		},
		{
			sLn:       72,
			sCol:      4,
			eLn:       72,
			eCol:      7,
			tokenType: "View",
		},
		{
			sLn:       73,
			sCol:      1,
			eLn:       73,
			eCol:      5,
			tokenType: "StartTag",
		},

		{
			sLn:       79,
			sCol:      5,
			eLn:       79,
			eCol:      7,
			tokenType: "StartTag",
		},
		{
			sLn:       79,
			sCol:      8,
			eLn:       79,
			eCol:      20,
			tokenType: "Text",
		},
		{
			sLn:       79,
			sCol:      21,
			eLn:       79,
			eCol:      33,
			tokenType: "Code",
		},
		{
			sLn:       79,
			sCol:      34,
			eLn:       79,
			eCol:      37,
			tokenType: "EndTag",
		},

		{
			sLn:       81,
			sCol:      9,
			eLn:       81,
			eCol:      15,
			tokenType: "Component",
		},
		{
			sLn:       81,
			sCol:      19,
			eLn:       81,
			eCol:      23,
			tokenType: "Code",
		},
		{
			sLn:       81,
			sCol:      27,
			eLn:       81,
			eCol:      28,
			tokenType: "Component",
		},
		{
			sLn:       90,
			sCol:      1,
			eLn:       90,
			eCol:      2,
			tokenType: "Comment",
		},
	}
	var encodedTokens = encodeTokens(convertTokens(tokens))
	return SemanticTokensResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Result: SemanticTokens{
			Data: encodedTokens,
		},
	}
}

type token struct {
	sLn       int
	sCol      int
	eLn       int
	eCol      int
	tokenType string
}

type deltaToken struct {
	line   int
	char   int
	length int
	index  int
	flags  int
}

type encodedToken []uint

func convertTokens(tokens []token) []deltaToken {
	var dTs = []deltaToken{}
	for _, t := range tokens {
		dTs = append(dTs, deltaToken{
			line:   t.sLn - 1,
			char:   t.sCol - 1,
			length: t.eCol - t.sCol + 1,
			index:  TokenTypes[TokenSyntax[t.tokenType]],
			flags:  0,
		})
	}
	return dTs
}

func encodeTokens(tokens []deltaToken) []uint {
	var encoded encodedToken
	for i, t := range tokens {
		if i-1 >= 0 {
			if t.line == tokens[i-1].line {
				t.line = 0
				t.char = t.char - tokens[i-1].char
			} else {
				t.line = t.line - tokens[i-1].line
			}
		}
		encoded = append(encoded, encodeToken(t)...)
	}
	return encoded
}

func encodeToken(token deltaToken) []uint {
	var encoded encodedToken
	encoded = append(encoded, uint(token.line))
	encoded = append(encoded, uint(token.char))
	encoded = append(encoded, uint(token.length))
	encoded = append(encoded, uint(token.index))
	encoded = append(encoded, uint(token.flags))
	return encoded
}
