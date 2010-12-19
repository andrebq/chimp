package chimp

import (
    "os"
)

type Task interface {
    Run(*TaskLogger) os.Error;
}

type Operation struct {
    Fn func() os.Error
    Name string
}

func (op *Operation) Run(l *TaskLogger) (err os.Error) {
    l.Info("Running operation: %s", op.Name)
    err = op.Fn()
    if err != nil {
        l.Error("Error: %s", err)
    }
    return
}

type Chain struct {
    Tasks []Task
}

func (c *Chain) Run(l *TaskLogger) (err os.Error) {
    for _, t := range c.Tasks {
        err = t.Run(l)
        if err != nil {
            return
        }
    }
    return
}

func NewChain(tasks ...Task) (*Chain) {
    return &Chain{Tasks:tasks}
}
