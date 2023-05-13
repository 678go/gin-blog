package main

import "gin-blog/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		return
	}
}
