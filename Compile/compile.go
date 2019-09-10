package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	_icon     = "icon.ico"
	_manifest = "manifest.xml"
)

var iconArg1, iconArg2, manifestArg1, manifestArg2 string
var outbuf, errbuf bytes.Buffer

func main() {
	// Check for icon.ico
	if _, err := os.Stat(_icon); err == nil {
		fmt.Println("LOG: icon.ico found")
		iconArg1 = "-ico"
		iconArg2 = "icon.ico"
	}

	// Check for manifest.xml
	if _, err := os.Stat(_manifest); os.IsNotExist(err) {
		fmt.Println("ERROR: manifest.xml is not found")
		os.Exit(25)
	}
	fmt.Println("LOG: manifest.xml found")
	manifestArg1 = "-manifest"
	manifestArg2 = "manifest.xml"
	fmt.Println("COMMAND:", "rsrc.exe", iconArg1, iconArg2, manifestArg1, manifestArg2)

	CMDrsrc := exec.Command("rsrc.exe", iconArg1, iconArg2, manifestArg1, manifestArg2)
	CMDrsrc.Stderr = os.Stderr
	CMDrsrc.Stdout = os.Stdout
	rsrcErr := CMDrsrc.Run()
	if rsrcErr != nil {
		fmt.Println("rsrc.exe failed", rsrcErr)
	}

	CMDbuild := exec.Command("go.exe", "build")
	buildErr := CMDbuild.Start()
	if buildErr != nil {
		fmt.Println("buildErr")
		log.Fatal(buildErr)
	}

}
