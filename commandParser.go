package splat

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type cmdArg struct {
	order  int
	value  string
	isFile bool
}

// Command represents the structure of a runnable Command
type Command struct {
	runner string
	args   []cmdArg
}

func extractCommand(text string) (command Command, err error) {
	rgxp, err := regexp.Compile(`\(\(<(.*)>\)\)`)
	if err != nil {
		return command, err
	}
	r := rgxp.FindStringSubmatch(text)
	cmdString := strings.TrimSpace(r[1])
	command.runner = extractRunner(cmdString)
	command.args = extractRunnerArgs(cmdString)
	return
}

func extractRunner(cmd string) (runner string) {
	rgxp, _ := regexp.Compile(`(.*)\((.*)\)`)
	runner = rgxp.FindAllStringSubmatch(cmd, -1)[0][1]
	return
}

func extractRunnerArgs(cmd string) (args []cmdArg) {
	args = []cmdArg{}
	rgxp, _ := regexp.Compile(`(.*)\((.*)\)`)
	argString := rgxp.FindAllStringSubmatch(cmd, -1)[0][2]
	if len(argString) > 0 {
		unformattedArgs := strings.Split(argString, ",")
		args = formatArgs(unformattedArgs)
	}
	return
}

func formatArgs(unformattedArgs []string) (args []cmdArg) {
	for i, arg := range unformattedArgs {
		value := strings.TrimSpace(arg)
		isFile, value := isFile(value)
		args = append(args, cmdArg{order: i, value: value, isFile: isFile})
	}
	return
}

func isFile(value string) (bool, string) {
	file := getAbsPath(value)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false, value
	}
	return true, file
}

func getAbsPath(value string) string {
	file, _ := filepath.Abs(value)
	return file
}
