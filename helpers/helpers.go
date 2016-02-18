package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type depInfo struct {
	Name       string
	URLLinux32 string
	URLLinux64 string
	URLOSX     string
	Directory  string
	Exist      bool
}

type dependencies []depInfo

//the meteor tarball, node, mongodb, api-server, caddy, and the lair app itself
var downloadLocNix64 = map[string]string{
	"meteor":   "",
	"node":     "https://nodejs.org/dist/v4.2.6/node-v4.2.6-linux-x64.tar.xz",
	"mongodb":  "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1404-3.0.6.tgz",
	"lair-api": "https://github.com/lair-framework/api-server/releases/download/v1.1.0/api-server_linux_amd64",
	"caddy":    "https://github.com/mholt/caddy/releases/download/v0.8.1/caddy_linux_amd64.tar.gz",
	"lair-app": "https://github.com/lair-framework/lair/releases/download/v2.0.4/lair-v2.0.4-linux-amd64.tar.gz",
}

var chkDirs = map[string]bool{
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

// CheckDirLayout sets the values in chkDir to true if the director Exists.
// if the directory exists, it is assumed that the dependency is installed
// and will not download or install the respective dependency
func CheckDirLayout(root string) error {

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
		fmt.Printf("Path: .%s\n", path)
		fmt.Printf("Exists: ")
		if exist {
			color.Green("true")
		} else {
			color.Red("false")
		}
	}
	return nil
}

// GetMissing returns a list of the missing dependencies that need to be downloaded
// and installed.
func GetMissing() (keys []string) {
	for i, j := range chkDirs {
		if j == false {
			keys = append(keys, i)
		}
	}
	return keys
}

// IsMissing returns true or false if the depdency already exists and is
// installed.
func IsMissing(dependency string) (exists bool) {
	return chkDirs[dependency]
}

// DownloadFile retrieves the file at the URL and saves to current directory
func DownloadFile(dep string) (err error) {

	// Get the filename to create from the trailing path info in url
	segments := strings.Split(downloadLocNix64[dep], "/")
	filepath := segments[len(segments)-1]
	//fmt.Println(downloadLocNix64[dep])

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(downloadLocNix64[dep])
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
