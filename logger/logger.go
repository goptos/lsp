package logger

import (
	"log"
	"os"
)

func New(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logFile, "[goptoslsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
