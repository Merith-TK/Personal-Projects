package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Set universal variables
var workspace = "/Documents/Workspace"
var err error
var path = "/PortableApps/CommonFiles/bin"

func main() {

	// If there is a specifed folder, use that was workspace
	if len(os.Args) >= 2 {
		workspace = os.Args[1]
	}

	// Get direct path for windows to `workspace`
	workspace, err = filepath.Abs(workspace)
	if err != nil {
		log.Fatal(err)
	}
	// Get direct path for windows to `path`
	path, err = filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	// Clean workspace path
	workspace = strings.Replace(workspace, "\\PortableApps\\msys64", "", -1)

	//Set PATH
	os.Setenv("PATH", path)
	
	//Execute Code with workspace
	fmt.Println("Starting Code Portable:", workspace)
	cmd := exec.Command("/PortableApps/Code/Code.exe", workspace)
	cmd.Start()
}
