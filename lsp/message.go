package lsp

type Method string

const (
	Cancel   Method = "$/cancelRequest"
	Shutdown Method = "shutdown"
)

type ErrorCode int

const (
	ParseError                     ErrorCode = -32700
	InvalidRequest                 ErrorCode = -32600
	MethodNotFound                 ErrorCode = -32601
	InvalidParams                  ErrorCode = -32602
	InternalError                  ErrorCode = -32603
	jsonrpcReservedErrorRangeStart ErrorCode = -32099
	ServerNotInitialized           ErrorCode = -32002
	UnknownErrorCode               ErrorCode = -32001
	jsonrpcReservedErrorRangeEnd   ErrorCode = -32000
	lspReservedErrorRangeStart     ErrorCode = -32899
	RequestFailed                  ErrorCode = -32803
	ServerCancelled                ErrorCode = -32802
	ContentModified                ErrorCode = -32801
	RequestCancelled               ErrorCode = -32800
	lspReservedErrorRangeEnd       ErrorCode = -32800
)

type Message struct {
	Method string `json:"method"`
	Rpc    string `json:"jsonrpc"`
}

type Request struct {
	Message
	Id int `json:"id"`

	// Params
}

type Response struct {
	Rpc string `json:"jsonrpc"`
	Id  *int   `json:"id,omitempty"`

	// Result
	// Error
}

type Notification struct {
	Message
}

type CancelRequest struct {
	Request
	Params CancelParams `json:"params"`
}

type CancelParams struct {
	Id int `json:"id"`
}

type CancelResponse struct {
	Response
	Error ResponseError `json:"error"`
}

type ResponseError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Data    *LSPAny   `json:"data,omitempty"`
}

type LSPAny struct {
}

func NewCancelResponse(req CancelRequest) CancelResponse {
	return CancelResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &req.Id,
		},
		Error: ResponseError{
			Code:    MethodNotFound,
			Message: "we are single threaded and cannot support cancellation requests",
		},
	}
}
