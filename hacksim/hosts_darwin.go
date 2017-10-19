package main

import "path"

func getHostsPath() string {
	return path.Join("/", "etc", "hosts")
}
