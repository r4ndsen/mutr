package main

import (
	"fmt"
	"os"
	"os/exec"
)

func toggleMute() error {
	// Command to get the current mute status
	getMuteStatus := exec.Command("osascript", "-e", "output muted of (get volume settings)")

	// Run the command and capture the output
	output, err := getMuteStatus.Output()
	if err != nil {
		return err
	}

	// Check if the output is "true" (currently muted)
	isMuted := string(output) == "true\n"

	var toggleCommand *exec.Cmd
	if isMuted {
		// If muted, unmute
		toggleCommand = exec.Command("osascript", "-e", "set volume without output muted")
	} else {
		// If not muted, mute
		toggleCommand = exec.Command("osascript", "-e", "set volume with output muted")
	}

	return toggleCommand.Run()
}

func main() {
	if err := toggleMute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
