package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	chkDirs := map[string]bool{
		"/deps":          false,
		"/deps/caddy":    false,
		"/deps/meteor":   false,
		"/deps/node":     false,
		"/deps/mongodb":  false,
		"/deps/lair-api": false,
		"/deps/lair-app": false,
		"/db":            false,
		"/db/mongo":      false,
	}

	// Get list of current directory recursively
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		// Make directory path relative by stripping absolute
		relDir := strings.Replace(path, root, "", -1)
		// If the found directory is needed for lair, mark its existence
		if _, ok := chkDirs[relDir]; ok {
			chkDirs[relDir] = true
		}
		return nil
	})

	for path, exist := range chkDirs {
		fmt.Printf("Path: %s\nExists: %v\n", path, exist)
	}

	// Download missing dependencies

	// Untar dependencies

	// Remove tar files

	// Start up the app

}

/*
Lair start in directory, detect if these things exist, if they dont, download
the meteor tarball, node, mongodb, api-server, caddy, and the lair app itself

Lair app will be on github or something in tarball

Lair api-server will be on github or something in tarball

Mongodb https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1404-3.0.6.tgz


Ask for info, set env variables, launch it
Default config and be configurable itself

come with yaml file with some defaults

Configure mongodb for oplog, make config file for mongod

If config is just strings with default config that works too, write on startup if dont exist

caddyFile := `
thign from boop {
}
`
*/
