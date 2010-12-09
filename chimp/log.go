package chimp

import (
    "log"
    "os"
    "io"
)

var ( 
    Stderr = os.Stderr
)

type TaskLogger struct {
    ln *log.Logger
}

func newTaskLogger(ln *log.Logger) *TaskLogger {
    tl := &TaskLogger{ln}
    tl.ln = ln
    return tl
}

func (tl *TaskLogger) Error(fmt string, v ...interface{}) {
    tl.ln.Printf("[ERROR] - " + fmt + "\n", v...)
}

func (tl *TaskLogger) Info(fmt string, v ...interface{}) {
    tl.ln.Printf("[INFO] - " + fmt + "\n", v...)
}

/**
Create a new loger which points to Stderr
*/
func NewLog(name string) *TaskLogger {
    ln := log.New(Stderr, "[" + name + "] ", log.Ldate | log.Ltime)
    return newTaskLogger(ln)
}

/**
Create a new loger which points to file
*/
func NewFileLog(name string, file io.Writer) *TaskLogger {
    ln := log.New(file, "[" + name + "] ", log.Ldate | log.Ltime)
    return newTaskLogger(ln)
}
