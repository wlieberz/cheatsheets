package main

import (
	"fmt"
	"os/exec"
)

func execScript(script string) {

	var formattedScript string
	formattedScript = fmt.Sprintf("./%v | tee ./output/%v_output.json", script, script)

	output, err := exec.Command("/bin/bash", "-c", formattedScript).Output()

	if err != nil {
		fmt.Printf("[ERROR] %s", err)
	}

	if output != nil {
		fmt.Printf("%s", output)
	}
}

func main() {

	scriptsToRun := make([]string, 3)
	scriptsToRun[0] = "task-1.sh"
	scriptsToRun[1] = "task-2.sh"
	scriptsToRun[2] = "task-3.sh"

	for _, script := range scriptsToRun {
		execScript(script)
	}

}
