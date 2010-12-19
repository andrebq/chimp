package chimp

import ( 
    "os" 
)

type CleanFiles struct {
    Files []string
}

func (cl *CleanFiles) Run(tl *TaskLogger) (err os.Error) {
    tl.Info("Removing files: %s", cl.Files)
    for _, f := range cl.Files {
        err := deleteFile(f)
        if err != nil {
            tl.Info("Unable to remove: %s reason: %s", f, err)
            return err
        }
    }
    return
}
