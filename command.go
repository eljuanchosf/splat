package splat

import (
	"fmt"
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
		err = fmt.Errorf("Cannot find key %s in file %s", args[1], args[0])
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
