package main

import (
	"os/Exec"
	"fmt"
	"log"
)

// fixed it in staging 
func main() {
     cmd := GetCommand()
     out, err := exec.Command(cmd[0], cmd[1:]...).Output()

    if err != nil {
        log.Fatal("ERROR:", err)
    }
    fmt.Printf("%s", out)
}
