package main

import (
	"os"
	"os/exec"
)

var _workspace = "/Documents/Workspace"
var _path = "/PortableApps/msys64/mingw64/bin"

func main() {
	// Setup Path
	os.Setenv("PATH", _path)

	// If a file or folder is provided, open that
	if len(os.Args) >= 2 {
		_workspace = os.Args[1]
	}

	// Launch VSCode
	cmd := exec.Command("/PortableApps/Code/Code.exe", _workspace)
	cmd.Start()
}
