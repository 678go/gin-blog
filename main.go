package main

import "gin-blog/cmd"
import _ "gin-blog/docs"

// @title gin-blog project
// @version 1.0
// @description this is gin-blog server.
// @host 127.0.0.1:8090
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
