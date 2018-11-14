package logger_go

import (
	"log"
	"os"
)

type Model struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func New(i, e, t, w string) *Model {
	m := &Model{}
	fileTrace, err := os.OpenFile(t+"/trace.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileTrace file: %v", err)
	}

	fileInfo, err := os.OpenFile(i+"/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileInfo file: %v", err)
	}

	fileWarning, err := os.OpenFile(w+"/warning.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileWarning file: %v", err)
	}

	fileError, err := os.OpenFile(e+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening fileError file: %v", err)
	}

	m.Trace = log.New(fileTrace,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Llongfile)

	m.Info = log.New(fileInfo,
		"INFO: ",
		log.Ldate|log.Ltime|log.Llongfile)

	m.Warning = log.New(fileWarning,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Llongfile)

	m.Error = log.New(fileError,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile)

	return m
}
