package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type jsonFile struct {
	Application     string            `json:"application"`
	ApplicationArgs string            `json:"applicationArgs"`
	Environment     map[string]string `json:"environment"`
}

var err error
var settings jsonFile

func main() {

	str, err := ioutil.ReadFile("settings.json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal([]byte(str), &settings)

	fmt.Println("APP:", settings.Application)
	fmt.Println("ARG:", settings.ApplicationArgs)
	for k, v := range settings.Environment {
		os.Setenv(k, v)
		fmt.Println("ENV:", k, "=", v)
	}

	args := make([]string, 0)
	if len(settings.ApplicationArgs) > 0 {
		args = append(args, settings.ApplicationArgs)
	}
	cmd := exec.Command(settings.Application, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
}
