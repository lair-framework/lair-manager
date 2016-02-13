package main

import (
	_ "crypto/sha512"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/lair-framework/lair-manager/helpers"
)

func main() {

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.CheckDirLayout(root)
	if err != nil {
		log.Fatal(err)
	}

	// Download them all
	err = helpers.DownloadFile("node")
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.DownloadFile("mongodb")
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.DownloadFile("api-server")
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.DownloadFile("caddy")
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.DownloadFile("lair-app")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(helpers.GetMissing())
	// Extract them all
	os.MkdirAll("./deps/mongodb", 0777)
	cmd := "tar"
	args := []string{"-zxvf", "mongodb-linux-x86_64-ubuntu1404-3.0.6.tgz", "-C", "./deps/mongodb", "--strip-components=1"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Extraction complete")
	// Delete tar files

	// Start up app

	// Download missing dependencies
	/*
		for path, exist := range chkDirs {
			fmt.Printf("")
		}
	*/
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
