package chimp

import (
    "os"
)

var (
    Goarch = os.Getenv("GOARCH")
    Goos = os.Getenv("GOOS")
    Gobin = os.Getenv("GOBIN")
)

type CompileTask struct {
    Files []string
    Target string
}

func (ct *CompileTask) Run(tl *TaskLogger) {

    objFileExt, compiler, linker := discoverGotools()
    
    objFiles := make([]string, len(ct.Files))
    
    for idx, file := range ct.Files {
        objFiles[idx] = file[:len(file) - 3] + "." + objFileExt
    }
    
    tl.Info("Compiling %s with %s", ct.Files, compiler)
    err := forkWait(compiler, ct.Files, ".")
    if err != nil {
        tl.Error("Unable to compile, reason: %s", err)
        return
    }
    
    paramLinker := make([]string, len(ct.Files) + 2)
    paramLinker[0] = "-o"
    paramLinker[1] = ct.Target
    copy(paramLinker[2:], objFiles)
    
    tl.Info("Linking %s to %s with %s", objFiles, paramLinker[1], linker)
    err = forkWait(linker, paramLinker, ".")
    if err != nil {
        tl.Error("Unable to link, reason: %s", err)
    }
}

func (ct *CompileTask) GetObjectFiles() []string {
    objFileExt, _, _ := discoverGotools()
    
    objFiles := make([]string, len(ct.Files))
    
    for idx, file := range ct.Files {
        objFiles[idx] = file[:len(file) - 3] + "." + objFileExt
    }
    
    return objFiles
}

func discoverGotools() (string, string, string) {
    switch Goarch {
        case "386":
            return "8", "8g", "8l"
            
        case "amd64":
            return "6", "6g", "6l"
    }
    panic("Unable to find Go Compiler")
}
