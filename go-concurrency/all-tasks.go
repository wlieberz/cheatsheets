package main

import (
	"fmt"
	"os/exec"
	"sync"
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

	var wg sync.WaitGroup

	for _, script := range scriptsToRun {
		wg.Add(1)

		// Need to copy the `script` var each time through the loop
		// so the go routine closure gets it's own copy of the var.
		// For more info, please see:
		// https://go.dev/doc/faq#closures_and_goroutines
		script := script

		go func() {
			defer wg.Done()
			execScript(script)
		}()
	}
	wg.Wait()

}
