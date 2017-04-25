package splat

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	ini "gopkg.in/ini.v1"
)

//CmdLookUp executes a lookup on a file or an eviroment variable
func CmdLookUp(cmdArgs []cmdArg) (value string, err error) {
	if cmdArgs[0].value == "ENV" {
		value, err = lookUpEnv(cmdArgs)
	} else {
		if cmdArgs[0].isFile == true {
			value, err = lookUpFile(cmdArgs)
		} else {
			err = errors.New("invalid lookup command")
		}
	}
	return
}

func lookUpFile(args []cmdArg) (value string, err error) {
	cfg, err := ini.Load(args[0].value)
	if err != nil {
		return value, err
	}
	if cfg.Section("").HasKey(args[1].value) {
		value = cfg.Section("").Key(args[1].value).String()
	} else {
		err = fmt.Errorf("Cannot find key %s in file %s", args[1].value, args[0].value)
		return value, err
	}
	return
}

func lookUpEnv(args []cmdArg) (value string, err error) {
	value, present := os.LookupEnv(args[1].value)
	if present == false {
		return value, fmt.Errorf("Cannot find variable %s in ENV", args[1].value)
	}
	return
}

// CmdFileContent gets the contents of a file
func CmdFileContent(cmdArgs []cmdArg) (value string, err error) {
	content, err := ioutil.ReadFile(cmdArgs[0].value)
	if err != nil {
		return value, err
	}
	value = string(content)
	return
}

// CmdRun executes an abitrary command and gets the STDOUT result
func CmdRun(cmdArgs []cmdArg) (value string, err error) {
	var args []string
	for _, arg := range cmdArgs[1:] {
		args = append(args, arg.value)
	}
	cmdOut, err := exec.Command(cmdArgs[0].value, args...).Output()
	if err != nil {
		return value, err
	}
	value = strings.TrimRight(string(cmdOut), "\n")
	return
}

// CmdCertificate generates a certificate based on the parameters associated in the argument
func CmdCertificate(cmd Command) (value string, err error) {
	return
}

func concatArgs(cmdArgs []cmdArg) (argLine string) {
	for _, arg := range cmdArgs {
		argLine += arg.value + " "
	}
	argLine = strings.TrimSpace(argLine)
	return
}
