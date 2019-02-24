package logger_go

import (
	"log"
	"os"
	"path"
	"sync"
)

var (
	once sync.Once
)

type Model struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func New(infoDir, errDir, traceDir, warnDir string) *Model {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	once.Do(func() {
		createDirectories(infoDir, errDir, traceDir, warnDir)
	})

	m := &Model{}
	fileTrace, err := createFile(traceDir, "trace.log")
	if err != nil {
		log.Fatalf("error opening fileTrace file: %v", err)
	}

	fileInfo, err := createFile(infoDir, "info.log")
	if err != nil {
		log.Fatalf("error opening fileInfo file: %v", err)
	}

	fileWarning, err := createFile(warnDir, "warning.log")
	if err != nil {
		log.Fatalf("error opening fileWarning file: %v", err)
	}

	fileError, err := createFile(errDir, "error.log")
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

// createDirectories revisa y crea (de ser necesario) los directorios
// de los log
func createDirectories(i, e, t, w string) {
	checkDirErr(checkDirectory(i))
	checkDirErr(checkDirectory(e))
	checkDirErr(checkDirectory(t))
	checkDirErr(checkDirectory(w))
}

// checkDirectory crea el directorio si no existe
func checkDirectory(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		err = os.MkdirAll(d, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// checkDirErr controla el error y termina el proceso
func checkDirErr(err error) {
	if err != nil {
		log.Fatalf("no se pudo crear el directorio: %v", err)
	}
}

// createFile crea el archivo de log
func createFile(d, s string) (*os.File, error) {
	return os.OpenFile(
		path.Join(d, s),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
}
