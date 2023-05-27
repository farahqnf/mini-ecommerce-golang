package helpers

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
