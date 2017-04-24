package splat

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	ini "gopkg.in/ini.v1"
)

// Command represents the structure of a runnable Command
type Command struct {
	runner string
	args   []string
}

//CmdLookUp executes a lookup on a file or an eviroment variable
func CmdLookUp(cmd Command) (value string, err error) {
	if cmd.args[0] == "ENV" {
		value, err = lookUpEnv(cmd.args)
	} else {
		value, err = lookUpFile(cmd.args)
	}
	return value, err
}

func lookUpFile(args []string) (value string, err error) {
	absPath, _ := filepath.Abs(args[0])
	if _, err = os.Stat(absPath); os.IsNotExist(err) {
		return value, err
	}
	cfg, err := ini.Load(absPath)
	if err != nil {
		return value, err
	}
	if cfg.Section("").HasKey(args[1]) {
		value = cfg.Section("").Key(args[1]).String()
	} else {
		err = fmt.Errorf("Cannot find key %s in file %s", args[1], absPath)
		return value, err
	}
	return value, nil
}

func lookUpEnv(args []string) (value string, err error) {
	value, present := os.LookupEnv(args[1])
	if present == false {
		return value, fmt.Errorf("Cannot find variable %s in ENV", args[1])
	}
	return
}

// CmdFileContent gets the contents of a file
func CmdFileContent(cmd Command) (value string, err error) {
	absPath, _ := filepath.Abs(cmd.args[0])
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return value, err
	}
	value = string(content)
	return value, nil
}

// CmdRun executes an abitrary command and gets the STDOUT result
func CmdRun(cmd Command) (value string, err error) {
	return
}

// CmdCertificate generates a certificate based on the parameters associated in the argument
func CmdCertificate(cmd Command) (value string, err error) {
	return
}
