package rpc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/goptos/lsp/lsp"
)

func splitFunction(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

type Connection struct {
	logger  *log.Logger
	reader  io.Reader
	writer  io.Writer
	handler func(*Connection, lsp.Method, []byte)
	scanner *bufio.Scanner
}

func NewConnection(l *log.Logger, r io.Reader, w io.Writer, h func(*Connection, lsp.Method, []byte)) *Connection {
	s := bufio.NewScanner(r)
	s.Split(splitFunction)
	return &Connection{
		logger:  l,
		reader:  r,
		writer:  w,
		handler: h,
		scanner: s,
	}
}

func (_self *Connection) Receive() {
	msg := _self.scanner.Bytes()
	method, content, err := _self.decodeMessage(msg)
	if err != nil {
		_self.logger.Printf("Couldn't decode message: %s\n", err)
	}
	_self.logger.Printf("Received << %s: %s", method, string(content))
	_self.handler(_self, lsp.Method(method), content)
}

func (_self *Connection) Send(res any) {
	var msg = []byte(_self.encodeMessage(res))
	_self.writer.Write(msg)
	_self.logger.Printf("Sent >> %v", strings.Replace(string(msg), "\r\n\r\n", " ", 1))
}

func (_self *Connection) GetLogger() *log.Logger {
	return _self.logger
}

func (_self *Connection) Log(s string, args ...interface{}) {
	_self.logger.Printf(" "+s, args...)
}

func (_self *Connection) Open() bool {
	return _self.scanner.Scan()
}

func (_self *Connection) Close() {
	_self.logger.Println("Ended")
	os.Exit(0)
}

func (_self *Connection) encodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func (*Connection) decodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, fmt.Errorf("did not find separator")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}
	var message lsp.Message
	err = json.Unmarshal(content[:contentLength], &message)
	if err != nil {
		return "", nil, err
	}
	return message.Method, content[:contentLength], nil
}
