package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var _workspace = "/Documents/Workspace"
var _GitPath string
var _FinalPath string

func main() {
	// Set Absolute (`C:\PortableApps`) paths

	_GitPath, err := filepath.Abs("/PortableApps/Git4Win/bin")
	_msysPath, err := filepath.Abs("/PortableApps/msys64/mingw64/bin")
	_GoPath, err := filepath.Abs("/PortableApps/msys64/go")
	if err != nil {
		log.Fatal(err)
	}

	// Set gopath
	os.Setenv("GOPATH", _GoPath)
	// Set PATH
	_Path := os.Getenv("PATH")
	lastPath := []string{_Path, _GitPath, _GoPath + "/bin", _msysPath}
	_FinalPath = strings.Join(lastPath, ";")

	os.Setenv("PATH", _FinalPath)

	// If a file or folder is provided, open that
	if len(os.Args) >= 2 {
		_workspace = os.Args[1]
	}

	// Launch VSCode
	cmd := exec.Command("/PortableApps/Code/Code.exe", _workspace)
	cmd.Start()
}
