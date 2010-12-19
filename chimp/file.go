package chimp

import (
    "exec"
    "os"
    "fmt"
)

func forkWait(program string, argv []string, dir string) (err os.Error) {  
    program, err = exec.LookPath(program)
    if err != nil { return }
    
    runArgs := make([]string, len(argv) + 1)
    copy(runArgs[1:], argv)
    runArgs[0] = program
    
    fmt.Fprintf(os.Stderr,"Running: %s \n", runArgs)
    
    cmd, err := exec.Run(runArgs[0], runArgs, os.Environ(), dir, 
        exec.PassThrough, exec.PassThrough, exec.PassThrough)
    _, err = cmd.Wait(0)
    return
}

func deleteFile(file string) (err os.Error) {
    return os.RemoveAll(file)
}
