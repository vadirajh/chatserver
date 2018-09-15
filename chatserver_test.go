package main 
import (
    "testing"
    "log"
    "os/exec"
)
func TestMain(t *testing.T) {
    out, err := exec.Command("./concurrent.sh").Output()
    if err != nil {
        log.Fatal(err)
    }
    t.Log(out)
} 
