package main

import (
	"regexp"
	"strings"
)

// Command represents the structure of a runnable Command
type Command struct {
	runner string
	args   []string
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

func runCommand(command Command) (result string) {
	return
}

func extractRunner(cmd string) (runner string) {
	rgxp, _ := regexp.Compile(`(.*)\((.*)\)`)
	runner = rgxp.FindAllStringSubmatch(cmd, -1)[0][1]
	return
}

func extractRunnerArgs(cmd string) (args []string) {
	args = []string{}
	rgxp, _ := regexp.Compile(`(.*)\((.*)\)`)
	argString := rgxp.FindAllStringSubmatch(cmd, -1)[0][2]
	if len(argString) > 0 {
		unformattedArgs := strings.Split(argString, ",")
		args = formatArgs(unformattedArgs)
	}
	return
}

func formatArgs(unformattedArgs []string) (args []string) {
	for _, arg := range unformattedArgs {
		args = append(args, strings.TrimSpace(arg))
	}
	return
}
