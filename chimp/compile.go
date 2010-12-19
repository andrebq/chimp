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
}

func (ct *CompileTask) Run(tl *TaskLogger) (err os.Error) {

    objFileExt, compiler, _ := discoverGotools()
    
    tl.Info("Compiling %s with %s", ct.Files, compiler)
    params := make([]string, len(ct.Files) + 2)
    params[0] = "-o"
    params[1] = "_go_." + objFileExt
    copy(params[2:], ct.Files)
    err = forkWait(compiler, params, ".")
    if err != nil {
        tl.Error("Unable to compile, reason: %s", err)
        return
    }
    
    return
}

type LinkTask struct {
    Ct *CompileTask
    Includes []string
    Target string
}

func (lt *LinkTask) Run(tl *TaskLogger) (err os.Error) {
    tl.Info("Starting linking process")
    tl.Info("Compile task files: %s", lt.Ct.Files)
    objFileExt, _, linker := discoverGotools()

    var paramLinker []string
    if lt.Includes == nil {
        paramLinker = make([]string, 3)
        paramLinker[0] = "-o"
        paramLinker[1] = lt.Target
        paramLinker[2] = "_go_." + objFileExt
    } else {
        paramLinker = make([]string, 3 + len(lt.Includes))
        paramLinker[0] = "-o"
        paramLinker[1] = lt.Target
        copy(paramLinker[2:], lt.Includes)
        paramLinker = append(paramLinker, "_go_." + objFileExt)
    }
    
    tl.Info("Linking with params: %s", paramLinker)
    err = forkWait(linker, paramLinker, ".")
    if err != nil {
        tl.Error("Unable to link, reason: %s", err)
    }
    return
}

func (lt *LinkTask) GetObjectFiles() []string {
    objFileExt, _, _ := discoverGotools()
    
    objFiles := make([]string, 1)
    objFiles[0] = "_go_." + objFileExt
    
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
