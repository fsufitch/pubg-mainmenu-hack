package main

import (
	"os"
	"path"
)

func getHostsPath() string {
	root := os.Getenv("SystemRoot")
	return path.Join(root, "System32", "drivers", "etc", "hosts")
}
