package scripts

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"ultimateWorkSpace/internal"
)

func ExecScript(name string) {
	fmt.Printf("Executing script: %s ...\n\n", name)
	scriptPath := internal.SCRIPTS_DIR + name + ".sh"
	if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Script %s does not exist", name)
		return
	}

	out, err := exec.Command("/bin/sh", scriptPath).CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %v\n", err)
		return
	}
	fmt.Printf("Script output:\n\n%s", string(out))
}
