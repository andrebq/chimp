package chimp

import (
    "os"
)

func SetGoEnv(arch, goos string) (err os.Error) {
    err = os.Setenv("GOARCH", arch)
    if err == nil {
        err = os.Setenv("GOOS", goos)
    }
    return
}
