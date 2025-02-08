package main

import "fmt"

func hello() {
	var message = `Welcome to workspace-runner!

Usage:
workspace-runner run <workspace-name>
workspace-runner list
workspace-runner server`

	fmt.Println(message)

}
