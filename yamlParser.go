package splat

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

// Doc represents the text structure of a Yaml document
type Doc struct {
	Lines []Line
}

// Line represents a line in a Yaml file
type Line struct {
	Order      int
	Key        string
	Value      string
	Indent     int
	HasCommand bool
}

// Parse parses a YAML document as text
func Parse(doc string) (Doc, error) {
	yaml := Doc{}
	scanner := bufio.NewScanner(strings.NewReader(doc))
	numLine := 0
	for scanner.Scan() {
		key, value := splitLine(scanner.Text())
		hasCommand := false
		if regexp.MustCompile(`\(\(<.*>\)\)`).MatchString(value) == true {
			hasCommand = true
		}

		indent := countLeadingSpace(key)
		docLine := Line{Order: numLine, Key: strings.TrimSpace(key), Value: strings.TrimSpace(value), Indent: indent, HasCommand: hasCommand}
		yaml.Lines = append(yaml.Lines, docLine)
		numLine++
	}
	return yaml, nil
}

// ParseFile parses the YAML file into a text representation
func ParseFile(fileName string) (Doc, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Doc{}, err
	}
	yaml, _ := Parse(string(bytes))
	return yaml, nil
}

func splitLine(line string) (key string, value string) {
	split := strings.SplitN(line, ":", 2)
	key = split[0]
	value = ""
	if len(split) == 2 {
		value = split[1]
	}
	return key, value
}

func countLeadingSpace(line string) int {
	i := 0
	for _, runeValue := range line {
		if runeValue == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
