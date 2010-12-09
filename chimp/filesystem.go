package chimp

type CleanFiles struct {
    Files []string
}

func (cl *CleanFiles) Run(tl *TaskLogger) {
    tl.Info("Removing files: %s", cl.Files)
    for _, f := range cl.Files {
        err := deleteFile(f)
        if err != nil {
            tl.Info("Unable to remove: %s reason: %s", f, err)
        }
    }
}
