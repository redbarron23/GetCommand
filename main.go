package main

import (
	"os/Exec"
)

func main() {
     cmd := GetCommand()
     out, err := exec.Command(cmd[0], cmd[1:]...).Output()
}
